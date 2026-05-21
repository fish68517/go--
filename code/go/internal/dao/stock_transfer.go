package dao

import (
	"errors"
	"fmt"
	"time"

	"github.com/mwqnice/oh-admin/internal/dto"
	"github.com/mwqnice/oh-admin/internal/model"
	"github.com/mwqnice/oh-admin/pkg/convert"
	"golang.org/x/net/context"
	"gorm.io/gorm"
)

func (d *Dao) GetStockTransferInList(ctx context.Context, params *dto.GetStockTransferListRequest) ([]*model.StockTransfer, int64, error) {
	stockIn := model.StockTransfer{}
	return stockIn.List(ctx, d.engine, params)
}
func (d *Dao) GetStockTransferItemInfo(ctx context.Context, drug *model.StockTransferItem) ([]*model.StockTransferItem, error) {
	return drug.Get(ctx, d.engine)
}
func (d *Dao) CreatetSockTransfer(ctx context.Context, pres *dto.StockTransferCreateDTO) error {
	return d.SockTransferCreate(ctx, d.engine, pres)
}
func (d *Dao) SockTransferCreate(ctx context.Context, db *gorm.DB, params *dto.StockTransferCreateDTO) error {

	transfer := &model.StockTransfer{
		OrderNo:         params.OrderNo,
		FromWarehouseID: params.FromWarehouseID,
		ToWarehouseID:   params.ToWarehouseID,
		OperatorID:      params.OperatorId,
		Remark:          params.Remark,
		CreateTime:      time.Now(),
		UpdateTime:      time.Now(),
	}
	var items []dto.StockTransferItems
	items = params.Items
	tx := db.WithContext(ctx).Begin()
	if err := tx.Create(&transfer).Error; err != nil {
		tx.Rollback()
		return err
	}

	for i := range items {
		items[i].TransferID = string(transfer.ID)
	}

	if err := tx.Create(&items).Error; err != nil {
		tx.Rollback()
		return err
	}
	// 处理库存变更（核心逻辑：调出减、调入加）
	for _, item := range params.Items {
		productID := item.ProductID
		quantity := item.Quantity

		// 6.1 调出仓库库存减少
		var fromInventory model.Inventory
		if err := tx.First(&fromInventory, "warehouse_id = ? AND product_id = ?",
			params.FromWarehouseID, productID).Error; err != nil {
			tx.Rollback()
			return fmt.Errorf("调出仓库商品库存不存在: 商品ID=%d, 错误=%w", productID, err)
		}
		if convert.Float64(fromInventory.Quantity) < convert.Float64(quantity) {
			tx.Rollback()
			return fmt.Errorf("调出仓库商品库存不足: 商品ID=%d, 现有=%d, 需调拨=%d",
				productID, fromInventory.Quantity, quantity)
		}
		if err := tx.Model(&model.Inventory{}).
			Where("id = ?", fromInventory.ID).
			Update("quantity", gorm.Expr("quantity - ?", quantity)).Error; err != nil {
			tx.Rollback()
			return fmt.Errorf("更新调出仓库库存失败: 商品ID=%d, 错误=%w", productID, err)
		}

		// 6.2 调入仓库库存增加（存在则更新，不存在则创建）
		var toInventory model.Inventory
		result := tx.First(&toInventory, "warehouse_id = ? AND product_id = ?",
			params.ToWarehouseID, productID)
		if result.Error != nil && !errors.Is(result.Error, gorm.ErrRecordNotFound) {
			tx.Rollback()
			return fmt.Errorf("查询调入仓库库存失败: 商品ID=%d, 错误=%w", productID, result.Error)
		}

		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			// 库存记录不存在，创建新记录
			newInventory := &model.Inventory{
				WarehouseID: params.ToWarehouseID,
				ProductID:   productID,
				ProductName: item.ProductName,
				Quantity:    quantity,
			}
			if err := tx.Create(newInventory).Error; err != nil {
				tx.Rollback()
				return fmt.Errorf("创建调入仓库库存记录失败: 商品ID=%d, 错误=%w", productID, err)
			}
		} else {
			// 库存记录存在，更新数量
			if err := tx.Model(&toInventory).
				Update("quantity", gorm.Expr("quantity + ?", quantity)).Error; err != nil {
				tx.Rollback()
				return fmt.Errorf("更新调入仓库库存失败: 商品ID=%d, 错误=%w", productID, err)
			}
		}
	}
	// 提交事务
	tx.Commit()

	return nil
}

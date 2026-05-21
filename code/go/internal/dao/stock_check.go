package dao

import (
	"context"
	"fmt"
	"time"

	"github.com/mwqnice/oh-admin/internal/dto"
	"github.com/mwqnice/oh-admin/internal/model"
	"github.com/mwqnice/oh-admin/pkg/convert"
	"gorm.io/gorm"
)

func (d *Dao) GetStockCheckItemInfo(ctx context.Context, drug *model.StockCheckItem) ([]*model.StockCheckItem, error) {
	return drug.Get(ctx, d.engine)
}
func (d *Dao) GetStockCheckList(ctx context.Context, params *dto.GetStockCheckListRequest) ([]*model.StockCheck, int64, error) {
	stockIn := model.StockCheck{}
	return stockIn.List(ctx, d.engine, params)
}
func (d *Dao) CreateStockCheck(ctx context.Context, pres *dto.StockCheckCreateDTO) error {
	return d.StockCheckCreate(ctx, d.engine, pres)
}
func (d *Dao) StockCheckCreate(ctx context.Context, db *gorm.DB, params *dto.StockCheckCreateDTO) error {

	stockIn := &model.StockCheck{
		OrderNo:     params.OrderNo,
		WarehouseID: params.WarehouseID,
		OperatorID:  params.OperatorID,
		CreateTime:  time.Now(),
		UpdateTime:  time.Now(),
		Remark:      params.Remark,
	}
	var items []dto.StockCheckItems
	items = params.Items
	tx := db.WithContext(ctx).Begin()
	// 插入调剂记录
	if err := tx.Create(&stockIn).Error; err != nil {
		tx.Rollback()
		return err
	}

	for i := range items {
		items[i].CheckID = convert.String(stockIn.ID)
	}

	if err := tx.Create(&items).Error; err != nil {
		tx.Rollback()
		return err
	}
	// 3. 批量更新库存（存在则更新，不存在则创建）
	for _, itemDTO := range params.Items {
		inventory := &model.Inventory{
			ProductID:   itemDTO.ProductID,
			WarehouseID: params.WarehouseID,
			Quantity:    itemDTO.ActualQuantity,
		}

		// 先尝试查询并锁定现有库存记录（避免并发问题）
		result := tx.Where("product_id = ? AND warehouse_id = ?", itemDTO.ProductID, params.WarehouseID).
			First(inventory)

		if result.Error != nil && result.Error != gorm.ErrRecordNotFound {
			return fmt.Errorf("查询库存记录失败: %w", result.Error)
		}

		// 更新库存数量（盘点后以实际盘点数量为准）
		inventory.Quantity = itemDTO.ActualQuantity
		if result.Error == gorm.ErrRecordNotFound {
			// 库存记录不存在，创建新记录
			if err := tx.Create(inventory).Error; err != nil {
				return fmt.Errorf("创建库存记录失败: %w", err)
			}
		} else {
			// 库存记录存在，更新数量
			if err := tx.Save(inventory).Error; err != nil {
				return fmt.Errorf("更新库存记录失败: %w", err)
			}
		}
	}

	// 提交事务
	tx.Commit()

	return nil
}

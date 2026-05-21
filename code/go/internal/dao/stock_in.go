package dao

import (
	"context"
	"time"

	"github.com/mwqnice/oh-admin/internal/dto"
	"github.com/mwqnice/oh-admin/internal/model"
	"github.com/mwqnice/oh-admin/pkg/convert"
	"gorm.io/gorm"
)

func (d *Dao) GetStockInItemInfo(ctx context.Context, drug *model.StockInItem) ([]*model.StockInItem, error) {
	return drug.Get(ctx, d.engine)
}
func (d *Dao) GetStockInList(ctx context.Context, params *dto.GetStockInListRequest) ([]*model.StockIn, int64, error) {
	stockIn := model.StockIn{}
	return stockIn.List(ctx, d.engine, params)
}
func (d *Dao) UpdateStockIn(ctx context.Context, pres *model.Prescription) error {
	return pres.Update(ctx, d.engine)
}

func (d *Dao) DeleteStockIn(ctx context.Context, id int64) error {
	prescription := model.Prescription{}
	return prescription.Delete(ctx, d.engine, int(id))
}
func (d *Dao) CreateStockIn(ctx context.Context, pres *dto.StockInCreateDTO) error {
	return d.StockInCreate(ctx, d.engine, pres)
}
func (d *Dao) StockInCreate(ctx context.Context, db *gorm.DB, params *dto.StockInCreateDTO) error {
	// 创建入库主记录
	stockIn := &model.StockIn{
		OrderNo:     params.OrderNo,
		WarehouseID: params.WarehouseId,
		OperatorId:  params.OperatorId,
		CreateTime:  time.Now(),
		UpdateTime:  time.Now(),
	}

	tx := db.WithContext(ctx).Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	// 创建入库单
	if err := tx.Create(&stockIn).Error; err != nil {
		tx.Rollback()
		return err
	}

	// 处理入库明细
	for i := range params.Items {
		item := &params.Items[i]
		item.StockInId = convert.String(stockIn.ID)

		// 查找现有库存
		inventory := model.Inventory{}
		if err := tx.Where("warehouse_id = ? AND product_id = ?",
			params.WarehouseId, item.ProductId).First(&inventory).Error; err == nil {

			// 更新现有库存
			newQty := inventory.Quantity + item.Quantity
			if err := tx.Model(&model.Inventory{}).
				Where("id = ?", inventory.ID).
				Update("quantity", newQty).Error; err != nil {
				tx.Rollback()
				return err
			}
		} else {
			// 创建新库存记录
			inventory = model.Inventory{
				ProductID:   item.ProductId,
				WarehouseID: params.WarehouseId,
				Quantity:    item.Quantity,
				ProductName: item.ProductName,
			}
			if err := tx.Create(&inventory).Error; err != nil {
				tx.Rollback()
				return err
			}
		}
	}

	// 批量创建入库明细
	if err := tx.Create(&params.Items).Error; err != nil {
		tx.Rollback()
		return err
	}

	// 提交事务
	tx.Commit()
	return nil
}

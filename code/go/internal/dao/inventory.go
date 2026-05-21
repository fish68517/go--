package dao

import (
	"context"
	"fmt"
	"math/rand"
	"time"

	"github.com/mwqnice/oh-admin/internal/dto"
	"github.com/mwqnice/oh-admin/internal/model"
	"github.com/mwqnice/oh-admin/pkg/convert"
	"gorm.io/gorm"
)

func (d *Dao) GetInventoryList(ctx context.Context, params *dto.GetInventoryListRequest) ([]*model.Inventory, int64, error) {
	link := model.Inventory{}
	return link.List(ctx, d.engine, params)
}
func (d *Dao) GetInventoryInfo(ctx context.Context, link *model.Inventory) (*model.Inventory, error) {
	return link.Get(ctx, d.engine)
}
func (d *Dao) CreateStockOut(ctx context.Context, pres *dto.InventoryOperationDTO) error {
	return d.StockOutCreate(ctx, d.engine, pres)
}
func (d *Dao) StockOutCreate(ctx context.Context, db *gorm.DB, params *dto.InventoryOperationDTO) error {
	date := time.Now().Format("20060102") // 日期部分
	randNum := rand.Intn(1000000)
	stockOut := &model.StockOut{
		OrderNo:     "pc" + fmt.Sprintf("ORDER-%s-%06d", date, randNum),
		WarehouseID: params.WarehouseID,
		OperatorID:  "1",
		CreateTime:  time.Now(),
		UpdateTime:  time.Now(),
		OutType:     "1",
		Remark:      "",
		CustomerID:  "",
	}

	tx := db.WithContext(ctx).Begin()
	if err := tx.Create(&stockOut).Error; err != nil {
		tx.Rollback()
		return err
	}
	stockOutItem := &model.StockOutItem{
		StockOutID:     convert.String(stockOut.ID),
		ProductID:      params.ProductID,
		Quantity:       params.Quantity,
		UnitPrice:      "",
		Amount:         "",
		BatchNo:        "",
		ProductionDate: "",
		ExpiryDate:     "",
		Remark:         "",
	}

	if err := tx.Create(&stockOutItem).Error; err != nil {
		tx.Rollback()
		return err
	}
	inventory := &model.Inventory{}
	result := tx.Where("product_id = ? AND warehouse_id = ?", params.ProductID, params.WarehouseID).
		Set("gorm:query_option", "FOR UPDATE"). // 悲观锁，防止并发更新
		First(inventory)

	if result.Error != nil && result.Error != gorm.ErrRecordNotFound {
		tx.Rollback()
		return fmt.Errorf("查询库存记录失败: %w", result.Error)
	}

	// 检查库存是否充足
	if result.Error == gorm.ErrRecordNotFound {
		tx.Rollback()
		return fmt.Errorf("产品 %s 在仓库 %s 中无库存记录", params.ProductID, params.WarehouseID)
	}

	if convert.Float64(inventory.Quantity) < convert.Float64(params.Quantity) {
		tx.Rollback()
		return fmt.Errorf("库存不足，当前库存: %d, 申请出库: %d", inventory.Quantity, params.Quantity)
	}

	// 更新库存数量
	inventory.Quantity = convert.String(convert.Float64(inventory.Quantity) - convert.Float64(params.Quantity))
	if err := tx.Save(inventory).Error; err != nil {
		tx.Rollback()
		return fmt.Errorf("更新库存记录失败: %w", err)
	}

	// 提交事务
	if err := tx.Commit().Error; err != nil {
		return fmt.Errorf("提交事务失败: %w", err)
	}

	// 提交事务
	tx.Commit()

	return nil
}

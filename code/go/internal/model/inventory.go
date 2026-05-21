package model

import (
	"context"
	"fmt"

	"github.com/mwqnice/oh-admin/internal/dto"
	"github.com/mwqnice/oh-admin/pkg/app"
	"gorm.io/gorm"
)

type Inventory struct {
	*Model
	ProductID   string `gorm:"column:product_id;not null;comment:商品ID;index:;uniqueIndex:uk_product_warehouse" json:"product_id"`
	ProductName string `gorm:"column:product_name;not null;comment:商品Name" json:"product_name"`
	WarehouseID string `gorm:"column:warehouse_id;not null;comment:仓库ID;index:;uniqueIndex:uk_product_warehouse" json:"warehouse_id"`
	Quantity    string `gorm:"column:quantity;type:decimal(12,2);default:0.00;not null;comment:库存数量" json:"quantity"`
}

// TableName 设置表名
func (m Inventory) TableName() string {
	return "inventory"
}
func (l *Inventory) List(ctx context.Context, db *gorm.DB, params *dto.GetInventoryListRequest) ([]*Inventory, int64, error) {
	db = db.WithContext(ctx).Table(l.TableName())
	if params.Name != "" {
		db = db.Where("product_name like ?", fmt.Sprintf("%%%s%%", params.Name))
	}
	var count int64
	db.Count(&count)
	var list []*Inventory
	if params.Page > 0 && params.PageSize > 0 {
		db = db.Offset(app.GetPageOffset(int(params.Page), int(params.PageSize))).Limit(int(params.PageSize))
	}
	if err := db.Order("id desc").Find(&list).Error; err != nil {
		return nil, 0, err
	}
	return list, count, nil
}
func (l *Inventory) Get(ctx context.Context, db *gorm.DB) (*Inventory, error) {
	db = db.WithContext(ctx).Table(l.TableName())
	fmt.Print(l.Model.ID)
	if l.Model != nil && l.Model.ID != 0 {
		db = db.Where("id = ? ", l.Model.ID)
	}
	var device *Inventory
	if err := db.First(&device).Error; err != nil && err != gorm.ErrRecordNotFound {
		return device, err
	}
	return device, nil
}

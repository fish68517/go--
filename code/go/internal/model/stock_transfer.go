package model

import (
	"context"
	"fmt"
	"time"

	"github.com/mwqnice/oh-admin/internal/dto"
	"github.com/mwqnice/oh-admin/pkg/app"
	"gorm.io/gorm"
)

// StockTransfer 调拨单主表模型
type StockTransfer struct {
	ID              int       `gorm:"primaryKey;column:id;comment:主键" json:"id"`
	OrderNo         string    `gorm:"uniqueIndex;column:order_no;size:30;comment:调拨单号" json:"order_no"`
	FromWarehouseID string    `gorm:"column:from_warehouse_id;index:idx_from_warehouse;comment:调出仓库ID" json:"from_warehouse_id"`
	ToWarehouseID   string    `gorm:"column:to_warehouse_id;index:idx_to_warehouse;comment:调入仓库ID" json:"to_warehouse_id"`
	OperatorID      string    `gorm:"column:operator_id;comment:操作人ID" json:"operator_id"`
	Remark          string    `gorm:"column:remark;size:500;comment:备注" json:"remark"`
	CreateTime      time.Time `gorm:"column:create_time;default:CURRENT_TIMESTAMP;comment:创建时间" json:"create_time"`
	UpdateTime      time.Time `gorm:"column:update_time;default:CURRENT_TIMESTAMP;onUpdate:CURRENT_TIMESTAMP;comment:更新时间" json:"update_time"`
}

// TableName 设置表名
func (m StockTransfer) TableName() string {
	return "stock_transfer"
}

// StockTransferItem 调拨明细表模型
type StockTransferItem struct {
	ID          int       `gorm:"primaryKey;column:id;comment:主键" json:"id"`
	TransferID  string    `gorm:"column:transfer_id;comment:调拨单ID" json:"transfer_id"`
	ProductName string    `gorm:"column:product_name;comment:商品Name" json:"product_name"`
	ProductID   string    `gorm:"column:product_id;index:idx_product;comment:商品ID" json:"product_id"`
	Quantity    float64   `gorm:"column:quantity;type:decimal(12,2);comment:数量" json:"quantity"`
	BatchNo     string    `gorm:"column:batch_no;size:30;comment:批次号" json:"batch_no"`
	Remark      string    `gorm:"column:remark;size:200;comment:备注" json:"remark"`
	CreateTime  time.Time `gorm:"column:create_time;default:CURRENT_TIMESTAMP;comment:创建时间" json:"create_time"`
}

// TableName 设置表名
func (m StockTransferItem) TableName() string {
	return "stock_transfer_items"
}
func (l *StockTransfer) List(ctx context.Context, db *gorm.DB, params *dto.GetStockTransferListRequest) ([]*StockTransfer, int64, error) {
	db = db.WithContext(ctx).Table(l.TableName())
	if params.Name != "" {
		db = db.Where("order_no like ?", fmt.Sprintf("%%%s%%", params.Name))
	}
	var count int64
	db.Count(&count)
	var list []*StockTransfer
	if params.Page > 0 && params.PageSize > 0 {
		db = db.Offset(app.GetPageOffset(int(params.Page), int(params.PageSize))).Limit(int(params.PageSize))
	}
	if err := db.Order("id desc").Find(&list).Error; err != nil {
		return nil, 0, err
	}
	return list, count, nil
}
func (l *StockTransferItem) Get(ctx context.Context, db *gorm.DB) ([]*StockTransferItem, error) {
	db = db.WithContext(ctx).Table(l.TableName())
	fmt.Println("current id", l.ID)
	if l.TransferID != "" {
		db = db.Where("transfer_id = ? ", l.TransferID)
	}
	var drug []*StockTransferItem
	if err := db.Find(&drug).Error; err != nil && err != gorm.ErrRecordNotFound {
		return drug, err
	}
	return drug, nil
}

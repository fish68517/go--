package model

import (
	"context"
	"fmt"
	"time"

	"github.com/mwqnice/oh-admin/internal/dto"
	"github.com/mwqnice/oh-admin/pkg/app"
	"gorm.io/gorm"
)

type StockIn struct {
	ID          int       `gorm:"primaryKey;column:id" json:"id"`
	OperatorId  string    `json:"operator_id" validate:"required,min=1" comment:"操作人ID"`
	OrderNo     string    `gorm:"unique;not null;column:order_no;size:30;comment:入库单号" json:"order_no"`
	WarehouseID string    `gorm:"not null;column:warehouse_id;index:idx_warehouse;comment:仓库ID" json:"warehouse_id"`
	Remark      string    `gorm:"column:remark;size:500;comment:备注" json:"remark,omitempty"`
	CreateTime  time.Time `gorm:"column:create_time;comment:创建时间" json:"create_time"`
	UpdateTime  time.Time `gorm:"column:update_time;comment:更新时间" json:"update_time"`
	Status      int       `gorm:"column:status;index:idx_status;comment:状态" json:"status"`
}
type StockInItem struct {
	Id             int        `gorm:"primaryKey;autoIncrement;comment:主键ID" json:"id"`
	StockInId      string     `gorm:"column:stock_in_id;not null;index:idx_stock_in;comment:关联入库单ID（外键关联stock_in表id）" json:"stock_in_id"`
	ProductId      string     `gorm:"column:product_id;not null;index:idx_product;comment:产品ID（关联产品表）" json:"product_id"`
	Quantity       string     `gorm:"type:decimal(12,2);not null;comment:入库数量" json:"quantity"`
	UnitPrice      string     `gorm:"column:unit_price;type:decimal(12,2);not null;comment:单价" json:"unit_price"`
	Amount         string     `gorm:"type:decimal(12,2);not null;comment:金额（quantity×unit_price）" json:"amount"`
	BatchNo        string     `gorm:"column:batch_no;type:varchar(30);comment:批次号（可选）" json:"batch_no"`
	ProductionDate *time.Time `gorm:"column:production_date;type:date;comment:生产日期（可选）" json:"production_date"`
	ExpiryDate     *time.Time `gorm:"column:expiry_date;type:date;comment:过期日期（可选）" json:"expiry_date"`
	Remark         string     `gorm:"type:varchar(200);comment:备注（可选）" json:"remark"`
	ProductName    string     `gorm:"column:product_id;not null;index:idx_product;comment:产品ID（关联产品表）" json:"product_id"`
}

// TableName 设置表名
func (l *StockIn) TableName() string {
	return "stock_in"
}
func (l *StockInItem) TableName() string {
	return "stock_in_items"
}
func (l *StockIn) List(ctx context.Context, db *gorm.DB, params *dto.GetStockInListRequest) ([]*StockIn, int64, error) {
	db = db.WithContext(ctx).Table(l.TableName())
	if params.Name != "" {
		db = db.Where("order_no like ?", fmt.Sprintf("%%%s%%", params.Name))
	}
	var count int64
	db.Count(&count)
	var list []*StockIn
	if params.Page > 0 && params.PageSize > 0 {
		db = db.Offset(app.GetPageOffset(int(params.Page), int(params.PageSize))).Limit(int(params.PageSize))
	}
	if err := db.Order("id desc").Find(&list).Error; err != nil {
		return nil, 0, err
	}
	return list, count, nil
}

// Delete 删除
func (l *StockIn) Delete(ctx context.Context, db *gorm.DB, id int) error {
	return db.WithContext(ctx).Where("id = ?", id).Delete(&l).Error
}

// Create 插入数据
func (l *StockIn) Create(ctx context.Context, db *gorm.DB) error {
	return db.WithContext(ctx).Create(&l).Error
}

// Update 更新
func (l *StockIn) Update(ctx context.Context, db *gorm.DB) error {
	return db.WithContext(ctx).Save(&l).Error
}
func (l *StockInItem) Get(ctx context.Context, db *gorm.DB) ([]*StockInItem, error) {
	db = db.WithContext(ctx).Table(l.TableName())

	if l.StockInId != "" {
		db = db.Where("stock_in_id = ? ", l.StockInId)
	}
	var drug []*StockInItem
	if err := db.Find(&drug).Error; err != nil && err != gorm.ErrRecordNotFound {
		return drug, err
	}
	return drug, nil
}

package model

import (
	"context"
	"fmt"
	"time"

	"github.com/mwqnice/oh-admin/internal/dto"
	"github.com/mwqnice/oh-admin/pkg/app"
	"gorm.io/gorm"
)

// StockCheck 盘点主表
type StockCheck struct {
	ID          int       `gorm:"primaryKey;column:id" json:"id"`
	OrderNo     string    `gorm:"unique;column:order_no;size:255;comment:盘点单号" json:"order_no"`
	WarehouseID string    `gorm:"column:warehouse_id;comment:仓库ID" json:"warehouse_id"`
	OperatorID  string    `gorm:"column:operator_id;comment:操作人ID" json:"operator_id"`
	Remark      string    `gorm:"column:remark;size:500;comment:备注" json:"remark"`
	CreateTime  time.Time `gorm:"column:create_time;default:CURRENT_TIMESTAMP;comment:创建时间" json:"create_time"`
	UpdateTime  time.Time `gorm:"column:update_time;default:CURRENT_TIMESTAMP;comment:更新时间" json:"update_time"`
}

// TableName 自定义表名
func (StockCheck) TableName() string {
	return "stock_check"
}

// StockCheckItem 盘点明细表
type StockCheckItem struct {
	ID             int        `gorm:"primaryKey;column:id" json:"id"`
	CheckID        string     `gorm:"column:check_id;comment:盘点单ID" json:"check_id"`
	ProductID      string     `gorm:"column:product_id;size:30;comment:商品ID" json:"product_id"`
	SystemQty      string     `gorm:"column:system_quantity;type:decimal(12,2);comment:系统数量" json:"system_qty"`
	ActualQty      string     `gorm:"column:actual_quantity;type:decimal(12,2);comment:实际数量" json:"actual_qty"`
	Difference     string     `gorm:"column:difference;type:decimal(12,2);comment:差异数量" json:"difference"`
	BatchNo        string     `gorm:"column:batch_no;size:30;comment:批次号" json:"batch_no"`
	ProductionDate *time.Time `gorm:"column:production_date;comment:生产日期" json:"production_date"`
	Remark         string     `gorm:"column:remark;size:200;comment:备注" json:"remark"`
}

func (l *StockCheckItem) Get(ctx context.Context, db *gorm.DB) ([]*StockCheckItem, error) {
	db = db.WithContext(ctx).Table(l.TableName())

	if l.CheckID != "" {
		db = db.Where("check_id = ? ", l.CheckID)
	}
	var drug []*StockCheckItem
	if err := db.Find(&drug).Error; err != nil && err != gorm.ErrRecordNotFound {
		return drug, err
	}
	return drug, nil
}

// TableName 自定义表名
func (StockCheckItem) TableName() string {
	return "stock_check_items"
}
func (l *StockCheck) List(ctx context.Context, db *gorm.DB, params *dto.GetStockCheckListRequest) ([]*StockCheck, int64, error) {
	db = db.WithContext(ctx).Table(l.TableName())
	if params.Name != "" {
		db = db.Where("order_no like ?", fmt.Sprintf("%%%s%%", params.Name))
	}
	var count int64
	db.Count(&count)
	var list []*StockCheck
	if params.Page > 0 && params.PageSize > 0 {
		db = db.Offset(app.GetPageOffset(int(params.Page), int(params.PageSize))).Limit(int(params.PageSize))
	}
	if err := db.Order("id desc").Find(&list).Error; err != nil {
		return nil, 0, err
	}
	return list, count, nil
}

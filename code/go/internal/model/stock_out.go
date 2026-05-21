package model

import (
	"context"
	"fmt"
	"time"

	"github.com/mwqnice/oh-admin/internal/dto"
	"github.com/mwqnice/oh-admin/pkg/app"
	"gorm.io/gorm"
)

// StockOut 出库单表模型
type StockOut struct {
	ID          int        `gorm:"column:id;type:int;primary_key;AUTO_INCREMENT" json:"id"`
	OrderNo     string     `gorm:"column:order_no;type:varchar(30);not null;uniqueIndex:order_no" json:"order_no"`                            // 出库单号
	WarehouseID string     `gorm:"column:warehouse_id;type:int;not null;index:idx_warehouse" json:"warehouse_id"`                             // 仓库ID
	CustomerID  string     `gorm:"column:customer_id;type:int" json:"customer_id,omitempty"`                                                  // 客户ID（销售出库时关联）
	OutType     string     `gorm:"column:out_type;type:tinyint;not null;index:idx_out_type" json:"out_type"`                                  // 出库类型（1-销售出库 2-退货出库 3-调拨出库 4-其他出库）
	OperatorID  string     `gorm:"column:operator_id;type:int;not null" json:"operator_id"`                                                   // 操作人ID
	Remark      string     `gorm:"column:remark;type:varchar(500)" json:"remark,omitempty"`                                                   // 备注
	CreateTime  time.Time  `gorm:"column:create_time;type:datetime;default:CURRENT_TIMESTAMP" json:"create_time"`                             // 创建时间
	UpdateTime  time.Time  `gorm:"column:update_time;type:datetime;default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP" json:"update_time"` // 更新时间
	OutTime     *time.Time `gorm:"column:out_time;type:datetime" json:"out_time,omitempty"`                                                   // 出库时间
}

// TableName 设置表名
func (s *StockOut) TableName() string {
	return "stock_out"
}

// StockOutItem 出库明细表模型
type StockOutItem struct {
	ID             int    `gorm:"column:id;type:int;primary_key;AUTO_INCREMENT" json:"id"`
	StockOutID     string `gorm:"column:stock_out_id;type:int;not null;index:idx_stock_out" json:"stock_out_id"` // 出库单ID
	ProductID      string `gorm:"column:product_id;type:int;not null;index:idx_product" json:"product_id"`       // 商品ID
	Quantity       string `gorm:"column:quantity;type:decimal(12,2);not null" json:"quantity"`                   // 数量
	UnitPrice      string `gorm:"column:unit_price;type:decimal(12,2);" json:"unit_price"`                       // 单价
	Amount         string `gorm:"column:amount;type:decimal(12,2); " json:"amount"`                              // 金额
	BatchNo        string `gorm:"column:batch_no;type:varchar(30);index:idx_batch" json:"batch_no,omitempty"`    // 批次号
	ProductionDate string `gorm:"column:production_date;type:date" json:"production_date,omitempty"`             // 生产日期
	ExpiryDate     string `gorm:"column:expiry_date;type:date" json:"expiry_date,omitempty"`                     // 过期日期
	Remark         string `gorm:"column:remark;type:varchar(200)" json:"remark,omitempty"`                       // 备注
}

// TableName 设置表名
func (s *StockOutItem) TableName() string {
	return "stock_out_items"
}
func (l *StockOut) List(ctx context.Context, db *gorm.DB, params *dto.GetStockOutListRequest) ([]*StockOut, int64, error) {
	db = db.WithContext(ctx).Table(l.TableName())
	if params.Name != "" {
		db = db.Where("order_no like ?", fmt.Sprintf("%%%s%%", params.Name))
	}
	var count int64
	db.Count(&count)
	var list []*StockOut
	if params.Page > 0 && params.PageSize > 0 {
		db = db.Offset(app.GetPageOffset(int(params.Page), int(params.PageSize))).Limit(int(params.PageSize))
	}
	if err := db.Order("id desc").Find(&list).Error; err != nil {
		return nil, 0, err
	}
	return list, count, nil
}
func (l *StockOutItem) Get(ctx context.Context, db *gorm.DB) ([]*StockOutItem, error) {
	db = db.WithContext(ctx).Table(l.TableName())

	if l.StockOutID != "" {
		db = db.Where("stock_out_id = ? ", l.StockOutID)
	}
	var drug []*StockOutItem
	if err := db.Find(&drug).Error; err != nil && err != gorm.ErrRecordNotFound {
		return drug, err
	}
	return drug, nil
}

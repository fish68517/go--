package dto

import "time"

type GetStockTransferListRequest struct {
	Name     string `form:"name" json:"name"`
	Status   int    `json:"status"`
	Page     int64  `form:"page" json:"page"`
	PageSize int64  `form:"page_size" json:"page_size"`
}

type StockTransferCreateDTO struct {
	OrderNo         string               `json:"order_no" validate:"required,max=30" comment:"入库单号"`
	FromWarehouseID string               `gorm:"column:from_warehouse_id;index:idx_from_warehouse;comment:调出仓库ID" json:"from_warehouse_id"`
	ToWarehouseID   string               `gorm:"column:to_warehouse_id;index:idx_to_warehouse;comment:调入仓库ID" json:"to_warehouse_id"`
	Remark          string               `gorm:"column:remark;size:500;comment:备注" json:"remark"`
	OperatorId      string               `json:"operator_id" validate:"required,min=1" comment:"操作人ID"`
	Items           []StockTransferItems `json:"items" validate:"required,min=1" comment:"入库明细列表"` // 至少包含一条明细
}

// StockTransfer 调拨单主表模型
type StockTransfer struct {
	ID              int       `gorm:"primaryKey;column:id;comment:主键" json:"id"`
	OrderNo         string    `gorm:"uniqueIndex;column:order_no;size:30;comment:调拨单号" json:"order_no"`
	FromWarehouseID int       `gorm:"column:from_warehouse_id;index:idx_from_warehouse;comment:调出仓库ID" json:"from_warehouse_id"`
	ToWarehouseID   int       `gorm:"column:to_warehouse_id;index:idx_to_warehouse;comment:调入仓库ID" json:"to_warehouse_id"`
	OperatorID      int       `gorm:"column:operator_id;comment:操作人ID" json:"operator_id"`
	Remark          string    `gorm:"column:remark;size:500;comment:备注" json:"remark"`
	CreateTime      time.Time `gorm:"column:create_time;default:CURRENT_TIMESTAMP;comment:创建时间" json:"create_time"`
	UpdateTime      time.Time `gorm:"column:update_time;default:CURRENT_TIMESTAMP;onUpdate:CURRENT_TIMESTAMP;comment:更新时间" json:"update_time"`
}

// StockTransferItem 调拨明细表模型
type StockTransferItems struct {
	ID          int       `gorm:"primaryKey;column:id;comment:主键" json:"id"`
	TransferID  string    `gorm:"column:transfer_id;index:idx_transfer;comment:调拨单ID" json:"transfer_id"`
	ProductID   string    `gorm:"column:product_id;index:idx_product;comment:商品ID" json:"product_id"`
	ProductName string    `json:"product_name" validate:"required,min=1" comment:"产品name"`
	Quantity    string    `gorm:"column:quantity;type:decimal(12,2);comment:数量" json:"quantity"`
	BatchNo     string    `gorm:"column:batch_no;size:30;comment:批次号" json:"batch_no"`
	Remark      string    `gorm:"column:remark;size:200;comment:备注" json:"remark"`
	CreateTime  time.Time `gorm:"column:create_time;default:CURRENT_TIMESTAMP;comment:创建时间" json:"create_time"`
}

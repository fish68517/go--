package dto

type GetStockCheckListRequest struct {
	Name     string `form:"name" json:"name"`
	Status   int    `json:"status"`
	Page     int64  `form:"page" json:"page"`
	PageSize int64  `form:"page_size" json:"page_size"`
}

// StockCheckCreateDTO 创建盘点单请求参数
type StockCheckCreateDTO struct {
	OrderNo     string            `json:"order_no" binding:"required"`         // 盘点单号
	WarehouseID string            `json:"warehouse_id" binding:"required"`     // 仓库ID
	OperatorID  string            `json:"operator_id" binding:"required"`      // 操作人ID
	Remark      string            `json:"remark"`                              // 备注
	Items       []StockCheckItems `json:"items" binding:"required,min=1,dive"` // 盘点明细项
}

// ItemDTO 盘点明细项DTO
type StockCheckItems struct {
	CheckID        string `gorm:"column:check_id;comment:盘点单ID" json:"check_id"`
	ProductID      string `json:"product_id" binding:"required"` // 商品ID
	Difference     string `gorm:"column:difference;type:decimal(12,2);comment:差异数量" json:"difference"`
	SystemQuantity string `json:"system_quantity" binding:"required"` // 系统数量
	ActualQuantity string `json:"actual_quantity" binding:"required"` // 实际数量
	BatchNo        string `json:"batch_no"`                           // 批次号
	Remark         string `json:"remark"`                             // 备注
}

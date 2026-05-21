package dto

type GetStockInListRequest struct {
	Name     string `form:"name" json:"name"`
	Status   int    `json:"status"`
	Page     int64  `form:"page" json:"page"`
	PageSize int64  `form:"page_size" json:"page_size"`
}

// StockInCreateDTO 入库单创建请求DTO
// 用于接收前端提交的入库单主表及明细数据
type StockInCreateDTO struct {
	OrderNo     string         `json:"order_no" validate:"required,max=30" comment:"入库单号"`
	WarehouseId string         `json:"warehouse_id" validate:"required,min=1" comment:"仓库ID"`
	OperatorId  string         `json:"operator_id" validate:"required,min=1" comment:"操作人ID"`
	Items       []StockInItems `json:"items" validate:"required,min=1" comment:"入库明细列表"` // 至少包含一条明细
}

// StockInItemCreateDTO 入库单明细创建请求DTO
// 用于接收前端提交的入库明细数据（无需包含stock_in_id，由后端关联主表ID）
type StockInItems struct {
	StockInId      string `gorm:"column:stock_in_id;not null;index:idx_stock_in;comment:关联入库单ID（外键关联stock_in表id）" json:"stock_in_id"`
	ProductName    string `json:"product_name" validate:"required,min=1" comment:"产品name"`
	ProductId      string `json:"product_id" validate:"required,min=1" comment:"产品ID"`
	Quantity       string `json:"quantity" validate:"required,gt=0" comment:"数量"`   // 大于0的正数
	UnitPrice      string `json:"unit_price" validate:"required,gt=0" comment:"单价"` // 大于0的正数
	Amount         string `json:"amount" validate:"required,gt=0" comment:"金额"`     // 数量×单价，大于0
	BatchNo        string `json:"batch_no" validate:"omitempty,max=30" comment:"批次号（可选）"`
	PurchaseOrigin string `json:"purchase_origin" validate:"omitempty,max=100" comment:"进货源地"`
	SupplierName    string `json:"supplier_name" validate:"omitempty,max=100" comment:"供应商名称"`
	ProductionDate string `json:"production_date" validate:"omitempty" comment:"生产日期（可选）"`
	ExpiryDate     string `json:"expiry_date" validate:"omitempty" comment:"过期日期（可选）"`
	Remark         string `json:"remark" validate:"omitempty,max=200" comment:"备注（可选）"`
}

// StockInUpdateDTO 入库单更新请求DTO
// 用于接收入库单更新数据（仅包含可修改字段）
type StockInUpdateDTO struct {
	Id          int    `json:"id" validate:"required,min=1" comment:"入库单ID"`
	OrderNo     string `json:"order_no" validate:"omitempty,max=30" comment:"入库单号（可选修改）"`
	WarehouseId int    `json:"warehouse_id" validate:"omitempty,min=1" comment:"仓库ID（可选修改）"`
	OperatorId  int    `json:"operator_id" validate:"omitempty,min=1" comment:"操作人ID（可选修改）"`
}

package dto

type GetInventoryListRequest struct {
	Name     string `form:"name" json:"name"`
	Status   int    `json:"status"`
	Page     int64  `form:"page" json:"page"`
	PageSize int64  `form:"page_size" json:"page_size"`
}
type InventoryOperationDTO struct {
	ID           string `json:"id"`            // 库存记录ID
	ProductID    string `json:"product_id"`    // 商品编号
	ProductName  string `json:"product_name"`  // 商品名称
	WarehouseID  string `json:"warehouse_id"`  // 仓库ID
	CurrentStock string `json:"current_stock"` // 当前库存数量
	Quantity     string `json:"quantity"`      // 操作数量(增减量)
}

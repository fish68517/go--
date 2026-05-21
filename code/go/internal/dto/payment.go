package dto

type DrugSourceResponse struct {
	StockInItemID  int    `json:"stock_in_item_id"`
	ProductID      string `json:"product_id"`
	ProductName    string `json:"product_name"`
	PurchaseOrigin string `json:"purchase_origin"`
	SupplierName    string `json:"supplier_name"`
	BatchNo        string `json:"batch_no"`
	UnitPrice      string `json:"unit_price"`
	Quantity       string `json:"quantity"`
	Remark         string `json:"remark"`
}

type PaymentItem struct {
	DrugProductNumber string  `json:"drug_product_number"`
	DrugProductName   string  `json:"drug_product_name"`
	PurchaseOrigin    string  `json:"purchase_origin"`
	BatchNo           string  `json:"batch_no"`
	Quantity          float64 `json:"quantity"`
	UnitPrice         float64 `json:"unit_price"`
	Amount            float64 `json:"amount"`
}

type PaymentDetailResponse struct {
	PrescriptionID     int           `json:"prescription_id"`
	PrescriptionNumber string        `json:"prescription_number"`
	PayStatus          string        `json:"pay_status"`
	TotalAmount        float64       `json:"total_amount"`
	PayAmount          float64       `json:"pay_amount"`
	PayMethod          string        `json:"pay_method"`
	MockTradeNo        string        `json:"mock_trade_no"`
	PaidAt             string        `json:"paid_at"`
	Items              []PaymentItem `json:"items"`
}

type MockPayRequest struct {
	PrescriptionID int    `json:"prescription_id" binding:"required"`
	PayMethod      string `json:"pay_method"`
}

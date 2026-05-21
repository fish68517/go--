package dto

import (
	"time"
)

// PrescriptionFlow represents a single step in the prescription flow.
type PrescriptionFlow struct {
	ID        uint   `gorm:"primaryKey"`
	RxNumber  string `gorm:"index"` // Prescription number
	Step      string
	Performer string
	Timestamp time.Time
}

// PrescriptionFlowResponse is the structure for the JSON response to the frontend.
type PrescriptionFlowResponse struct {
	Code    int                  `json:"code"`
	Message string               `json:"message"`
	Data    PrescriptionFlowData `json:"data"`
}

type PrescriptionFlowData struct {
	Process    []PrescriptionFlow `json:"process"`
	Medicines  []TraceMedicineItem `json:"medicines"`
	Payment    TracePayment        `json:"payment"`
	Blockchain TraceBlockchain     `json:"blockchain"`
}

type TraceMedicineItem struct {
	DrugProductNumber string  `json:"drug_product_number"`
	DrugProductName   string  `json:"drug_product_name"`
	PurchaseOrigin    string  `json:"purchase_origin"`
	BatchNo           string  `json:"batch_no"`
	Quantity          float64 `json:"quantity"`
	UnitPrice         float64 `json:"unit_price"`
	TotalPrice        float64 `json:"total_price"`
}

type TracePayment struct {
	PayStatus   string  `json:"pay_status"`
	TotalAmount float64 `json:"total_amount"`
	PayAmount   float64 `json:"pay_amount"`
	PayMethod   string  `json:"pay_method"`
	MockTradeNo string  `json:"mock_trade_no"`
	PaidAt      string  `json:"paid_at"`
}

type TraceBlockchain struct {
	ChainStatus string `json:"chain_status"`
	TxID        string `json:"tx_id"`
	PayloadHash string `json:"payload_hash"`
	Error       string `json:"error"`
}

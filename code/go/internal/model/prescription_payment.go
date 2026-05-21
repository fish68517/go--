package model

import "time"

type PrescriptionPayment struct {
	ID                 int        `gorm:"primaryKey;column:id" json:"id"`
	PrescriptionID     int        `gorm:"column:prescription_id" json:"prescription_id"`
	PrescriptionNumber string     `gorm:"column:prescription_number" json:"prescription_number"`
	TotalAmount        float64    `gorm:"column:total_amount" json:"total_amount"`
	PayAmount          float64    `gorm:"column:pay_amount" json:"pay_amount"`
	PayMethod          string     `gorm:"column:pay_method" json:"pay_method"`
	PayStatus          string     `gorm:"column:pay_status" json:"pay_status"`
	MockTradeNo        string     `gorm:"column:mock_trade_no" json:"mock_trade_no"`
	PaidAt             *time.Time `gorm:"column:paid_at" json:"paid_at"`
	CreatedAt          time.Time  `gorm:"column:created_at" json:"created_at"`
	UpdatedAt          time.Time  `gorm:"column:updated_at" json:"updated_at"`
}

func (l *PrescriptionPayment) TableName() string {
	return "prescription_payment"
}

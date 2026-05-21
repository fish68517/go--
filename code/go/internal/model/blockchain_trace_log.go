package model

import "time"

type BlockchainTraceLog struct {
	ID                 int       `gorm:"primaryKey;column:id" json:"id"`
	PrescriptionID     int       `gorm:"column:prescription_id" json:"prescription_id"`
	PrescriptionNumber string    `gorm:"column:prescription_number" json:"prescription_number"`
	PayloadHash        string    `gorm:"column:payload_hash" json:"payload_hash"`
	TxID               string    `gorm:"column:tx_id" json:"tx_id"`
	ChainStatus        string    `gorm:"column:chain_status" json:"chain_status"`
	ErrorMessage       string    `gorm:"column:error_message" json:"error_message"`
	CreatedAt          time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt          time.Time `gorm:"column:updated_at" json:"updated_at"`
}

func (l *BlockchainTraceLog) TableName() string {
	return "blockchain_trace_log"
}

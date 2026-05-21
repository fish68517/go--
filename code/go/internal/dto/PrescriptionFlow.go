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
	Process []PrescriptionFlow `json:"process"`
}

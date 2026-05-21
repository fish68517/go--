package model

import (
	"context"
	"gorm.io/gorm"
)

type PrescriptionDrug struct {
	ID                     int     `db:"id" json:"id"`
	PrescriptionID         int     `db:"prescription_id" json:"prescription_id"`
	HospitalID             string  `db:"hospital_id" json:"hospital_id"`
	DrugProductNumber      string  `db:"drug_product_number" json:"drug_product_number"`
	DrugProductName        string  `db:"drug_product_name" json:"drug_product_name"`
	MeasurementUnit        string  `db:"measurement_unit" json:"measurement_unit"`
	DrugProductDescription string  `db:"drug_product_description" json:"drug_product_description"`
	Dose                   int     `db:"dose" json:"dose"`
	DrugWeight             float64 `db:"durg_weight" json:"drug_weight"` // Note: Fixed typo from `durg_weight` to `drug_weight`
	DrugWeights            float64 `db:"drug_weights" json:"drug_weights"`
	DrugPrice              float64 `db:"drug_price" json:"drug_price"`
	TotalPrices            float64 `db:"total_prices" json:"total_prices"`
}
type QrCode struct {
	QrCode string
}
type QrcodeVw struct {
	ID                  int
	HospitalID          string
	HospitalName        string
	PrescriptionNumber  string
	PatientName         string
	PatientAddress      string
	PackNum             string
	DeScheme            string
	PackAcount          string
	AdministrationCount string
	BNum                string
}

func (l *QrcodeVw) TableName() string {
	return "qrcode_vw"
}
func (l *PrescriptionDrug) TableName() string {
	return "prescription_drug"
}
func (drug *PrescriptionDrug) Create(ctx context.Context, db *gorm.DB) error {
	return db.WithContext(ctx).Create(&drug).Error
}

type MedicineGroup struct {
	Row1 *PrescriptionDrug
	Row2 *PrescriptionDrug
}

func (l *QrcodeVw) GetQrcodeVw(ctx context.Context, db *gorm.DB) (*QrcodeVw, error) {
	db = db.WithContext(ctx).Table(l.TableName())

	if l.ID != 0 {
		db = db.Where("id = ? ", l.ID)
	}
	var qrcode *QrcodeVw
	if err := db.First(&qrcode).Error; err != nil && err != gorm.ErrRecordNotFound {
		return qrcode, err
	}
	return qrcode, nil
}
func (l *PrescriptionDrug) Get(ctx context.Context, db *gorm.DB) ([]*PrescriptionDrug, error) {
	db = db.WithContext(ctx).Table(l.TableName())

	if l.PrescriptionID != 0 {
		db = db.Where("prescription_id = ? ", l.PrescriptionID)
	}
	var drug []*PrescriptionDrug
	if err := db.Find(&drug).Error; err != nil && err != gorm.ErrRecordNotFound {
		return drug, err
	}
	return drug, nil
}

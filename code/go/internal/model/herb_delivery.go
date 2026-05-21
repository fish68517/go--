package model

import (
	"context"
	"fmt"
	"github.com/mwqnice/oh-admin/internal/dto"
	"github.com/mwqnice/oh-admin/pkg/app"
	"gorm.io/gorm"
)

type HerbDeliveryView struct {
	ID                   int    `json:"id" db:"id"`
	HospitalName         string `json:"hospital_name" db:"hospital_name"`
	PatientName          string `json:"patient_name" db:"patient_name"`
	PrescriptionNumber   string `json:"prescription_number" db:"prescription_number"`
	PrescriptionID       int    `json:"prescription_id" db:"prescription_id"`
	DeliveryPersonnel    string `json:"delivery_personnel" db:"delivery_personnel"`
	DeliveryTime         string `json:"delivery_time" db:"delivery_time"`
	ProcessingEmployeeID int    `json:"processing_employee_id" db:"processing_employee_id"`
	WordContent          string `json:"word_content" db:"word_content"`
	Remarks              string `json:"remarks" db:"remarks"`
	DeliveryStatus       int    `json:"delivery_status" db:"delivery_status"` // 假设状态是整数类型
	LogisticsNumber      string `json:"logistics_number" db:"logistics_number"`
}
type HerbDelivery struct {
	ID                   int    `gorm:"primaryKey;autoIncrement"`
	PrescriptionID       int    `gorm:"comment:'处方ID'"`
	DeliveryPersonnel    string `gorm:"not null;comment:'发货人'"`
	DeliveryTime         string `gorm:"comment:'发货时间'"`
	DeliveryStatus       int    `gorm:"comment:'发货状态'"`
	Remarks              string `gorm:"comment:'备注'"`
	ProcessingEmployeeID int    `gorm:"comment:'处理员工ID'"`
	LogisticsNumber      string `gorm:"comment:'物流单号'"`
	WordContent          string `gorm:"comment:'word内容'"`
	Barcode              string `gorm:"comment:'Barcode'"`
}

func (l *HerbDelivery) TableName() string {
	return "herb_delivery"
}
func (l *HerbDeliveryView) TableName() string {
	return "herb_delivery_view"
}
func (l *HerbDeliveryView) List(ctx context.Context, db *gorm.DB, params *dto.GetDeliveryListRequest) ([]*HerbDeliveryView, int64, error) {
	db = db.WithContext(ctx).Table(l.TableName())
	if params.Name != "" {
		db = db.Where("prescription_number like ?", fmt.Sprintf("%%%s%%", params.Name))
	}
	var count int64
	db.Count(&count)
	var list []*HerbDeliveryView
	if params.Page > 0 && params.PageSize > 0 {
		db = db.Offset(app.GetPageOffset(int(params.Page), int(params.PageSize))).Limit(int(params.PageSize))
	}
	if err := db.Order("id desc").Find(&list).Error; err != nil {
		return nil, 0, err
	}
	return list, count, nil
}
func (l *HerbDelivery) Create(ctx context.Context, db *gorm.DB) error {
	tx := db.Begin()

	// 插入调剂记录
	if err := tx.Create(&l).Error; err != nil {
		tx.Rollback()
		return err
	}
	// 更新处方状态
	var prescription Prescription
	if err := tx.Model(&Prescription{}).Where("id = ?", l.PrescriptionID).Update("current_state", "发货").Find(&prescription).Error; err != nil {
		tx.Rollback()
		return err
	}
	// 提交事务
	tx.Commit()
	return nil
}

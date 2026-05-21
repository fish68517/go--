package model

import (
	"context"
	"fmt"
	"github.com/mwqnice/oh-admin/internal/dto"
	"github.com/mwqnice/oh-admin/pkg/app"
	"gorm.io/gorm"
)

type MedicinePackingView struct {
	ID                 int    `json:"id" db:"id"`
	WordContent        string `json:"word_content" db:"word_content"`
	HospitalName       string `json:"hospital_name" db:"hospital_name"`
	PatientName        string `json:"patient_name" db:"patient_name"`
	PrescriptionNumber string `json:"prescription_number" db:"prescription_number"`
	Barcode            string `json:"barcode" db:"barcode"`
	StartTime          string `json:"start_time" db:"start_time"`
	EmployeeID         int    `json:"employee_id" db:"employee_id"`
	PackingPersonnel   string `json:"packing_personnel" db:"packing_personnel"`
	PackingStatus      int    `json:"packing_status" db:"packing_status"` // 假设状态是整数类型
	EndTime            string `json:"end_time" db:"end_time"`
}
type MedicinePacking struct {
	ID               int    `gorm:"primaryKey;autoIncrement"`
	Barcode          string `gorm:"comment:'条码'"`
	WordContent      string `gorm:"comment:'word内容'"`
	PrescriptionID   int    `gorm:"type:varchar(50);comment:'处方ID'"`
	PackingPersonnel string `gorm:"comment:'包装人员'"`
	StartTime        string `gorm:"comment:'开始时间'"`
	EndTime          string `gorm:"comment:'包装完成时间'"`
	PackingStatus    int    `gorm:"comment:'包装状态'"`
	EmployeeID       int    `gorm:"comment:'员工ID'"`
}

func (l *MedicinePacking) TableName() string {
	return "medicine_packing"
}
func (l *MedicinePackingView) TableName() string {
	return "medicine_packing_view"
}
func (l *MedicinePackingView) List(ctx context.Context, db *gorm.DB, params *dto.GetPackInfoListRequest) ([]*MedicinePackingView, int64, error) {
	db = db.WithContext(ctx).Table(l.TableName())
	if params.Name != "" {
		db = db.Where("prescription_number like ?", fmt.Sprintf("%%%s%%", params.Name))
	}
	var count int64
	db.Count(&count)
	var list []*MedicinePackingView
	if params.Page > 0 && params.PageSize > 0 {
		db = db.Offset(app.GetPageOffset(int(params.Page), int(params.PageSize))).Limit(int(params.PageSize))
	}
	if err := db.Order("id desc").Find(&list).Error; err != nil {
		return nil, 0, err
	}
	return list, count, nil
}
func (l *MedicinePacking) Create(ctx context.Context, db *gorm.DB) error {
	tx := db.Begin()

	// 插入调剂记录
	if err := tx.Create(&l).Error; err != nil {
		tx.Rollback()
		return err
	}
	// 更新处方状态
	var prescription Prescription
	if err := tx.Model(&Prescription{}).Where("id = ?", l.PrescriptionID).Update("current_state", "包装").Find(&prescription).Error; err != nil {
		tx.Rollback()
		return err
	}
	// 提交事务
	tx.Commit()
	return nil
}

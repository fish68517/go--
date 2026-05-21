package model

import (
	"context"
	"fmt"
	"github.com/mwqnice/oh-admin/internal/dto"
	"github.com/mwqnice/oh-admin/pkg/app"
	"gorm.io/gorm"
)

type AuditView struct {
	ID                 int    `json:"id" db:"id"`
	HospitalName       string `json:"hospital_name" db:"hospital_name"`
	PatientName        string `json:"patient_name" db:"patient_name"`
	PrescriptionNumber string `json:"prescription_number" db:"prescription_number"`
	Barcode            string `json:"barcode" db:"barcode"`
	WordContent        string `json:"word_content" db:"word_content"`
	Reviewer           string `json:"reviewer" db:"reviewer"`
	AuditDatetime      string `json:"audit_datetime" db:"audit_datetime"`
	AuditStatus        int    `json:"audit_status" db:"audit_status"`
	EmployeeID         int    `json:"employeeID" db:"employeeID"`
}
type HerbDecoctionAudit struct {
	ID             uint   `gorm:"primaryKey;autoIncrement" json:"id"`     // 自增主键
	PrescriptionID int    `gorm:"not null" json:"prescription_id"`        // 处方ID
	Barcode        string `gorm:"size:50" json:"barcode"`                 // 条形码
	Reviewer       string `gorm:"size:100" json:"reviewer"`               // 审核人
	WordContent    string `gorm:"size:50" json:"word_content"`            // 文字内容
	EmployeeID     int    `json:"employee_id"`                            // 员工ID
	AuditDatetime  string `gorm:"not null" json:"audit_datetime"`         // 审核日期时间
	AuditStatus    int    `gorm:"not null;default:0" json:"audit_status"` // 审核状态
}

func (l *HerbDecoctionAudit) TableName() string {
	return "herb_decoction_audit"
}
func (l *AuditView) TableName() string {
	return "herb_decoction_audit_view"
}
func (l *AuditView) List(ctx context.Context, db *gorm.DB, params *dto.GetAuditListRequest) ([]*AuditView, int64, error) {
	db = db.WithContext(ctx).Table(l.TableName())
	if params.Name != "" {
		db = db.Where("prescription_number like ?", fmt.Sprintf("%%%s%%", params.Name))
	}
	var count int64
	db.Count(&count)
	var list []*AuditView
	if params.Page > 0 && params.PageSize > 0 {
		db = db.Offset(app.GetPageOffset(int(params.Page), int(params.PageSize))).Limit(int(params.PageSize))
	}
	if err := db.Order("id desc").Find(&list).Error; err != nil {
		return nil, 0, err
	}
	return list, count, nil
}
func (l *HerbDecoctionAudit) Create(ctx context.Context, db *gorm.DB) error {
	tx := db.Begin()

	// 插入调剂记录
	if err := tx.Create(&l).Error; err != nil {
		tx.Rollback()
		return err
	}
	// 更新处方状态
	var prescription Prescription
	if err := tx.Model(&Prescription{}).Where("id = ?", l.PrescriptionID).Update("current_state", "复核").Find(&prescription).Error; err != nil {
		tx.Rollback()
		return err
	}
	// 提交事务
	tx.Commit()
	return nil
}

package model

import (
	"context"
	"fmt"
	"github.com/mwqnice/oh-admin/internal/dto"
	"github.com/mwqnice/oh-admin/pkg/app"
	"gorm.io/gorm"
)

type DecoctionView struct {
	ID                 int    `json:"id" db:"id"`
	HospitalName       string `json:"hospital_name" db:"hospital_name"`
	PatientName        string `json:"patient_name" db:"patient_name"`
	PrescriptionNumber string `json:"prescription_number" db:"prescription_number"`
	Barcode            string `json:"barcode" db:"barcode"`
	DecoctionManager   string `json:"decoction_manager" db:"decoction_manager"`
	DecoctionTime      string `json:"decoction_time" db:"decoction_time"`
	StartTime          string `json:"start_time" db:"start_time"`
	EndTime            string `json:"end_time" db:"end_time"`
	DecoctionStatus    int    `json:"decoction_status" db:"decoction_status"` // 假设状态是整数类型
	EmployeeID         int    `json:"employee_id" db:"employee_id"`
	MachineID          int    `json:"machine_id" db:"machine_id"`
	Remark             string `json:"remark" db:"remark"`
}
type HerbalDecoctionInfo struct {
	ID               int    `gorm:"primaryKey;autoIncrement"`
	PrescriptionID   int    `gorm:"not null"`
	MachineID        int    `gorm:"not null"`
	DecoctionManager string `gorm:"comment:'煎药工姓名'"`
	DecoctionTime    int    `gorm:"comment:'煎药时间'"`
	DecoctionStatus  int    `gorm:"not null;default:0;comment:'煎药状态'"`
	StartTime        string `gorm:"not null;comment:'开始时间'"`
	EndTime          string `gorm:"comment:'结束时间'"`
	WordContent      string `gorm:"comment:'word内容'"`
	EmployeeID       int    `gorm:"comment:'员工ID'"`
	Barcode          string `gorm:"comment:'条码'"`
	Remark           string `gorm:"type:varchar(255);comment:'备注'"`
}

func (l *HerbalDecoctionInfo) TableName() string {
	return "herbal_decoction_info"
}
func (l *DecoctionView) TableName() string {
	return "herbal_decoction_info_view"
}
func (l *DecoctionView) List(ctx context.Context, db *gorm.DB, params *dto.GetDecoInfoListRequest) ([]*DecoctionView, int64, error) {
	db = db.WithContext(ctx).Table(l.TableName())
	if params.Name != "" {
		db = db.Where("prescription_number like ?", fmt.Sprintf("%%%s%%", params.Name))
	}
	var count int64
	db.Count(&count)
	var list []*DecoctionView
	if params.Page > 0 && params.PageSize > 0 {
		db = db.Offset(app.GetPageOffset(int(params.Page), int(params.PageSize))).Limit(int(params.PageSize))
	}
	if err := db.Order("id desc").Find(&list).Error; err != nil {
		return nil, 0, err
	}
	return list, count, nil
}
func (l *HerbalDecoctionInfo) Create(ctx context.Context, db *gorm.DB) error {
	tx := db.Begin()

	// 插入调剂记录
	if err := tx.Create(&l).Error; err != nil {
		tx.Rollback()
		return err
	}
	// 更新处方状态
	var prescription Prescription
	if err := tx.Model(&Prescription{}).Where("id = ?", l.PrescriptionID).Update("current_state", "煎药").Find(&prescription).Error; err != nil {
		tx.Rollback()
		return err
	}
	// 提交事务
	tx.Commit()
	return nil
}

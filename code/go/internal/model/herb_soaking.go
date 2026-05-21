package model

import (
	"context"
	"fmt"
	"github.com/mwqnice/oh-admin/internal/dto"
	"github.com/mwqnice/oh-admin/pkg/app"
	"gorm.io/gorm"
)

// SoakView 结构体对应 MySQL 查询结果
type SoakView struct {
	ID                 int    `json:"id" db:"id"`
	HospitalName       string `json:"hospital_name" db:"hospital_name"`
	PatientName        string `json:"patient_name" db:"patient_name"`
	PrescriptionNumber string `json:"prescription_number" db:"prescription_number"`
	Barcode            string `json:"barcode" db:"barcode"`
	SoakingPerson      string `json:"soaking_person" db:"soaking_person"`
	StartTime          string `json:"start_time" db:"start_time"`
	EndTime            string `json:"end_time" db:"end_time"`
	Duration           int    `json:"duration" db:"duration"` // 假设浸泡时长是整数类型
	WordContent        string `json:"word_content" db:"word_content"`
	Remarks            string `json:"remarks" db:"remarks"`
	PrescriptionID     int    `json:"prescription_id" db:"prescription_id"`
	EmployeeID         int    `json:"employee_id" db:"employee_id"`
}
type HerbSoaking struct {
	ID             uint   `gorm:"primaryKey;autoIncrement" json:"id"`       // 自增主键
	PrescriptionID int    `gorm:"not null" json:"prescription_id"`          // 处方ID
	WordContent    string `gorm:"size:50" json:"word_content"`              // 文字内容
	StartTime      string `gorm:"not null" json:"start_time"`               // 开始时间
	Duration       int    `gorm:"default:null" json:"duration"`             // 持续时间（分钟）
	WarningStatus  int    `gorm:"default:null" json:"warning_status"`       // 警告状态
	Remarks        string `gorm:"size:1000" json:"remarks"`                 // 备注
	EndTime        string `gorm:"default:null" json:"end_time,omitempty"`   // 结束时间（可为空）
	SoakingPerson  string `gorm:"size:100" json:"soaking_person"`           // 浸泡人
	EmployeeID     int    `gorm:"default:null" json:"employee_id"`          // 员工ID
	Barcode        string `gorm:"size:50" json:"barcode"`                   // 条形码
	Mark           string `gorm:"size:50" json:"mark"`                      // 标记
	WarningTime    string `gorm:"not null" json:"warning_time"`             // 警告时间
	WarningType    string `gorm:"size:100" json:"warning_type"`             // 警告类型
	SoakingStatus  int    `gorm:"not null;default:0" json:"soaking_status"` // 浸泡状态
}

// TableName 指定模型对应的数据库表名（可选，如果模型名与表名不一致时需要）
func (HerbSoaking) TableName() string {
	return "herb_soaking"
}

func (l *SoakView) TableName() string {
	return "herb_soaking_view"
}
func (l *SoakView) List(ctx context.Context, db *gorm.DB, params *dto.GetSoakListRequest) ([]*SoakView, int64, error) {
	db = db.WithContext(ctx).Table(l.TableName())
	if params.Name != "" {
		db = db.Where("prescription_number like ?", fmt.Sprintf("%%%s%%", params.Name))
	}
	var count int64
	db.Count(&count)
	var list []*SoakView
	if params.Page > 0 && params.PageSize > 0 {
		db = db.Offset(app.GetPageOffset(int(params.Page), int(params.PageSize))).Limit(int(params.PageSize))
	}
	if err := db.Order("id desc").Find(&list).Error; err != nil {
		return nil, 0, err
	}
	return list, count, nil
}
func (l *HerbSoaking) Create(ctx context.Context, db *gorm.DB) error {
	tx := db.Begin()

	// 插入调剂记录
	if err := tx.Create(&l).Error; err != nil {
		tx.Rollback()
		return err
	}
	// 更新处方状态
	var prescription Prescription
	if err := tx.Model(&Prescription{}).Where("id = ?", l.PrescriptionID).Update("current_state", "泡药").Find(&prescription).Error; err != nil {
		tx.Rollback()
		return err
	}
	// 提交事务
	tx.Commit()
	return nil
}

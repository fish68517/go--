package model

import (
	"context"
	"fmt"
	"github.com/mwqnice/oh-admin/internal/dto"
	"github.com/mwqnice/oh-admin/pkg/app"
	"gorm.io/gorm"
)

type Adjustment struct {
	ID                  int    `json:"id" db:"id"`
	WordContent         string `json:"word_content" db:"word_content"`
	WordDate            string `json:"word_date" db:"word_date"`
	WorkloadDescription string `json:"workload_description" db:"workload_description"`
	WordPerson          string `json:"word_person" db:"word_person"`                   // 使用 float64 来存储 decimal 类型
	EmployeeID          int    `json:"employee_id,omitempty" db:"employee_id"`         // 使用指针来处理可能为 NULL 的整数
	PrescriptionID      int    `json:"prescription_id,omitempty" db:"prescription_id"` // 同样使用指针
	Status              int    `json:"status" db:"status"`
	Barcode             string `json:"barcode" db:"barcode"`
	EndDate             string `json:"end_date,omitempty" db:"end_date"` // 使用指针来处理可能为 NULL 的日期时间
	// 注意：这里没有包含 PRIMARY KEY，因为它在 Go 结构体中不是必需的
}

// TableName 定义了模型的表名（可选，如果结构体名与表名一致则可以省略）
func (Adjustment) TableName() string {
	return "adjustment"
}

// AdjustmentView 结构体对应 MySQL 的 adjustment_view 视图
type AdjustmentView struct {
	Id                  int    `json:"id" db:"id"`
	HospitalName        string `json:"hospital_name" db:"hospital_name"`
	PatientName         string `json:"patient_name" db:"patient_name"`
	PrescriptionNumber  string `json:"prescription_number" db:"prescription_number"`
	WordDate            string `json:"word_date" db:"word_date"`
	EndDate             string `json:"end_date" db:"end_date"` // 使用 sql.NullTime 来处理可能为 NULL 的日期时间
	WordPerson          string `json:"word_person" db:"word_person"`
	Status              int    `json:"status" db:"status"` // 假设状态是整数类型
	WordContent         string `json:"word_content" db:"word_content"`
	WorkloadDescription string `json:"workload_description" db:"workload_description"`
	Barcode             string `json:"barcode" db:"barcode"` // 如果 barcode 可能为 NULL，则使用 sql.NullString
}

func (l *AdjustmentView) TableName() string {
	return "adjustment_view"
}
func (l *AdjustmentView) List(ctx context.Context, db *gorm.DB, params *dto.GetAdjustListRequest) ([]*AdjustmentView, int64, error) {
	db = db.WithContext(ctx).Table(l.TableName())
	if params.Name != "" {
		db = db.Where("prescription_number like ?", fmt.Sprintf("%%%s%%", params.Name))
	}
	var count int64
	db.Count(&count)
	var list []*AdjustmentView
	if params.Page > 0 && params.PageSize > 0 {
		db = db.Offset(app.GetPageOffset(int(params.Page), int(params.PageSize))).Limit(int(params.PageSize))
	}
	if err := db.Order("id desc").Find(&list).Error; err != nil {
		return nil, 0, err
	}
	return list, count, nil
}
func (l *Adjustment) Create(ctx context.Context, db *gorm.DB) error {
	// 开始事务
	tx := db.Begin()

	// 插入调剂记录
	if err := tx.Create(&l).Error; err != nil {
		tx.Rollback()
		return err
	}
	// 更新处方状态
	var prescription Prescription
	if err := tx.Model(&Prescription{}).Where("id = ?", l.PrescriptionID).Update("current_state", "调剂").Find(&prescription).Error; err != nil {
		tx.Rollback()
		return err
	}
	// 提交事务
	tx.Commit()
	return nil
}

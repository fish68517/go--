package model

import (
	"context"
	"fmt"

	"github.com/mwqnice/oh-admin/internal/dto"
	"github.com/mwqnice/oh-admin/pkg/app"
	"gorm.io/gorm"
)

// TDrugDosageLimit 药品剂量限制模型，对应数据表t_drug_dosage_limit
type TDrugDosageLimit struct {
	*Model             // 自增主键
	DrugID     string  `gorm:"column:drug_id;type:varchar(32);not null" json:"drug_id"`         // 关联药品ID
	MaxDosage  float64 `gorm:"column:max_dosage;type:decimal(10,2);not null" json:"max_dosage"` // 最大剂量
	DosageDesc string  `gorm:"column:dosage_desc;type:text" json:"dosage_desc"`                 // 剂量说明
	IsEnabled  int     `gorm:"column:is_enabled;not null;default:1" json:"is_enabled"`          // 是否启用（1=启用，0=禁用）
}

// TableName 指定数据表名
func (TDrugDosageLimit) TableName() string {
	return "t_drug_dosage_limit"
}
func (l *TDrugDosageLimit) List(ctx context.Context, db *gorm.DB, params *dto.GetDrugDosageListRequest) ([]*TDrugDosageLimit, int64, error) {
	db = db.WithContext(ctx).Table(l.TableName())
	if params.Name != "" {
		db = db.Where("drug_id like ?", fmt.Sprintf("%%%s%%", params.Name))
	}
	var count int64
	db.Count(&count)
	var list []*TDrugDosageLimit
	if params.Page > 0 && params.PageSize > 0 {
		db = db.Offset(app.GetPageOffset(int(params.Page), int(params.PageSize))).Limit(int(params.PageSize))
	}
	if err := db.Order("id desc").Find(&list).Error; err != nil {
		return nil, 0, err
	}
	return list, count, nil
}

// Get 根据条件查询单条数据
func (l *TDrugDosageLimit) Get(ctx context.Context, db *gorm.DB) (*TDrugDosageLimit, error) {
	db = db.WithContext(ctx).Table(l.TableName())

	if l.Model != nil && l.ID != 0 {
		db = db.Where("id = ? ", l.ID)
	}
	var link *TDrugDosageLimit
	if err := db.First(&link).Error; err != nil && err != gorm.ErrRecordNotFound {
		return link, err
	}
	return link, nil
}
func (l *TDrugDosageLimit) GetSingleDosageLimit(ctx context.Context, db *gorm.DB, params string) (*TDrugDosageLimit, error) {
	db = db.WithContext(ctx).Table(l.TableName())
	if params != "" {
		db = db.Where("drug_id = ? ", params)
	}

	var link *TDrugDosageLimit
	if err := db.First(&link).Error; err != nil && err != gorm.ErrRecordNotFound {
		return link, err
	}
	return link, nil
}
func (l *TDrugDosageLimit) Create(ctx context.Context, db *gorm.DB) error {
	return db.WithContext(ctx).Create(&l).Error
}

// Update 更新
func (l *TDrugDosageLimit) Update(ctx context.Context, db *gorm.DB) error {
	return db.WithContext(ctx).Save(&l).Error
}

// Delete 删除
func (l *TDrugDosageLimit) Delete(ctx context.Context, db *gorm.DB, id int) error {
	return db.WithContext(ctx).Where("id = ?", id).Delete(&l).Error
}

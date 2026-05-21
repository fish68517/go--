package model

import (
	"context"
	"github.com/mwqnice/oh-admin/internal/dto"
	"github.com/mwqnice/oh-admin/pkg/app"
	"gorm.io/gorm"
)

type SamInfo struct {
	ID                   int    `json:"id" db:"id"`
	PrescriptionNumber   string `json:"prescription_number" db:"prescription_number"`
	PrescriptionWeight   string `json:"prescription_weight" db:"prescription_weight"`
	PrescriptionSjWeight string `json:"prescription_sj_weight" db:"prescription_sj_weight"`
	DrugCount            int    `json:"drug_count" db:"drug_count"`
	DrugSjCount          int    `json:"drug_sj_count" db:"drug_sj_count"`
	Dosage               int    `json:"dosage" db:"dosage"`
	OperateName          string `json:"operate_name" db:"operate_name"`
	OperateTime          string `json:"operate_time" db:"operate_time"`
	OperateUpdateTime    string `json:"operate_update_time" db:"operate_update_time"`
	Remark               string `json:"remark" db:"remark"`
	Status               int    `json:"state" db:"status"`
}

func (l *SamInfo) TableName() string {
	return "sam_info"
}
func (l *SamInfo) List(ctx context.Context, db *gorm.DB, params *dto.GetSampleListRequest) ([]*SamInfo, int64, error) {
	db = db.WithContext(ctx).Table(l.TableName())
	if params.Status != 0 {
		db = db.Where("status = ?", params.Status)
	}
	if params.PrescriptionNumber != "" {
		db = db.Where("prescription_number = ?", params.PrescriptionNumber)
	}
	var count int64
	db.Count(&count)
	var list []*SamInfo
	if params.Page > 0 && params.PageSize > 0 {
		db = db.Offset(app.GetPageOffset(int(params.Page), int(params.PageSize))).Limit(int(params.PageSize))
	}
	if err := db.Order("id asc").Find(&list).Error; err != nil {
		return nil, 0, err
	}
	return list, count, nil
}
func (l *SamInfo) Create(ctx context.Context, db *gorm.DB) error {
	return db.WithContext(ctx).Create(&l).Error
}

// Get 根据条件查询单条数据
func (l *SamInfo) Get(ctx context.Context, db *gorm.DB) (*SamInfo, error) {
	db = db.WithContext(ctx).Table(l.TableName())

	if l.ID != 0 {
		db = db.Where("id = ? ", l.ID)
	}
	var sam *SamInfo
	if err := db.First(&sam).Error; err != nil && err != gorm.ErrRecordNotFound {
		return sam, err
	}
	return sam, nil
}
func (l *SamInfo) Update(ctx context.Context, db *gorm.DB) error {
	return db.WithContext(ctx).Save(&l).Error
}
func (l *SamInfo) Delete(ctx context.Context, db *gorm.DB, id int) error {
	return db.WithContext(ctx).Where("id = ?", id).Delete(&l).Error
}

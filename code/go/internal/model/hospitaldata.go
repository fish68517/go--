package model

import (
	"context"
	"fmt"
	"github.com/mwqnice/oh-admin/internal/dto"
	"github.com/mwqnice/oh-admin/pkg/app"
	"gorm.io/gorm"
)

type HospitalData struct {
	HospitalName         string `gorm:"column:hospital_name;NOT NULL" json:"hospital_name"`                   // 医院名称
	PrescriptionCount    int    `gorm:"column:prescription_count;NOT NULL" json:"prescription_count"`         // 开方数量
	PrescriptionFee      string `gorm:"column:prescription_fee;NOT NULL" json:"prescription_fee"`             // 处方加工费
	DrugFee              string `gorm:"column:drug_fee;NOT NULL" json:"drug_fee"`                             // 药品费用
	PrescriptionTotalFee string `gorm:"column:prescription_total_fee;NOT NULL" json:"prescription_total_fee"` // 总费用（非数据库字段，计算得出）
	DotTime              string `gorm:"column:dot_time;NOT NULL" json:"dot_time"`                             // 日期
}

func (l *HospitalData) TableName() string {
	return "hospital_tj_vw"
}
func (l *HospitalData) List(ctx context.Context, db *gorm.DB, params *dto.GetHosptialDataRequest) ([]*HospitalData, int64, error) {
	db = db.WithContext(ctx).Table(l.TableName())
	fmt.Print(params.DateEnd)
	if params.DateStart != "" && params.DateEnd != "" {
		db = db.Where("dot_time >= ? and dot_time <= ? ", params.DateStart, params.DateEnd)
	}
	var count int64
	db.Count(&count)
	var list []*HospitalData
	if params.Page > 0 && params.PageSize > 0 {
		db = db.Offset(app.GetPageOffset(int(params.Page), int(params.PageSize))).Limit(int(params.PageSize))
	}
	if err := db.Select(" hospital_name,sum(prescription_count) prescription_count,sum(prescription_fee) prescription_fee,sum(drug_fee) drug_fee,sum(prescription_total_fee) prescription_total_fee").Group("hospital_name").Scan(&list).Error; err != nil {
		return nil, 0, err
	}
	return list, count, nil
}

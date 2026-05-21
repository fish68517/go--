package model

import (
	"context"
	"fmt"

	"github.com/mwqnice/oh-admin/internal/dto"
	"github.com/mwqnice/oh-admin/pkg/app"
	"gorm.io/gorm"
)

// InspectionRecord 设备巡检记录表模型
type InspectionRecord struct {
	*Model
	EquipmentType      string `gorm:"column:equipment_type;type:varchar(255);default:null" json:"equipment_type"`  // 设备类型
	EquipmentID        string `gorm:"column:equipment_id;type:varchar(50);default:null" json:"equipment_id"`       // 设备编号
	HealthStatus       int    `gorm:"column:health_status;type:int;default:null" json:"health_status"`             // 卫生状态：1-良好，0-较差
	DisinfectionStatus int    `gorm:"column:disinfection_status;type:int;default:null" json:"disinfection_status"` // 消毒状态：1-已消毒，0-未消毒
	Status             int    `gorm:"column:status;type:int;default:null" json:"status"`                           // 运行状态：1-正常，0-异常
	InspectionTime     string `gorm:"column:inspection_time;type:datetime;default:null" json:"inspection_time"`    // 巡检时间
	Inspector          string `gorm:"column:inspector;type:varchar(255);default:null" json:"inspector"`            // 巡检人员// 忽略GORM软删除字段
}

// TableName 指定数据表名
func (InspectionRecord) TableName() string {
	return "inspection_record"
}
func (l *InspectionRecord) List(ctx context.Context, db *gorm.DB, params *dto.GetDeviceRecordListRequest) ([]*InspectionRecord, int64, error) {
	db = db.WithContext(ctx).Table(l.TableName())
	if params.Name != "" {
		db = db.Where("equipment_id like ?", fmt.Sprintf("%%%s%%", params.Name))
	}
	var count int64
	db.Count(&count)
	var list []*InspectionRecord
	if params.Page > 0 && params.PageSize > 0 {
		db = db.Offset(app.GetPageOffset(int(params.Page), int(params.PageSize))).Limit(int(params.PageSize))
	}
	if err := db.Order("id desc").Find(&list).Error; err != nil {
		return nil, 0, err
	}
	return list, count, nil
}

// Get 根据条件查询单条数据
func (l *InspectionRecord) Get(ctx context.Context, db *gorm.DB) (*InspectionRecord, error) {
	db = db.WithContext(ctx).Table(l.TableName())

	if l.Model != nil && l.ID != 0 {
		db = db.Where("id = ? ", l.ID)
	}
	var link *InspectionRecord
	if err := db.First(&link).Error; err != nil && err != gorm.ErrRecordNotFound {
		return link, err
	}
	return link, nil
}
func (l *InspectionRecord) Create(ctx context.Context, db *gorm.DB) error {
	return db.WithContext(ctx).Create(&l).Error
}

// Update 更新
func (l *InspectionRecord) Update(ctx context.Context, db *gorm.DB) error {
	return db.WithContext(ctx).Save(&l).Error
}

// Delete 删除
func (l *InspectionRecord) Delete(ctx context.Context, db *gorm.DB, id int) error {
	return db.WithContext(ctx).Where("id = ?", id).Delete(&l).Error
}

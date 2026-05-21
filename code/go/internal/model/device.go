package model

import (
	"context"
	"fmt"

	"github.com/mwqnice/oh-admin/internal/dto"
	"github.com/mwqnice/oh-admin/pkg/app"
	"gorm.io/gorm"
)

// Device 结构体映射设备信息表
type Device struct {
	*Model
	ID            int    `gorm:"primaryKey;column:id" json:"id"`
	EquipmentType string `gorm:"column:equipment_type;type:varchar(50)" json:"equipment_type"`
	DeviceName    string `gorm:"column:device_name;type:varchar(50)" json:"device_name"`
	DeviceRoom    string `gorm:"column:device_room;type:varchar(50)" json:"device_room"`
	UnitNumber    string `gorm:"column:unit_number;type:varchar(10)" json:"unit_number"`
	Remark        string `gorm:"column:remark" json:"remark"`
}

func (l *Device) TableName() string {
	return "device"
}
func (l *Device) List(ctx context.Context, db *gorm.DB, params *dto.DeviceListRequest) ([]*Device, int64, error) {
	db = db.WithContext(ctx).Table(l.TableName())
	if params.Name != "" {
		db = db.Where("device_name like ?", fmt.Sprintf("%%%s%%", params.Name))
	}
	var count int64
	db.Count(&count)
	var list []*Device
	if params.Page > 0 && params.PageSize > 0 {
		db = db.Offset(app.GetPageOffset(int(params.Page), int(params.PageSize))).Limit(int(params.PageSize))
	}
	if err := db.Order("id desc").Find(&list).Error; err != nil {
		return nil, 0, err
	}
	return list, count, nil
}

// Create 插入数据
func (l *Device) Create(ctx context.Context, db *gorm.DB) error {
	return db.WithContext(ctx).Create(&l).Error
}

// Update 更新
func (l *Device) Update(ctx context.Context, db *gorm.DB) error {
	return db.WithContext(ctx).Save(&l).Error
}

// Delete 删除
func (l *Device) Delete(ctx context.Context, db *gorm.DB, id int) error {
	return db.WithContext(ctx).Where("id = ?", id).Delete(&l).Error
}
func (l *Device) Get(ctx context.Context, db *gorm.DB) (*Device, error) {
	db = db.WithContext(ctx).Table(l.TableName())
	fmt.Print(l.Model.ID)
	if l.Model != nil && l.Model.ID != 0 {
		db = db.Where("id = ? ", l.Model.ID)
	}
	var device *Device
	if err := db.First(&device).Error; err != nil && err != gorm.ErrRecordNotFound {
		return device, err
	}
	return device, nil
}

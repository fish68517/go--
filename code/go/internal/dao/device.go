package dao

import (
	"context"

	"github.com/mwqnice/oh-admin/internal/dto"
	"github.com/mwqnice/oh-admin/internal/model"
)

func (d *Dao) GetDeviceListList(ctx context.Context, params *dto.DeviceListRequest) ([]*model.Device, int64, error) {
	device := model.Device{}
	return device.List(ctx, d.engine, params)
}
func (d *Dao) GetDeviceInfo(ctx context.Context, link *model.Device) (*model.Device, error) {
	return link.Get(ctx, d.engine)
}
func (d *Dao) CreateDevice(ctx context.Context, link *model.Device) error {
	return link.Create(ctx, d.engine)
}
func (d *Dao) UpdateDevice(ctx context.Context, link *model.Device) error {
	return link.Update(ctx, d.engine)
}

func (d *Dao) DeleteDevice(ctx context.Context, id int64) error {
	device := model.Device{}
	return device.Delete(ctx, d.engine, int(id))
}

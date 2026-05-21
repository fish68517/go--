package service

import (
	"context"
	"errors"
	"fmt"

	"github.com/mwqnice/oh-admin/internal/dto"
	"github.com/mwqnice/oh-admin/internal/model"
	"github.com/mwqnice/oh-admin/pkg/convert"
)

func (svc *Service) GetDeviceInfo(ctx context.Context, id int64) (*model.Device, error) {
	fmt.Println("==22", id)
	return svc.dao.GetDeviceInfo(ctx, &model.Device{Model: &model.Model{ID: int(id)}})
}
func (svc *Service) GetDeviceDataList(ctx context.Context, params *dto.DeviceListRequest) ([]*model.Device, int64, error) {
	return svc.dao.GetDeviceListList(ctx, params)
}
func (svc *Service) CreateDevice(ctx context.Context, params *dto.CreateDeviceRequest) error {
	device := &model.Device{
		EquipmentType: params.EquipmentType,
		DeviceName:    params.DeviceName,
		DeviceRoom:    params.DeviceRoom,
		UnitNumber:    params.UnitNumber,
		Remark:        params.Remark,
	}
	return svc.dao.CreateDevice(ctx, device)
}

func (svc *Service) UpdateDevice(ctx context.Context, params *dto.UpdateDeviceRequest) error {
	fmt.Printf("%#v\n", params)
	link, err := svc.dao.GetDeviceInfo(ctx, &model.Device{Model: &model.Model{ID: convert.Int(params.ID)}})
	if err != nil {
		return err
	}
	if link.Model == nil {
		return errors.New("该记录不存在")
	}
	link.EquipmentType = params.EquipmentType
	link.DeviceName = params.DeviceName
	link.DeviceRoom = params.DeviceRoom
	link.UnitNumber = params.UnitNumber
	link.Remark = params.Remark
	return svc.dao.UpdateDevice(ctx, link)
}

func (svc *Service) DeleteDevice(ctx context.Context, id string) error {
	return svc.dao.DeleteDevice(ctx, convert.Int64(id))
}

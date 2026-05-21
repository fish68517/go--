package service

import (
	"context"
	"errors"
	"fmt"

	"github.com/mwqnice/oh-admin/internal/dto"
	"github.com/mwqnice/oh-admin/internal/model"
	"github.com/mwqnice/oh-admin/pkg/convert"
)

func (svc *Service) GetDeviceRecordList(ctx context.Context, params *dto.GetDeviceRecordListRequest) ([]*model.InspectionRecord, int64, error) {
	return svc.dao.GetDeviceRecordList(ctx, params)
}
func (svc *Service) GetDeviceRecordInfo(ctx context.Context, id int64) (*model.InspectionRecord, error) {

	return svc.dao.GetDeviceRecordInfo(ctx, &model.InspectionRecord{Model: &model.Model{ID: int(id)}})
}
func (svc *Service) CreateDeviceRecord(ctx context.Context, params *dto.InspectionRecordCreateRequest) error {
	link := &model.InspectionRecord{
		EquipmentID:        params.EquipmentID,
		EquipmentType:      params.EquipmentType,
		HealthStatus:       convert.Int(params.HealthStatus),
		DisinfectionStatus: convert.Int(params.DisinfectionStatus),
		Status:             convert.Int(params.Status),
		Inspector:          params.Inspector,
		InspectionTime:     params.InspectionTime,
	}
	return svc.dao.CreateDeviceRecord(ctx, link)
}
func (svc *Service) UpdateDeviceRecord(ctx context.Context, params *dto.InspectionRecordUpdateRequest) error {
	fmt.Printf("%#v\n", params)
	link, err := svc.dao.GetDeviceRecordInfo(ctx, &model.InspectionRecord{Model: &model.Model{ID: convert.Int(params.ID)}})
	if err != nil {
		return err
	}
	if link.Model == nil {
		return errors.New("该记录不存在")
	}
	link.EquipmentType = params.EquipmentType
	link.EquipmentID = params.EquipmentID
	link.HealthStatus = convert.Int(params.HealthStatus)
	link.DisinfectionStatus = convert.Int(params.DisinfectionStatus)
	link.Status = convert.Int(params.Status)
	link.Inspector = params.Inspector
	link.InspectionTime = params.InspectionTime
	return svc.dao.UpdateDeviceRecord(ctx, link)
}
func (svc *Service) DeleteDeviceRecord(ctx context.Context, id string) error {
	return svc.dao.DeleteDeviceRecord(ctx, convert.Int64(id))
}

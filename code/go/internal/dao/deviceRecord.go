package dao

import (
	"context"

	"github.com/mwqnice/oh-admin/internal/dto"
	"github.com/mwqnice/oh-admin/internal/model"
)

func (d *Dao) GetDeviceRecordList(ctx context.Context, params *dto.GetDeviceRecordListRequest) ([]*model.InspectionRecord, int64, error) {
	link := model.InspectionRecord{}
	return link.List(ctx, d.engine, params)
}
func (d *Dao) CreateDeviceRecord(ctx context.Context, link *model.InspectionRecord) error {
	return link.Create(ctx, d.engine)
}

func (d *Dao) GetDeviceRecordInfo(ctx context.Context, link *model.InspectionRecord) (*model.InspectionRecord, error) {
	return link.Get(ctx, d.engine)
}

func (d *Dao) UpdateDeviceRecord(ctx context.Context, link *model.InspectionRecord) error {
	return link.Update(ctx, d.engine)
}

func (d *Dao) DeleteDeviceRecord(ctx context.Context, id int64) error {
	link := model.InspectionRecord{}
	return link.Delete(ctx, d.engine, int(id))
}

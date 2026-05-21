package dao

import (
	"context"
	"github.com/mwqnice/oh-admin/internal/dto"
	"github.com/mwqnice/oh-admin/internal/model"
)

func (d *Dao) GetHosptialDataList(ctx context.Context, params *dto.GetHosptialDataRequest) ([]*model.HospitalData, int64, error) {
	hd := model.HospitalData{}
	return hd.List(ctx, d.engine, params)
}

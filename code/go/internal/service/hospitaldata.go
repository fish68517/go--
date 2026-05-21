package service

import (
	"context"
	"github.com/mwqnice/oh-admin/internal/dto"
	"github.com/mwqnice/oh-admin/internal/model"
)

func (svc *Service) GetHosptialDataList(ctx context.Context, params *dto.GetHosptialDataRequest) ([]*model.HospitalData, int64, error) {
	return svc.dao.GetHosptialDataList(ctx, params)
}

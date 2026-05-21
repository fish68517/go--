package service

import (
	"context"
	"github.com/mwqnice/oh-admin/internal/dto"
	"github.com/mwqnice/oh-admin/internal/model"
	"time"
)

func (svc *Service) GetHerbDecoctionInfoList(ctx context.Context, params *dto.GetDecoInfoListRequest) ([]*model.DecoctionView, int64, error) {
	return svc.dao.GetDecList(ctx, params)
}
func (svc *Service) CreateDecoctionInfo(ctx context.Context, params *dto.PrescriptionFlowRequest) error {
	deco := &model.HerbalDecoctionInfo{
		WordContent:      "煎药",
		StartTime:        time.Now().Format("2006-01-02 15:04:05"),
		EndTime:          time.Now().Format("2006-01-02 15:04:05"),
		DecoctionManager: params.OperateName,
		EmployeeID:       params.EmployeeId,
		PrescriptionID:   params.PrescriptionID,
		DecoctionStatus:  0,
		Barcode:          params.Barcode,
	}
	return svc.dao.CreateHerbalDeco(ctx, deco)
}

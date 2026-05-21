package service

import (
	"context"
	"github.com/mwqnice/oh-admin/internal/dto"
	"github.com/mwqnice/oh-admin/internal/model"
	"time"
)

func (svc *Service) GetSoakList(ctx context.Context, params *dto.GetSoakListRequest) ([]*model.SoakView, int64, error) {
	return svc.dao.GetSoakList(ctx, params)
}

func (svc *Service) CreateSoak(ctx context.Context, params *dto.PrescriptionFlowRequest) error {
	soak := &model.HerbSoaking{
		WordContent:    "泡药",
		StartTime:      time.Now().Format("2006-01-02 15:04:05"),
		EndTime:        time.Now().Format("2006-01-02 15:04:05"),
		SoakingPerson:  params.OperateName,
		EmployeeID:     params.EmployeeId,
		PrescriptionID: params.PrescriptionID,
		WarningTime:    time.Now().Format("2006-01-02 15:04:05"),
		SoakingStatus:  0,
		Barcode:        params.Barcode,
	}
	return svc.dao.CreateSoak(ctx, soak)
}

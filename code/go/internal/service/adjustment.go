package service

import (
	"context"
	"github.com/mwqnice/oh-admin/internal/dto"
	"github.com/mwqnice/oh-admin/internal/model"
	"time"
)

func (svc *Service) GetAdjustmentList(ctx context.Context, params *dto.GetAdjustListRequest) ([]*model.AdjustmentView, int64, error) {
	return svc.dao.GetAdjustList(ctx, params)
}
func (svc *Service) CreateAdjust(ctx context.Context, params *dto.PrescriptionFlowRequest) error {
	link := &model.Adjustment{
		WordContent:    "调剂",
		WordDate:       time.Now().Format("2006-01-02 15:04:05"),
		WordPerson:     params.OperateName,
		EmployeeID:     params.EmployeeId,
		PrescriptionID: params.PrescriptionID,
		Status:         0,
		EndDate:        time.Now().Format("2006-01-02 15:04:05"),
		Barcode:        params.Barcode,
	}
	return svc.dao.CreateAdjust(ctx, link)
}

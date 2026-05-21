package service

import (
	"context"

	"github.com/mwqnice/oh-admin/internal/dto"
	"github.com/mwqnice/oh-admin/internal/model"
)

func (svc *Service) GetPrescriptionProcessList(ctx context.Context, params *dto.GetPrescriptionListRequest) ([]*model.Prescription, int64, error) {
	return svc.dao.GetPrescriptionList(ctx, params)
}

//	func (svc *Service) TraceChineseMedicine(ctx context.Context, rx_number string) ([]*model.PrescriptionView, error) {
//		return svc.dao.GetPrescriptionList(ctx, rx_number)
//	}
func (svc *Service) Create(ctx context.Context, params *dto.PrescriptionFlowRequest) error {

	return svc.dao.CreateLink(ctx, nil)
}

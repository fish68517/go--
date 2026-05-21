package service

import (
	"context"
	"github.com/mwqnice/oh-admin/internal/dto"
	"github.com/mwqnice/oh-admin/internal/model"
	"time"
)

func (svc *Service) GetDeliveryList(ctx context.Context, params *dto.GetDeliveryListRequest) ([]*model.HerbDeliveryView, int64, error) {
	return svc.dao.GetHerbDeliveryList(ctx, params)
}
func (svc *Service) CreateDeliver(ctx context.Context, params *dto.PrescriptionFlowRequest) error {
	link := &model.HerbDelivery{
		WordContent:          "发货",
		DeliveryTime:         time.Now().Format("2006-01-02 15:04:05"),
		DeliveryPersonnel:    params.OperateName,
		ProcessingEmployeeID: params.EmployeeId,
		PrescriptionID:       params.PrescriptionID,
		DeliveryStatus:       0,
		Barcode:              params.Barcode,
	}
	return svc.dao.CreateHerbDelivery(ctx, link)
}

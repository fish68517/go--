package service

import (
	"context"
	"github.com/mwqnice/oh-admin/internal/dto"
	"github.com/mwqnice/oh-admin/internal/model"
	"time"
)

func (svc *Service) GetPackList(ctx context.Context, params *dto.GetPackInfoListRequest) ([]*model.MedicinePackingView, int64, error) {
	return svc.dao.GetPackList(ctx, params)
}
func (svc *Service) CreatePack(ctx context.Context, params *dto.PrescriptionFlowRequest) error {
	pack := &model.MedicinePacking{
		WordContent:      "包装",
		StartTime:        time.Now().Format("2006-01-02 15:04:05"),
		PackingPersonnel: params.OperateName,
		EmployeeID:       params.EmployeeId,
		PrescriptionID:   params.PrescriptionID,
		PackingStatus:    0,
		Barcode:          params.Barcode,
		EndTime:          time.Now().Format("2006-01-02 15:04:05"),
	}
	return svc.dao.CreateMedicinePack(ctx, pack)
}

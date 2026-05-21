package service

import (
	"context"
	"github.com/mwqnice/oh-admin/internal/model"
)

func (svc *Service) GetPrescriptionSummaryList(ctx context.Context) (*model.PrescriptionSummary, error) {
	return svc.dao.GetPrescriptionSummaryList(ctx)
}
func (svc *Service) GetPrescriptionDosageCountList(ctx context.Context) ([]*model.PrescriptionDosageCount, error) {
	return svc.dao.GetPrescriptionDosageCountList(ctx)
}
func (svc *Service) GetPrescriptionDecoctionCountList(ctx context.Context) ([]*model.PrescriptionDecoctionCount, error) {
	return svc.dao.GetPrescriptionDecoctionCountList(ctx)
}
func (svc *Service) GetPrescriptionDetailList(ctx context.Context) ([]*model.PrescriptionDetail, error) {
	return svc.dao.GetPrescriptionDetailList(ctx)
}
func (svc *Service) GetDrugTotalQuantityList(ctx context.Context) ([]*model.DrugTotalQuantity, error) {
	return svc.dao.GetDrugTotalQuantityList(ctx)
}
func (svc *Service) GetPrescriptionCountByDeliveryDateList(ctx context.Context) ([]*model.PrescriptionCountByDeliveryDate, error) {
	return svc.dao.GetPrescriptionCountByDeliveryDateList(ctx)
}
func (svc *Service) GetPrescriptionCountByDateDateList(ctx context.Context) ([]*model.PrescriptionCountByDate, error) {
	return svc.dao.GetPrescriptionCountByDateList(ctx)
}
func (svc *Service) GetPrescriptionCountByHospital(ctx context.Context) ([]*model.PrescriptionCountByHospital, error) {
	return svc.dao.GetPrescriptionCountByHospital(ctx)
}

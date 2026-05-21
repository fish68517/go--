package dao

import (
	"context"
	"github.com/mwqnice/oh-admin/internal/model"
)

func (d *Dao) GetPrescriptionSummaryList(ctx context.Context) (*model.PrescriptionSummary, error) {
	result := model.PrescriptionSummary{}
	return result.List(ctx, d.engine)
}
func (d *Dao) GetPrescriptionDosageCountList(ctx context.Context) ([]*model.PrescriptionDosageCount, error) {
	result := model.PrescriptionDosageCount{}
	return result.List(ctx, d.engine)
}
func (d *Dao) GetPrescriptionDecoctionCountList(ctx context.Context) ([]*model.PrescriptionDecoctionCount, error) {
	result := model.PrescriptionDecoctionCount{}
	return result.List(ctx, d.engine)
}
func (d *Dao) GetPrescriptionDetailList(ctx context.Context) ([]*model.PrescriptionDetail, error) {
	result := model.PrescriptionDetail{}
	return result.List(ctx, d.engine)
}
func (d *Dao) GetDrugTotalQuantityList(ctx context.Context) ([]*model.DrugTotalQuantity, error) {
	result := model.DrugTotalQuantity{}
	return result.List(ctx, d.engine)
}
func (d *Dao) GetPrescriptionCountByDeliveryDateList(ctx context.Context) ([]*model.PrescriptionCountByDeliveryDate, error) {
	result := model.PrescriptionCountByDeliveryDate{}
	return result.List(ctx, d.engine)
}
func (d *Dao) GetPrescriptionCountByDateList(ctx context.Context) ([]*model.PrescriptionCountByDate, error) {
	result := model.PrescriptionCountByDate{}
	return result.List(ctx, d.engine)
}
func (d *Dao) GetPrescriptionCountByHospital(ctx context.Context) ([]*model.PrescriptionCountByHospital, error) {
	result := model.PrescriptionCountByHospital{}
	return result.List(ctx, d.engine)
}

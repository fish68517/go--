package dao

import (
	"context"

	"github.com/mwqnice/oh-admin/internal/dto"
	"github.com/mwqnice/oh-admin/internal/model"
)

func (d *Dao) GetPrescriptionViewList(ctx context.Context, params *dto.GetPrescriptionListRequest) ([]*model.PrescriptionView, int64, error) {
	prescription := model.PrescriptionView{}
	return prescription.PrescriptionViewList(ctx, d.engine, params)
}
func (d *Dao) GetPrescriptionView(ctx context.Context, params *dto.PrescriptionRequest) (*model.PrescriptionView, error) {
	prescription := model.PrescriptionView{}
	return prescription.GetPrescriptionView(ctx, d.engine, params)
}

func (d *Dao) GetPrescriptionPCViewList(ctx context.Context, params *dto.GetPrescriptionSearchListRequest) ([]*model.PrescriptionView, int64, error) {
	prescription := model.PrescriptionView{}
	return prescription.PrescriptionPCViewList(ctx, d.engine, params)
}

// GetPrescriptionList 获取处方信息列表
func (d *Dao) GetPrescriptionList(ctx context.Context, params *dto.GetPrescriptionListRequest) ([]*model.Prescription, int64, error) {
	prescription := model.Prescription{}
	return prescription.List(ctx, d.engine, params)
}

// GetWorkStatisticsList 获取工作量列表
func (d *Dao) GetWorkStatisticsList(ctx context.Context, params *dto.GetWorkStatistListRequest) ([]*model.WorkloadVw, int64, error) {
	workloadVw := model.WorkloadVw{}
	return workloadVw.WorkStatistViewList(ctx, d.engine, params)
}

// GetPrescriptionInfo 获取处方信息
func (d *Dao) GetPrescriptionInfo(ctx context.Context, pres *model.Prescription) (*model.Prescription, error) {
	return pres.Get(ctx, d.engine)
}
func (d *Dao) GetHosptialInfo(ctx context.Context, pres *model.HerbHospital) (*model.HerbHospital, error) {
	return pres.Get(ctx, d.engine)
}

func (d *Dao) UpdatePrescription(ctx context.Context, pres *model.Prescription) error {
	return pres.Update(ctx, d.engine)
}

func (d *Dao) DeletePrescription(ctx context.Context, id int64) error {
	prescription := model.Prescription{}
	return prescription.Delete(ctx, d.engine, int(id))
}

func (d *Dao) CreatePrescription(ctx context.Context, pres *model.Prescription) error {
	return pres.Create(ctx, d.engine)
}

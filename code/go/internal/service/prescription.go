package service

import (
	"context"
	"errors"
	"fmt"

	"github.com/mwqnice/oh-admin/internal/dto"
	"github.com/mwqnice/oh-admin/internal/model"
	"github.com/mwqnice/oh-admin/pkg/convert"
)

func (svc *Service) CreateHosptialPrescription(ctx context.Context, params *dto.PrescriptionDTO) error {
	return svc.dao.CreateHosptialPrescription(ctx, params)
}
func (svc *Service) PrescriptionDelete(ctx context.Context, params *dto.PrescriptionDeleteRequest) error {
	return svc.dao.PrescriptionDeleteDAO(ctx, params)
}
func (svc *Service) GetPrescriptionViewList(ctx context.Context, params *dto.GetPrescriptionListRequest) ([]*model.PrescriptionView, int64, error) {
	return svc.dao.GetPrescriptionViewList(ctx, params)
}
func (svc *Service) GetPrescriptionPCViewList(ctx context.Context, params *dto.GetPrescriptionSearchListRequest) ([]*model.PrescriptionView, int64, error) {
	return svc.dao.GetPrescriptionPCViewList(ctx, params)
}

// GetPrescriptionList 获取处方列表
func (svc *Service) GetPrescriptionList(ctx context.Context, params *dto.GetPrescriptionListRequest) ([]*model.Prescription, int64, error) {
	return svc.dao.GetPrescriptionList(ctx, params)
}

// GetWorkStatisticsList 获取工作量列表
func (svc *Service) GetWorkStatisticsList(ctx context.Context, params *dto.GetWorkStatistListRequest) ([]*model.WorkloadVw, int64, error) {
	return svc.dao.GetWorkStatisticsList(ctx, params)
}

// DeletePrescription 删除处方
func (svc *Service) DeletePrescription(ctx context.Context, id string) error {
	return svc.dao.DeletePrescription(ctx, convert.Int64(id))
}
func (svc *Service) CreatePrescription(ctx context.Context, params *dto.CreatePrescriptionRequest) error {
	prescription := &model.Prescription{
		HospitalName:         params.HospitalName,
		PatientAddress:       params.PatientAddress,
		PrescriptionType:     convert.Int(params.PrescriptionType),
		DoTime:               params.DoTime,
		DeletionNumber:       params.DeletionNumber,
		PrescriptionNumber:   params.PrescriptionNumber,
		PatientName:          params.PatientName,
		PatientSex:           convert.Int(params.PatientSex),
		PatientAge:           convert.Int(params.PatientAge),
		PatientPhone:         params.PatientPhone,
		DepartmentName:       params.DeletionNumber,
		DiagnosisResult:      params.PrescriptionNumber,
		Dosage:               convert.Int(params.Dosage),
		AdministrationCount:  convert.Int(params.AdministrationCount),
		AdministrationMethod: params.AdministrationMethod,
		SoakTime:             convert.Int(params.SoakTime),
		InpatientArea:        params.InpatientArea,
		WardName:             params.WardName,
		SickBed:              params.SickBed,
		PackageCount:         convert.Int(params.PackageCount),
		DecoctionScheme:      params.DecoctionScheme,
		IsDecoction:          convert.Int(params.IsDecoction),
		DoctorName:           params.DoctorName,
		Footnote:             params.Footnote,
		DrugPickupTime:       params.DrugPickupTime,
		SoakWaterAmount:      convert.Int(params.SoakWaterAmount),
		AdministrationWay:    params.DecoctionScheme,
		DrugPickupNumber:     params.DrugPickupNumber,
		Remarks:              params.Remarks,
		AdditionalRemarksA:   params.AdditionalRemarksA,
		AdditionalRemarksB:   params.DecoctionScheme,
		CurrentState:         "0",
	}

	return svc.dao.CreatePrescription(ctx, prescription)
}
func (svc *Service) GetPrescriptionState(ctx context.Context, params *dto.PrescriptionRequest) (string, error) {
	prescription, err := svc.dao.GetPrescriptionInfo(ctx, &model.Prescription{ID: convert.Int(params.Id)})
	if err != nil {
		return "", err
	}
	if prescription == nil {
		return "", errors.New("该记录不存在")
	}
	return prescription.CurrentState, nil
}
func (svc *Service) GetPrescriptionView(ctx context.Context, params *dto.PrescriptionRequest) (*model.PrescriptionView, error) {
	prescription, err := svc.dao.GetPrescriptionView(ctx, params)
	if err != nil {
		return prescription, err
	}
	return prescription, nil
}

// UpdatePrescription 更新处方
func (svc *Service) UpdatePrescription(ctx context.Context, params *dto.UpdatePrescriptionRequest) error {
	fmt.Printf("%#v\n", params)
	prescription, err := svc.dao.GetPrescriptionInfo(ctx, &model.Prescription{ID: convert.Int(params.Id)})
	if err != nil {
		return err
	}
	if prescription == nil {
		return errors.New("该记录不存在")
	}
	prescription.HospitalName = params.HospitalName
	prescription.PatientAddress = params.PatientAddress
	prescription.PrescriptionType = convert.Int(params.PrescriptionType)
	prescription.DoTime = params.DoTime
	prescription.DeletionNumber = params.DeletionNumber
	prescription.PrescriptionNumber = params.PrescriptionNumber
	prescription.PatientName = params.PatientName
	prescription.PatientSex = convert.Int(params.PatientSex)
	prescription.PatientAge = convert.Int(params.PatientAge)
	prescription.PatientAge = convert.Int(params.PatientAge)
	prescription.PatientPhone = params.PatientPhone
	prescription.DepartmentName = params.DeletionNumber
	prescription.DiagnosisResult = params.PrescriptionNumber
	prescription.Dosage = convert.Int(params.Dosage)
	prescription.AdministrationCount = convert.Int(params.AdministrationCount)
	prescription.AdministrationMethod = params.AdministrationMethod
	prescription.SoakTime = convert.Int(params.SoakTime)
	prescription.InpatientArea = params.InpatientArea
	prescription.WardName = params.WardName
	prescription.SickBed = params.SickBed
	prescription.PackageCount = convert.Int(params.PackageCount)
	prescription.DecoctionScheme = params.DecoctionScheme
	prescription.IsDecoction = convert.Int(params.IsDecoction)
	prescription.IsDecoction = convert.Int(params.IsDecoction)
	prescription.DoctorName = params.DoctorName
	prescription.Footnote = params.Footnote
	prescription.DrugPickupTime = params.DrugPickupTime
	prescription.SoakWaterAmount = convert.Int(params.SoakWaterAmount)
	prescription.AdministrationWay = params.DecoctionScheme
	prescription.DrugPickupNumber = params.DrugPickupNumber
	prescription.Remarks = params.Remarks
	prescription.AdditionalRemarksA = params.AdditionalRemarksA
	prescription.AdditionalRemarksB = params.DecoctionScheme
	prescription.CurrentState = "0"

	return svc.dao.UpdatePrescription(ctx, prescription)
}

// GetPrescriptionInfo 获取处方详情根据id
func (svc *Service) GetPrescriptionInfo(ctx context.Context, id int) (*model.Prescription, error) {
	return svc.dao.GetPrescriptionInfo(ctx, &model.Prescription{ID: id})
}

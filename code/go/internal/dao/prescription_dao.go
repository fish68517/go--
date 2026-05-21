package dao

import (
	"context"
	"time"

	"github.com/mwqnice/oh-admin/internal/dto"
	"github.com/mwqnice/oh-admin/internal/model"
	"github.com/mwqnice/oh-admin/pkg/convert"
	"gorm.io/gorm"
)

// CreatePrescription 创建处方
func (d *Dao) CreateHosptialPrescription(ctx context.Context, pres *dto.PrescriptionDTO) error {
	return d.PresCreate(ctx, d.engine, pres)
}
func (d *Dao) PrescriptionDeleteDAO(ctx context.Context, pres *dto.PrescriptionDeleteRequest) error {
	return d.PresDelete(ctx, d.engine, pres)
}

func (d *Dao) PresCreate(ctx context.Context, db *gorm.DB, params *dto.PrescriptionDTO) error {
	var hp *model.HerbHospital
	if err := db.Where("hospital_number = ? ", params.HospitalId).First(&hp).Error; err != nil && err != gorm.ErrRecordNotFound {
	}

	prescription := &model.Prescription{
		HospitalID:           params.HospitalId,
		HospitalName:         hp.HospitalName,
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
		CurrentState:         "接方",
	}
	var drug []dto.PrescriptionDrug

	drug = params.Drug
	tx := db.WithContext(ctx).Begin()
	// 插入调剂记录
	if err := tx.Create(&prescription).Error; err != nil {
		tx.Rollback()
		return err
	}

	for i := range drug {
		drug[i].PrescriptionID = prescription.ID
		drug[i].HospitalID = prescription.HospitalID
	}

	if err := tx.Create(&drug).Error; err != nil {
		tx.Rollback()
		return err
	}

	audit := &model.PrescriptionAudit{
		PrescriptionID: prescription.ID,
		CreateTime:     time.Now().Format("2006-01-02 15:04:05"),
	}
	if err := tx.Create(&audit).Error; err != nil {
		tx.Rollback()
		return err
	}
	// 提交事务
	tx.Commit()

	return nil
}
func (d *Dao) PresDelete(ctx context.Context, db *gorm.DB, params *dto.PrescriptionDeleteRequest) error {

	tx := db.WithContext(ctx).Begin()

	if err := tx.Where("ID = ?", params.Id).Delete(&model.Prescription{}).Error; err != nil {
		tx.Rollback()
		return err
	}
	if err := tx.Where("prescription_id = ?", params.Id).Delete(&model.PrescriptionDrug{}).Error; err != nil {
		tx.Rollback()
		return err
	}
	// 提交事务
	tx.Commit()
	return nil
}

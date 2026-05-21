package dao

import (
	"context"
	"strconv"
	"strings"
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
	var hp model.HerbHospital
	if err := db.Where("hospital_number = ? ", params.HospitalId).First(&hp).Error; err != nil && err != gorm.ErrRecordNotFound {
		return err
	}
	hospitalName := strings.TrimSpace(hp.HospitalName)
	if hospitalName == "" {
		hospitalName = params.HospitalName
	}
	doTime := strings.TrimSpace(params.DoTime)
	if doTime == "" {
		doTime = time.Now().Format("2006-01-02 15:04:05")
	}

	prescription := &model.Prescription{
		HospitalID:           params.HospitalId,
		HospitalName:         hospitalName,
		PatientAddress:       params.PatientAddress,
		PrescriptionType:     convert.Int(params.PrescriptionType),
		DoTime:               doTime,
		DeletionNumber:       params.DeletionNumber,
		PrescriptionNumber:   params.PrescriptionNumber,
		PatientName:          params.PatientName,
		PatientSex:           convert.Int(params.PatientSex),
		PatientAge:           convert.Int(params.PatientAge),
		PatientPhone:         params.PatientPhone,
		DepartmentName:       params.DepartmentName,
		DiagnosisResult:      params.DiagnosisResult,
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
		DrugCount:            len(params.Drug),
	}

	tx := db.WithContext(ctx).Begin()
	// 插入调剂记录
	if err := tx.Create(&prescription).Error; err != nil {
		tx.Rollback()
		return err
	}

	drugs := make([]model.PrescriptionDrug, 0, len(params.Drug))
	totalAmount := 0.0
	for i := range params.Drug {
		item := params.Drug[i]
		drug := model.PrescriptionDrug{
			PrescriptionID:         prescription.ID,
			HospitalID:             prescription.HospitalID,
			StockInItemID:          item.StockInItemID,
			PurchaseOrigin:         item.PurchaseOrigin,
			BatchNo:                item.BatchNo,
			DrugProductNumber:      item.DrugProductNumber,
			DrugProductName:        item.DrugProductName,
			MeasurementUnit:        item.MeasurementUnit,
			DrugProductDescription: item.DrugProductDescription,
			Dose:                   convert.Int(item.Dose),
			DrugWeight:             parseFloat64(item.DrugWeight),
			DrugWeights:            parseFloat64(item.DrugWeights),
			DrugPrice:              float64(item.DrugPrice),
			TotalPrices:            float64(item.TotalPrices),
		}
		if drug.MeasurementUnit == "" {
			drug.MeasurementUnit = "g"
		}
		if drug.StockInItemID > 0 {
			var stockItem model.StockInItem
			if err := tx.Where("id = ?", drug.StockInItemID).First(&stockItem).Error; err != nil {
				tx.Rollback()
				return err
			}
			if drug.PurchaseOrigin == "" {
				drug.PurchaseOrigin = stockItem.PurchaseOrigin
				if drug.PurchaseOrigin == "" {
					drug.PurchaseOrigin = extractOriginFromRemarkDAO(stockItem.Remark)
				}
			}
			if drug.BatchNo == "" {
				drug.BatchNo = stockItem.BatchNo
			}
			if drug.DrugPrice == 0 {
				drug.DrugPrice = parseFloat64(stockItem.UnitPrice)
			}
			if drug.DrugProductNumber == "" {
				drug.DrugProductNumber = stockItem.ProductId
			}
			if drug.DrugProductName == "" {
				drug.DrugProductName = stockItem.ProductName
			}
		}
		if drug.DrugWeights == 0 && drug.DrugWeight > 0 && drug.Dose > 0 {
			drug.DrugWeights = drug.DrugWeight * float64(drug.Dose)
		}
		if drug.TotalPrices == 0 {
			drug.TotalPrices = drug.DrugPrice * drug.DrugWeights
		}
		totalAmount += drug.TotalPrices
		drugs = append(drugs, drug)
	}

	if len(drugs) > 0 {
		if err := tx.Create(&drugs).Error; err != nil {
			tx.Rollback()
			return err
		}
	}

	payment := &model.PrescriptionPayment{
		PrescriptionID:     prescription.ID,
		PrescriptionNumber: prescription.PrescriptionNumber,
		TotalAmount:        totalAmount,
		PayAmount:          0,
		PayStatus:          "UNPAID",
	}
	if err := tx.Create(&payment).Error; err != nil {
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

func parseFloat64(value string) float64 {
	number, _ := strconv.ParseFloat(strings.TrimSpace(value), 64)
	return number
}

func extractOriginFromRemarkDAO(remark string) string {
	remark = strings.TrimSpace(remark)
	if remark == "" {
		return ""
	}
	start := strings.LastIndex(remark, "（")
	end := strings.LastIndex(remark, "）")
	if start >= 0 && end > start {
		return strings.TrimSpace(remark[start+len("（") : end])
	}
	start = strings.LastIndex(remark, "(")
	end = strings.LastIndex(remark, ")")
	if start >= 0 && end > start {
		return strings.TrimSpace(remark[start+1 : end])
	}
	return remark
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
	if err := tx.Where("prescription_id = ?", params.Id).Delete(&model.PrescriptionPayment{}).Error; err != nil {
		tx.Rollback()
		return err
	}
	if err := tx.Where("prescription_id = ?", params.Id).Delete(&model.BlockchainTraceLog{}).Error; err != nil {
		tx.Rollback()
		return err
	}
	// 提交事务
	tx.Commit()
	return nil
}

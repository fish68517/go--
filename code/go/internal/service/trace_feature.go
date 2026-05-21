package service

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/mwqnice/oh-admin/internal/dto"
	"github.com/mwqnice/oh-admin/internal/model"
)

func (svc *Service) GetDrugSources(ctx context.Context, productID, productName string) ([]dto.DrugSourceResponse, error) {
	items, err := svc.dao.GetDrugSources(ctx, strings.TrimSpace(productID), strings.TrimSpace(productName))
	if err != nil {
		return nil, err
	}
	sources := make([]dto.DrugSourceResponse, 0, len(items))
	for _, item := range items {
		purchaseOrigin := strings.TrimSpace(item.PurchaseOrigin)
		if purchaseOrigin == "" {
			purchaseOrigin = extractOriginFromRemark(item.Remark)
		}
		sources = append(sources, dto.DrugSourceResponse{
			StockInItemID:  item.Id,
			ProductID:      item.ProductId,
			ProductName:    item.ProductName,
			PurchaseOrigin: purchaseOrigin,
			SupplierName:    item.SupplierName,
			BatchNo:        item.BatchNo,
			UnitPrice:      item.UnitPrice,
			Quantity:       item.Quantity,
			Remark:         item.Remark,
		})
	}
	return sources, nil
}

func (svc *Service) GetPrescriptionPaymentDetail(ctx context.Context, prescriptionID int) (*dto.PaymentDetailResponse, error) {
	prescription, err := svc.dao.GetPrescriptionInfo(ctx, &model.Prescription{ID: prescriptionID})
	if err != nil {
		return nil, err
	}
	if prescription == nil || prescription.ID == 0 {
		return nil, fmt.Errorf("处方不存在")
	}
	drugs, err := svc.dao.GetPrescriptionDrugInfo(ctx, &model.PrescriptionDrug{PrescriptionID: prescriptionID})
	if err != nil {
		return nil, err
	}
	payment, err := svc.dao.GetPrescriptionPaymentByPrescriptionID(ctx, prescriptionID)
	if err != nil {
		return nil, err
	}
	return buildPaymentDetail(prescription, drugs, payment), nil
}

func (svc *Service) MockPayPrescription(ctx context.Context, params *dto.MockPayRequest) (*dto.PaymentDetailResponse, error) {
	if params.PayMethod == "" {
		params.PayMethod = "MOCK"
	}
	detail, err := svc.GetPrescriptionPaymentDetail(ctx, params.PrescriptionID)
	if err != nil {
		return nil, err
	}
	payment, err := svc.dao.GetPrescriptionPaymentByPrescriptionID(ctx, params.PrescriptionID)
	if err != nil {
		return nil, err
	}
	now := time.Now()
	if payment == nil {
		payment = &model.PrescriptionPayment{
			PrescriptionID:     detail.PrescriptionID,
			PrescriptionNumber: detail.PrescriptionNumber,
		}
	}
	if payment.PayStatus != "PAID" {
		payment.TotalAmount = detail.TotalAmount
		payment.PayAmount = detail.TotalAmount
		payment.PayMethod = params.PayMethod
		payment.PayStatus = "PAID"
		payment.MockTradeNo = fmt.Sprintf("MOCKPAY%s%03d", now.Format("20060102150405"), payment.PrescriptionID)
		payment.PaidAt = &now
		if err := svc.dao.SavePrescriptionPayment(ctx, payment); err != nil {
			return nil, err
		}
	}
	return svc.GetPrescriptionPaymentDetail(ctx, params.PrescriptionID)
}

func (svc *Service) BuildTraceMedicineItems(ctx context.Context, prescriptionID int) ([]dto.TraceMedicineItem, error) {
	drugs, err := svc.dao.GetPrescriptionDrugInfo(ctx, &model.PrescriptionDrug{PrescriptionID: prescriptionID})
	if err != nil {
		return nil, err
	}
	items := make([]dto.TraceMedicineItem, 0, len(drugs))
	for _, drug := range drugs {
		totalPrice := drug.TotalPrices
		if totalPrice == 0 {
			totalPrice = drug.DrugPrice * drug.DrugWeights
		}
		items = append(items, dto.TraceMedicineItem{
			DrugProductNumber: drug.DrugProductNumber,
			DrugProductName:   drug.DrugProductName,
			PurchaseOrigin:    drug.PurchaseOrigin,
			BatchNo:           drug.BatchNo,
			Quantity:          drug.DrugWeights,
			UnitPrice:         drug.DrugPrice,
			TotalPrice:        totalPrice,
		})
	}
	return items, nil
}

func (svc *Service) BuildTracePayment(ctx context.Context, prescriptionID int) (dto.TracePayment, error) {
	detail, err := svc.GetPrescriptionPaymentDetail(ctx, prescriptionID)
	if err != nil {
		return dto.TracePayment{}, err
	}
	return dto.TracePayment{
		PayStatus:   detail.PayStatus,
		TotalAmount: detail.TotalAmount,
		PayAmount:   detail.PayAmount,
		PayMethod:   detail.PayMethod,
		MockTradeNo: detail.MockTradeNo,
		PaidAt:      detail.PaidAt,
	}, nil
}

func (svc *Service) BuildTraceBlockchain(ctx context.Context, prescriptionID int) (dto.TraceBlockchain, error) {
	chainLog, err := svc.dao.GetLatestBlockchainTraceLog(ctx, prescriptionID)
	if err != nil || chainLog == nil {
		return dto.TraceBlockchain{}, err
	}
	return dto.TraceBlockchain{
		ChainStatus: chainLog.ChainStatus,
		TxID:        chainLog.TxID,
		PayloadHash: chainLog.PayloadHash,
		Error:       chainLog.ErrorMessage,
	}, nil
}

func (svc *Service) SaveBlockchainTraceLog(ctx context.Context, chainLog *model.BlockchainTraceLog) error {
	return svc.dao.SaveBlockchainTraceLog(ctx, chainLog)
}

func buildPaymentDetail(prescription *model.Prescription, drugs []*model.PrescriptionDrug, payment *model.PrescriptionPayment) *dto.PaymentDetailResponse {
	items := make([]dto.PaymentItem, 0, len(drugs))
	totalAmount := 0.0
	for _, drug := range drugs {
		amount := drug.TotalPrices
		if amount == 0 {
			amount = drug.DrugPrice * drug.DrugWeights
		}
		totalAmount += amount
		items = append(items, dto.PaymentItem{
			DrugProductNumber: drug.DrugProductNumber,
			DrugProductName:   drug.DrugProductName,
			PurchaseOrigin:    drug.PurchaseOrigin,
			BatchNo:           drug.BatchNo,
			Quantity:          drug.DrugWeights,
			UnitPrice:         drug.DrugPrice,
			Amount:            amount,
		})
	}
	result := &dto.PaymentDetailResponse{
		PrescriptionID:     prescription.ID,
		PrescriptionNumber: prescription.PrescriptionNumber,
		PayStatus:          "UNPAID",
		TotalAmount:        totalAmount,
		Items:              items,
	}
	if payment == nil {
		return result
	}
	result.PayStatus = payment.PayStatus
	result.TotalAmount = payment.TotalAmount
	if result.TotalAmount == 0 {
		result.TotalAmount = totalAmount
	}
	result.PayAmount = payment.PayAmount
	result.PayMethod = payment.PayMethod
	result.MockTradeNo = payment.MockTradeNo
	if payment.PaidAt != nil {
		result.PaidAt = payment.PaidAt.Format("2006-01-02 15:04:05")
	}
	return result
}

func extractOriginFromRemark(remark string) string {
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

package service

import (
	"context"
	"github.com/mwqnice/oh-admin/internal/model"
)

// GetPrescriptionInfo 获取药品详情
func (svc *Service) GetPrescriptionDrugInfo(ctx context.Context, id int) ([]*model.PrescriptionDrug, error) {
	return svc.dao.GetPrescriptionDrugInfo(ctx, &model.PrescriptionDrug{PrescriptionID: id})
}
func (svc *Service) GetQrcodeVwInfo(ctx context.Context, id int) (*model.QrcodeVw, error) {
	return svc.dao.GetQrcodeVw(ctx, &model.QrcodeVw{ID: id})
}

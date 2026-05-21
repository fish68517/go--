package dao

import (
	"context"
	"github.com/mwqnice/oh-admin/internal/dto"
	"github.com/mwqnice/oh-admin/internal/model"
)

// GetPrescriptionAuditList 获取审核信息列表
func (d *Dao) GetPrescriptionAuditList(ctx context.Context, params *dto.GetPrescriptionAuditListRequest) ([]*model.PrescriptionAuditView, int64, error) {
	prescription := model.PrescriptionAuditView{}
	return prescription.List(ctx, d.engine, params)
}

// PrescriptionAudit 更新审核表
func (d *Dao) PrescriptionAudit(ctx context.Context, audit *model.PrescriptionAudit) error {
	return audit.Update(ctx, d.engine)
}

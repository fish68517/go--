package service

import (
	"context"
	"time"

	"github.com/mwqnice/oh-admin/internal/dto"
	"github.com/mwqnice/oh-admin/internal/model"
)

func (svc *Service) GetPrescriptionAduitList(ctx context.Context, params *dto.GetPrescriptionAuditListRequest) ([]*model.PrescriptionAuditView, int64, error) {
	return svc.dao.GetPrescriptionAuditList(ctx, params)
}

// AuditPrescription 审核处方
func (svc *Service) AuditPrescription(ctx context.Context, id int, employId int, username string) error {
	var PrescriptionAudit model.PrescriptionAudit
	PrescriptionAudit.PrescriptionID = id
	PrescriptionAudit.EmployeeID = employId
	PrescriptionAudit.Reviewer = username
	PrescriptionAudit.AuditStatus = 1
	PrescriptionAudit.ReviewTime = time.Now().Format("2006-01-02 15:04:05")
	return svc.dao.PrescriptionAudit(ctx, &PrescriptionAudit)
}
func (svc *Service) AuditFailePrescription(ctx context.Context, id int, employId int, username string) error {
	var PrescriptionAudit model.PrescriptionAudit
	PrescriptionAudit.PrescriptionID = id
	PrescriptionAudit.EmployeeID = employId
	PrescriptionAudit.Reviewer = username
	PrescriptionAudit.AuditStatus = 2
	PrescriptionAudit.ReviewTime = time.Now().Format("2006-01-02 15:04:05")
	return svc.dao.PrescriptionAudit(ctx, &PrescriptionAudit)
}

//查询是否超剂量

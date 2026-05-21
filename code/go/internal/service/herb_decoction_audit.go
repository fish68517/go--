package service

import (
	"context"
	"github.com/mwqnice/oh-admin/internal/dto"
	"github.com/mwqnice/oh-admin/internal/model"
	"time"
)

func (svc *Service) GetHerbDecoctionAuditmentList(ctx context.Context, params *dto.GetAuditListRequest) ([]*model.AuditView, int64, error) {
	return svc.dao.GetHerbDecoctionAuditList(ctx, params)
}
func (svc *Service) CreateHerbDecoctionAudit(ctx context.Context, params *dto.PrescriptionFlowRequest) error {
	audit := &model.HerbDecoctionAudit{
		WordContent:    "复核",
		AuditDatetime:  time.Now().Format("2006-01-02 15:04:05"),
		Reviewer:       params.OperateName,
		EmployeeID:     params.EmployeeId,
		PrescriptionID: params.PrescriptionID,
		AuditStatus:    0,

		Barcode: params.Barcode,
	}
	return svc.dao.CreateAudit(ctx, audit)
}

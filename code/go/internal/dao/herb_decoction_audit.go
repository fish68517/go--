package dao

import (
	"context"
	"github.com/mwqnice/oh-admin/internal/dto"
	"github.com/mwqnice/oh-admin/internal/model"
)

func (d *Dao) GetHerbDecoctionAuditList(ctx context.Context, params *dto.GetAuditListRequest) ([]*model.AuditView, int64, error) {
	adjust := model.AuditView{}
	return adjust.List(ctx, d.engine, params)
}
func (d *Dao) CreateAudit(ctx context.Context, aduit *model.HerbDecoctionAudit) error {
	return aduit.Create(ctx, d.engine)
}

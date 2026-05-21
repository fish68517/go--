package dao

import (
	"context"
	"github.com/mwqnice/oh-admin/internal/dto"
	"github.com/mwqnice/oh-admin/internal/model"
)

func (d *Dao) GetAdjustList(ctx context.Context, params *dto.GetAdjustListRequest) ([]*model.AdjustmentView, int64, error) {
	adjust := model.AdjustmentView{}
	return adjust.List(ctx, d.engine, params)
}
func (d *Dao) CreateAdjust(ctx context.Context, adjust *model.Adjustment) error {
	return adjust.Create(ctx, d.engine)
}

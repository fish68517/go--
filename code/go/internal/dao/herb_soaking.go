package dao

import (
	"context"
	"github.com/mwqnice/oh-admin/internal/dto"
	"github.com/mwqnice/oh-admin/internal/model"
)

func (d *Dao) GetSoakList(ctx context.Context, params *dto.GetSoakListRequest) ([]*model.SoakView, int64, error) {
	adjust := model.SoakView{}
	return adjust.List(ctx, d.engine, params)
}
func (d *Dao) CreateSoak(ctx context.Context, soak *model.HerbSoaking) error {
	return soak.Create(ctx, d.engine)
}

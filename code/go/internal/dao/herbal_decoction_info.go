package dao

import (
	"context"
	"github.com/mwqnice/oh-admin/internal/dto"
	"github.com/mwqnice/oh-admin/internal/model"
)

func (d *Dao) GetDecList(ctx context.Context, params *dto.GetDecoInfoListRequest) ([]*model.DecoctionView, int64, error) {
	adjust := model.DecoctionView{}
	return adjust.List(ctx, d.engine, params)
}
func (d *Dao) CreateHerbalDeco(ctx context.Context, deco *model.HerbalDecoctionInfo) error {
	return deco.Create(ctx, d.engine)
}

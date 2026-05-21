package dao

import (
	"context"
	"github.com/mwqnice/oh-admin/internal/dto"
	"github.com/mwqnice/oh-admin/internal/model"
)

func (d *Dao) GetHerbDeliveryList(ctx context.Context, params *dto.GetDeliveryListRequest) ([]*model.HerbDeliveryView, int64, error) {
	adjust := model.HerbDeliveryView{}
	return adjust.List(ctx, d.engine, params)
}
func (d *Dao) CreateHerbDelivery(ctx context.Context, delivery *model.HerbDelivery) error {
	return delivery.Create(ctx, d.engine)
}

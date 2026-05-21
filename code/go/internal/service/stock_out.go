package service

import (
	"context"

	"github.com/mwqnice/oh-admin/internal/dto"
	"github.com/mwqnice/oh-admin/internal/model"
)

func (svc *Service) StockOutList(ctx context.Context, params *dto.GetStockOutListRequest) ([]*model.StockOut, int64, error) {
	return svc.dao.GetStockOutList(ctx, params)
}
func (svc *Service) GetStockOutItem(ctx context.Context, id string) ([]*model.StockOutItem, error) {
	return svc.dao.GetStockOutItemInfo(ctx, &model.StockOutItem{StockOutID: id})
}

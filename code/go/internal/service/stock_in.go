package service

import (
	"context"

	"github.com/mwqnice/oh-admin/internal/dto"
	"github.com/mwqnice/oh-admin/internal/model"
)

func (svc *Service) StockInList(ctx context.Context, params *dto.GetStockInListRequest) ([]*model.StockIn, int64, error) {
	return svc.dao.GetStockInList(ctx, params)
}
func (svc *Service) CreateStockIn(ctx context.Context, params *dto.StockInCreateDTO) error {
	return svc.dao.CreateStockIn(ctx, params)
}
func (svc *Service) GetStockInItem(ctx context.Context, id string) ([]*model.StockInItem, error) {
	return svc.dao.GetStockInItemInfo(ctx, &model.StockInItem{StockInId: id})
}

package service

import (
	"context"

	"github.com/mwqnice/oh-admin/internal/dto"
	"github.com/mwqnice/oh-admin/internal/model"
)

func (svc *Service) StockCheckList(ctx context.Context, params *dto.GetStockCheckListRequest) ([]*model.StockCheck, int64, error) {
	return svc.dao.GetStockCheckList(ctx, params)
}
func (svc *Service) CreateStockCheck(ctx context.Context, params *dto.StockCheckCreateDTO) error {
	return svc.dao.CreateStockCheck(ctx, params)
}
func (svc *Service) GetStockCheckItem(ctx context.Context, id string) ([]*model.StockCheckItem, error) {
	return svc.dao.GetStockCheckItemInfo(ctx, &model.StockCheckItem{CheckID: id})
}

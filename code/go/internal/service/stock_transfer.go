package service

import (
	"context"

	"github.com/mwqnice/oh-admin/internal/dto"
	"github.com/mwqnice/oh-admin/internal/model"
)

func (svc *Service) StockTransferList(ctx context.Context, params *dto.GetStockTransferListRequest) ([]*model.StockTransfer, int64, error) {
	return svc.dao.GetStockTransferInList(ctx, params)
}
func (svc *Service) CreateStockTransfer(ctx context.Context, params *dto.StockTransferCreateDTO) error {
	return svc.dao.CreatetSockTransfer(ctx, params)
}
func (svc *Service) GetStockTransferItem(ctx context.Context, id string) ([]*model.StockTransferItem, error) {

	return svc.dao.GetStockTransferItemInfo(ctx, &model.StockTransferItem{TransferID: id})
}

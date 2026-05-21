package service

import (
	"context"
	"fmt"

	"github.com/mwqnice/oh-admin/internal/dto"
	"github.com/mwqnice/oh-admin/internal/model"
)

func (svc *Service) InventoryList(ctx context.Context, params *dto.GetInventoryListRequest) ([]*model.Inventory, int64, error) {
	return svc.dao.GetInventoryList(ctx, params)
}
func (svc *Service) GetInventoryInfo(ctx context.Context, id int64) (*model.Inventory, error) {
	fmt.Println("==22", id)
	return svc.dao.GetInventoryInfo(ctx, &model.Inventory{Model: &model.Model{ID: int(id)}})
}
func (svc *Service) CreateInventory(ctx context.Context, params *dto.InventoryOperationDTO) error {
	return svc.dao.CreateStockOut(ctx, params)
}

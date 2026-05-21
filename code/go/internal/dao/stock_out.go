package dao

import (
	"context"

	"github.com/mwqnice/oh-admin/internal/dto"
	"github.com/mwqnice/oh-admin/internal/model"
)

func (d *Dao) GetStockOutList(ctx context.Context, params *dto.GetStockOutListRequest) ([]*model.StockOut, int64, error) {
	stockIn := model.StockOut{}
	return stockIn.List(ctx, d.engine, params)
}
func (d *Dao) GetStockOutItemInfo(ctx context.Context, drug *model.StockOutItem) ([]*model.StockOutItem, error) {
	return drug.Get(ctx, d.engine)
}

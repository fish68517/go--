package dao

import (
	"context"
	"github.com/mwqnice/oh-admin/internal/dto"
	"github.com/mwqnice/oh-admin/internal/model"
)

func (d *Dao) GetPackList(ctx context.Context, params *dto.GetPackInfoListRequest) ([]*model.MedicinePackingView, int64, error) {
	adjust := model.MedicinePackingView{}
	return adjust.List(ctx, d.engine, params)
}
func (d *Dao) CreateMedicinePack(ctx context.Context, pack *model.MedicinePacking) error {
	return pack.Create(ctx, d.engine)
}

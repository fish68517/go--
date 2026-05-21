package dao

import (
	"context"

	"github.com/mwqnice/oh-admin/internal/dto"
	"github.com/mwqnice/oh-admin/internal/model"
)

// GetLinkList 获取菜单列表
func (d *Dao) GetTDrugDosageLimitList(ctx context.Context, params *dto.GetDrugDosageListRequest) ([]*model.TDrugDosageLimit, int64, error) {
	link := model.TDrugDosageLimit{}
	return link.List(ctx, d.engine, params)
}

func (d *Dao) CreateTDrugDosageLimit(ctx context.Context, link *model.TDrugDosageLimit) error {
	return link.Create(ctx, d.engine)
}

func (d *Dao) GetTDrugDosageLimitInfo(ctx context.Context, link *model.TDrugDosageLimit) (*model.TDrugDosageLimit, error) {
	return link.Get(ctx, d.engine)
}
func (d *Dao) GetSingleTDrugDosageLimitInfo(ctx context.Context, link *model.TDrugDosageLimit, param string) (*model.TDrugDosageLimit, error) {
	return link.GetSingleDosageLimit(ctx, d.engine, param)
}

func (d *Dao) UpdateTDrugDosageLimit(ctx context.Context, link *model.TDrugDosageLimit) error {
	return link.Update(ctx, d.engine)
}

func (d *Dao) DeleteTDrugDosageLimit(ctx context.Context, id int64) error {
	link := model.TDrugDosageLimit{}
	return link.Delete(ctx, d.engine, int(id))
}

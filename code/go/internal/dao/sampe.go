package dao

import (
	"context"
	"github.com/mwqnice/oh-admin/internal/dto"
	"github.com/mwqnice/oh-admin/internal/model"
)

func (d *Dao) GetSampeList(ctx context.Context, params *dto.GetSampleListRequest) ([]*model.SamInfo, int64, error) {
	sampe := model.SamInfo{}
	return sampe.List(ctx, d.engine, params)
}
func (d *Dao) CreateSampe(ctx context.Context, samp *model.SamInfo) error {
	return samp.Create(ctx, d.engine)
}
func (d *Dao) GetSampeListById(ctx context.Context, sam *model.SamInfo) (*model.SamInfo, error) {
	return sam.Get(ctx, d.engine)
}
func (d *Dao) UpdateSampe(ctx context.Context, sam *model.SamInfo) error {
	return sam.Update(ctx, d.engine)
}

// DeletePrescription 删除处方信息
func (d *Dao) DeleteSampe(ctx context.Context, id int64) error {
	sampe := model.SamInfo{}
	return sampe.Delete(ctx, d.engine, int(id))
}

package dao

import (
	"context"

	"github.com/mwqnice/oh-admin/internal/dto"
	"github.com/mwqnice/oh-admin/internal/model"
)

func (d *Dao) CreateTisaneRecord(ctx context.Context, link *model.Link) error {
	return link.Create(ctx, d.engine)
}

func (d *Dao) GeTisaneView(ctx context.Context, params *dto.PrescriptionRequest) (*model.PrescriptionView, error) {
	prescription := model.PrescriptionView{}
	return prescription.GetPrescriptionView(ctx, d.engine, params)
}

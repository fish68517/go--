package dao

import (
	"context"
	"github.com/mwqnice/oh-admin/internal/model"
)

func (d *Dao) GetPrescriptionDrugInfo(ctx context.Context, drug *model.PrescriptionDrug) ([]*model.PrescriptionDrug, error) {
	return drug.Get(ctx, d.engine)
}
func (d *Dao) GetQrcodeVw(ctx context.Context, qr *model.QrcodeVw) (*model.QrcodeVw, error) {
	return qr.GetQrcodeVw(ctx, d.engine)
}

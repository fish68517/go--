package service

import (
	"context"
	"errors"
	"fmt"

	"github.com/mwqnice/oh-admin/internal/dto"
	"github.com/mwqnice/oh-admin/internal/model"
	"github.com/mwqnice/oh-admin/pkg/convert"
)

func (svc *Service) GetDrugDosageList(ctx context.Context, params *dto.GetDrugDosageListRequest) ([]*model.TDrugDosageLimit, int64, error) {
	return svc.dao.GetTDrugDosageLimitList(ctx, params)
}
func (svc *Service) GetDrugDosageInfo(ctx context.Context, id int64) (*model.TDrugDosageLimit, error) {

	return svc.dao.GetTDrugDosageLimitInfo(ctx, &model.TDrugDosageLimit{Model: &model.Model{ID: int(id)}})
}
func (svc *Service) GetSingleDrugDosageInfo(ctx context.Context, id int64, param string) (*model.TDrugDosageLimit, error) {

	return svc.dao.GetSingleTDrugDosageLimitInfo(ctx, &model.TDrugDosageLimit{Model: &model.Model{ID: int(id)}}, param)
}

func (svc *Service) CreateTDrugDosage(ctx context.Context, params *dto.CreateDrugDosageRequest) error {
	link := &model.TDrugDosageLimit{
		DrugID:     params.DrugID,
		MaxDosage:  convert.Float64(params.MaxDosage),
		DosageDesc: params.DosageDesc,
		IsEnabled:  convert.Int(params.IsEnabled),
	}
	return svc.dao.CreateTDrugDosageLimit(ctx, link)
}
func (svc *Service) UpdateTDrugDosage(ctx context.Context, params *dto.UpdateDrugDosageRequest) error {

	link, err := svc.dao.GetSingleTDrugDosageLimitInfo(ctx, &model.TDrugDosageLimit{Model: &model.Model{ID: convert.Int(params.Id)}}, "")

	if err != nil {
		return err
	}
	if link.Model == nil {
		return errors.New("该记录不存在")
	}
	fmt.Print(link)
	link.DosageDesc = params.DosageDesc
	link.MaxDosage = convert.Float64(params.MaxDosage)
	link.IsEnabled = convert.Int(params.IsEnabled)
	link.ID = convert.Int(params.Id)
	return svc.dao.UpdateTDrugDosageLimit(ctx, link)

}
func (svc *Service) DeleteTDrugDosage(ctx context.Context, id string) error {
	return svc.dao.DeleteTDrugDosageLimit(ctx, convert.Int64(id))
}

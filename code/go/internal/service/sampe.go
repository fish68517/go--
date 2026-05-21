package service

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/mwqnice/oh-admin/internal/dto"
	"github.com/mwqnice/oh-admin/internal/model"
	"github.com/mwqnice/oh-admin/pkg/convert"
)

func (svc *Service) GetSampeList(ctx context.Context, params *dto.GetSampleListRequest) ([]*model.SamInfo, int64, error) {
	return svc.dao.GetSampeList(ctx, params)
}
func (svc *Service) CreateSampe(ctx context.Context, params *dto.SampleDTO) error {
	same := &model.SamInfo{}
	convert.SwapTo(params, &same)
	return svc.dao.CreateSampe(ctx, same)
}
func (svc *Service) CreateAppSampe(ctx context.Context, params *dto.SampleAppDTO) error {
	fmt.Print(params.PrescriptionWeight, params.PrescriptionSJWeight)
	same := &model.SamInfo{
		PrescriptionNumber:   params.PrescriptionNumber,
		PrescriptionSjWeight: string(params.PrescriptionSJWeight),
		PrescriptionWeight:   string(params.PrescriptionWeight),
		DrugCount:            params.DrugCount,
		DrugSjCount:          params.DrugSJCount,
		Dosage:               params.Dosage,
		OperateName:          params.OperateName,
		OperateTime:          params.OperateTime,
		Remark:               params.Remark,
	}
	convert.SwapTo(params, &same)
	return svc.dao.CreateSampe(ctx, same)
}
func (svc *Service) UpdateSampe(ctx context.Context, params *dto.UpdateSampleRequest) error {
	same := &model.SamInfo{}
	same.ID = convert.Int(params.Id)
	sameInfo, err := svc.dao.GetSampeListById(ctx, same)
	if err != nil {
		return err
	}
	if sameInfo == nil {
		return errors.New("该记录不存在")
	}
	sameInfo.PrescriptionNumber = params.PrescriptionNumber
	sameInfo.PrescriptionWeight = params.PrescriptionWeight
	sameInfo.PrescriptionSjWeight = params.PrescriptionSJWeight
	sameInfo.DrugCount = convert.Int(params.DrugCount)
	sameInfo.DrugSjCount = convert.Int(params.DrugSJCount)
	sameInfo.Status = convert.Int(params.Status)
	sameInfo.Remark = params.Remark
	sameInfo.OperateUpdateTime = time.Now().Format("2006-01-02 15:04:05")
	sameInfo.OperateName = params.OperateName
	return svc.dao.UpdateSampe(ctx, sameInfo)
}
func (svc *Service) GetSampeListById(ctx context.Context, id int) (*model.SamInfo, error) {
	return svc.dao.GetSampeListById(ctx, &model.SamInfo{ID: id})
}
func (svc *Service) DeleteSampe(ctx context.Context, id string) error {
	return svc.dao.DeleteSampe(ctx, convert.Int64(id))
}

package service

import (
	"context"
	"errors"
	"fmt"

	"github.com/mwqnice/oh-admin/internal/dto"
	"github.com/mwqnice/oh-admin/internal/model"
	"github.com/mwqnice/oh-admin/pkg/convert"
)

// GetLinkList 获取友链列表
func (svc *Service) GetLinkList(ctx context.Context, params *dto.GetLinkListRequest) ([]*model.Link, int64, error) {
	return svc.dao.GetLinkList(ctx, params)
}

func (svc *Service) GetLinkInfo(ctx context.Context, id int64) (*model.Link, error) {
	fmt.Print("222333", id)
	return svc.dao.GetLinkInfo(ctx, &model.Link{Model: &model.Model{ID: int(id)}})
}
func (svc *Service) CreateLink(ctx context.Context, params *dto.CreateLinkRequest) error {
	link := &model.Link{
		Name:   params.Name,
		Url:    params.Url,
		Image:  params.Image,
		Sort:   convert.Int(params.Sort),
		Status: convert.Int(params.Status),
	}
	return svc.dao.CreateLink(ctx, link)
}
func (svc *Service) UpdateLink(ctx context.Context, params *dto.UpdateLinkRequest) error {
	fmt.Printf("%#v\n", params)
	link, err := svc.dao.GetLinkInfo(ctx, &model.Link{Model: &model.Model{ID: convert.Int(params.Id)}})
	if err != nil {
		return err
	}
	if link.Model == nil {
		return errors.New("该记录不存在")
	}
	link.Name = params.Name
	link.Url = params.Url
	link.Image = params.Image
	link.Status = convert.Int(params.Status)
	link.Sort = convert.Int(params.Sort)
	return svc.dao.UpdateLink(ctx, link)
}
func (svc *Service) DeleteLink(ctx context.Context, id string) error {
	return svc.dao.DeleteLink(ctx, convert.Int64(id))
}

// SetLinkStatus 设置友链状态
func (svc *Service) SetLinkStatus(ctx context.Context, params *dto.SetLinkStatusRequest) error {
	link, err := svc.dao.GetLinkInfo(ctx, &model.Link{Model: &model.Model{ID: convert.Int(params.Id)}})
	if err != nil {
		return err
	}
	if link.Model == nil {
		return errors.New("该记录不存在")
	}
	link.Status = convert.Int(params.Status)
	return svc.dao.UpdateLink(ctx, link)
}

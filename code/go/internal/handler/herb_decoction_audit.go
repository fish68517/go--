package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/mwqnice/oh-admin/internal/dto"
	"github.com/mwqnice/oh-admin/internal/service"
	"github.com/mwqnice/oh-admin/pkg/app"
	"github.com/mwqnice/oh-admin/pkg/errcode"
)

type herbDecoctionAuditHandler struct {
	svc service.Service
}

var HerbDecoctionAuditHandler = new(herbDecoctionAuditHandler)

func (c *herbDecoctionAuditHandler) Index(ctx *gin.Context) {
	response := app.NewResponse(ctx)
	response.BuildTpl(ctx, "herb_decoction_audit_index.html").WriteTpl(gin.H{})
}
func (c *herbDecoctionAuditHandler) List(ctx *gin.Context) {
	response := app.NewResponse(ctx)
	params := &dto.GetAuditListRequest{}
	valid, errs := app.BindAndValid(ctx, params)
	if !valid {
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}

	// 调用查询列表方法
	svc := service.New(ctx.Request.Context())
	list, total, err := svc.GetHerbDecoctionAuditmentList(svc.GetCtx(), params)
	if err != nil {
		response.ToErrorResponse(errcode.Fail.WithDetails(err.Error()))
		return
	}
	response.ToResponseList(list, int(total))
	return
}

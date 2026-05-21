package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/mwqnice/oh-admin/internal/dto"
	"github.com/mwqnice/oh-admin/internal/service"
	"github.com/mwqnice/oh-admin/pkg/app"
	"github.com/mwqnice/oh-admin/pkg/errcode"
)

type stockoutDataHandler struct {
	svc service.Service
}

var StockOutDataHandler = new(stockoutDataHandler)

func (c *stockoutDataHandler) Detail(ctx *gin.Context) {
	response := app.NewResponse(ctx)
	svc := service.New(ctx.Request.Context())
	id := ctx.Query("id")
	drugs, _ := svc.GetStockOutItem(svc.GetCtx(), id)

	// 渲染模板
	response.BuildTpl(ctx, "stock_out_detail.html").WriteTpl(gin.H{
		"medicines": drugs,
	})

}
func (c *stockoutDataHandler) Index(ctx *gin.Context) {
	response := app.NewResponse(ctx)
	response.BuildTpl(ctx, "stock_out_index.html").WriteTpl(gin.H{})
}
func (c *stockoutDataHandler) List(ctx *gin.Context) {
	response := app.NewResponse(ctx)
	params := &dto.GetStockOutListRequest{}
	valid, errs := app.BindAndValid(ctx, params)
	if !valid {
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}

	// 调用查询列表方法
	svc := service.New(ctx.Request.Context())
	list, total, err := svc.StockOutList(svc.GetCtx(), params)
	if err != nil {
		response.ToErrorResponse(errcode.Fail.WithDetails(err.Error()))
		return
	}
	response.ToResponseList(list, int(total))
	return

}

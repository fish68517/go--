package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mwqnice/oh-admin/internal/dto"
	"github.com/mwqnice/oh-admin/internal/service"
	"github.com/mwqnice/oh-admin/pkg/app"
	"github.com/mwqnice/oh-admin/pkg/errcode"
)

type stockCheckDataHandler struct {
	svc service.Service
}

var StockCheckDataHandler = new(stockCheckDataHandler)

func (c *stockCheckDataHandler) Detail(ctx *gin.Context) {
	response := app.NewResponse(ctx)
	svc := service.New(ctx.Request.Context())

	id := ctx.Query("id")
	drugs, _ := svc.GetStockCheckItem(svc.GetCtx(), id)
	// 渲染模板
	response.BuildTpl(ctx, "stock_check_detail.html").WriteTpl(gin.H{
		"medicines": drugs,
	})

}
func (c *stockCheckDataHandler) Index(ctx *gin.Context) {
	response := app.NewResponse(ctx)
	response.BuildTpl(ctx, "stock_check_index.html").WriteTpl(gin.H{})
}
func (c *stockCheckDataHandler) List(ctx *gin.Context) {
	response := app.NewResponse(ctx)
	params := &dto.GetStockCheckListRequest{}
	valid, errs := app.BindAndValid(ctx, params)
	if !valid {
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}

	// 调用查询列表方法
	svc := service.New(ctx.Request.Context())
	list, total, err := svc.StockCheckList(svc.GetCtx(), params)
	if err != nil {
		response.ToErrorResponse(errcode.Fail.WithDetails(err.Error()))
		return
	}
	response.ToResponseList(list, int(total))
	return

}
func (c *stockCheckDataHandler) Add(ctx *gin.Context) {
	response := app.NewResponse(ctx)
	svc := service.New(ctx.Request.Context())
	if ctx.Request.Method == http.MethodPost {
		params := &dto.StockCheckCreateDTO{}
		valid, errs := app.BindAndValid(ctx, params)
		if !valid {
			response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
			return
		}
		err := svc.CreateStockCheck(svc.GetCtx(), params)
		if err != nil {
			response.ToErrorResponse(errcode.Fail.WithDetails(err.Error()))
			return
		}

		response.ToResponse(dto.SuccessResponse{
			Msg: "添加成功",
		})
		return
	} else {

		// 渲染模板
		response.BuildTpl(ctx, "stock_check_add.html").WriteTpl(gin.H{})
	}
}

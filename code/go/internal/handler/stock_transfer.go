package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mwqnice/oh-admin/internal/dto"
	"github.com/mwqnice/oh-admin/internal/service"
	"github.com/mwqnice/oh-admin/pkg/app"
	"github.com/mwqnice/oh-admin/pkg/errcode"
)

type stockTransferDataHandler struct {
	svc service.Service
}

var StockTransferDataHandler = new(stockTransferDataHandler)

func (c *stockTransferDataHandler) Index(ctx *gin.Context) {
	response := app.NewResponse(ctx)
	response.BuildTpl(ctx, "stock_transfer_index.html").WriteTpl(gin.H{})
}
func (c *stockTransferDataHandler) List(ctx *gin.Context) {
	response := app.NewResponse(ctx)
	params := &dto.GetStockTransferListRequest{}
	valid, errs := app.BindAndValid(ctx, params)
	if !valid {
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}

	// 调用查询列表方法
	svc := service.New(ctx.Request.Context())
	list, total, err := svc.StockTransferList(svc.GetCtx(), params)
	if err != nil {
		response.ToErrorResponse(errcode.Fail.WithDetails(err.Error()))
		return
	}
	response.ToResponseList(list, int(total))
	return

}
func (c *stockTransferDataHandler) Add(ctx *gin.Context) {
	response := app.NewResponse(ctx)
	svc := service.New(ctx.Request.Context())
	if ctx.Request.Method == http.MethodPost {
		params := &dto.StockTransferCreateDTO{}
		valid, errs := app.BindAndValid(ctx, params)
		if !valid {
			response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
			return
		}
		err := svc.CreateStockTransfer(svc.GetCtx(), params)
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
		response.BuildTpl(ctx, "stock_transfer_add.html").WriteTpl(gin.H{})
	}
}
func (c *stockTransferDataHandler) Detail(ctx *gin.Context) {
	response := app.NewResponse(ctx)
	svc := service.New(ctx.Request.Context())
	id := ctx.Query("id")
	fmt.Print("id== ", id)
	drugs, _ := svc.GetStockTransferItem(svc.GetCtx(), id)
	// 渲染模板
	response.BuildTpl(ctx, "stock_transfer_detail.html").WriteTpl(gin.H{
		"medicines": drugs,
	})

}

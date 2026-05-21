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

type stockInDataHandler struct {
	svc service.Service
}

var StockInDataHandler = new(stockInDataHandler)

func (c *stockInDataHandler) Index(ctx *gin.Context) {
	response := app.NewResponse(ctx)
	response.BuildTpl(ctx, "stock_in_index.html").WriteTpl(gin.H{})
}
func (c *stockInDataHandler) Detail(ctx *gin.Context) {

	response := app.NewResponse(ctx)
	svc := service.New(ctx.Request.Context())
	id := ctx.Query("id")
	fmt.Println("====", id)
	drugs, _ := svc.GetStockInItem(svc.GetCtx(), id)
	// 渲染模板
	response.BuildTpl(ctx, "stock_in_detail.html").WriteTpl(gin.H{
		"medicines": drugs,
	})

}
func (c *stockInDataHandler) List(ctx *gin.Context) {
	response := app.NewResponse(ctx)
	params := &dto.GetStockInListRequest{}
	valid, errs := app.BindAndValid(ctx, params)
	if !valid {
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}

	// 调用查询列表方法
	svc := service.New(ctx.Request.Context())
	list, total, err := svc.StockInList(svc.GetCtx(), params)
	if err != nil {
		response.ToErrorResponse(errcode.Fail.WithDetails(err.Error()))
		return
	}
	response.ToResponseList(list, int(total))
	return

}
func (c *stockInDataHandler) Add(ctx *gin.Context) {
	response := app.NewResponse(ctx)
	svc := service.New(ctx.Request.Context())
	if ctx.Request.Method == http.MethodPost {
		params := &dto.StockInCreateDTO{}
		valid, errs := app.BindAndValid(ctx, params)
		if !valid {
			response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
			return
		}
		err := svc.CreateStockIn(svc.GetCtx(), params)
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
		response.BuildTpl(ctx, "stock_in_add.html").WriteTpl(gin.H{})
	}
}

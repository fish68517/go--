package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mwqnice/oh-admin/internal/dto"
	"github.com/mwqnice/oh-admin/internal/service"
	"github.com/mwqnice/oh-admin/pkg/app"
	"github.com/mwqnice/oh-admin/pkg/convert"
	"github.com/mwqnice/oh-admin/pkg/errcode"
)

type inventoryDataHandler struct {
	svc service.Service
}

var InventoryDataHandler = new(inventoryDataHandler)

func (c *inventoryDataHandler) Index(ctx *gin.Context) {
	response := app.NewResponse(ctx)
	response.BuildTpl(ctx, "inventory_index.html").WriteTpl(gin.H{})
}
func (c *inventoryDataHandler) List(ctx *gin.Context) {
	response := app.NewResponse(ctx)
	params := &dto.GetInventoryListRequest{}
	valid, errs := app.BindAndValid(ctx, params)
	if !valid {
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}

	// 调用查询列表方法
	svc := service.New(ctx.Request.Context())
	list, total, err := svc.InventoryList(svc.GetCtx(), params)
	if err != nil {
		response.ToErrorResponse(errcode.Fail.WithDetails(err.Error()))
		return
	}
	response.ToResponseList(list, int(total))
	return

}
func (c *inventoryDataHandler) Edit(ctx *gin.Context) {
	response := app.NewResponse(ctx)
	svc := service.New(ctx.Request.Context())
	if ctx.Request.Method == http.MethodPost {
		params := &dto.InventoryOperationDTO{}
		valid, errs := app.BindAndValid(ctx, params)
		if !valid {
			response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
			return
		}

		err := svc.CreateInventory(svc.GetCtx(), params)
		if err != nil {
			response.ToErrorResponse(errcode.Fail.WithDetails(err.Error()))
			return
		}

		response.ToResponse(dto.SuccessResponse{
			Msg: "出库成功",
		})
		return
	} else {
		id := ctx.Query("id")
		fmt.Println("==", id)
		info, _ := svc.GetInventoryInfo(svc.GetCtx(), int64(convert.Int(id)))

		// 渲染模板
		response.BuildTpl(ctx, "inventory_edit.html").WriteTpl(gin.H{
			"info": info,
		})
	}
}

package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/goccy/go-json"
	"github.com/mwqnice/oh-admin/internal/dto"
	"github.com/mwqnice/oh-admin/internal/model"
	"github.com/mwqnice/oh-admin/internal/service"
	"github.com/mwqnice/oh-admin/pkg/app"
	"github.com/mwqnice/oh-admin/pkg/convert"
	"github.com/mwqnice/oh-admin/pkg/errcode"
	"net/http"
)

// 定义结构体来表示 JSON 数据
type Response1 struct {
	Code int    `json:"code"`
	Data []Step `json:"data"`
}

type Step struct {
	Step      string `json:"step"`
	Performer string `json:"performer"`
	Timestamp string `json:"timestamp"`
}
type prescriptionHandler struct {
	svc service.Service
}

var PrescriptionHandler = new(prescriptionHandler)

func (c *prescriptionHandler) Index(ctx *gin.Context) {
	response := app.NewResponse(ctx)
	response.BuildTpl(ctx, "prescription_index.html").WriteTpl(gin.H{})
}
func (c *prescriptionHandler) PrescriptionQueryIndex(ctx *gin.Context) {
	response := app.NewResponse(ctx)
	response.BuildTpl(ctx, "prescription_query_index.html").WriteTpl(gin.H{})
}
func (c *prescriptionHandler) WorkStatisticsIndex(ctx *gin.Context) {
	response := app.NewResponse(ctx)
	response.BuildTpl(ctx, "work_statistics_index.html").WriteTpl(gin.H{})
}
func (c *prescriptionHandler) WorkStatisticsList(ctx *gin.Context) {
	response := app.NewResponse(ctx)
	params := &dto.GetWorkStatistListRequest{}
	valid, errs := app.BindAndValid(ctx, params)
	if !valid {
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}

	// 调用查询列表方法
	svc := service.New(ctx.Request.Context())
	list, total, err := svc.GetWorkStatisticsList(svc.GetCtx(), params)
	if err != nil {
		response.ToErrorResponse(errcode.Fail.WithDetails(err.Error()))
		return
	}
	response.ToResponseList(list, int(total))
	return
}
func (c *prescriptionHandler) SearchView(ctx *gin.Context) {
	// 定义要返回的固定数据
	fixedData := Response1{
		Code: 200,
		Data: []Step{
			{Step: "接方", Performer: "李丽", Timestamp: "2025-01-10 13:10"},
			{Step: "审核", Performer: "张伟", Timestamp: "2025-01-11 13:10"},
			{Step: "泡药", Performer: "李丽", Timestamp: "2025-01-12 13:10"},
			{Step: "煎药", Performer: "王芳", Timestamp: "2025-01-13 13:10"},
			{Step: "包装", Performer: "李丽", Timestamp: "2025-01-14 13:10"},
			{Step: "发货", Performer: "赵雷", Timestamp: "2025-01-15 13:10"},
		},
	}
	ctx.JSON(http.StatusOK, fixedData)
	return
}
func (c *prescriptionHandler) List(ctx *gin.Context) {
	response := app.NewResponse(ctx)
	params := &dto.GetPrescriptionListRequest{}
	valid, errs := app.BindAndValid(ctx, params)
	if !valid {
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}

	// 调用查询列表方法
	svc := service.New(ctx.Request.Context())
	list, total, err := svc.GetPrescriptionList(svc.GetCtx(), params)
	if err != nil {
		response.ToErrorResponse(errcode.Fail.WithDetails(err.Error()))
		return
	}
	response.ToResponseList(list, int(total))
	return
}
func (c *prescriptionHandler) Add(ctx *gin.Context) {
	response := app.NewResponse(ctx)
	svc := service.New(ctx.Request.Context())
	if ctx.Request.Method == http.MethodPost {
		params := &dto.PrescriptionDTO{}
		valid, errs := app.BindAndValid(ctx, params)
		if !valid {
			response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
			return
		}
		err := svc.CreateHosptialPrescription(svc.GetCtx(), params)
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
		response.BuildTpl(ctx, "prescription_add.html").WriteTpl(gin.H{})
	}
}

//	func (c *prescriptionHandler) Add(ctx *gin.Context) {
//		response := app.NewResponse(ctx)
//		svc := service.New(ctx.Request.Context())
//		if ctx.Request.Method == http.MethodPost {
//			params := &dto.CreatePrescriptionRequest{}
//			valid, errs := app.BindAndValid(ctx, params)
//			if !valid {
//				response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
//				return
//			}
//			err := svc.CreatePrescription(svc.GetCtx(), params)
//			if err != nil {
//				response.ToErrorResponse(errcode.Fail.WithDetails(err.Error()))
//				return
//			}
//
//			response.ToResponse(dto.SuccessResponse{
//				Msg: "添加成功",
//			})
//			return
//		} else {
//
//			// 渲染模板
//			response.BuildTpl(ctx, "prescription_add.html").WriteTpl(gin.H{})
//		}
//	}
func (c *prescriptionHandler) Edit(ctx *gin.Context) {
	response := app.NewResponse(ctx)
	svc := service.New(ctx.Request.Context())
	if ctx.Request.Method == http.MethodPost {
		params := &dto.UpdatePrescriptionRequest{}
		valid, errs := app.BindAndValid(ctx, params)
		if !valid {
			response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
			return
		}

		err := svc.UpdatePrescription(svc.GetCtx(), params)
		if err != nil {
			response.ToErrorResponse(errcode.Fail.WithDetails(err.Error()))
			return
		}

		response.ToResponse(dto.SuccessResponse{
			Msg: "修改成功",
		})
		return
	} else {
		id := ctx.Query("id")
		info, _ := svc.GetPrescriptionInfo(svc.GetCtx(), convert.Int(id))
		b, err := json.Marshal(info)
		if err != nil {
			fmt.Println("error:", err)
		}
		fmt.Print("test=", string(b))
		// 渲染模板
		response.BuildTpl(ctx, "prescription_edit.html").WriteTpl(gin.H{
			"info": info,
		})
	}
}
func (c *prescriptionHandler) Detail(ctx *gin.Context) {
	response := app.NewResponse(ctx)
	svc := service.New(ctx.Request.Context())

	id := ctx.Query("id")

	info, _ := svc.GetPrescriptionInfo(svc.GetCtx(), convert.Int(id))
	if info == nil {
		response.ToErrorResponse(errcode.InvalidParams)
		return
	}
	drugs, _ := svc.GetPrescriptionDrugInfo(svc.GetCtx(), convert.Int(info.ID))
	qrlist, _ := svc.GetQrcodeVwInfo(svc.GetCtx(), convert.Int(info.ID))
	payment, _ := svc.GetPrescriptionPaymentDetail(svc.GetCtx(), convert.Int(info.ID))
	drug := preprocessMedicines(drugs)
	qcode := preprocessQr(qrlist)
	// 渲染模板
	response.BuildTpl(ctx, "prescription_detail.html").WriteTpl(gin.H{
		"info":      info,
		"medicines": drug,
		"drugItems": drugs,
		"payment":   payment,
		"qr":        qcode,
	})

}
func preprocessQr(qrlist *model.QrcodeVw) model.QrCode {
	var qrcode model.QrCode
	if qrlist == nil {
		return qrcode
	}
	qrcode.QrCode = qrlist.PackNum + qrlist.DeScheme + qrlist.PackAcount + qrlist.BNum
	fmt.Println(qrcode.QrCode)
	return qrcode
}
func preprocessMedicines(medicines []*model.PrescriptionDrug) []model.MedicineGroup {
	var groups []model.MedicineGroup
	for i := 0; i < len(medicines); i += 2 {
		var group model.MedicineGroup
		if i < len(medicines)-1 {
			group.Row1 = medicines[i]
			group.Row2 = medicines[i+1]
		} else {
			group.Row1 = medicines[i]
			group.Row2 = nil // 最后一组可能只有一个药品
		}
		groups = append(groups, group)
	}
	return groups
}

// Delete 删除
func (c *prescriptionHandler) Delete(ctx *gin.Context) {
	response := app.NewResponse(ctx)
	id := ctx.Param("id")
	if id == "" {
		response.ToErrorResponse(errcode.InvalidParams)
		return
	}
	svc := service.New(ctx.Request.Context())
	err := svc.DeletePrescription(svc.GetCtx(), id)
	if err != nil {
		response.ToErrorResponse(errcode.Fail.WithDetails(err.Error()))
		return
	}

	response.ToResponse(dto.SuccessResponse{
		Msg: "删除成功",
	})
	return
}

func (c *prescriptionHandler) DrugSources(ctx *gin.Context) {
	response := app.NewResponse(ctx)
	svc := service.New(ctx.Request.Context())
	list, err := svc.GetDrugSources(svc.GetCtx(), ctx.Query("product_id"), ctx.Query("product_name"))
	if err != nil {
		response.ToErrorResponse(errcode.Fail.WithDetails(err.Error()))
		return
	}
	response.ToResponse(dto.SuccessResponse{
		Code: 0,
		Msg:  "success",
		Data: list,
	})
}

func (c *prescriptionHandler) PaymentDetail(ctx *gin.Context) {
	response := app.NewResponse(ctx)
	id := convert.Int(ctx.Query("id"))
	if id == 0 {
		response.ToErrorResponse(errcode.InvalidParams)
		return
	}
	svc := service.New(ctx.Request.Context())
	detail, err := svc.GetPrescriptionPaymentDetail(svc.GetCtx(), id)
	if err != nil {
		response.ToErrorResponse(errcode.Fail.WithDetails(err.Error()))
		return
	}
	response.ToResponse(dto.SuccessResponse{
		Code: 0,
		Msg:  "success",
		Data: detail,
	})
}

func (c *prescriptionHandler) MockPay(ctx *gin.Context) {
	response := app.NewResponse(ctx)
	params := &dto.MockPayRequest{}
	if err := ctx.ShouldBindJSON(params); err != nil || params.PrescriptionID == 0 {
		response.ToErrorResponse(errcode.InvalidParams)
		return
	}
	svc := service.New(ctx.Request.Context())
	detail, err := svc.MockPayPrescription(svc.GetCtx(), params)
	if err != nil {
		response.ToErrorResponse(errcode.Fail.WithDetails(err.Error()))
		return
	}
	response.ToResponse(dto.SuccessResponse{
		Code: 0,
		Msg:  "支付成功",
		Data: detail,
	})
}

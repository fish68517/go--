package handler

import (
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/mwqnice/oh-admin/internal/dto"
	"github.com/mwqnice/oh-admin/internal/service"
	"github.com/mwqnice/oh-admin/pkg/app"
	"github.com/mwqnice/oh-admin/pkg/convert"
	"github.com/mwqnice/oh-admin/pkg/errcode"
	"gorm.io/gorm"
)

type prescriptionAuditHandler struct {
	svc service.Service
}

var PrescriptionAuditHandler = new(prescriptionAuditHandler)

func (c *prescriptionAuditHandler) Index(ctx *gin.Context) {
	response := app.NewResponse(ctx)
	response.BuildTpl(ctx, "prescription_audit_index.html").WriteTpl(gin.H{})
}
func (c *prescriptionAuditHandler) List(ctx *gin.Context) {
	response := app.NewResponse(ctx)
	params := &dto.GetPrescriptionAuditListRequest{}
	valid, errs := app.BindAndValid(ctx, params)
	if !valid {
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}

	// 调用查询列表方法
	svc := service.New(ctx.Request.Context())
	list, total, err := svc.GetPrescriptionAduitList(svc.GetCtx(), params)
	if err != nil {
		response.ToErrorResponse(errcode.Fail.WithDetails(err.Error()))
		return
	}
	response.ToResponseList(list, int(total))
	return
}

// Audit 审核
func (c *prescriptionAuditHandler) Audit(ctx *gin.Context) {
	response := app.NewResponse(ctx)
	id := ctx.Param("id")
	if id == "" {
		response.ToErrorResponse(errcode.InvalidParams)
		return
	}
	svc := service.New(ctx.Request.Context())
	// 1. 查询处方关联的所有药品
	prescriptionDrugs, err := svc.GetPrescriptionDrugInfo(svc.GetCtx(), convert.Int(id))
	if err != nil {
		response.ToErrorResponse(errcode.Fail.WithDetails(err.Error()))
		return
	}
	if len(prescriptionDrugs) == 0 {
		response.ToResponse(dto.SuccessResponse{
			Code: 409000,
			Msg:  "处方未关联任何药品",
		})
		return

	}
	// 2. 逐个校验药品是否超剂量
	var errorDrugs []string
	for _, drug := range prescriptionDrugs {
		// 跳过无剂量的药品（按需调整逻辑）
		if drug.DrugWeight <= 0 {
			continue
		}

		// 查询该药品的启用剂量限制
		dosageLimit, err := svc.GetSingleDrugDosageInfo(svc.GetCtx(), 0, drug.DrugProductNumber)
		if err != nil {
			// 无剂量限制配置时，默认允许通过（可按需调整为“无配置则审核失败”）
			if err == gorm.ErrRecordNotFound {
				continue
			}
		}
		if dosageLimit.MaxDosage == 0 {
			continue
		}

		// 校验实际剂量是否超过最大限制
		if drug.DrugWeight > dosageLimit.MaxDosage {
			errorDrugs = append(errorDrugs,
				strings.Join([]string{drug.DrugProductName, "（药品ID：" + drug.DrugProductNumber + "）"}, "")+
					"，实际剂量："+convert.String(drug.DrugWeight)+"，最大限制："+convert.String(dosageLimit.MaxDosage))
		}

	}

	userInfo, _ := svc.GetAdminUserInfo(svc.GetCtx(), svc.GetAdminLoginUid(ctx))

	// 3. 处理审核结果
	if len(errorDrugs) > 0 {
		// 审核失败：更新处方状态，返回错误信息
		response.ToResponse(dto.SuccessResponse{
			Code: 409005,
			Msg:  "审核失败,请调整修改（超剂量）" + strings.Join(errorDrugs, "\n"),
		})
		return
	}
	err = svc.AuditPrescription(svc.GetCtx(), convert.Int(id), userInfo.ID, userInfo.Username)
	if err != nil {
		response.ToErrorResponse(errcode.Fail.WithDetails(err.Error()))
		return
	}

	response.ToResponse(dto.SuccessResponse{
		Msg: "审核成功",
	})
	return
}

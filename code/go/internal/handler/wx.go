package handler

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"reflect"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/mwqnice/oh-admin/global"
	"github.com/mwqnice/oh-admin/internal/dto"
	"github.com/mwqnice/oh-admin/internal/service"
	"github.com/mwqnice/oh-admin/pkg/app"
	"github.com/mwqnice/oh-admin/pkg/convert"
	"github.com/mwqnice/oh-admin/pkg/errcode"
)

type wxHandler struct {
	svc service.Service
}

var WxHandler = new(wxHandler)

func (c *wxHandler) Index(ctx *gin.Context) {
	response := app.NewResponse(ctx)
	response.BuildTpl(ctx, "prescription_scan_index.html").WriteTpl(gin.H{})
}
func (c *wxHandler) PcList(ctx *gin.Context) {
	response := app.NewResponse(ctx)
	params := &dto.GetPrescriptionSearchListRequest{}
	valid, errs := app.BindAndValid(ctx, params)
	if !valid {
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}

	// 调用查询列表方法
	svc := service.New(ctx.Request.Context())
	list, total, err := svc.GetPrescriptionPCViewList(svc.GetCtx(), params)
	if err != nil {
		response.ToErrorResponse(errcode.Fail.WithDetails(err.Error()))
		return
	}
	response.ToResponseList(list, int(total))
	return
}
func (c *wxHandler) List(ctx *gin.Context) {
	response := app.NewResponse(ctx)
	params := &dto.GetPrescriptionListRequest{}
	valid, errs := app.BindAndValid(ctx, params)
	if !valid {
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}

	// 调用查询列表方法
	svc := service.New(ctx.Request.Context())
	list, total, err := svc.GetPrescriptionViewList(svc.GetCtx(), params)
	if err != nil {
		response.ToErrorResponse(errcode.Fail.WithDetails(err.Error()))
		return
	}
	response.ToResponseList(list, int(total))
	return
}
func ParseBillNumber(barcode string) (string, error) {
	// 检查条码长度是否足够
	if len(barcode) < 17 {
		return "", fmt.Errorf("barcode too short: %s", barcode)
	}

	// 截取第8位到第17位
	billNumberStr := barcode[7:17]

	// 移除前面的0
	billNumberStr = strings.TrimLeft(billNumberStr, "0")

	// 如果移除0后字符串为空，说明原字符串全是0，根据业务逻辑可能需要特殊处理
	if billNumberStr == "" {
		// 这里可以返回错误，或者返回一个默认值，或者进行其他处理
		// 返回错误示例：
		return "", fmt.Errorf("bill number is all zeros: %s", barcode[7:17])
		// 返回默认值示例（如果需要）：
		// return "0", nil
	}

	// 返回解析后的单据号
	return billNumberStr, nil
}
func (h *wxHandler) UploadChain(ctx *gin.Context) {
	response := app.NewResponse(ctx)
	id := ctx.Param("id")
	fmt.Print("id=", id)
	if id == "" {
		response.ToErrorResponse(errcode.InvalidParams)
		return
	}
	svc := service.New(ctx.Request.Context())
	prescriptionView, _ := svc.GetPrescriptionView(svc.GetCtx(), &dto.PrescriptionRequest{Id: convert.Int(id)})
	fmt.Print("prescriptionView=", prescriptionView)
	tisaneRecord := &dto.TisaneRecord{
		HospitalName:       prescriptionView.HospitalName,
		PatientName:        prescriptionView.PatientName,
		PrescriptionNumber: prescriptionView.PrescriptionNumber,
		DoPerson:           prescriptionView.DoPerson,
		DoTime:             prescriptionView.DoTime,
		PresAduitReviewer:  prescriptionView.PresAduitReviewer,
		PresAduitTime:      prescriptionView.PresAduitTime,
		AdjustmentReviewer: prescriptionView.AdjustmentReviewer,
		AdjustmentTime:     prescriptionView.AdjustmentTime,
		AuditDatetime:      prescriptionView.AuditDatetime,
		AuditReviewer:      prescriptionView.AuditReviewer,
		SoakingPerson:      prescriptionView.SoakingPerson,
		SoakStartTime:      prescriptionView.SoakStartTime,
		SoakEndTime:        prescriptionView.SoakEndTime,
		DecoctionPerson:    prescriptionView.DecoctionPerson,
		DecoctionStartTime: prescriptionView.DecoctionStartTime,
		DecoctionEndTime:   prescriptionView.DecoctionEndTime,
		PackingPersonnel:   prescriptionView.PackingPersonnel,
		PackStartTime:      prescriptionView.PackStartTime,
		PackEndTime:        prescriptionView.PackEndTime,
		DeliveryPersonnel:  prescriptionView.DeliveryPersonnel,
		DeliveryTime:       prescriptionView.DeliveryTime,
		LogisticsNumber:    prescriptionView.LogisticsNumber,
	}
	// 声明一个 TisaneRecord 类型的切片（数组）
	var tisaneRecords []*dto.TisaneRecord
	err := ValidateTimes(tisaneRecord)
	if err != nil {
		response.ToResponse(dto.SuccessResponse{
			Msg:  "数据不完整=" + err.Error(),
			Code: 87890,
		})
		return
	}
	// 将 tisaneRecord 添加到切片中
	tisaneRecords = append(tisaneRecords, tisaneRecord)
	tisaneRecordJson, _ := json.Marshal(tisaneRecords)
	fmt.Println("json串：--", string(tisaneRecordJson))
	url := global.AppSetting.UpChain
	respBody, err := PostJSON(url, tisaneRecordJson)
	if err != nil {
		fmt.Printf("错误: %v\n", err)
		return
	}
	var rp dto.SuccessResponse

	// 将 JSON 数据解析到结构体中
	err = json.Unmarshal(respBody, &rp)
	if err != nil {
		log.Fatalf("解析 JSON 数据失败: %v", err)
	}
	// 判断 Code 是否等于 201
	if rp.Code == http.StatusOK {
		//更新处方状态
		response.ToResponse(dto.SuccessResponse{
			Msg:  "上链成功,交易ID=" + reflect.ValueOf(rp.Data).String(),
			Code: 0,
		})
		return
	} else {
		response.ToResponse(dto.SuccessResponse{
			Msg:  "上链失败=" + reflect.ValueOf(rp.Data).String(),
			Code: rp.Code,
			Data: rp.Data,
		})
		return
	}

}
func ValidateTimes(t *dto.TisaneRecord) error {
	timeFields := []struct {
		Name  string
		Value string
	}{
		{"DoTime", t.DoTime},
		{"PresAduitTime", t.PresAduitTime},
		{"AdjustmentTime", t.AdjustmentTime},
		{"SoakStartTime", t.SoakStartTime},
		{"DecoctionStartTime", t.DecoctionStartTime},
		{"PackStartTime", t.PackStartTime},
		{"DeliveryTime", t.DeliveryTime},
	}

	for _, field := range timeFields {
		if field.Value == "" {
			return errors.New(fmt.Sprintf("Error: %s is empty", field.Name))
		}
		// Optionally, you can also validate if the time string is in a correct format
		// _, err := time.Parse(time.RFC3339, field.Value)
		// if err != nil {
		// 	return errors.New(fmt.Sprintf("Error: %s has incorrect format", field.Name))
		// }
	}
	return nil
}

// PostJSON 发送一个包含 JSON 数据的 HTTP POST 请求
func PostJSON(url string, data []byte) ([]byte, error) {

	// 创建一个新的 HTTP 请求
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(data))
	if err != nil {
		return nil, fmt.Errorf("无法创建请求: %v", err)
	}

	// 设置请求头
	req.Header.Set("Content-Type", "application/json")

	// 发送请求并获取响应
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("无法发送请求: %v", err)
	}
	defer resp.Body.Close()

	// 读取响应体
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("无法读取响应体: %v", err)
	}

	return body, nil
}
func (h *wxHandler) ScanBarcode(ctx *gin.Context) {
	response := app.NewResponse(ctx)
	svc := service.New(ctx.Request.Context())
	if ctx.Request.Method == http.MethodPost {
		params := &dto.ScanPcRequest{}
		valid, errs := app.BindAndValid(ctx, params)
		if !valid {
			response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
			return
		}
		id, err := ParseBillNumber(params.Barcode)
		fmt.Println("id:", id)

		if err != nil {
			fmt.Println("Error parsing BNum:", err)
			return
		}

		prescriptionId := id
		adminLoginUid := svc.GetAdminLoginUid(ctx)
		state, _ := svc.GetPrescriptionState(svc.GetCtx(), &dto.PrescriptionRequest{Id: convert.Int(prescriptionId)})
		// 获取用户信息
		userInfo, _ := svc.GetGetAdminUserWithRole(svc.GetCtx(), adminLoginUid)

		employeeId := adminLoginUid
		userName := userInfo.Realname
		roleName := userInfo.Name
		switch roleName {
		case "调剂员":
			if state != "审核" {
				response.ToErrorResponse(errcode.StateMessage.WithDetails(state))
				return
			}
			request := &dto.PrescriptionFlowRequest{}
			request.PrescriptionID = convert.Int(prescriptionId)
			request.Barcode = params.Barcode
			request.OperateName = userName
			request.EmployeeId = employeeId
			request.WordContent = "调剂"
			err := svc.CreateAdjust(svc.GetCtx(), request)
			if err != nil {
				response.ToErrorResponse(errcode.Fail.WithDetails(err.Error()))
				return
			}
			if err != nil {
				response.ToErrorResponse(errcode.Fail.WithDetails(err.Error()))
				return
			}
			break
		case "复核员":
			if state != "调剂" {
				response.ToErrorResponse(errcode.StateMessage.WithDetails(state))
				return
			}
			request := &dto.PrescriptionFlowRequest{}
			request.PrescriptionID = convert.Int(prescriptionId)
			request.Barcode = params.Barcode
			request.OperateName = userName
			request.EmployeeId = employeeId
			request.WordContent = "复核"
			err := svc.CreateHerbDecoctionAudit(svc.GetCtx(), request)
			if err != nil {
				response.ToErrorResponse(errcode.Fail.WithDetails(err.Error()))
				return
			}
			if err != nil {
				response.ToErrorResponse(errcode.Fail.WithDetails(err.Error()))
				return
			}
			break
		case "泡药员":
			if state != "复核" {
				response.ToErrorResponse(errcode.StateMessage.WithDetails(state))
				return
			}
			request := &dto.PrescriptionFlowRequest{}
			request.PrescriptionID = convert.Int(prescriptionId)
			request.Barcode = params.Barcode
			request.OperateName = userName
			request.EmployeeId = employeeId
			request.WordContent = "泡药"
			err := svc.CreateSoak(svc.GetCtx(), request)
			if err != nil {
				response.ToErrorResponse(errcode.Fail.WithDetails(err.Error()))
				return
			}
			if err != nil {
				response.ToErrorResponse(errcode.Fail.WithDetails(err.Error()))
				return
			}
			break
		case "煎药员":
			if state != "泡药" {
				response.ToErrorResponse(errcode.StateMessage.WithDetails(state))
				return
			}
			request := &dto.PrescriptionFlowRequest{}
			request.PrescriptionID = convert.Int(prescriptionId)
			request.Barcode = params.Barcode
			request.OperateName = userName
			request.EmployeeId = employeeId
			request.WordContent = "煎药"
			err := svc.CreateDecoctionInfo(svc.GetCtx(), request)
			if err != nil {
				response.ToErrorResponse(errcode.Fail.WithDetails(err.Error()))
				return
			}
			if err != nil {
				response.ToErrorResponse(errcode.Fail.WithDetails(err.Error()))
				return
			}
			break
		case "包装员":
			if state != "煎药" {
				response.ToErrorResponse(errcode.StateMessage.WithDetails(state))
				return
			}
			request := &dto.PrescriptionFlowRequest{}
			request.PrescriptionID = convert.Int(prescriptionId)
			request.Barcode = params.Barcode
			request.OperateName = userName
			request.EmployeeId = employeeId
			request.WordContent = "包装"
			err := svc.CreatePack(svc.GetCtx(), request)
			if err != nil {
				response.ToErrorResponse(errcode.Fail.WithDetails(err.Error()))
				return
			}
			if err != nil {
				response.ToErrorResponse(errcode.Fail.WithDetails(err.Error()))
				return
			}
			break
		case "发货员":
			if state != "包装" {
				response.ToErrorResponse(errcode.StateMessage.WithDetails(state))
				return
			}
			request := &dto.PrescriptionFlowRequest{}
			request.PrescriptionID = convert.Int(prescriptionId)
			request.Barcode = params.Barcode
			request.OperateName = userName
			request.EmployeeId = employeeId
			request.WordContent = "发货"
			err := svc.CreateDeliver(svc.GetCtx(), request)
			if err != nil {
				response.ToErrorResponse(errcode.Fail.WithDetails(err.Error()))
				return
			}
			if err != nil {
				response.ToErrorResponse(errcode.Fail.WithDetails(err.Error()))
				return
			}
		default:
			response.ToErrorResponse(errcode.ScanNotFound.WithDetails("改角色无权限扫描！"))
			return
			break
		}
		response.ToResponse(dto.SuccessResponse{
			Msg: "扫描成功",
		})
		return
	} else {
		// 渲染模板
		response.BuildTpl(ctx, "prescription_scan_add.html").WriteTpl(gin.H{})
	}
}

// app登陆接口
type AppLoginResponse struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}
type AppResponse struct {
	Code         int         `json:"code"`
	Message      string      `json:"message"`
	Prescription interface{} `json:"prescription"`
	Drug         interface{} `json:"drug"`
}

func (h *wxHandler) Login(ctx *gin.Context) {
	svc := service.New(ctx.Request.Context())
	response := app.NewResponse(ctx)
	params := &dto.CheckAppPwdRequest{}
	valid, errs := app.BindAndValid(ctx, params)
	if !valid {
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}
	fmt.Print(params.UserName, params.Password)
	err, userinfo := svc.CheckAdminPwdReturnUserInfo(svc.GetCtx(), params)
	if err != nil {
		response.ToErrorResponse(errcode.Fail.WithDetails(err.Error()))
		return
	}
	response.ToResponse(AppLoginResponse{
		Message: "success",
		Code:    200,
		Data:    userinfo,
	})
	return

}

type JyStatusItem struct {
	UpdateTime     string `json:"updateTime"`     // 时间戳字符串
	Name           string `json:"name"`           // 患者姓名
	StatusStr      string `json:"statusStr"`      // 状态描述
	PrescriptionNo string `json:"prescriptionNo"` // 处方号
}

func (h *wxHandler) PrescriptionList(ctx *gin.Context) {
	response := app.NewResponse(ctx)
	params := &dto.GetPrescriptionListRequest{}
	valid, errs := app.BindAndValid(ctx, params)
	if !valid {
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}

	// 调用查询列表方法
	svc := service.New(ctx.Request.Context())
	list, _, err := svc.GetPrescriptionList(svc.GetCtx(), params)
	if err != nil {
		response.ToErrorResponse(errcode.Fail.WithDetails(err.Error()))
		return
	}
	response.ToResponse(AppLoginResponse{
		Message: "success",
		Code:    200,
		Data:    list,
	})

	return

}
func (h *wxHandler) PrescriptionDetail(ctx *gin.Context) {
	response := app.NewResponse(ctx)
	params := &dto.GetPrescriptionListRequest{}
	id := ctx.Param("id")
	if id == "" {
		response.ToErrorResponse(errcode.InvalidParams)
		return
	}
	fmt.Print(id)
	params.Name = id
	// 调用查询列表方法
	svc := service.New(ctx.Request.Context())
	prescription, _, err := svc.GetPrescriptionList(svc.GetCtx(), params)
	if err != nil {
		response.ToErrorResponse(errcode.Fail.WithDetails(err.Error()))
		return
	}
	drugs, _ := svc.GetPrescriptionDrugInfo(svc.GetCtx(), convert.Int(prescription[0].ID))
	response.ToResponse(AppResponse{
		Message:      "success",
		Code:         200,
		Prescription: prescription,
		Drug:         drugs,
	})

	return

}
func (h *wxHandler) AppScan(ctx *gin.Context) {
	response := app.NewResponse(ctx)
	svc := service.New(ctx.Request.Context())
	params := &dto.ScanAppRequest{}
	valid, errs := app.BindAndValid(ctx, params)
	if !valid {
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}
	id, err := ParseBillNumber(params.Barcode)
	if err != nil {
		fmt.Println("Error parsing BNum:", err)
		return
	}

	prescriptionId := id
	adminLoginUid := params.UserId
	state, _ := svc.GetPrescriptionState(svc.GetCtx(), &dto.PrescriptionRequest{Id: convert.Int(prescriptionId)})
	// 获取用户信息
	userInfo, _ := svc.GetGetAdminUserWithRole(svc.GetCtx(), adminLoginUid)

	employeeId := adminLoginUid
	userName := userInfo.Realname
	roleName := userInfo.Name
	switch roleName {
	case "调剂员":
		if state != "审核" {
			response.ToErrorResponse(errcode.StateMessage.WithDetails(state))
			return
		}
		request := &dto.PrescriptionFlowRequest{}
		request.PrescriptionID = convert.Int(prescriptionId)
		request.Barcode = params.Barcode
		request.OperateName = userName
		request.EmployeeId = employeeId
		request.WordContent = "调剂"
		err := svc.CreateAdjust(svc.GetCtx(), request)
		if err != nil {
			response.ToErrorResponse(errcode.Fail.WithDetails(err.Error()))
			return
		}
		if err != nil {
			response.ToErrorResponse(errcode.Fail.WithDetails(err.Error()))
			return
		}
		break
	case "复核员":
		if state != "调剂" {
			response.ToErrorResponse(errcode.StateMessage.WithDetails(state))
			return
		}
		request := &dto.PrescriptionFlowRequest{}
		request.PrescriptionID = convert.Int(prescriptionId)
		request.Barcode = params.Barcode
		request.OperateName = userName
		request.EmployeeId = employeeId
		request.WordContent = "复核"
		err := svc.CreateHerbDecoctionAudit(svc.GetCtx(), request)
		if err != nil {
			response.ToErrorResponse(errcode.Fail.WithDetails(err.Error()))
			return
		}
		if err != nil {
			response.ToErrorResponse(errcode.Fail.WithDetails(err.Error()))
			return
		}
		break
	case "泡药员":
		if state != "复核" {
			response.ToErrorResponse(errcode.StateMessage.WithDetails(state))
			return
		}
		request := &dto.PrescriptionFlowRequest{}
		request.PrescriptionID = convert.Int(prescriptionId)
		request.Barcode = params.Barcode
		request.OperateName = userName
		request.EmployeeId = employeeId
		request.WordContent = "泡药"
		err := svc.CreateSoak(svc.GetCtx(), request)
		if err != nil {
			response.ToErrorResponse(errcode.Fail.WithDetails(err.Error()))
			return
		}
		if err != nil {
			response.ToErrorResponse(errcode.Fail.WithDetails(err.Error()))
			return
		}
		break
	case "煎药员":
		if state != "泡药" {
			response.ToErrorResponse(errcode.StateMessage.WithDetails(state))
			return
		}
		request := &dto.PrescriptionFlowRequest{}
		request.PrescriptionID = convert.Int(prescriptionId)
		request.Barcode = params.Barcode
		request.OperateName = userName
		request.EmployeeId = employeeId
		request.WordContent = "煎药"
		err := svc.CreateDecoctionInfo(svc.GetCtx(), request)
		if err != nil {
			response.ToErrorResponse(errcode.Fail.WithDetails(err.Error()))
			return
		}
		if err != nil {
			response.ToErrorResponse(errcode.Fail.WithDetails(err.Error()))
			return
		}
		break
	case "包装员":
		if state != "煎药" {
			response.ToErrorResponse(errcode.StateMessage.WithDetails(state))
			return
		}
		request := &dto.PrescriptionFlowRequest{}
		request.PrescriptionID = convert.Int(prescriptionId)
		request.Barcode = params.Barcode
		request.OperateName = userName
		request.EmployeeId = employeeId
		request.WordContent = "包装"
		err := svc.CreatePack(svc.GetCtx(), request)
		if err != nil {
			response.ToErrorResponse(errcode.Fail.WithDetails(err.Error()))
			return
		}
		if err != nil {
			response.ToErrorResponse(errcode.Fail.WithDetails(err.Error()))
			return
		}
		break
	case "发货员":
		if state != "包装" {
			response.ToErrorResponse(errcode.StateMessage.WithDetails(state))
			return
		}
		request := &dto.PrescriptionFlowRequest{}
		request.PrescriptionID = convert.Int(prescriptionId)
		request.Barcode = params.Barcode
		request.OperateName = userName
		request.EmployeeId = employeeId
		request.WordContent = "发货"
		err := svc.CreateDeliver(svc.GetCtx(), request)
		if err != nil {
			response.ToErrorResponse(errcode.Fail.WithDetails(err.Error()))
			return
		}
		if err != nil {
			response.ToErrorResponse(errcode.Fail.WithDetails(err.Error()))
			return
		}
	default:
		response.ToErrorResponse(errcode.ScanNotFound.WithDetails("改角色无权限扫描！"))
		return
		break
	}
	response.ToResponse(dto.SuccessResponse{
		Msg: "扫描成功",
	})
	return

}
func (h *wxHandler) PrescriptionDelete(ctx *gin.Context) {
	response := app.NewResponse(ctx)
	params := &dto.PrescriptionDeleteRequest{}
	valid, errs := app.BindAndValid(ctx, params)
	if !valid {
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}

	svc := service.New(ctx.Request.Context())
	err := svc.PrescriptionDelete(svc.GetCtx(), params)
	if err != nil {
		response.ToErrorResponse(errcode.Fail.WithDetails(err.Error()))
		return
	}
	response.ToResponse(AppLoginResponse{
		Message: "success",
		Code:    200,
	})

	return

}

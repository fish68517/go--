package handler

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/mwqnice/oh-admin/internal/dto"
	"github.com/mwqnice/oh-admin/internal/service"
	"github.com/mwqnice/oh-admin/pkg/app"
)

type trace_chinese_medicineHandler struct {
	svc service.Service
}

var Trace_chinese_medicineHandler = new(trace_chinese_medicineHandler)

func (c *trace_chinese_medicineHandler) Index(ctx *gin.Context) {
	response := app.NewResponse(ctx)
	response.BuildTpl(ctx, "trace_chinese_medicine_index.html").WriteTpl(gin.H{})
}
func (c *trace_chinese_medicineHandler) Flow(ctx *gin.Context) {
	response := app.NewResponse(ctx)
	rxNumber := ctx.Query("rx_number") // 从查询参数中获取处方号
	if rxNumber == "" {
		response.ToResponse(dto.SuccessResponse{
			Msg:  "处方号不能为空",
			Code: 7890,
			Data: nil,
		})
		return
	}
	// 调用上传方法
	svc := service.New(ctx.Request.Context())
	record, err := svc.GetPrescriptionView(svc.GetCtx(), &dto.PrescriptionRequest{PrescriptionNumber: rxNumber})
	if err != nil {

		response.ToResponse(dto.SuccessResponse{
			Msg:  "查询失败",
			Code: 7890,
			Data: nil,
		})
		return
	}
	layout := "2006-01-02 15:04:05"
	parsedoTime, err := time.Parse(layout, record.DoTime)
	if err != nil {
		response.ToResponse(dto.SuccessResponse{
			Msg:  "查询失败,处方状态接方",
			Code: 7891,
			Data: nil,
		})
		return
	}
	parsePresAduitTime, err := time.Parse(layout, record.PresAduitTime)
	if err != nil {
		response.ToResponse(dto.SuccessResponse{
			Msg:  "查询失败,处方状态审核",
			Code: 7891,
			Data: nil,
		})
		return
	}
	parseAdjustmentTime, err := time.Parse(layout, record.AdjustmentTime)
	if err != nil {
		response.ToResponse(dto.SuccessResponse{
			Msg:  "查询失败,处方状态调剂",
			Code: 7891,
			Data: nil,
		})
		return
	}
	parseAuditDatetime, err := time.Parse(layout, record.AuditDatetime)
	if err != nil {
		response.ToResponse(dto.SuccessResponse{
			Msg:  "查询失败,处方状态复核",
			Code: 7891,
			Data: nil,
		})
		return
	}
	parseSoakStartTime, err := time.Parse(layout, record.SoakStartTime)
	if err != nil {
		response.ToResponse(dto.SuccessResponse{
			Msg:  "查询失败,处方状态泡药",
			Code: 7891,
			Data: nil,
		})
		return
	}
	parseDecoctionStartTime, err := time.Parse(layout, record.DecoctionStartTime)
	if err != nil {
		response.ToResponse(dto.SuccessResponse{
			Msg:  "查询失败,处方状态煎药",
			Code: 7891,
			Data: nil,
		})
		return
	}
	parsePackStartTime, err := time.Parse(layout, record.PackStartTime)
	if err != nil {
		response.ToResponse(dto.SuccessResponse{
			Msg:  "查询失败,处方状态包装",
			Code: 7891,
			Data: nil,
		})
		return
	}
	parseDeliveryTime, err := time.Parse(layout, record.DeliveryTime)
	if err != nil {
		response.ToResponse(dto.SuccessResponse{
			Msg:  "查询失败,处方状态未发货",
			Code: 7891,
			Data: nil,
		})
		return
	}

	mockData := dto.PrescriptionFlowResponse{
		Code:    0,
		Message: "success",
		Data: dto.PrescriptionFlowData{
			Process: []dto.PrescriptionFlow{
				{Step: "接方", Performer: record.DoPerson, Timestamp: parsedoTime},
				{Step: "审核", Performer: record.PresAduitReviewer, Timestamp: parsePresAduitTime},
				{Step: "调剂", Performer: record.PresAduitReviewer, Timestamp: parseAdjustmentTime},
				{Step: "复核", Performer: record.AuditReviewer, Timestamp: parseAuditDatetime},
				{Step: "泡药", Performer: record.SoakingPerson, Timestamp: parseSoakStartTime},
				{Step: "煎药", Performer: record.DecoctionPerson, Timestamp: parseDecoctionStartTime},
				{Step: "包装", Performer: record.PackingPersonnel, Timestamp: parsePackStartTime},
				{Step: "发货", Performer: record.DeliveryPersonnel, Timestamp: parseDeliveryTime},
			},
		},
	}
	ctx.JSON(http.StatusOK, mockData)
	return

}

//从区块链上查询数据
//func (c *trace_chinese_medicineHandler) Flow(ctx *gin.Context) {
//	response := app.NewResponse(ctx)
//	rxNumber := ctx.Query("rx_number") // 从查询参数中获取处方号
//	if rxNumber == "" {
//		response.ToErrorResponse(errcode.InvalidParams)
//		return
//	}
//	baseURL := "http://116.198.196.57:9000/queryRecipes"
//	// 定义查询参数
//	params := map[string]string{
//		"recipe_code": rxNumber,
//	}
//
//	// 调用带参数的 HTTP GET 方法
//	body, err := HTTPGetWithParams(baseURL, params)
//	if err != nil {
//		response.ToResponse(dto.SuccessResponse{
//			Msg:  "查询失败",
//			Code: 8009,
//		})
//		return
//	}
//	fmt.Printf("Response Body: %s\n", body)
//	var rp dto.SuccessResponse
//	// 将 JSON 数据解析到结构体中
//	err = json.Unmarshal(body, &rp)
//	if err != nil {
//		log.Fatalf("解析 JSON 数据失败: %v", err)
//	}
//	if rp.Code == 200 {
//		var record dto.TisaneRecord
//		layout := "2006-01-02 15:04:05"
//		err = json.Unmarshal([]byte(convert.String(rp.Data)), &record)
//		if err != nil {
//			log.Fatalf("解析 JSON 数据失败: %v", err)
//		}
//		parsedoTime, err := time.Parse(layout, record.DoTime)
//		if err != nil {
//			fmt.Println("Error parsing time:", err)
//			return
//		}
//		parsePresAduitTime, err := time.Parse(layout, record.PresAduitTime)
//		if err != nil {
//			fmt.Println("Error parsing time:", err)
//			return
//		}
//		parseAdjustmentTime, err := time.Parse(layout, record.AdjustmentTime)
//		if err != nil {
//			fmt.Println("Error parsing time:", err)
//			return
//		}
//		parseAuditDatetime, err := time.Parse(layout, record.AuditDatetime)
//		if err != nil {
//			fmt.Println("Error parsing time:", err)
//			return
//		}
//		parseSoakStartTime, err := time.Parse(layout, record.SoakStartTime)
//		if err != nil {
//			fmt.Println("Error parsing time:", err)
//			return
//		}
//		parseDecoctionStartTime, err := time.Parse(layout, record.DecoctionStartTime)
//		if err != nil {
//			fmt.Println("Error parsing time:", err)
//			return
//		}
//		parsePackStartTime, err := time.Parse(layout, record.PackStartTime)
//		if err != nil {
//			fmt.Println("Error parsing time:", err)
//			return
//		}
//		parseDeliveryTime, err := time.Parse(layout, record.DeliveryTime)
//		if err != nil {
//			fmt.Println("Error parsing time:", err)
//			return
//		}
//
//		mockData := dto.PrescriptionFlowResponse{
//			Code:    0,
//			Message: "success",
//			Data: dto.PrescriptionFlowData{
//				Process: []dto.PrescriptionFlow{
//					{Step: "接方", Performer: record.DoPerson, Timestamp: parsedoTime},
//					{Step: "审核", Performer: record.PresAduitReviewer, Timestamp: parsePresAduitTime},
//					{Step: "调剂", Performer: record.PresAduitReviewer, Timestamp: parseAdjustmentTime},
//					{Step: "复核", Performer: record.AuditReviewer, Timestamp: parseAuditDatetime},
//					{Step: "泡药", Performer: record.SoakingPerson, Timestamp: parseSoakStartTime},
//					{Step: "煎药", Performer: record.DecoctionPerson, Timestamp: parseDecoctionStartTime},
//					{Step: "包装", Performer: record.PackingPersonnel, Timestamp: parsePackStartTime},
//					{Step: "发货", Performer: record.DeliveryPersonnel, Timestamp: parseDeliveryTime},
//				},
//			},
//		}
//		ctx.JSON(http.StatusOK, mockData)
//		return
//	} else {
//		response.ToResponse(dto.SuccessResponse{
//			Msg:  "查询失败" + reflect.ValueOf(rp.Data).String(),
//			Code: rp.Code,
//			Data: rp.Data,
//		})
//		return
//	}
//
//}

// HTTPGetWithParams 发送带参数的 HTTP GET 请求
func HTTPGetWithParams(baseURL string, params map[string]string) ([]byte, error) {
	// 创建一个 URL 对象
	u, err := url.Parse(baseURL)
	if err != nil {
		return nil, err
	}

	// 添加查询参数
	q := u.Query()
	for key, value := range params {
		q.Set(key, value)
	}
	u.RawQuery = q.Encode()

	// 创建一个 HTTP 客户端
	client := &http.Client{
		Timeout: 10 * time.Second, // 设置请求超时时间
	}

	// 发送 HTTP GET 请求
	resp, err := client.Get(u.String())
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// 检查 HTTP 响应码
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to get data: %s", resp.Status)
	}

	// 读取响应体
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return bodyBytes, nil
}

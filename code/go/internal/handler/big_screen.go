package handler

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/mwqnice/oh-admin/internal/model"
	"github.com/mwqnice/oh-admin/internal/service"
)

type bigScreenHandler struct {
	svc service.Service
}

var BigScreenHandler = new(bigScreenHandler)

func (c *bigScreenHandler) BigScreenData(ctx *gin.Context) {

	// 调用上传方法
	svc := service.New(ctx.Request.Context())
	getPrescriptionSummaryList, _ := svc.GetPrescriptionSummaryList(ctx)
	getPrescriptionDosageCountList, _ := svc.GetPrescriptionDosageCountList(ctx)
	getPrescriptionDecoctionCountList, _ := svc.GetPrescriptionDecoctionCountList(ctx)
	getPrescriptionDetailList, _ := svc.GetPrescriptionDetailList(ctx)
	getDrugTotalQuantityList, _ := svc.GetDrugTotalQuantityList(ctx)
	getPrescriptionCountByDeliveryDateList, _ := svc.GetPrescriptionCountByDeliveryDateList(ctx)
	getPrescriptionCountByDateDateList, _ := svc.GetPrescriptionCountByDateDateList(ctx)
	getPrescriptionCountByHospital, _ := svc.GetPrescriptionCountByHospital(ctx)
	if getPrescriptionSummaryList == nil {
		getPrescriptionSummaryList = &model.PrescriptionSummary{}
	}
	response := Response{
		Success:   true,
		Message:   "",
		Code:      200,
		Timestamp: time.Now().UnixMilli(),
		Result: Result{
			State: State{
				OverviewData: []OverviewDataItem{
					{Title: "测试", Value: "2219"},
					{Title: "测试", Value: "32107"},
				},
				TotalData: TotalData{
					ModalSale:     getPrescriptionSummaryList.PrescriptionTotal,     //接方数量
					EquipmentSale: getPrescriptionSummaryList.PrescriptionPackTotal, //调剂数量
					ProduceTotal:  getPrescriptionSummaryList.PrescriptionTotal,     //今日接方数量
					Name:          "大屏总览数据",
					MaterialSale:  getPrescriptionSummaryList.PrescriptionFinishTotal, //复核数量
					ShowDate:      time.Now().Format("2006-01-02 15:04:05"),
					UseTotal:      getPrescriptionSummaryList.PrescriptionFinishTotal, //今日已完成数量
					DetonatorSale: getPrescriptionSummaryList.PrescriptionAduitTotal,  //审方数量
				},
				MarketData:   ConvertPrescriptionData(getPrescriptionDosageCountList),
				ProduceData:  ConvertDrugTotalQuantityCount(getDrugTotalQuantityList),
				UseData:      ConvertPrescriptionDecoctionCount(getPrescriptionDecoctionCountList),
				CustomerData: ConvertPrescriptionCountByHospital(getPrescriptionCountByHospital),
				SaleData:     ConvertPrescriptionCountByDateCount(getPrescriptionCountByDateDateList, getPrescriptionCountByDeliveryDateList),
				MapData:      getPrescriptionDetailList,
			},
		},
	}
	ctx.JSON(http.StatusOK, response)

	return
}

// 优化后的转换函数
func ConvertPrescriptionData(data []*model.PrescriptionDosageCount) []NamedData {
	var marketData []NamedData
	// 假设所有处方贴数都属于同一个类别，这里我们只创建一个 NamedData 实例
	namedData := NamedData{
		Name: "处方贴数",
		Data: make([]DataItem, 0, len(data)), // 预先分配切片容量以提高性能
	}
	for _, item := range data {
		namedData.Data = append(namedData.Data, DataItem{
			Title: item.DoTime.Format("2006-01-02"),
			Value: fmt.Sprintf("%.1f", float64(item.PrescriptionDosageCount)), // 假设需要保留一位小数，根据实际情况调整
		})
	}
	// 因为所有数据都添加到了同一个 NamedData 实例中，所以直接将这个实例添加到切片中
	marketData = append(marketData, namedData)
	return marketData
}
func ConvertPrescriptionDecoctionCount(data []*model.PrescriptionDecoctionCount) []NamedData {
	var marketData []NamedData
	// 假设所有处方贴数都属于同一个类别，这里我们只创建一个 NamedData 实例
	namedData := NamedData{
		Name: "调配方式",
		Data: make([]DataItem, 0, len(data)), // 预先分配切片容量以提高性能
	}
	for _, item := range data {
		namedData.Data = append(namedData.Data, DataItem{
			Title: item.DecoctionType,
			Value: fmt.Sprintf("%.1f", float64(item.DecoctionCount)), // 假设需要保留一位小数，根据实际情况调整
		})
	}
	// 因为所有数据都添加到了同一个 NamedData 实例中，所以直接将这个实例添加到切片中
	marketData = append(marketData, namedData)
	return marketData
}
func ConvertPrescriptionCountByHospital(data []*model.PrescriptionCountByHospital) []NamedData {
	var marketData []NamedData
	// 假设所有处方贴数都属于同一个类别，这里我们只创建一个 NamedData 实例
	namedData := NamedData{
		Name: "医院处方占比",
		Data: make([]DataItem, 0, len(data)), // 预先分配切片容量以提高性能
	}
	for _, item := range data {
		namedData.Data = append(namedData.Data, DataItem{
			Title: item.HospitalName,
			Value: fmt.Sprintf("%.1f", float64(item.PrescriptionCount)), // 假设需要保留一位小数，根据实际情况调整
		})
	}
	// 因为所有数据都添加到了同一个 NamedData 实例中，所以直接将这个实例添加到切片中
	marketData = append(marketData, namedData)
	return marketData
}
func ConvertDrugTotalQuantityCount(data []*model.DrugTotalQuantity) []NamedData {
	var marketData []NamedData
	// 假设所有处方贴数都属于同一个类别，这里我们只创建一个 NamedData 实例
	namedData := NamedData{
		Name: "饮品使用量",
		Data: make([]DataItem, 0, len(data)), // 预先分配切片容量以提高性能
	}
	for _, item := range data {
		namedData.Data = append(namedData.Data, DataItem{
			Title: item.DrugProductName,
			Value: fmt.Sprintf("%.1f", float64(item.TotalQuantity)), // 假设需要保留一位小数，根据实际情况调整
		})
	}
	// 因为所有数据都添加到了同一个 NamedData 实例中，所以直接将这个实例添加到切片中
	marketData = append(marketData, namedData)
	return marketData
}
func ConvertPrescriptionCountByDateCount(data []*model.PrescriptionCountByDate, data_2 []*model.PrescriptionCountByDeliveryDate) []NamedData {
	var marketData []NamedData
	// 假设所有处方贴数都属于同一个类别，这里我们只创建一个 NamedData 实例
	namedData := NamedData{
		Name: "接方数量",
		Data: make([]DataItem, 0, len(data)), // 预先分配切片容量以提高性能
	}
	for _, item := range data {
		namedData.Data = append(namedData.Data, DataItem{
			Title: item.DoTime.Format("2006-01-02"),
			Value: fmt.Sprintf("%.1f", float64(item.PrescriptionCount)), // 假设需要保留一位小数，根据实际情况调整
		})
	}
	// 因为所有数据都添加到了同一个 NamedData 实例中，所以直接将这个实例添加到切片中
	marketData = append(marketData, namedData)
	namedData_2 := NamedData{
		Name: "完成数量",
		Data: make([]DataItem, 0, len(data)), // 预先分配切片容量以提高性能
	}
	for _, item := range data_2 {
		namedData_2.Data = append(namedData_2.Data, DataItem{
			Title: item.DeliveryDate.Format("2006-01-02"),
			Value: fmt.Sprintf("%.1f", float64(item.PrescriptionCount)), // 假设需要保留一位小数，根据实际情况调整
		})
	}
	marketData = append(marketData, namedData_2)
	return marketData
}

// 定义所有数据结构
type Response struct {
	Success   bool   `json:"success"`
	Message   string `json:"message"`
	Code      int    `json:"code"`
	Result    Result `json:"result"`
	Timestamp int64  `json:"timestamp"`
}

type Result struct {
	State State `json:"state"`
}

type State struct {
	OverviewData []OverviewDataItem          `json:"overviewData"`
	TotalData    TotalData                   `json:"totalData"`
	MarketData   []NamedData                 `json:"marketData"`
	ProduceData  []NamedData                 `json:"produceData"`
	UseData      []NamedData                 `json:"useData"`
	CustomerData []NamedData                 `json:"customerData"`
	SaleData     []NamedData                 `json:"saleData"`
	MapData      []*model.PrescriptionDetail `json:"mapData"`
}

type OverviewDataItem struct {
	Title string `json:"title"`
	Value string `json:"value"`
}

type TotalData struct {
	ModalSale     int    `json:"modalSale"`
	EquipmentSale int    `json:"equipmentSale"`
	ProduceTotal  int    `json:"produceTotal"`
	Name          string `json:"name"`
	MaterialSale  int    `json:"materialSale"`
	ShowDate      string `json:"showDate"`
	UseTotal      int    `json:"useTotal"`
	DetonatorSale int    `json:"detonatorSale"`
}

type NamedData struct {
	Name string     `json:"name"`
	Data []DataItem `json:"data"`
}

type DataItem struct {
	Title string `json:"title"`
	Value string `json:"value"`
}

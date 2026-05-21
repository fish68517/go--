package dto

type GetDeviceRecordListRequest struct {
	Name     string `form:"name" json:"name"`
	Status   int    `json:"status"`
	Page     int64  `form:"page" json:"page"`
	PageSize int64  `form:"page_size" json:"page_size"`
}

type InspectionRecordUpdateRequest struct {
	ID                 string `json:"id" binding:"required" comment:"记录ID"`
	EquipmentType      string `json:"equipment_type" binding:"required" comment:"设备类型"`
	EquipmentID        string `json:"equipment_id" binding:"required,max=50" comment:"设备编号"`
	HealthStatus       string `json:"health_status"  comment:"卫生状态(1-良好 0-较差)"`
	DisinfectionStatus string `json:"disinfection_status"  comment:"消毒状态(1-已消毒 0-未消毒)"`
	Status             string `json:"status"  comment:"运行状态(1-正常 0-异常)"`
	InspectionTime     string `json:"inspection_time" binding:"required" comment:"巡检时间"`
	Inspector          string `json:"inspector" binding:"required" comment:"巡检人员"`
}

// InspectionRecordCreateRequest 新增设备巡检记录请求
type InspectionRecordCreateRequest struct {
	EquipmentType      string `json:"equipment_type" binding:"required" comment:"设备类型"`
	EquipmentID        string `json:"equipment_id" binding:"required,max=50" comment:"设备编号"`
	HealthStatus       string `json:"health_status"  comment:"卫生状态(1-良好 0-较差)"`
	DisinfectionStatus string `json:"disinfection_status"  comment:"消毒状态(1-已消毒 0-未消毒)"`
	Status             string `json:"status"  comment:"运行状态(1-正常 0-异常)"`
	InspectionTime     string `json:"inspection_time" binding:"required" comment:"巡检时间"`
	Inspector          string `json:"inspector" binding:"required" comment:"巡检人员"`
}

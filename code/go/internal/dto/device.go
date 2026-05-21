package dto

type DeviceListRequest struct {
	Name     string `form:"name" json:"name"`
	Status   int    `json:"status"`
	Page     int64  `form:"page" json:"page"`
	PageSize int64  `form:"page_size" json:"page_size"`
}
type CreateDeviceRequest struct {
	EquipmentType string `json:"equipment_type" binding:"required"`
	DeviceName    string `json:"device_name" binding:"required"`
	DeviceRoom    string `json:"device_room"`
	UnitNumber    string `json:"unit_number"`
	Remark        string `json:"remark"`
}

// UpdateDeviceRequest 更新设备请求
type UpdateDeviceRequest struct {
	ID            string `json:"id" binding:"required"` // 设备ID(必填)
	EquipmentType string `json:"equipment_type"`        // 设备类型
	DeviceName    string `json:"device_name"`           // 设备名称
	DeviceRoom    string `json:"device_room"`           // 设备所在房间
	UnitNumber    string `json:"unit_number"`           // 单元编号
	Remark        string `json:"remark"`                // 备注
}

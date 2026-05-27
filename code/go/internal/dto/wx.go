package dto

type ScanPcRequest struct {
	Barcode string `form:"barcode" binding:"required" json:"barcode"` // 条码
}
type ScanAppRequest struct {
	Barcode string `form:"barcode" binding:"required" json:"barcode"` // 条码
	UserId  int    `form:"user_id" binding:"required" json:"user_id"` // 条码
}
type PrescriptionFlowRequest struct {
	PrescriptionID int    `json:"prescription_id" binding:"required"`
	Barcode        string `json:"bacode" binding:"required"`
	OperateName    string `json:"operate_name" binding:"required"`
	EmployeeId     int    `json:"employee_id" binding:"required"`
	StartTime      string `json:"start_time" binding:"required"`
	WordContent    string `json:"word_content" binding:"required"`
	Remark         string `json:"remark" binding:"required"`
	// 其他字段根据具体需求添加...
}

type TisaneRecord struct {
	HospitalName       string `json:"hospital_name"`
	PatientName        string `json:"patient_name"`
	PrescriptionNumber string `json:"prescription_number"`
	DoPerson           string `json:"do_person"`
	DoTime             string `json:"do_time"`
	PresAduitReviewer  string `json:"pres_aduit_reviewer,omitempty"`
	PresAduitTime      string `json:"pres_aduit_time,omitempty"`
	AdjustmentReviewer string `json:"adjustment_reviewer,omitempty"`
	AdjustmentTime     string `json:"adjustment_time,omitempty"`
	AuditDatetime      string `json:"audit_datetime,omitempty"`
	AuditReviewer      string `json:"audit_reviewer,omitempty"`
	SoakingPerson      string `json:"soaking_person,omitempty"`
	SoakStartTime      string `json:"soak_start_time,omitempty"`
	SoakEndTime        string `json:"soak_end_time,omitempty"`
	DecoctionPerson    string `json:"decoction_person,omitempty"`
	DecoctionStartTime string `json:"decoction_start_time,omitempty"`
	DecoctionEndTime   string `json:"decoction_end_time,omitempty"`
	PackingPersonnel   string `json:"packing_personnel,omitempty"`
	PackStartTime      string `json:"pack_start_time,omitempty"`
	PackEndTime        string `json:"pack_end_time,omitempty"`
	DeliveryPersonnel  string `json:"delivery_personnel,omitempty"`
	DeliveryTime       string `json:"delivery_time,omitempty"`
	LogisticsNumber    string `json:"logistics_number,omitempty"`
}

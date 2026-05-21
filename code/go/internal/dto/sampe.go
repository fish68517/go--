package dto

type GetSampleListRequest struct {
	PrescriptionNumber string `form:"prescription_number" json:"prescription_number"` // 处方号
	Status             int    `form:"status" json:"status"`
	StartTime          string `form:"start_time" json:"start_time"`
	EndTime            string `form:"end_time" json:"end_time"`
	Page               int64  `form:"page" json:"page"`
	PageSize           int64  `form:"page_size" json:"page_size"`
}
type UpdateSampleRequest struct {
	Id                   string `form:"id" binding:"required" json:"id"`
	PrescriptionNumber   string `form:"prescription_number"  json:"prescription_number"`
	PrescriptionWeight   string `form:"prescription_weight"  json:"prescription_weight"`
	PrescriptionSJWeight string `form:"prescription_sj_weight"  json:"prescription_sj_weight"`
	DrugCount            string `form:"drug_count"  json:"drug_count"`
	DrugSJCount          string `form:"drug_sj_count"  json:"drug_sj_count"`
	Dosage               string `form:"dosage" json:"dosage"`
	Status               string `form:"status" json:"status"`
	OperateTime          string `form:"operate_time" json:"operate_time"`
	OperateName          string `form:"operate_name" json:"operate_name"`
	Remark               string `form:"remark" json:"remark"`
}

type SampleAppDTO struct {
	PrescriptionNumber   string `json:"prescription_number"`
	PrescriptionWeight   int    `json:"prescription_weight"`
	PrescriptionSJWeight int    `json:"prescription_sj_weight"`
	DrugCount            int    `json:"drug_count"`
	DrugSJCount          int    `json:"drug_sj_count"`
	Dosage               int    `json:"dosage"`
	Status               string `json:"status"`
	OperateTime          string `json:"operate_time"`
	OperateName          string `json:"operate_name"`
	Remark               string `json:"remark"`
}
type SampleDTO struct {
	PrescriptionNumber   string `json:"prescription_number"`
	PrescriptionWeight   string `json:"prescription_weight"`
	PrescriptionSJWeight string `json:"prescription_sj_weight"`
	DrugCount            string `json:"drug_count"`
	DrugSJCount          string `json:"drug_sj_count"`
	Dosage               string `json:"dosage"`
	Status               string `json:"status"`
	OperateTime          string `json:"operate_time"`
	OperateName          string `json:"operate_name"`
	Remark               string `json:"remark"`
}

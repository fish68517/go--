package dto

// CreateDrugDosageRequest 创建药品剂量限制请求
type CreateDrugDosageRequest struct {
	DrugID     string `json:"drug_id" binding:"required"`         // 药品ID（必填）
	MaxDosage  string `json:"max_dosage" binding:"required,gt=0"` // 最大剂量（必填，必须大于0）
	DosageDesc string `json:"dosage_desc"`                        // 剂量说明
	IsEnabled  string `json:"is_enabled" `                        // 是否启用（可选，默认1）
}

// UpdateDrugDosageRequest 更新药品剂量限制请求
type UpdateDrugDosageRequest struct {
	Id         string `json:"id"`
	DrugId     string `json:"drug_id"`
	MaxDosage  string `json:"max_dosage"`
	DosageDesc string `json:"dosage_desc"`
	IsEnabled  string `json:"is_enabled"`
}

type GetDrugDosageListRequest struct {
	Name     string `form:"name" json:"name"`
	Status   int    `json:"status"`
	Page     int64  `form:"page" json:"page"`
	PageSize int64  `form:"page_size" json:"page_size"`
}

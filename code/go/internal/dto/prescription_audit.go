package dto

// GetPrescriptionAuditListRequest 获取处方审核列表
type GetPrescriptionAuditListRequest struct {
	Name     string `form:"name" json:"name"`
	Status   int    `json:"status"`
	Page     int64  `form:"page" json:"page"`
	PageSize int64  `form:"page_size" json:"page_size"`
}
type UpdatePrescriptionAuditRequest struct {
	ID          string `form:"id" json:"id"`
	Reviewer    string `form:"reviewer" json:"reviewer"`
	ReviewTime  string `form:"reviewtime" json:"reviewtime"`
	AuditStatus int    `form:"auditstatus" json:"auditstatus"`
}

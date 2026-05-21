package dto

type GetHosptialDataRequest struct {
	DateStart string `form:"date_start" json:"date_start"`
	DateEnd   string `form:"date_end" json:"date_end"`
	Page      int64  `form:"page" json:"page"`
	PageSize  int64  `form:"page_size" json:"page_size"`
}

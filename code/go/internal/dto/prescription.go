/**
 * @describe linkDto
 * @author fengxh
 * @since 2025/1/8
 * @File :prescription
 */
package dto

import "time"

// 处方查询
type GetPrescriptionSearchListRequest struct {
	State              string    `form:"state" json:"state"`                             // 状态
	HospitalName       string    `form:"hospital_name" json:"hospital_name"`             // 医院名称
	PatientName        string    `form:"patient_name" json:"patient_name"`               // 患者姓名
	PrescriptionNumber string    `form:"prescription_number" json:"prescription_number"` // 处方编号
	StartTime          time.Time `form:"start_time" json:"start_time"`                   // 开始时间
	EndTime            time.Time `form:"end_time" json:"end_time"`                       // 结束时间
	Page               int64     `form:"page" json:"page"`                               // 分页页码
	PageSize           int64     `form:"page_size" json:"page_size"`                     // 每页数量
}

// GetPrescriptionListRequest 获取处方列表
type GetPrescriptionListRequest struct {
	Name     string `form:"name" json:"name"`
	Status   int    `json:"status"`
	Page     int64  `form:"page" json:"page"`
	PageSize int64  `form:"page_size" json:"page_size"`
}
type PrescriptionDeleteRequest struct {
	Id int `form:"id" json:"id"`
}

// GetWorkStatistListRequest 获取工作量统计列表
type GetWorkStatistListRequest struct {
	WordPerson    string `form:"word_person" json:"word_person"`
	WordContent   string `form:"word_content" json:"word_content"`
	WordDateStart string `form:"word_date_start" json:"word_date_start"`
	WordDateEnd   string `form:"word_date_end" json:"word_date_end"`
	Page          int64  `form:"page" json:"page"`
	PageSize      int64  `form:"page_size" json:"page_size"`
}

// CreatePrescriptionRequest 添加处方
type CreatePrescriptionRequest struct {
	HospitalName         string `form:"hospitalName" binding:"required" json:"hospitalName"`                 // 医院名称
	PatientAddress       string `form:"patientAddress" json:"patientAddress"`                                // 地址
	PrescriptionType     string `form:"prescriptionType" json:"prescriptionType"`                            // 处方类型
	DoTime               string `form:"doTime"  json:"doTime"`                                               // 处方日期
	DeletionNumber       string `form:"deletionNumber" json:"deletionNumber"`                                // 门诊序号
	PrescriptionNumber   string `form:"prescriptionNumber" binding:"required" json:"prescriptionNumber"`     // 处方号
	PatientName          string `form:"patientName" binding:"required" json:"patientName"`                   // 患者姓名
	PatientSex           string `form:"patientSex" binding:"required" json:"patientSex"`                     // 性别
	PatientAge           string `form:"patientAge" binding:"required" json:"patientAge"`                     // 年龄
	PatientPhone         string `form:"patientPhone" json:"patientPhone"`                                    // 联系电话
	DepartmentName       string `form:"departmentName"  json:"departmentName"`                               // 科室
	DiagnosisResult      string `form:"diagnosisResult" binding:"required" json:"diagnosisResult"`           // 诊断结果
	Dosage               string `form:"dosage" binding:"required" json:"dosage"`                             // 贴数
	AdministrationCount  string `form:"AdministrationCount" binding:"required" json:"AdministrationCount"`   // 次数
	AdministrationMethod string `form:"AdministrationMethod" binding:"required" json:"AdministrationMethod"` //服用方式
	SoakTime             string `form:"SoakTime" json:"SoakTime"`                                            // 侵泡时间
	InpatientArea        string `form:"inpatientArea" json:"inpatientArea"`                                  // 病区号
	WardName             string `form:"wardName" json:"wardName"`                                            // 病房号
	SickBed              string `form:"sickBed" json:"sickBed"`                                              // 病床号
	PackageCount         string `form:"packageCount"  binding:"required" json:"packageCount"`                // 包装量
	DecoctionScheme      string `form:"decoctionScheme"  binding:"required" json:"decoctionScheme"`          //煎药方案
	IsDecoction          string `form:"isDecoction"  binding:"required" json:"isDecoction"`                  //是否代煎
	DoctorName           string `form:"DoctorName"  binding:"required" json:"doctorName"`                    //医生
	Footnote             string `form:"Footnote" json:"Footnote"`                                            // 病床号
	DrugPickupTime       string `form:"DrugPickupTime" json:"DrugPickupTime"`                                // 病床号
	SoakWaterAmount      string `form:"SoakWaterAmount" json:"SoakWaterAmount"`
	DecoctionMethod      string `form:"decoctionMethod" json:"decoctionMethod"`       //煎药方法
	AdministrationWay    string `form:"AdministrationWay" json:"AdministrationWay"`   // 服用方法
	DrugPickupNumber     string `form:"DrugPickupNumber" json:"DrugPickupNumber"`     // 取药号
	Remarks              string `form:"Remarks" json:"Remarks"`                       // 备注
	AdditionalRemarksA   string `form:"AdditionalRemarksA" json:"AdditionalRemarksA"` // 备注a
	AdditionalRemarksB   string `form:"AdditionalRemarksB" json:"AdditionalRemarksB"` // 备注b
}
type PrescriptionRequest struct {
	Id                 int    `form:"id" binding:"required" json:"id"`
	PrescriptionNumber string `form:"prescriptionNumber"  json:"prescriptionNumber"`
}

// UpdatePrescriptionRequest 更新处方
type UpdatePrescriptionRequest struct {
	Id                   string `form:"id" binding:"required" json:"id"`
	HospitalName         string `form:"hospitalName" binding:"required" json:"hospitalName"`                 // 医院名称
	PatientAddress       string `form:"patientAddress" json:"patientAddress"`                                // 地址
	PrescriptionType     string `form:"prescriptionType" json:"prescriptionType"`                            // 处方类型
	DoTime               string `form:"doTime"  json:"doTime"`                                               // 处方日期
	DeletionNumber       string `form:"deletionNumber" json:"deletionNumber"`                                // 门诊序号
	PrescriptionNumber   string `form:"prescriptionNumber" binding:"required" json:"prescriptionNumber"`     // 处方号
	PatientName          string `form:"patientName" binding:"required" json:"patientName"`                   // 患者姓名
	PatientSex           string `form:"patientSex" binding:"required" json:"patientSex"`                     // 性别
	PatientAge           string `form:"patientAge" binding:"required" json:"patientAge"`                     // 年龄
	PatientPhone         string `form:"patientPhone" json:"patientPhone"`                                    // 联系电话
	DepartmentName       string `form:"departmentName" binding:"required" json:"departmentName"`             // 科室
	DiagnosisResult      string `form:"diagnosisResult" binding:"required" json:"diagnosisResult"`           // 诊断结果
	Dosage               string `form:"dosage" binding:"required" json:"dosage"`                             // 贴数
	AdministrationCount  string `form:"AdministrationCount" binding:"required" json:"AdministrationCount"`   // 次数
	AdministrationMethod string `form:"AdministrationMethod" binding:"required" json:"AdministrationMethod"` //服用方式
	SoakTime             string `form:"SoakTime" json:"SoakTime"`                                            // 侵泡时间
	InpatientArea        string `form:"inpatientArea" json:"inpatientArea"`                                  // 病区号
	WardName             string `form:"wardName" json:"wardName"`                                            // 病房号
	SickBed              string `form:"sickBed" json:"sickBed"`                                              // 病床号
	PackageCount         string `form:"packageCount"  binding:"required" json:"packageCount"`                // 包装量
	DecoctionScheme      string `form:"decoctionScheme"  binding:"required" json:"decoctionScheme"`          //煎药方案
	IsDecoction          string `form:"isDecoction"  binding:"required" json:"isDecoction"`                  //是否代煎
	DoctorName           string `form:"DoctorName"  binding:"required" json:"doctorName"`                    //医生
	Footnote             string `form:"Footnote" json:"Footnote"`                                            // 病床号
	DrugPickupTime       string `form:"DrugPickupTime" json:"DrugPickupTime"`                                // 病床号
	SoakWaterAmount      string `form:"SoakWaterAmount" json:"SoakWaterAmount"`
	DecoctionMethod      string `form:"decoctionMethod" json:"decoctionMethod"`       //煎药方法
	AdministrationWay    string `form:"AdministrationWay" json:"AdministrationWay"`   // 服用方法
	DrugPickupNumber     string `form:"DrugPickupNumber" json:"DrugPickupNumber"`     // 取药号
	Remarks              string `form:"Remarks" json:"Remarks"`                       // 备注
	AdditionalRemarksA   string `form:"AdditionalRemarksA" json:"AdditionalRemarksA"` // 备注a
	AdditionalRemarksB   string `form:"AdditionalRemarksB" json:"AdditionalRemarksB"` // 备注b                        // 显示顺序
}

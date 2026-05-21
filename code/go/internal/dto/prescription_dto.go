package dto

type PrescriptionDTO struct {
	HospitalId           string             `form:"hospitalId" binding:"required" json:"hospitalId"`                     // 医院名称
	HospitalName         string             `form:"hospitalName"  json:"hospitalName"`                                   // 医院名称
	PatientAddress       string             `form:"patientAddress" json:"patientAddress"`                                // 地址
	PrescriptionType     string             `form:"prescriptionType" json:"prescriptionType"`                            // 处方类型
	DoTime               string             `form:"doTime"  json:"doTime"`                                               // 处方日期
	DeletionNumber       string             `form:"deletionNumber" json:"deletionNumber"`                                // 门诊序号
	PrescriptionNumber   string             `form:"prescriptionNumber" binding:"required" json:"prescriptionNumber"`     // 处方号
	PatientName          string             `form:"patientName" binding:"required" json:"patientName"`                   // 患者姓名
	PatientSex           string             `form:"patientSex" binding:"required" json:"patientSex"`                     // 性别
	PatientAge           string             `form:"patientAge" binding:"required" json:"patientAge"`                     // 年龄
	PatientPhone         string             `form:"patientPhone" json:"patientPhone"`                                    // 联系电话
	DepartmentName       string             `form:"departmentName" binding:"required" json:"departmentName"`             // 科室
	DiagnosisResult      string             `form:"diagnosisResult" binding:"required" json:"diagnosisResult"`           // 诊断结果
	Dosage               string             `form:"dosage" binding:"required" json:"dosage"`                             // 贴数
	AdministrationCount  string             `form:"AdministrationCount" binding:"required" json:"AdministrationCount"`   // 次数
	AdministrationMethod string             `form:"AdministrationMethod" binding:"required" json:"AdministrationMethod"` //服用方式
	SoakTime             string             `form:"SoakTime" json:"SoakTime"`                                            // 侵泡时间
	InpatientArea        string             `form:"inpatientArea" json:"inpatientArea"`                                  // 病区号
	WardName             string             `form:"wardName" json:"wardName"`                                            // 病房号
	SickBed              string             `form:"sickBed" json:"sickBed"`                                              // 病床号
	PackageCount         string             `form:"packageCount"  binding:"required" json:"packageCount"`                // 包装量
	DecoctionScheme      string             `form:"decoctionScheme"  binding:"required" json:"decoctionScheme"`          //煎药方案
	IsDecoction          string             `form:"isDecoction"  binding:"required" json:"isDecoction"`                  //是否代煎
	DoctorName           string             `form:"DoctorName"  binding:"required" json:"doctorName"`                    //医生
	Footnote             string             `form:"Footnote" json:"Footnote"`                                            // 病床号
	DrugPickupTime       string             `form:"DrugPickupTime" json:"DrugPickupTime"`                                // 病床号
	SoakWaterAmount      string             `form:"SoakWaterAmount" json:"SoakWaterAmount"`
	DecoctionMethod      string             `form:"decoctionMethod" json:"decoctionMethod"`       //煎药方法
	AdministrationWay    string             `form:"AdministrationWay" json:"AdministrationWay"`   // 服用方法
	DrugPickupNumber     string             `form:"DrugPickupNumber" json:"DrugPickupNumber"`     // 取药号
	Remarks              string             `form:"Remarks" json:"Remarks"`                       // 备注
	AdditionalRemarksA   string             `form:"AdditionalRemarksA" json:"AdditionalRemarksA"` // 备注a
	AdditionalRemarksB   string             `form:"AdditionalRemarksB" json:"AdditionalRemarksB"` // 备注b
	Drug                 []PrescriptionDrug `json:"drug"`
}

type PrescriptionDrug struct {
	PrescriptionID         int     `db:"prescription_id" json:"prescription_id"`
	HospitalID             string  `db:"hospital_id" json:"hospital_id"`
	StockInItemID          int     `db:"stock_in_item_id" json:"stock_in_item_id"`
	PurchaseOrigin         string  `db:"purchase_origin" json:"purchase_origin"`
	BatchNo                string  `db:"batch_no" json:"batch_no"`
	DrugProductNumber      string  `db:"drug_product_number" json:"drug_product_number"`
	DrugProductName        string  `db:"drug_product_name" json:"drug_product_name"`
	MeasurementUnit        string  `db:"measurement_unit" json:"measurement_unit"`
	DrugProductDescription string  `db:"drug_product_description" json:"drug_product_description"`
	Dose                   string  `db:"dose" json:"dose"`
	DrugWeight             string  `db:"drug_weight" json:"drug_weight"`
	DrugWeights            string  `db:"drug_weights" json:"drug_weights"`
	DrugPrice              float32 `db:"drug_price" json:"drug_price"`
	TotalPrices            float32 `db:"total_prices" json:"total_prices"`
}

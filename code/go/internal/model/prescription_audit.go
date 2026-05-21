package model

import (
	"context"
	"fmt"
	"github.com/mwqnice/oh-admin/internal/dto"
	"github.com/mwqnice/oh-admin/pkg/app"
	"gorm.io/gorm"
)

type PrescriptionAudit struct {
	ID              int    `json:"ID" db:"ID"`
	PrescriptionID  int    `json:"PrescriptionID" db:"prescription_id"`
	Reviewer        string `json:"Reviewer" db:"reviewer"`
	ReviewTime      string `json:"ReviewTime" db:"review_time"`
	AuditStatus     int    `json:"AuditStatus" db:"audit_status"`
	RejectionReason string `json:"RejectionReason" db:"rejection_reason"`
	PrintStatus     int    `json:"PrintStatus" db:"print_status"`
	EmployeeID      int    `json:"EmployeeID" db:"employee_id"`
	CreateTime      string `json:"create_time" db:"create_time"`
}

// Create 插入数据
func (audit *PrescriptionAudit) Create(ctx context.Context, db *gorm.DB) error {
	return db.WithContext(ctx).Create(&audit).Error
}

// Update 更新
func (audit *PrescriptionAudit) Update(ctx context.Context, db *gorm.DB) error {

	tx := db.Begin()
	columns := make(map[string]interface{})
	columns["reviewer"] = audit.Reviewer
	columns["review_time"] = audit.ReviewTime
	columns["audit_status"] = audit.AuditStatus
	columns["employee_id"] = audit.EmployeeID
	if err := tx.WithContext(ctx).Model(&audit).Where("prescription_id=?", audit.PrescriptionID).UpdateColumns(columns).Error; err != nil {
		tx.Rollback()
		return err
	}
	// 更新处方状态
	var prescription Prescription
	if err := tx.Model(&Prescription{}).Where("id = ?", audit.PrescriptionID).Update("current_state", "审核").Find(&prescription).Error; err != nil {
		tx.Rollback()
		return err
	}
	// 提交事务
	tx.Commit()
	return nil
}

type PrescriptionAuditView struct {
	ID                   int    `gorm:"column:id;NOT NULL" json:"id"`                             // 审计ID
	PrescriptionID       int    `gorm:"column:PrescriptionID;NOT NULL" json:"prescriptionId"`     // 处方ID
	Reviewer             string `gorm:"column:Reviewer" json:"reviewer"`                          // 审核者
	ReviewTime           string `gorm:"column:ReviewTime" json:"reviewTime"`                      // 审核时间
	AuditStatus          string `gorm:"column:AuditStatus" json:"auditStatus"`                    // 审核状态
	RejectionReason      string `gorm:"column:RejectionReason" json:"rejectionReason"`            // 拒绝原因
	PrintStatus          string `gorm:"column:PrintStatus" json:"printStatus"`                    // 打印状态
	EmployeeID           int    `gorm:"column:EmployeeID" json:"employeeId"`                      // 员工ID
	DeletionNumber       string `gorm:"column:deletion_number" json:"deletionNumber"`             // 删除编号
	IsDecoction          int    `gorm:"column:Is_decoction" json:"isDecoction"`                   // 是否煎煮
	BarcodeScan          string `gorm:"column:barcode_scan" json:"barcodeScan"`                   // 条码扫描
	HospitalID           string `gorm:"column:hospital_id" json:"hospitalId"`                     // 医院ID
	HospitalName         string `gorm:"column:hospital_name" json:"hospitalName"`                 // 医院名称
	PrescriptionNumber   string `gorm:"column:prescription_number" json:"prescriptionNumber"`     // 处方编号
	DecoctionMethod      string `gorm:"column:decoction_method" json:"decoctionMethod"`           // 煎煮方法
	PrescriptionType     string `gorm:"column:prescription_type" json:"prescriptionType"`         // 处方类型
	PatientName          string `gorm:"column:patient_name" json:"patientName"`                   // 患者姓名
	PatientSex           string `gorm:"column:patient_sex" json:"patientSex"`                     // 患者性别
	PatientAge           int    `gorm:"column:patient_age" json:"patientAge"`                     // 患者年龄
	PatientPhone         string `gorm:"column:patient_phone" json:"patientPhone"`                 // 患者电话
	PatientAddress       string `gorm:"column:patient_address" json:"patientAddress"`             // 患者地址
	DepartmentName       string `gorm:"column:department_name" json:"departmentName"`             // 科室名称
	InpatientArea        string `gorm:"column:inpatient_area" json:"inpatientArea"`               // 住院区
	WardName             string `gorm:"column:ward_name" json:"wardName"`                         // 病房名称
	SickBed              string `gorm:"column:sick_bed" json:"sickBed"`                           // 病床号
	DiagnosisResult      string `gorm:"column:diagnosis_result" json:"diagnosisResult"`           // 诊断结果
	Dosage               string `gorm:"column:dosage" json:"dosage"`                              // 用法用量
	AdministrationMethod string `gorm:"column:administration_method" json:"administrationMethod"` // 服用方法
	AdministrationCount  int    `gorm:"column:administration_count" json:"administrationCount"`   // 服用次数
	PackageCount         int    `gorm:"column:package_count" json:"packageCount"`                 // 包装数量
	DecoctionScheme      string `gorm:"column:decoction_scheme" json:"decoctionScheme"`           // 煎煮方案
	OneTimeDosage        string `gorm:"column:one_time_dosage" json:"oneTimeDosage"`              // 一次剂量
	TwoTimeDosage        string `gorm:"column:two_time_dosage" json:"twoTimeDosage"`              // 二次剂量
	SoakWaterAmount      string `gorm:"column:soak_water_amount" json:"soak_water_amount"`        // 浸泡水量
	SoakTime             string `gorm:"column:soak_time" json:"soak_time"`                        // 浸泡时间
	LabelNumber          string `gorm:"column:label_number" json:"label_number"`                  // 标签编号
	Remarks              string `gorm:"column:remarks" json:"remarks"`                            // 备注
	DoctorName           string `gorm:"column:doctor_name" json:"doctor_name"`                    // 医生姓名
	Footnote             string `gorm:"column:footnote" json:"footnote"`                          // 脚注
	DrugPickupTime       string `gorm:"column:drug_pickup_time" json:"drug_pickup_time"`          // 取药时间
	DrugPickupNumber     string `gorm:"column:drug_pickup_number" json:"drug_pickup_number"`      // 取药编号
	OrderTime            string `gorm:"column:order_time" json:"order_time"`                      // 订单时间
	CurrentState         string `gorm:"column:current_state" json:"current_state"`                // 当前状态
	DoTime               string `gorm:"column:do_time" json:"do_time"`                            // 执行时间
	DoPerson             string `gorm:"column:do_person" json:"do_person"`                        // 执行人
	DistributionCompany  string `gorm:"column:distribution_company" json:"distribution_company"`  // 配送公司
	DistributionAddress  string `gorm:"column:distribution_address" json:"distribution_address"`  // 配送地址
	DistributionPhone    string `gorm:"column:distribution_phone" json:"distribution_phone"`      // 配送电话
	DistributionType     string `gorm:"column:distribution_type" json:"distribution_type"`        // 配送类型
	AdministrationWay    string `gorm:"column:administration_way" json:"administration_way"`      // 服用方式
	AdditionalRemarksA   string `gorm:"column:additional_remarks_a" json:"additional_remarks_a"`  // 附加备注A
	AdditionalRemarksB   string `gorm:"column:additional_remarks_b" json:"additional_remarks_b"`  // 附加备注B
	DrugConfirmation     string `gorm:"column:drug_confirmation" json:"drug_confirmation"`        // 药物确认
	LogisticsState       string `gorm:"column:logistics_state" json:"logistics_state"`            // 物流状态
	QueryTime            string `gorm:"column:query_time" json:"query_time"`                      // 查询时间
	QueryPerson          string `gorm:"column:query_person" json:"query_person"`                  // 查询人
	DrugCount            int    `gorm:"column:drug_count" json:"drug_count"`                      // 药物数量
	MachineRoom          string `gorm:"column:machine_room" json:"machine_room"`                  // 机房
	StatePostBack        string `gorm:"column:state_post_back" json:"state_post_back"`            // 状态回传
	SubsideTime          string `gorm:"column:subside_time" json:"subside_time"`                  // 沉淀时间
}

func (l *PrescriptionAuditView) TableName() string {
	return "prescription_audit_view"
}
func (l *PrescriptionAuditView) List(ctx context.Context, db *gorm.DB, params *dto.GetPrescriptionAuditListRequest) ([]*PrescriptionAuditView, int64, error) {
	db = db.WithContext(ctx).Table(l.TableName())
	if params.Status > 0 {
		db = db.Where("print_status = ?", params.Status)
	}
	if params.Name != "" {
		db = db.Where("prescription_number like ?", fmt.Sprintf("%%%s%%", params.Name))
	}
	var count int64
	db.Count(&count)
	var list []*PrescriptionAuditView
	if params.Page > 0 && params.PageSize > 0 {
		db = db.Offset(app.GetPageOffset(int(params.Page), int(params.PageSize))).Limit(int(params.PageSize))
	}
	if err := db.Order("id desc").Find(&list).Error; err != nil {
		return nil, 0, err
	}
	return list, count, nil
}

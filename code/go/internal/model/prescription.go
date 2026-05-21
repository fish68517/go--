package model

import (
	"fmt"

	_ "github.com/mwqnice/oh-admin/global"
	"github.com/mwqnice/oh-admin/internal/dto"
	"github.com/mwqnice/oh-admin/pkg/app"
	"golang.org/x/net/context"
	"gorm.io/gorm"
)

type Prescription struct {
	// 唯一标识符
	ID int `gorm:"primaryKey;autoIncrement" json:"id"`
	// 门诊序号
	DeletionNumber string `json:"deletionNumber,omitempty" db:"DeletionNumber"`
	// 处方类型
	PrescriptionType int `json:"prescriptionType" db:"prescriptionType"`
	// 是否为代煎（1表示是，0表示否）
	IsDecoction int `json:"isDecoction" db:"IsDecoction"`
	// 条形码扫描结果
	BarcodeScan string `json:"barcodeScan,omitempty" db:"BarcodeScan"`
	// 医院ID
	HospitalID string `json:"hospitalID" db:"HospitalID"`
	// 医院名称
	HospitalName string `json:"hospitalName,omitempty" db:"HospitalName"`
	// 处方编号
	PrescriptionNumber string `json:"prescriptionNumber" db:"PrescriptionNumber"`
	// 煎药方法
	DecoctionMethod string `json:"decoctionMethod,omitempty" db:"DecoctionMethod"`
	// 患者姓名
	PatientName string `json:"patientName,omitempty" db:"PatientName"`
	// 患者性别（1表示男，0表示女）
	PatientSex int `json:"patientSex,omitempty" db:"PatientSex"`
	// 患者年龄
	PatientAge int `json:"patientAge,omitempty" db:"PatientAge"`
	// 患者电话
	PatientPhone string `json:"patientPhone,omitempty" db:"PatientPhone"`
	// 患者地址
	PatientAddress string `json:"patientAddress,omitempty" db:"PatientAddress"`
	// 科室名称
	DepartmentName string `json:"departmentName,omitempty" db:"DepartmentName"`
	// 住院部
	InpatientArea string `json:"inpatientArea,omitempty" db:"InpatientArea"`
	// 病房名称
	WardName string `json:"wardName,omitempty" db:"WardName"`
	// 病床号
	SickBed string `json:"sickBed,omitempty" db:"SickBed"`
	// 诊断结果
	DiagnosisResult string `json:"diagnosisResult,omitempty" db:"DiagnosisResult"`
	// 用量
	Dosage int `json:"dosage" db:"Dosage"`
	// 服用方法
	AdministrationMethod string `json:"administrationMethod,omitempty" db:"AdministrationMethod"`
	// 服用次数
	AdministrationCount int `json:"administrationCount" db:"AdministrationCount"`
	// 包装数量
	PackageCount int `json:"packageCount" db:"PackageCount"`
	// 煎药方案
	DecoctionScheme string `json:"decoctionScheme" db:"DecoctionScheme"`
	// 每次用量
	OneTimeDosage int `json:"oneTimeDosage" db:"OneTimeDosage"`
	// 两次用量（若适用）
	TwoTimeDosage int `json:"twoTimeDosage" db:"TwoTimeDosage"`
	// 浸泡用水量
	SoakWaterAmount int `json:"soakWaterAmount,omitempty" db:"SoakWaterAmount"`
	// 浸泡时间
	SoakTime int `json:"soakTime" db:"SoakTime"`
	// 标签编号
	LabelNumber int `json:"labelNumber,omitempty" db:"LabelNumber"`
	// 备注
	Remarks string `json:"remarks,omitempty" db:"Remarks"`
	// 医生姓名
	DoctorName string `json:"doctorName,omitempty" db:"DoctorName"`
	// 脚注
	Footnote string `json:"footnote,omitempty" db:"Footnote"`
	// 取药时间
	DrugPickupTime string `json:"drugPickupTime" db:"DrugPickupTime"`
	// 取药编号
	DrugPickupNumber string `json:"drugPickupNumber,omitempty" db:"DrugPickupNumber"`
	// 下单时间
	OrderTime string `json:"OrderTime,omitempty" db:"OrderTime"`
	// 当前状态
	CurrentState string `json:"currentState" db:"CurrentState"`
	// 执行时间
	DoTime string `json:"doTime,omitempty" db:"DoTime"`
	// 执行人
	DoPerson string `json:"doPerson" db:"DoPerson"`
	// 配送公司
	DistributionCompany string `json:"distributionCompany,omitempty" db:"DistributionCompany"`
	// 配送地址
	DistributionAddress string `json:"distributionAddress,omitempty" db:"DistributionAddress"`
	// 配送电话
	DistributionPhone string `json:"distributionPhone,omitempty" db:"DistributionPhone"`
	// 配送类型
	DistributionType string `json:"distributionType,omitempty" db:"DistributionType"`
	// 服用方式（额外信息）
	AdministrationWay string `json:"administrationWay,omitempty" db:"AdministrationWay"`
	// 附加备注A
	AdditionalRemarksA string `json:"additionalRemarksA,omitempty" db:"AdditionalRemarksA"`
	// 附加备注B
	AdditionalRemarksB string `json:"additionalRemarksB,omitempty" db:"AdditionalRemarksB"`
	// 药物确认（1表示已确认，0表示未确认）
	DrugConfirmation int `json:"drugConfirmation,omitempty" db:"DrugConfirmation"`
	// 物流状态
	LogisticsState string `json:"logisticsState,omitempty" db:"LogisticsState"`
	// 查询时间
	QueryTime string `json:"queryTime,omitempty" db:"QueryTime"`
	// 查询人
	QueryPerson string `json:"queryPerson,omitempty" db:"QueryPerson"`
	// 药物数量
	DrugCount int `json:"drugCount,omitempty" db:"DrugCount"`
	// 机房（可能指煎药房）
	MachineRoom string `json:"machineRoom,omitempty" db:"MachineRoom"`
	// 状态回传信息
	StatePostBack string `json:"statePostBack,omitempty" db:"StatePostBack"`
	// 沉淀时间（可能指药物沉淀所需时间）
	SubsideTime int `json:"subsideTime,omitempty" db:"SubsideTime"`
}
type PrescriptionView struct {
	Id                 int    `json:"id"`
	HospitalName       string `json:"hospital_name"`
	PatientName        string `json:"patient_name"`
	PrescriptionNumber string `json:"prescription_number"`
	Dosage             string `json:"dosage"`
	DecoctionMethod    string `json:"decoction_method"`
	DecoctionScheme    string `json:"decoction_scheme"`
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
	MachineID          int    `json:"machine_id,omitempty"`
	PackingPersonnel   string `json:"packing_personnel,omitempty"`
	PackStartTime      string `json:"pack_start_time,omitempty"`
	PackEndTime        string `json:"pack_end_time,omitempty"`
	DeliveryPersonnel  string `json:"delivery_personnel,omitempty"`
	DeliveryTime       string `json:"delivery_time,omitempty"`
	LogisticsNumber    string `json:"logistics_number,omitempty"`
	CurrentState       string `json:"current_state"`
}
type WorkloadVw struct {
	WordPerson  string `json:"word_person"`
	WordContent string `json:"word_content"`
	WordDate    string `json:"word_date"`
	Dosage      string `json:"dosage"`
}
type HerbHospital struct {
	ID                int    `json:"id" db:"id"`
	HospitalNumber    string `json:"hospital_number" db:"hospital_number"`
	HospitalName      string `json:"hospital_name" db:"hospital_name"`
	HospitalShortName string `json:"hospital_short_name" db:"hospital_short_name"`
	ContactPerson     string `json:"contact_person" db:"contact_person"`
	Address           string `json:"address" db:"address"`
	AreaCode          int    `json:"area_code" db:"area_code"`
	Phone             string `json:"phone" db:"phone"`
}

func (l *WorkloadVw) TableName() string {
	return "workload_vw"
}
func (l *PrescriptionView) TableName() string {
	return "wx_prescription_view"
}
func (l *Prescription) TableName() string {
	return "prescription"
}
func (l *HerbHospital) TableName() string {
	return "herb_hospital"
}
func (l *HerbHospital) Get(ctx context.Context, db *gorm.DB) (*HerbHospital, error) {
	db = db.WithContext(ctx).Table(l.TableName())

	if l.HospitalNumber != "" {
		db = db.Where("hospital_number = ? ", l.ID)
	}
	var hp *HerbHospital
	if err := db.First(&hp).Error; err != nil && err != gorm.ErrRecordNotFound {
		return hp, err
	}
	return hp, nil
}
func (l *HerbHospital) List(ctx context.Context, db *gorm.DB, params *dto.GetDeviceRecordListRequest) ([]*InspectionRecord, int64, error) {
	db = db.WithContext(ctx).Table(l.TableName())
	if params.Name != "" {
		db = db.Where("equipment_id like ?", fmt.Sprintf("%%%s%%", params.Name))
	}
	var count int64
	db.Count(&count)
	var list []*InspectionRecord
	if params.Page > 0 && params.PageSize > 0 {
		db = db.Offset(app.GetPageOffset(int(params.Page), int(params.PageSize))).Limit(int(params.PageSize))
	}
	if err := db.Order("id desc").Find(&list).Error; err != nil {
		return nil, 0, err
	}
	return list, count, nil
}

// Get 根据条件查询单条数据
func (l *Prescription) Get(ctx context.Context, db *gorm.DB) (*Prescription, error) {
	db = db.WithContext(ctx).Table(l.TableName())

	if l.ID != 0 {
		db = db.Where("id = ? ", l.ID)
	}
	var prescription *Prescription
	if err := db.First(&prescription).Error; err != nil && err != gorm.ErrRecordNotFound {
		return prescription, err
	}
	return prescription, nil
}
func (l *PrescriptionView) PrescriptionPCViewList(ctx context.Context, db *gorm.DB, params *dto.GetPrescriptionSearchListRequest) ([]*PrescriptionView, int64, error) {
	db = db.WithContext(ctx).Table(l.TableName())
	if params.State != "0" {
		db = db.Where("current_state = ?", params.State)
	}
	if params.PrescriptionNumber != "" {
		db = db.Where("prescription_number = ?", params.PrescriptionNumber)
	}
	if params.PatientName != "" {
		db = db.Where("patient_name = ?", params.PatientName)
	}
	if params.HospitalName != "" {
		db = db.Where("hospital_name = ?", params.HospitalName)
	}
	var count int64
	db.Count(&count)
	var list []*PrescriptionView
	if params.Page > 0 && params.PageSize > 0 {
		db = db.Offset(app.GetPageOffset(int(params.Page), int(params.PageSize))).Limit(int(params.PageSize))
	}
	if err := db.Order("do_time asc").Find(&list).Error; err != nil {
		return nil, 0, err
	}
	return list, count, nil
}
func (l *PrescriptionView) PrescriptionViewList(ctx context.Context, db *gorm.DB, params *dto.GetPrescriptionListRequest) ([]*PrescriptionView, int64, error) {
	db = db.WithContext(ctx).Table(l.TableName())
	if params.Status > 0 {
		db = db.Where("current_state = ?", params.Status)
	}
	if params.Name != "" {
		db = db.Where("prescription_number like ?", fmt.Sprintf("%%%s%%", params.Name))
	}
	var count int64
	db.Count(&count)
	var list []*PrescriptionView
	if params.Page > 0 && params.PageSize > 0 {
		db = db.Offset(app.GetPageOffset(int(params.Page), int(params.PageSize))).Limit(int(params.PageSize))
	}
	if err := db.Order("do_time asc").Find(&list).Error; err != nil {
		return nil, 0, err
	}
	return list, count, nil
}
func (l *PrescriptionView) GetPrescriptionView(ctx context.Context, db *gorm.DB, params *dto.PrescriptionRequest) (*PrescriptionView, error) {
	db = db.WithContext(ctx).Table(l.TableName())
	if params.Id > 0 {
		db = db.Where("id = ?", params.Id)
	}
	if params.PrescriptionNumber != "" {
		db = db.Where("prescription_number = ?", params.PrescriptionNumber)
	}
	var list *PrescriptionView
	if err := db.Find(&list).Error; err != nil {
		return nil, err
	}
	return list, nil
}
func (l *WorkloadVw) WorkStatistViewList(ctx context.Context, db *gorm.DB, params *dto.GetWorkStatistListRequest) ([]*WorkloadVw, int64, error) {
	db = db.WithContext(ctx).Table(l.TableName())
	if params.WordPerson != "" {
		db = db.Where("word_person = ?", params.WordPerson)
	}
	if params.WordContent != "" {
		db = db.Where("word_content = ?", params.WordContent)
	}
	if params.WordDateStart != "" && params.WordDateEnd != "" {
		db = db.Where("word_date >= ? and word_date <= ? ", params.WordDateStart, params.WordDateEnd)
	}
	var count int64
	db.Count(&count)
	var list []*WorkloadVw
	if params.Page > 0 && params.PageSize > 0 {
		db = db.Offset(app.GetPageOffset(int(params.Page), int(params.PageSize))).Limit(int(params.PageSize))
	}
	if err := db.Select(" word_person,word_content,sum(dosage) dosage").Group(" word_person,word_content").Scan(&list).Error; err != nil {
		return nil, 0, err
	}
	return list, count, nil
}
func (l *Prescription) List(ctx context.Context, db *gorm.DB, params *dto.GetPrescriptionListRequest) ([]*Prescription, int64, error) {
	db = db.WithContext(ctx).Table(l.TableName())
	if params.Status > 0 {
		db = db.Where("current_state = ?", params.Status)
	}
	if params.Name != "" {
		db = db.Where("prescription_number like ?", fmt.Sprintf("%%%s%%", params.Name))
	}
	var count int64
	db.Count(&count)
	var list []*Prescription
	if params.Page > 0 && params.PageSize > 0 {
		db = db.Offset(app.GetPageOffset(int(params.Page), int(params.PageSize))).Limit(int(params.PageSize))
	}
	if err := db.Order("id desc").Find(&list).Error; err != nil {
		return nil, 0, err
	}
	return list, count, nil
}

// Delete 删除
func (l *Prescription) Delete(ctx context.Context, db *gorm.DB, id int) error {
	return db.WithContext(ctx).Where("id = ?", id).Delete(&l).Error
}

// Create 插入数据
func (l *Prescription) Create(ctx context.Context, db *gorm.DB) error {
	return db.WithContext(ctx).Create(&l).Error
}

// Update 更新
func (l *Prescription) Update(ctx context.Context, db *gorm.DB) error {
	return db.WithContext(ctx).Save(&l).Error
}

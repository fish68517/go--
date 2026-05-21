package model

import (
	"context"
	"gorm.io/gorm"
	"time"
)

// PrescriptionSummary 表示处方统计信息的结构体
type PrescriptionSummary struct {
	PrescriptionTotal          int `json:"prescription_total"`
	PrescriptionFinishTotal    int `json:"prescription_finish_total"`
	PrescriptionAduitTotal     int `json:"prescription_aduit_total"` // 注意：这里保留了原始查询中的拼写，但可能是错误的，应该是Audit
	PrescriptionSoakTotal      int `json:"prescription_soak_total"`
	PrescriptionDecoctionTotal int `json:"prescription_decoction_total"`
	PrescriptionPackTotal      int `json:"prescription_pack_total"`
}
type PrescriptionDosageCount struct {
	DoTime                  time.Time `json:"do_time"`
	PrescriptionDosageCount int       `json:"prescription_dosage_count"`
}
type PrescriptionCountByHospital struct {
	HospitalName      string `json:"hospital_name"`
	PrescriptionCount int    `json:"prescription_count"`
}
type PrescriptionDecoctionCount struct {
	DecoctionType  string `json:"decoction_type"`  //
	DecoctionCount int    `json:"decoction_count"` //
}

// PrescriptionDetail 表示处方详细信息的结构体
type PrescriptionDetail struct {
	HospitalName       string  `json:"hospital_name"`
	DecoctionType      string  `json:"decoction_type"`
	PrescriptionNumber string  `json:"prescription_number"`
	PatientName        string  `json:"patient_name"`
	Dosage             float64 `json:"dosage"` // 假设剂量为浮点数，根据实际情况调整类型
	CurrentState       string  `json:"current_state"`
	DrugCount          int     `json:"drug_count"`
}
type DrugTotalQuantity struct {
	DrugProductName string  `json:"drug_product_name"`
	TotalQuantity   float64 `json:"total_quantity"` // 假设药品重量为浮点数，根据实际情况调整类型
}
type PrescriptionCountByDate struct {
	DoTime            time.Time `json:"do_time"`
	PrescriptionCount int       `json:"prescription_count"`
}
type PrescriptionCountByDeliveryDate struct {
	DeliveryDate      time.Time `json:"delivery_date"`
	PrescriptionCount int       `json:"prescription_count"`
}

func (l *PrescriptionSummary) TableName() string {
	return "prescription_total_screen"
}
func (l *PrescriptionDosageCount) TableName() string {
	return "vw_prescription_dosage_screen"
}
func (l *PrescriptionCountByHospital) TableName() string {
	return "hospital_prescption_tj_vw"
}
func (l *PrescriptionDecoctionCount) TableName() string {
	return "prescription_is_decoction_screen"
}
func (l *PrescriptionDetail) TableName() string {
	return "prescription_list_screen"
}
func (l *DrugTotalQuantity) TableName() string {
	return "drug_used_total_screen"
}

func (l *PrescriptionCountByDate) TableName() string {
	return "prescription_do_seven_day_screen"
}

func (l *PrescriptionCountByDeliveryDate) TableName() string {
	return "prescription_finish_seven_day_screen"
}

func (l *PrescriptionSummary) List(ctx context.Context, db *gorm.DB) (*PrescriptionSummary, error) {
	db = db.WithContext(ctx).Table(l.TableName())
	var list *PrescriptionSummary
	if err := db.First(&list).Error; err != nil {
		return nil, err
	}
	return list, nil
}
func (l *PrescriptionDosageCount) List(ctx context.Context, db *gorm.DB) ([]*PrescriptionDosageCount, error) {
	db = db.WithContext(ctx).Table(l.TableName())
	var list []*PrescriptionDosageCount
	if err := db.Find(&list).Error; err != nil {
		return nil, err
	}
	return list, nil
}
func (l *PrescriptionDecoctionCount) List(ctx context.Context, db *gorm.DB) ([]*PrescriptionDecoctionCount, error) {
	db = db.WithContext(ctx).Table(l.TableName())
	var list []*PrescriptionDecoctionCount
	if err := db.Find(&list).Error; err != nil {
		return nil, err
	}
	return list, nil
}
func (l *PrescriptionDetail) List(ctx context.Context, db *gorm.DB) ([]*PrescriptionDetail, error) {
	db = db.WithContext(ctx).Table(l.TableName())
	var list []*PrescriptionDetail
	if err := db.Find(&list).Error; err != nil {
		return nil, err
	}
	return list, nil
}
func (l *DrugTotalQuantity) List(ctx context.Context, db *gorm.DB) ([]*DrugTotalQuantity, error) {
	db = db.WithContext(ctx).Table(l.TableName())
	var list []*DrugTotalQuantity
	if err := db.Find(&list).Error; err != nil {
		return nil, err
	}
	return list, nil
}
func (l *PrescriptionCountByDeliveryDate) List(ctx context.Context, db *gorm.DB) ([]*PrescriptionCountByDeliveryDate, error) {
	db = db.WithContext(ctx).Table(l.TableName())
	var list []*PrescriptionCountByDeliveryDate
	if err := db.Find(&list).Error; err != nil {
		return nil, err
	}
	return list, nil
}
func (l *PrescriptionCountByDate) List(ctx context.Context, db *gorm.DB) ([]*PrescriptionCountByDate, error) {
	db = db.WithContext(ctx).Table(l.TableName())
	var list []*PrescriptionCountByDate
	if err := db.Find(&list).Error; err != nil {
		return nil, err
	}
	return list, nil
}
func (l *PrescriptionCountByHospital) List(ctx context.Context, db *gorm.DB) ([]*PrescriptionCountByHospital, error) {
	db = db.WithContext(ctx).Table(l.TableName())
	var list []*PrescriptionCountByHospital
	if err := db.Find(&list).Error; err != nil {
		return nil, err
	}
	return list, nil
}

type Response struct {
	Success   bool   `json:"success"`
	Message   string `json:"message"`
	Code      int    `json:"code"`
	Result    Result `json:"result"`
	Timestamp int64  `json:"timestamp"`
}

type Result struct {
	State State `json:"state"`
}

type State struct {
	OverviewData []OverviewDataItem `json:"overviewData"`
	TotalData    TotalData          `json:"totalData"`
	MarketData   []NamedData        `json:"marketData"`
	ProduceData  []NamedData        `json:"produceData"`
	UseData      []NamedData        `json:"useData"`
	CustomerData []NamedData        `json:"customerData"`
	SaleData     []NamedData        `json:"saleData"`
}

type OverviewDataItem struct {
	Title string `json:"title"`
	Value string `json:"value"`
}

type TotalData struct {
	ModalSale     int    `json:"modalSale"`
	EquipmentSale int    `json:"equipmentSale"`
	ProduceTotal  string `json:"produceTotal"`
	Name          string `json:"name"`
	MaterialSale  int    `json:"materialSale"`
	ShowDate      string `json:"showDate"`
	UseTotal      string `json:"useTotal"`
	DetonatorSale int    `json:"detonatorSale"`
}

type NamedData struct {
	Name string     `json:"name"`
	Data []DataItem `json:"data"`
}

type DataItem struct {
	Title string `json:"title"`
	Value string `json:"value"`
}

package dao

import (
	"context"
	"fmt"

	"github.com/mwqnice/oh-admin/internal/model"
	"gorm.io/gorm"
)

func (d *Dao) GetDrugSources(ctx context.Context, productID, productName string) ([]*model.StockInItem, error) {
	db := d.engine.WithContext(ctx).Table((&model.StockInItem{}).TableName())
	if productID != "" {
		db = db.Where("product_id = ?", productID)
	}
	if productName != "" {
		db = db.Where("product_name LIKE ? OR remark LIKE ?", fmt.Sprintf("%%%s%%", productName), fmt.Sprintf("%%%s%%", productName))
	}
	var list []*model.StockInItem
	if err := db.Order("id DESC").Limit(50).Find(&list).Error; err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return list, nil
}

func (d *Dao) GetStockInItemByID(ctx context.Context, id int) (*model.StockInItem, error) {
	var item model.StockInItem
	if err := d.engine.WithContext(ctx).Where("id = ?", id).First(&item).Error; err != nil {
		return nil, err
	}
	return &item, nil
}

func (d *Dao) GetPrescriptionPaymentByPrescriptionID(ctx context.Context, prescriptionID int) (*model.PrescriptionPayment, error) {
	var payment model.PrescriptionPayment
	err := d.engine.WithContext(ctx).Where("prescription_id = ?", prescriptionID).First(&payment).Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return &payment, nil
}

func (d *Dao) SavePrescriptionPayment(ctx context.Context, payment *model.PrescriptionPayment) error {
	db := d.engine.WithContext(ctx)
	if payment.ID == 0 {
		return db.Create(payment).Error
	}
	return db.Save(payment).Error
}

func (d *Dao) GetLatestBlockchainTraceLog(ctx context.Context, prescriptionID int) (*model.BlockchainTraceLog, error) {
	var chainLog model.BlockchainTraceLog
	err := d.engine.WithContext(ctx).Where("prescription_id = ?", prescriptionID).Order("id DESC").First(&chainLog).Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return &chainLog, nil
}

func (d *Dao) SaveBlockchainTraceLog(ctx context.Context, chainLog *model.BlockchainTraceLog) error {
	db := d.engine.WithContext(ctx)
	if chainLog.ID == 0 {
		return db.Create(chainLog).Error
	}
	return db.Save(chainLog).Error
}

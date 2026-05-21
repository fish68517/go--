package model

import "gorm.io/gorm"

func EnsureTraceSchema(db *gorm.DB) error {
	if err := addColumnIfMissing(db, "stock_in_items", "purchase_origin", "ALTER TABLE stock_in_items ADD COLUMN purchase_origin varchar(100) DEFAULT NULL COMMENT '进货源地' AFTER batch_no"); err != nil {
		return err
	}
	if err := addColumnIfMissing(db, "stock_in_items", "supplier_name", "ALTER TABLE stock_in_items ADD COLUMN supplier_name varchar(100) DEFAULT NULL COMMENT '供应商名称' AFTER purchase_origin"); err != nil {
		return err
	}
	if err := addColumnIfMissing(db, "prescription_drug", "stock_in_item_id", "ALTER TABLE prescription_drug ADD COLUMN stock_in_item_id int DEFAULT NULL COMMENT '关联入库明细ID' AFTER hospital_id"); err != nil {
		return err
	}
	if err := addColumnIfMissing(db, "prescription_drug", "purchase_origin", "ALTER TABLE prescription_drug ADD COLUMN purchase_origin varchar(100) DEFAULT NULL COMMENT '进货源地' AFTER stock_in_item_id"); err != nil {
		return err
	}
	if err := addColumnIfMissing(db, "prescription_drug", "batch_no", "ALTER TABLE prescription_drug ADD COLUMN batch_no varchar(50) DEFAULT NULL COMMENT '入库批次号' AFTER purchase_origin"); err != nil {
		return err
	}
	if err := createTableIfMissing(db, "prescription_payment", `
CREATE TABLE prescription_payment (
  id int NOT NULL AUTO_INCREMENT,
  prescription_id int NOT NULL COMMENT '处方ID',
  prescription_number varchar(60) NOT NULL COMMENT '处方号',
  total_amount decimal(12,2) NOT NULL DEFAULT 0.00 COMMENT '应付总金额',
  pay_amount decimal(12,2) NOT NULL DEFAULT 0.00 COMMENT '实付金额',
  pay_method varchar(30) DEFAULT NULL COMMENT '模拟支付方式',
  pay_status varchar(20) NOT NULL DEFAULT 'UNPAID' COMMENT '支付状态：UNPAID/PAID/CANCELLED',
  mock_trade_no varchar(64) DEFAULT NULL COMMENT '模拟交易号',
  paid_at datetime DEFAULT NULL COMMENT '支付时间',
  created_at datetime DEFAULT CURRENT_TIMESTAMP,
  updated_at datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (id),
  UNIQUE KEY uk_prescription_payment (prescription_id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='处方模拟支付记录'`); err != nil {
		return err
	}
	if err := createTableIfMissing(db, "blockchain_trace_log", `
CREATE TABLE blockchain_trace_log (
  id int NOT NULL AUTO_INCREMENT,
  prescription_id int NOT NULL COMMENT '处方ID',
  prescription_number varchar(60) NOT NULL COMMENT '处方号',
  payload_hash varchar(64) NOT NULL COMMENT '上链内容SHA256',
  tx_id varchar(128) DEFAULT NULL COMMENT '链上交易ID',
  chain_status varchar(20) NOT NULL DEFAULT 'PENDING' COMMENT 'PENDING/SUCCESS/FAILED',
  error_message varchar(500) DEFAULT NULL,
  created_at datetime DEFAULT CURRENT_TIMESTAMP,
  updated_at datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (id),
  KEY idx_prescription_number (prescription_number),
  KEY idx_prescription_id (prescription_id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='处方溯源上链日志'`); err != nil {
		return err
	}
	return db.Exec(`
CREATE OR REPLACE VIEW prescription_list_screen AS
SELECT
  p.hospital_name AS hospital_name,
  CASE WHEN p.Is_decoction = 1 THEN '代煎' ELSE '自煎' END AS decoction_type,
  p.prescription_number AS prescription_number,
  p.patient_name AS patient_name,
  p.dosage AS dosage,
  p.current_state AS current_state,
  p.drug_count AS drug_count,
  COALESCE(pay.pay_status, 'UNPAID') AS pay_status,
  COALESCE(pay.pay_amount, 0) AS pay_amount,
  COALESCE(chain.chain_status, '') AS chain_status,
  COALESCE(chain.tx_id, '') AS tx_id
FROM prescription p
LEFT JOIN prescription_payment pay ON pay.prescription_id = p.ID
LEFT JOIN (
  SELECT b1.*
  FROM blockchain_trace_log b1
  INNER JOIN (
    SELECT prescription_id, MAX(id) AS id
    FROM blockchain_trace_log
    GROUP BY prescription_id
  ) latest ON latest.id = b1.id
) chain ON chain.prescription_id = p.ID
WHERE CAST(p.do_time AS DATE) = CURDATE()`).Error
}

func addColumnIfMissing(db *gorm.DB, tableName, columnName, statement string) error {
	var count int64
	if err := db.Raw(`
SELECT COUNT(*)
FROM information_schema.COLUMNS
WHERE TABLE_SCHEMA = DATABASE()
  AND TABLE_NAME = ?
  AND COLUMN_NAME = ?`, tableName, columnName).Scan(&count).Error; err != nil {
		return err
	}
	if count > 0 {
		return nil
	}
	return db.Exec(statement).Error
}

func createTableIfMissing(db *gorm.DB, tableName, statement string) error {
	var count int64
	if err := db.Raw(`
SELECT COUNT(*)
FROM information_schema.TABLES
WHERE TABLE_SCHEMA = DATABASE()
  AND TABLE_NAME = ?`, tableName).Scan(&count).Error; err != nil {
		return err
	}
	if count > 0 {
		return nil
	}
	return db.Exec(statement).Error
}

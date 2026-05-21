/*
 Navicat Premium Data Transfer

 Source Server         : localhost_3309
 Source Server Type    : MySQL
 Source Server Version : 80032
 Source Host           : localhost:3309
 Source Schema         : oh-admin

 Target Server Type    : MySQL
 Target Server Version : 80032
 File Encoding         : 65001

 Date: 03/11/2025 08:27:10
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for adjustment
-- ----------------------------
DROP TABLE IF EXISTS `adjustment`;
CREATE TABLE `adjustment` (
  `id` int NOT NULL AUTO_INCREMENT,
  `prescription_id` int DEFAULT NULL,
  `word_content` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
  `word_date` varchar(30) DEFAULT NULL,
  `workload_description` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
  `word_person` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
  `employee_id` int DEFAULT NULL,
  `status` int NOT NULL DEFAULT '0',
  `barcode` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
  `end_date` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=9 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='调剂信息表';

-- ----------------------------
-- Records of adjustment
-- ----------------------------
BEGIN;
INSERT INTO `adjustment` (`id`, `prescription_id`, `word_content`, `word_date`, `workload_description`, `word_person`, `employee_id`, `status`, `barcode`, `end_date`) VALUES (6, 68, '调剂', '2025-02-14 14:45:48', '', '调剂员工', 7, 0, '14032000000000068  ', '2025-02-14 14:48:08');
INSERT INTO `adjustment` (`id`, `prescription_id`, `word_content`, `word_date`, `workload_description`, `word_person`, `employee_id`, `status`, `barcode`, `end_date`) VALUES (7, 69, '调剂', '2025-02-14 15:52:08', '', '调剂员工', 7, 0, '14032000000000069  ', '2025-02-14 15:52:08');
INSERT INTO `adjustment` (`id`, `prescription_id`, `word_content`, `word_date`, `workload_description`, `word_person`, `employee_id`, `status`, `barcode`, `end_date`) VALUES (8, 79, '调剂', '2025-03-12 22:56:29', '', '调剂员工', 7, 0, '14031800000000079', '2025-03-12 22:56:29');
COMMIT;

-- ----------------------------
-- Table structure for device
-- ----------------------------
DROP TABLE IF EXISTS `device`;
CREATE TABLE `device` (
  `id` int NOT NULL AUTO_INCREMENT,
  `equipment_type` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
  `device_name` varchar(50) DEFAULT NULL,
  `device_room` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
  `unit_number` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
  `remark` varchar(255) DEFAULT NULL,
  `created_time` datetime DEFAULT NULL,
  `updated_time` datetime DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=10 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='设备信息表';

-- ----------------------------
-- Records of device
-- ----------------------------
BEGIN;
INSERT INTO `device` (`id`, `equipment_type`, `device_name`, `device_room`, `unit_number`, `remark`, `created_time`, `updated_time`) VALUES (2, '煎药机', 'JY1002', '第一煎药室', '01', '21', '0001-01-01 00:00:00', '2025-10-27 23:09:09');
INSERT INTO `device` (`id`, `equipment_type`, `device_name`, `device_room`, `unit_number`, `remark`, `created_time`, `updated_time`) VALUES (3, '煎药机', 'JY1003', '第一煎药室', '01', NULL, NULL, NULL);
INSERT INTO `device` (`id`, `equipment_type`, `device_name`, `device_room`, `unit_number`, `remark`, `created_time`, `updated_time`) VALUES (5, '包装机', 'BZ1001', '第二煎药室', '01', NULL, NULL, NULL);
INSERT INTO `device` (`id`, `equipment_type`, `device_name`, `device_room`, `unit_number`, `remark`, `created_time`, `updated_time`) VALUES (6, '打印机', 'DY1001', '第二煎药室', '01', NULL, NULL, NULL);
INSERT INTO `device` (`id`, `equipment_type`, `device_name`, `device_room`, `unit_number`, `remark`, `created_time`, `updated_time`) VALUES (7, '煎药机', '煎药机01', '第二煎药室', 'JY10001', '1', '2025-10-27 22:02:34', '2025-10-27 22:02:34');
COMMIT;

-- ----------------------------
-- Table structure for drug_basics
-- ----------------------------
DROP TABLE IF EXISTS `drug_basics`;
CREATE TABLE `drug_basics` (
  `id` int NOT NULL AUTO_INCREMENT,
  `product_batch_number` varchar(20) DEFAULT NULL,
  `drug_type_category` varchar(50) DEFAULT NULL,
  `drug_identification_code` varchar(50) DEFAULT NULL,
  `purchase_unit` varchar(20) DEFAULT NULL,
  `drug_name` varchar(100) DEFAULT NULL,
  `drug_specification` varchar(50) DEFAULT NULL,
  `shelf_position_number` varchar(20) DEFAULT NULL,
  `unit_price` decimal(18,2) DEFAULT NULL,
  `mnemonic_code` varchar(20) DEFAULT NULL,
  `remarks1` varchar(255) DEFAULT NULL,
  `production_area` varchar(50) DEFAULT NULL,
  `storage_duration` varchar(20) DEFAULT NULL,
  `batch_number_for_reference` varchar(60) DEFAULT NULL,
  `system_creation_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='药品信息表';

-- ----------------------------
-- Records of drug_basics
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for drug_match
-- ----------------------------
DROP TABLE IF EXISTS `drug_match`;
CREATE TABLE `drug_match` (
  `id` int NOT NULL AUTO_INCREMENT,
  `hospital_id` int DEFAULT NULL,
  `hospital_hnum` varchar(50) DEFAULT NULL,
  `drug_hospital_num` varchar(50) DEFAULT NULL,
  `drug_code_num` varchar(50) DEFAULT NULL,
  `priority` int DEFAULT NULL,
  `status` int NOT NULL DEFAULT '0',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='药品匹配信息表';

-- ----------------------------
-- Records of drug_match
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for herb_decoction_audit
-- ----------------------------
DROP TABLE IF EXISTS `herb_decoction_audit`;
CREATE TABLE `herb_decoction_audit` (
  `id` int NOT NULL AUTO_INCREMENT,
  `prescription_id` int NOT NULL,
  `barcode` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
  `reviewer` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
  `word_content` varchar(50) DEFAULT NULL,
  `employee_id` int DEFAULT NULL,
  `audit_datetime` varchar(30) NOT NULL,
  `audit_status` int NOT NULL DEFAULT '0',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=7 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='复核信息表';

-- ----------------------------
-- Records of herb_decoction_audit
-- ----------------------------
BEGIN;
INSERT INTO `herb_decoction_audit` (`id`, `prescription_id`, `barcode`, `reviewer`, `word_content`, `employee_id`, `audit_datetime`, `audit_status`) VALUES (5, 68, '14032000000000068  ', '复核员工', '复核', 6, '2025-02-14 14:48:30', 0);
INSERT INTO `herb_decoction_audit` (`id`, `prescription_id`, `barcode`, `reviewer`, `word_content`, `employee_id`, `audit_datetime`, `audit_status`) VALUES (6, 79, '14031800000000079', '复核员工', '复核', 6, '2025-03-12 22:56:54', 0);
COMMIT;

-- ----------------------------
-- Table structure for herb_delivery
-- ----------------------------
DROP TABLE IF EXISTS `herb_delivery`;
CREATE TABLE `herb_delivery` (
  `id` int NOT NULL AUTO_INCREMENT,
  `prescription_id` int DEFAULT NULL,
  `delivery_personnel` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '发货人',
  `delivery_time` varchar(30) DEFAULT NULL COMMENT '发货时间',
  `delivery_status` varchar(50) DEFAULT NULL,
  `barcode` varchar(255) DEFAULT NULL,
  `processing_employee_id` int DEFAULT NULL,
  `logistics_number` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL COMMENT '物流单号',
  `word_content` varchar(50) DEFAULT NULL,
  `remarks` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=7 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='发货信息表';

-- ----------------------------
-- Records of herb_delivery
-- ----------------------------
BEGIN;
INSERT INTO `herb_delivery` (`id`, `prescription_id`, `delivery_personnel`, `delivery_time`, `delivery_status`, `barcode`, `processing_employee_id`, `logistics_number`, `word_content`, `remarks`) VALUES (5, 68, '发货员工', '2025-02-23 00:55:13', '0', '14032000000000068  ', 2, '', '发货', '');
INSERT INTO `herb_delivery` (`id`, `prescription_id`, `delivery_personnel`, `delivery_time`, `delivery_status`, `barcode`, `processing_employee_id`, `logistics_number`, `word_content`, `remarks`) VALUES (6, 79, '发货员工', '2025-03-12 22:59:12', '0', '14031800000000079', 2, '', '发货', '');
COMMIT;

-- ----------------------------
-- Table structure for herb_hospital
-- ----------------------------
DROP TABLE IF EXISTS `herb_hospital`;
CREATE TABLE `herb_hospital` (
  `id` int NOT NULL AUTO_INCREMENT,
  `hospital_number` varchar(50) NOT NULL,
  `hospital_name` varchar(50) DEFAULT NULL,
  `hospital_short_name` varchar(50) DEFAULT NULL,
  `contact_person` varchar(10) DEFAULT NULL,
  `address` varchar(100) DEFAULT NULL,
  `area_code` int DEFAULT NULL,
  `phone` varchar(11) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='医院信息表';

-- ----------------------------
-- Records of herb_hospital
-- ----------------------------
BEGIN;
INSERT INTO `herb_hospital` (`id`, `hospital_number`, `hospital_name`, `hospital_short_name`, `contact_person`, `address`, `area_code`, `phone`) VALUES (1, '10001', '河南省人民医院', '', '李丽', NULL, NULL, NULL);
INSERT INTO `herb_hospital` (`id`, `hospital_number`, `hospital_name`, `hospital_short_name`, `contact_person`, `address`, `area_code`, `phone`) VALUES (2, '10002', '河南省中医院', NULL, '王菲', NULL, NULL, NULL);
COMMIT;

-- ----------------------------
-- Table structure for herb_soaking
-- ----------------------------
DROP TABLE IF EXISTS `herb_soaking`;
CREATE TABLE `herb_soaking` (
  `id` int NOT NULL AUTO_INCREMENT,
  `prescription_id` int NOT NULL,
  `word_content` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
  `start_time` varchar(30) NOT NULL,
  `duration` int DEFAULT NULL,
  `warning_status` int DEFAULT NULL,
  `remarks` varchar(1000) DEFAULT NULL,
  `end_time` varchar(30) DEFAULT NULL,
  `soaking_person` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
  `employee_id` int DEFAULT NULL,
  `barcode` varchar(50) DEFAULT NULL,
  `mark` varchar(50) DEFAULT NULL,
  `warning_time` datetime NOT NULL,
  `warning_type` varchar(100) DEFAULT NULL,
  `soaking_status` int NOT NULL DEFAULT '0',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=8 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='泡药信息表';

-- ----------------------------
-- Records of herb_soaking
-- ----------------------------
BEGIN;
INSERT INTO `herb_soaking` (`id`, `prescription_id`, `word_content`, `start_time`, `duration`, `warning_status`, `remarks`, `end_time`, `soaking_person`, `employee_id`, `barcode`, `mark`, `warning_time`, `warning_type`, `soaking_status`) VALUES (6, 68, '泡药', '2025-02-14 14:53:27', NULL, NULL, '', '2025-02-14 14:53:27', '泡药员工', 5, '14032000000000068  ', '', '2025-02-14 14:53:27', '', 0);
INSERT INTO `herb_soaking` (`id`, `prescription_id`, `word_content`, `start_time`, `duration`, `warning_status`, `remarks`, `end_time`, `soaking_person`, `employee_id`, `barcode`, `mark`, `warning_time`, `warning_type`, `soaking_status`) VALUES (7, 79, '泡药', '2025-03-12 22:57:14', NULL, NULL, '', '2025-03-12 22:57:14', '泡药员工', 5, '14031800000000079', '', '2025-03-12 22:57:14', '', 0);
COMMIT;

-- ----------------------------
-- Table structure for herbal_decoction_info
-- ----------------------------
DROP TABLE IF EXISTS `herbal_decoction_info`;
CREATE TABLE `herbal_decoction_info` (
  `id` int NOT NULL AUTO_INCREMENT,
  `prescription_id` int NOT NULL,
  `machine_id` int NOT NULL,
  `decoction_manager` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL COMMENT '煎药工姓名',
  `decoction_time` int DEFAULT NULL COMMENT '煎药时间',
  `decoction_status` int NOT NULL DEFAULT '0' COMMENT '煎药状态',
  `start_time` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '开始时间',
  `end_time` varchar(255) DEFAULT NULL,
  `word_content` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
  `employee_id` int DEFAULT NULL,
  `barcode` varchar(50) DEFAULT NULL,
  `remark` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=7 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='煎药信息表';

-- ----------------------------
-- Records of herbal_decoction_info
-- ----------------------------
BEGIN;
INSERT INTO `herbal_decoction_info` (`id`, `prescription_id`, `machine_id`, `decoction_manager`, `decoction_time`, `decoction_status`, `start_time`, `end_time`, `word_content`, `employee_id`, `barcode`, `remark`) VALUES (5, 68, 0, '煎药员工', 0, 0, '2025-02-14 14:53:58', '2025-02-14 14:53:58', '煎药', 4, '14032000000000068  ', '');
INSERT INTO `herbal_decoction_info` (`id`, `prescription_id`, `machine_id`, `decoction_manager`, `decoction_time`, `decoction_status`, `start_time`, `end_time`, `word_content`, `employee_id`, `barcode`, `remark`) VALUES (6, 79, 0, '煎药员工', 0, 0, '2025-03-12 22:57:38', '2025-03-12 22:57:38', '煎药', 4, '14031800000000079', '');
COMMIT;

-- ----------------------------
-- Table structure for inspection_record
-- ----------------------------
DROP TABLE IF EXISTS `inspection_record`;
CREATE TABLE `inspection_record` (
  `id` int NOT NULL AUTO_INCREMENT,
  `equipment_type` varchar(255) DEFAULT NULL,
  `equipment_id` varchar(50) DEFAULT NULL,
  `health_status` int DEFAULT NULL COMMENT '卫生状态',
  `disinfection_status` int DEFAULT NULL COMMENT '消毒状态',
  `status` int DEFAULT NULL COMMENT '运行状态',
  `inspection_time` datetime DEFAULT NULL,
  `inspector` varchar(255) DEFAULT NULL,
  `created_time` datetime DEFAULT NULL,
  `updated_time` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=17 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='设备巡检记录表';

-- ----------------------------
-- Records of inspection_record
-- ----------------------------
BEGIN;
INSERT INTO `inspection_record` (`id`, `equipment_type`, `equipment_id`, `health_status`, `disinfection_status`, `status`, `inspection_time`, `inspector`, `created_time`, `updated_time`) VALUES (1, '服务器', 'SRV-2023-001', 1, 1, 1, '2025-10-26 09:30:00', '张三', NULL, NULL);
INSERT INTO `inspection_record` (`id`, `equipment_type`, `equipment_id`, `health_status`, `disinfection_status`, `status`, `inspection_time`, `inspector`, `created_time`, `updated_time`) VALUES (2, '服务器', 'SRV-2023-002', 1, 1, 0, '2025-10-26 09:45:00', '张三', NULL, NULL);
INSERT INTO `inspection_record` (`id`, `equipment_type`, `equipment_id`, `health_status`, `disinfection_status`, `status`, `inspection_time`, `inspector`, `created_time`, `updated_time`) VALUES (3, '服务器', 'SRV-2023-003', 0, 0, 1, '2025-10-26 10:00:00', '李四', NULL, NULL);
INSERT INTO `inspection_record` (`id`, `equipment_type`, `equipment_id`, `health_status`, `disinfection_status`, `status`, `inspection_time`, `inspector`, `created_time`, `updated_time`) VALUES (4, '交换机', 'SW-2023-001', 1, 1, 1, '2025-10-26 10:30:00', '李四', NULL, NULL);
INSERT INTO `inspection_record` (`id`, `equipment_type`, `equipment_id`, `health_status`, `disinfection_status`, `status`, `inspection_time`, `inspector`, `created_time`, `updated_time`) VALUES (5, '交换机', 'SW-2023-002', 1, 0, 1, '2025-10-26 11:00:00', '王五', NULL, NULL);
INSERT INTO `inspection_record` (`id`, `equipment_type`, `equipment_id`, `health_status`, `disinfection_status`, `status`, `inspection_time`, `inspector`, `created_time`, `updated_time`) VALUES (6, '路由器', 'RT-2023-001', 0, 1, 1, '2025-10-26 11:30:00', '王五', NULL, NULL);
INSERT INTO `inspection_record` (`id`, `equipment_type`, `equipment_id`, `health_status`, `disinfection_status`, `status`, `inspection_time`, `inspector`, `created_time`, `updated_time`) VALUES (7, '摄像头', 'CAM-2023-001', 1, 1, 1, '2025-10-27 09:00:00', '赵六', NULL, NULL);
INSERT INTO `inspection_record` (`id`, `equipment_type`, `equipment_id`, `health_status`, `disinfection_status`, `status`, `inspection_time`, `inspector`, `created_time`, `updated_time`) VALUES (8, '摄像头', 'CAM-2023-002', 1, 1, 0, '2025-10-27 09:15:00', '赵六', NULL, NULL);
INSERT INTO `inspection_record` (`id`, `equipment_type`, `equipment_id`, `health_status`, `disinfection_status`, `status`, `inspection_time`, `inspector`, `created_time`, `updated_time`) VALUES (9, '门禁控制器', 'ACC-2023-001', 0, 0, 1, '2025-10-27 10:00:00', '钱七', NULL, NULL);
INSERT INTO `inspection_record` (`id`, `equipment_type`, `equipment_id`, `health_status`, `disinfection_status`, `status`, `inspection_time`, `inspector`, `created_time`, `updated_time`) VALUES (10, '空调', 'AIR-2023-001', 1, 0, 1, '2025-10-27 14:30:00', '钱七', NULL, NULL);
INSERT INTO `inspection_record` (`id`, `equipment_type`, `equipment_id`, `health_status`, `disinfection_status`, `status`, `inspection_time`, `inspector`, `created_time`, `updated_time`) VALUES (11, 'UPS电源', 'UPS-2023-001', 0, 1, 0, '2025-10-28 15:33:00', '孙八', '0001-01-01 00:00:00', '2025-10-28 15:37:58');
INSERT INTO `inspection_record` (`id`, `equipment_type`, `equipment_id`, `health_status`, `disinfection_status`, `status`, `inspection_time`, `inspector`, `created_time`, `updated_time`) VALUES (12, 'UPS电源', 'UPS-2023-002', 1, 0, 1, '2025-10-21 15:38:00', '孙八', '0001-01-01 00:00:00', '2025-10-28 15:38:44');
INSERT INTO `inspection_record` (`id`, `equipment_type`, `equipment_id`, `health_status`, `disinfection_status`, `status`, `inspection_time`, `inspector`, `created_time`, `updated_time`) VALUES (13, '煎药机', 'JY0012100', 1, 1, 1, '2025-10-19 00:00:00', 'admin', '2025-10-28 15:42:53', '2025-10-28 15:42:53');
INSERT INTO `inspection_record` (`id`, `equipment_type`, `equipment_id`, `health_status`, `disinfection_status`, `status`, `inspection_time`, `inspector`, `created_time`, `updated_time`) VALUES (14, '包装机', 'BZ001', 2, 1, 1, '2025-10-26 00:00:00', 'admin', '2025-10-28 15:45:08', '2025-10-28 15:45:08');
INSERT INTO `inspection_record` (`id`, `equipment_type`, `equipment_id`, `health_status`, `disinfection_status`, `status`, `inspection_time`, `inspector`, `created_time`, `updated_time`) VALUES (15, '煎药机', 'BZ0901', NULL, NULL, NULL, '2025-10-28 16:41:14', 'pda巡检', '2025-10-28 16:41:17', '2025-10-28 16:41:17');
INSERT INTO `inspection_record` (`id`, `equipment_type`, `equipment_id`, `health_status`, `disinfection_status`, `status`, `inspection_time`, `inspector`, `created_time`, `updated_time`) VALUES (16, '煎药机', 'jy0900', 1, 0, 0, '2025-10-31 08:51:00', 'pda巡检', '2025-10-29 10:17:56', '2025-10-31 08:51:29');
COMMIT;

-- ----------------------------
-- Table structure for inventory
-- ----------------------------
DROP TABLE IF EXISTS `inventory`;
CREATE TABLE `inventory` (
  `id` int NOT NULL AUTO_INCREMENT,
  `product_id` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '商品ID',
  `product_name` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
  `warehouse_id` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '仓库ID',
  `quantity` varchar(50) NOT NULL DEFAULT '0.00' COMMENT '库存数量',
  `created_time` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `updated_time` datetime DEFAULT NULL,
  `remark` varchar(50) DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE KEY `uk_product_warehouse` (`product_id`,`warehouse_id`)
) ENGINE=InnoDB AUTO_INCREMENT=17 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='库存表';

-- ----------------------------
-- Records of inventory
-- ----------------------------
BEGIN;
INSERT INTO `inventory` (`id`, `product_id`, `product_name`, `warehouse_id`, `quantity`, `created_time`, `updated_time`, `remark`) VALUES (1, '90023', '党参', '1', '100', '2025-10-30 04:57:06', '2025-10-31 23:10:10', NULL);
INSERT INTO `inventory` (`id`, `product_id`, `product_name`, `warehouse_id`, `quantity`, `created_time`, `updated_time`, `remark`) VALUES (6, '10902', '甘草', '1', '2', '2025-10-29 23:18:10', '2025-10-29 23:18:10', NULL);
INSERT INTO `inventory` (`id`, `product_id`, `product_name`, `warehouse_id`, `quantity`, `created_time`, `updated_time`, `remark`) VALUES (7, '0902389', '测试', '1', '2', '2025-10-30 08:35:13', '2025-10-29 23:18:10', NULL);
INSERT INTO `inventory` (`id`, `product_id`, `product_name`, `warehouse_id`, `quantity`, `created_time`, `updated_time`, `remark`) VALUES (8, '90023', '党参', '2', '9', '2025-10-30 08:35:05', '2025-10-31 11:13:12', NULL);
INSERT INTO `inventory` (`id`, `product_id`, `product_name`, `warehouse_id`, `quantity`, `created_time`, `updated_time`, `remark`) VALUES (9, 'RS001', '发半夏', '1', '12200', '2025-10-31 14:41:45', '2025-10-31 22:41:45', NULL);
INSERT INTO `inventory` (`id`, `product_id`, `product_name`, `warehouse_id`, `quantity`, `created_time`, `updated_time`, `remark`) VALUES (10, 'HQ002', '槟郎', '1', '100', '2025-10-31 14:06:58', '2025-10-31 21:37:31', NULL);
INSERT INTO `inventory` (`id`, `product_id`, `product_name`, `warehouse_id`, `quantity`, `created_time`, `updated_time`, `remark`) VALUES (11, 'DG003', '炙甘草', '1', '200', '2025-10-31 14:07:03', '2025-10-31 21:37:31', NULL);
INSERT INTO `inventory` (`id`, `product_id`, `product_name`, `warehouse_id`, `quantity`, `created_time`, `updated_time`, `remark`) VALUES (12, 'GQ004', '野菊', '1', '200', '2025-10-31 14:07:19', '2025-11-01 18:23:07', NULL);
INSERT INTO `inventory` (`id`, `product_id`, `product_name`, `warehouse_id`, `quantity`, `created_time`, `updated_time`, `remark`) VALUES (13, 'SQ005', '通草', '1', '300', '2025-10-31 14:07:34', '2025-10-31 21:37:31', NULL);
INSERT INTO `inventory` (`id`, `product_id`, `product_name`, `warehouse_id`, `quantity`, `created_time`, `updated_time`, `remark`) VALUES (14, 'JYH06', '山楂', '1', '100', '2025-10-31 15:03:57', '2025-10-31 23:03:57', NULL);
INSERT INTO `inventory` (`id`, `product_id`, `product_name`, `warehouse_id`, `quantity`, `created_time`, `updated_time`, `remark`) VALUES (16, 'JYH06', '山楂', '2', '70', '2025-10-31 23:03:57', '2025-11-01 18:22:39', NULL);
COMMIT;

-- ----------------------------
-- Table structure for medicine_packing
-- ----------------------------
DROP TABLE IF EXISTS `medicine_packing`;
CREATE TABLE `medicine_packing` (
  `id` int NOT NULL AUTO_INCREMENT,
  `barcode` varchar(50) DEFAULT NULL,
  `word_content` varchar(50) DEFAULT NULL,
  `prescription_id` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
  `packing_personnel` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL COMMENT '包装人员',
  `start_time` varchar(30) DEFAULT NULL,
  `end_time` varchar(30) DEFAULT NULL COMMENT '包装完成时间',
  `packing_status` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL COMMENT '包装状态',
  `employee_id` int DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=7 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='包装信息表';

-- ----------------------------
-- Records of medicine_packing
-- ----------------------------
BEGIN;
INSERT INTO `medicine_packing` (`id`, `barcode`, `word_content`, `prescription_id`, `packing_personnel`, `start_time`, `end_time`, `packing_status`, `employee_id`) VALUES (5, '14032000000000068  ', '包装', '68', '包装员工', '2025-02-14 14:54:34', '2025-02-14 14:54:34', '0', 3);
INSERT INTO `medicine_packing` (`id`, `barcode`, `word_content`, `prescription_id`, `packing_personnel`, `start_time`, `end_time`, `packing_status`, `employee_id`) VALUES (6, '14031800000000079', '包装', '79', '包装员工', '2025-03-12 22:58:50', '2025-03-12 22:58:50', '0', 3);
COMMIT;

-- ----------------------------
-- Table structure for oh_admin_menu
-- ----------------------------
DROP TABLE IF EXISTS `oh_admin_menu`;
CREATE TABLE `oh_admin_menu` (
  `id` int NOT NULL AUTO_INCREMENT COMMENT '用户id',
  `name` varchar(30) NOT NULL DEFAULT '' COMMENT '菜单名称',
  `icon` varchar(50) DEFAULT NULL COMMENT '图标',
  `url` varchar(100) DEFAULT NULL COMMENT 'URL地址',
  `pid` int NOT NULL DEFAULT '0' COMMENT '上级ID',
  `type` tinyint NOT NULL DEFAULT '0' COMMENT '类型：1模块 2导航 3菜单 4节点',
  `permission` varchar(255) DEFAULT '' COMMENT '权限标识',
  `is_show` tinyint(1) DEFAULT '1' COMMENT '是否显示：1显示 2不显示',
  `sort` int DEFAULT NULL COMMENT '显示顺序',
  `remark` varchar(255) DEFAULT NULL COMMENT '菜单备注',
  `status` tinyint(1) DEFAULT '1' COMMENT '状态1-在用 2-禁用',
  `created_time` datetime DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_time` datetime DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间',
  PRIMARY KEY (`id`),
  KEY `idx_pid` (`pid`)
) ENGINE=InnoDB AUTO_INCREMENT=67 DEFAULT CHARSET=utf8mb3;

-- ----------------------------
-- Records of oh_admin_menu
-- ----------------------------
BEGIN;
INSERT INTO `oh_admin_menu` (`id`, `name`, `icon`, `url`, `pid`, `type`, `permission`, `is_show`, `sort`, `remark`, `status`, `created_time`, `updated_time`) VALUES (1, '系统管理', 'layui-icon-component', '#', 0, 0, 'sys:sysconfig', 1, 9, '', 1, '2023-03-15 07:20:52', '2025-10-25 17:16:38');
INSERT INTO `oh_admin_menu` (`id`, `name`, `icon`, `url`, `pid`, `type`, `permission`, `is_show`, `sort`, `remark`, `status`, `created_time`, `updated_time`) VALUES (2, '权限管理', 'layui-icon-component', NULL, 1, 0, '', 1, 2, NULL, 1, '2023-03-15 07:21:28', '2023-03-24 07:19:18');
INSERT INTO `oh_admin_menu` (`id`, `name`, `icon`, `url`, `pid`, `type`, `permission`, `is_show`, `sort`, `remark`, `status`, `created_time`, `updated_time`) VALUES (3, '用户管理', 'layui-icon-component', 'user/index', 2, 0, '', 1, 3, NULL, 1, '2023-03-15 07:23:02', '2023-03-24 07:19:19');
INSERT INTO `oh_admin_menu` (`id`, `name`, `icon`, `url`, `pid`, `type`, `permission`, `is_show`, `sort`, `remark`, `status`, `created_time`, `updated_time`) VALUES (4, '角色管理', 'layui-icon-component', 'role/index', 2, 0, '', 1, 3, NULL, 1, '2023-03-15 07:23:02', '2023-03-27 02:26:41');
INSERT INTO `oh_admin_menu` (`id`, `name`, `icon`, `url`, `pid`, `type`, `permission`, `is_show`, `sort`, `remark`, `status`, `created_time`, `updated_time`) VALUES (5, '菜单管理', 'layui-icon-component', 'menu/index', 2, 0, '', 1, 3, NULL, 1, '2023-03-15 07:23:02', '2023-03-27 03:28:15');
INSERT INTO `oh_admin_menu` (`id`, `name`, `icon`, `url`, `pid`, `type`, `permission`, `is_show`, `sort`, `remark`, `status`, `created_time`, `updated_time`) VALUES (6, '业务管理', '', '#', 0, 0, 'sys:process', 1, 1, '', 1, '2023-03-30 15:00:28', '2025-02-18 09:02:11');
INSERT INTO `oh_admin_menu` (`id`, `name`, `icon`, `url`, `pid`, `type`, `permission`, `is_show`, `sort`, `remark`, `status`, `created_time`, `updated_time`) VALUES (17, '友链列表', 'layui-icon-util', 'link/index', 16, 0, 'sys:link:index', 1, 1, '', 1, '2023-03-30 15:00:28', '2023-04-03 08:55:21');
INSERT INTO `oh_admin_menu` (`id`, `name`, `icon`, `url`, `pid`, `type`, `permission`, `is_show`, `sort`, `remark`, `status`, `created_time`, `updated_time`) VALUES (18, '处方管理', 'layui-icon-tabs', '#', 6, 0, 'sys:cf', 1, 2, '处方管理模块', 1, '2025-01-08 07:06:33', '2025-10-25 16:53:03');
INSERT INTO `oh_admin_menu` (`id`, `name`, `icon`, `url`, `pid`, `type`, `permission`, `is_show`, `sort`, `remark`, `status`, `created_time`, `updated_time`) VALUES (19, '接方信息', 'layui-icon-rate-half', 'prescription/index', 18, 0, '1', 1, 3, '', 1, '2025-01-08 07:21:00', '2025-01-08 07:29:03');
INSERT INTO `oh_admin_menu` (`id`, `name`, `icon`, `url`, `pid`, `type`, `permission`, `is_show`, `sort`, `remark`, `status`, `created_time`, `updated_time`) VALUES (20, '审方信息', 'layui-icon-util', 'prescription_audit/index', 51, 0, 'sys:prescription_audit', 1, 4, '', 1, '2025-01-08 07:27:59', '2025-10-25 16:54:21');
INSERT INTO `oh_admin_menu` (`id`, `name`, `icon`, `url`, `pid`, `type`, `permission`, `is_show`, `sort`, `remark`, `status`, `created_time`, `updated_time`) VALUES (21, '调剂信息', 'layui-icon-home', 'adjustment/index', 36, 0, 'sys:adjustment', 1, 4, '', 1, '2025-01-11 11:40:53', '2025-10-25 16:45:19');
INSERT INTO `oh_admin_menu` (`id`, `name`, `icon`, `url`, `pid`, `type`, `permission`, `is_show`, `sort`, `remark`, `status`, `created_time`, `updated_time`) VALUES (22, '复核信息', 'layui-icon-survey', 'herb_decoction_audit/index', 36, 0, 'sys:herb_decoction_audit', 1, 5, '', 1, '2025-01-11 11:43:24', '2025-10-25 16:45:50');
INSERT INTO `oh_admin_menu` (`id`, `name`, `icon`, `url`, `pid`, `type`, `permission`, `is_show`, `sort`, `remark`, `status`, `created_time`, `updated_time`) VALUES (23, '泡药管理', 'layui-icon-component', 'herb_soaking/index', 36, 0, 'sys:herb_soaking', 1, 6, '', 1, '2025-01-11 11:44:30', '2025-10-25 16:46:30');
INSERT INTO `oh_admin_menu` (`id`, `name`, `icon`, `url`, `pid`, `type`, `permission`, `is_show`, `sort`, `remark`, `status`, `created_time`, `updated_time`) VALUES (24, '煎药信息', 'layui-icon-util', 'herbal_decoction_info/index', 36, 0, 'sys:herbal_decoction_info', 1, 7, '', 1, '2025-01-11 11:44:58', '2025-10-25 16:47:05');
INSERT INTO `oh_admin_menu` (`id`, `name`, `icon`, `url`, `pid`, `type`, `permission`, `is_show`, `sort`, `remark`, `status`, `created_time`, `updated_time`) VALUES (25, '包装信息', 'layui-icon-tabs', 'medicine_packing/index', 36, 0, 'sys:medicine_packing', 1, 8, '', 1, '2025-01-11 11:46:06', '2025-10-25 16:47:25');
INSERT INTO `oh_admin_menu` (`id`, `name`, `icon`, `url`, `pid`, `type`, `permission`, `is_show`, `sort`, `remark`, `status`, `created_time`, `updated_time`) VALUES (26, '发货信息', 'layui-icon-auz', 'herb_delivery/index', 36, 0, 'sys:herb_delivery', 1, 9, '', 1, '2025-01-11 11:46:38', '2025-10-25 16:47:54');
INSERT INTO `oh_admin_menu` (`id`, `name`, `icon`, `url`, `pid`, `type`, `permission`, `is_show`, `sort`, `remark`, `status`, `created_time`, `updated_time`) VALUES (27, '查询统计', 'layui-icon-chart-screen', '#', 0, 0, 'sys:query', 1, 3, '', 1, '2025-01-11 11:54:40', '2025-02-18 08:53:46');
INSERT INTO `oh_admin_menu` (`id`, `name`, `icon`, `url`, `pid`, `type`, `permission`, `is_show`, `sort`, `remark`, `status`, `created_time`, `updated_time`) VALUES (28, '设备管理', 'layui-icon-password', '#', 0, 0, 'sys:device', 1, 6, '', 1, '2025-01-11 12:00:54', '2025-10-25 16:59:02');
INSERT INTO `oh_admin_menu` (`id`, `name`, `icon`, `url`, `pid`, `type`, `permission`, `is_show`, `sort`, `remark`, `status`, `created_time`, `updated_time`) VALUES (30, '处方查询', 'layui-icon-fire', 'prescription_query/index', 27, 0, 'sys:presctionQuery', 1, 1, '', 1, '2025-01-12 16:52:48', '2025-02-03 21:10:18');
INSERT INTO `oh_admin_menu` (`id`, `name`, `icon`, `url`, `pid`, `type`, `permission`, `is_show`, `sort`, `remark`, `status`, `created_time`, `updated_time`) VALUES (31, '工作量统计', 'layui-icon-tabs', 'work_statistics/index', 27, 0, 'sys:/index', 1, 2, '', 1, '2025-01-12 16:54:25', '2025-02-03 21:49:53');
INSERT INTO `oh_admin_menu` (`id`, `name`, `icon`, `url`, `pid`, `type`, `permission`, `is_show`, `sort`, `remark`, `status`, `created_time`, `updated_time`) VALUES (33, '医院数据汇总', 'layui-icon-util', 'hospitaldata/index', 27, 0, 'sys:hospitaldata', 1, 4, '', 1, '2025-01-12 16:58:49', '2025-02-18 14:07:42');
INSERT INTO `oh_admin_menu` (`id`, `name`, `icon`, `url`, `pid`, `type`, `permission`, `is_show`, `sort`, `remark`, `status`, `created_time`, `updated_time`) VALUES (34, '质量管控', 'layui-icon-auz', '#', 6, 0, 'zlgk', 1, 4, '', 1, '2025-01-12 17:01:47', '2025-10-25 16:53:27');
INSERT INTO `oh_admin_menu` (`id`, `name`, `icon`, `url`, `pid`, `type`, `permission`, `is_show`, `sort`, `remark`, `status`, `created_time`, `updated_time`) VALUES (35, '抽样信息', 'layui-icon-survey', 'samp_info/index', 34, 0, 'sys:samp_info', 1, 1, '', 1, '2025-01-12 17:03:07', '2025-02-15 14:17:27');
INSERT INTO `oh_admin_menu` (`id`, `name`, `icon`, `url`, `pid`, `type`, `permission`, `is_show`, `sort`, `remark`, `status`, `created_time`, `updated_time`) VALUES (36, '煎药流程管控', 'layui-icon-tabs', '#', 6, 0, 'sys:jylc', 1, 32, '', 1, '2025-01-15 09:13:31', '2025-10-25 16:48:26');
INSERT INTO `oh_admin_menu` (`id`, `name`, `icon`, `url`, `pid`, `type`, `permission`, `is_show`, `sort`, `remark`, `status`, `created_time`, `updated_time`) VALUES (37, '处方扫描', 'layui-icon-cellphone', 'prescription_scan/add', 36, 0, 'sys:prescription_scan', 1, 1, '', 1, '2025-01-16 12:42:52', '2025-01-16 12:42:52');
INSERT INTO `oh_admin_menu` (`id`, `name`, `icon`, `url`, `pid`, `type`, `permission`, `is_show`, `sort`, `remark`, `status`, `created_time`, `updated_time`) VALUES (38, '区块链业务管理', 'layui-icon-vercode', '#', 6, 0, 'sys:chain', 2, 6, '', 2, '2025-01-16 23:14:46', '2025-10-07 16:56:53');
INSERT INTO `oh_admin_menu` (`id`, `name`, `icon`, `url`, `pid`, `type`, `permission`, `is_show`, `sort`, `remark`, `status`, `created_time`, `updated_time`) VALUES (39, '煎药过程数据上链', 'layui-icon-website', 'prescription_scan/index', 38, 0, 'sys:prescription_scan', 1, 1, '', 1, '2025-01-16 23:17:13', '2025-01-24 21:38:55');
INSERT INTO `oh_admin_menu` (`id`, `name`, `icon`, `url`, `pid`, `type`, `permission`, `is_show`, `sort`, `remark`, `status`, `created_time`, `updated_time`) VALUES (40, '区块链浏览器', 'layui-icon-fire', 'blockchain_explorer/index', 38, 0, 'sys:blockchain_explorer', 1, 2, '', 1, '2025-01-16 23:23:36', '2025-01-17 00:56:50');
INSERT INTO `oh_admin_menu` (`id`, `name`, `icon`, `url`, `pid`, `type`, `permission`, `is_show`, `sort`, `remark`, `status`, `created_time`, `updated_time`) VALUES (41, '智慧中药溯源', 'layui-icon-util', 'trace_chinese_medicine/index', 38, 0, 'sys:trace_chinese_medicine', 1, 3, '', 1, '2025-01-16 23:28:55', '2025-01-16 23:28:55');
INSERT INTO `oh_admin_menu` (`id`, `name`, `icon`, `url`, `pid`, `type`, `permission`, `is_show`, `sort`, `remark`, `status`, `created_time`, `updated_time`) VALUES (49, '处方录入', 'layui-icon-util', 'prescription/add', 18, 0, 'sys:prescription_add', 1, 0, '', 1, '2025-10-25 16:24:58', '2025-10-25 16:26:40');
INSERT INTO `oh_admin_menu` (`id`, `name`, `icon`, `url`, `pid`, `type`, `permission`, `is_show`, `sort`, `remark`, `status`, `created_time`, `updated_time`) VALUES (50, '抽检录入', 'layui-icon-login-wechat', 'samp_info/add', 34, 0, 'sys:samp_info_add', 1, 0, '', 1, '2025-10-25 16:30:43', '2025-10-25 16:32:12');
INSERT INTO `oh_admin_menu` (`id`, `name`, `icon`, `url`, `pid`, `type`, `permission`, `is_show`, `sort`, `remark`, `status`, `created_time`, `updated_time`) VALUES (51, '审方管理', 'layui-icon-util', '#', 6, 0, 'sys_audit', 1, 3, '', 1, '2025-10-25 16:44:28', '2025-10-25 16:53:13');
INSERT INTO `oh_admin_menu` (`id`, `name`, `icon`, `url`, `pid`, `type`, `permission`, `is_show`, `sort`, `remark`, `status`, `created_time`, `updated_time`) VALUES (52, '基础数据管理', 'layui-icon-util', '#', 0, 0, 'sys:base', 1, 7, '', 1, '2025-10-25 17:00:09', '2025-10-25 17:00:09');
INSERT INTO `oh_admin_menu` (`id`, `name`, `icon`, `url`, `pid`, `type`, `permission`, `is_show`, `sort`, `remark`, `status`, `created_time`, `updated_time`) VALUES (53, '设备信息', 'layui-icon-util', 'device/index', 28, 0, 'sys:device_index', 1, 1, '', 1, '2025-10-25 17:03:38', '2025-10-25 17:03:38');
INSERT INTO `oh_admin_menu` (`id`, `name`, `icon`, `url`, `pid`, `type`, `permission`, `is_show`, `sort`, `remark`, `status`, `created_time`, `updated_time`) VALUES (54, '设备巡检', 'layui-icon-util', 'deviceRecord/index', 28, 0, 'sys:deviceRecord', 1, 2, '', 1, '2025-10-25 17:08:55', '2025-10-25 17:08:55');
INSERT INTO `oh_admin_menu` (`id`, `name`, `icon`, `url`, `pid`, `type`, `permission`, `is_show`, `sort`, `remark`, `status`, `created_time`, `updated_time`) VALUES (55, '巡检记录统计', 'layui-icon-util', 'deviceRecord/query', 28, 0, 'sys:deviceRecord_query', 1, 4, '', 1, '2025-10-25 17:10:49', '2025-10-25 17:11:54');
INSERT INTO `oh_admin_menu` (`id`, `name`, `icon`, `url`, `pid`, `type`, `permission`, `is_show`, `sort`, `remark`, `status`, `created_time`, `updated_time`) VALUES (56, '库存管理', 'layui-icon-util', '#', 0, 0, 'sys:warehourse', 1, 8, '', 1, '2025-10-26 15:41:24', '2025-10-26 15:41:24');
INSERT INTO `oh_admin_menu` (`id`, `name`, `icon`, `url`, `pid`, `type`, `permission`, `is_show`, `sort`, `remark`, `status`, `created_time`, `updated_time`) VALUES (57, '入库管理', 'layui-icon-util', 'stock_in/index', 56, 0, 'sys:stock_in', 1, 1, '', 1, '2025-10-26 15:42:21', '2025-10-26 18:12:58');
INSERT INTO `oh_admin_menu` (`id`, `name`, `icon`, `url`, `pid`, `type`, `permission`, `is_show`, `sort`, `remark`, `status`, `created_time`, `updated_time`) VALUES (58, '调拨管理', 'layui-icon-login-wechat', 'stock_transfer/index', 56, 0, 'sys:stock_transfer', 1, 2, '', 1, '2025-10-26 15:44:12', '2025-10-26 16:01:12');
INSERT INTO `oh_admin_menu` (`id`, `name`, `icon`, `url`, `pid`, `type`, `permission`, `is_show`, `sort`, `remark`, `status`, `created_time`, `updated_time`) VALUES (59, '库存信息', 'layui-icon-cellphone', 'inventory/index', 56, 0, 'sys:inventory', 1, 3, '', 1, '2025-10-26 15:56:30', '2025-10-26 16:01:23');
INSERT INTO `oh_admin_menu` (`id`, `name`, `icon`, `url`, `pid`, `type`, `permission`, `is_show`, `sort`, `remark`, `status`, `created_time`, `updated_time`) VALUES (60, '盘点数据', 'layui-icon-util', 'stock_check/index', 56, 0, 'sys:stock_check', 1, 4, '', 1, '2025-10-26 16:00:52', '2025-10-26 16:00:52');
INSERT INTO `oh_admin_menu` (`id`, `name`, `icon`, `url`, `pid`, `type`, `permission`, `is_show`, `sort`, `remark`, `status`, `created_time`, `updated_time`) VALUES (61, '药品信息', 'layui-icon-rate', 'drug_basics/index', 52, 0, 'sys:drug_basics', 1, 1, '', 1, '2025-10-26 16:03:36', '2025-10-26 16:03:36');
INSERT INTO `oh_admin_menu` (`id`, `name`, `icon`, `url`, `pid`, `type`, `permission`, `is_show`, `sort`, `remark`, `status`, `created_time`, `updated_time`) VALUES (62, '配伍禁忌信息', 'layui-icon-util', 'compatibility_taboo/index', 52, 0, 'sys:compatibility_taboo', 1, 2, '', 1, '2025-10-26 16:06:15', '2025-10-26 16:08:23');
INSERT INTO `oh_admin_menu` (`id`, `name`, `icon`, `url`, `pid`, `type`, `permission`, `is_show`, `sort`, `remark`, `status`, `created_time`, `updated_time`) VALUES (63, '药品剂量限制信息', 'layui-icon-util', 'drug_dosage/index', 52, 0, 'sys:drug_dosage_limit', 1, 4, '', 1, '2025-10-26 16:08:04', '2025-10-26 16:08:04');
INSERT INTO `oh_admin_menu` (`id`, `name`, `icon`, `url`, `pid`, `type`, `permission`, `is_show`, `sort`, `remark`, `status`, `created_time`, `updated_time`) VALUES (64, '库房管理', 'layui-icon-util', 'warehouse/index', 52, 0, 'sys:warehouse', 1, 4, '', 1, '2025-10-26 18:34:04', '2025-10-26 18:34:04');
INSERT INTO `oh_admin_menu` (`id`, `name`, `icon`, `url`, `pid`, `type`, `permission`, `is_show`, `sort`, `remark`, `status`, `created_time`, `updated_time`) VALUES (65, '医院信息', 'layui-icon-util', 'hospital/index', 52, 0, 'sys:hospital', 1, 8, '', 1, '2025-10-26 21:41:33', '2025-10-26 21:41:33');
INSERT INTO `oh_admin_menu` (`id`, `name`, `icon`, `url`, `pid`, `type`, `permission`, `is_show`, `sort`, `remark`, `status`, `created_time`, `updated_time`) VALUES (66, '出库记录', 'layui-icon-util', 'stock_out/index', 56, 0, 'sys:stock_out', 1, 8, '', 1, '2025-10-31 11:15:41', '2025-10-31 11:16:26');
COMMIT;

-- ----------------------------
-- Table structure for oh_admin_role
-- ----------------------------
DROP TABLE IF EXISTS `oh_admin_role`;
CREATE TABLE `oh_admin_role` (
  `id` int NOT NULL AUTO_INCREMENT COMMENT '用户id',
  `name` varchar(50) NOT NULL DEFAULT '' COMMENT '角色名称',
  `code` varchar(50) NOT NULL DEFAULT '' COMMENT '角色编码',
  `status` tinyint(1) DEFAULT '1' COMMENT '状态 1-启用2-禁用',
  `sort` int DEFAULT '0' COMMENT '排序',
  `created_time` datetime DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_time` datetime DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=12 DEFAULT CHARSET=utf8mb3;

-- ----------------------------
-- Records of oh_admin_role
-- ----------------------------
BEGIN;
INSERT INTO `oh_admin_role` (`id`, `name`, `code`, `status`, `sort`, `created_time`, `updated_time`) VALUES (1, '超级管理员', 'super', 1, 3, '2023-03-27 02:01:56', '2023-03-30 16:29:13');
INSERT INTO `oh_admin_role` (`id`, `name`, `code`, `status`, `sort`, `created_time`, `updated_time`) VALUES (2, '管理员', 'admin', 1, 0, '2023-03-27 14:46:24', '2023-03-27 18:02:14');
INSERT INTO `oh_admin_role` (`id`, `name`, `code`, `status`, `sort`, `created_time`, `updated_time`) VALUES (6, '调剂员', 'adjustment', 1, 0, '2025-01-14 23:31:28', '2025-01-14 23:33:20');
INSERT INTO `oh_admin_role` (`id`, `name`, `code`, `status`, `sort`, `created_time`, `updated_time`) VALUES (7, '复核员', 'herb_decoction_audit', 1, 0, '2025-01-14 23:32:33', '2025-01-14 23:34:27');
INSERT INTO `oh_admin_role` (`id`, `name`, `code`, `status`, `sort`, `created_time`, `updated_time`) VALUES (8, '泡药员', 'herb_soaking', 1, 0, '2025-01-14 23:35:03', '2025-01-14 23:35:03');
INSERT INTO `oh_admin_role` (`id`, `name`, `code`, `status`, `sort`, `created_time`, `updated_time`) VALUES (9, '煎药员', 'herbal_decoction_info', 1, 0, '2025-01-14 23:35:38', '2025-01-14 23:35:38');
INSERT INTO `oh_admin_role` (`id`, `name`, `code`, `status`, `sort`, `created_time`, `updated_time`) VALUES (10, '包装员', 'medicine_packing', 1, 0, '2025-01-14 23:36:07', '2025-01-14 23:36:07');
INSERT INTO `oh_admin_role` (`id`, `name`, `code`, `status`, `sort`, `created_time`, `updated_time`) VALUES (11, '发货员', 'herb_delivery', 1, 0, '2025-01-14 23:36:31', '2025-01-14 23:36:31');
COMMIT;

-- ----------------------------
-- Table structure for oh_admin_role_menu
-- ----------------------------
DROP TABLE IF EXISTS `oh_admin_role_menu`;
CREATE TABLE `oh_admin_role_menu` (
  `id` int NOT NULL AUTO_INCREMENT COMMENT '用户id',
  `menu_id` int NOT NULL DEFAULT '0' COMMENT '菜单id',
  `role_id` int NOT NULL DEFAULT '0' COMMENT '角色id',
  `created_time` datetime DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_time` datetime DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间',
  PRIMARY KEY (`id`),
  KEY `idx_role` (`role_id`),
  KEY `idx_menu` (`menu_id`)
) ENGINE=InnoDB AUTO_INCREMENT=949 DEFAULT CHARSET=utf8mb3;

-- ----------------------------
-- Records of oh_admin_role_menu
-- ----------------------------
BEGIN;
INSERT INTO `oh_admin_role_menu` (`id`, `menu_id`, `role_id`, `created_time`, `updated_time`) VALUES (80, 6, 4, '2025-01-08 09:49:14', '2025-01-08 09:49:14');
INSERT INTO `oh_admin_role_menu` (`id`, `menu_id`, `role_id`, `created_time`, `updated_time`) VALUES (81, 16, 4, '2025-01-08 09:49:14', '2025-01-08 09:49:14');
INSERT INTO `oh_admin_role_menu` (`id`, `menu_id`, `role_id`, `created_time`, `updated_time`) VALUES (82, 17, 4, '2025-01-08 09:49:14', '2025-01-08 09:49:14');
INSERT INTO `oh_admin_role_menu` (`id`, `menu_id`, `role_id`, `created_time`, `updated_time`) VALUES (83, 18, 4, '2025-01-08 09:49:14', '2025-01-08 09:49:14');
INSERT INTO `oh_admin_role_menu` (`id`, `menu_id`, `role_id`, `created_time`, `updated_time`) VALUES (84, 19, 4, '2025-01-08 09:49:14', '2025-01-08 09:49:14');
INSERT INTO `oh_admin_role_menu` (`id`, `menu_id`, `role_id`, `created_time`, `updated_time`) VALUES (85, 20, 4, '2025-01-08 09:49:14', '2025-01-08 09:49:14');
INSERT INTO `oh_admin_role_menu` (`id`, `menu_id`, `role_id`, `created_time`, `updated_time`) VALUES (222, 6, 6, '2025-01-16 13:07:30', '2025-01-16 13:07:30');
INSERT INTO `oh_admin_role_menu` (`id`, `menu_id`, `role_id`, `created_time`, `updated_time`) VALUES (223, 18, 6, '2025-01-16 13:07:30', '2025-01-16 13:07:30');
INSERT INTO `oh_admin_role_menu` (`id`, `menu_id`, `role_id`, `created_time`, `updated_time`) VALUES (224, 19, 6, '2025-01-16 13:07:30', '2025-01-16 13:07:30');
INSERT INTO `oh_admin_role_menu` (`id`, `menu_id`, `role_id`, `created_time`, `updated_time`) VALUES (225, 21, 6, '2025-01-16 13:07:30', '2025-01-16 13:07:30');
INSERT INTO `oh_admin_role_menu` (`id`, `menu_id`, `role_id`, `created_time`, `updated_time`) VALUES (226, 36, 6, '2025-01-16 13:07:30', '2025-01-16 13:07:30');
INSERT INTO `oh_admin_role_menu` (`id`, `menu_id`, `role_id`, `created_time`, `updated_time`) VALUES (227, 37, 6, '2025-01-16 13:07:30', '2025-01-16 13:07:30');
INSERT INTO `oh_admin_role_menu` (`id`, `menu_id`, `role_id`, `created_time`, `updated_time`) VALUES (228, 6, 7, '2025-01-16 13:07:37', '2025-01-16 13:07:37');
INSERT INTO `oh_admin_role_menu` (`id`, `menu_id`, `role_id`, `created_time`, `updated_time`) VALUES (229, 18, 7, '2025-01-16 13:07:37', '2025-01-16 13:07:37');
INSERT INTO `oh_admin_role_menu` (`id`, `menu_id`, `role_id`, `created_time`, `updated_time`) VALUES (230, 19, 7, '2025-01-16 13:07:37', '2025-01-16 13:07:37');
INSERT INTO `oh_admin_role_menu` (`id`, `menu_id`, `role_id`, `created_time`, `updated_time`) VALUES (231, 22, 7, '2025-01-16 13:07:37', '2025-01-16 13:07:37');
INSERT INTO `oh_admin_role_menu` (`id`, `menu_id`, `role_id`, `created_time`, `updated_time`) VALUES (232, 36, 7, '2025-01-16 13:07:37', '2025-01-16 13:07:37');
INSERT INTO `oh_admin_role_menu` (`id`, `menu_id`, `role_id`, `created_time`, `updated_time`) VALUES (233, 37, 7, '2025-01-16 13:07:37', '2025-01-16 13:07:37');
INSERT INTO `oh_admin_role_menu` (`id`, `menu_id`, `role_id`, `created_time`, `updated_time`) VALUES (234, 6, 8, '2025-01-16 13:07:46', '2025-01-16 13:07:46');
INSERT INTO `oh_admin_role_menu` (`id`, `menu_id`, `role_id`, `created_time`, `updated_time`) VALUES (235, 18, 8, '2025-01-16 13:07:46', '2025-01-16 13:07:46');
INSERT INTO `oh_admin_role_menu` (`id`, `menu_id`, `role_id`, `created_time`, `updated_time`) VALUES (236, 19, 8, '2025-01-16 13:07:46', '2025-01-16 13:07:46');
INSERT INTO `oh_admin_role_menu` (`id`, `menu_id`, `role_id`, `created_time`, `updated_time`) VALUES (237, 23, 8, '2025-01-16 13:07:46', '2025-01-16 13:07:46');
INSERT INTO `oh_admin_role_menu` (`id`, `menu_id`, `role_id`, `created_time`, `updated_time`) VALUES (238, 36, 8, '2025-01-16 13:07:46', '2025-01-16 13:07:46');
INSERT INTO `oh_admin_role_menu` (`id`, `menu_id`, `role_id`, `created_time`, `updated_time`) VALUES (239, 37, 8, '2025-01-16 13:07:46', '2025-01-16 13:07:46');
INSERT INTO `oh_admin_role_menu` (`id`, `menu_id`, `role_id`, `created_time`, `updated_time`) VALUES (240, 6, 9, '2025-01-16 13:07:57', '2025-01-16 13:07:57');
INSERT INTO `oh_admin_role_menu` (`id`, `menu_id`, `role_id`, `created_time`, `updated_time`) VALUES (241, 18, 9, '2025-01-16 13:07:57', '2025-01-16 13:07:57');
INSERT INTO `oh_admin_role_menu` (`id`, `menu_id`, `role_id`, `created_time`, `updated_time`) VALUES (242, 24, 9, '2025-01-16 13:07:57', '2025-01-16 13:07:57');
INSERT INTO `oh_admin_role_menu` (`id`, `menu_id`, `role_id`, `created_time`, `updated_time`) VALUES (243, 36, 9, '2025-01-16 13:07:57', '2025-01-16 13:07:57');
INSERT INTO `oh_admin_role_menu` (`id`, `menu_id`, `role_id`, `created_time`, `updated_time`) VALUES (244, 37, 9, '2025-01-16 13:07:57', '2025-01-16 13:07:57');
INSERT INTO `oh_admin_role_menu` (`id`, `menu_id`, `role_id`, `created_time`, `updated_time`) VALUES (245, 6, 10, '2025-01-16 13:08:07', '2025-01-16 13:08:07');
INSERT INTO `oh_admin_role_menu` (`id`, `menu_id`, `role_id`, `created_time`, `updated_time`) VALUES (246, 18, 10, '2025-01-16 13:08:07', '2025-01-16 13:08:07');
INSERT INTO `oh_admin_role_menu` (`id`, `menu_id`, `role_id`, `created_time`, `updated_time`) VALUES (247, 19, 10, '2025-01-16 13:08:07', '2025-01-16 13:08:07');
INSERT INTO `oh_admin_role_menu` (`id`, `menu_id`, `role_id`, `created_time`, `updated_time`) VALUES (248, 25, 10, '2025-01-16 13:08:07', '2025-01-16 13:08:07');
INSERT INTO `oh_admin_role_menu` (`id`, `menu_id`, `role_id`, `created_time`, `updated_time`) VALUES (249, 36, 10, '2025-01-16 13:08:07', '2025-01-16 13:08:07');
INSERT INTO `oh_admin_role_menu` (`id`, `menu_id`, `role_id`, `created_time`, `updated_time`) VALUES (250, 37, 10, '2025-01-16 13:08:07', '2025-01-16 13:08:07');
INSERT INTO `oh_admin_role_menu` (`id`, `menu_id`, `role_id`, `created_time`, `updated_time`) VALUES (251, 6, 11, '2025-01-16 13:08:17', '2025-01-16 13:08:17');
INSERT INTO `oh_admin_role_menu` (`id`, `menu_id`, `role_id`, `created_time`, `updated_time`) VALUES (252, 18, 11, '2025-01-16 13:08:17', '2025-01-16 13:08:17');
INSERT INTO `oh_admin_role_menu` (`id`, `menu_id`, `role_id`, `created_time`, `updated_time`) VALUES (253, 19, 11, '2025-01-16 13:08:17', '2025-01-16 13:08:17');
INSERT INTO `oh_admin_role_menu` (`id`, `menu_id`, `role_id`, `created_time`, `updated_time`) VALUES (254, 26, 11, '2025-01-16 13:08:17', '2025-01-16 13:08:17');
INSERT INTO `oh_admin_role_menu` (`id`, `menu_id`, `role_id`, `created_time`, `updated_time`) VALUES (255, 36, 11, '2025-01-16 13:08:17', '2025-01-16 13:08:17');
INSERT INTO `oh_admin_role_menu` (`id`, `menu_id`, `role_id`, `created_time`, `updated_time`) VALUES (256, 37, 11, '2025-01-16 13:08:17', '2025-01-16 13:08:17');
INSERT INTO `oh_admin_role_menu` (`id`, `menu_id`, `role_id`, `created_time`, `updated_time`) VALUES (616, 6, 2, '2025-10-25 16:25:30', '2025-10-25 16:25:30');
INSERT INTO `oh_admin_role_menu` (`id`, `menu_id`, `role_id`, `created_time`, `updated_time`) VALUES (617, 16, 2, '2025-10-25 16:25:30', '2025-10-25 16:25:30');
INSERT INTO `oh_admin_role_menu` (`id`, `menu_id`, `role_id`, `created_time`, `updated_time`) VALUES (618, 17, 2, '2025-10-25 16:25:30', '2025-10-25 16:25:30');
INSERT INTO `oh_admin_role_menu` (`id`, `menu_id`, `role_id`, `created_time`, `updated_time`) VALUES (619, 18, 2, '2025-10-25 16:25:30', '2025-10-25 16:25:30');
INSERT INTO `oh_admin_role_menu` (`id`, `menu_id`, `role_id`, `created_time`, `updated_time`) VALUES (620, 49, 2, '2025-10-25 16:25:30', '2025-10-25 16:25:30');
INSERT INTO `oh_admin_role_menu` (`id`, `menu_id`, `role_id`, `created_time`, `updated_time`) VALUES (621, 20, 2, '2025-10-25 16:25:30', '2025-10-25 16:25:30');
INSERT INTO `oh_admin_role_menu` (`id`, `menu_id`, `role_id`, `created_time`, `updated_time`) VALUES (622, 19, 2, '2025-10-25 16:25:30', '2025-10-25 16:25:30');
INSERT INTO `oh_admin_role_menu` (`id`, `menu_id`, `role_id`, `created_time`, `updated_time`) VALUES (907, 6, 1, '2025-10-31 11:16:00', '2025-10-31 11:16:00');
INSERT INTO `oh_admin_role_menu` (`id`, `menu_id`, `role_id`, `created_time`, `updated_time`) VALUES (908, 18, 1, '2025-10-31 11:16:00', '2025-10-31 11:16:00');
INSERT INTO `oh_admin_role_menu` (`id`, `menu_id`, `role_id`, `created_time`, `updated_time`) VALUES (909, 49, 1, '2025-10-31 11:16:00', '2025-10-31 11:16:00');
INSERT INTO `oh_admin_role_menu` (`id`, `menu_id`, `role_id`, `created_time`, `updated_time`) VALUES (910, 19, 1, '2025-10-31 11:16:00', '2025-10-31 11:16:00');
INSERT INTO `oh_admin_role_menu` (`id`, `menu_id`, `role_id`, `created_time`, `updated_time`) VALUES (911, 51, 1, '2025-10-31 11:16:00', '2025-10-31 11:16:00');
INSERT INTO `oh_admin_role_menu` (`id`, `menu_id`, `role_id`, `created_time`, `updated_time`) VALUES (912, 20, 1, '2025-10-31 11:16:00', '2025-10-31 11:16:00');
INSERT INTO `oh_admin_role_menu` (`id`, `menu_id`, `role_id`, `created_time`, `updated_time`) VALUES (913, 34, 1, '2025-10-31 11:16:00', '2025-10-31 11:16:00');
INSERT INTO `oh_admin_role_menu` (`id`, `menu_id`, `role_id`, `created_time`, `updated_time`) VALUES (914, 50, 1, '2025-10-31 11:16:00', '2025-10-31 11:16:00');
INSERT INTO `oh_admin_role_menu` (`id`, `menu_id`, `role_id`, `created_time`, `updated_time`) VALUES (915, 35, 1, '2025-10-31 11:16:00', '2025-10-31 11:16:00');
INSERT INTO `oh_admin_role_menu` (`id`, `menu_id`, `role_id`, `created_time`, `updated_time`) VALUES (916, 36, 1, '2025-10-31 11:16:00', '2025-10-31 11:16:00');
INSERT INTO `oh_admin_role_menu` (`id`, `menu_id`, `role_id`, `created_time`, `updated_time`) VALUES (917, 37, 1, '2025-10-31 11:16:00', '2025-10-31 11:16:00');
INSERT INTO `oh_admin_role_menu` (`id`, `menu_id`, `role_id`, `created_time`, `updated_time`) VALUES (918, 21, 1, '2025-10-31 11:16:00', '2025-10-31 11:16:00');
INSERT INTO `oh_admin_role_menu` (`id`, `menu_id`, `role_id`, `created_time`, `updated_time`) VALUES (919, 22, 1, '2025-10-31 11:16:00', '2025-10-31 11:16:00');
INSERT INTO `oh_admin_role_menu` (`id`, `menu_id`, `role_id`, `created_time`, `updated_time`) VALUES (920, 23, 1, '2025-10-31 11:16:00', '2025-10-31 11:16:00');
INSERT INTO `oh_admin_role_menu` (`id`, `menu_id`, `role_id`, `created_time`, `updated_time`) VALUES (921, 24, 1, '2025-10-31 11:16:00', '2025-10-31 11:16:00');
INSERT INTO `oh_admin_role_menu` (`id`, `menu_id`, `role_id`, `created_time`, `updated_time`) VALUES (922, 25, 1, '2025-10-31 11:16:00', '2025-10-31 11:16:00');
INSERT INTO `oh_admin_role_menu` (`id`, `menu_id`, `role_id`, `created_time`, `updated_time`) VALUES (923, 26, 1, '2025-10-31 11:16:00', '2025-10-31 11:16:00');
INSERT INTO `oh_admin_role_menu` (`id`, `menu_id`, `role_id`, `created_time`, `updated_time`) VALUES (924, 27, 1, '2025-10-31 11:16:00', '2025-10-31 11:16:00');
INSERT INTO `oh_admin_role_menu` (`id`, `menu_id`, `role_id`, `created_time`, `updated_time`) VALUES (925, 30, 1, '2025-10-31 11:16:00', '2025-10-31 11:16:00');
INSERT INTO `oh_admin_role_menu` (`id`, `menu_id`, `role_id`, `created_time`, `updated_time`) VALUES (926, 31, 1, '2025-10-31 11:16:00', '2025-10-31 11:16:00');
INSERT INTO `oh_admin_role_menu` (`id`, `menu_id`, `role_id`, `created_time`, `updated_time`) VALUES (927, 33, 1, '2025-10-31 11:16:00', '2025-10-31 11:16:00');
INSERT INTO `oh_admin_role_menu` (`id`, `menu_id`, `role_id`, `created_time`, `updated_time`) VALUES (928, 28, 1, '2025-10-31 11:16:00', '2025-10-31 11:16:00');
INSERT INTO `oh_admin_role_menu` (`id`, `menu_id`, `role_id`, `created_time`, `updated_time`) VALUES (929, 53, 1, '2025-10-31 11:16:00', '2025-10-31 11:16:00');
INSERT INTO `oh_admin_role_menu` (`id`, `menu_id`, `role_id`, `created_time`, `updated_time`) VALUES (930, 54, 1, '2025-10-31 11:16:00', '2025-10-31 11:16:00');
INSERT INTO `oh_admin_role_menu` (`id`, `menu_id`, `role_id`, `created_time`, `updated_time`) VALUES (931, 55, 1, '2025-10-31 11:16:00', '2025-10-31 11:16:00');
INSERT INTO `oh_admin_role_menu` (`id`, `menu_id`, `role_id`, `created_time`, `updated_time`) VALUES (932, 52, 1, '2025-10-31 11:16:00', '2025-10-31 11:16:00');
INSERT INTO `oh_admin_role_menu` (`id`, `menu_id`, `role_id`, `created_time`, `updated_time`) VALUES (933, 61, 1, '2025-10-31 11:16:00', '2025-10-31 11:16:00');
INSERT INTO `oh_admin_role_menu` (`id`, `menu_id`, `role_id`, `created_time`, `updated_time`) VALUES (934, 62, 1, '2025-10-31 11:16:00', '2025-10-31 11:16:00');
INSERT INTO `oh_admin_role_menu` (`id`, `menu_id`, `role_id`, `created_time`, `updated_time`) VALUES (935, 63, 1, '2025-10-31 11:16:00', '2025-10-31 11:16:00');
INSERT INTO `oh_admin_role_menu` (`id`, `menu_id`, `role_id`, `created_time`, `updated_time`) VALUES (936, 64, 1, '2025-10-31 11:16:00', '2025-10-31 11:16:00');
INSERT INTO `oh_admin_role_menu` (`id`, `menu_id`, `role_id`, `created_time`, `updated_time`) VALUES (937, 65, 1, '2025-10-31 11:16:00', '2025-10-31 11:16:00');
INSERT INTO `oh_admin_role_menu` (`id`, `menu_id`, `role_id`, `created_time`, `updated_time`) VALUES (938, 56, 1, '2025-10-31 11:16:00', '2025-10-31 11:16:00');
INSERT INTO `oh_admin_role_menu` (`id`, `menu_id`, `role_id`, `created_time`, `updated_time`) VALUES (939, 57, 1, '2025-10-31 11:16:00', '2025-10-31 11:16:00');
INSERT INTO `oh_admin_role_menu` (`id`, `menu_id`, `role_id`, `created_time`, `updated_time`) VALUES (940, 66, 1, '2025-10-31 11:16:00', '2025-10-31 11:16:00');
INSERT INTO `oh_admin_role_menu` (`id`, `menu_id`, `role_id`, `created_time`, `updated_time`) VALUES (941, 58, 1, '2025-10-31 11:16:00', '2025-10-31 11:16:00');
INSERT INTO `oh_admin_role_menu` (`id`, `menu_id`, `role_id`, `created_time`, `updated_time`) VALUES (942, 59, 1, '2025-10-31 11:16:00', '2025-10-31 11:16:00');
INSERT INTO `oh_admin_role_menu` (`id`, `menu_id`, `role_id`, `created_time`, `updated_time`) VALUES (943, 60, 1, '2025-10-31 11:16:00', '2025-10-31 11:16:00');
INSERT INTO `oh_admin_role_menu` (`id`, `menu_id`, `role_id`, `created_time`, `updated_time`) VALUES (944, 1, 1, '2025-10-31 11:16:00', '2025-10-31 11:16:00');
INSERT INTO `oh_admin_role_menu` (`id`, `menu_id`, `role_id`, `created_time`, `updated_time`) VALUES (945, 2, 1, '2025-10-31 11:16:00', '2025-10-31 11:16:00');
INSERT INTO `oh_admin_role_menu` (`id`, `menu_id`, `role_id`, `created_time`, `updated_time`) VALUES (946, 5, 1, '2025-10-31 11:16:00', '2025-10-31 11:16:00');
INSERT INTO `oh_admin_role_menu` (`id`, `menu_id`, `role_id`, `created_time`, `updated_time`) VALUES (947, 4, 1, '2025-10-31 11:16:00', '2025-10-31 11:16:00');
INSERT INTO `oh_admin_role_menu` (`id`, `menu_id`, `role_id`, `created_time`, `updated_time`) VALUES (948, 3, 1, '2025-10-31 11:16:00', '2025-10-31 11:16:00');
COMMIT;

-- ----------------------------
-- Table structure for oh_admin_user
-- ----------------------------
DROP TABLE IF EXISTS `oh_admin_user`;
CREATE TABLE `oh_admin_user` (
  `id` int NOT NULL AUTO_INCREMENT COMMENT '用户id',
  `realname` varchar(50) NOT NULL DEFAULT '' COMMENT '真实姓名',
  `username` varchar(100) NOT NULL DEFAULT '' COMMENT '登录用户名',
  `gender` tinyint(1) NOT NULL DEFAULT '0' COMMENT '性别:1男 2女 3保密',
  `avatar` varchar(255) NOT NULL DEFAULT '' COMMENT '头像',
  `mobile` varchar(11) NOT NULL DEFAULT '' COMMENT '手机号',
  `email` varchar(100) NOT NULL DEFAULT '' COMMENT '邮箱地址',
  `address` varchar(255) DEFAULT '' COMMENT '地址',
  `password` varchar(150) NOT NULL DEFAULT '' COMMENT '登录密码',
  `salt` varchar(30) NOT NULL DEFAULT '' COMMENT '盐加密',
  `intro` varchar(255) DEFAULT '' COMMENT '备注',
  `status` tinyint(1) NOT NULL DEFAULT '1' COMMENT '状态：1正常 2禁用',
  `login_num` int NOT NULL DEFAULT '0' COMMENT '登录次数',
  `login_ip` varchar(20) NOT NULL DEFAULT '' COMMENT '最近登录ip',
  `login_time` int NOT NULL DEFAULT '0' COMMENT '最近登录时间',
  `created_time` datetime DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_time` datetime DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间',
  `app_token` varchar(100) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL DEFAULT '',
  PRIMARY KEY (`id`),
  KEY `idx_username` (`username`)
) ENGINE=InnoDB AUTO_INCREMENT=9 DEFAULT CHARSET=utf8mb3;

-- ----------------------------
-- Records of oh_admin_user
-- ----------------------------
BEGIN;
INSERT INTO `oh_admin_user` (`id`, `realname`, `username`, `gender`, `avatar`, `mobile`, `email`, `address`, `password`, `salt`, `intro`, `status`, `login_num`, `login_ip`, `login_time`, `created_time`, `updated_time`, `app_token`) VALUES (1, '管理员', 'admin', 2, '', '123456', 'sd_mwq@163.com', '北京市', '52af3ce8a82f62707789fe00899ed3f0', '123456', 'test', 1, 196, '::1', 1761922459, '2023-03-16 09:38:04', '2025-10-31 22:54:19', 'asdeerfsd');
INSERT INTO `oh_admin_user` (`id`, `realname`, `username`, `gender`, `avatar`, `mobile`, `email`, `address`, `password`, `salt`, `intro`, `status`, `login_num`, `login_ip`, `login_time`, `created_time`, `updated_time`, `app_token`) VALUES (2, '发货员工', '10000', 1, '/static/uploads/images/user/20250107/1736216388.png', '15539266902', '3231@qq.com', '321', '52af3ce8a82f62707789fe00899ed3f0', '123456', '  321', 1, 7, '::1', 1759825699, '2025-01-07 10:19:50', '2025-10-07 16:28:19', 'asdeerfsd');
INSERT INTO `oh_admin_user` (`id`, `realname`, `username`, `gender`, `avatar`, `mobile`, `email`, `address`, `password`, `salt`, `intro`, `status`, `login_num`, `login_ip`, `login_time`, `created_time`, `updated_time`, `app_token`) VALUES (3, '包装员工', '9000', 2, '/static/uploads/images/user/20250108/1736293145.png', '13526798791', 'xx@qq.com', 'hebi', '2d408aeb7257b85d5ad3ff131081e800', 'j9rq3w', '    ', 1, 12, '::1', 1741791523, '2025-01-08 07:39:43', '2025-10-05 06:43:59', 'asdeerfsd');
INSERT INTO `oh_admin_user` (`id`, `realname`, `username`, `gender`, `avatar`, `mobile`, `email`, `address`, `password`, `salt`, `intro`, `status`, `login_num`, `login_ip`, `login_time`, `created_time`, `updated_time`, `app_token`) VALUES (4, '煎药员工', '8000', 1, '/static/uploads/images/user/20250114/1736869022.png', '15539266904', 'lmm@163.com', '郑州二七区', '01a7c22d3150ac793ed5e4a63c6242c7', 'cojs7h', '  ', 1, 6, '::1', 1741791474, '2025-01-14 23:38:19', '2025-10-05 06:44:02', 'asdeerfsd');
INSERT INTO `oh_admin_user` (`id`, `realname`, `username`, `gender`, `avatar`, `mobile`, `email`, `address`, `password`, `salt`, `intro`, `status`, `login_num`, `login_ip`, `login_time`, `created_time`, `updated_time`, `app_token`) VALUES (5, '泡药员工', '7000', 0, '/static/uploads/images/user/20250114/1736869189.png', '13526798791', 'xx@163.com', '郑州高新区', '26de36ea532ebb2c100f74444074a839', 'o8o6fu', '   ', 1, 5, '::1', 1741791428, '2025-01-14 23:39:50', '2025-10-05 06:44:06', 'asdeerfsd1');
INSERT INTO `oh_admin_user` (`id`, `realname`, `username`, `gender`, `avatar`, `mobile`, `email`, `address`, `password`, `salt`, `intro`, `status`, `login_num`, `login_ip`, `login_time`, `created_time`, `updated_time`, `app_token`) VALUES (6, '复核员工', '6000', 1, '/static/uploads/images/user/20250114/1736869201.png', '13526789098', 'x@163.com', 'xxx', 'c2d6b797205651899aa5173a3f6787ad', 'j8jtto', '  ', 1, 6, '::1', 1741791405, '2025-01-14 23:40:41', '2025-10-05 06:44:09', 'asdeerfsd34');
INSERT INTO `oh_admin_user` (`id`, `realname`, `username`, `gender`, `avatar`, `mobile`, `email`, `address`, `password`, `salt`, `intro`, `status`, `login_num`, `login_ip`, `login_time`, `created_time`, `updated_time`, `app_token`) VALUES (7, '调剂员工', '5000', 1, '/static/uploads/images/user/20250116/1737003710.png', '15539266902', 'beiqiaoyuanmu@163.com', 'xx', 'da431ad0f82f854573c3dda16c26186c', '2ncb3p', ' ', 1, 9, '::1', 1759825279, '2025-01-16 13:02:32', '2025-10-07 16:21:19', 'asdeerfsd');
COMMIT;

-- ----------------------------
-- Table structure for oh_admin_user_role
-- ----------------------------
DROP TABLE IF EXISTS `oh_admin_user_role`;
CREATE TABLE `oh_admin_user_role` (
  `id` int NOT NULL AUTO_INCREMENT COMMENT '用户id',
  `user_id` int NOT NULL DEFAULT '0' COMMENT '用户id',
  `role_id` int NOT NULL DEFAULT '0' COMMENT '角色id',
  `created_time` datetime DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_time` datetime DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间',
  PRIMARY KEY (`id`),
  KEY `idx_user_role` (`user_id`,`role_id`)
) ENGINE=InnoDB AUTO_INCREMENT=24 DEFAULT CHARSET=utf8mb3;

-- ----------------------------
-- Records of oh_admin_user_role
-- ----------------------------
BEGIN;
INSERT INTO `oh_admin_user_role` (`id`, `user_id`, `role_id`, `created_time`, `updated_time`) VALUES (1, 1, 1, '2023-03-27 02:02:14', NULL);
INSERT INTO `oh_admin_user_role` (`id`, `user_id`, `role_id`, `created_time`, `updated_time`) VALUES (16, 7, 6, '2025-01-16 13:04:44', '2025-01-16 13:04:44');
INSERT INTO `oh_admin_user_role` (`id`, `user_id`, `role_id`, `created_time`, `updated_time`) VALUES (17, 6, 7, '2025-01-16 13:04:56', '2025-01-16 13:04:56');
INSERT INTO `oh_admin_user_role` (`id`, `user_id`, `role_id`, `created_time`, `updated_time`) VALUES (18, 5, 8, '2025-01-16 13:05:10', '2025-01-16 13:05:10');
INSERT INTO `oh_admin_user_role` (`id`, `user_id`, `role_id`, `created_time`, `updated_time`) VALUES (19, 4, 9, '2025-01-16 13:05:25', '2025-01-16 13:05:25');
INSERT INTO `oh_admin_user_role` (`id`, `user_id`, `role_id`, `created_time`, `updated_time`) VALUES (20, 3, 10, '2025-01-16 13:05:43', '2025-01-16 13:05:43');
INSERT INTO `oh_admin_user_role` (`id`, `user_id`, `role_id`, `created_time`, `updated_time`) VALUES (21, 2, 11, '2025-01-16 13:06:19', '2025-01-16 13:06:19');
INSERT INTO `oh_admin_user_role` (`id`, `user_id`, `role_id`, `created_time`, `updated_time`) VALUES (23, 8, 6, '2025-10-24 23:04:40', '2025-10-24 23:04:40');
COMMIT;

-- ----------------------------
-- Table structure for oh_link
-- ----------------------------
DROP TABLE IF EXISTS `oh_link`;
CREATE TABLE `oh_link` (
  `id` int NOT NULL AUTO_INCREMENT COMMENT '用户id',
  `name` varchar(100) NOT NULL DEFAULT '' COMMENT '友链名称',
  `url` varchar(255) NOT NULL DEFAULT '' COMMENT '友链地址',
  `image` varchar(255) DEFAULT '' COMMENT 'logo',
  `status` tinyint(1) NOT NULL DEFAULT '1' COMMENT '状态 1-启用 2-禁用',
  `sort` int NOT NULL DEFAULT '0' COMMENT '排序',
  `created_time` datetime DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_time` datetime DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8mb3;

-- ----------------------------
-- Records of oh_link
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for prescription
-- ----------------------------
DROP TABLE IF EXISTS `prescription`;
CREATE TABLE `prescription` (
  `ID` int NOT NULL AUTO_INCREMENT,
  `deletion_number` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
  `Is_decoction` int NOT NULL,
  `barcode_scan` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
  `hospital_id` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `hospital_name` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
  `prescription_number` varchar(60) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `decoction_method` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
  `prescription_type` int DEFAULT NULL,
  `patient_name` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
  `patient_sex` int DEFAULT NULL,
  `patient_age` int DEFAULT NULL,
  `patient_phone` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
  `patient_address` text CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci,
  `department_name` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
  `inpatient_area` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
  `ward_name` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
  `sick_bed` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
  `diagnosis_result` text CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci,
  `dosage` int NOT NULL,
  `administration_method` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
  `administration_count` int NOT NULL,
  `package_count` int NOT NULL,
  `decoction_scheme` int NOT NULL,
  `one_time_dosage` int NOT NULL,
  `two_time_dosage` int NOT NULL,
  `soak_water_amount` int DEFAULT NULL,
  `soak_time` int NOT NULL,
  `label_number` int DEFAULT NULL,
  `remarks` text CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci,
  `doctor_name` varchar(10) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
  `footnote` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
  `drug_pickup_time` varchar(255) NOT NULL,
  `drug_pickup_number` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
  `order_time` varchar(255) DEFAULT NULL,
  `do_time` varchar(255) DEFAULT NULL,
  `do_person` varchar(10) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `distribution_company` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
  `distribution_address` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
  `distribution_phone` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
  `distribution_type` varchar(10) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
  `administration_way` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
  `additional_remarks_a` varchar(150) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
  `additional_remarks_b` varchar(150) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
  `drug_confirmation` int DEFAULT NULL,
  `logistics_state` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
  `query_time` varchar(255) DEFAULT NULL,
  `query_person` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
  `drug_count` int DEFAULT NULL,
  `machine_room` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
  `state_post_back` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
  `subside_time` int DEFAULT NULL,
  `current_state` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '接方',
  PRIMARY KEY (`ID`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=86 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='处方信息表';

-- ----------------------------
-- Records of prescription
-- ----------------------------
BEGIN;
INSERT INTO `prescription` (`ID`, `deletion_number`, `Is_decoction`, `barcode_scan`, `hospital_id`, `hospital_name`, `prescription_number`, `decoction_method`, `prescription_type`, `patient_name`, `patient_sex`, `patient_age`, `patient_phone`, `patient_address`, `department_name`, `inpatient_area`, `ward_name`, `sick_bed`, `diagnosis_result`, `dosage`, `administration_method`, `administration_count`, `package_count`, `decoction_scheme`, `one_time_dosage`, `two_time_dosage`, `soak_water_amount`, `soak_time`, `label_number`, `remarks`, `doctor_name`, `footnote`, `drug_pickup_time`, `drug_pickup_number`, `order_time`, `do_time`, `do_person`, `distribution_company`, `distribution_address`, `distribution_phone`, `distribution_type`, `administration_way`, `additional_remarks_a`, `additional_remarks_b`, `drug_confirmation`, `logistics_state`, `query_time`, `query_person`, `drug_count`, `machine_room`, `state_post_back`, `subside_time`, `current_state`) VALUES (68, '1', 1, '', '10001', '河南省人民医院', '2025688783', '', 0, '诸葛四郎', 1, 22, '', '', '1', '', '', '', '20', 7, '内服', 2, 200, 3, 0, 0, 0, 0, 0, '', '赵越', '', '', '', '', '2025-02-26 13:00:00', '', '', '', '', '', '3', '', '3', 0, '', '', '', 6, '', '', 0, '发货');
INSERT INTO `prescription` (`ID`, `deletion_number`, `Is_decoction`, `barcode_scan`, `hospital_id`, `hospital_name`, `prescription_number`, `decoction_method`, `prescription_type`, `patient_name`, `patient_sex`, `patient_age`, `patient_phone`, `patient_address`, `department_name`, `inpatient_area`, `ward_name`, `sick_bed`, `diagnosis_result`, `dosage`, `administration_method`, `administration_count`, `package_count`, `decoction_scheme`, `one_time_dosage`, `two_time_dosage`, `soak_water_amount`, `soak_time`, `label_number`, `remarks`, `doctor_name`, `footnote`, `drug_pickup_time`, `drug_pickup_number`, `order_time`, `do_time`, `do_person`, `distribution_company`, `distribution_address`, `distribution_phone`, `distribution_type`, `administration_way`, `additional_remarks_a`, `additional_remarks_b`, `drug_confirmation`, `logistics_state`, `query_time`, `query_person`, `drug_count`, `machine_room`, `state_post_back`, `subside_time`, `current_state`) VALUES (69, '2', 0, '', '10001', '郑州人民医院', '20250214002', '', 0, '徐山', 1, 19, '', '', '2', '', '', '', '20250214002', 7, '内服', 2, 200, 3, 0, 0, 200, 40, 0, '', '医生', '', '2025-02-15', '', '', '2025-02-27 0:00:00', '', '', '', '', '', '3', '', '3', 0, '', '', '', 12, '', '', 0, '审核');
INSERT INTO `prescription` (`ID`, `deletion_number`, `Is_decoction`, `barcode_scan`, `hospital_id`, `hospital_name`, `prescription_number`, `decoction_method`, `prescription_type`, `patient_name`, `patient_sex`, `patient_age`, `patient_phone`, `patient_address`, `department_name`, `inpatient_area`, `ward_name`, `sick_bed`, `diagnosis_result`, `dosage`, `administration_method`, `administration_count`, `package_count`, `decoction_scheme`, `one_time_dosage`, `two_time_dosage`, `soak_water_amount`, `soak_time`, `label_number`, `remarks`, `doctor_name`, `footnote`, `drug_pickup_time`, `drug_pickup_number`, `order_time`, `do_time`, `do_person`, `distribution_company`, `distribution_address`, `distribution_phone`, `distribution_type`, `administration_way`, `additional_remarks_a`, `additional_remarks_b`, `drug_confirmation`, `logistics_state`, `query_time`, `query_person`, `drug_count`, `machine_room`, `state_post_back`, `subside_time`, `current_state`) VALUES (70, '3', 0, '', '10001', '河南省人民医院', '20250203001', '', 0, '赵亮亮', 1, 22, '', '', '1', '', '', '', '20250203001', 7, '内服', 2, 200, 3, 0, 0, 0, 0, 0, '', 'xx', '', '', '', '', '2025-02-28 0:00:00', '', '', '', '', '', '3', '', '3', 0, '', '', '', 32, '', '', 0, '接方');
INSERT INTO `prescription` (`ID`, `deletion_number`, `Is_decoction`, `barcode_scan`, `hospital_id`, `hospital_name`, `prescription_number`, `decoction_method`, `prescription_type`, `patient_name`, `patient_sex`, `patient_age`, `patient_phone`, `patient_address`, `department_name`, `inpatient_area`, `ward_name`, `sick_bed`, `diagnosis_result`, `dosage`, `administration_method`, `administration_count`, `package_count`, `decoction_scheme`, `one_time_dosage`, `two_time_dosage`, `soak_water_amount`, `soak_time`, `label_number`, `remarks`, `doctor_name`, `footnote`, `drug_pickup_time`, `drug_pickup_number`, `order_time`, `do_time`, `do_person`, `distribution_company`, `distribution_address`, `distribution_phone`, `distribution_type`, `administration_way`, `additional_remarks_a`, `additional_remarks_b`, `drug_confirmation`, `logistics_state`, `query_time`, `query_person`, `drug_count`, `machine_room`, `state_post_back`, `subside_time`, `current_state`) VALUES (71, '4', 1, '', '10002', '河南省中医院', '22025099323', '', 0, '张一飞', 1, 45, '', '', '1', '', '', '', 'wewqe', 2, '内服', 2, 200, 3, 0, 0, 0, 0, 0, '', '2ewr', '', '', '', '', '2025-02-28 0:00:00', '', '', '', '', '', '3', '', '3', 0, '', '', '', 10, '', '', 0, '接方');
INSERT INTO `prescription` (`ID`, `deletion_number`, `Is_decoction`, `barcode_scan`, `hospital_id`, `hospital_name`, `prescription_number`, `decoction_method`, `prescription_type`, `patient_name`, `patient_sex`, `patient_age`, `patient_phone`, `patient_address`, `department_name`, `inpatient_area`, `ward_name`, `sick_bed`, `diagnosis_result`, `dosage`, `administration_method`, `administration_count`, `package_count`, `decoction_scheme`, `one_time_dosage`, `two_time_dosage`, `soak_water_amount`, `soak_time`, `label_number`, `remarks`, `doctor_name`, `footnote`, `drug_pickup_time`, `drug_pickup_number`, `order_time`, `do_time`, `do_person`, `distribution_company`, `distribution_address`, `distribution_phone`, `distribution_type`, `administration_way`, `additional_remarks_a`, `additional_remarks_b`, `drug_confirmation`, `logistics_state`, `query_time`, `query_person`, `drug_count`, `machine_room`, `state_post_back`, `subside_time`, `current_state`) VALUES (72, '', 1, '', '10001', '河南省人民医院', '202509003213', '', 1, '张莉莉', 1, 22, '', '', '', '', '', '', '202509003213', 2, '内服', 7, 200, 6, 0, 0, 0, 0, 0, '', '王维', '', '', '', '', '2025-02-21  0:00:00', '', '', '', '', '', '6', '', '6', 0, '', '', '', 0, '', '', 0, '接方');
INSERT INTO `prescription` (`ID`, `deletion_number`, `Is_decoction`, `barcode_scan`, `hospital_id`, `hospital_name`, `prescription_number`, `decoction_method`, `prescription_type`, `patient_name`, `patient_sex`, `patient_age`, `patient_phone`, `patient_address`, `department_name`, `inpatient_area`, `ward_name`, `sick_bed`, `diagnosis_result`, `dosage`, `administration_method`, `administration_count`, `package_count`, `decoction_scheme`, `one_time_dosage`, `two_time_dosage`, `soak_water_amount`, `soak_time`, `label_number`, `remarks`, `doctor_name`, `footnote`, `drug_pickup_time`, `drug_pickup_number`, `order_time`, `do_time`, `do_person`, `distribution_company`, `distribution_address`, `distribution_phone`, `distribution_type`, `administration_way`, `additional_remarks_a`, `additional_remarks_b`, `drug_confirmation`, `logistics_state`, `query_time`, `query_person`, `drug_count`, `machine_room`, `state_post_back`, `subside_time`, `current_state`) VALUES (73, '2', 0, '', '10001', '河南省人民医院', '20980000', '', 0, '赵莉莉', 2, 18, '', '', '2', '', '', '', '20980000', 7, '内服', 2, 200, 4, 0, 0, 0, 0, 0, '', '莉莉', '', '', '', '', '2025-02-23  0:00:00', '', '', '', '', '', '4', '', '4', 0, '', '', '', 0, '', '', 0, '接方');
INSERT INTO `prescription` (`ID`, `deletion_number`, `Is_decoction`, `barcode_scan`, `hospital_id`, `hospital_name`, `prescription_number`, `decoction_method`, `prescription_type`, `patient_name`, `patient_sex`, `patient_age`, `patient_phone`, `patient_address`, `department_name`, `inpatient_area`, `ward_name`, `sick_bed`, `diagnosis_result`, `dosage`, `administration_method`, `administration_count`, `package_count`, `decoction_scheme`, `one_time_dosage`, `two_time_dosage`, `soak_water_amount`, `soak_time`, `label_number`, `remarks`, `doctor_name`, `footnote`, `drug_pickup_time`, `drug_pickup_number`, `order_time`, `do_time`, `do_person`, `distribution_company`, `distribution_address`, `distribution_phone`, `distribution_type`, `administration_way`, `additional_remarks_a`, `additional_remarks_b`, `drug_confirmation`, `logistics_state`, `query_time`, `query_person`, `drug_count`, `machine_room`, `state_post_back`, `subside_time`, `current_state`) VALUES (74, '', 1, '', '10001', '河南省人民医院', '2312321334324', '', 1, '王菲', 1, 2, '', '', '', '', '', '', '2312321334324', 7, '内服', 2, 200, 3, 0, 0, 0, 0, 0, '', '测试', '', '', '', '', '2025-02-24  0:00:00', '', '', '', '', '', '3', '', '3', 0, '', '', '', 0, '', '', 0, '接方');
INSERT INTO `prescription` (`ID`, `deletion_number`, `Is_decoction`, `barcode_scan`, `hospital_id`, `hospital_name`, `prescription_number`, `decoction_method`, `prescription_type`, `patient_name`, `patient_sex`, `patient_age`, `patient_phone`, `patient_address`, `department_name`, `inpatient_area`, `ward_name`, `sick_bed`, `diagnosis_result`, `dosage`, `administration_method`, `administration_count`, `package_count`, `decoction_scheme`, `one_time_dosage`, `two_time_dosage`, `soak_water_amount`, `soak_time`, `label_number`, `remarks`, `doctor_name`, `footnote`, `drug_pickup_time`, `drug_pickup_number`, `order_time`, `do_time`, `do_person`, `distribution_company`, `distribution_address`, `distribution_phone`, `distribution_type`, `administration_way`, `additional_remarks_a`, `additional_remarks_b`, `drug_confirmation`, `logistics_state`, `query_time`, `query_person`, `drug_count`, `machine_room`, `state_post_back`, `subside_time`, `current_state`) VALUES (76, '', 1, '', '10001', '河南省人民医院', '32321321', '', 0, '测试', 1, 2, '', '', '', '', '', '', '32321321', 14, '内服', 2, 200, 3, 0, 0, 0, 0, 0, '', '22', '', '', '', '', '2025-02-26  0:00:00', '', '', '', '', '', '3', '', '3', 0, '', '', '', 0, '', '', 0, '接方');
INSERT INTO `prescription` (`ID`, `deletion_number`, `Is_decoction`, `barcode_scan`, `hospital_id`, `hospital_name`, `prescription_number`, `decoction_method`, `prescription_type`, `patient_name`, `patient_sex`, `patient_age`, `patient_phone`, `patient_address`, `department_name`, `inpatient_area`, `ward_name`, `sick_bed`, `diagnosis_result`, `dosage`, `administration_method`, `administration_count`, `package_count`, `decoction_scheme`, `one_time_dosage`, `two_time_dosage`, `soak_water_amount`, `soak_time`, `label_number`, `remarks`, `doctor_name`, `footnote`, `drug_pickup_time`, `drug_pickup_number`, `order_time`, `do_time`, `do_person`, `distribution_company`, `distribution_address`, `distribution_phone`, `distribution_type`, `administration_way`, `additional_remarks_a`, `additional_remarks_b`, `drug_confirmation`, `logistics_state`, `query_time`, `query_person`, `drug_count`, `machine_room`, `state_post_back`, `subside_time`, `current_state`) VALUES (77, '', 1, '', '10001', '河南省人民医院', '9899931231', '', 0, '张丽莉', 2, 28, '', '', '', '', '', '', '9899931231', 14, '内服', 2, 200, 3, 0, 0, 0, 0, 0, '', 'xx', '', '', '', '', '2025-02-28  0:00:00', '', '', '', '', '', '3', '', '3', 0, '', '', '', 0, '', '', 0, '接方');
INSERT INTO `prescription` (`ID`, `deletion_number`, `Is_decoction`, `barcode_scan`, `hospital_id`, `hospital_name`, `prescription_number`, `decoction_method`, `prescription_type`, `patient_name`, `patient_sex`, `patient_age`, `patient_phone`, `patient_address`, `department_name`, `inpatient_area`, `ward_name`, `sick_bed`, `diagnosis_result`, `dosage`, `administration_method`, `administration_count`, `package_count`, `decoction_scheme`, `one_time_dosage`, `two_time_dosage`, `soak_water_amount`, `soak_time`, `label_number`, `remarks`, `doctor_name`, `footnote`, `drug_pickup_time`, `drug_pickup_number`, `order_time`, `do_time`, `do_person`, `distribution_company`, `distribution_address`, `distribution_phone`, `distribution_type`, `administration_way`, `additional_remarks_a`, `additional_remarks_b`, `drug_confirmation`, `logistics_state`, `query_time`, `query_person`, `drug_count`, `machine_room`, `state_post_back`, `subside_time`, `current_state`) VALUES (78, '', 0, '', '10001', '河南省人民医院', '32139999321', '', 1, '赵云里', 1, 22, '', '', '', '', '', '', '32139999321', 16, '内服', 2, 200, 3, 0, 0, 0, 0, 0, '', '312', '', '', '', '', '2025-02-28  0:00:00', '', '', '', '', '', '3', '', '3', 0, '', '', '', 0, '', '', 0, '审核');
INSERT INTO `prescription` (`ID`, `deletion_number`, `Is_decoction`, `barcode_scan`, `hospital_id`, `hospital_name`, `prescription_number`, `decoction_method`, `prescription_type`, `patient_name`, `patient_sex`, `patient_age`, `patient_phone`, `patient_address`, `department_name`, `inpatient_area`, `ward_name`, `sick_bed`, `diagnosis_result`, `dosage`, `administration_method`, `administration_count`, `package_count`, `decoction_scheme`, `one_time_dosage`, `two_time_dosage`, `soak_water_amount`, `soak_time`, `label_number`, `remarks`, `doctor_name`, `footnote`, `drug_pickup_time`, `drug_pickup_number`, `order_time`, `do_time`, `do_person`, `distribution_company`, `distribution_address`, `distribution_phone`, `distribution_type`, `administration_way`, `additional_remarks_a`, `additional_remarks_b`, `drug_confirmation`, `logistics_state`, `query_time`, `query_person`, `drug_count`, `machine_room`, `state_post_back`, `subside_time`, `current_state`) VALUES (79, '', 1, '', '10002', '河南省中医院', '20250312002', '', 0, '张悦', 2, 22, '1556223232', '', '', '一病区', '090002', '测试病床号', '20250312001', 7, '内服', 2, 180, 3, 0, 0, 400, 40, 0, '', '王梅', '', '2025-3-13 11:20:00', '', '', '2025-03-12  0:00:00', '', '', '', '', '', '3', '', '3', 0, '', '', '', 0, '', '', 0, '发货');
INSERT INTO `prescription` (`ID`, `deletion_number`, `Is_decoction`, `barcode_scan`, `hospital_id`, `hospital_name`, `prescription_number`, `decoction_method`, `prescription_type`, `patient_name`, `patient_sex`, `patient_age`, `patient_phone`, `patient_address`, `department_name`, `inpatient_area`, `ward_name`, `sick_bed`, `diagnosis_result`, `dosage`, `administration_method`, `administration_count`, `package_count`, `decoction_scheme`, `one_time_dosage`, `two_time_dosage`, `soak_water_amount`, `soak_time`, `label_number`, `remarks`, `doctor_name`, `footnote`, `drug_pickup_time`, `drug_pickup_number`, `order_time`, `do_time`, `do_person`, `distribution_company`, `distribution_address`, `distribution_phone`, `distribution_type`, `administration_way`, `additional_remarks_a`, `additional_remarks_b`, `drug_confirmation`, `logistics_state`, `query_time`, `query_person`, `drug_count`, `machine_room`, `state_post_back`, `subside_time`, `current_state`) VALUES (80, '', 2, '', '10001', '河南省人民医院', '9023213', '', 0, 'xx', 1, 22, '', 'xxxxx', '', '', '', '', '9023213', 2, '内服', 2, 200, 3, 0, 0, 0, 0, 0, '', 'xdddd', '', '', '', '', '2025-10-21', '', '', '', '', '', '3', '', '3', 0, '', '', '', 0, '', '', 0, '审核');
INSERT INTO `prescription` (`ID`, `deletion_number`, `Is_decoction`, `barcode_scan`, `hospital_id`, `hospital_name`, `prescription_number`, `decoction_method`, `prescription_type`, `patient_name`, `patient_sex`, `patient_age`, `patient_phone`, `patient_address`, `department_name`, `inpatient_area`, `ward_name`, `sick_bed`, `diagnosis_result`, `dosage`, `administration_method`, `administration_count`, `package_count`, `decoction_scheme`, `one_time_dosage`, `two_time_dosage`, `soak_water_amount`, `soak_time`, `label_number`, `remarks`, `doctor_name`, `footnote`, `drug_pickup_time`, `drug_pickup_number`, `order_time`, `do_time`, `do_person`, `distribution_company`, `distribution_address`, `distribution_phone`, `distribution_type`, `administration_way`, `additional_remarks_a`, `additional_remarks_b`, `drug_confirmation`, `logistics_state`, `query_time`, `query_person`, `drug_count`, `machine_room`, `state_post_back`, `subside_time`, `current_state`) VALUES (81, '1', 1, '', '10001', '河南省人民医院', '20251029001', '', 0, 'xx', 1, 22, '', '', '1', '', '', '', '20251029001', 7, '内服', 2, 200, 3, 0, 0, 0, 0, 0, '', '222', '', '', '', '', '2025-10-29', '', '', '', '', '', '3', '', '3', 0, '', '', '', 0, '', '', 0, '审核');
INSERT INTO `prescription` (`ID`, `deletion_number`, `Is_decoction`, `barcode_scan`, `hospital_id`, `hospital_name`, `prescription_number`, `decoction_method`, `prescription_type`, `patient_name`, `patient_sex`, `patient_age`, `patient_phone`, `patient_address`, `department_name`, `inpatient_area`, `ward_name`, `sick_bed`, `diagnosis_result`, `dosage`, `administration_method`, `administration_count`, `package_count`, `decoction_scheme`, `one_time_dosage`, `two_time_dosage`, `soak_water_amount`, `soak_time`, `label_number`, `remarks`, `doctor_name`, `footnote`, `drug_pickup_time`, `drug_pickup_number`, `order_time`, `do_time`, `do_person`, `distribution_company`, `distribution_address`, `distribution_phone`, `distribution_type`, `administration_way`, `additional_remarks_a`, `additional_remarks_b`, `drug_confirmation`, `logistics_state`, `query_time`, `query_person`, `drug_count`, `machine_room`, `state_post_back`, `subside_time`, `current_state`) VALUES (82, '', 1, '', '10001', '河南省人民医院', '200800', '', 0, 'xxxxxx', 1, 22, '', '', '', '', '', '', '200800', 2, '内服', 2, 200, 3, 0, 0, 0, 0, 0, '', 'yyy', '', '', '', '', '2025-10-25', '', '', '', '', '', '3', '', '3', 0, '', '', '', 0, '', '', 0, '审核');
INSERT INTO `prescription` (`ID`, `deletion_number`, `Is_decoction`, `barcode_scan`, `hospital_id`, `hospital_name`, `prescription_number`, `decoction_method`, `prescription_type`, `patient_name`, `patient_sex`, `patient_age`, `patient_phone`, `patient_address`, `department_name`, `inpatient_area`, `ward_name`, `sick_bed`, `diagnosis_result`, `dosage`, `administration_method`, `administration_count`, `package_count`, `decoction_scheme`, `one_time_dosage`, `two_time_dosage`, `soak_water_amount`, `soak_time`, `label_number`, `remarks`, `doctor_name`, `footnote`, `drug_pickup_time`, `drug_pickup_number`, `order_time`, `do_time`, `do_person`, `distribution_company`, `distribution_address`, `distribution_phone`, `distribution_type`, `administration_way`, `additional_remarks_a`, `additional_remarks_b`, `drug_confirmation`, `logistics_state`, `query_time`, `query_person`, `drug_count`, `machine_room`, `state_post_back`, `subside_time`, `current_state`) VALUES (83, '', 1, '', '10001', '河南省人民医院', '200800', '', 0, '测试数据', 1, 22, '', '', '', '', '', '', '200800', 2, '内服', 2, 200, 3, 0, 0, 0, 0, 0, '', 'yyy', '', '', '', '', '2025-10-25', '', '', '', '', '', '3', '', '3', 0, '', '', '', 0, '', '', 0, '接方');
INSERT INTO `prescription` (`ID`, `deletion_number`, `Is_decoction`, `barcode_scan`, `hospital_id`, `hospital_name`, `prescription_number`, `decoction_method`, `prescription_type`, `patient_name`, `patient_sex`, `patient_age`, `patient_phone`, `patient_address`, `department_name`, `inpatient_area`, `ward_name`, `sick_bed`, `diagnosis_result`, `dosage`, `administration_method`, `administration_count`, `package_count`, `decoction_scheme`, `one_time_dosage`, `two_time_dosage`, `soak_water_amount`, `soak_time`, `label_number`, `remarks`, `doctor_name`, `footnote`, `drug_pickup_time`, `drug_pickup_number`, `order_time`, `do_time`, `do_person`, `distribution_company`, `distribution_address`, `distribution_phone`, `distribution_type`, `administration_way`, `additional_remarks_a`, `additional_remarks_b`, `drug_confirmation`, `logistics_state`, `query_time`, `query_person`, `drug_count`, `machine_room`, `state_post_back`, `subside_time`, `current_state`) VALUES (84, '', 2, '', '10001', '河南省人民医院', '21313', '', 0, '谢谢', 1, 2, '', '', '', '', '', '', '21313', 2, '内服', 1, 200, 3, 0, 0, 0, 0, 0, '', '2', '', '', '', '', '', '', '', '', '', '', '3', '', '3', 0, '', '', '', 0, '', '', 0, '接方');
INSERT INTO `prescription` (`ID`, `deletion_number`, `Is_decoction`, `barcode_scan`, `hospital_id`, `hospital_name`, `prescription_number`, `decoction_method`, `prescription_type`, `patient_name`, `patient_sex`, `patient_age`, `patient_phone`, `patient_address`, `department_name`, `inpatient_area`, `ward_name`, `sick_bed`, `diagnosis_result`, `dosage`, `administration_method`, `administration_count`, `package_count`, `decoction_scheme`, `one_time_dosage`, `two_time_dosage`, `soak_water_amount`, `soak_time`, `label_number`, `remarks`, `doctor_name`, `footnote`, `drug_pickup_time`, `drug_pickup_number`, `order_time`, `do_time`, `do_person`, `distribution_company`, `distribution_address`, `distribution_phone`, `distribution_type`, `administration_way`, `additional_remarks_a`, `additional_remarks_b`, `drug_confirmation`, `logistics_state`, `query_time`, `query_person`, `drug_count`, `machine_room`, `state_post_back`, `subside_time`, `current_state`) VALUES (85, '', 2, '', '10001', '河南省人民医院', '231', '', 0, '23', 1, 3213, '', '', '', '', '', '', '231', 2, '内服', 232, 200, 3, 0, 0, 0, 0, 0, '', '323', '', '', '', '', '', '', '', '', '', '', '3', '', '3', 0, '', '', '', 0, '', '', 0, '接方');
COMMIT;

-- ----------------------------
-- Table structure for prescription_audit
-- ----------------------------
DROP TABLE IF EXISTS `prescription_audit`;
CREATE TABLE `prescription_audit` (
  `ID` int NOT NULL AUTO_INCREMENT,
  `prescription_id` int NOT NULL,
  `reviewer` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
  `review_time` varchar(255) DEFAULT NULL,
  `audit_status` int NOT NULL,
  `rejection_reason` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
  `print_status` int NOT NULL DEFAULT '0',
  `employee_id` int DEFAULT NULL,
  `create_time` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`ID`)
) ENGINE=InnoDB AUTO_INCREMENT=76 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='处方审核表';

-- ----------------------------
-- Records of prescription_audit
-- ----------------------------
BEGIN;
INSERT INTO `prescription_audit` (`ID`, `prescription_id`, `reviewer`, `review_time`, `audit_status`, `rejection_reason`, `print_status`, `employee_id`, `create_time`) VALUES (59, 68, 'admin', '2025-02-14 14:45:48', 1, '', 0, 1, '2025-02-14 14:42:51');
INSERT INTO `prescription_audit` (`ID`, `prescription_id`, `reviewer`, `review_time`, `audit_status`, `rejection_reason`, `print_status`, `employee_id`, `create_time`) VALUES (60, 69, 'admin', '2025-03-10 11:41:02', 1, '', 0, 1, '2025-02-14 14:45:33');
INSERT INTO `prescription_audit` (`ID`, `prescription_id`, `reviewer`, `review_time`, `audit_status`, `rejection_reason`, `print_status`, `employee_id`, `create_time`) VALUES (61, 70, '', '', 0, '', 0, 0, '2025-02-17 22:03:54');
INSERT INTO `prescription_audit` (`ID`, `prescription_id`, `reviewer`, `review_time`, `audit_status`, `rejection_reason`, `print_status`, `employee_id`, `create_time`) VALUES (62, 71, '', '', 0, '', 0, 0, '2025-02-24 13:42:50');
INSERT INTO `prescription_audit` (`ID`, `prescription_id`, `reviewer`, `review_time`, `audit_status`, `rejection_reason`, `print_status`, `employee_id`, `create_time`) VALUES (63, 72, '', '', 0, '', 0, 0, '2025-02-28 17:10:11');
INSERT INTO `prescription_audit` (`ID`, `prescription_id`, `reviewer`, `review_time`, `audit_status`, `rejection_reason`, `print_status`, `employee_id`, `create_time`) VALUES (64, 73, '', '', 0, '', 0, 0, '2025-02-28 17:11:19');
INSERT INTO `prescription_audit` (`ID`, `prescription_id`, `reviewer`, `review_time`, `audit_status`, `rejection_reason`, `print_status`, `employee_id`, `create_time`) VALUES (65, 74, '', '', 0, '', 0, 0, '2025-02-28 17:12:20');
INSERT INTO `prescription_audit` (`ID`, `prescription_id`, `reviewer`, `review_time`, `audit_status`, `rejection_reason`, `print_status`, `employee_id`, `create_time`) VALUES (66, 76, '', '', 0, '', 0, 0, '2025-02-28 17:13:31');
INSERT INTO `prescription_audit` (`ID`, `prescription_id`, `reviewer`, `review_time`, `audit_status`, `rejection_reason`, `print_status`, `employee_id`, `create_time`) VALUES (67, 77, '', '', 0, '', 0, 0, '2025-02-28 17:14:47');
INSERT INTO `prescription_audit` (`ID`, `prescription_id`, `reviewer`, `review_time`, `audit_status`, `rejection_reason`, `print_status`, `employee_id`, `create_time`) VALUES (68, 78, 'admin', '2025-10-29 13:02:30', 1, '', 0, 1, '2025-02-28 17:15:49');
INSERT INTO `prescription_audit` (`ID`, `prescription_id`, `reviewer`, `review_time`, `audit_status`, `rejection_reason`, `print_status`, `employee_id`, `create_time`) VALUES (69, 79, 'admin', '2025-03-12 22:55:56', 1, '', 0, 1, '2025-03-12 22:55:16');
INSERT INTO `prescription_audit` (`ID`, `prescription_id`, `reviewer`, `review_time`, `audit_status`, `rejection_reason`, `print_status`, `employee_id`, `create_time`) VALUES (70, 80, 'admin', '2025-10-29 13:02:49', 1, '', 0, 1, '2025-10-25 16:28:09');
INSERT INTO `prescription_audit` (`ID`, `prescription_id`, `reviewer`, `review_time`, `audit_status`, `rejection_reason`, `print_status`, `employee_id`, `create_time`) VALUES (71, 81, 'admin', '2025-10-29 12:20:16', 2, '', 0, 1, '2025-10-29 12:08:54');
INSERT INTO `prescription_audit` (`ID`, `prescription_id`, `reviewer`, `review_time`, `audit_status`, `rejection_reason`, `print_status`, `employee_id`, `create_time`) VALUES (72, 82, 'admin', '2025-10-29 15:04:24', 1, '', 0, 1, '2025-10-29 15:03:41');
INSERT INTO `prescription_audit` (`ID`, `prescription_id`, `reviewer`, `review_time`, `audit_status`, `rejection_reason`, `print_status`, `employee_id`, `create_time`) VALUES (73, 83, '', '', 0, '', 0, 0, '2025-10-29 15:04:48');
INSERT INTO `prescription_audit` (`ID`, `prescription_id`, `reviewer`, `review_time`, `audit_status`, `rejection_reason`, `print_status`, `employee_id`, `create_time`) VALUES (74, 84, '', '', 0, '', 0, 0, '2025-10-29 20:20:58');
INSERT INTO `prescription_audit` (`ID`, `prescription_id`, `reviewer`, `review_time`, `audit_status`, `rejection_reason`, `print_status`, `employee_id`, `create_time`) VALUES (75, 85, '', '', 0, '', 0, 0, '2025-10-29 21:11:18');
COMMIT;

-- ----------------------------
-- Table structure for prescription_drug
-- ----------------------------
DROP TABLE IF EXISTS `prescription_drug`;
CREATE TABLE `prescription_drug` (
  `id` int NOT NULL AUTO_INCREMENT,
  `prescription_id` varchar(15) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
  `hospital_id` varchar(15) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `stock_in_item_id` int DEFAULT NULL COMMENT '关联入库明细ID',
  `purchase_origin` varchar(100) DEFAULT NULL COMMENT '进货源地',
  `batch_no` varchar(50) DEFAULT NULL COMMENT '入库批次号',
  `drug_product_number` varchar(60) DEFAULT NULL,
  `drug_product_name` varchar(60) DEFAULT NULL,
  `measurement_unit` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
  `drug_product_description` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
  `dose` int DEFAULT NULL,
  `drug_weight` double DEFAULT NULL,
  `drug_weights` double DEFAULT NULL,
  `drug_price` double DEFAULT NULL,
  `total_prices` double DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=65 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='处方药品表';

-- ----------------------------
-- Records of prescription_drug
-- ----------------------------
BEGIN;
INSERT INTO `prescription_drug` (`id`, `prescription_id`, `hospital_id`, `drug_product_number`, `drug_product_name`, `measurement_unit`, `drug_product_description`, `dose`, `drug_weight`, `drug_weights`, `drug_price`, `total_prices`) VALUES (25, '68', '', '20880080', '甘草', '', '', 6, 2, 12, 2.9, 10);
INSERT INTO `prescription_drug` (`id`, `prescription_id`, `hospital_id`, `drug_product_number`, `drug_product_name`, `measurement_unit`, `drug_product_description`, `dose`, `drug_weight`, `drug_weights`, `drug_price`, `total_prices`) VALUES (26, '68', '', '20899092', '测试药品1', '', '', 7, 1, 7, 2, 12);
INSERT INTO `prescription_drug` (`id`, `prescription_id`, `hospital_id`, `drug_product_number`, `drug_product_name`, `measurement_unit`, `drug_product_description`, `dose`, `drug_weight`, `drug_weights`, `drug_price`, `total_prices`) VALUES (27, '68', '', '20989002', '测试药品2', '', '', 7, 2, 14, 3.6, 10);
INSERT INTO `prescription_drug` (`id`, `prescription_id`, `hospital_id`, `drug_product_number`, `drug_product_name`, `measurement_unit`, `drug_product_description`, `dose`, `drug_weight`, `drug_weights`, `drug_price`, `total_prices`) VALUES (28, '68', '', '30909003', '测试药品3', '', '', 7, 3, 21, 0, 10);
INSERT INTO `prescription_drug` (`id`, `prescription_id`, `hospital_id`, `drug_product_number`, `drug_product_name`, `measurement_unit`, `drug_product_description`, `dose`, `drug_weight`, `drug_weights`, `drug_price`, `total_prices`) VALUES (29, '68', '', '30900032', '测试药品4', '', '', 7, 4, 28, 4, 18);
INSERT INTO `prescription_drug` (`id`, `prescription_id`, `hospital_id`, `drug_product_number`, `drug_product_name`, `measurement_unit`, `drug_product_description`, `dose`, `drug_weight`, `drug_weights`, `drug_price`, `total_prices`) VALUES (30, '68', '', '209009903', '测试药品5', '', '', 7, 1, 7, 4, 14.9);
INSERT INTO `prescription_drug` (`id`, `prescription_id`, `hospital_id`, `drug_product_number`, `drug_product_name`, `measurement_unit`, `drug_product_description`, `dose`, `drug_weight`, `drug_weights`, `drug_price`, `total_prices`) VALUES (31, '68', '', '9008203', '测试药品6', '', '', 7, 3, 21, 2.8, 10);
INSERT INTO `prescription_drug` (`id`, `prescription_id`, `hospital_id`, `drug_product_number`, `drug_product_name`, `measurement_unit`, `drug_product_description`, `dose`, `drug_weight`, `drug_weights`, `drug_price`, `total_prices`) VALUES (32, '69', '', '2090000', '测试药品', '', '', 7, 2, 15, 3, 10);
INSERT INTO `prescription_drug` (`id`, `prescription_id`, `hospital_id`, `drug_product_number`, `drug_product_name`, `measurement_unit`, `drug_product_description`, `dose`, `drug_weight`, `drug_weights`, `drug_price`, `total_prices`) VALUES (33, '69', '', '9008900', '测试药品02', '', '', 7, 2, 14, 3, 10);
INSERT INTO `prescription_drug` (`id`, `prescription_id`, `hospital_id`, `drug_product_number`, `drug_product_name`, `measurement_unit`, `drug_product_description`, `dose`, `drug_weight`, `drug_weights`, `drug_price`, `total_prices`) VALUES (34, '69', '', '09000', '测试药品03', '', '', 7, 2, 18, 3, 10);
INSERT INTO `prescription_drug` (`id`, `prescription_id`, `hospital_id`, `drug_product_number`, `drug_product_name`, `measurement_unit`, `drug_product_description`, `dose`, `drug_weight`, `drug_weights`, `drug_price`, `total_prices`) VALUES (35, '70', '', 'ZHCY0575', '熟地黄', '', '', 7, 0.29, 2.03, 3, 6);
INSERT INTO `prescription_drug` (`id`, `prescription_id`, `hospital_id`, `drug_product_number`, `drug_product_name`, `measurement_unit`, `drug_product_description`, `dose`, `drug_weight`, `drug_weights`, `drug_price`, `total_prices`) VALUES (36, '70', '', 'ZHCY0057', '白附片', '', '', 7, 9.6, 67.2, 4, 240);
INSERT INTO `prescription_drug` (`id`, `prescription_id`, `hospital_id`, `drug_product_number`, `drug_product_name`, `measurement_unit`, `drug_product_description`, `dose`, `drug_weight`, `drug_weights`, `drug_price`, `total_prices`) VALUES (37, '70', '', 'ZHCY1148M', '阿胶珠', '', '', 7, 15.53, 108.71, 5, 879);
INSERT INTO `prescription_drug` (`id`, `prescription_id`, `hospital_id`, `drug_product_number`, `drug_product_name`, `measurement_unit`, `drug_product_description`, `dose`, `drug_weight`, `drug_weights`, `drug_price`, `total_prices`) VALUES (38, '71', '10002', 'wow', 'wewqe', '', '', 2, 2, 4, 0, 0);
INSERT INTO `prescription_drug` (`id`, `prescription_id`, `hospital_id`, `drug_product_number`, `drug_product_name`, `measurement_unit`, `drug_product_description`, `dose`, `drug_weight`, `drug_weights`, `drug_price`, `total_prices`) VALUES (39, '72', '10001', '98900890', '测试药品', '', '', 7, 2, 14, 0, 0);
INSERT INTO `prescription_drug` (`id`, `prescription_id`, `hospital_id`, `drug_product_number`, `drug_product_name`, `measurement_unit`, `drug_product_description`, `dose`, `drug_weight`, `drug_weights`, `drug_price`, `total_prices`) VALUES (40, '73', '10001', '8789900', '测试药品2', '', '', 7, 2, 14, 0, 0);
INSERT INTO `prescription_drug` (`id`, `prescription_id`, `hospital_id`, `drug_product_number`, `drug_product_name`, `measurement_unit`, `drug_product_description`, `dose`, `drug_weight`, `drug_weights`, `drug_price`, `total_prices`) VALUES (41, '74', '10001', '232333434', '人参', '', '', 7, 2, 14, 0, 0);
INSERT INTO `prescription_drug` (`id`, `prescription_id`, `hospital_id`, `drug_product_number`, `drug_product_name`, `measurement_unit`, `drug_product_description`, `dose`, `drug_weight`, `drug_weights`, `drug_price`, `total_prices`) VALUES (42, '76', '10001', '2312213213', '党参', '', '', 14, 2, 28, 0, 0);
INSERT INTO `prescription_drug` (`id`, `prescription_id`, `hospital_id`, `drug_product_number`, `drug_product_name`, `measurement_unit`, `drug_product_description`, `dose`, `drug_weight`, `drug_weights`, `drug_price`, `total_prices`) VALUES (43, '77', '10001', '3223233', '人参', '', '', 14, 14, 196, 0, 0);
INSERT INTO `prescription_drug` (`id`, `prescription_id`, `hospital_id`, `drug_product_number`, `drug_product_name`, `measurement_unit`, `drug_product_description`, `dose`, `drug_weight`, `drug_weights`, `drug_price`, `total_prices`) VALUES (44, '77', '10001', '88990', '甘草', '', '', 14, 16, 224, 0, 0);
INSERT INTO `prescription_drug` (`id`, `prescription_id`, `hospital_id`, `drug_product_number`, `drug_product_name`, `measurement_unit`, `drug_product_description`, `dose`, `drug_weight`, `drug_weights`, `drug_price`, `total_prices`) VALUES (45, '77', '10001', '23213', '山药', '', '', 17, 3, 51, 0, 0);
INSERT INTO `prescription_drug` (`id`, `prescription_id`, `hospital_id`, `drug_product_number`, `drug_product_name`, `measurement_unit`, `drug_product_description`, `dose`, `drug_weight`, `drug_weights`, `drug_price`, `total_prices`) VALUES (46, '78', '10001', '32132134', '当归', '', '', 2, 23, 46, 0, 0);
INSERT INTO `prescription_drug` (`id`, `prescription_id`, `hospital_id`, `drug_product_number`, `drug_product_name`, `measurement_unit`, `drug_product_description`, `dose`, `drug_weight`, `drug_weights`, `drug_price`, `total_prices`) VALUES (47, '78', '10001', '312324', '红参', '', '', 2, 34, 68, 0, 0);
INSERT INTO `prescription_drug` (`id`, `prescription_id`, `hospital_id`, `drug_product_number`, `drug_product_name`, `measurement_unit`, `drug_product_description`, `dose`, `drug_weight`, `drug_weights`, `drug_price`, `total_prices`) VALUES (48, '79', '10002', '090001', '甘草', '', '', 7, 75, 525, 0, 0);
INSERT INTO `prescription_drug` (`id`, `prescription_id`, `hospital_id`, `drug_product_number`, `drug_product_name`, `measurement_unit`, `drug_product_description`, `dose`, `drug_weight`, `drug_weights`, `drug_price`, `total_prices`) VALUES (49, '79', '10002', '090080', '红参', '', '', 7, 5, 35, 0, 0);
INSERT INTO `prescription_drug` (`id`, `prescription_id`, `hospital_id`, `drug_product_number`, `drug_product_name`, `measurement_unit`, `drug_product_description`, `dose`, `drug_weight`, `drug_weights`, `drug_price`, `total_prices`) VALUES (50, '79', '10002', '090809', '党参', '', '', 7, 5, 35, 0, 0);
INSERT INTO `prescription_drug` (`id`, `prescription_id`, `hospital_id`, `drug_product_number`, `drug_product_name`, `measurement_unit`, `drug_product_description`, `dose`, `drug_weight`, `drug_weights`, `drug_price`, `total_prices`) VALUES (51, '79', '10002', '090090', '白芍', '', '', 7, 5, 35, 0, 0);
INSERT INTO `prescription_drug` (`id`, `prescription_id`, `hospital_id`, `drug_product_number`, `drug_product_name`, `measurement_unit`, `drug_product_description`, `dose`, `drug_weight`, `drug_weights`, `drug_price`, `total_prices`) VALUES (52, '80', '10001', '10903231', '甘草', '', '', 2, 2, 4, 0, 0);
INSERT INTO `prescription_drug` (`id`, `prescription_id`, `hospital_id`, `drug_product_number`, `drug_product_name`, `measurement_unit`, `drug_product_description`, `dose`, `drug_weight`, `drug_weights`, `drug_price`, `total_prices`) VALUES (53, '80', '10001', '0901122', '白', '', '', 3, 2, 6, 0, 0);
INSERT INTO `prescription_drug` (`id`, `prescription_id`, `hospital_id`, `drug_product_number`, `drug_product_name`, `measurement_unit`, `drug_product_description`, `dose`, `drug_weight`, `drug_weights`, `drug_price`, `total_prices`) VALUES (54, '81', '10001', 'DR001', '甘草', '', '', 7, 20, 140, 0, 0);
INSERT INTO `prescription_drug` (`id`, `prescription_id`, `hospital_id`, `drug_product_number`, `drug_product_name`, `measurement_unit`, `drug_product_description`, `dose`, `drug_weight`, `drug_weights`, `drug_price`, `total_prices`) VALUES (55, '81', '10001', '09000', '党参', '', '', 7, 10, 70, 0, 0);
INSERT INTO `prescription_drug` (`id`, `prescription_id`, `hospital_id`, `drug_product_number`, `drug_product_name`, `measurement_unit`, `drug_product_description`, `dose`, `drug_weight`, `drug_weights`, `drug_price`, `total_prices`) VALUES (56, '81', '10001', '09012', '测试1', '', '', 7, 20, 140, 0, 0);
INSERT INTO `prescription_drug` (`id`, `prescription_id`, `hospital_id`, `drug_product_number`, `drug_product_name`, `measurement_unit`, `drug_product_description`, `dose`, `drug_weight`, `drug_weights`, `drug_price`, `total_prices`) VALUES (57, '81', '10001', 'DR002', '测试2', '', '', 7, 2, 14, 0, 0);
INSERT INTO `prescription_drug` (`id`, `prescription_id`, `hospital_id`, `drug_product_number`, `drug_product_name`, `measurement_unit`, `drug_product_description`, `dose`, `drug_weight`, `drug_weights`, `drug_price`, `total_prices`) VALUES (58, '81', '10001', 'DR002', '测试3', '', '', 7, 10, 70, 0, 0);
INSERT INTO `prescription_drug` (`id`, `prescription_id`, `hospital_id`, `drug_product_number`, `drug_product_name`, `measurement_unit`, `drug_product_description`, `dose`, `drug_weight`, `drug_weights`, `drug_price`, `total_prices`) VALUES (59, '82', '10001', 'DR240', '甘草', '', '', 1, 20, 20, 0, 0);
INSERT INTO `prescription_drug` (`id`, `prescription_id`, `hospital_id`, `drug_product_number`, `drug_product_name`, `measurement_unit`, `drug_product_description`, `dose`, `drug_weight`, `drug_weights`, `drug_price`, `total_prices`) VALUES (60, '82', '10001', '0901021', '党参', '', '', 2, 2, 4, 0, 0);
INSERT INTO `prescription_drug` (`id`, `prescription_id`, `hospital_id`, `drug_product_number`, `drug_product_name`, `measurement_unit`, `drug_product_description`, `dose`, `drug_weight`, `drug_weights`, `drug_price`, `total_prices`) VALUES (61, '83', '10001', 'DR240', '甘草', '', '', 1, 28, 28, 0, 0);
INSERT INTO `prescription_drug` (`id`, `prescription_id`, `hospital_id`, `drug_product_number`, `drug_product_name`, `measurement_unit`, `drug_product_description`, `dose`, `drug_weight`, `drug_weights`, `drug_price`, `total_prices`) VALUES (62, '83', '10001', '0901021', '党参', '', '', 2, 2, 4, 0, 0);
INSERT INTO `prescription_drug` (`id`, `prescription_id`, `hospital_id`, `drug_product_number`, `drug_product_name`, `measurement_unit`, `drug_product_description`, `dose`, `drug_weight`, `drug_weights`, `drug_price`, `total_prices`) VALUES (63, '84', '10001', '100212', '甘草', '', '', 1, 1, 1, 0, 0);
INSERT INTO `prescription_drug` (`id`, `prescription_id`, `hospital_id`, `drug_product_number`, `drug_product_name`, `measurement_unit`, `drug_product_description`, `dose`, `drug_weight`, `drug_weights`, `drug_price`, `total_prices`) VALUES (64, '85', '10001', '2132', 'xxq', '', '', 2, 2, 4, 0, 0);
COMMIT;

-- ----------------------------
-- Table structure for proess_fee
-- ----------------------------
DROP TABLE IF EXISTS `proess_fee`;
CREATE TABLE `proess_fee` (
  `id` int NOT NULL AUTO_INCREMENT,
  `mehod` varchar(30) DEFAULT NULL,
  `hospital_id` int DEFAULT NULL,
  `hospital_name` varchar(30) DEFAULT NULL,
  `fee` double DEFAULT NULL,
  `contact_name` varchar(40) DEFAULT NULL,
  `tel` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='结算方信息';

-- ----------------------------
-- Records of proess_fee
-- ----------------------------
BEGIN;
INSERT INTO `proess_fee` (`id`, `mehod`, `hospital_id`, `hospital_name`, `fee`, `contact_name`, `tel`) VALUES (1, '1', 10001, '河南省人民医院', 3, '张三', '13526798791');
COMMIT;

-- ----------------------------
-- Table structure for sam_info
-- ----------------------------
DROP TABLE IF EXISTS `sam_info`;
CREATE TABLE `sam_info` (
  `id` int NOT NULL AUTO_INCREMENT,
  `prescription_number` varchar(60) DEFAULT NULL,
  `prescription_weight` varchar(255) DEFAULT NULL,
  `prescription_sj_weight` varchar(255) DEFAULT NULL,
  `drug_count` int DEFAULT NULL,
  `drug_sj_count` int DEFAULT NULL,
  `dosage` int DEFAULT NULL,
  `operate_name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
  `operate_time` varchar(255) DEFAULT NULL,
  `operate_update_time` varchar(255) DEFAULT NULL,
  `remark` varchar(80) DEFAULT NULL,
  `status` int DEFAULT '0',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=10 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='质量检测表';

-- ----------------------------
-- Records of sam_info
-- ----------------------------
BEGIN;
INSERT INTO `sam_info` (`id`, `prescription_number`, `prescription_weight`, `prescription_sj_weight`, `drug_count`, `drug_sj_count`, `dosage`, `operate_name`, `operate_time`, `operate_update_time`, `remark`, `status`) VALUES (1, '20250214002', '237', '236', 10, 10, 7, '王华', '2025-02-15 22:17:56', '2025-02-15 22:17:59', '测试1', 1);
INSERT INTO `sam_info` (`id`, `prescription_number`, `prescription_weight`, `prescription_sj_weight`, `drug_count`, `drug_sj_count`, `dosage`, `operate_name`, `operate_time`, `operate_update_time`, `remark`, `status`) VALUES (4, '20900232', '890032', '22', 22, 22, 0, 'admin', '2025-09-21', '2025-10-25 16:33:44', '', 1);
INSERT INTO `sam_info` (`id`, `prescription_number`, `prescription_weight`, `prescription_sj_weight`, `drug_count`, `drug_sj_count`, `dosage`, `operate_name`, `operate_time`, `operate_update_time`, `remark`, `status`) VALUES (5, '2090090', '23', '23', 0, 0, 0, 'admin', '2025-10-11', '', '', 0);
INSERT INTO `sam_info` (`id`, `prescription_number`, `prescription_weight`, `prescription_sj_weight`, `drug_count`, `drug_sj_count`, `dosage`, `operate_name`, `operate_time`, `operate_update_time`, `remark`, `status`) VALUES (8, 'BZ0901', '', '', 10, 10, 10, 'admin', '2025-10-28', '', 'Xx', 0);
INSERT INTO `sam_info` (`id`, `prescription_number`, `prescription_weight`, `prescription_sj_weight`, `drug_count`, `drug_sj_count`, `dosage`, `operate_name`, `operate_time`, `operate_update_time`, `remark`, `status`) VALUES (9, '04032000000000080', '', '', 10, 10, 7, 'admin', '2025-10-29', '', '测试', 0);
COMMIT;

-- ----------------------------
-- Table structure for stock_adjustment
-- ----------------------------
DROP TABLE IF EXISTS `stock_adjustment`;
CREATE TABLE `stock_adjustment` (
  `id` int NOT NULL AUTO_INCREMENT,
  `order_no` varchar(30) NOT NULL COMMENT '调整单号',
  `warehouse_id` int NOT NULL COMMENT '仓库ID',
  `adjustment_type` tinyint NOT NULL COMMENT '调整类型（1-报损 2-报溢）',
  `operator_id` int NOT NULL COMMENT '操作人ID',
  `reason` varchar(500) NOT NULL COMMENT '调整原因',
  `remark` varchar(500) DEFAULT NULL COMMENT '备注',
  `create_time` datetime DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_time` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE KEY `order_no` (`order_no`),
  KEY `idx_warehouse` (`warehouse_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='报损报溢表';

-- ----------------------------
-- Records of stock_adjustment
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for stock_adjustment_items
-- ----------------------------
DROP TABLE IF EXISTS `stock_adjustment_items`;
CREATE TABLE `stock_adjustment_items` (
  `id` int NOT NULL AUTO_INCREMENT,
  `adjustment_id` int NOT NULL COMMENT '调整单ID',
  `product_id` int NOT NULL COMMENT '商品ID',
  `quantity` decimal(12,2) NOT NULL COMMENT '数量',
  `unit_price` decimal(12,2) NOT NULL COMMENT '单价',
  `amount` decimal(12,2) NOT NULL COMMENT '金额',
  `batch_no` varchar(30) DEFAULT NULL COMMENT '批次号',
  `remark` varchar(200) DEFAULT NULL COMMENT '备注',
  PRIMARY KEY (`id`) USING BTREE,
  KEY `idx_adjustment` (`adjustment_id`),
  KEY `idx_product` (`product_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='报损报溢明细表';

-- ----------------------------
-- Records of stock_adjustment_items
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for stock_check
-- ----------------------------
DROP TABLE IF EXISTS `stock_check`;
CREATE TABLE `stock_check` (
  `id` int NOT NULL AUTO_INCREMENT,
  `order_no` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '盘点单号',
  `warehouse_id` int NOT NULL COMMENT '仓库ID',
  `operator_id` int NOT NULL COMMENT '操作人ID',
  `remark` varchar(500) DEFAULT NULL COMMENT '备注',
  `create_time` datetime DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_time` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE KEY `order_no` (`order_no`),
  KEY `idx_warehouse` (`warehouse_id`)
) ENGINE=InnoDB AUTO_INCREMENT=11 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='盘点表';

-- ----------------------------
-- Records of stock_check
-- ----------------------------
BEGIN;
INSERT INTO `stock_check` (`id`, `order_no`, `warehouse_id`, `operator_id`, `remark`, `create_time`, `update_time`) VALUES (1, 'P202510300001', 1, 101, '季度例行盘点', '2025-10-30 07:03:56', '2025-10-30 07:03:56');
INSERT INTO `stock_check` (`id`, `order_no`, `warehouse_id`, `operator_id`, `remark`, `create_time`, `update_time`) VALUES (2, 'P202510300002', 2, 102, '临时抽查盘点', '2025-10-30 07:03:56', '2025-10-30 07:03:56');
INSERT INTO `stock_check` (`id`, `order_no`, `warehouse_id`, `operator_id`, `remark`, `create_time`, `update_time`) VALUES (3, 'P202510300003', 3, 103, '年度大盘点', '2025-10-30 07:03:56', '2025-10-30 07:03:56');
INSERT INTO `stock_check` (`id`, `order_no`, `warehouse_id`, `operator_id`, `remark`, `create_time`, `update_time`) VALUES (9, 'PC1761830222636', 1, 2, '测试', '2025-10-30 21:25:31', '2025-10-30 21:25:31');
INSERT INTO `stock_check` (`id`, `order_no`, `warehouse_id`, `operator_id`, `remark`, `create_time`, `update_time`) VALUES (10, 'PC1761923343908', 1, 1, '1', '2025-10-31 23:10:10', '2025-10-31 23:10:10');
COMMIT;

-- ----------------------------
-- Table structure for stock_check_items
-- ----------------------------
DROP TABLE IF EXISTS `stock_check_items`;
CREATE TABLE `stock_check_items` (
  `id` int NOT NULL AUTO_INCREMENT,
  `check_id` int NOT NULL COMMENT '盘点单ID',
  `product_id` varchar(30) NOT NULL COMMENT '商品ID',
  `system_quantity` decimal(12,2) NOT NULL COMMENT '系统数量',
  `actual_quantity` decimal(12,2) NOT NULL COMMENT '实际数量',
  `difference` decimal(12,2) NOT NULL COMMENT '差异数量',
  `batch_no` varchar(30) DEFAULT NULL COMMENT '批次号',
  `remark` varchar(200) DEFAULT NULL COMMENT '备注',
  PRIMARY KEY (`id`) USING BTREE,
  KEY `idx_check` (`check_id`),
  KEY `idx_product` (`product_id`)
) ENGINE=InnoDB AUTO_INCREMENT=13 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='盘点明细表';

-- ----------------------------
-- Records of stock_check_items
-- ----------------------------
BEGIN;
INSERT INTO `stock_check_items` (`id`, `check_id`, `product_id`, `system_quantity`, `actual_quantity`, `difference`, `batch_no`, `remark`) VALUES (1, 1, 'P001', 100.00, 98.00, -2.00, 'B20240101', '数量差异');
INSERT INTO `stock_check_items` (`id`, `check_id`, `product_id`, `system_quantity`, `actual_quantity`, `difference`, `batch_no`, `remark`) VALUES (2, 1, 'P002', 200.00, 200.00, 0.00, 'B20240201', '无差异');
INSERT INTO `stock_check_items` (`id`, `check_id`, `product_id`, `system_quantity`, `actual_quantity`, `difference`, `batch_no`, `remark`) VALUES (3, 1, 'P003', 150.00, 155.00, 5.00, 'B20240301', '数量差异');
INSERT INTO `stock_check_items` (`id`, `check_id`, `product_id`, `system_quantity`, `actual_quantity`, `difference`, `batch_no`, `remark`) VALUES (4, 2, 'P004', 300.00, 298.00, -2.00, 'B20240401', '数量差异');
INSERT INTO `stock_check_items` (`id`, `check_id`, `product_id`, `system_quantity`, `actual_quantity`, `difference`, `batch_no`, `remark`) VALUES (5, 2, 'P005', 400.00, 400.00, 0.00, 'B20240501', '无差异');
INSERT INTO `stock_check_items` (`id`, `check_id`, `product_id`, `system_quantity`, `actual_quantity`, `difference`, `batch_no`, `remark`) VALUES (6, 3, 'P006', 500.00, 505.00, 5.00, 'B20240601', '数量差异');
INSERT INTO `stock_check_items` (`id`, `check_id`, `product_id`, `system_quantity`, `actual_quantity`, `difference`, `batch_no`, `remark`) VALUES (7, 3, 'P007', 600.00, 600.00, 0.00, 'B20240701', '无差异');
INSERT INTO `stock_check_items` (`id`, `check_id`, `product_id`, `system_quantity`, `actual_quantity`, `difference`, `batch_no`, `remark`) VALUES (8, 3, 'P008', 700.00, 695.00, -5.00, 'B20240801', '数量差异');
INSERT INTO `stock_check_items` (`id`, `check_id`, `product_id`, `system_quantity`, `actual_quantity`, `difference`, `batch_no`, `remark`) VALUES (11, 9, '90023', 88.00, 90.00, 2.00, '2025001', '');
INSERT INTO `stock_check_items` (`id`, `check_id`, `product_id`, `system_quantity`, `actual_quantity`, `difference`, `batch_no`, `remark`) VALUES (12, 10, '90023', 90.00, 100.00, 10.00, '90231', '');
COMMIT;

-- ----------------------------
-- Table structure for stock_in
-- ----------------------------
DROP TABLE IF EXISTS `stock_in`;
CREATE TABLE `stock_in` (
  `id` int NOT NULL AUTO_INCREMENT,
  `order_no` varchar(30) NOT NULL COMMENT '入库单号',
  `warehouse_id` int NOT NULL,
  `operator_id` int NOT NULL COMMENT '操作人ID',
  `create_time` datetime DEFAULT CURRENT_TIMESTAMP,
  `update_time` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `remark` varchar(50) DEFAULT NULL,
  `status` int DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE KEY `order_no` (`order_no`),
  KEY `idx_warehouse` (`warehouse_id`)
) ENGINE=InnoDB AUTO_INCREMENT=45 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='入库单表';

-- ----------------------------
-- Records of stock_in
-- ----------------------------
BEGIN;
INSERT INTO `stock_in` (`id`, `order_no`, `warehouse_id`, `operator_id`, `create_time`, `update_time`, `remark`, `status`) VALUES (1, 'RK202510001', 1, 101, '2025-10-01 09:30:00', '2025-10-01 09:30:00', NULL, NULL);
INSERT INTO `stock_in` (`id`, `order_no`, `warehouse_id`, `operator_id`, `create_time`, `update_time`, `remark`, `status`) VALUES (2, 'RK202510002', 1, 101, '2025-10-03 14:15:00', '2025-10-03 14:15:00', NULL, NULL);
INSERT INTO `stock_in` (`id`, `order_no`, `warehouse_id`, `operator_id`, `create_time`, `update_time`, `remark`, `status`) VALUES (3, 'RK202510003', 2, 102, '2025-10-05 10:00:00', '2025-10-05 10:00:00', NULL, NULL);
INSERT INTO `stock_in` (`id`, `order_no`, `warehouse_id`, `operator_id`, `create_time`, `update_time`, `remark`, `status`) VALUES (4, 'RK202510004', 1, 103, '2025-10-07 16:40:00', '2025-10-07 16:40:00', NULL, NULL);
INSERT INTO `stock_in` (`id`, `order_no`, `warehouse_id`, `operator_id`, `create_time`, `update_time`, `remark`, `status`) VALUES (5, 'RK202510005', 2, 102, '2025-10-09 11:20:00', '2025-10-09 11:20:00', NULL, NULL);
INSERT INTO `stock_in` (`id`, `order_no`, `warehouse_id`, `operator_id`, `create_time`, `update_time`, `remark`, `status`) VALUES (6, 'RK202510006', 1, 101, '2025-10-12 08:50:00', '2025-10-12 08:50:00', NULL, NULL);
INSERT INTO `stock_in` (`id`, `order_no`, `warehouse_id`, `operator_id`, `create_time`, `update_time`, `remark`, `status`) VALUES (7, 'RK202510007', 3, 104, '2025-10-15 15:30:00', '2025-10-15 15:30:00', NULL, NULL);
INSERT INTO `stock_in` (`id`, `order_no`, `warehouse_id`, `operator_id`, `create_time`, `update_time`, `remark`, `status`) VALUES (8, 'RK202510008', 1, 103, '2025-10-18 13:10:00', '2025-10-18 13:10:00', NULL, NULL);
INSERT INTO `stock_in` (`id`, `order_no`, `warehouse_id`, `operator_id`, `create_time`, `update_time`, `remark`, `status`) VALUES (9, 'RK202510009', 2, 102, '2025-10-21 09:45:00', '2025-10-21 09:45:00', NULL, NULL);
INSERT INTO `stock_in` (`id`, `order_no`, `warehouse_id`, `operator_id`, `create_time`, `update_time`, `remark`, `status`) VALUES (10, 'RK202510010', 3, 104, '2025-10-24 16:20:00', '2025-10-24 16:20:00', NULL, NULL);
INSERT INTO `stock_in` (`id`, `order_no`, `warehouse_id`, `operator_id`, `create_time`, `update_time`, `remark`, `status`) VALUES (37, '209031', 1, 101, '2025-10-29 23:10:46', '2025-10-29 23:10:46', '', 0);
INSERT INTO `stock_in` (`id`, `order_no`, `warehouse_id`, `operator_id`, `create_time`, `update_time`, `remark`, `status`) VALUES (40, '209931', 1, 102, '2025-10-29 23:18:10', '2025-10-29 23:18:10', '', 0);
INSERT INTO `stock_in` (`id`, `order_no`, `warehouse_id`, `operator_id`, `create_time`, `update_time`, `remark`, `status`) VALUES (41, 'RK1761917664796', 1, 102, '2025-10-31 21:37:31', '2025-10-31 21:37:31', '', 0);
INSERT INTO `stock_in` (`id`, `order_no`, `warehouse_id`, `operator_id`, `create_time`, `update_time`, `remark`, `status`) VALUES (44, 'RK1761921214909', 1, 101, '2025-10-31 22:41:46', '2025-10-31 22:41:46', '', 0);
COMMIT;

-- ----------------------------
-- Table structure for stock_in_items
-- ----------------------------
DROP TABLE IF EXISTS `stock_in_items`;
CREATE TABLE `stock_in_items` (
  `id` int NOT NULL AUTO_INCREMENT,
  `stock_in_id` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `product_id` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `product_name` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
  `quantity` decimal(12,2) NOT NULL,
  `unit_price` decimal(12,2) NOT NULL,
  `amount` decimal(12,2) NOT NULL,
  `batch_no` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL COMMENT '批次号',
  `purchase_origin` varchar(100) DEFAULT NULL COMMENT '进货源地',
  `supplier_name` varchar(100) DEFAULT NULL COMMENT '供应商名称',
  `production_date` datetime DEFAULT NULL COMMENT '生产日期',
  `expiry_date` datetime DEFAULT NULL COMMENT '过期日期',
  `remark` varchar(200) DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  KEY `idx_stock_in` (`stock_in_id`),
  KEY `idx_product` (`product_id`)
) ENGINE=InnoDB AUTO_INCREMENT=62 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='入库明细表';

-- ----------------------------
-- Records of stock_in_items
-- ----------------------------
BEGIN;
INSERT INTO `stock_in_items` (`id`, `stock_in_id`, `product_id`, `product_name`, `quantity`, `unit_price`, `amount`, `batch_no`, `production_date`, `expiry_date`, `remark`) VALUES (1, '1', '1001', '法半夏', 50.00, 12.50, 625.00, '20251001001', '2025-09-15 00:00:00', '2028-09-14 00:00:00', '当归（甘肃产）');
INSERT INTO `stock_in_items` (`id`, `stock_in_id`, `product_id`, `product_name`, `quantity`, `unit_price`, `amount`, `batch_no`, `production_date`, `expiry_date`, `remark`) VALUES (2, '1', '1002', '槟郎', 30.00, 9.80, 294.00, '20251001002', '2025-09-20 00:00:00', '2028-09-19 00:00:00', '黄芪（内蒙古产）');
INSERT INTO `stock_in_items` (`id`, `stock_in_id`, `product_id`, `product_name`, `quantity`, `unit_price`, `amount`, `batch_no`, `production_date`, `expiry_date`, `remark`) VALUES (3, '1', '1003', '通草', 20.00, 15.60, 312.00, '20251001003', '2025-09-10 00:00:00', '2027-09-09 00:00:00', '党参');
INSERT INTO `stock_in_items` (`id`, `stock_in_id`, `product_id`, `product_name`, `quantity`, `unit_price`, `amount`, `batch_no`, `production_date`, `expiry_date`, `remark`) VALUES (4, '2', '1001', NULL, 20.00, 12.30, 246.00, '20251003001', '2025-09-25 00:00:00', '2028-09-24 00:00:00', '当归（新批次）');
INSERT INTO `stock_in_items` (`id`, `stock_in_id`, `product_id`, `product_name`, `quantity`, `unit_price`, `amount`, `batch_no`, `production_date`, `expiry_date`, `remark`) VALUES (5, '2', '1002', NULL, 40.00, 9.70, 388.00, '20251003002', '2025-09-30 00:00:00', '2028-09-29 00:00:00', '黄芪（特级）');
INSERT INTO `stock_in_items` (`id`, `stock_in_id`, `product_id`, `product_name`, `quantity`, `unit_price`, `amount`, `batch_no`, `production_date`, `expiry_date`, `remark`) VALUES (6, '3', '1004', NULL, 15.00, 28.50, 427.50, '20251005001', '2025-09-05 00:00:00', '2027-09-04 00:00:00', '枸杞（宁夏）');
INSERT INTO `stock_in_items` (`id`, `stock_in_id`, `product_id`, `product_name`, `quantity`, `unit_price`, `amount`, `batch_no`, `production_date`, `expiry_date`, `remark`) VALUES (7, '3', '1005', NULL, 10.00, 45.00, 450.00, '20251005002', '2025-08-20 00:00:00', '2027-08-19 00:00:00', '金银花');
INSERT INTO `stock_in_items` (`id`, `stock_in_id`, `product_id`, `product_name`, `quantity`, `unit_price`, `amount`, `batch_no`, `production_date`, `expiry_date`, `remark`) VALUES (8, '3', '1006', NULL, 8.00, 68.00, 544.00, '20251005003', '2025-09-12 00:00:00', '2028-09-11 00:00:00', '三七');
INSERT INTO `stock_in_items` (`id`, `stock_in_id`, `product_id`, `product_name`, `quantity`, `unit_price`, `amount`, `batch_no`, `production_date`, `expiry_date`, `remark`) VALUES (9, '4', '1007', NULL, 5.00, 120.00, 600.00, '20251007001', '2025-09-18 00:00:00', '2028-09-17 00:00:00', '川贝母');
INSERT INTO `stock_in_items` (`id`, `stock_in_id`, `product_id`, `product_name`, `quantity`, `unit_price`, `amount`, `batch_no`, `production_date`, `expiry_date`, `remark`) VALUES (10, '4', '1008', NULL, 12.00, 35.50, 426.00, '20251007002', '2025-09-22 00:00:00', '2027-09-21 00:00:00', '麦冬');
INSERT INTO `stock_in_items` (`id`, `stock_in_id`, `product_id`, `product_name`, `quantity`, `unit_price`, `amount`, `batch_no`, `production_date`, `expiry_date`, `remark`) VALUES (11, '5', '1009', NULL, 25.00, 8.60, 215.00, '20251009001', '2025-09-08 00:00:00', '2027-09-07 00:00:00', '甘草');
INSERT INTO `stock_in_items` (`id`, `stock_in_id`, `product_id`, `product_name`, `quantity`, `unit_price`, `amount`, `batch_no`, `production_date`, `expiry_date`, `remark`) VALUES (12, '5', '1010', NULL, 18.00, 22.30, 401.40, '20251009002', '2025-09-14 00:00:00', '2028-09-13 00:00:00', '红枣');
INSERT INTO `stock_in_items` (`id`, `stock_in_id`, `product_id`, `product_name`, `quantity`, `unit_price`, `amount`, `batch_no`, `production_date`, `expiry_date`, `remark`) VALUES (13, '6', '1001', NULL, 30.00, 12.40, 372.00, '20251012001', '2025-10-01 00:00:00', '2028-09-30 00:00:00', '当归（批量采购）');
INSERT INTO `stock_in_items` (`id`, `stock_in_id`, `product_id`, `product_name`, `quantity`, `unit_price`, `amount`, `batch_no`, `production_date`, `expiry_date`, `remark`) VALUES (14, '6', '1003', NULL, 25.00, 15.50, 387.50, '20251012002', '2025-09-28 00:00:00', '2027-09-27 00:00:00', '党参（新货）');
INSERT INTO `stock_in_items` (`id`, `stock_in_id`, `product_id`, `product_name`, `quantity`, `unit_price`, `amount`, `batch_no`, `production_date`, `expiry_date`, `remark`) VALUES (15, '6', '1004', NULL, 20.00, 28.30, 566.00, '20251012003', '2025-09-26 00:00:00', '2027-09-25 00:00:00', '枸杞（批量）');
INSERT INTO `stock_in_items` (`id`, `stock_in_id`, `product_id`, `product_name`, `quantity`, `unit_price`, `amount`, `batch_no`, `production_date`, `expiry_date`, `remark`) VALUES (16, '7', '1011', NULL, 10.00, 58.00, 580.00, '20251015001', '2025-09-10 00:00:00', '2028-09-09 00:00:00', '茯苓');
INSERT INTO `stock_in_items` (`id`, `stock_in_id`, `product_id`, `product_name`, `quantity`, `unit_price`, `amount`, `batch_no`, `production_date`, `expiry_date`, `remark`) VALUES (17, '7', '1012', NULL, 15.00, 18.70, 280.50, '20251015002', '2025-09-15 00:00:00', '2027-09-14 00:00:00', '白术');
INSERT INTO `stock_in_items` (`id`, `stock_in_id`, `product_id`, `product_name`, `quantity`, `unit_price`, `amount`, `batch_no`, `production_date`, `expiry_date`, `remark`) VALUES (18, '8', '1002', NULL, 50.00, 9.60, 480.00, '20251018001', '2025-10-05 00:00:00', '2028-10-04 00:00:00', '黄芪（季度储备）');
INSERT INTO `stock_in_items` (`id`, `stock_in_id`, `product_id`, `product_name`, `quantity`, `unit_price`, `amount`, `batch_no`, `production_date`, `expiry_date`, `remark`) VALUES (19, '8', '1007', NULL, 8.00, 118.00, 944.00, '20251018002', '2025-10-03 00:00:00', '2028-10-02 00:00:00', '川贝母（补充）');
INSERT INTO `stock_in_items` (`id`, `stock_in_id`, `product_id`, `product_name`, `quantity`, `unit_price`, `amount`, `batch_no`, `production_date`, `expiry_date`, `remark`) VALUES (20, '8', '1013', NULL, 30.00, 7.50, 225.00, '20251018003', '2025-09-20 00:00:00', '2027-09-19 00:00:00', '陈皮');
INSERT INTO `stock_in_items` (`id`, `stock_in_id`, `product_id`, `product_name`, `quantity`, `unit_price`, `amount`, `batch_no`, `production_date`, `expiry_date`, `remark`) VALUES (21, '9', '1014', NULL, 3.00, 380.00, 1140.00, '20251021001', '2025-08-15 00:00:00', '2028-08-14 00:00:00', '燕窝');
INSERT INTO `stock_in_items` (`id`, `stock_in_id`, `product_id`, `product_name`, `quantity`, `unit_price`, `amount`, `batch_no`, `production_date`, `expiry_date`, `remark`) VALUES (22, '9', '1015', NULL, 5.00, 260.00, 1300.00, '20251021002', '2025-08-25 00:00:00', '2028-08-24 00:00:00', '海参');
INSERT INTO `stock_in_items` (`id`, `stock_in_id`, `product_id`, `product_name`, `quantity`, `unit_price`, `amount`, `batch_no`, `production_date`, `expiry_date`, `remark`) VALUES (23, '10', '1008', NULL, 20.00, 35.00, 700.00, '20251024001', '2025-10-10 00:00:00', '2027-10-09 00:00:00', '麦冬（应急）');
INSERT INTO `stock_in_items` (`id`, `stock_in_id`, `product_id`, `product_name`, `quantity`, `unit_price`, `amount`, `batch_no`, `production_date`, `expiry_date`, `remark`) VALUES (24, '10', '1013', NULL, 25.00, 7.60, 190.00, '20251024002', '2025-10-08 00:00:00', '2027-10-07 00:00:00', '陈皮（补充）');
INSERT INTO `stock_in_items` (`id`, `stock_in_id`, `product_id`, `product_name`, `quantity`, `unit_price`, `amount`, `batch_no`, `production_date`, `expiry_date`, `remark`) VALUES (47, '37', 'by10001', '测试', 100.00, 10.00, 1000.00, '2025001', '2025-10-08 00:00:00', '2025-10-08 00:00:00', '');
INSERT INTO `stock_in_items` (`id`, `stock_in_id`, `product_id`, `product_name`, `quantity`, `unit_price`, `amount`, `batch_no`, `production_date`, `expiry_date`, `remark`) VALUES (52, '40', '10902', '甘草', 2.00, 18.00, 36.00, '2025001', '2025-10-08 00:00:00', '2025-10-08 00:00:00', '');
INSERT INTO `stock_in_items` (`id`, `stock_in_id`, `product_id`, `product_name`, `quantity`, `unit_price`, `amount`, `batch_no`, `production_date`, `expiry_date`, `remark`) VALUES (53, '40', '09023', '测试', 2.00, 3.00, 6.00, '', '2025-10-08 00:00:00', '2025-10-08 00:00:00', '');
INSERT INTO `stock_in_items` (`id`, `stock_in_id`, `product_id`, `product_name`, `quantity`, `unit_price`, `amount`, `batch_no`, `production_date`, `expiry_date`, `remark`) VALUES (54, '41', 'RS001', '', 12.00, 13.00, 156.00, '2025078', '2025-10-31 00:00:00', '2026-01-02 00:00:00', '');
INSERT INTO `stock_in_items` (`id`, `stock_in_id`, `product_id`, `product_name`, `quantity`, `unit_price`, `amount`, `batch_no`, `production_date`, `expiry_date`, `remark`) VALUES (55, '41', 'HQ002', '', 100.00, 16.00, 1600.00, '202508989', '2025-10-31 00:00:00', '2026-01-22 00:00:00', '');
INSERT INTO `stock_in_items` (`id`, `stock_in_id`, `product_id`, `product_name`, `quantity`, `unit_price`, `amount`, `batch_no`, `production_date`, `expiry_date`, `remark`) VALUES (56, '41', 'DG003', '', 200.00, 14.00, 2800.00, '2090033', '2025-10-31 00:00:00', '2025-11-14 00:00:00', '');
INSERT INTO `stock_in_items` (`id`, `stock_in_id`, `product_id`, `product_name`, `quantity`, `unit_price`, `amount`, `batch_no`, `production_date`, `expiry_date`, `remark`) VALUES (57, '41', 'GQ004', '', 400.00, 15.00, 6000.00, '20250608', '2025-10-31 00:00:00', '2025-12-19 00:00:00', '');
INSERT INTO `stock_in_items` (`id`, `stock_in_id`, `product_id`, `product_name`, `quantity`, `unit_price`, `amount`, `batch_no`, `production_date`, `expiry_date`, `remark`) VALUES (58, '41', 'SQ005', '', 300.00, 13.00, 3900.00, '20250709', '2025-10-31 00:00:00', '2025-12-05 00:00:00', '');
INSERT INTO `stock_in_items` (`id`, `stock_in_id`, `product_id`, `product_name`, `quantity`, `unit_price`, `amount`, `batch_no`, `production_date`, `expiry_date`, `remark`) VALUES (59, '41', 'JYH06', '', 200.00, 10.00, 2000.00, '20250802', '2025-10-31 00:00:00', '2026-01-02 00:00:00', '');
INSERT INTO `stock_in_items` (`id`, `stock_in_id`, `product_id`, `product_name`, `quantity`, `unit_price`, `amount`, `batch_no`, `production_date`, `expiry_date`, `remark`) VALUES (61, '44', 'RS001', '', 200.00, 10.00, 2000.00, '2090', '2025-10-31 00:00:00', '2025-10-31 00:00:00', '');
COMMIT;

-- ----------------------------
-- Table structure for stock_out
-- ----------------------------
DROP TABLE IF EXISTS `stock_out`;
CREATE TABLE `stock_out` (
  `id` int NOT NULL AUTO_INCREMENT,
  `order_no` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '出库单号',
  `warehouse_id` int NOT NULL COMMENT '仓库ID',
  `customer_id` varchar(255) DEFAULT NULL COMMENT '客户ID（销售出库时关联）',
  `out_type` tinyint NOT NULL COMMENT '出库类型（1-销售出库 2-退货出库 3-调拨出库 4-其他出库）',
  `operator_id` int NOT NULL COMMENT '操作人ID',
  `remark` varchar(500) DEFAULT NULL COMMENT '备注',
  `create_time` datetime DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_time` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `out_time` datetime DEFAULT NULL COMMENT '出库时间',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE KEY `order_no` (`order_no`),
  KEY `idx_warehouse` (`warehouse_id`),
  KEY `idx_out_type` (`out_type`)
) ENGINE=InnoDB AUTO_INCREMENT=17 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='出库单表';

-- ----------------------------
-- Records of stock_out
-- ----------------------------
BEGIN;
INSERT INTO `stock_out` (`id`, `order_no`, `warehouse_id`, `customer_id`, `out_type`, `operator_id`, `remark`, `create_time`, `update_time`, `out_time`) VALUES (6, 'pcORDER-20251031-781184', 2, '', 1, 1, '', '2025-10-31 11:09:47', '2025-10-31 11:09:47', NULL);
INSERT INTO `stock_out` (`id`, `order_no`, `warehouse_id`, `customer_id`, `out_type`, `operator_id`, `remark`, `create_time`, `update_time`, `out_time`) VALUES (7, 'pcORDER-20251031-535706', 2, '', 1, 1, '', '2025-10-31 11:12:20', '2025-10-31 11:12:20', NULL);
INSERT INTO `stock_out` (`id`, `order_no`, `warehouse_id`, `customer_id`, `out_type`, `operator_id`, `remark`, `create_time`, `update_time`, `out_time`) VALUES (8, 'pcORDER-20251031-691209', 2, '', 1, 1, '', '2025-10-31 11:13:13', '2025-10-31 11:13:13', NULL);
INSERT INTO `stock_out` (`id`, `order_no`, `warehouse_id`, `customer_id`, `out_type`, `operator_id`, `remark`, `create_time`, `update_time`, `out_time`) VALUES (14, 'pcORDER-20251031-839295', 2, '', 1, 1, '', '2025-10-31 23:08:44', '2025-10-31 23:08:44', NULL);
INSERT INTO `stock_out` (`id`, `order_no`, `warehouse_id`, `customer_id`, `out_type`, `operator_id`, `remark`, `create_time`, `update_time`, `out_time`) VALUES (15, 'pcORDER-20251101-076403', 2, '', 1, 1, '', '2025-11-01 18:22:39', '2025-11-01 18:22:39', NULL);
INSERT INTO `stock_out` (`id`, `order_no`, `warehouse_id`, `customer_id`, `out_type`, `operator_id`, `remark`, `create_time`, `update_time`, `out_time`) VALUES (16, 'pcORDER-20251101-275235', 1, '', 1, 1, '', '2025-11-01 18:23:07', '2025-11-01 18:23:07', NULL);
COMMIT;

-- ----------------------------
-- Table structure for stock_out_items
-- ----------------------------
DROP TABLE IF EXISTS `stock_out_items`;
CREATE TABLE `stock_out_items` (
  `id` int NOT NULL AUTO_INCREMENT,
  `stock_out_id` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '出库单ID',
  `product_id` varchar(255) NOT NULL COMMENT '商品ID',
  `quantity` varchar(255) NOT NULL COMMENT '数量',
  `unit_price` varchar(255) DEFAULT NULL COMMENT '单价',
  `amount` varchar(255) DEFAULT NULL COMMENT '金额',
  `batch_no` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL COMMENT '批次号',
  `production_date` varchar(255) DEFAULT NULL COMMENT '生产日期',
  `expiry_date` varchar(255) DEFAULT NULL COMMENT '过期日期',
  `remark` varchar(200) DEFAULT NULL COMMENT '备注',
  PRIMARY KEY (`id`) USING BTREE,
  KEY `idx_stock_out` (`stock_out_id`),
  KEY `idx_product` (`product_id`),
  KEY `idx_batch` (`batch_no`)
) ENGINE=InnoDB AUTO_INCREMENT=11 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='出库明细表';

-- ----------------------------
-- Records of stock_out_items
-- ----------------------------
BEGIN;
INSERT INTO `stock_out_items` (`id`, `stock_out_id`, `product_id`, `quantity`, `unit_price`, `amount`, `batch_no`, `production_date`, `expiry_date`, `remark`) VALUES (1, '6\0', '90023', '1', '', '', '', '', '', '');
INSERT INTO `stock_out_items` (`id`, `stock_out_id`, `product_id`, `quantity`, `unit_price`, `amount`, `batch_no`, `production_date`, `expiry_date`, `remark`) VALUES (2, '7', '90023', '1', '', '', '', '', '', '');
INSERT INTO `stock_out_items` (`id`, `stock_out_id`, `product_id`, `quantity`, `unit_price`, `amount`, `batch_no`, `production_date`, `expiry_date`, `remark`) VALUES (3, '8', '90023', '1', '', '', '', '', '', '');
INSERT INTO `stock_out_items` (`id`, `stock_out_id`, `product_id`, `quantity`, `unit_price`, `amount`, `batch_no`, `production_date`, `expiry_date`, `remark`) VALUES (8, '14', 'JYH06', '20', '', '', '', '', '', '');
INSERT INTO `stock_out_items` (`id`, `stock_out_id`, `product_id`, `quantity`, `unit_price`, `amount`, `batch_no`, `production_date`, `expiry_date`, `remark`) VALUES (9, '15', 'JYH06', '10', '', '', '', '', '', '');
INSERT INTO `stock_out_items` (`id`, `stock_out_id`, `product_id`, `quantity`, `unit_price`, `amount`, `batch_no`, `production_date`, `expiry_date`, `remark`) VALUES (10, '16', 'GQ004', '200', '', '', '', '', '', '');
COMMIT;

-- ----------------------------
-- Table structure for stock_transfer
-- ----------------------------
DROP TABLE IF EXISTS `stock_transfer`;
CREATE TABLE `stock_transfer` (
  `id` int NOT NULL AUTO_INCREMENT,
  `order_no` varchar(30) NOT NULL COMMENT '调拨单号',
  `from_warehouse_id` int NOT NULL COMMENT '调出仓库ID',
  `to_warehouse_id` int NOT NULL COMMENT '调入仓库ID',
  `operator_id` int NOT NULL COMMENT '操作人ID',
  `remark` varchar(500) DEFAULT NULL COMMENT '备注',
  `create_time` datetime DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_time` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE KEY `order_no` (`order_no`),
  KEY `idx_from_warehouse` (`from_warehouse_id`),
  KEY `idx_to_warehouse` (`to_warehouse_id`)
) ENGINE=InnoDB AUTO_INCREMENT=21 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='调拨单表';

-- ----------------------------
-- Records of stock_transfer
-- ----------------------------
BEGIN;
INSERT INTO `stock_transfer` (`id`, `order_no`, `from_warehouse_id`, `to_warehouse_id`, `operator_id`, `remark`, `create_time`, `update_time`) VALUES (1, 'TR-20251030-001', 1, 2, 101, '常规仓库调拨', '2025-10-28 09:30:00', '2025-10-28 09:30:00');
INSERT INTO `stock_transfer` (`id`, `order_no`, `from_warehouse_id`, `to_warehouse_id`, `operator_id`, `remark`, `create_time`, `update_time`) VALUES (2, 'TR-20251030-002', 2, 3, 102, '区域调货补充', '2025-10-28 14:15:00', '2025-10-28 14:15:00');
INSERT INTO `stock_transfer` (`id`, `order_no`, `from_warehouse_id`, `to_warehouse_id`, `operator_id`, `remark`, `create_time`, `update_time`) VALUES (3, 'TR-20251030-003', 3, 1, 103, '临期商品转移', '2025-10-29 10:00:00', '2025-10-29 10:00:00');
INSERT INTO `stock_transfer` (`id`, `order_no`, `from_warehouse_id`, `to_warehouse_id`, `operator_id`, `remark`, `create_time`, `update_time`) VALUES (4, 'TR-20251030-004', 1, 4, 101, '样品调拨', '2025-10-29 16:40:00', '2025-10-29 16:40:00');
INSERT INTO `stock_transfer` (`id`, `order_no`, `from_warehouse_id`, `to_warehouse_id`, `operator_id`, `remark`, `create_time`, `update_time`) VALUES (5, 'TR-20251030-005', 4, 2, 104, '季度库存调整', '2025-10-30 08:20:00', '2025-10-30 08:20:00');
INSERT INTO `stock_transfer` (`id`, `order_no`, `from_warehouse_id`, `to_warehouse_id`, `operator_id`, `remark`, `create_time`, `update_time`) VALUES (18, '209031', 1, 2, 102, '20250012', '2025-10-30 12:57:06', '2025-10-30 12:57:06');
INSERT INTO `stock_transfer` (`id`, `order_no`, `from_warehouse_id`, `to_warehouse_id`, `operator_id`, `remark`, `create_time`, `update_time`) VALUES (20, 'db90313', 1, 2, 101, '21', '2025-10-31 23:03:57', '2025-10-31 23:03:57');
COMMIT;

-- ----------------------------
-- Table structure for stock_transfer_items
-- ----------------------------
DROP TABLE IF EXISTS `stock_transfer_items`;
CREATE TABLE `stock_transfer_items` (
  `id` int NOT NULL AUTO_INCREMENT,
  `transfer_id` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '调拨单ID',
  `product_id` varchar(255) NOT NULL COMMENT '商品ID',
  `quantity` decimal(10,2) NOT NULL COMMENT '数量',
  `batch_no` varchar(30) DEFAULT NULL COMMENT '批次号',
  `remark` varchar(200) DEFAULT NULL COMMENT '备注',
  `product_name` varchar(50) DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  KEY `idx_transfer` (`transfer_id`),
  KEY `idx_product` (`product_id`)
) ENGINE=InnoDB AUTO_INCREMENT=18 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='调拨明细表';

-- ----------------------------
-- Records of stock_transfer_items
-- ----------------------------
BEGIN;
INSERT INTO `stock_transfer_items` (`id`, `transfer_id`, `product_id`, `quantity`, `batch_no`, `remark`, `product_name`) VALUES (1, '1', '1001', 50.00, 'B20251001', '日用品-洗发水', NULL);
INSERT INTO `stock_transfer_items` (`id`, `transfer_id`, `product_id`, `quantity`, `batch_no`, `remark`, `product_name`) VALUES (2, '1', '1002', 30.00, 'B20251002', '日用品-沐浴露', NULL);
INSERT INTO `stock_transfer_items` (`id`, `transfer_id`, `product_id`, `quantity`, `batch_no`, `remark`, `product_name`) VALUES (3, '2', '2001', 20.00, 'B20251003', '食品-方便面', NULL);
INSERT INTO `stock_transfer_items` (`id`, `transfer_id`, `product_id`, `quantity`, `batch_no`, `remark`, `product_name`) VALUES (4, '2', '2002', 15.00, 'B20251004', '食品-矿泉水', NULL);
INSERT INTO `stock_transfer_items` (`id`, `transfer_id`, `product_id`, `quantity`, `batch_no`, `remark`, `product_name`) VALUES (5, '2', '2003', 35.00, 'B20251005', '食品-面包', NULL);
INSERT INTO `stock_transfer_items` (`id`, `transfer_id`, `product_id`, `quantity`, `batch_no`, `remark`, `product_name`) VALUES (6, '3', '3001', 100.00, 'B20251006', '电子产品-充电线', NULL);
INSERT INTO `stock_transfer_items` (`id`, `transfer_id`, `product_id`, `quantity`, `batch_no`, `remark`, `product_name`) VALUES (7, '4', '4001', 5.00, 'B20251007', '样品-化妆品', NULL);
INSERT INTO `stock_transfer_items` (`id`, `transfer_id`, `product_id`, `quantity`, `batch_no`, `remark`, `product_name`) VALUES (8, '4', '4002', 8.00, 'B20251008', '样品-护肤品', NULL);
INSERT INTO `stock_transfer_items` (`id`, `transfer_id`, `product_id`, `quantity`, `batch_no`, `remark`, `product_name`) VALUES (9, '5', '5001', 60.00, 'B20251009', '服装-T恤', NULL);
INSERT INTO `stock_transfer_items` (`id`, `transfer_id`, `product_id`, `quantity`, `batch_no`, `remark`, `product_name`) VALUES (10, '5', '5002', 25.00, 'B20251010', '服装-裤子', NULL);
INSERT INTO `stock_transfer_items` (`id`, `transfer_id`, `product_id`, `quantity`, `batch_no`, `remark`, `product_name`) VALUES (11, '5', '5003', 15.00, 'B20251011', '服装-帽子', NULL);
INSERT INTO `stock_transfer_items` (`id`, `transfer_id`, `product_id`, `quantity`, `batch_no`, `remark`, `product_name`) VALUES (16, '18', '90023', 12.00, '20250021', '', '党参');
INSERT INTO `stock_transfer_items` (`id`, `transfer_id`, `product_id`, `quantity`, `batch_no`, `remark`, `product_name`) VALUES (17, '', 'JYH06', 100.00, '20990', '', '山楂');
COMMIT;

-- ----------------------------
-- Table structure for t_compatibility_taboo
-- ----------------------------
DROP TABLE IF EXISTS `t_compatibility_taboo`;
CREATE TABLE `t_compatibility_taboo` (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '禁忌记录唯一标识（自增主键）',
  `drug_id1` varchar(32) NOT NULL COMMENT '禁忌药品1（关联t_drug_basic.drug_id）',
  `drug_id2` varchar(32) NOT NULL COMMENT '禁忌药品2（关联t_drug_basic.drug_id）',
  `taboo_category` varchar(30) NOT NULL COMMENT '禁忌类别（如“十八反”“十九畏”“现代药理禁忌”）',
  `taboo_level` varchar(20) NOT NULL COMMENT '禁忌级别（如“禁止同用”“慎用”“不宜长期同用”）',
  `taboo_desc` text COMMENT '禁忌说明（如“乌头反半夏：两者同用可能增强毒性”）',
  `is_enabled` tinyint NOT NULL DEFAULT '1' COMMENT '是否启用（1=启用，0=禁用）',
  `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '记录创建时间',
  `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '记录更新时间',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE KEY `uk_drug_pair` (`drug_id1`,`drug_id2`),
  KEY `fk_taboo_drug1` (`drug_id1`),
  KEY `fk_taboo_drug2` (`drug_id2`),
  KEY `idx_taboo_category` (`taboo_category`),
  KEY `idx_taboo_level` (`taboo_level`)
) ENGINE=InnoDB AUTO_INCREMENT=25 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='中药配伍禁忌表';

-- ----------------------------
-- Records of t_compatibility_taboo
-- ----------------------------
BEGIN;
INSERT INTO `t_compatibility_taboo` (`id`, `drug_id1`, `drug_id2`, `taboo_category`, `taboo_level`, `taboo_desc`, `is_enabled`, `create_time`, `update_time`) VALUES (1, 'DR001', 'DR002', '十八反', '禁止同用', '乌头（附子）反半夏：两者同用可能增强乌头类毒性，引发心律失常等不良反应（依据《中国药典》2020年版一部）', 1, '2025-10-26 07:07:51', '2025-10-26 07:07:51');
INSERT INTO `t_compatibility_taboo` (`id`, `drug_id1`, `drug_id2`, `taboo_category`, `taboo_level`, `taboo_desc`, `is_enabled`, `create_time`, `update_time`) VALUES (2, 'DR001', 'DR003', '十八反', '禁止同用', '乌头（附子）反瓜蒌：传统认为两者配伍会产生毒性，现代研究提示可能增加心脏毒性风险', 1, '2025-10-26 07:07:51', '2025-10-26 07:07:51');
INSERT INTO `t_compatibility_taboo` (`id`, `drug_id1`, `drug_id2`, `taboo_category`, `taboo_level`, `taboo_desc`, `is_enabled`, `create_time`, `update_time`) VALUES (3, 'DR001', 'DR004', '十八反', '禁止同用', '乌头（附子）反川贝母：配伍后可能导致毒性成分叠加，引发中毒反应', 1, '2025-10-26 07:07:51', '2025-10-26 07:07:51');
INSERT INTO `t_compatibility_taboo` (`id`, `drug_id1`, `drug_id2`, `taboo_category`, `taboo_level`, `taboo_desc`, `is_enabled`, `create_time`, `update_time`) VALUES (4, 'DR001', 'DR005', '十八反', '禁止同用', '乌头（附子）反白蔹：属于经典十八反禁忌，临床应避免同用', 1, '2025-10-26 07:07:51', '2025-10-26 07:07:51');
INSERT INTO `t_compatibility_taboo` (`id`, `drug_id1`, `drug_id2`, `taboo_category`, `taboo_level`, `taboo_desc`, `is_enabled`, `create_time`, `update_time`) VALUES (5, 'DR001', 'DR006', '十八反', '禁止同用', '乌头（附子）反白及：配伍可能增强乌头碱毒性，导致神经系统及心脏损害', 1, '2025-10-26 07:07:51', '2025-10-26 07:07:51');
INSERT INTO `t_compatibility_taboo` (`id`, `drug_id1`, `drug_id2`, `taboo_category`, `taboo_level`, `taboo_desc`, `is_enabled`, `create_time`, `update_time`) VALUES (6, 'DR007', 'DR008', '十八反', '禁止同用', '甘草反甘遂：两者配伍可能拮抗药效或增加胃肠道毒性（依据《中国药典》配伍禁忌项）', 1, '2025-10-26 07:07:51', '2025-10-26 07:07:51');
INSERT INTO `t_compatibility_taboo` (`id`, `drug_id1`, `drug_id2`, `taboo_category`, `taboo_level`, `taboo_desc`, `is_enabled`, `create_time`, `update_time`) VALUES (7, 'DR007', 'DR009', '十八反', '禁止同用', '甘草反大戟：传统认为配伍会产生不良反应，现代研究提示可能影响药物代谢', 1, '2025-10-26 07:07:51', '2025-10-26 07:07:51');
INSERT INTO `t_compatibility_taboo` (`id`, `drug_id1`, `drug_id2`, `taboo_category`, `taboo_level`, `taboo_desc`, `is_enabled`, `create_time`, `update_time`) VALUES (8, 'DR007', 'DR010', '十八反', '禁止同用', '甘草反海藻：两者同用可能引发甲状腺功能异常或加重肝肾负担', 1, '2025-10-26 07:07:51', '2025-10-26 07:07:51');
INSERT INTO `t_compatibility_taboo` (`id`, `drug_id1`, `drug_id2`, `taboo_category`, `taboo_level`, `taboo_desc`, `is_enabled`, `create_time`, `update_time`) VALUES (9, 'DR007', 'DR011', '十八反', '禁止同用', '甘草反芫花：经典十八反禁忌，配伍可能增强芫花的泻下及毒性作用', 1, '2025-10-26 07:07:51', '2025-10-26 07:07:51');
INSERT INTO `t_compatibility_taboo` (`id`, `drug_id1`, `drug_id2`, `taboo_category`, `taboo_level`, `taboo_desc`, `is_enabled`, `create_time`, `update_time`) VALUES (10, 'DR012', 'DR013', '十八反', '禁止同用', '藜芦反人参：两者配伍可能降低人参补气功效，且藜芦的催吐作用会加重人参的胃肠刺激', 1, '2025-10-26 07:07:51', '2025-10-26 07:07:51');
INSERT INTO `t_compatibility_taboo` (`id`, `drug_id1`, `drug_id2`, `taboo_category`, `taboo_level`, `taboo_desc`, `is_enabled`, `create_time`, `update_time`) VALUES (11, 'DR012', 'DR014', '十八反', '禁止同用', '藜芦反沙参：属于十八反禁忌，临床应避免配伍使用', 1, '2025-10-26 07:07:51', '2025-10-26 07:07:51');
INSERT INTO `t_compatibility_taboo` (`id`, `drug_id1`, `drug_id2`, `taboo_category`, `taboo_level`, `taboo_desc`, `is_enabled`, `create_time`, `update_time`) VALUES (12, 'DR012', 'DR015', '十八反', '禁止同用', '藜芦反丹参：可能影响丹参的活血功效，且增加不良反应风险', 1, '2025-10-26 07:07:51', '2025-10-26 07:07:51');
INSERT INTO `t_compatibility_taboo` (`id`, `drug_id1`, `drug_id2`, `taboo_category`, `taboo_level`, `taboo_desc`, `is_enabled`, `create_time`, `update_time`) VALUES (13, 'DR012', 'DR016', '十八反', '禁止同用', '藜芦反玄参：传统认为两者配伍会产生毒性，现代临床亦不推荐同用', 1, '2025-10-26 07:07:51', '2025-10-26 07:07:51');
INSERT INTO `t_compatibility_taboo` (`id`, `drug_id1`, `drug_id2`, `taboo_category`, `taboo_level`, `taboo_desc`, `is_enabled`, `create_time`, `update_time`) VALUES (14, 'DR012', 'DR017', '十八反', '禁止同用', '藜芦反细辛：配伍可能引发中枢神经系统兴奋或抑制异常', 1, '2025-10-26 07:07:51', '2025-10-26 07:07:51');
INSERT INTO `t_compatibility_taboo` (`id`, `drug_id1`, `drug_id2`, `taboo_category`, `taboo_level`, `taboo_desc`, `is_enabled`, `create_time`, `update_time`) VALUES (15, 'DR012', 'DR018', '十八反', '禁止同用', '藜芦反白芍：经典禁忌，两者同用可能导致药效冲突或毒性增强', 1, '2025-10-26 07:07:51', '2025-10-26 07:07:51');
INSERT INTO `t_compatibility_taboo` (`id`, `drug_id1`, `drug_id2`, `taboo_category`, `taboo_level`, `taboo_desc`, `is_enabled`, `create_time`, `update_time`) VALUES (16, 'DR019', 'DR020', '十九畏', '禁止同用', '硫黄畏朴硝：两者配伍可能发生化学反应，生成有毒物质（如硫化物）', 1, '2025-10-26 07:07:51', '2025-10-26 07:07:51');
INSERT INTO `t_compatibility_taboo` (`id`, `drug_id1`, `drug_id2`, `taboo_category`, `taboo_level`, `taboo_desc`, `is_enabled`, `create_time`, `update_time`) VALUES (17, 'DR021', 'DR022', '十九畏', '禁止同用', '水银畏砒霜：配伍会生成剧毒的砷汞化合物，严禁同用', 1, '2025-10-26 07:07:51', '2025-10-26 07:07:51');
INSERT INTO `t_compatibility_taboo` (`id`, `drug_id1`, `drug_id2`, `taboo_category`, `taboo_level`, `taboo_desc`, `is_enabled`, `create_time`, `update_time`) VALUES (18, 'DR023', 'DR024', '十九畏', '禁止同用', '狼毒畏密陀僧：两者均有大毒，配伍会增强毒性，损伤脏腑', 1, '2025-10-26 07:07:51', '2025-10-26 07:07:51');
INSERT INTO `t_compatibility_taboo` (`id`, `drug_id1`, `drug_id2`, `taboo_category`, `taboo_level`, `taboo_desc`, `is_enabled`, `create_time`, `update_time`) VALUES (19, 'DR025', 'DR026', '十九畏', '禁止同用', '巴豆畏牵牛子：同用会剧烈刺激肠道，引发严重腹泻甚至脱水', 1, '2025-10-26 07:07:51', '2025-10-26 07:07:51');
INSERT INTO `t_compatibility_taboo` (`id`, `drug_id1`, `drug_id2`, `taboo_category`, `taboo_level`, `taboo_desc`, `is_enabled`, `create_time`, `update_time`) VALUES (20, 'DR027', 'DR028', '十九畏', '禁止同用', '丁香畏郁金：可能降低丁香温胃功效，且引发胃肠不适', 1, '2025-10-26 07:07:51', '2025-10-26 07:07:51');
INSERT INTO `t_compatibility_taboo` (`id`, `drug_id1`, `drug_id2`, `taboo_category`, `taboo_level`, `taboo_desc`, `is_enabled`, `create_time`, `update_time`) VALUES (21, 'DR001', 'DR029', '十九畏', '禁止同用', '川乌（乌头类）畏犀角：配伍可能产生毒性复合物，加重肝肾损伤', 1, '2025-10-26 07:07:51', '2025-10-26 07:07:51');
INSERT INTO `t_compatibility_taboo` (`id`, `drug_id1`, `drug_id2`, `taboo_category`, `taboo_level`, `taboo_desc`, `is_enabled`, `create_time`, `update_time`) VALUES (22, 'DR030', 'DR031', '十九畏', '禁止同用', '牙硝畏三棱：两者配伍可能刺激胃肠道，引发呕吐、腹痛', 1, '2025-10-26 07:07:51', '2025-10-26 07:07:51');
INSERT INTO `t_compatibility_taboo` (`id`, `drug_id1`, `drug_id2`, `taboo_category`, `taboo_level`, `taboo_desc`, `is_enabled`, `create_time`, `update_time`) VALUES (23, 'DR032', 'DR033', '十九畏', '禁止同用', '官桂畏赤石脂：可能降低官桂温阳功效，且影响药物吸收', 1, '2025-10-26 07:07:51', '2025-10-26 07:07:51');
INSERT INTO `t_compatibility_taboo` (`id`, `drug_id1`, `drug_id2`, `taboo_category`, `taboo_level`, `taboo_desc`, `is_enabled`, `create_time`, `update_time`) VALUES (24, 'DR013', 'DR034', '十九畏', '禁止同用', '人参畏五灵脂：配伍会降低人参补气作用，且可能引发消化不良', 1, '2025-10-26 07:07:51', '2025-10-26 07:07:51');
COMMIT;

-- ----------------------------
-- Table structure for t_drug_dosage_limit
-- ----------------------------
DROP TABLE IF EXISTS `t_drug_dosage_limit`;
CREATE TABLE `t_drug_dosage_limit` (
  `ID` bigint NOT NULL AUTO_INCREMENT COMMENT '剂量限制记录唯一标识（自增主键）',
  `drug_id` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '关联药品（t_drug_basic.drug_id）',
  `max_dosage` decimal(10,2) NOT NULL COMMENT '最大剂量（超此值触发超剂量预警）',
  `dosage_desc` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL COMMENT '剂量说明（如“炮制后使用可降低毒性”）',
  `is_enabled` tinyint NOT NULL DEFAULT '1' COMMENT '是否启用（1=启用，0=禁用）',
  `created_time` datetime DEFAULT CURRENT_TIMESTAMP COMMENT '记录创建时间',
  `updated_time` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '记录更新时间',
  PRIMARY KEY (`ID`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=12 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='药品剂量限制表（用于超剂量检查）';

-- ----------------------------
-- Records of t_drug_dosage_limit
-- ----------------------------
BEGIN;
INSERT INTO `t_drug_dosage_limit` (`ID`, `drug_id`, `max_dosage`, `dosage_desc`, `is_enabled`, `created_time`, `updated_time`) VALUES (1, 'drg101', 12.00, '1', 1, '2025-10-26 07:15:06', '2025-11-01 13:50:27');
INSERT INTO `t_drug_dosage_limit` (`ID`, `drug_id`, `max_dosage`, `dosage_desc`, `is_enabled`, `created_time`, `updated_time`) VALUES (2, 'DR002', 9.00, '成人常规内服（煎汤，炮制后）最大剂量；生品仅限外用，内服禁用超量', 1, '2025-10-26 07:15:06', '2025-10-26 07:15:06');
INSERT INTO `t_drug_dosage_limit` (`ID`, `drug_id`, `max_dosage`, `dosage_desc`, `is_enabled`, `created_time`, `updated_time`) VALUES (3, 'DR007', 10.00, '成人常规内服（煎汤）最大剂量；长期使用（>7天）建议不超过6g/日', 1, '2025-10-26 07:15:06', '2025-10-26 07:15:06');
INSERT INTO `t_drug_dosage_limit` (`ID`, `drug_id`, `max_dosage`, `dosage_desc`, `is_enabled`, `created_time`, `updated_time`) VALUES (4, 'DR013', 9.00, '成人常规内服（煎汤）最大剂量；急救场景可增至30g（需医师指导）', 1, '2025-10-26 07:15:06', '2025-10-26 07:15:06');
INSERT INTO `t_drug_dosage_limit` (`ID`, `drug_id`, `max_dosage`, `dosage_desc`, `is_enabled`, `created_time`, `updated_time`) VALUES (5, 'DR004', 10.00, '成人常规内服（煎汤）最大剂量；研末冲服每日不超过4g', 1, '2025-10-26 07:15:06', '2025-10-26 07:15:06');
INSERT INTO `t_drug_dosage_limit` (`ID`, `drug_id`, `max_dosage`, `dosage_desc`, `is_enabled`, `created_time`, `updated_time`) VALUES (6, 'DR008', 1.50, '成人内服（炮制后，入丸散）最大剂量，不宜入汤剂；孕妇禁用', 1, '2025-10-26 07:15:06', '2025-10-26 07:15:06');
INSERT INTO `t_drug_dosage_limit` (`ID`, `drug_id`, `max_dosage`, `dosage_desc`, `is_enabled`, `created_time`, `updated_time`) VALUES (7, 'DR012', 3.00, '成人内服（煎汤）最大剂量，毒性较强，需严格控制剂量', 1, '2025-10-26 07:15:06', '2025-10-26 07:15:06');
INSERT INTO `t_drug_dosage_limit` (`ID`, `drug_id`, `max_dosage`, `dosage_desc`, `is_enabled`, `created_time`, `updated_time`) VALUES (8, 'DR025', 0.60, '成人内服（去壳去油，入丸散）最大剂量，外用适量', 1, '2025-10-26 07:15:06', '2025-10-26 07:15:06');
COMMIT;

-- ----------------------------
-- Table structure for warehouses
-- ----------------------------
DROP TABLE IF EXISTS `warehouses`;
CREATE TABLE `warehouses` (
  `id` int NOT NULL AUTO_INCREMENT,
  `warehouse_code` varchar(20) NOT NULL,
  `warehouse_name` varchar(100) NOT NULL,
  `warehouse_type` varchar(50) DEFAULT NULL COMMENT '类型：仓库、药房',
  `location` varchar(200) DEFAULT NULL,
  `contact_person` varchar(50) DEFAULT NULL,
  `contact_phone` varchar(20) DEFAULT NULL,
  `status` tinyint DEFAULT '1' COMMENT '1-启用 0-禁用',
  `create_time` datetime DEFAULT CURRENT_TIMESTAMP,
  `update_time` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE KEY `warehouse_code` (`warehouse_code`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='仓库信息表';

-- ----------------------------
-- Records of warehouses
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for prescription_payment
-- ----------------------------
DROP TABLE IF EXISTS `prescription_payment`;
CREATE TABLE `prescription_payment` (
  `id` int NOT NULL AUTO_INCREMENT,
  `prescription_id` int NOT NULL COMMENT '处方ID',
  `prescription_number` varchar(60) NOT NULL COMMENT '处方号',
  `total_amount` decimal(12,2) NOT NULL DEFAULT 0.00 COMMENT '应付总金额',
  `pay_amount` decimal(12,2) NOT NULL DEFAULT 0.00 COMMENT '实付金额',
  `pay_method` varchar(30) DEFAULT NULL COMMENT '模拟支付方式',
  `pay_status` varchar(20) NOT NULL DEFAULT 'UNPAID' COMMENT '支付状态：UNPAID/PAID/CANCELLED',
  `mock_trade_no` varchar(64) DEFAULT NULL COMMENT '模拟交易号',
  `paid_at` datetime DEFAULT NULL COMMENT '支付时间',
  `created_at` datetime DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `uk_prescription_payment` (`prescription_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='处方模拟支付记录';

-- ----------------------------
-- Table structure for blockchain_trace_log
-- ----------------------------
DROP TABLE IF EXISTS `blockchain_trace_log`;
CREATE TABLE `blockchain_trace_log` (
  `id` int NOT NULL AUTO_INCREMENT,
  `prescription_id` int NOT NULL COMMENT '处方ID',
  `prescription_number` varchar(60) NOT NULL COMMENT '处方号',
  `payload_hash` varchar(64) NOT NULL COMMENT '上链内容SHA256',
  `tx_id` varchar(128) DEFAULT NULL COMMENT '链上交易ID',
  `chain_status` varchar(20) NOT NULL DEFAULT 'PENDING' COMMENT 'PENDING/SUCCESS/FAILED',
  `error_message` varchar(500) DEFAULT NULL,
  `created_at` datetime DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  KEY `idx_prescription_number` (`prescription_number`),
  KEY `idx_prescription_id` (`prescription_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='处方溯源上链日志';

-- ----------------------------
-- View structure for adjustment_view
-- ----------------------------
DROP VIEW IF EXISTS `adjustment_view`;
CREATE ALGORITHM = UNDEFINED SQL SECURITY DEFINER VIEW `adjustment_view` AS select `a`.`id` AS `id`,`p`.`hospital_name` AS `hospital_name`,`p`.`patient_name` AS `patient_name`,`p`.`prescription_number` AS `prescription_number`,`a`.`word_date` AS `word_date`,`a`.`end_date` AS `end_date`,`a`.`word_person` AS `word_person`,`a`.`status` AS `status`,`a`.`word_content` AS `word_content`,`a`.`workload_description` AS `workload_description`,`a`.`barcode` AS `barcode` from (`adjustment` `a` join `prescription` `p` on((`a`.`prescription_id` = `p`.`ID`)));

-- ----------------------------
-- View structure for drug_used_total_screen
-- ----------------------------
DROP VIEW IF EXISTS `drug_used_total_screen`;
CREATE ALGORITHM = UNDEFINED SQL SECURITY DEFINER VIEW `drug_used_total_screen` AS select `drug`.`drug_product_name` AS `drug_product_name`,sum(`drug`.`drug_weights`) AS `total_quantity` from `prescription_drug` `drug` group by `drug`.`drug_product_name` order by `total_quantity` desc limit 10;

-- ----------------------------
-- View structure for herb_decoction_audit_view
-- ----------------------------
DROP VIEW IF EXISTS `herb_decoction_audit_view`;
CREATE ALGORITHM = UNDEFINED SQL SECURITY DEFINER VIEW `herb_decoction_audit_view` AS select `audit`.`id` AS `id`,`p`.`hospital_name` AS `hospital_name`,`p`.`patient_name` AS `patient_name`,`p`.`prescription_number` AS `prescription_number`,`audit`.`barcode` AS `barcode`,`audit`.`word_content` AS `word_content`,`audit`.`reviewer` AS `reviewer`,`audit`.`audit_datetime` AS `audit_datetime`,`audit`.`audit_status` AS `audit_status`,`audit`.`employee_id` AS `employeeID` from (`prescription` `p` join `herb_decoction_audit` `audit` on((`p`.`ID` = `audit`.`prescription_id`)));

-- ----------------------------
-- View structure for herb_delivery_view
-- ----------------------------
DROP VIEW IF EXISTS `herb_delivery_view`;
CREATE ALGORITHM = UNDEFINED SQL SECURITY DEFINER VIEW `herb_delivery_view` AS select `a`.`id` AS `id`,`p`.`hospital_name` AS `hospital_name`,`p`.`patient_name` AS `patient_name`,`p`.`prescription_number` AS `prescription_number`,`a`.`prescription_id` AS `prescription_id`,`a`.`delivery_personnel` AS `delivery_personnel`,`a`.`delivery_time` AS `delivery_time`,`a`.`processing_employee_id` AS `processing_employee_id`,`a`.`word_content` AS `word_content`,`a`.`remarks` AS `remarks`,`a`.`delivery_status` AS `delivery_status`,`a`.`logistics_number` AS `logistics_number` from (`herb_delivery` `a` join `prescription` `p` on((`a`.`prescription_id` = `p`.`ID`)));

-- ----------------------------
-- View structure for herb_soaking_view
-- ----------------------------
DROP VIEW IF EXISTS `herb_soaking_view`;
CREATE ALGORITHM = UNDEFINED SQL SECURITY DEFINER VIEW `herb_soaking_view` AS select `soak`.`id` AS `id`,`p`.`hospital_name` AS `hospital_name`,`p`.`patient_name` AS `patient_name`,`p`.`prescription_number` AS `prescription_number`,`soak`.`barcode` AS `barcode`,`soak`.`soaking_person` AS `soaking_person`,`soak`.`start_time` AS `start_time`,`soak`.`end_time` AS `end_time`,`soak`.`duration` AS `duration`,`soak`.`word_content` AS `word_content`,`soak`.`remarks` AS `remarks`,`soak`.`prescription_id` AS `prescription_id`,`soak`.`employee_id` AS `employee_id` from (`prescription` `p` join `herb_soaking` `soak` on((`p`.`ID` = `soak`.`prescription_id`)));

-- ----------------------------
-- View structure for herbal_decoction_info_view
-- ----------------------------
DROP VIEW IF EXISTS `herbal_decoction_info_view`;
CREATE ALGORITHM = UNDEFINED SQL SECURITY DEFINER VIEW `herbal_decoction_info_view` AS select `decoction`.`id` AS `id`,`p`.`hospital_name` AS `hospital_name`,`p`.`patient_name` AS `patient_name`,`p`.`prescription_number` AS `prescription_number`,`decoction`.`barcode` AS `barcode`,`decoction`.`word_content` AS `word_content`,`decoction`.`decoction_manager` AS `decoction_manager`,`decoction`.`decoction_time` AS `decoction_time`,`decoction`.`start_time` AS `start_time`,`decoction`.`end_time` AS `end_time`,`decoction`.`decoction_status` AS `decoction_status`,`decoction`.`employee_id` AS `employee_id`,`decoction`.`machine_id` AS `machine_id`,`decoction`.`remark` AS `remark` from (`herbal_decoction_info` `decoction` join `prescription` `p` on((`decoction`.`prescription_id` = `p`.`ID`)));

-- ----------------------------
-- View structure for hospital_prescption_tj_vw
-- ----------------------------
DROP VIEW IF EXISTS `hospital_prescption_tj_vw`;
CREATE ALGORITHM = UNDEFINED SQL SECURITY DEFINER VIEW `hospital_prescption_tj_vw` AS select `prescription`.`hospital_name` AS `hospital_name`,count(0) AS `prescription_count` from `prescription` group by `prescription`.`hospital_name` order by `prescription_count` desc limit 10;

-- ----------------------------
-- View structure for hospital_tj_vw
-- ----------------------------
DROP VIEW IF EXISTS `hospital_tj_vw`;
CREATE ALGORITHM = UNDEFINED SQL SECURITY DEFINER VIEW `hospital_tj_vw` AS select `t`.`医院` AS `hospital_name`,`t`.`开方数量` AS `prescription_count`,`t`.`处方加工费` AS `prescription_fee`,`t`.`药品费用` AS `drug_fee`,(`t`.`处方加工费` + `t`.`药品费用`) AS `prescription_total_fee`,`t`.`日期` AS `dot_time` from (select `p`.`hospital_name` AS `医院`,count(distinct `p`.`ID`) AS `开方数量`,(`p`.`dosage` * `fe`.`fee`) AS `处方加工费`,sum(`pd`.`total_prices`) AS `药品费用`,`p`.`do_time` AS `日期` from ((`prescription` `p` join `proess_fee` `fe` on((`fe`.`hospital_id` = `p`.`hospital_id`))) join `prescription_drug` `pd` on((`p`.`ID` = `pd`.`prescription_id`))) group by `p`.`hospital_name`,`fe`.`fee`,`p`.`dosage`,`p`.`do_time`) `t`;

-- ----------------------------
-- View structure for medicine_packing_view
-- ----------------------------
DROP VIEW IF EXISTS `medicine_packing_view`;
CREATE ALGORITHM = UNDEFINED SQL SECURITY DEFINER VIEW `medicine_packing_view` AS select `a`.`id` AS `id`,`a`.`word_content` AS `word_content`,`p`.`hospital_name` AS `hospital_name`,`p`.`patient_name` AS `patient_name`,`p`.`prescription_number` AS `prescription_number`,`a`.`barcode` AS `barcode`,`a`.`start_time` AS `start_time`,`a`.`employee_id` AS `employee_id`,`a`.`packing_personnel` AS `packing_personnel`,`a`.`packing_status` AS `packing_status`,`a`.`end_time` AS `end_time` from (`medicine_packing` `a` join `prescription` `p` on((`a`.`prescription_id` = `p`.`ID`)));

-- ----------------------------
-- View structure for oh_admin_user_role_view
-- ----------------------------
DROP VIEW IF EXISTS `oh_admin_user_role_view`;
CREATE ALGORITHM = UNDEFINED SQL SECURITY DEFINER VIEW `oh_admin_user_role_view` AS select `ur`.`id` AS `id`,`ur`.`realname` AS `realname`,`ur`.`password` AS `password`,`ur`.`username` AS `username`,`adr`.`name` AS `name`,`ur`.`app_token` AS `token`,`ur`.`status` AS `status` from ((`oh_admin_user` `ur` left join `oh_admin_user_role` `ro` on((`ro`.`user_id` = `ur`.`id`))) left join `oh_admin_role` `adr` on((`adr`.`id` = `ro`.`role_id`)));

-- ----------------------------
-- View structure for prescription_audit_view
-- ----------------------------
DROP VIEW IF EXISTS `prescription_audit_view`;
CREATE ALGORITHM = UNDEFINED SQL SECURITY DEFINER VIEW `prescription_audit_view` AS select `pa`.`ID` AS `PrescriptionID`,`pa`.`prescription_id` AS `id`,`pa`.`reviewer` AS `Reviewer`,`pa`.`review_time` AS `ReviewTime`,`pa`.`audit_status` AS `AuditStatus`,`pa`.`rejection_reason` AS `RejectionReason`,`pa`.`print_status` AS `PrintStatus`,`pa`.`employee_id` AS `EmployeeID`,`pa`.`create_time` AS `create_time`,`p`.`deletion_number` AS `deletion_number`,`p`.`Is_decoction` AS `Is_decoction`,`p`.`barcode_scan` AS `barcode_scan`,`p`.`hospital_id` AS `hospital_id`,`p`.`hospital_name` AS `hospital_name`,`p`.`prescription_number` AS `prescription_number`,`p`.`decoction_method` AS `decoction_method`,`p`.`prescription_type` AS `prescription_type`,`p`.`patient_name` AS `patient_name`,`p`.`patient_sex` AS `patient_sex`,`p`.`patient_age` AS `patient_age`,`p`.`patient_phone` AS `patient_phone`,`p`.`patient_address` AS `patient_address`,`p`.`department_name` AS `department_name`,`p`.`inpatient_area` AS `inpatient_area`,`p`.`ward_name` AS `ward_name`,`p`.`sick_bed` AS `sick_bed`,`p`.`diagnosis_result` AS `diagnosis_result`,`p`.`dosage` AS `dosage`,`p`.`administration_method` AS `administration_method`,`p`.`administration_count` AS `administration_count`,`p`.`package_count` AS `package_count`,`p`.`decoction_scheme` AS `decoction_scheme`,`p`.`one_time_dosage` AS `one_time_dosage`,`p`.`two_time_dosage` AS `two_time_dosage`,`p`.`soak_water_amount` AS `soak_water_amount`,`p`.`soak_time` AS `soak_time`,`p`.`label_number` AS `label_number`,`p`.`remarks` AS `remarks`,`p`.`doctor_name` AS `doctor_name`,`p`.`footnote` AS `footnote`,`p`.`drug_pickup_time` AS `drug_pickup_time`,`p`.`drug_pickup_number` AS `drug_pickup_number`,`p`.`order_time` AS `order_time`,`p`.`current_state` AS `current_state`,`p`.`do_time` AS `do_time`,`p`.`do_person` AS `do_person`,`p`.`distribution_company` AS `distribution_company`,`p`.`distribution_address` AS `distribution_address`,`p`.`distribution_phone` AS `distribution_phone`,`p`.`distribution_type` AS `distribution_type`,`p`.`administration_way` AS `administration_way`,`p`.`additional_remarks_a` AS `additional_remarks_a`,`p`.`additional_remarks_b` AS `additional_remarks_b`,`p`.`drug_confirmation` AS `drug_confirmation`,`p`.`logistics_state` AS `logistics_state`,`p`.`query_time` AS `query_time`,`p`.`query_person` AS `query_person`,`p`.`drug_count` AS `drug_count`,`p`.`machine_room` AS `machine_room`,`p`.`state_post_back` AS `state_post_back`,`p`.`subside_time` AS `subside_time` from (`prescription_audit` `pa` left join `prescription` `p` on((`pa`.`prescription_id` = `p`.`ID`)));

-- ----------------------------
-- View structure for prescription_do_seven_day_screen
-- ----------------------------
DROP VIEW IF EXISTS `prescription_do_seven_day_screen`;
CREATE ALGORITHM = UNDEFINED SQL SECURITY DEFINER VIEW `prescription_do_seven_day_screen` AS select `dates`.`do_time` AS `do_time`,coalesce(count(`prescription`.`ID`),0) AS `prescription_count` from ((select (curdate() - interval (`a`.`a` + (10 * `b`.`a`)) day) AS `do_time` from ((select 0 AS `a` union all select 1 AS `1` union all select 2 AS `2` union all select 3 AS `3` union all select 4 AS `4` union all select 5 AS `5` union all select 6 AS `6` union all select 7 AS `7` union all select 8 AS `8` union all select 9 AS `9`) `a` join (select 0 AS `a` union all select 1 AS `1` union all select 2 AS `2` union all select 3 AS `3` union all select 4 AS `4` union all select 5 AS `5` union all select 6 AS `6` union all select 7 AS `7` union all select 8 AS `8` union all select 9 AS `9`) `b`) where ((curdate() - interval (`a`.`a` + (10 * `b`.`a`)) day) >= (curdate() - interval 7 day))) `dates` left join `prescription` on((cast(`prescription`.`do_time` as date) = `dates`.`do_time`))) group by `dates`.`do_time` order by `dates`.`do_time` desc;

-- ----------------------------
-- View structure for prescription_finish_seven_day_screen
-- ----------------------------
DROP VIEW IF EXISTS `prescription_finish_seven_day_screen`;
CREATE ALGORITHM = UNDEFINED SQL SECURITY DEFINER VIEW `prescription_finish_seven_day_screen` AS select `subquery`.`delivery_date` AS `delivery_date`,coalesce(count(distinct `prescription`.`ID`),0) AS `prescription_count` from (((select (curdate() - interval ((`a`.`a` + (10 * `b`.`a`)) + (100 * `c`.`a`)) day) AS `delivery_date` from (((select 0 AS `a` union all select 1 AS `1` union all select 2 AS `2` union all select 3 AS `3` union all select 4 AS `4` union all select 5 AS `5` union all select 6 AS `6` union all select 7 AS `7` union all select 8 AS `8` union all select 9 AS `9`) `a` join (select 0 AS `a` union all select 1 AS `1` union all select 2 AS `2` union all select 3 AS `3` union all select 4 AS `4` union all select 5 AS `5` union all select 6 AS `6` union all select 7 AS `7` union all select 8 AS `8` union all select 9 AS `9`) `b`) join (select 0 AS `a` union all select 1 AS `1` union all select 2 AS `2` union all select 3 AS `3` union all select 4 AS `4` union all select 5 AS `5` union all select 6 AS `6` union all select 7 AS `7` union all select 8 AS `8` union all select 9 AS `9`) `c`) where ((curdate() - interval ((`a`.`a` + (10 * `b`.`a`)) + (100 * `c`.`a`)) day) >= (curdate() - interval 7 day))) `subquery` left join `herb_delivery` on((cast(`herb_delivery`.`delivery_time` as date) = `subquery`.`delivery_date`))) left join `prescription` on((`herb_delivery`.`prescription_id` = `prescription`.`ID`))) group by `subquery`.`delivery_date` order by `subquery`.`delivery_date` desc;

-- ----------------------------
-- View structure for prescription_is_decoction_screen
-- ----------------------------
DROP VIEW IF EXISTS `prescription_is_decoction_screen`;
CREATE ALGORITHM = UNDEFINED SQL SECURITY DEFINER VIEW `prescription_is_decoction_screen` AS select '代煎' AS `decoction_type`,sum((case when (`prescription`.`Is_decoction` = 1) then 1 else 0 end)) AS `decoction_count` from `prescription` where (cast(`prescription`.`do_time` as date) = curdate()) union all select '代配' AS `decoction_type`,sum((case when (`prescription`.`Is_decoction` = 0) then 1 else 0 end)) AS `decoction_count` from `prescription` where (cast(`prescription`.`do_time` as date) = curdate());

-- ----------------------------
-- View structure for prescription_list_screen
-- ----------------------------
DROP VIEW IF EXISTS `prescription_list_screen`;
CREATE ALGORITHM = UNDEFINED SQL SECURITY DEFINER VIEW `prescription_list_screen` AS select `p`.`hospital_name` AS `hospital_name`,(case when (`p`.`Is_decoction` = 1) then '代煎' else '自煎' end) AS `decoction_type`,`p`.`prescription_number` AS `prescription_number`,`p`.`patient_name` AS `patient_name`,`p`.`dosage` AS `dosage`,`p`.`current_state` AS `current_state`,`p`.`drug_count` AS `drug_count`,coalesce(`pay`.`pay_status`,'UNPAID') AS `pay_status`,coalesce(`pay`.`pay_amount`,0) AS `pay_amount`,coalesce(`chain`.`chain_status`,'') AS `chain_status`,coalesce(`chain`.`tx_id`,'') AS `tx_id` from ((`prescription` `p` left join `prescription_payment` `pay` on((`pay`.`prescription_id` = `p`.`ID`))) left join (select `b1`.* from (`blockchain_trace_log` `b1` join (select `blockchain_trace_log`.`prescription_id` AS `prescription_id`,max(`blockchain_trace_log`.`id`) AS `id` from `blockchain_trace_log` group by `blockchain_trace_log`.`prescription_id`) `latest` on((`latest`.`id` = `b1`.`id`)))) `chain` on((`chain`.`prescription_id` = `p`.`ID`))) where (cast(`p`.`do_time` as date) = curdate());

-- ----------------------------
-- View structure for prescription_total_screen
-- ----------------------------
DROP VIEW IF EXISTS `prescription_total_screen`;
CREATE ALGORITHM = UNDEFINED SQL SECURITY DEFINER VIEW `prescription_total_screen` AS select sum((case when (`prescription`.`current_state` = '接方') then 1 else 0 end)) AS `prescription_total`,sum((case when (`prescription`.`current_state` = '发货') then 1 else 0 end)) AS `prescription_finish_total`,sum((case when (`prescription`.`current_state` = '审核') then 1 else 0 end)) AS `prescription_aduit_total`,sum((case when (`prescription`.`current_state` = '泡药') then 1 else 0 end)) AS `prescription_soak_total`,sum((case when (`prescription`.`current_state` = '煎药') then 1 else 0 end)) AS `prescription_decoction_total`,sum((case when (`prescription`.`current_state` = '包装') then 1 else 0 end)) AS `prescription_pack_total` from `prescription` where ((`prescription`.`do_time` >= curdate()) and (`prescription`.`do_time` < (curdate() + interval 1 day)));

-- ----------------------------
-- View structure for qrcode_vw
-- ----------------------------
DROP VIEW IF EXISTS `qrcode_vw`;
CREATE ALGORITHM = UNDEFINED SQL SECURITY DEFINER VIEW `qrcode_vw` AS select `p`.`ID` AS `id`,`p`.`hospital_id` AS `hospital_id`,`p`.`hospital_name` AS `hospital_name`,`p`.`prescription_number` AS `prescription_number`,`p`.`patient_name` AS `patient_name`,`p`.`patient_address` AS `patient_address`,right(concat('0',lpad(cast((`p`.`dosage` * `p`.`administration_count`) as char charset utf8mb4),20,'0')),2) AS `PackNum`,right(concat('0',lpad(cast(`p`.`decoction_scheme` as char charset utf8mb4),20,'0')),2) AS `DeScheme`,right(concat('000',lpad(cast(`p`.`package_count` as char charset utf8mb4),20,'0')),3) AS `PackAcount`,`p`.`administration_count` AS `administration_count`,right(concat('0000000',lpad(cast(`p`.`ID` as char charset utf8mb4),20,'0')),10) AS `BNum` from `prescription` `p`;

-- ----------------------------
-- View structure for vw_prescription_dosage_screen
-- ----------------------------
DROP VIEW IF EXISTS `vw_prescription_dosage_screen`;
CREATE ALGORITHM = UNDEFINED SQL SECURITY DEFINER VIEW `vw_prescription_dosage_screen` AS with recursive `DateSeries` as (select (curdate() - interval 6 day) AS `do_time` union all select (`DateSeries`.`do_time` + interval 1 day) AS `do_time + INTERVAL 1 DAY` from `DateSeries` where (`DateSeries`.`do_time` < curdate())) select `ds`.`do_time` AS `do_time`,coalesce(sum(`p`.`dosage`),0) AS `prescription_dosage_count` from (`DateSeries` `ds` left join `prescription` `p` on((cast(`p`.`do_time` as date) = `ds`.`do_time`))) group by `ds`.`do_time` order by `ds`.`do_time`;

-- ----------------------------
-- View structure for workload_vw
-- ----------------------------
DROP VIEW IF EXISTS `workload_vw`;
CREATE ALGORITHM = UNDEFINED SQL SECURITY DEFINER VIEW `workload_vw` AS select `adu`.`word_person` AS `word_person`,`adu`.`word_content` AS `word_content`,sum(`p`.`dosage`) AS `dosage`,`adu`.`word_date` AS `word_date` from (`prescription` `p` join `adjustment` `adu` on((`adu`.`prescription_id` = `p`.`ID`))) group by `adu`.`word_person`,`adu`.`word_content`,`p`.`dosage`,`adu`.`word_date` union all select `audit`.`reviewer` AS `word_person`,`audit`.`word_content` AS `word_content`,`p`.`dosage` AS `dosage`,`audit`.`audit_datetime` AS `word_date` from (`prescription` `p` join `herb_decoction_audit` `audit` on((`audit`.`prescription_id` = `p`.`ID`))) group by `word_person`,`audit`.`word_content`,`p`.`dosage`,`word_date` union all select `soak`.`soaking_person` AS `word_person`,`soak`.`word_content` AS `word_content`,`p`.`dosage` AS `dosage`,`soak`.`start_time` AS `word_date` from (`herb_soaking` `soak` join `prescription` `p` on((`soak`.`prescription_id` = `p`.`ID`))) union all select `deco`.`decoction_manager` AS `word_person`,`deco`.`word_content` AS `word_content`,`p`.`dosage` AS `dosage`,`deco`.`start_time` AS `word_date` from (`herbal_decoction_info` `deco` join `prescription` `p` on((`deco`.`prescription_id` = `p`.`ID`))) union all select `pack`.`packing_personnel` AS `word_person`,`pack`.`word_content` AS `word_content`,`p`.`dosage` AS `dosage`,`pack`.`start_time` AS `word_date` from (`medicine_packing` `pack` join `prescription` `p` on((`pack`.`prescription_id` = `p`.`ID`))) union all select `delivery`.`delivery_personnel` AS `word_person`,`delivery`.`word_content` AS `word_content`,`p`.`dosage` AS `dosage`,`delivery`.`delivery_time` AS `delivery_time` from (`herb_delivery` `delivery` join `prescription` `p` on((`delivery`.`prescription_id` = `p`.`ID`)));

-- ----------------------------
-- View structure for wx_prescription_view
-- ----------------------------
DROP VIEW IF EXISTS `wx_prescription_view`;
CREATE ALGORITHM = UNDEFINED SQL SECURITY DEFINER VIEW `wx_prescription_view` AS select `pres`.`ID` AS `id`,`pres`.`hospital_name` AS `hospital_name`,`pres`.`patient_name` AS `patient_name`,`pres`.`prescription_number` AS `prescription_number`,`pres`.`dosage` AS `dosage`,`pres`.`decoction_method` AS `decoction_method`,`pres`.`decoction_scheme` AS `decoction_scheme`,`pres`.`do_person` AS `do_person`,`pres`.`do_time` AS `do_time`,`pres_aduit`.`reviewer` AS `pres_aduit_reviewer`,`pres_aduit`.`review_time` AS `pres_aduit_time`,`adj`.`word_person` AS `adjustment_reviewer`,`adj`.`word_date` AS `adjustment_time`,`dec_aduit`.`audit_datetime` AS `audit_datetime`,`dec_aduit`.`reviewer` AS `audit_reviewer`,`soak`.`soaking_person` AS `soaking_person`,`soak`.`start_time` AS `soak_start_time`,`soak`.`end_time` AS `soak_end_time`,`deoc`.`decoction_manager` AS `decoction_person`,`deoc`.`start_time` AS `decoction_start_time`,`deoc`.`end_time` AS `decoction_end_time`,`deoc`.`machine_id` AS `machine_id`,`pack`.`packing_personnel` AS `packing_personnel`,`pack`.`start_time` AS `pack_start_time`,`pack`.`end_time` AS `pack_end_time`,`delivery`.`delivery_personnel` AS `delivery_personnel`,`delivery`.`delivery_time` AS `delivery_time`,`delivery`.`logistics_number` AS `logistics_number`,`pres`.`current_state` AS `current_state` from (((((((`prescription` `pres` left join `prescription_audit` `pres_aduit` on((`pres`.`ID` = `pres_aduit`.`prescription_id`))) left join `adjustment` `adj` on((`adj`.`prescription_id` = `pres`.`ID`))) left join `herb_decoction_audit` `dec_aduit` on((`pres`.`ID` = `dec_aduit`.`prescription_id`))) left join `herb_soaking` `soak` on((`soak`.`prescription_id` = `pres`.`ID`))) left join `herbal_decoction_info` `deoc` on((`deoc`.`prescription_id` = `pres`.`ID`))) left join `medicine_packing` `pack` on((`pres`.`ID` = `pack`.`prescription_id`))) left join `herb_delivery` `delivery` on((`delivery`.`prescription_id` = `pres`.`ID`)));

SET FOREIGN_KEY_CHECKS = 1;

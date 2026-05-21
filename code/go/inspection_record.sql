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

 Date: 03/11/2025 16:15:37
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

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

SET FOREIGN_KEY_CHECKS = 1;

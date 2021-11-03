/*
 Navicat Premium Data Transfer

 Source Server         : MySQL-本地
 Source Server Type    : MySQL
 Source Server Version : 80021
 Source Host           : localhost:3306
 Source Schema         : micro_mall_shop

 Target Server Type    : MySQL
 Target Server Version : 80021
 File Encoding         : 65001

 Date: 18/07/2021 13:24:41
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for shop_business
-- ----------------------------
DROP TABLE IF EXISTS `shop_business`;
CREATE TABLE `shop_business` (
  `shop_id` bigint NOT NULL AUTO_INCREMENT COMMENT '店铺ID',
  `nick_name` varchar(512) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '简称',
  `shop_code` char(36) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '店铺唯一code',
  `full_name` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '店铺全称',
  `register_addr` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '注册地址',
  `business_addr` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '实际经营地址',
  `legal_person` bigint NOT NULL COMMENT '店铺法人',
  `business_license` char(36) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '经营许可证',
  `tax_card_no` char(36) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '纳税号',
  `business_desc` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '经营描述',
  `social_credit_code` char(36) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '统一社会信用代码',
  `organization_code` char(36) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '组织机构代码',
  `state` tinyint NOT NULL DEFAULT '2' COMMENT '状态，0-未审核，1-审核不通过，2-审核通过',
  `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间',
  PRIMARY KEY (`shop_id`) USING BTREE,
  UNIQUE KEY `legal_person_nick_name_index` (`legal_person`,`nick_name`) USING BTREE COMMENT '法人店铺名索引',
  UNIQUE KEY `shop_code_index` (`shop_code`) USING BTREE COMMENT '店铺唯一code',
  KEY `legal_person_index` (`legal_person`) USING BTREE COMMENT '店铺f法人'
) ENGINE=InnoDB AUTO_INCREMENT=1000 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='店铺主体登记表';

SET FOREIGN_KEY_CHECKS = 1;

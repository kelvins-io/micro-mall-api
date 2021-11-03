/*
 Navicat Premium Data Transfer

 Source Server         : MySQL-本地
 Source Server Type    : MySQL
 Source Server Version : 80021
 Source Host           : localhost:3306
 Source Schema         : micro_mall_sku

 Target Server Type    : MySQL
 Target Server Version : 80021
 File Encoding         : 65001

 Date: 18/07/2021 13:24:34
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for sku_inventory
-- ----------------------------
DROP TABLE IF EXISTS `sku_inventory`;
CREATE TABLE `sku_inventory` (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '商品库存ID',
  `sku_code` char(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '商品编码',
  `amount` bigint DEFAULT NULL COMMENT '库存数量',
  `price` decimal(32,16) DEFAULT NULL COMMENT '入库单价',
  `shop_id` bigint NOT NULL COMMENT '所属店铺ID',
  `version` int NOT NULL DEFAULT '1' COMMENT '商品版本',
  `last_tx_id` char(60) COLLATE utf8mb4_general_ci NOT NULL DEFAULT 'dd13b4aa-4121-4898-a2b5-bcfebccb713b' COMMENT '最后一次更新事务ID',
  `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE KEY `sku_code_index` (`sku_code`) USING BTREE COMMENT '商品编码code',
  KEY `shop_id_index` (`shop_id`) USING BTREE COMMENT '店铺ID索引',
  KEY `last_tx_id_index` (`last_tx_id`) USING BTREE COMMENT '最后一次修改事务ID索引'
) ENGINE=InnoDB AUTO_INCREMENT=193 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='商品库存表';

-- ----------------------------
-- Table structure for sku_inventory_record
-- ----------------------------
DROP TABLE IF EXISTS `sku_inventory_record`;
CREATE TABLE `sku_inventory_record` (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '自责ID',
  `out_trade_no` char(40) COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '外部订单号',
  `shop_id` bigint DEFAULT NULL COMMENT '店铺ID',
  `sku_code` char(40) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '商品sku',
  `op_type` tinyint DEFAULT '0' COMMENT '操作类型，0-入库，1-出库，2-冻结',
  `op_uid` bigint DEFAULT NULL COMMENT '操作的用户ID',
  `op_ip` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '操作IP地址',
  `amount_before` bigint DEFAULT NULL COMMENT '变化之前数量',
  `amount` bigint DEFAULT NULL COMMENT '操作数量',
  `op_tx_id` char(60) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '操作的事务ID',
  `state` tinyint DEFAULT '0' COMMENT '状态，0-有效，1-锁定中，2-无效',
  `verify` tinyint DEFAULT '0' COMMENT '是否核实，0-未核实，1-已核实',
  `create_time` datetime DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_time` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间',
  PRIMARY KEY (`id`) USING BTREE,
  KEY `op_tx_id_index` (`op_tx_id`) USING BTREE COMMENT '操作事务ID',
  KEY `shop_id_index` (`shop_id`) USING BTREE COMMENT '店铺ID',
  KEY `sku_code_index` (`sku_code`) USING BTREE COMMENT '商品sku',
  KEY `out_trade_no_index` (`out_trade_no`) USING BTREE COMMENT '外部订单号',
  KEY `verify_op_type_index` (`verify`,`op_type`) USING BTREE COMMENT '操作类型-库存验证'
) ENGINE=InnoDB AUTO_INCREMENT=1000 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='商品库存记录';

-- ----------------------------
-- Table structure for sku_price_history
-- ----------------------------
DROP TABLE IF EXISTS `sku_price_history`;
CREATE TABLE `sku_price_history` (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '自增ID',
  `shop_id` bigint NOT NULL COMMENT '调价的店铺id',
  `sku_code` char(40) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '商品sku_code',
  `price` decimal(32,16) NOT NULL COMMENT '调整后价格',
  `reason` text CHARACTER SET utf8 COLLATE utf8_general_ci COMMENT '调价说明',
  `version` int DEFAULT NULL COMMENT '调整版本',
  `op_uid` bigint DEFAULT NULL COMMENT '操作员UID',
  `op_ip` char(16) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL COMMENT '操作员IP',
  `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`) USING BTREE,
  KEY `sku_code_index` (`sku_code`) USING BTREE COMMENT '商品sku_code索引',
  KEY `shop_id_sku_code_index` (`shop_id`,`sku_code`) USING BTREE COMMENT '唯一索引'
) ENGINE=InnoDB AUTO_INCREMENT=380 DEFAULT CHARSET=utf8 COMMENT='商品价格历史记录';

-- ----------------------------
-- Table structure for sku_property
-- ----------------------------
DROP TABLE IF EXISTS `sku_property`;
CREATE TABLE `sku_property` (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `code` char(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '商品唯一编号',
  `price` decimal(32,16) DEFAULT NULL COMMENT '商品当前价格',
  `name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '商品名称',
  `desc` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci COMMENT '商品描述',
  `production` varchar(1024) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '生产企业',
  `supplier` varchar(1024) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '供应商',
  `category` int DEFAULT NULL COMMENT '商品类别',
  `title` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '商品标题',
  `sub_title` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '商品副标题',
  `color` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '商品颜色',
  `color_code` int DEFAULT NULL COMMENT '商品颜色代码',
  `specification` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '商品规格',
  `desc_link` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '商品介绍链接',
  `state` tinyint DEFAULT '0' COMMENT '商品状态，0-有效，1-无效，2-锁定',
  `create_time` datetime NOT NULL COMMENT '创建时间',
  `update_time` datetime NOT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`) USING BTREE,
  KEY `sku_code_index` (`code`) USING BTREE COMMENT '商品sku索引',
  KEY `sku_name_index` (`name`) USING BTREE COMMENT '商品名索引'
) ENGINE=InnoDB AUTO_INCREMENT=295 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='商品详情属性表';

SET FOREIGN_KEY_CHECKS = 1;

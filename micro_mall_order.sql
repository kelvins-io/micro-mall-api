/*
 Navicat Premium Data Transfer

 Source Server         : MySQL-本地
 Source Server Type    : MySQL
 Source Server Version : 80021
 Source Host           : localhost:3306
 Source Schema         : micro_mall_order

 Target Server Type    : MySQL
 Target Server Version : 80021
 File Encoding         : 65001

 Date: 18/07/2021 13:24:56
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for order
-- ----------------------------
DROP TABLE IF EXISTS `order`;
CREATE TABLE `order` (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '自增ID',
  `tx_code` char(40) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '交易号',
  `order_code` char(40) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '订单code',
  `uid` bigint NOT NULL COMMENT '用户UID',
  `order_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '下单时间',
  `description` text CHARACTER SET utf8 COLLATE utf8_general_ci COMMENT '订单描述',
  `client_ip` char(16) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL COMMENT '客户端IP',
  `device_code` varchar(512) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL COMMENT '客户端设备code',
  `shop_id` bigint NOT NULL COMMENT '门店ID',
  `state` tinyint NOT NULL DEFAULT '0' COMMENT '订单状态，0-有效，1-锁定中，2-无效',
  `pay_expire` datetime NOT NULL COMMENT '支付有效期，默认30分钟内有效',
  `pay_state` tinyint NOT NULL DEFAULT '0' COMMENT '支付状态，0-未支付，1-支付中，2-支付失败，3-已支付，4-支付过期取消',
  `amount` int DEFAULT NULL COMMENT '订单关联商品数量',
  `money` decimal(48,4) NOT NULL DEFAULT '0.0000' COMMENT '订单总金额',
  `coin_type` tinyint DEFAULT '0' COMMENT ' 订单币种，0-CNY，1-USD',
  `logistics_delivery_id` int DEFAULT NULL COMMENT '物流投递ID',
  `inventory_verify` tinyint DEFAULT '0' COMMENT '库存核实，0-未核实，1-核实',
  `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间',
  PRIMARY KEY (`id`) USING BTREE,
  KEY `uid_index` (`uid`) USING BTREE COMMENT '用户UID索引',
  KEY `shop_id_index` (`shop_id`) USING BTREE COMMENT '店铺ID索引',
  KEY `create_time_index` (`create_time`) USING BTREE COMMENT '订单创建时间',
  KEY `pay_expire_index` (`pay_expire`) USING BTREE COMMENT '订单支付过期时间',
  KEY `pay_state_order_code_index` (`pay_state`,`order_code`) USING BTREE COMMENT '支付状态-订单号',
  KEY `state,order_code_index` (`state`,`order_code`) USING BTREE COMMENT '订单状态-订单号',
  KEY `inventory_verify_order_code` (`inventory_verify`,`order_code`) USING BTREE COMMENT '订单库存核实-订单号',
  KEY `order_code_index` (`order_code`) USING BTREE COMMENT '订单code索引',
  KEY `tx_code_order_code_index` (`tx_code`,`order_code`) USING BTREE COMMENT '交易号订单号唯一索引'
) ENGINE=InnoDB AUTO_INCREMENT=1000 DEFAULT CHARSET=utf8 COMMENT='订单表';

-- ----------------------------
-- Table structure for order_estimate
-- ----------------------------
DROP TABLE IF EXISTS `order_estimate`;
CREATE TABLE `order_estimate` (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '自增ID',
  `estimate_code` char(40) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '评论code',
  `sku_code` char(40) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '商品sku',
  `order_code` char(40) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '订单code',
  `uid` bigint DEFAULT NULL COMMENT '用户uid',
  `shop_id` bigint DEFAULT NULL COMMENT '店铺ID',
  `content` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci COMMENT '内容',
  `star` int DEFAULT NULL COMMENT '星级',
  `state` tinyint DEFAULT '0' COMMENT '状态，0-有效',
  `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE KEY `estimate_code_shop_id` (`estimate_code`,`shop_id`) USING BTREE COMMENT '评论code-店铺ID',
  KEY `uid_index` (`uid`) USING BTREE COMMENT '用户ID',
  KEY `shop_id_index` (`shop_id`) USING BTREE COMMENT '店铺ID',
  KEY `order_code_index` (`order_code`) USING BTREE COMMENT '订单code'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='订单评论表';

-- ----------------------------
-- Table structure for order_scene_shop
-- ----------------------------
DROP TABLE IF EXISTS `order_scene_shop`;
CREATE TABLE `order_scene_shop` (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '自增ID',
  `order_code` char(40) DEFAULT NULL COMMENT '订单code',
  `shop_id` bigint DEFAULT NULL COMMENT '店铺ID',
  `shop_name` varchar(512) DEFAULT NULL COMMENT '店铺名',
  `shop_area_code` varchar(255) DEFAULT NULL COMMENT '店铺区域code',
  `shop_address` text COMMENT '店铺地址',
  `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1000 DEFAULT CHARSET=utf8 COMMENT='订单店铺信息';

-- ----------------------------
-- Table structure for order_sku
-- ----------------------------
DROP TABLE IF EXISTS `order_sku`;
CREATE TABLE `order_sku` (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '自增ID',
  `order_code` char(64) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '对应订单code',
  `shop_id` bigint NOT NULL COMMENT '店铺ID',
  `sku_code` char(64) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '商品sku',
  `price` decimal(32,16) NOT NULL DEFAULT '0.0000000000000000' COMMENT '商品单价',
  `amount` int NOT NULL COMMENT '商品数量',
  `name` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL COMMENT '商品名称',
  `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间',
  PRIMARY KEY (`id`) USING BTREE,
  KEY `name_index` (`name`) USING BTREE COMMENT '商品名称索引',
  KEY `shop_id_index` (`shop_id`) USING BTREE COMMENT '店铺索引',
  KEY `shop_order_code_index` (`shop_id`,`order_code`) USING BTREE COMMENT '店铺-订单code',
  KEY `order_code_index` (`order_code`) USING BTREE COMMENT '订单code索引'
) ENGINE=InnoDB AUTO_INCREMENT=1000 DEFAULT CHARSET=utf8 COMMENT='订单商品明细';

SET FOREIGN_KEY_CHECKS = 1;

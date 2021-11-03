/*
 Navicat Premium Data Transfer

 Source Server         : 本地MySQL
 Source Server Type    : MySQL
 Source Server Version : 80026
 Source Host           : localhost:3306
 Source Schema         : micro_mall_pay

 Target Server Type    : MySQL
 Target Server Version : 80026
 File Encoding         : 65001

 Date: 03/11/2021 04:43:44
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for account
-- ----------------------------
DROP TABLE IF EXISTS `account`;
CREATE TABLE `account` (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '自增ID',
  `account_code` char(50) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '账户主键',
  `owner` char(36) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '账户所有者',
  `balance` decimal(64,4) DEFAULT NULL COMMENT '账户余额',
  `coin_type` tinyint NOT NULL DEFAULT '0' COMMENT '币种类型，0-rmb，1-usdt',
  `coin_desc` varchar(64) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL COMMENT '币种描述',
  `state` tinyint DEFAULT NULL COMMENT '状态，1无效，2锁定，3正常',
  `account_type` tinyint NOT NULL COMMENT '账户类型，1-个人账户，2-公司账户，3-系统账户',
  `last_tx_id` char(60) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '99' COMMENT '最后一次事务ID',
  `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE KEY `account_index` (`owner`,`account_type`,`coin_type`) USING BTREE COMMENT '账户索引',
  KEY `create_time_index` (`create_time`) USING BTREE COMMENT '创建时间'
) ENGINE=InnoDB AUTO_INCREMENT=1000 DEFAULT CHARSET=utf8mb3 COMMENT='账户表';

-- ----------------------------
-- Table structure for pay_record
-- ----------------------------
DROP TABLE IF EXISTS `pay_record`;
CREATE TABLE `pay_record` (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '自增ID',
  `tx_id` char(40) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '批次交易号',
  `out_trade_no` char(40) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '外部商户订单号',
  `notify_url` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL COMMENT '交易结果通知地址',
  `description` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL COMMENT '交易描述',
  `merchant` char(40) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '交易商户ID',
  `attach` text CHARACTER SET utf8 COLLATE utf8_general_ci COMMENT '交易留言',
  `user` char(40) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '交易用户ID',
  `amount` decimal(64,4) NOT NULL COMMENT '交易数量',
  `coin_type` tinyint NOT NULL DEFAULT '0' COMMENT '交易币种，0-cny,1-usd',
  `reduction` decimal(64,4) DEFAULT NULL COMMENT '满减优惠',
  `pay_type` tinyint NOT NULL COMMENT '交易类型，1入账，2退款',
  `pay_state` tinyint DEFAULT NULL COMMENT '支付状态，0-未支付，1-支付中，2-支付失败，3-支付成功',
  `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE KEY `out_trade_no_index` (`out_trade_no`) USING BTREE COMMENT '外部商户单号',
  KEY `merchant_index` (`merchant`) USING BTREE COMMENT '外部商户ID',
  KEY `user_index` (`user`) USING BTREE COMMENT '外部用户ID',
  KEY `tx_id_index` (`tx_id`) USING BTREE COMMENT '批次交易号'
) ENGINE=InnoDB AUTO_INCREMENT=1000 DEFAULT CHARSET=utf8mb3 COMMENT='支付记录';

-- ----------------------------
-- Table structure for transaction
-- ----------------------------
DROP TABLE IF EXISTS `transaction`;
CREATE TABLE `transaction` (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '交易ID',
  `from_account_code` char(36) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '0' COMMENT '转出账户ID',
  `from_balance` decimal(64,4) DEFAULT '0.0000' COMMENT '转出后账户余额',
  `to_account_code` char(36) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '0' COMMENT '转入账户ID',
  `to_balance` decimal(64,4) DEFAULT NULL COMMENT '转入后账户余额',
  `amount` decimal(64,4) DEFAULT NULL COMMENT '交易金额',
  `meta` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '转账说明',
  `scene` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '支付场景',
  `op_uid` bigint NOT NULL COMMENT '操作用户UID',
  `op_ip` varchar(16) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '操作的IP',
  `tx_id` char(36) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '对应交易号',
  `fingerprint` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '防篡改指纹',
  `pay_type` tinyint DEFAULT '0' COMMENT '支付方式，0系统操作，1-银行卡，2-信用卡,3-支付宝,4-微信支付,5-京东支付',
  `pay_desc` varchar(36) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '支付方式描述',
  `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=1000 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='交易流水表';

SET FOREIGN_KEY_CHECKS = 1;

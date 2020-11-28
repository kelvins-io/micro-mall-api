/*
 Navicat Premium Data Transfer

 Source Server         : MySQL-本地
 Source Server Type    : MySQL
 Source Server Version : 80021
 Source Host           : localhost:3306
 Source Schema         : micro_mall_user

 Target Server Type    : MySQL
 Target Server Version : 80021
 File Encoding         : 65001

 Date: 27/11/2020 13:14:56
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for merchant
-- ----------------------------
DROP TABLE IF EXISTS `merchant`;
CREATE TABLE `merchant` (
  `merchant_id` bigint NOT NULL AUTO_INCREMENT COMMENT '商户号ID',
  `merchant_code` char(36) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '商户唯一code',
  `uid` bigint NOT NULL COMMENT '用户ID',
  `register_addr` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '注册地址',
  `health_card_no` char(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '健康证号',
  `identity` tinyint DEFAULT NULL COMMENT '身份属性，1-临时店员，2-正式店员，3-经理，4-店长',
  `state` tinyint DEFAULT NULL COMMENT '状态，0-未审核，1-审核中，2-审核不通过，3-已审核',
  `tax_card_no` char(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '纳税账户号',
  `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间',
  PRIMARY KEY (`merchant_id`) USING BTREE,
  UNIQUE KEY `uid_index` (`uid`) USING BTREE COMMENT '商户用户ID',
  KEY `merchant_code_index` (`merchant_code`) USING BTREE COMMENT '商户code唯一索引'
) ENGINE=InnoDB AUTO_INCREMENT=1082 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='商户属性表';

-- ----------------------------
-- Table structure for user
-- ----------------------------
DROP TABLE IF EXISTS `user`;
CREATE TABLE `user` (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '自增ID',
  `account_id` char(36) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '账户ID，全局唯一',
  `user_name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '用户名',
  `password` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '用户密码md5值',
  `password_salt` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '密码salt值',
  `sex` tinyint(1) DEFAULT NULL COMMENT '性别，1-男，2-女',
  `phone` char(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '手机号',
  `country_code` char(5) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '手机区号',
  `email` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '邮箱',
  `state` tinyint(1) NOT NULL DEFAULT '3' COMMENT '状态，0-未激活，1-审核中，2-审核未通过，3-已审核',
  `id_card_no` char(18) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '身份证号',
  `inviter` bigint DEFAULT NULL COMMENT '邀请人uid',
  `invite_code` char(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '邀请码',
  `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间',
  `contact_addr` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci COMMENT '联系地址',
  `age` int DEFAULT NULL COMMENT '年龄',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE KEY `account_id_index` (`account_id`) USING BTREE COMMENT '账户ID索引',
  UNIQUE KEY `country_code_phone_index` (`country_code`,`phone`) USING BTREE COMMENT '手机号索引',
  KEY `user_name_index` (`user_name`) USING BTREE COMMENT '用户名索引',
  KEY `email_index` (`email`) USING BTREE COMMENT '邮箱索引',
  KEY `id_card_no_index` (`id_card_no`) USING BTREE COMMENT '身份证号索引',
  KEY `invite_code_index` (`invite_code`) USING BTREE COMMENT '邀请码索引'
) ENGINE=InnoDB AUTO_INCREMENT=79292 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='用户信息表';

-- ----------------------------
-- Table structure for user_logistics_delivery
-- ----------------------------
DROP TABLE IF EXISTS `user_logistics_delivery`;
CREATE TABLE `user_logistics_delivery` (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '自增ID',
  `uid` bigint DEFAULT NULL COMMENT '用户ID',
  `delivery_user` varchar(512) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '交付人',
  `country_code` varchar(10) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '86' COMMENT '区号',
  `phone` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '手机号',
  `area` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '交付区域',
  `area_detailed` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci COMMENT '详细地址',
  `label` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci COMMENT '标签，多个以|分割开',
  `is_default` tinyint DEFAULT '0' COMMENT '是否为默认，1-默认',
  `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`) USING BTREE,
  KEY `uid_index` (`uid`) USING BTREE COMMENT '用户ID'
) ENGINE=InnoDB AUTO_INCREMENT=134 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='用户物流交付信息';

SET FOREIGN_KEY_CHECKS = 1;

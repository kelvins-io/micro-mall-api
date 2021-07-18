/*
 Navicat Premium Data Transfer

 Source Server         : MySQL-本地
 Source Server Type    : MySQL
 Source Server Version : 80021
 Source Host           : localhost:3306
 Source Schema         : micro_mall

 Target Server Type    : MySQL
 Target Server Version : 80021
 File Encoding         : 65001

 Date: 18/07/2021 13:25:29
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for config_kv_store
-- ----------------------------
DROP TABLE IF EXISTS `config_kv_store`;
CREATE TABLE `config_kv_store` (
  `id` int NOT NULL AUTO_INCREMENT COMMENT '主键',
  `config_key` varchar(255) NOT NULL COMMENT '配置键',
  `config_value` varchar(255) NOT NULL COMMENT '配置值',
  `prefix` varchar(255) NOT NULL COMMENT '配置前缀',
  `suffix` varchar(255) NOT NULL COMMENT '配置后缀',
  `status` tinyint NOT NULL DEFAULT '1' COMMENT '是否启用 1是 0否',
  `is_delete` tinyint NOT NULL DEFAULT '0' COMMENT '是否删除 1是 0否',
  `create_time` datetime DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_time` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `uq_config_key` (`config_key`)
) ENGINE=InnoDB AUTO_INCREMENT=35 DEFAULT CHARSET=utf8 COMMENT='参数配置';

-- ----------------------------
-- Table structure for verify_code_record
-- ----------------------------
DROP TABLE IF EXISTS `verify_code_record`;
CREATE TABLE `verify_code_record` (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '自增id',
  `uid` bigint NOT NULL COMMENT '用户UID',
  `business_type` tinyint DEFAULT NULL COMMENT '验证类型，1-注册登录，2-购买商品',
  `verify_code` char(6) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '验证码',
  `expire` int DEFAULT NULL COMMENT '过期时间unix',
  `country_code` char(5) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '验证码下发手机国际码',
  `phone` char(11) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '验证码下发手机号',
  `email` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '验证码下发邮箱',
  `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间',
  PRIMARY KEY (`id`),
  KEY `country_code_phone_index` (`country_code`,`phone`) USING BTREE COMMENT '手机号索引',
  KEY `email_index` (`email`) USING BTREE COMMENT '邮箱索引',
  KEY `verify_code_index` (`verify_code`) USING BTREE COMMENT '验证码索引'
) ENGINE=InnoDB AUTO_INCREMENT=1968 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='验证码记录表';

SET FOREIGN_KEY_CHECKS = 1;

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

SET FOREIGN_KEY_CHECKS = 1;

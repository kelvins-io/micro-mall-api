/*
 Navicat Premium Data Transfer

 Source Server         : MySQL-本地
 Source Server Type    : MySQL
 Source Server Version : 80021
 Source Host           : localhost:3306
 Source Schema         : micro_mall_trolley

 Target Server Type    : MySQL
 Target Server Version : 80021
 File Encoding         : 65001

 Date: 18/07/2021 13:24:26
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for user_trolley
-- ----------------------------
DROP TABLE IF EXISTS `user_trolley`;
CREATE TABLE `user_trolley` (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '自增ID',
  `uid` bigint NOT NULL COMMENT '用户ID',
  `shop_id` bigint NOT NULL COMMENT '店铺ID',
  `sku_code` char(40) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '商品sku',
  `count` int NOT NULL DEFAULT '1' COMMENT '商品数量',
  `join_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '加入时间',
  `selected` tinyint(1) DEFAULT '1' COMMENT '是否选中，1-未选中，2-选中',
  `state` tinyint DEFAULT '1' COMMENT '状态，1-有效，2-移除',
  `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`) USING BTREE,
  KEY `sku_code_index` (`sku_code`) USING BTREE COMMENT 'sku索引',
  KEY `shop_id_sku_index` (`shop_id`,`sku_code`) USING BTREE COMMENT '店铺=sku索引',
  KEY `shop_id_sku_uid_index` (`uid`,`shop_id`,`sku_code`) USING BTREE COMMENT '唯一索引'
) ENGINE=InnoDB AUTO_INCREMENT=23 DEFAULT CHARSET=utf8 COMMENT='购物车';

SET FOREIGN_KEY_CHECKS = 1;

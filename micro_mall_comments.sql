/*
 Navicat Premium Data Transfer

 Source Server         : MySQL-本地
 Source Server Type    : MySQL
 Source Server Version : 80021
 Source Host           : localhost:3306
 Source Schema         : micro_mall_comments

 Target Server Type    : MySQL
 Target Server Version : 80021
 File Encoding         : 65001

 Date: 18/07/2021 13:25:15
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for comments_logistics
-- ----------------------------
DROP TABLE IF EXISTS `comments_logistics`;
CREATE TABLE `comments_logistics` (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '自增ID',
  `logistics_code` char(40) NOT NULL COMMENT '物流单号',
  `uid` bigint NOT NULL COMMENT '用户ID',
  `fedex_pack_star` tinyint DEFAULT NULL COMMENT '物流包装星级',
  `fedex_pack_content` text COMMENT '物流包装评价',
  `delivery_speed_star` tinyint DEFAULT NULL COMMENT '送货速度星级',
  `delivery_speed_content` text COMMENT '送货速度评价',
  `delivery_service` tinyint DEFAULT NULL COMMENT '配送服务星级',
  `delivery_service_content` text CHARACTER SET utf8 COLLATE utf8_general_ci COMMENT '配送服务评价',
  `comment` text COMMENT '评价内容',
  `anonymity` tinyint NOT NULL DEFAULT '0' COMMENT '是否匿名，0-匿名，1-实名',
  `state` tinyint NOT NULL DEFAULT '0' COMMENT '状态，0-有效，1-无效',
  `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`),
  KEY `logistics_code_index` (`logistics_code`) USING BTREE COMMENT '物流单号',
  KEY `uid_index` (`uid`) USING BTREE COMMENT '用户ID',
  KEY `create_time_index` (`create_time`) USING BTREE COMMENT '创建时间索引'
) ENGINE=InnoDB AUTO_INCREMENT=105 DEFAULT CHARSET=utf8 COMMENT='物流评论';

-- ----------------------------
-- Table structure for comments_order
-- ----------------------------
DROP TABLE IF EXISTS `comments_order`;
CREATE TABLE `comments_order` (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '自增ID',
  `comment_code` char(40) NOT NULL COMMENT '评论code',
  `uid` bigint DEFAULT NULL COMMENT '用户ID',
  `shop_id` bigint DEFAULT NULL COMMENT '店铺id',
  `order_code` char(40) DEFAULT NULL COMMENT '订单code',
  `star` int DEFAULT NULL COMMENT '星级',
  `content` text COMMENT '评价类容',
  `img_list` text COMMENT '评价图片or视频',
  `anonymity` tinyint NOT NULL DEFAULT '0' COMMENT '是否匿名，0-匿名，1-实名',
  `state` tinyint NOT NULL DEFAULT '0' COMMENT '状态，0-有效，1-无效',
  `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`),
  KEY `shop_id_index` (`shop_id`) USING BTREE COMMENT '店铺索引',
  KEY `uid_index` (`uid`) USING BTREE COMMENT '用户索引',
  KEY `comment_code_index` (`comment_code`) USING BTREE COMMENT '评论code',
  KEY `create_time_index` (`create_time`) USING BTREE COMMENT '创建时间'
) ENGINE=InnoDB AUTO_INCREMENT=105 DEFAULT CHARSET=utf8 COMMENT='订单评论';

-- ----------------------------
-- Table structure for comments_tags
-- ----------------------------
DROP TABLE IF EXISTS `comments_tags`;
CREATE TABLE `comments_tags` (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '自增ID',
  `tag_code` char(40) NOT NULL COMMENT '标签code',
  `classification_major` varchar(255) DEFAULT NULL COMMENT '主要分类',
  `classification_medium` varchar(255) DEFAULT NULL COMMENT '中等分类',
  `classification_minor` varchar(255) DEFAULT NULL COMMENT '次要分类',
  `content` text COMMENT '内容',
  `state` tinyint NOT NULL DEFAULT '0' COMMENT '状态，0-有效，1-无效',
  `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`),
  KEY `tag_code_index` (`tag_code`) USING BTREE COMMENT 'tag索引',
  KEY `classification_major_index` (`classification_major`) USING BTREE COMMENT '主要分类',
  KEY `classification_medium_index` (`classification_medium`) USING BTREE COMMENT '次要分类'
) ENGINE=InnoDB AUTO_INCREMENT=118 DEFAULT CHARSET=utf8 COMMENT='标签表';

SET FOREIGN_KEY_CHECKS = 1;

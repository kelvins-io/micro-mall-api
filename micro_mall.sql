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

 Date: 18/10/2020 20:56:19
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for account
-- ----------------------------
DROP TABLE IF EXISTS `account`;
CREATE TABLE `account` (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '自增ID',
  `account_code` char(50) NOT NULL COMMENT '账户主键',
  `owner` char(36) NOT NULL COMMENT '账户所有者',
  `balance` decimal(32,16) DEFAULT NULL COMMENT '账户余额',
  `coin_type` tinyint NOT NULL DEFAULT '1' COMMENT '币种类型，1-rmb，2-usdt',
  `coin_desc` varchar(64) DEFAULT NULL COMMENT '币种描述',
  `state` tinyint DEFAULT NULL COMMENT '状态，1无效，2锁定，3正常',
  `account_type` tinyint NOT NULL COMMENT '账户类型，1-个人账户，2-公司账户，3-系统账户',
  `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE KEY `account_index` (`owner`,`account_type`,`coin_type`) USING BTREE COMMENT '账户索引',
  KEY `create_time_index` (`create_time`) USING BTREE COMMENT '创建时间'
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8 COMMENT='账户表';

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
  PRIMARY KEY (`merchant_id`),
  UNIQUE KEY `uid_index` (`uid`) USING BTREE COMMENT '商户用户ID',
  KEY `merchant_code_index` (`merchant_code`) USING BTREE COMMENT '商户code唯一索引'
) ENGINE=InnoDB AUTO_INCREMENT=1028 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='商户属性表';

-- ----------------------------
-- Table structure for order
-- ----------------------------
DROP TABLE IF EXISTS `order`;
CREATE TABLE `order` (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '自增ID',
  `tx_code` char(40) NOT NULL COMMENT '交易号',
  `order_code` char(40) NOT NULL COMMENT '订单code',
  `uid` bigint NOT NULL COMMENT '用户UID',
  `order_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '下单时间',
  `description` varchar(255) DEFAULT NULL COMMENT '订单描述',
  `client_ip` char(16) DEFAULT NULL COMMENT '客户端IP',
  `device_code` varchar(512) DEFAULT NULL COMMENT '客户端设备code',
  `shop_id` bigint NOT NULL COMMENT '门店ID',
  `shop_name` varchar(255) NOT NULL COMMENT '门店名称',
  `shop_area_code` varchar(255) DEFAULT NULL COMMENT '门店区域编号',
  `shop_address` text COMMENT '门店地址',
  `state` tinyint NOT NULL DEFAULT '0' COMMENT '订单状态，0-有效，1-锁定中，2-无效',
  `pay_expire` datetime NOT NULL COMMENT '支付有效期，默认30分钟内有效',
  `pay_state` tinyint NOT NULL DEFAULT '0' COMMENT '支付状态，0-未支付，1-支付中，2-支付失败，3-已支付',
  `amount` int DEFAULT NULL COMMENT '订单关联商品数量',
  `total_amount` decimal(32,16) NOT NULL DEFAULT '0.0000000000000000' COMMENT '订单总金额',
  `coin_type` tinyint DEFAULT '1' COMMENT ' 订单币种，1-CNY，2-USD',
  `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `order_code_index` (`order_code`) COMMENT '订单code索引',
  UNIQUE KEY `tx_code_order_code_index` (`tx_code`,`order_code`) USING BTREE COMMENT '交易号订单号唯一索引',
  KEY `uid_index` (`uid`) USING BTREE COMMENT '用户UID索引',
  KEY `shop_id_index` (`shop_id`) USING BTREE COMMENT '店铺ID索引',
  KEY `shop_name_index` (`shop_name`) USING BTREE COMMENT '店铺名索引',
  KEY `description_index` (`description`) USING BTREE COMMENT '订单描述索引',
  KEY `order_time_index` (`order_time`) USING BTREE COMMENT '订单创建时间索引'
) ENGINE=InnoDB AUTO_INCREMENT=142 DEFAULT CHARSET=utf8 COMMENT='订单表';

-- ----------------------------
-- Table structure for order_sku
-- ----------------------------
DROP TABLE IF EXISTS `order_sku`;
CREATE TABLE `order_sku` (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '自增ID',
  `order_code` char(64) NOT NULL COMMENT '对应订单code',
  `shop_id` bigint NOT NULL COMMENT '店铺ID',
  `sku_code` char(64) NOT NULL COMMENT '商品sku',
  `price` decimal(32,16) NOT NULL DEFAULT '0.0000000000000000' COMMENT '商品单价',
  `amount` int NOT NULL COMMENT '商品数量',
  `name` varchar(255) DEFAULT NULL COMMENT '商品名称',
  `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `order_unique` (`order_code`,`shop_id`,`sku_code`) USING BTREE COMMENT '唯一索引',
  KEY `name_index` (`name`) USING BTREE COMMENT '商品名称索引',
  KEY `shop_id_index` (`shop_id`) USING BTREE COMMENT '店铺索引',
  KEY `sku_code_index` (`sku_code`) USING BTREE COMMENT '商品sku索引'
) ENGINE=InnoDB AUTO_INCREMENT=148 DEFAULT CHARSET=utf8 COMMENT='订单商品明细';

-- ----------------------------
-- Table structure for pay_record
-- ----------------------------
DROP TABLE IF EXISTS `pay_record`;
CREATE TABLE `pay_record` (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '自增ID',
  `tx_id` char(40) NOT NULL COMMENT '批次交易号',
  `out_trade_no` char(40) NOT NULL COMMENT '外部商户订单号',
  `time_expire` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '支付过期时间',
  `notify_url` varchar(255) DEFAULT NULL COMMENT '交易结果通知地址',
  `description` varchar(255) DEFAULT NULL COMMENT '交易描述',
  `merchant` char(40) NOT NULL COMMENT '交易商户ID',
  `attach` text COMMENT '交易留言',
  `user` char(40) NOT NULL COMMENT '交易用户ID',
  `amount` decimal(32,16) NOT NULL COMMENT '交易数量',
  `coin_type` tinyint NOT NULL DEFAULT '0' COMMENT '交易币种，0-cny,1-usd',
  `reduction` decimal(32,16) DEFAULT NULL COMMENT '满减优惠',
  `pay_type` tinyint NOT NULL COMMENT '交易类型，1入账，2退款',
  `pay_state` tinyint DEFAULT NULL COMMENT '支付状态，0-未支付，1-支付中，2-支付失败，3-支付成功',
  `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间',
  PRIMARY KEY (`id`),
  KEY `out_trade_no_index` (`out_trade_no`) USING BTREE COMMENT '外部商户单号',
  KEY `merchant_index` (`merchant`) USING BTREE COMMENT '外部商户ID',
  KEY `user_index` (`user`) USING BTREE COMMENT '外部用户ID',
  KEY `tx_id_index` (`tx_id`) USING BTREE COMMENT '批次交易号'
) ENGINE=InnoDB AUTO_INCREMENT=156 DEFAULT CHARSET=utf8 COMMENT='支付记录';

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
  `state` tinyint NOT NULL DEFAULT '0' COMMENT '状态，0-未审核，1-审核不通过，2-审核通过',
  `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间',
  PRIMARY KEY (`shop_id`),
  UNIQUE KEY `legal_person_nick_name_index` (`legal_person`,`nick_name`) USING BTREE COMMENT '法人店铺名索引',
  UNIQUE KEY `shop_code_index` (`shop_code`) USING BTREE COMMENT '店铺唯一code',
  KEY `legal_person_index` (`legal_person`) USING BTREE COMMENT '店铺f法人'
) ENGINE=InnoDB AUTO_INCREMENT=30046 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='店铺主体登记表';

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
  `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `sku_code_index` (`sku_code`) USING BTREE COMMENT '商品编码code',
  UNIQUE KEY `sku_code_shop_id_index` (`sku_code`,`shop_id`) USING BTREE COMMENT '商品code，店铺ID索引',
  KEY `shop_id_index` (`shop_id`) USING BTREE COMMENT '店铺ID索引'
) ENGINE=InnoDB AUTO_INCREMENT=121 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='商品库存表';

-- ----------------------------
-- Table structure for sku_price_history
-- ----------------------------
DROP TABLE IF EXISTS `sku_price_history`;
CREATE TABLE `sku_price_history` (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '自增ID',
  `shop_id` bigint NOT NULL COMMENT '调价的店铺id',
  `sku_code` char(40) NOT NULL COMMENT '商品sku_code',
  `price` decimal(32,16) NOT NULL COMMENT '商品价格',
  `tsp` int NOT NULL COMMENT '价格变化时的时间戳',
  `reason` text COMMENT '调价说明',
  `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `op_uid` bigint DEFAULT NULL COMMENT '操作员UID',
  `op_ip` char(16) DEFAULT NULL COMMENT '操作员IP',
  PRIMARY KEY (`id`),
  UNIQUE KEY `shop_id_sku_code_index` (`shop_id`,`sku_code`) USING BTREE COMMENT '唯一索引',
  KEY `sku_code_index` (`sku_code`) USING BTREE COMMENT '商品sku_code索引',
  KEY `timestamp_index` (`tsp`) USING BTREE COMMENT '调价时间索引'
) ENGINE=InnoDB AUTO_INCREMENT=308 DEFAULT CHARSET=utf8 COMMENT='商品价格历史记录';

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
  PRIMARY KEY (`id`),
  KEY `sku_code_index` (`code`) USING BTREE COMMENT '商品sku索引',
  KEY `sku_name_index` (`name`) USING BTREE COMMENT '商品名索引'
) ENGINE=InnoDB AUTO_INCREMENT=221 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='商品详情属性表';

-- ----------------------------
-- Table structure for transaction
-- ----------------------------
DROP TABLE IF EXISTS `transaction`;
CREATE TABLE `transaction` (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '交易ID',
  `from_account_code` char(36) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '0' COMMENT '转出账户ID',
  `from_balance` decimal(32,16) DEFAULT '0.0000000000000000' COMMENT '转出后账户余额',
  `to_account_code` char(36) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '0' COMMENT '转入账户ID',
  `to_balance` decimal(32,16) DEFAULT NULL COMMENT '转入后账户余额',
  `amount` decimal(32,16) DEFAULT NULL COMMENT '交易金额',
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
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=22 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='交易流水表';

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
  `phone` char(11) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '手机号',
  `country_code` char(5) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '手机区号',
  `email` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '邮箱',
  `state` tinyint(1) DEFAULT NULL COMMENT '状态，0-未激活，1-审核中，2-审核未通过，3-已审核',
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
  UNIQUE KEY `id_card_no_index` (`id_card_no`) USING BTREE COMMENT '身份证号索引',
  KEY `user_name_index` (`user_name`) USING BTREE COMMENT '用户名索引',
  KEY `email_index` (`email`) USING BTREE COMMENT '邮箱索引'
) ENGINE=InnoDB AUTO_INCREMENT=10032 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='用户信息表';

-- ----------------------------
-- Table structure for user_trolley
-- ----------------------------
DROP TABLE IF EXISTS `user_trolley`;
CREATE TABLE `user_trolley` (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '自增ID',
  `uid` bigint NOT NULL COMMENT '用户ID',
  `shop_id` bigint NOT NULL COMMENT '店铺ID',
  `sku_code` char(40) NOT NULL COMMENT '商品sku',
  `count` int NOT NULL DEFAULT '1' COMMENT '商品数量',
  `join_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '加入时间',
  `selected` tinyint(1) DEFAULT '1' COMMENT '是否选中，1-未选中，2-选中',
  `state` tinyint DEFAULT '1' COMMENT '状态，1-有效，2-移除',
  `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`),
  KEY `sku_code_index` (`sku_code`) USING BTREE COMMENT 'sku索引',
  KEY `shop_id_sku_index` (`shop_id`,`sku_code`) USING BTREE COMMENT '店铺=sku索引',
  KEY `shop_id_sku_uid_index` (`uid`,`shop_id`,`sku_code`) USING BTREE COMMENT '唯一索引'
) ENGINE=InnoDB AUTO_INCREMENT=16 DEFAULT CHARSET=utf8 COMMENT='购物车';

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
) ENGINE=InnoDB AUTO_INCREMENT=1064 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='验证码记录表';

SET FOREIGN_KEY_CHECKS = 1;

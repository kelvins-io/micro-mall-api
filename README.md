# micro-mall-api

#### 介绍
微商城-api

#### 软件架构
go + xorm + mysql + redis + rabbitmq + grpc + etcd + rocketmq + protobuf    
服务间通信采用gRPC，服务注册采用etcd，普通消息事件采用rabbitmq，交易订单消息事件采用rocketmq   
protobuf v3   

#### 模块分类
接入层（gateway，BFF）   
https://gitee.com/cristiane/micro-mall-api   

用户服务   
https://gitee.com/cristiane/micro-mall-users   
https://gitee.com/cristiane/micro-mall-users-proto   

店铺服务   
https://gitee.com/cristiane/micro-mall-shop   
https://gitee.com/cristiane/micro-mall-shop-proto   

商品服务   
https://gitee.com/cristiane/micro-mall-sku   
https://gitee.com/cristiane/micro-mall-shop-proto   

购物车服务   
https://gitee.com/cristiane/micro-mall-trolley   
https://gitee.com/cristiane/micro-mall-trolley-proto   

订单服务   
https://gitee.com/cristiane/micro-mall-order   
https://gitee.com/cristiane/micro-mall-order-proto   

支付服务   
https://gitee.com/cristiane/micro-mall-pay   
https://gitee.com/cristiane/micro-mall-pay-proto   

物流系统   
https://gitee.com/cristiane/micro-mall-logistics   
https://gitee.com/cristiane/micro-mall-logistics-proto   

评价系统   
https://gitee.com/cristiane/micro-mall-estimate   
https://gitee.com/cristiane/micro-mall-estimate-proto   

////依赖   
web模板：https://gitee.com/cristiane/web_gin_template   
gRPC应用模板：https://gitee.com/kelvins-io/kelvins-template   
脚手架：https://gitee.com/kelvins-io/kelvins


#### 安装教程
开发环境地址：  http://127.0.0.1:52001/   
监控地址：http://localhost:52002/debug/pprof/   


返回码code：   
4003 	 token无效   
4009 	 商户认证资料已存在   
4011 	 商户未提交过店铺认证资料   
4007 	 用户密码错误   
4012 	 商品唯一code已存在系统   
4013 	 商品唯一code不存在   
200 	 ok   
500 	 服务器出错   
400 	 请求参数错误   
4006 	 用户已存在   
50002 	 验证码为空   
4015 	 店铺ID已存在   
50004 	 验证码过期   
50000 	 Duplicate entry   
4008 	 商户未提交过认证资料   
4001 	 ID为空   
4002 	 token为空   
4004 	 token过期   
4005 	 用户不存在   
50003 	 验证码无效    
4010 	 店铺认证资料已存在   
50001 	 邮件发送错误   
4014 	 店铺ID不存在   
4016     邀请码不存在   
接口列表：   
首页   
GET    /               
返回body   

```
{
	"code": 200,
	"data": "Welcome to micro-mall-api",
	"msg": "ok"
}
```


在线检测          
GET    /ping                      

返回body   

```
{
	"code": 200,
	"data": "2020-09-11T21:55:28.873726+08:00",
	"msg": "ok"
}
```

发送验证码   
POST   /api/v1/common/verify_code/send   
请求参数：   

参数 | 含义 |  类型 | 备注  
---|------|------|---
country_code |国际码 | string | 86
phone |手机号 | string | 11位手机号
business_type |业务类型 | int | 1注册，2登录，3修改或重置密码
receive_email |接收验证码邮箱 | string | xxxx@xx.com

返回body：   
```
{"code":200,"data":"ok","msg":"ok"}
```

注册用户   
POST   /api/v1/common/register   
请求参数：   

参数 | 含义 |  类型 | 备注  
---|------|------|---
user_name |用户名 | string | 
password |手机号 | string | 11位手机号
sex |性别 | int | 1男，2女
email |接收验证码邮箱 | string | xxxx@xx.com
country_code |国际码 | string | 
phone |手机号 | string | 11位手机号
verify_code |验证码 | string | 6位验证码
id_card_no | 身份证号 | string | 选填
invite_code |邀请码 | string | xxx

返回body：   

```
{"code":200,"data":{"invite_code":"46e4eabbf000065"},"msg":"ok"}
```

验证码登陆   
POST   /api/v1/common/login/verify_code   
请求参数：   

参数 | 含义 |  类型 | 备注  
---|------|------|---
country_code |国际码 | string | 
phone |手机号 | string | 11位手机号
verify_code |验证码 | string | 6位验证码

返回body：   

```
{"code":200,"data":"token","msg":"ok"}
```

密码登陆   
POST   /api/v1/common/login/pwd   
请求参数：   

参数 | 含义 |  类型 | 备注  
---|------|------|---
country_code |国际码 | string | 
phone |手机号 | string | 11位手机号
password | 密码 | string | 可传md5值

返回body：   
```
{"code":200,"data":{},"msg":"ok"}
```

重置用户密码   
PUT    /api/v1/user/password/reset   
header token   
请求参数：   

参数 | 含义 |  类型 | 备注  
---|------|------|---
verify_code |验证码 | string | 6位验证码
password | 密码 | string | 可传md5值

返回body：   
```
{"code":200,"data":"token","msg":"ok"}
```

获取用户信息   
GET    /api/v1/user/user_info     
header token   

返回body： 
```
{
	"code": 200,
	"data": {
		"id": 10009,
		"account_id": "ae23bab6-c31b-4f61-ad5e-2521a9a4917d",
		"user_name": "杨强",
		"sex": 1,
		"phone": "18319430520",
		"country_code": "86",
		"email": "1225807604@qq.com",
		"state": 0,
		"id_card_no": "524348787893748475",
		"inviter": 0,
		"invite_code": "46a576fc4000065",
		"contact_addr": "深圳市南山区南头街道桃园路南贸市场三栋208",
		"age": 34,
		"create_time": "2020-09-04 19:10:07",
		"update_time": "2020-09-06 12:10:15"
	},
	"msg": "ok"
}
```

提交商户认证资料   
PUT    /api/v1/user/merchants/material   
header token   
请求参数：    

参数 | 含义 |  类型 | 备注  
---|------|------|---
operation_type | 操作类型 | int | 0-创建，1-更新，2删除，3-审核
register_addr | 注册地址 | string | 
health_card_no | 从业人员健康证 | string | 11-29位字符
register_addr | 注册地址 | string | 真实注册地址
identity | 身份标识 | int | 身份属性，1-临时店员，2-正式店员，3-经理，4-店长
tax_card_no | 纳税人证号 | string | 大于16位字符

返回body： 

```
{"code":200,"data":{"merchant_id":111},"msg":"ok"}

```

添加商品到购物车   
PUT    /api/v1/user/trolley/sku/join   
header token   
请求参数：   

参数 | 含义 |  类型 | 备注  
---|------|------|---
sku_code | string | int | 商品唯一sku_code
shop_id | 店铺ID | int | 商品所属店铺ID
count | 数量 | int | 大于0
time | 加入时间 | string | 2020-09-05 13:25:43
selected | 是否选中 | bool | true,false

返回body： 
```
{"code":200,"data":"ok","msg":"ok"}
```

从购物车中移除商品   
DELETE /api/v1/user/trolley/sku/remove   
header token   
请求参数：   

参数 | 含义 |  类型 | 备注  
---|------|------|---
sku_code | string | int | 商品唯一sku_code
shop_id | 店铺ID | int | 商品所属店铺ID

返回body： 
```
{"code":200,"data":"ok","msg":"ok"}
```

获取用户购物车   
GET    /api/v1/user/trolley/sku/list   
header token   
返回body： 
```
{"code":200,"data":{"list":[{"sku_code":"df1a9633-b060-4682-9502-bc934f89392b","shop_id":29914,"count":534252790,"time":"2020-09-11 23:01:25","selected":true}]},"msg":"ok"}
```

商户申请店铺   
POST   /api/v1/shop_business/shop/apply   
header token   
请求参数：   

参数 | 含义 |  类型 | 备注  
---|------|------|---
operation_type | 操作类型 | int | 0-创建，1-更新，2删除，3-审核
shop_id | 店铺ID | int | 商品所属店铺ID
nick_name | 店铺简称 | string | 不能为空
full_name | 店铺完整名称 | string | 不能为空
register_addr | 店铺地址 | string | 不能为空
merchant_id | 店铺商户ID（法人） | int | 请先申请商户
business_addr | 业务地址 | string | 具体地址
business_license | 商业许可证号 | string | 
tax_card_no | 纳税号 | string | 
business_desc | 经营业务描述 | string | 尽可能详细可以加快审核
social_credit_code | 同一信用代码 | string | 不能为空
organization_code | 组织结构代码 | string | 不能为空

返回body： 

```
{"code":200,"data":{"shop_id":111},"msg":"ok"}

```

店铺质押保证金   
PUT    /api/v1/shop_business/shop/pledge   
暂时未实现   


商品上架   
POST   /api/v1/sku_business/sku/put_away   
header token   
请求参数：   

参数 | 含义 |  类型 | 备注  
---|------|------|---
operation_type | 操作类型 | int | 0-创建，1-更新，2删除，3-审核
sku_code | 商品sku | string | 商品唯一code
name | 商品名称 | string | 不能为空
price | 价格 | string | 数字字符串
title | 商品标题 | string | 不能为空
sub_title | 商品副标题 | string | 
desc | 商品描述 | string | 
production | 生产商 | string | 不能为空
supplier | 供应商 | string | 
category | 商品分类 | string | 不能为空
color | 颜色 | string | 如白色，红色
color_code | 颜色代码 | int | 细分颜色代码
specification | 商品规格 | string | 产品等级描述
desc_link | 商品描述链接 | string | 不能为空
state | 状态 | int | 商品所属店铺ID
amount | 上架数量 | int | 大于0
shop_id | 店铺ID | int | 商品所属店铺ID

返回body： 

```
{"code":200,"data":{},"msg":"ok"}

```

补充商品扩展信息   
PUT    /api/v1/sku_business/sku/supplement   
header token   
请求参数：   

参数 | 含义 |  类型 | 备注  
---|------|------|---
operation_type | 操作类型 | int | 0-创建，1-更新，2删除，3-审核
sku_code | 商品sku | string | 商品唯一code
shop_id | 店铺ID | int | 商品所属店铺ID
name | 商品名称 | string | 不能为空
size | 商品尺寸 | string | 描述商品大小，187cm x 112cm
shape | 商品形状 | string | 袋装，箱装，
production_country | 产地 | string | 不能为空
production_date | 生成日期 | string | 如2020-12-11 09:09
shelf_life | 有效期 | string | 描述过期截止时间

返回body： 

```
{"code":200,"data":{},"msg":"ok"}

```

获取店铺上架商品列表   
GET    /api/v1/sku_business/sku/list   
header token   
返回body： 

```
{"code":200,"data":{"list":[{"sku_code":"df1a9633-b060-4682-9502-bc934f89392b","shop_id":29914,"count":534252790,"time":"2020-09-11 23:01:25","selected":true}]},"msg":"ok"}

```


#### 使用说明
配置数据库sql, rabbitmq, redis，邮件，etcd   
请先将根目录micro-mall.sql导入数据库创建相应的表   

有问题联系：565608463@qq.com   
# micro-mall-api

#### 介绍
微商城-api，基于gRPC构建的微服务商城，包含用户，商品，购物车，订单，支付共计16个微服务并通过micro-mall-api聚合

#### 项目问题交流
QQ群：578859618 （micro-mall-api交流群）  
![avatar](./交流群.JPG)
邮件：1225807604@qq.com   

#### 软件架构
micro-mall系列需要etcd集群，集群有问题无法运行任何一个项目，请先搭建好！   
gin + xorm + mysql + redis + rabbitmq + grpc + etcd + MongoDB + protobuf + prometheus     
服务间通信采用gRPC（protobuf v3 ），服务注册/发现采用etcd，消息事件采用rabbitmq， 搜索采用elasticsearch 
    
用户鉴权   
jwt

存储说明：   
MySQL 主存储，事务处理   
MongoDB：备份仓库，商品价格变化，商品详情（如./sku_property_ex.json），历史记录   
Redis：数据缓存，消息事件结果，用户在线状态，分布式锁支持   
rabbitMQ：消息事件中转站，订阅   
ETCD：配置项，微服务注册，发现，分布式锁支持   

监控说明：   
pprof接口   
elastic_metrics接口  
prometheus_metrics接口   

架构示意图：   
![avatar](./微商城系统架构设计.png)

### 如何构建开发环境
micro-mall-xxx系列服务，希望开发者有中高级go后端开发经验，了解电商业务，mysql redis MQ使用经验     
你需要安装golang并配置golang开发环境   
然后看看下面的环节      

#### 服务注册说明
由于micro-mall系列服务是通过etcd来注册的，所以是需要etcd集群的，搭建步骤参考本仓库的etcd集群部署文档    
export ETCDV3_SERVER_URLS=http://10.211.55.12:2379,http://10.211.55.13:2379  #地址仅供示范     
export ETCDCTL_API=3   

#### 都有哪些服务
micro-mall-api   接入层   
micro-mall-users  用户服务   
micro-mall-users-consumer  用户事件消费者   
micro-mall-order   订单服务   
micro-mall-order-cron   订单定时任务   
micro-mall-order-consumer   订单事件消费者   
micro-mall-shop   店铺服务   
micro-mall-trolley   购物车服务   
micro-mall-sku   商品库服务   
micro-mall-sku-cron   商品库定时任务   
micro-mall-pay   支付服务   
micro-mall-pay-consumer  支付事件消费者   
micro-mall-comments   评论服务   
micro-mall-logistics   物流服务   
micro-mall-search    搜索服务   
micro-mall-search-cron   搜索定时任务   

#### 关于go mod
请各位一定配置go proxy   
GOPROXY="https://goproxy.baidu.com,https://goproxy.io,direct"   

#### 克隆仓库
将这些服务（目前共16个服务以及它们依赖的proto仓库，在模块分类环节可以了解到）clone到本地   

#### 服务启动端口说明
除了micro-mall-api服务需要在/etc/app.ini中配置端口外，其余需要占用tcp端口的服务都是在运行时自动生成随机端口号并注册到etcd集群中   

#### 运行环境说明
export GO_ENV=dev   #本地开发环境   
可选值：dev,test,release,prod   

#### 都有哪些依赖
部分依赖文件安装需要科学上网环境，演示安装步骤都是Mac环境下(同时也建议开发者使用Linux或Mac环境)，Windows请参考安装或自行Google安装   
go 1.13.15   
goland or vscode   
mysql，redis，rabbitmq（如果用redis做MQ则不需要），etcd集群环境，MongoDB，elasticsearch       
protoc   安装方法如下   
```
wget https://github.com/google/protobuf/releases/download/v3.5.1/protobuf-all-3.5.1.zip
unzip protobuf-all-3.5.1.zip
cd protobuf-3.5.1/
./configure
make
make install
# 如果报错请执行ldconfig
```
go get -u google.golang.org/grpc   
go get -u github.com/golang/protobuf/protoc-gen-go   
mv /usr/local/go/path/bin/protoc-gen-go /usr/local/go/bin/（Mac）   
go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway   
go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger   
go get -u github.com/jteeuwen/go-bindata/...   
go get github.com/elazarl/go-bindata-assetfs/...   
python 2.7或3.5   

#### 数据库设计
micro-mall-系列采用分库存储，各服务拥有独立的数据库，独立的缓存，独立的事件消息   
将本仓库根目录下的micro_mall*.sql（共9个sql文件） 导入同名数据库中初始化表结构   
MongoDB需要注册一个admin用户和数据库micro_mall_sku      

#### host配置   
为什么需要配置host ？   
：因为调用服务时会以：域名+从etcd获取的端口号 作为endpoint进行tcp拨号，如 tcp://micro-mall-users:50987 为了让grpc能正常解析域名对应的IP地址，所以需要配置dns，当然此处是有优化空间的（比如注册服务时同时将启动服务的机器IP地址和端口一起注册到etcd，拨号连接时就可以用ip+端口进行拨号了，但是这种方案也是有弊端的）   

127.0.0.1  micro-mall-users   
127.0.0.1  micro-mall-order   
127.0.0.1  micro-mall-shop   
127.0.0.1  micro-mall-sku   
127.0.0.1  micro-mall-pay   
127.0.0.1  micro-mall-comments   
127.0.0.1  micro-mall-logistics    
127.0.0.1 micro-mall-trolley    
127.0.0.1  micro-mall-search    
    ....   
配置host是为了让服务根据服务名找到IP地址   
Windows下编辑C:\Windows\System32\drivers\etc   
Mac，Linux下编辑 /etc/hosts   
当然也可以使用https://github.com/oldj/SwitchHosts/releases    
这个工具可视化配置   

#### pb.gw 代码生成
看看下面说明

##### 为什么要生成pb,gw代码
因为项目使用了gRPC+protobuf，grpc-gateway协议转换中间件，swagger文档托管中间件    

##### 有哪些仓库需要生成
大部分列出的服务基本上都需要，基本上micro-mall-xxx-proto就是micro-mall-xxx系列服务的依赖proto仓库   
那么就需要在这些仓库下面生成proto仓库的代码文件   

##### 如何生成
生成pb，gw代码需要本地安装protoc,protoc-gen-go，grpc-gateway 可以参考https://segmentfault.com/a/1190000013339403 grpc系列文章   
生成proto.pb.go代码时请将https://gitee.com/kelvins-io/common clone下来并放到gopath路径中(GOPATH/src/gitee.com/kelvins-io/)   
生成proto.pb.go代码方式：   
在micro-mall-xxx根目录执行python genpb.py .../micro-xxx-proto   
例如：为micro-mall-api项目添加micro-users-proto的pb代码   
cd $GOPATH   
cd src/gitee.com/cristiane/ #没有则创建   
git clone https://gitee.com/cristiane/micro-mall-users-proto.git   
git clone https://gitee.com/cristiane/micro-mall-api.git   
cd micro-mall-api   
没有Python环境的需要安装python，Mac的话自带python   
python genpb.py ../micro-mall-users-proto   
没有报错，且检查proto目录是否创建micro-mall-users-proto目录   

#### 配置文件
看看下面环节   

##### 需要配置什么
基本上就是日志路径，MySQL，Redis，rabbitmq，elasticsearch，email这些      

##### 如何配置
在需要运行的项目根目录下etc/app.ini更改自己开发环境的配置info（可以参考默认提供的）      


#### 运行项目
在需要项目根目录运行go mod vendor安装依赖（不要运行go mod tidy）      
go run main.go   


### 模块分类
接入层（gateway，BFF）   
https://gitee.com/cristiane/micro-mall-api   

用户服务   
https://gitee.com/cristiane/micro-mall-users   
https://gitee.com/cristiane/micro-mall-users-proto   
https://gitee.com/cristiane/micro-mall-users-consumer   

店铺服务   
https://gitee.com/cristiane/micro-mall-shop   
https://gitee.com/cristiane/micro-mall-shop-proto   

商品服务   
https://gitee.com/cristiane/micro-mall-sku   
https://gitee.com/cristiane/micro-mall-sku-proto   
https://gitee.com/cristiane/micro-mall-sku-cron   

购物车服务   
https://gitee.com/cristiane/micro-mall-trolley   
https://gitee.com/cristiane/micro-mall-trolley-proto   

订单服务   
https://gitee.com/cristiane/micro-mall-order   
https://gitee.com/cristiane/micro-mall-order-proto   
https://gitee.com/cristiane/micro-mall-order-cron   
https://gitee.com/cristiane/micro-mall-order-consumer   

支付服务   
https://gitee.com/cristiane/micro-mall-pay   
https://gitee.com/cristiane/micro-mall-pay-proto   
https://gitee.com/cristiane/micro-mall-pay-consumer   

物流系统   
https://gitee.com/cristiane/micro-mall-logistics   
https://gitee.com/cristiane/micro-mall-logistics-proto   

搜索服务   
https://gitee.com/cristiane/micro-mall-search   
https://gitee.com/cristiane/micro-mall-search-cron   
https://gitee.com/cristiane/micro-mall-search-proto   

评论服务   
https://gitee.com/cristiane/micro-mall-comments   
https://gitee.com/cristiane/micro-mall-comments-proto   

### 赞赏   
作者是来自一个贫困山区的孩子，靠社会资助上了大学，工作后不忘回馈社会
花费N个周末和夜晚才写了这么一个开源项目，其中孤独不言而喻，如果你喜欢项目的话，给作者买一盒膏药缓解下腰间盘突出吧   
![avatar](./微信赞赏码.JPG)
支付宝
![avatar](./支付宝赞赏码.JPG)
### 接口文档
开发环境地址：  http://127.0.0.1:52001/   
监控地址：   
pprof：http://localhost:52002/debug/pprof/   
Elastic：http://localhost:52002/debug/vars   
Prometheus：http://localhost:52002/metrics   

返回码code：   
200 		 ok   
500 		 服务器出错   
4001 		 ID为空   
4002 		 token为空   
4007 		 用户密码错误   
4012 		 商品唯一code已存在系统   
4015 		 店铺ID已存在   
4016 		 邀请码不存在   
600000 		 金额格式解析错误   
4003 		 token无效   
50003 		 验证码无效   
600002 		 用户账户被锁定   
600004 		 商户账户不存在   
400 		 请求参数错误   
50004 		 验证码过期   
600010 		 事务执行失败   
4005 		 用户不存在   
50002 		 验证码为空   
4008 		 商户未提交过认证资料   
4011 		 商户未提交过店铺认证资料   
600001 		 用户余额不足   
600003 		 用户账户不存在   
4009 		 商户认证资料已存在   
4014 		 店铺ID不存在   
600011 		 交易号不存在   
4004 		 token过期   
4006 		 用户已存在   
50001 		 邮件发送错误   
50000 		 Duplicate entry   
4010 		 店铺认证资料已存在   
4013 		 商品唯一code不存在   
50005 		 商品库存不够   
600005 		 商户账户被锁定    


接口列表：   
####【说明】post请求没指明content-type的接口表单和json都支持   
1 首页   
GET    /               
返回body   

```
{
	"code": 200,
	"data": "Welcome to micro-mall-api",
	"msg": "ok"
}
```


2 在线检测          
GET    /ping                      

返回body   

```
{
	"code": 200,
	"data": "2020-09-11T21:55:28.873726+08:00",
	"msg": "ok"
}
```

3 发送验证码   
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

4 注册用户   
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

5 验证码登陆   
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

6 密码登陆   
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

7 重置用户密码   
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

8 获取用户信息   
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

8.1 列举用户    
GET /api/v1/user/user_info/list?page_size=500&page_num=1&token=xxx   
header token   

返回body：   
```json
{
	"code": 200,
	"data": {
		"user_info_list": [{
			"country_code": "86",
			"phone": "01606450639"
		}, {
			"country_code": "86",
			"phone": "01606450640"
		}, {
			"country_code": "86",
			"phone": "01606450641"
		}]
	},
	"msg": "ok"
}
```

9 提交商户认证资料   
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

10 添加商品到购物车   
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

11 从购物车中移除商品   
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

12  获取用户购物车   
GET    /api/v1/user/trolley/sku/list   
header token   
返回body： 
```
{"code":200,"data":{"list":[{"sku_code":"df1a9633-b060-4682-9502-bc934f89392b","shop_id":29914,"count":534252790,"time":"2020-09-11 23:01:25","selected":true}]},"msg":"ok"}
```

13  商户申请店铺   
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

14  店铺质押保证金   
PUT    /api/v1/shop_business/shop/pledge   
暂时未实现   


15  商品上架   
POST   /api/v1/sku_business/sku/put_away   
header token   
请求参数：   

参数 | 含义 |  类型 | 备注  
---|------|------|---
operation_type | 操作类型 | int | 0-创建，1-更新，2删除，3-审核，4增加库存
sku_code | 商品sku | string | 商品唯一code
name | 商品名称 | string | 不能为空
price | 价格 | string | 数字字符串
title | 商品标题 | string | 不能为空
sub_title | 商品副标题 | string | 
desc | 商品描述 | string | 商品描述
production | 生产商 | string | 不能为空
supplier | 供应商 | string | 供应商
category | 商品分类 | string | 不能为空
color | 颜色 | string | 如白色，红色
color_code | 颜色代码 | int | 细分颜色代码
specification | 商品规格 | string | 产品等级描述
desc_link | 商品描述链接 | string | 不能为空
state | 状态 | int | 状态
amount | 上架数量 | int | 大于0
shop_id | 店铺ID | int | 商品所属店铺ID

#### operation_type等于4时，参数只需要shop_id,sku_code,amount

返回body： 

```
{"code":200,"data":{},"msg":"ok"}

```

16   补充商品扩展信息   
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

17   获取店铺上架商品列表   
GET    /api/v1/sku_business/sku/list   
header token   
返回body： 

```
{"code":200,"data":{"list":[{"sku_code":"df1a9633-b060-4682-9502-bc934f89392b","shop_id":29914,"count":534252790,"time":"2020-09-11 23:01:25","selected":true}]},"msg":"ok"}

```

18   添加商品到购物车   
post /user/trolley/sku/join   
header token  
请求参数：   

参数 | 含义 |  类型 | 备注  
---|------|------|---
sku_code | 商品sku | string | 商品唯一code
shop_id | 店铺ID | int | 商品所属店铺ID
count | 数量 | int | 最少为1
time | 时间 | string | 如2020-12-11 09:09
selected | 是否选中 | bool | true表示选中，false表示未选中

返回body   
```

```

19   从购物车移除商品
get /user/trolley/sku/remove   
header token  
请求参数：   

参数 | 含义 |  类型 | 备注  
---|------|------|---
sku_code | 商品sku | string | 商品唯一code
shop_id | 店铺ID | int | 商品所属店铺ID

返回body   
```

```

20   获取用户购物车列表   
get /user/trolley/sku/list   
header token  

返回body   
```

```

21  创建订单   
post /user/order/create   
header token  
请求参数：   

```
{
	"uid": 100098,
	"client_ip": "127.0.0.1",
	"description": "网络购物",
	"device_id": "iphone-x",
	"detail": [{
		"shop_id": 29912,
		"coin_type": 1,
		"goods": [{
			"sku_code": "38d9d035-00ed-40ed-aa83-abe90b59c055",
			"price": "184.32",
			"amount": 5,
			"name": "盼盼铜锣烧"
		}, {
			"sku_code": "b363e9f4-3bae-4103-86a6-5e4b83b70303",
			"price": "184.32",
			"amount": 5,
			"name": "盼盼铜锣烧"
		}],
		"scene_info": {
			"store_info": {
				"id": 29912,
				"name": "良品铺子京东旗舰店1",
				"area_code": "深圳",
				"address": "深圳市宝安区"
			}
		}
	}, {
		"shop_id": 29911,
		"coin_type": 0,
		"goods": [{
			"sku_code": "b882a5c9-564a-4912-a5d4-ce77de71577c",
			"price": "184.32",
			"amount": 5,
			"name": "盼盼铜锣烧-2"
		}],
		"scene_info": {
			"store_info": {
				"id": 29911,
				"name": "良品铺子京东旗舰店-2",
				"area_code": "广州",
				"address": "广州市海珠区"
			}
		}
	}]
}
```
返回body   
```

```

22  订单支付   
post /user/order/trade   
header token  
请求参数：   

参数 | 含义 |  类型 | 备注  
---|------|------|---
tx_code | 订单交易号 | string | 不能为空

返回body   
```

```

23  申请物流
post /user/logistics/apply   
header token  
请求参数：   

参数 | 含义 |  类型 | 备注  
---|------|------|---
out_trade_no | 订单交易号 | string | 不能为空
courier | 承运人 | string | 如，微商城快递
courier_type | 承运类型 | string | 0-普通，1-铁路，2-空运，3-加急，4-延迟
receive_type | 收件类型 | string | 0-普通，1-本人接收，2-代理接收
send_user | 发送方 | string | 李云龙
send_addr | 发送地址 | string | 河北省邯郸市东方路198号怡和豪庭10栋
send_phone | 发送方联系方式 | string | 如，13683749374
send_time | 发送时间 | string | 如，2020-10-10 10:10:10
receive_user | 接收方 | string | 赵富贵
receive_addr | 接收方地址 | string | 四川省成都市武侯区98号
receive_phone | 接收方联系方式 | string | 如，0838-10182827
goods | 需要承运的货物 | string | 如，下面序列化后的值

goods示范
```
[{
	"sku_code": "2131d-f111-45e1-b68a-d602c2f0f1b3",
	"name": "怡宝矿泉水",
	"kind": "饮用水",
	"count": 98
}]
```

24  用户配置收货地址   
post json   
/api/v1/user/setting/address

请求body   
```
{
	"id": 101,
	"delivery_user": "张6丰",
	"mobile_phone": "15501707785",
	"area": "广东省广州市",
	"detailed_area": "上海路步行街111号",
	"label": ["公司", "住宅", "生活"],
	"is_default": true,
	"operation_type": 0
}
```

返回body   
```
{"code":200,"data":"","msg":"ok"}
```

24 用户查询收货地址列表   
get    
/api/v1/user/setting/address?delivery_id=xx  
  
返回body   
```
{
	"code": 200,
	"data": [{
		"id": 105,
		"delivery_user": "张6丰",
		"mobile_phone": "15501707785",
		"area": "广东省广州市",
		"detailed_area": "上海路步行街111号",
		"label": ["公司", "住宅", "生活"],
		"is_default": true
	}, {
		"id": 106,
		"delivery_user": "张6丰",
		"mobile_phone": "15501707785",
		"area": "广东省广州市",
		"detailed_area": "上海路步行街111号",
		"label": ["公司", "住宅", "生活"],
		"is_default": false
	}, {
		"id": 107,
		"delivery_user": "张6丰",
		"mobile_phone": "15501707785",
		"area": "广东省广州市",
		"detailed_area": "上海路步行街111号",
		"label": ["公司", "住宅", "生活"],
		"is_default": false
	}],
	"msg": "ok"
}
```

25  商品库存搜索   
get /search/sku_inventory?keyword=剃须刀   

返回body   
```
{
	"code": 200,
	"data": [{
		"info": {
			"sku_code": "2cf90b0f-4fc3-49cc-8df7-de8942c1f128",
			"name": "飞科剃须刀",
			"price": "699.0000000000000000",
			"title": "飞科剃须刀",
			"sub_title": "飞科(FLYCO) 男士电动剃须刀 全身水洗干湿双剃刮胡刀 浮动贴面三刀头 FS372，减价促销",
			"desc": "飞科(FLYCO) 男士电动剃须刀 全身水洗干湿双剃刮胡刀 浮动贴面三刀头 FS372",
			"production": "上海飞科用具有限公司",
			"supplier": "飞科京东旗舰店",
			"category": 11010,
			"color": "黑色",
			"color_code": 199,
			"specification": "旋转式剃须刀，三刀头，刀头进口",
			"desc_link": "https://item.jd.com/1750531.html",
			"state": 1,
			"version": 1,
			"amount": 100
		},
		"score": 13.864214
	}, {
		"info": {
			"sku_code": "9475963f-317f-4a9a-b513-9dcc76da2672",
			"name": "飞利浦剃须刀",
			"price": "599.0000000000000000",
			"title": "飞利浦剃须刀",
			"sub_title": "飞利浦剃须刀，减价促销",
			"desc": "飞利浦（PHILIPS）男士电动剃须刀多功能理容剃胡刀刮胡刀礼盒装（配鬓角 鼻毛修剪器）S5082/61",
			"production": "广州飞利浦科技有限公司",
			"supplier": "飞利浦微商城旗舰店",
			"category": 11010,
			"color": "黑色",
			"color_code": 199,
			"specification": "旋转式剃须刀，三刀头，刀头进口",
			"desc_link": "https://item.jd.com/1750531.html",
			"state": 1,
			"version": 1,
			"amount": 100
		},
		"score": 13.504881
	}],
	"msg": "ok"
}
```

26 店铺搜索   
get /search/shop?keyword=交个朋友   

返回body   
```
{
	"code": 200,
	"data": [{
		"info": {
			"shop_id": 30063,
			"merchant_id": 1037,
			"nick_name": "广州市交个朋友科技有限公司",
			"full_name": "广州市交个朋友科技有限公司",
			"register_addr": "深圳市宝安区宝源二区73栋111号",
			"business_addr": "深圳市宝安区宝源二区73栋111号",
			"business_license": "qX2MkznWrlvO4sIp7",
			"tax_card_no": "qX2MkznWrlvO4sIp7",
			"business_desc": "qX2MkznWrlvO4sIp7",
			"social_credit_code": "qX2MkznWrlvO4sIp7",
			"organization_code": "qX2MkznWrlvO4sIp7",
			"shop_code": "7e0be82d-6fdd-4a89-a228-d6f3378b82da"
		},
		"score": 3.157851
	}, {
		"info": {
			"shop_id": 30066,
			"merchant_id": 1037,
			"nick_name": "广州市交个朋友科技有限公司（南京分公司）",
			"full_name": "广州市交个朋友科技有限公司（南京分公司）",
			"register_addr": "深圳市宝安区宝源二区73栋111号",
			"business_addr": "深圳市宝安区宝源二区73栋111号",
			"business_license": "qX2MkznWrlvO4sIp7",
			"tax_card_no": "qX2MkznWrlvO4sIp7",
			"business_desc": "qX2MkznWrlvO4sIp7",
			"social_credit_code": "qX2MkznWrlvO4sIp7",
			"organization_code": "qX2MkznWrlvO4sIp7",
			"shop_code": "07964e6c-16f9-4e3d-8212-bb336e9ad75a"
		},
		"score": 2.7211857
	}, {
		"info": {
			"shop_id": 30065,
			"merchant_id": 1037,
			"nick_name": "广州市交个朋友科技有限公司（北京分公司）",
			"full_name": "广州市交个朋友科技有限公司（北京分公司）",
			"register_addr": "深圳市宝安区宝源二区73栋111号",
			"business_addr": "深圳市宝安区宝源二区73栋111号",
			"business_license": "qX2MkznWrlvO4sIp7",
			"tax_card_no": "qX2MkznWrlvO4sIp7",
			"business_desc": "qX2MkznWrlvO4sIp7",
			"social_credit_code": "qX2MkznWrlvO4sIp7",
			"organization_code": "qX2MkznWrlvO4sIp7",
			"shop_code": "2b464f04-c360-463f-8cbf-44e9adfde329"
		},
		"score": 2.7211857
	}, {
		"info": {
			"shop_id": 30064,
			"merchant_id": 1037,
			"nick_name": "广州市交个朋友科技有限公司（深圳分公司）",
			"full_name": "广州市交个朋友科技有限公司（深圳分公司）",
			"register_addr": "深圳市宝安区宝源二区73栋111号",
			"business_addr": "深圳市宝安区宝源二区73栋111号",
			"business_license": "qX2MkznWrlvO4sIp7",
			"tax_card_no": "qX2MkznWrlvO4sIp7",
			"business_desc": "qX2MkznWrlvO4sIp7",
			"social_credit_code": "qX2MkznWrlvO4sIp7",
			"organization_code": "qX2MkznWrlvO4sIp7",
			"shop_code": "5c2ae7b2-113f-491d-b4b5-6f81089aec6a"
		},
		"score": 2.7211857
	}, {
		"info": {
			"shop_id": 30067,
			"merchant_id": 1037,
			"nick_name": "福建赚它一个亿科技有限公司",
			"full_name": "福建赚它一个亿科技有限公司",
			"register_addr": "深圳市宝安区宝源二区73栋111号",
			"business_addr": "深圳市宝安区宝源二区73栋111号",
			"business_license": "qX2MkznWrlvO4sIp7",
			"tax_card_no": "qX2MkznWrlvO4sIp7",
			"business_desc": "qX2MkznWrlvO4sIp7",
			"social_credit_code": "qX2MkznWrlvO4sIp7",
			"organization_code": "qX2MkznWrlvO4sIp7",
			"shop_code": "d55eaf41-88a8-4d73-a324-70fcc8f64e2d"
		},
		"score": 0.6836133
	}],
	"msg": "ok"
}
```

获取店铺订单报告   
post  /user/order/report   
header token   

参数 | 含义 |  类型 | 备注  
---|------|------|---
shop_id | 店铺ID | int | 不能为空
start_time | 统计开始时间 | string | 如，2019-11-22 08:46:41
end_time | 统计结束时间 | string | 如，2020-12-04 18:46:41
page_size | 分页大小 | int | 500，最小1
page_num | 分页号 | int | 最小1

返回body   
```
{
	"code": 200,
	"data": {
		"report_file_path": "http://localhost:52001/static/order-report-30070-1606124289.xlsx"
	},
	"msg": "ok"
}
```
report_file_path 报告的下载地址   

用户账户充值   
post  /user/account/charge   
header token   

参数 | 含义 |  类型 | 备注  
---|------|------|---
account_type | 账户类型 | int | 0-个人账户，2-公司账户，3-系统账户
amount | 金额 | string | 如，99.09
coin_type | 币种 | int | 如，0-RMB，1-USDT
device_code | 设备 | string | vivo NEX
device_platform | 平台 | string | Android

返回body   
```
{"code":200,"data":"","msg":"ok"}
```

订单评价   
post/json  /user/comments/order/create   
header token   

```
{
	"anonymity": false,
	"OrderCommentsInfo": {
		"shop_id": 30072,
		"order_code": "000be2f2-489c-4e19-8e2a-731319c98aab",
		"star": 1,
		"content": "经常在这家店购买，没毛病",
		"img_list": ["image1"],
		"comment_id": ""
	},
	"LogisticsCommentsInfo": {
		"logistics_code": "f7e7cf5c-ae54-46bc-a0b3-623f446be29f",
		"fedex_pack": 3,
		"fedex_pack_label": ["打包不结实"],
		"delivery_speed": 3,
		"delivery_speed_label": ["送货速度慢"],
		"delivery_service": 3,
		"delivery_service_label": ["配送服务不到位"],
		"comment": "配送人员没送到家门口"
	}
}
```

返回body   
```
{"code":200,"data":"","msg":"ok"}
```

获取店铺评价   
get  /user/comments/shop/list?shop_id=111    
返回body   

```
{
	"code": 200,
	"data": [{
		"shop_id": 30072,
		"order_code": "00038f56-7123-4af6-96b7-b7fceeb12415",
		"star": 1,
		"content": "商品很快就送到手里了，物美价廉",
		"img_list": ["image1"],
		"comment_id": "7e80704a-2731-44fb-9450-a0c8bbb68441"
	}],
	"msg": "ok"
}
```

修改评价标签    
post  /user/comments/tags/modify   
请求参数：   

参数 | 含义 |  类型 | 备注  
---|------|------|---
operation_type | 操作类型 | int | 0-新建，1-修改
tag_code | 标签ID | string | 修改时需要，如，0099acd
classification_major | 主要分类 | string | 如，商品
classification_medium | 次要分类 | string | 如，仓库
classification_minor | 细致分类 | string | 如，配送
content | 平台 | string | 标签内容

返回body   
```
{"code":200,"data":"","msg":"ok"}
```

获取标签列表   
get /user/comments/tags/list   
请求参数：   

参数 | 含义 |  类型 | 备注  
---|------|------|---
tag_code | 标签ID | string | 如，0099acd
classification_major | 主要分类 | string | 如，商品
classification_medium | 次要分类 | string | 如，仓库

返回body   

```
{
	"code": 200,
	"data": [{
		"tag_code": "1221d8e7-ab5f-42da-831d-455dd5a023d3",
		"classification_major": "店铺",
		"classification_medium": "商品",
		"classification_minor": "包装",
		"content": "商品保证破损"
	}],
	"msg": "ok"
}
```
  
### 共同开源
1 请star收藏项目   
2 提出issue   
3 fork项目   
4 clone fork后的项目到本地开发   
5 提交pr   

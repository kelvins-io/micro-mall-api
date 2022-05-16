# micro-mall-api
[![micro-mall-api](logo.png)](https://gitee.com/cristiane)

#### 介绍
微商城-api，基于gRPC构建的微服务商城，包含用户，商品，店铺，购物车，订单，支付，物流，评论，搜索共计22个微服务并通过micro-mall-api聚合

#### 框架，库依赖
kelvins框架支持：https://gitee.com/kelvins-io/kelvins   
模板生成，服务治理：https://gitee.com/kelvins-io/kelvins-tools   
g2cache多级缓存库支持：https://gitee.com/kelvins-io/g2cache   

#### 项目问题交流
QQ群：578859618 （micro-mall-api交流群）  
![avatar](./交流群.JPG)   

#### 如何赞助
任何商业用途（包括卖课平台，二次开发商用，毕业设计等）都需要作者授权，否则视为侵权追究法律责任。   
开源项目的发展离不开大家的鼓励和赞赏，扫描下方二维码鼓励一下吧   
![avatar](./微信赞赏码.JPG)

支付宝   
![avatar](./支付宝赞赏码.JPG)

#### 赞助商列表
昵称 | 赞助金额 |  时间 | 留言  
---|------|------|---
雨化田 | 100元 | 2021-1-25 | 一起加入
thomas | 100元 | 2021-2-18 | 指导
皮卡猪 | 250元 | 2021-2-20 | 支持大佬
*抹 | 20 | 2021-3-19 | 资金有限，支持下
*康 | 66.66 | 2021-4-15 | 加油
Bleem | -goland正版license | 2021-4-18 | 落地验证码限制以及缓存实施
Christible | 66.00 | 2021-4-26 | 大神，膜拜。资金有限
剑峰 | 50.00 | 2021-5-10 | 支持下
mu | 100.00 | 2021-6-9 | 意思意思
osc | -200.00 | 2021-7-9 | 落地docker构建方案 
这个杀手有点冷 | 150.00 | 2021-7-11 | 很好的一个项目
pick刘 | 50.00 | 2021-8-3 | 有料
IT詹天佑 | 88.00 | 2021-8-15 | 喜欢这一点代码
*浩 | 20.00 | 2021-8-25 | 请喝一杯奶茶
Doyle | 100.00 | 2021-8-31 | 一点小意思
星辰大海 | 200.00 | 2021-9-5 | 很好的项目
黔驴技穷 | 100.00 | 2021-10-2 | 问题咨询
_天行健_ | 100.00 | 2021-10-15 | 部署咨询
东正 | 20.00 | 2021-10-20 | 咨询解答
Jackson | 20.00 | 2021-10-28 | 部署咨询
農民 | 20.00 | 2021-11-2 | 一杯茶颜悦色
ps | 50.00 | 2021-11-6 | 项目咨询
Mark | 66.00 | 2021-11-14 | etcd搭建协助
Micky | 20.00 | 2021-11-18 | 赞助
井 | 50.00 | 2021-11-19 | 赞助
农民GG | 500.00 | 2021-11-24 | 支持开源项目
Z*k | 20.00 | 2021-11-27 | 支持开源项目
*左 | 20.00 | 2021-11-29 | 赞助
奈何桥 | 300.00 | 2021-12-2 | 运行部署
p神 | 1000.00 | 2021-12-5 | 个人赞助
曹大 | 1000.00 | 2021-12-11 | 个人赞助
*铭 | 50.00 | 2021-12-19 | 赞助
番茄可乐 | 100.00 | 2022-3-4 | 赞助
东窗事发 | 80.00 | 2022-3-10 | 运行支持
*(男) | 20.00 | 2022-3-17 | 喜欢
沈** | 1000.00 | 2022-3-29 | 私有化
（*） | 66.66 | 2022-4-3 | 多谢开源项目
（*生） | 20.00 | 2022-4-4 | 进交流群
A | 66.00 | 2022-4-23 | 直播打赏
*李 | 20.00 | 2022-5-7 | 进群
m*e | 1.00 | 2022-5-7 | 进群
Jenkins | 60.00 | 2022-5-7 | 远程协助
佟 | 50.00 | 2022-5-9 | 赞助


#### 软件架构
micro-mall系列需要etcd集群，集群有问题无法运行任何一个项目，请先搭建好！   
接入层micro-mall-api采用HTTP2（兼容HTTP1.1）对接客户端   
micro-mall-api基于gRPC对接其它服务   
gin + xorm + mysql + redis + rabbitmq + elasticsearch + grpc + etcd + MongoDB + protobuf + prometheus     
服务间通信采用gRPC（protobuf v3 ），服务注册/发现采用etcd，消息事件采用rabbitmq/redis， 搜索采用elasticsearch 
    
用户鉴权   
jwt

存储说明：   
MySQL 主存储，事务处理   
MongoDB：备份仓库，商品价格变化，商品详情（如./sku_property_ex.json），历史记录，搜索数据仓   
Redis：数据缓存，消息事件结果，用户在线状态，分布式锁支持   
rabbitMQ：消息事件中转站，订阅   
ETCD：配置项，微服务注册，发现，分布式锁支持   

监控说明：   
pprof接口   
prometheus_metrics接口   

架构示意图：   
![avatar](./微商城系统架构设计.png)

我的架构图都是在[processON在线画图，点击领7天会员福利](https://www.processon.com/i/5ac1db5ee4b0cf24e963bd80) 上面画的，有超多UML，软件架构，产品原型模板可以选择。

### 项目目录结构
```
.
├── batch-clone-project.sh  一键克隆本项目 shell
├── build-project-proto.sh  开发环境编译项目proto shell
├── build.sh  一键构建项目 shell
├── batch-run.sh  一键构建运行项目 shell
├── build-run-all.sh  一键构建运行micro-mall所有项目 shell
├── stop.sh  停止当前项目 shell
├── stop-all.sh  一键停止micro-mall所有项目 shell
├── clean-logs-all.sh  一键清理micro-mall所有目录下logs shell
├── git-push-all.sh  一键git push micro-mall所有项目（开发者专用） shell
├── Dockerfile  docker构建文件
├── docker-build-run.sh  docker构建运行 shell
├── docker-compose.yml  基础组件compose构建文件
├── docker-compose-build.yml  micro-mall系列项目compose构建文件
├── docker-etcd-build.sh  etcd集群构建 shell
├── golang-install.sh  centos环境golang安装 shell
├── docker-install.sh  centos环境docker安装 shell
├── etcd_cluster
│   └── docker-compose-etcd.yml etcd集群构建docker镜像
├── test-gateway-http.sh  运行client目录下测试用例 shell
├── LICENSE   授权文件
├── README.md   导航指南
├── app   
│   ├── app.go    app初始化
│   └── web.go    web服务初始化
├── client
│   ├── header.go   测试用例request header
│   ├── main
│   │   └── user_name.go    模拟注册的生成用户名
│   ├── micro-mall-api_test.go    接口测试用例
│   ├── url.go    接口测试URL
│   └── user_name.go    模拟生成用户名
├── config
│   ├── config.go   解析配置文件
│   └── setting
│       └── setting.go    系统配置setting
├── etc
│   ├── app.ini   app配置文件
│   └── app.ini.example   app配置文件example
├── genpb.py    pb,gw文件生成脚本
├── go.mod
├── go.sum
├── internal
│   ├── config
│   │   └── config.go   内部解析配置
│   ├── logging
│   │   └── log.go    启动日志
│   ├── metrics_mux
│   │   ├── elastic_metrics.go    elastic监控
│   │   ├── pprof.go    golang内置pprof监控
│   │   └── prometheus_metrics.go   prometheus监控
│   ├── setup
│   │   ├── mysql.go  mysql启动
│   │   ├── queue.go   queue启动
│   │   ├── redis.go  redis启动
│   │   └── server_mux.go  server监控mux入口
│   └── util
│       ├── mysql_callback
│       │   └── time_callback.go    mysql time callback
│       └── mysql_model
│           └── mysql_model.go    mysql mode 
├── logs  日志文件目录
├── main.go   main入口文件
├── micro-mall-api    main.go编译后文件
├── micro-mall-api.pid    micro-mall-api进程pid
├── micro-mall-db.yaml    映射db表结构到orm struct
├── micro_mall.sql    micro-mall数据库SQL
├── micro_mall_comments.sql   micro-mall-comments评论服务数据库SQL
├── micro_mall_logistics.sql    micro-mall-logistics物流服务数据库SQL
├── micro_mall_order.sql    micro-mall-order订单服务数据库SQL
├── micro_mall_pay.sql    micro-mall-pay支付服务数据库SQL
├── micro_mall_shop.sql   micro-mall-shop店铺服务数据库SQL
├── micro_mall_sku.sql    micro-mall-sku商品库服务数据库SQL
├── micro_mall_trolley.sql    micro-mall-trolley购物车服务数据库SQL
├── micro_mall_user.sql   micro-mall-users用户服务数据库SQL
├── middleware
│   ├── cors.go   cros跨域中间件
│   └── user.go   user验证中间件
├── model
│   ├── args
│   │   ├── const.go    请求参数常量
│   │   └── form.go   请求参数struct
│   └── mysql
│       └── table.go    MySQL表名常量
├── pkg
│   ├── app
│   │   ├── form.go   请求表单参数
│   │   └── response.go   http响应
│   ├── code
│   │   ├── code.go   响应码定义
│   │   ├── error.go    错误定义
│   │   └── msg.go    错误码
│   └── util
│       ├── cache
│       │   ├── big_cache.go    big_cache缓存
│       │   └── redis.go    redis缓存
│       ├── client.go   调用grpc客户端
│       ├── email
│       │   ├── email.go    邮件发送
│       │   └── email_helper.go   邮件发送helper
│       ├── groutine
│       │   ├── attach_panic.go   函数panic包装器
│       │   ├── errgroup.go   errgroup并发
│       │   └── grpool.go   异步任务调度
│       ├── kprocess
│       │   └── tableflip.go    进程重启
│       ├── slice.go    切片工具类
│       ├── time.go   时间工具类
│       ├── token.go    jwt token
│       └── user_name.go    用户名生成
├── repository
│   └── verify_code_record.go   验证码存储
├── router
│   ├── api
│   │   └── v1
│   │       ├── comments.go   评论controller
│   │       ├── index.go    首页controller
│   │       ├── logistics.go    物流controller
│   │       ├── merchants.go    商户controller
│   │       ├── order.go    订单controller
│   │       ├── search.go   搜索controller
│   │       ├── shop.go   店铺controller
│   │       ├── sku.go    商品库controller
│   │       ├── user.go   用户controller
│   │       ├── user_account.go   用户账户controller
│   │       ├── user_setting.go   用户设置controller
│   │       └── user_trolley.go   用户购物车controller
│   ├── process
│   │   └── prometheus.go   gin prometheus监控接口
│   └── router.go   app路由
├── service
│   ├── comments.go   评论业务调用gRPC
│   ├── common.go   公共服务
│   ├── excel_handler.go    Excel处理
│   ├── logistics.go    物流业务调用gRPC
│   ├── merchants.go    商户业务调用gRPC
│   ├── push_notice.go    队列推送
│   ├── shop_business.go    店铺业务调用gRPC
│   ├── sku_business.go   商品库业务调用gRPC
│   ├── trade_order.go    支付订单业务调用gRPC
│   ├── trade_order_report.go   支付订单业务报告
│   ├── user.go   用户业务调用gRPC
│   ├── user_account.go   用户账户调用gRPC
│   ├── user_setting.go   用户设置调用gRPC
│   └── user_trolley.go   用户购物车调用gRPC
├── sku_property_ex.json    商品库MongoDB存储参考
├── startup
│   ├── config.go   启动自定义配置
│   ├── register.go   注册自定义服务
│   └── vars.go   自定义配置变量
├── static  静态文件目录，订单报告
├── vars
│   ├── server.go   server全局变量
│   ├── setting.go    全局setting
│   └── vars.go   全局变量
├── 交流群.JPG
├── 微信赞赏码.JPG
├── 支付宝赞赏码.JPG
├── 微商城需求文档.pdf
└── 微商城系统架构设计.png
```

### 如何构建开发环境
micro-mall-xxx系列服务，希望开发者有中高级go后端开发经验，了解电商业务，mysql redis MQ使用经验     
你需要安装golang并配置golang开发环境（设置GOPATH,GOROOT,GOBIN，参考golang-install.sh）    
然后看看下面的环节      

#### 关于go mod
请各位一定配置go proxy   
GOPROXY="https://goproxy.cn,https://goproxy.io,direct"

#### 利用docker构建etcd集群
仅限于裸机运行项目   
需要先安装docker以及docker-compose，参考docker-install.sh   
建议本机也安装etcd工具（方便检测集群是否构建成功）   
运行sh docker-etcd-build.sh


#### 服务注册说明
仅限于裸机运行项目   
由于micro-mall系列服务是通过etcd来注册的，所以是需要etcd集群的，搭建步骤参考本仓库的etcd集群部署文档    
下面配置<利用docker构建etcd集群>的地址   
export  ETCDV3_SERVER_URLS=http://127.0.0.1:12379,http://127.0.0.1:22379,http://127.0.0.1:32379   
export  ETCDV3_SERVER_URL=http://127.0.0.1:12379,http://127.0.0.1:22379,http://127.0.0.1:32379   
export ETCDCTL_API=3 


#### 都有哪些服务
micro-mall-api   接入层   
micro-mall-users  用户服务   
micro-mall-users-cron  用户定时服务   
micro-mall-users-consumer  用户事件消费者   
micro-mall-order   订单服务   
micro-mall-order-cron   订单定时任务   
micro-mall-order-consumer   订单事件消费者   
micro-mall-shop   店铺服务   
micro-mall-shop-cron   店铺定时服务   
micro-mall-trolley   购物车服务   
micro-mall-sku   商品库服务   
micro-mall-sku-cron   商品库定时任务   
micro-mall-pay   支付服务   
micro-mall-pay-consumer  支付事件消费者   
micro-mall-comments   评论服务   
micro-mall-logistics   物流服务   
micro-mall-search    搜索服务   
micro-mall-search-cron   搜索定时任务   
micro-mall-search-shop-consumer   店铺信息同步消费   
micro-mall-search-sku-consumer   商品库存信息同步消费   
micro-mall-search-users-consumer   用户信息同步消费   
micro-mall-search-order-consumer   交易订单信息同步消费

#### 克隆仓库
克隆之前GOPATH要配置好   
将这些服务（目前共16个服务以及它们依赖的proto仓库，在模块分类环节可以了解到）clone到本地    
为了方便clone仓库，特此提供了一键clone micro-mall系列脚本（需要安装有git），在本仓库根目录   
mac/linux：sh batch-clone-project.sh   
windows请使用git bash shell运行：sh batch-clone-project.sh      

#### 服务启动端口说明
均可以在./etc/app.ini中配置，并在运行日志nohup.out显示当前使用的端口，进程ID   

#### 为什么本项目建议你安装python环境
这是因为本项目提供的生成pb.gw工具是python脚本，因此需要你有python环境。   
为了降低你的使用难度，我们会在后期减轻对python的依赖，进而用go或shell替代，请耐心等待。

#### 开发项目都有哪些依赖
部分依赖文件安装需要科学上网环境，演示安装步骤都是Mac环境下(同时也建议开发者使用Linux或Mac环境)，Windows请参考安装或自行Google安装   
go 1.13.15+   
goland or vscode   
mysql，redis，rabbitmq，etcd-cluster，mongo（可以不开启），elasticsearch       
protoc   安装方法如下   
```
wget https://github.com/google/protobuf/releases/download/v3.14.0/protobuf-all-3.14.0.zip
unzip protobuf-all-3.14.0.zip
cd protobuf-3.14.0/
./configure
make
make install
# 如果报错请执行ldconfig
```
go get -u google.golang.org/grpc@v1.40.0   
go get -u google.golang.org/protobuf@v1.25.0   
go get -u github.com/golang/protobuf/protoc-gen-go@v.1.4.3    
go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway@v1.14.3   
go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger   
go get -u github.com/jteeuwen/go-bindata/...   
go get -u github.com/elazarl/go-bindata-assetfs/...   
python 2.7或3.5   

#### 数据库设计
仅限于裸机运行项目   
1. 安装MySQL   
Mac下：brew install mysql   
2. 启动并初始MySQL   
Mac下：brew services start mysql   
micro-mall-系列采用分库存储，各服务拥有独立的数据库，独立的缓存，独立的事件消息   
将本仓库根目录下的micro_mall*.sql（共9个sql文件） 导入同名数据库中初始化表结构
3. 安装Redis   
Mac下：brew install redis   
4. 启动并配置   
Mac下：brew services start redis  
5. 安装MongoDB（可以不安装）   
Mac下：brew install mongodb-community   
6. 启动并配置   
Mac下：brew services start mongodb-community   
```shell
use admin  
db.createUser({
  user: 'admin',
  pwd: 'xxxx',
  roles:[{
    role: 'root',
    db: 'admin'
  }]
})
```
修改配置文件，追加如下：   
```shell
security:
  authorization: enabled
```
MongoDB需要注册一个admin用户和数据库micro_mall_sku      

#### pb.gw 代码生成
看看下面说明

##### 为什么要生成pb,gw代码
因为项目使用了gRPC+protobuf，grpc-gateway协议转换代理中间件，swagger文档托管中间件    

##### 有哪些仓库需要生成
大部分列出的服务基本上都需要，基本上micro-mall-xxx-proto就是micro-mall-xxx系列服务的依赖proto仓库   
那么就需要在这些仓库下面生成proto仓库的代码文件   

##### 如何生成
生成pb，gw代码需要本地安装protoc,protoc-gen-go，grpc-gateway 可以参考https://segmentfault.com/a/1190000013339403 grpc系列文章   
生成proto.pb.go代码时请将https://gitee.com/kelvins-io/common clone下来并放到gopath路径中(GOPATH/src/gitee.com/kelvins-io/)（如果你执行过batch-clone-project.sh脚本或已经存在了请跳过）   
生成proto.pb.go代码方式：   
在micro-mall-xxx根目录执行python genpb.py .../micro-xxx-proto   
例如：为micro-mall-api项目添加micro-users-proto的pb代码   
cd $GOPATH   
cd src/gitee.com/cristiane/ #没有则需要创建   
git clone https://gitee.com/cristiane/micro-mall-users-proto.git   
git clone https://gitee.com/cristiane/micro-mall-api.git   
cd micro-mall-api   
没有Python环境的需要安装python，Mac的话自带python   
python genpb.py ../micro-mall-users-proto   
没有报错，且检查proto目录是否创建micro-mall-users-proto目录   

#### micro-mall-api仓库现有依赖proto如何一键更新成代码
如果你不想手动执行python脚本来生成，请执行sh build-project-proto.sh      

#### 配置文件
看看下面环节   

##### 需要配置什么
基本上就是日志路径，MySQL，Redis，rabbitmq，elasticsearch，MongoDB，email这些      

##### 如何配置
在需要运行的项目根目录下etc/app.ini更改自己开发环境的配置info（可以参考默认提供的）      

##### rabbitmq如何配置
仅限于裸机运行项目   
1. 安装   
Mac下：brew install rabbitmq   

2. 启动：   
Mac下：   
rabbitmq-plugins enable rabbitmq_management   
sudo rabbitmq-server   

3. 配置
   在rabbit-manger页面配置：   
   如果mq使用rabbitmq（也可不用rabbitmq），则需要   
   创建用户，密码：micro-mall/szJ9aePR   
   创建vhost：micro-mall（下面的exchange和queue均在此vhost下面）   
   为当前用户赋予micro-mall虚拟host访问权限    
   创建exchange：user_register_notice，模式为direct    
   创建exchange：user_state_notice，模式为direct     
   创建exchange：trade_order_notice，模式为direct   
   创建exchange：trade_order_pay_callback，模式为direct   
   创建exchange：trade_pay_notice，模式为direct   
   创建exchange：shop_info_search_notice，模式为direct   
   创建exchange：sku_inventory_search_notice，模式为direct   
   创建exchange：user_info_search_notice，模式为direct   
   创建exchange：trade_order_info_search_notice，模式为direct   
   创建queue：user_register_notice，持久化为true   
   创建queue：user_state_notice，持久化为true   
   创建queue：trade_order_notice，持久化为true   
   创建queue：trade_order_pay_callback，持久化为true   
   创建queue：trade_pay_notice，持久化为true   
   创建queue：shop_info_search_notice，持久化为true   
   创建queue：sku_inventory_search_notice，持久化为true   
   创建queue：user_info_search_notice，持久化为true   
   创建queue：trade_order_info_search_notice，持久化为true

4. 填写配置文件
配置文件app.ini中关于rabbitmq配置说明（仅供参考）  
```toml
[kelvins-queue-amqp]
Broker = "amqp://用户名:密码R@地址:端口/创建的vhost"
DefaultQueue = "队列名"
ResultBackend = "redis://redis的密码@127.0.0.1:6379/10" #队列消费结果保存地址
ResultsExpireIn = 36000 # 队列任务消费结果的保存过期时间
Exchange = "交换机"
ExchangeType = "交换机类型"
BindingKey = "绑定key，一般与队列名一样"
PrefetchCount = 5 #每次从队列取任务个数
TaskRetryCount = 3 #任务失败后重试上限
TaskRetryTimeout = 36000 #任务失败后重试超时，满足斐波拉契排列
```

#### elasticsearch配置
仅限于裸机运行项目   
1. 安装elasticsearch   
Mac下：   
使用brew install elasticsearch   
其它平台参考网络教程   
2. 启动并安装对应版本的中文分词插件   
通过docker启动项目的也需要进入elasticsearch容器执行如下操作    
Mac下：   
brew services start elasticsearch   
```shell
elasticsearch-plugin install https://github.com/medcl/elasticsearch-analysis-ik/releases/download/v7.5.2/elasticsearch-analysis-ik-7.5.2.zip
```
3. 创建索引   
通过docker启动项目的也需要进入elasticsearch容器执行如下操作   
```shell
# 用户信息索引
curl -X PUT "localhost:9200/micro-mall-user-info?pretty"
# 商户申请信息索引
curl -X PUT "localhost:9200/micro-mall-merchants-material-info?pretty"
# 店铺信息索引
curl -X PUT "localhost:9200/micro-mall-shop?pretty"
# 交易订单息索引
curl -X PUT "localhost:9200/micro-mall-trade-order?pretty"
# 商品库存息索引
curl -X PUT "localhost:9200/micro-mall-sku-inventory?pretty"
# 列举索引
curl -X GET "localhost:9200/_cat/indices?v"
health status index                              uuid                   pri rep docs.count docs.deleted store.size pri.store.size
yellow open   micro-mall-shop                    93u2W6uLTn--EuzifHKFgQ   1   1         25            2       57kb           57kb
yellow open   micro-mall-user-info               WUFGM9IrR2u9d3egU-krWg   1   1      15562         5262      2.7mb          2.7mb
yellow open   micro-mall-sku-inventory           XWQYjy4PTZahB7CpHbLeiw   1   1        164            1    132.5kb        132.5kb
yellow open   micro-mall-trade-order             XWQYjy4PTZahB7CpHbLeiw   1   1        164            1    132.5kb        132.5kb
yellow open   micro-mall-merchants-material-info WAzVixxOQ4-QPFjGWEWnWA   1   1          8            0     57.9kb         57.9kb
```

#### 负载均衡
所有服务均可启动多个实例   

### 管理服务

#### 如何启动
仅限于裸机启动   
运行前请确保你已经阅读完了【如何构建开发环境】，并执行了batch-clone-project.sh    
构建且运行系列服务：   
sh build-run-all.sh   
运行系列服务   
sh start-all.sh   
重启系列服务（Windows不支持）   
sh restart-all.sh   
停止系列服务   
sh stop-all.sh   

#### 如何用docker来一键构建运行micro-mall系列项目
如果你有docker以及compose环境（docker环境可以避免安装依赖的中间件）那么你可以：   
sh docker-build-run.sh

#### 在此特别鸣谢osc为micro-mall系列项目支持docker编排方案

#### 在线服务治理
参考kelvins-tools：https://gitee.com/kelvins-io/kelvins-tools   

### 模块分类
接入层（gateway，BFF）   
https://gitee.com/cristiane/micro-mall-api   

用户服务   
https://gitee.com/cristiane/micro-mall-users   
https://gitee.com/cristiane/micro-mall-users-proto   
https://gitee.com/cristiane/micro-mall-users-consumer   
https://gitee.com/cristiane/micro-mall-users-cron   

店铺服务   
https://gitee.com/cristiane/micro-mall-shop   
https://gitee.com/cristiane/micro-mall-shop-cron      
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
https://gitee.com/cristiane/micro-mall-search-sku-consumer      
https://gitee.com/cristiane/micro-mall-search-shop-consumer    
https://gitee.com/cristiane/micro-mall-search-users-consumer    
https://gitee.com/cristiane/micro-mall-search-order-consumer    
https://gitee.com/cristiane/micro-mall-search-proto   

评论服务   
https://gitee.com/cristiane/micro-mall-comments   
https://gitee.com/cristiane/micro-mall-comments-proto

### 测试报告
![avatar](./micro-mall测试报告.png)
测试环境如下：所有服务，中间件，虚拟机，IDE，测试客户端均运行在这台机器上   
![avatar](./测试环境.png)

### micro-mall-api启用HTTP2
启用http2后，调用接口，使用wireshark抓包结果   
![avatar](./micro-mall-api启用H2wireshark抓包.png)

#### RPC接口压测报告
使用ghz压测接口   
![avatar](./ghz压测RPC接口.png)

### 接口文档
开发环境地址：  http://127.0.0.1:52001/   
监控地址：   
pprof：http://localhost:52001/debug/pprof/   
Prometheus：http://localhost:52001/debug/metrics   

返回HTTP响应码遵循HTTP通用code含义   

code | 含义
---|---
200 | 成功
400 | authorization
401 | 用户未登录
403 | 用户被禁止登录
429 | 请求被限速
500 | 服务器内部错误

HTTP 200返回body（非200时不用解析body） JSON结构：   
```json
{
   "code": 200,
   "msg": "",
   "data": {
      
   }
}
```
code含义如下：  

code | 含义
---|---
200 		 |  ok
400 		 |  请求参数错误
401 		 |  请求太多，稍后再试
500 		 |  服务器出错
40001 		 |  ID为空
40002 		 |  用户token为空
40003 		 |  用户token无效
40004 		 |  用户token过期
40005 		 |  用户不存在
40006 		 |  用户已存在
40007 		 |  用户密码错误
40008 		 |  商户未提交过认证资料
40009 		 |  商户认证资料已存在
40010 		 |  店铺认证资料已存在
40011 		 |  商户未提交过店铺认证资料
40012 		 |  商品sku已入库
40013 		 |  商品sku未入库
40014 		 |  店铺ID不存在
40015 		 |  店铺ID已存在
40016 		 |  邀请码不存在
40017 		 |  时间参数格式错误
40018 		 |  用户状态未审核或被禁止登录
50000 		 |  数据重复
50001 		 |  邮件发送错误
50002 		 |  验证码为空
50003 		 |  验证码无效
50004 		 |  验证码过期
50005 		 |  商品库存不够
50006 		 |  验证码仍在请求时间间隔内
50007 		 |  验证码在请求时间段达到最大限制
50008 		 |  用户验证码被禁止（错误尝试次数太多，请24小时后重试）
600000 		 |  金额格式解析错误
600001 		 |  用户余额不足
600002 		 |  用户账户被锁定
600003 		 |  用户账户不存在
600004 		 |  商户账户不存在
600005 		 |  商户账户被锁定
600010 		 |  事务执行失败
600011 		 |  交易号不存在
600012 		 |  订单正在支付中
600013 		 |  订单已完成支付
600014 		 |  物流记录已存在
600015 		 |  物流记录不存在
600016 		 |  用户被禁止登录（登录失败次数过多，请24小时候重试）
600017 		 |  支付时间过期
600018 		 |  订单事务号为空
600019 		 |  订单已存在
600020 		 |  用户设置信息已存在
600021 		 |  用户设置记录不存在
600022 		 |  用户物流收货地址不存在
600023 		 |  交易订单不匹配当前用户
600024 		 |  商品价格版本不存在
600025 		 |  订单状态无效
600026 		 |  订单状态被锁定
600027 		 |  订单过期
600028 		 |  订单已完成支付
600029 		 |  用户账户无效
600032 		 |  评论标签已存在
600033 		 |  评论标签不存在
600034 		 |  用户订单不存在
600035 		 |  外部交易号为空
600037 		 |  店铺状态未审核或被冻结
600038 		 |  交易单号正在支付中

接口列表（由于未及时更新，以实际接口返回为准）：   
####【说明】post请求没指明content-type的接口表单和json都支持   
1 首页   
GET    /               
返回body   

```json
{
	"code": 200,
	"data": "Welcome to micro-mall-api",
	"msg": "ok"
}
```

2  调试          
GET    /ping    
返回body   

```json
{
	"code": 200,
	"data": "2020-09-11T21:55:28.873726+08:00",
	"msg": "ok"
}
```

GET    /debug/XX  #程序内部指标   
GET     /debug/metrics  #prometheus监控   

3 发送验证码   
POST   /api/v1/verify_code/send   
请求参数：   

参数 | 含义 |  类型 | 备注  
---|------|------|---
country_code |国际码 | string | 86
phone |手机号 | string | 11位手机号
business_type |业务类型 | int | 1注册，2登录，3修改或重置密码
receive_email |接收验证码邮箱，需要在配置项主动配置email | string | xxxx@xx.com

同一手机号一段时间内同一业务只能获取一次验证码   

返回body：   
```json
{"code":200,"data":"ok","msg":"ok"}
```

4 注册用户   
POST   /api/v1/register   
请求参数：   

参数 | 含义 |  类型 | 备注  
---|------|------|---
user_name |用户名 | string | 
password |密码 | string | 可以是密文
sex |性别 | int | 1男，2女
email | 邮箱 | string | xxxx@xx.com
country_code |国际码 | string | 86
phone |手机号 | string | 11位手机号
verify_code |验证码 | string | 6位验证码
id_card_no | 身份证号 | string | 选填
account_id | 账号 | string | 选填，最长36位
invite_code |邀请码 | string | 选填

返回body：   

```json
{"code":200,"data":{"invite_code":"46e4eabbf000065"},"msg":"ok"}
```

5 验证码登陆   
POST   /api/v1/login/verify_code   
请求参数：   

参数 | 含义 |  类型 | 备注  
---|------|------|---
country_code |国际码 | string | 86
phone |手机号 | string | 11位手机号
verify_code |验证码 | string | 6位验证码

返回body：   

```json
{
   "code": 200,
   "data": {
      "user_info": {
         "id": 79860,
         "account_id": "38728714129",
         "user_name": "韦唯",
         "sex": 1,
         "phone": "38728714129",
         "country_code": "86",
         "email": "565608463@qq.com",
         "state": 3,
         "id_card_no": "100000001651933865",
         "inviter": 0,
         "invite_code": "5a4d78e48000065",
         "contact_addr": "深圳市安河桥路18号安和大院",
         "age": 33,
         "create_time": "2022-05-07 22:31:05",
         "update_time": "2022-05-07 22:31:05"
      },
      "token": "eyJhbGciOiJIUzM4NCIsInR5cCI6IkpXVCJ9.eyJ1c2VyX25hbWUiOiLpn6bllK8iLCJ1aWQiOjc5ODYwLCJleHAiOjE2NTUyOTc4NTgsImlzcyI6IndlYl9naW5fdGVtcGxhdGUifQ.Vr44kIbx5N48EkHli5P4M8COHo9gS3u52xhtTHuXBZiP-0csj4kyq0fl9L4KPeV9"
   },
   "msg": "ok"
}
```

6 密码登陆   
POST   /api/v1/login/pwd   
请求参数：   

参数 | 含义 |  类型 | 备注  
---|------|------|---
country_code |国际码 | string | 86
phone |手机号 | string | 11位手机号
password | 密码 | string | 可传md5值

返回body：   
```json
{
   "code": 200,
   "data": {
      "user_info": {
         "id": 79860,
         "account_id": "38728714129",
         "user_name": "韦唯",
         "sex": 1,
         "phone": "38728714129",
         "country_code": "86",
         "email": "565608463@qq.com",
         "state": 3,
         "id_card_no": "100000001651933865",
         "inviter": 0,
         "invite_code": "5a4d78e48000065",
         "contact_addr": "深圳市安河桥路18号安和大院",
         "age": 33,
         "create_time": "2022-05-07 22:31:05",
         "update_time": "2022-05-07 22:31:05"
      },
      "token": "eyJhbGciOiJIUzM4NCIsInR5cCI6IkpXVCJ9.eyJ1c2VyX25hbWUiOiLpn6bllK8iLCJ1aWQiOjc5ODYwLCJleHAiOjE2NTUyOTc4NTgsImlzcyI6IndlYl9naW5fdGVtcGxhdGUifQ.Vr44kIbx5N48EkHli5P4M8COHo9gS3u52xhtTHuXBZiP-0csj4kyq0fl9L4KPeV9"
   },
   "msg": "ok"
}
```

6 账号登录   
POST /api/v1/login/account   

请求参数：   

参数 | 含义 |  类型 | 备注
---|------|------|---
account_id | 账号 | string | 最长36位
password | 密码 | string | 可传md5值

返回body：
```json
{
   "code": 200,
   "data": {
      "user_info": {
         "id": 79860,
         "account_id": "38728714129",
         "user_name": "韦唯",
         "sex": 1,
         "phone": "38728714129",
         "country_code": "86",
         "email": "565608463@qq.com",
         "state": 3,
         "id_card_no": "100000001651933865",
         "inviter": 0,
         "invite_code": "5a4d78e48000065",
         "contact_addr": "深圳市安河桥路18号安和大院",
         "age": 33,
         "create_time": "2022-05-07 22:31:05",
         "update_time": "2022-05-07 22:31:05"
      },
      "token": "eyJhbGciOiJIUzM4NCIsInR5cCI6IkpXVCJ9.eyJ1c2VyX25hbWUiOiLpn6bllK8iLCJ1aWQiOjc5ODYwLCJleHAiOjE2NTUyOTc4NTgsImlzcyI6IndlYl9naW5fdGVtcGxhdGUifQ.Vr44kIbx5N48EkHli5P4M8COHo9gS3u52xhtTHuXBZiP-0csj4kyq0fl9L4KPeV9"
   },
   "msg": "ok"
}
```


7 重置用户密码   
PUT    /api/v1/user/password/reset   
header/query  token   
请求参数：   

参数 | 含义 |  类型 | 备注  
---|------|------|---
verify_code |验证码 | string | 6位验证码
password | 密码 | string | 可传md5值

返回body：   
```json
{"code":200,"data":"token","msg":"ok"}
```

8 获取用户信息   
GET    /api/v1/user/user_info     
header/query  token   

返回body： 
```json
{
	"code": 200,
	"data": {
		"id": 10009,
		"account_id": "ae23bab6-c31b-4f61-ad5e-2521a9a4917d",
		"user_name": "王尼玛",
		"sex": 1,
		"phone": "15501707783",
		"country_code": "86",
		"email": "xxx@qq.com",
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
header/query  token   

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
header/query  token   
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

```json
{"code":200,"data":{"merchant_id":111},"msg":"ok"}

```

10 添加商品到购物车   
PUT    /api/v1/user/trolley/sku/join   
header/query  token   
请求参数：   

参数 | 含义 |  类型 | 备注  
---|------|------|---
sku_code | 商品code | string | 商品唯一sku_code
shop_id | 店铺ID | int | 商品所属店铺ID
count | 数量 | int | 大于0
time | 加入时间 | string | 2020-09-05 13:25:43
selected | 是否选中 | bool | true,false

返回body： 
```json
{"code":200,"data":"ok","msg":"ok"}
```

11 从购物车中移除商品   
DELETE /api/v1/user/trolley/sku/remove   
header/query  token   
请求参数：   

参数 | 含义 |  类型 | 备注  
---|------|------|---
sku_code | 商品code | string | 商品唯一sku_code
shop_id | 店铺ID | int | 商品所属店铺ID
amount | 数量 | int | 要移除的数量，-1表示全部移除该商品

返回body： 
```json
{"code":200,"data":"ok","msg":"ok"}
```

12  获取用户购物车   
GET    /api/v1/user/trolley/sku/list   
header/query  token   
返回body： 
```json
{"code":200,"data":{"list":[{"sku_code":"df1a9633-b060-4682-9502-bc934f89392b","shop_id":29914,"count":534252790,"time":"2020-09-11 23:01:25","selected":true}]},"msg":"ok"}
```

13  商户申请店铺   
POST   /api/v1/shop_business/shop/apply   
header/query  token   
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

```json
{"code":200,"data":{"shop_id":111},"msg":"ok"}

```

14  店铺质押保证金   
PUT    /api/v1/shop_business/shop/pledge   
暂时未实现   


15  商品上架   
POST   /api/v1/sku_business/sku/put_away   
header/query  token   
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

operation_type等于4时，参数只需要shop_id,sku_code,amount

返回body： 

```json
{"code":200,"data":{},"msg":"ok"}
```

16   补充商品扩展信息   
PUT    /api/v1/sku_business/sku/supplement   
header/query  token   
请求参数：   

参数 | 含义 |  类型 | 备注  
---|------|------|---
operation_type | 操作类型 | int | 0-创建，1-更新，2删除，3-审核
sku_code | 商品sku | string | 商品唯一code，针对已上架的商品
shop_id | 店铺ID | int | 商品所属店铺ID
name | 商品名称 | string | 不能为空
size | 商品尺寸 | string | 描述商品大小，187cm x 112cm
shape | 商品形状 | string | 袋装，箱装，
production_country | 产地 | string | 不能为空
production_date | 生成日期 | string | 如2020-12-11 09:09
shelf_life | 有效期 | string | 描述过期截止时间

返回body： 

```json
{"code":200,"data":{},"msg":"ok"}

```

17   获取店铺上架商品列表   
GET    /api/v1/sku_business/sku/list   
header/query  token   
返回body： 

```json
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
```json

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
```json

```

20   获取用户购物车列表   
get /user/trolley/sku/list   
header token  

返回body   
```json

```

21  创建订单   
post /user/order/create   
header token  
请求参数：   

```json
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
一次可购买多个商品   

返回body   
```json
{"code": 200,"data": "","msg": ""}
```

22  订单支付   
post /user/order/trade   
header token  
请求参数：   

参数 | 含义 |  类型 | 备注  
---|------|------|---
tx_code | 订单交易号 | string | 不能为空

订单交易号需要是自己创建的订单交易号   

返回body   
```json
{"code": 200,"data": "","msg": ""}
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
send_user | 发送方 | string | 汉东省京州市月牙湖房地产开发公司
send_user_id | 发送方用户id | int64 | 用户ID
send_addr | 发送地址 | string | 汉东省京州市月牙湖房地产开发公司
send_phone | 发送方联系方式 | string | 如，13683749374
send_time | 发送时间 | string | 如，2020-10-10 10:10:10
receive_user | 接收方 | string | 赵富贵
receive_user_id | 接收方用户id | int64 | 用户ID
receive_addr | 接收方地址 | string | 四川省成都市武侯区98号
receive_phone | 接收方联系方式 | string | 如，0838-10182827
goods | 需要承运的货物 | string | 如，下面序列化后的值

goods示范
```json
[{
	"sku_code": "2131d-f111-45e1-b68a-d602c2f0f1b3",
	"name": "怡宝矿泉水",
	"kind": "饮用水",
	"count": 3
}]
```

24  用户配置收货地址   
post json   
/api/v1/user/setting/address

请求body   
```json
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
```json
{"code":200,"data":"","msg":"ok"}
```

24 用户查询收货地址列表   
get    
/api/v1/user/setting/address?delivery_id=xx  
  
返回body   
```json
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
```json
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
```json
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
	}],
	"msg": "ok"
}
```

获取店铺订单报告   
get  /user/order/report   
header/query  token   

参数 | 含义 |  类型 | 备注  
---|------|------|---
shop_id | 店铺ID | int | 不能为空
start_time | 统计开始时间 | string | 如，2019-11-22 08:46:41
end_time | 统计结束时间 | string | 如，2020-12-04 18:46:41
page_size | 分页大小 | int | 500，最小1
page_num | 分页号 | int | 最小1

返回body：Excel报告   


用户账户充值   
post  /user/account/charge   
header/query  token   

参数 | 含义 |  类型 | 备注  
---|------|------|---
account_type | 账户类型 | int | 0-个人账户，2-公司账户，3-系统账户
amount | 金额 | string | 如，99.09
coin_type | 币种 | int | 如，0-RMB，1-USDT
device_code | 设备 | string | vivo NEX
device_platform | 平台 | string | Android
out_trade_no | 外部交易号 | string | uuid

返回body   
```json
{"code":200,"data":"","msg":"ok"}
```

订单评价   
post/json  /user/comments/order/create   
header/query  token   

```json
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
```json
{"code":200,"data":"","msg":"ok"}
```

获取店铺评价   
get  /user/comments/shop/list?shop_id=111    
返回body   

```json
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
```json
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

```json
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

列举用户   
get /user/user_info/list?page_size=3&page_num=1&token=xxx   
请求参数：   

参数 | 含义 |  类型 | 备注  
---|------|------|---
page_size | 页码大小 | int | 大于0
page_num | 页码 | int | 大于0
token | 授权码 | string | 需要用户服务特别授权码
```json
{
	"code": 200,
	"data": {
		"user_info_list": [{
			"country_code": "1015",
			"phone": "19494479989"
		}, {
			"country_code": "1017",
			"phone": "90687911165"
		}]
	},
	"msg": "ok"
}
```

用户信息搜索   
get /search/user_info?query=王友      
响应：   
```json
{
	"code": 200,
	"data": [{
		"info": {
			"uid": 10141,
			"account_id": "bf00de2c-1612-47a4-bddf-0e97b912d8ff",
			"user_name": "王友",
			"sex": 1,
			"country_code": "86",
			"phone": "41606450736",
			"email": "mybaishati@gmail.com",
			"state": 3,
			"id_card_no": "100000001606450736",
			"inviter_code": "495bdcec1000065",
			"contact_addr": "南京市王友大院",
			"age": 33,
			"create_time": "2020-11-27T12:18:56+08:00"
		},
		"score": 11.652336
	}, {
		"info": {
			"uid": 12371,
			"account_id": "eee4a1e7-79c0-4f6c-bd9e-af01e139d7bf",
			"user_name": "甘友",
			"sex": 1,
			"country_code": "86",
			"phone": "41606451085538176000",
			"email": "mybaishati@gmail.com",
			"state": 3,
			"id_card_no": "100000001606451085",
			"inviter_code": "495be574d000065",
			"contact_addr": "南京市甘友大院",
			"age": 33,
			"create_time": "2020-11-27T12:24:45+08:00"
		},
		"score": 6.071039
	}],
	"msg": "ok"
}
```
商户认证信息搜索   
get /search/merchant_info?query=王友   
响应：
```json
{
	"code": 200,
	"data": [{
		"info": {
			"uid": 79857,
			"material_id": 1172,
			"register_addr": "京港市上海路111号",
			"health_card_no": "R8nJ65TDUGAlqrwerSdb9",
			"identity": 1,
			"state": 3,
			"tax_card_no": "qX2Mr545kznWrlvO4sIp7",
			"create_time": "2021-10-03T01:50:15+08:00"
		},
		"score": 5.01196
	}],
	"msg": "ok"
}
```
交易订单搜索   
get /search/merchant_info?query=龟苓膏   
响应：     
```json
{
  "code": 200,
  "data": [{
    "order_code": "b250d641-ce9c-4f91-882d-b2b2b6fb1d6e",
    "shop_id": 30071,
    "money": "共计：50.4600",
    "description": "国庆大促",
    "create_time": "2021-10-04 20:10:04",
    "pay_state": "支付取消",
    "pay_time": "2021-10-04 21:00:30",
    "shop_address": "舟山市来福路90号来福花园5栋1单元2309",
    "shop_name": "大发发连锁超市",
    "goods": [{
      "goods_name": "高乐低儿童遥控玩具",
      "price": "2.9000000000000000",
      "sku_code": "dd13b4aa-4121-4898-a2b5-bcfebccb713b",
      "amount": 1
    }, {
      "goods_name": "富源不爱龟苓膏",
      "price": "23.7800000000000000",
      "sku_code": "a3e5da0a-d3aa-43e2-a7b8-2c5e264e2a09",
      "amount": 2
    }],
    "score": 18.350704
  }, {
    "order_code": "178aa53b-925e-4a0f-9786-93a47411143d",
    "shop_id": 30071,
    "money": "共计：50.4600",
    "description": "国庆大促",
    "create_time": "2021-10-04 21:02:39",
    "pay_state": "未支付",
    "pay_time": "2021-10-04 21:02:39",
    "shop_address": "舟山市来福路90号来福花园5栋1单元2309",
    "shop_name": "大发发连锁超市",
    "goods": [{
      "goods_name": "高乐低儿童遥控玩具",
      "price": "2.9000000000000000",
      "sku_code": "dd13b4aa-4121-4898-a2b5-bcfebccb713b",
      "amount": 1
    }, {
      "goods_name": "富源不爱龟苓膏",
      "price": "23.7800000000000000",
      "sku_code": "a3e5da0a-d3aa-43e2-a7b8-2c5e264e2a09",
      "amount": 2
    }],
    "score": 18.350704
  }, {
    "order_code": "296d0410-75a9-424d-bfa3-8767364abf67",
    "shop_id": 30072,
    "money": "共计：39.8000",
    "description": "国庆大促",
    "create_time": "2021-10-04 21:02:39",
    "pay_state": "未支付",
    "pay_time": "2021-10-04 21:02:39",
    "shop_address": "辽宁省葫芦岛市迎宾路33号轻工业产业园67号",
    "shop_name": "辽宁省葫芦岛电子商务公司",
    "goods": [{
      "goods_name": "张宇解百纳龟苓膏",
      "price": "19.9000000000000000",
      "sku_code": "dd13b4aa-4121-a2b5-a2b5-bcfebccb4898",
      "amount": 1
    }, {
      "goods_name": "张鱼解百纳2008窖藏特曲红酒",
      "price": "19.9000000000000000",
      "sku_code": "dd13b4aa-4121-a2b5-a2b5-bcfebccb4898",
      "amount": 1
    }],
    "score": 16.219776
  }],
  "msg": "ok"
}
```

### 赞助

**1 感谢 jetbrains 为本项目提供的 goland 激活码**   
![avatar](./icon-goland.png)

### 共同开源
1 请star收藏项目   
2 提出issue   
3 fork项目   
4 clone fork后的项目到本地开发   
5 提交pr   
6 参与开源代码被合并后将获得goland 正版license   
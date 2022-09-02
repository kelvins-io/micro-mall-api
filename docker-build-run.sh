#! /bin/bash
project_names=(
"micro-mall-api"
"micro-mall-users"
"micro-mall-users-cron"
"micro-mall-users-consumer"
"micro-mall-shop"
"micro-mall-shop-cron"
"micro-mall-trolley"
"micro-mall-sku"
"micro-mall-sku-cron"
"micro-mall-order"
"micro-mall-order-cron"
"micro-mall-order-consumer"
"micro-mall-pay"
"micro-mall-pay-consumer"
"micro-mall-logistics"
"micro-mall-comments"
"micro-mall-search"
"micro-mall-search-cron"
"micro-mall-search-order-consumer"
"micro-mall-search-shop-consumer"
"micro-mall-search-sku-consumer"
"micro-mall-search-users-consumer"
)

echo 每个仓库服务配置默认会使用etc/app-docker.ini.example里面对应服务的配置
echo 如果你有自定义配置-如邮箱发送设置-请提前修改每个仓库下etc/app-docker.ini.example以便构建的时候自动覆盖
echo 容器环境下每个仓库下etc/app-docker.ini.example配置文件会覆盖为执行目录下的app.ini

echo 开始配置环境变量
# 配置环境遍历
# shellcheck disable=SC2091
$(go env -w GO111MODULE=on)
# shellcheck disable=SC2091
$(go env -w GOPROXY=https://goproxy.cn,https://goproxy.io,direct)
#记录api项目的路径
api_path=$(pwd)

# 循环目录 并编译可执行文件
function loopPathBuild() {
  for file in ${project_names[*]}; do
      cd "$file" || exit
      echo "$file"
      if [ ! -d "vendor" ]; then
        go mod vendor
      fi
      cp -rf ./etc/app-docker.ini.example ./etc/app.ini
      CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o "${file}" main.go
      cd ../
  done
}
# 返回上一级
cd ../
# 遍历所有目录并编译可执行文件
echo 开始初始化项目配置文件以及编译项目可执行文件
# shellcheck disable=SC2046
loopPathBuild $(pwd)

# 返回api项目 开始执行docker流程
cd "$api_path" || exit

echo 配置容器环境变量
export COMPOSE_HTTP_TIMEOUT=500
export DOCKER_CLIENT_TIMEOUT=500
export COMPOSE_PARALLEL_LIMIT=1024

echo 重置容器网络
echo y | docker network prune

echo 新建容器专用网络
docker network create mall

echo 构建中间件服务容器基础环境
echo 这些中间件服务的标准端口会映射到物理机端口-当然你也可以自行注释掉
echo 中间件服务容器端口映射到物理机是为了方便在物理机上就能给其安装插件或初始配置
docker-compose up -d
# shellcheck disable=SC2009
echo "显示基础组件运行情况"
ps -ef | grep mysql5_7
ps -ef | grep redis
ps -ef | grep rabbitmq
ps -ef | grep mongo
ps -ef | grep elasticsearch

echo "配置rabbitmq，下面注释的命令手动执行"
#docker exec -it $(docker ps -aqf "name=rabbitmq") /bin/bash
#rabbitmqctl add_vhost micro-mall
#rabbitmqctl set_permissions -p micro-mall root ".*" ".*" ".*"
#rabbitmqadmin -u root -p micro-mall -V micro-mall declare exchange name=user_register_notice type=direct
#rabbitmqadmin -u root -p micro-mall -V micro-mall declare exchange name=user_state_notice type=direct
#rabbitmqadmin -u root -p micro-mall -V micro-mall declare exchange name=trade_order_notice type=direct
#rabbitmqadmin -u root -p micro-mall -V micro-mall declare exchange name=trade_order_pay_callback type=direct
#rabbitmqadmin -u root -p micro-mall -V micro-mall declare exchange name=trade_pay_notice type=direct
#rabbitmqadmin -u root -p micro-mall -V micro-mall declare exchange name=shop_info_search_notice type=direct
#rabbitmqadmin -u root -p micro-mall -V micro-mall declare exchange name=sku_inventory_search_notice type=direct
#rabbitmqadmin -u root -p micro-mall -V micro-mall declare exchange name=user_info_search_notice type=direct
#rabbitmqadmin -u root -p micro-mall -V micro-mall declare exchange name=trade_order_info_search_notice type=direct
#rabbitmqadmin -u root -p micro-mall -V micro-mall declare queue name=user_register_notice
#rabbitmqadmin -u root -p micro-mall -V micro-mall declare queue name=user_state_notice
#rabbitmqadmin -u root -p micro-mall -V micro-mall declare queue name=trade_order_notice
#rabbitmqadmin -u root -p micro-mall -V micro-mall declare queue name=trade_order_pay_callback
#rabbitmqadmin -u root -p micro-mall -V micro-mall declare queue name=trade_pay_notice
#rabbitmqadmin -u root -p micro-mall -V micro-mall declare queue name=shop_info_search_notice
#rabbitmqadmin -u root -p micro-mall -V micro-mall declare queue name=sku_inventory_search_notice
#rabbitmqadmin -u root -p micro-mall -V micro-mall declare queue name=user_info_search_notice
#rabbitmqadmin -u root -p micro-mall -V micro-mall declare queue name=trade_order_info_search_notice
#exit

echo "配置elasticsearch，下面注释的命令手动执行"
#docker exec -it $(docker ps -aqf "name=elasticsearch") /bin/bash
#echo y |elasticsearch-plugin install https://github.com/medcl/elasticsearch-analysis-ik/releases/download/v7.5.2/elasticsearch-analysis-ik-7.5.2.zip
#curl -X PUT "localhost:9200/micro-mall-user-info?pretty"
#curl -X PUT "localhost:9200/micro-mall-merchants-material-info?pretty"
#curl -X PUT "localhost:9200/micro-mall-shop?pretty"
#curl -X PUT "localhost:9200/micro-mall-trade-order?pretty"
#curl -X PUT "localhost:9200/micro-mall-sku-inventory?pretty"
#curl -X GET "localhost:9200/_cat/indices?v"
#exit

echo "构建并运行容器项目"
docker-compose -f docker-compose-build.yml up -d

# shellcheck disable=SC2009
ps -ef | grep micro-mall
echo "开启micro-mall的旅行吧"
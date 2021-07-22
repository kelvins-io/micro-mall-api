#! /bin/bash
project_names=(
"micro-mall-api"
"micro-mall-users"
"micro-mall-users-consumer"
"micro-mall-shop"
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
)

echo 编译可执行文件
# 配置环境遍历
# shellcheck disable=SC2091
$(go env -w GO111MODULE=on)
# shellcheck disable=SC2091
$(go env -w GOPROXY=https://goproxy.io,direct)
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
# shellcheck disable=SC2046
loopPathBuild $(pwd)

# 返回api项目 开始执行docker流程
cd "$api_path" || exit

echo 配置环境变量
export COMPOSE_HTTP_TIMEOUT=500
export DOCKER_CLIENT_TIMEOUT=500
export COMPOSE_PARALLEL_LIMIT=1024

echo 重置网络
echo y | docker network prune

echo 新建专用网络
docker network create mall
echo 构建基础环境
docker-compose up -d
echo 运行项目
docker-compose -f docker-compose-build.yml up -d

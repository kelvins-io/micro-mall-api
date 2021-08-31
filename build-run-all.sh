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
#记录api项目的路径
# shellcheck disable=SC2034
api_path=$(pwd)
# 循环目录 构建运行
function loopPathBuildRun() {
  for file in ${project_names[*]}; do
      cd "$file" || exit
      echo "=> $file"
      sh build-run.sh
      cd ../
  done
}

# 返回上一级
cd ../
# 遍历所有目录构建运行
# shellcheck disable=SC2046
loopPathBuildRun $(pwd)
# shellcheck disable=SC2028
echo "\n"
echo "显示 micro-mall-* 进程运行状态"
# shellcheck disable=SC2009
ps -ef | grep micro-mall
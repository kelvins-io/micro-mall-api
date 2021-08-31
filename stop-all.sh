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

# 遍历所有目录停止进程
function loopPathStop() {
  for file in ${project_names[*]}; do
      cd "$file" || exit
      echo "=> $file"
      sh stop.sh
      cd ../
  done
}

# 返回上一级
cd ../

# shellcheck disable=SC2046
loopPathStop $(pwd)
# shellcheck disable=SC2028
echo "\n"
echo "显示 micro-mall-* 进程运行状态"
# shellcheck disable=SC2009
ps -ef | grep micro-mall

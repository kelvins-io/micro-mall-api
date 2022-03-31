#! /bin/bash
project_names=(
"micro-mall-api"
"micro-mall-users"
"micro-mall-users-consumer"
"micro-mall-users-cron"
"micro-mall-shop"
"micro-mall-shop-cron"
"micro-mall-trolley"
"micro-mall-sku"
"micro-mall-sku-cron"
"micro-mall-order"
"micro-mall-order-cron"
"micro-mall-order-consumer"
"micro-mall-order-consumer1"
"micro-mall-order-consumer2"
"micro-mall-pay"
"micro-mall-pay-consumer"
"micro-mall-pay-consumer1"
"micro-mall-pay-consumer2"
"micro-mall-logistics"
"micro-mall-comments"
"micro-mall-search"
"micro-mall-search-cron"
"micro-mall-search-shop-consumer"
"micro-mall-search-sku-consumer"
"micro-mall-search-users-consumer"
"micro-mall-search-order-consumer"
"micro-mall-search-order-consumer1"
"micro-mall-search-order-consumer2"
)

# 遍历所有目录重启进程
echo "重启进程不支持Windows平台"

function loopPathRestart() {
  for file in ${project_names[*]}; do
      cd "$file" || exit
      echo "=> $file"
      sh restart.sh
      cd ../
  done
}

# 返回上一级
cd ../

# shellcheck disable=SC2046
loopPathRestart $(pwd)
# shellcheck disable=SC2028
echo "\n"
echo "重启完成，显示 micro-mall-* 进程运行状态"
# shellcheck disable=SC2009
ps -ef | grep micro-mall
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
"micro-mall-pay"
"micro-mall-pay-consumer"
"micro-mall-logistics"
"micro-mall-comments"
"micro-mall-search"
"micro-mall-search-cron"
"micro-mall-search-shop-consumer"
"micro-mall-search-sku-consumer"
"micro-mall-search-users-consumer"
)

# 遍历所有目录 启动
function loopPathStart() {
  for file in ${project_names[*]}; do
      cd "$file" || exit
      echo "=> $file"
      sh start.sh
      cd ../
  done
}

# 返回上一级
cd ../

# shellcheck disable=SC2046
loopPathStart $(pwd)
# shellcheck disable=SC2028
echo "\n"
echo "启动完成，显示 micro-mall-* 进程运行状态"
# shellcheck disable=SC2009
ps -ef | grep micro-mall

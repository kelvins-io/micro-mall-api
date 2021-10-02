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

# 遍历所有目录 清楚bin
function loopPathCleanBin() {
  for file in ${project_names[*]}; do
      cd "$file" || exit
      echo "=> $file"
      sh clean-bin.sh
      cd ../
  done
}

# 返回上一级
cd ../
echo "开始遍历目录下bin文件"
# shellcheck disable=SC2046
loopPathCleanBin $(pwd)
# shellcheck disable=SC2028
echo "清理完成"
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

# 遍历所有目录 清理logs
function loopPathCleanLogs() {
  for file in ${project_names[*]}; do
      cd "$file" || exit
      echo "=> $file"
      sh clean-logs.sh
      cd ../
  done
}

# 返回上一级
cd ../
# 遍历所有目录 清理logs
echo "开始遍历目录下logs目录"
# shellcheck disable=SC2046
loopPathCleanLogs $(pwd)
# shellcheck disable=SC2028
echo "清理完成"
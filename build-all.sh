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

# 遍历所有目录编译
function loopPathBuild() {
  for file in ${project_names[*]}; do
      cd "$file" || exit
      echo "=> $file"
      sh build.sh
      cd ../
  done
}

# 返回上一级
cd ../

# shellcheck disable=SC2046
loopPathBuild $(pwd)
# shellcheck disable=SC2028
echo "\n"
echo "编译完成"
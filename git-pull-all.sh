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

function loopPathGitPull() {
  for file in ${project_names[*]}; do
      cd "$file" || exit
      echo "=> $file"
      git status
      git fetch origin
      git pull
      cd ../
  done
}

# 返回上一级
cd ../
# 遍历所有目录 执行Git pull
# shellcheck disable=SC2046
loopPathGitPull $(pwd)
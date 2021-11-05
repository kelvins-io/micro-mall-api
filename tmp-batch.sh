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
"micro-mall-search-order-consumer"
)

# 遍历所有目录 执行任务
function loopPathExec() {
  for file in ${project_names[*]}; do
      cd "$file" || exit
      echo "=> $file"
      git remote remove origin
      git remote add origin https://gitee.com/cristiane/$file.git
      git push --set-upstream origin master
      # shellcheck disable=SC2028
      echo "执行一些批量任务"
      cd ../
  done
}

# 返回上一级
cd ../

# shellcheck disable=SC2046
loopPathExec $(pwd)
# shellcheck disable=SC2028

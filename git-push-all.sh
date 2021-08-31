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

# shellcheck disable=SC2236
if [ ! -n "$1" ]; then
  echo "请输入git commit 信息"
  exit 1
else
  echo ""
fi

function loopPathGitPush() {
  for file in ${project_names[*]}; do
      cd "$file" || exit
      echo "=> $file"
      git status
      # shellcheck disable=SC2046
      statusConfirm $(pwd)
      git add .
      git commit -m "$1"
      git push origin
      git push github
      cd ../
  done
}

function statusConfirm() {
    read -r -p "确认提交吗? [Y-y/N-n] " input
    case $input in
      [yY][eE][sS]|[yY])
      echo "Yes"
      ;;

      [nN][oO]|[nN])
      echo "No"
      exit 1
          ;;

      *)
      echo "Invalid input..."
      exit 1
      ;;
    esac
}

# 返回上一级
cd ../
# 遍历所有目录 执行git push
# shellcheck disable=SC2046
loopPathGitPush $(pwd)
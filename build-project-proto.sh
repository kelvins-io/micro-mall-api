#! /bin/bash
git branch

echo 本shell需要你安装了python环境用来一键生成pb,gw代码
read -r -p "你安装python了吗? [Y-y/N-n] " input
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

python genpb.py ../micro-mall-users-proto/
python genpb.py ../micro-mall-trolley-proto/
python genpb.py ../micro-mall-sku-proto/
python genpb.py ../micro-mall-order-proto/
python genpb.py ../micro-mall-pay-proto/
python genpb.py ../micro-mall-logistics-proto/
python genpb.py ../micro-mall-comments-proto/
python genpb.py ../micro-mall-shop-proto/
# shellcheck disable=SC2034
# shellcheck disable=SC2153

if [ ! -d "$GOPATH" ]; then
  echo "没有找到GOPATH目录，退出"
  exit
fi


if [ ! -d "$GOPATH/src/gitee.com/kelvins-io" ]; then
  mkdir "$GOPATH/src/gitee.com/kelvins-io"
fi
cd "$GOPATH/src/gitee.com/kelvins-io" || exit
echo "开始clone kelvins-io/common 仓库"
git clone git@gitee.com:kelvins-io/common.git


if [ ! -d "$GOPATH/src/gitee.com/cristiane" ]; then
  mkdir "$GOPATH/src/gitee.com/cristiane"
fi
cd "$GOPATH/src/gitee.com/cristiane" || exit
echo "开始clone micro-mall-api 仓库"
git clone git@gitee.com:cristiane/micro-mall-api.git
echo "开始clone micro-mall-users 仓库"
git clone git@gitee.com:cristiane/micro-mall-users.git
git clone git@gitee.com:cristiane/micro-mall-users-proto.git
git clone git@gitee.com:cristiane/micro-mall-users-consumer.git
echo "开始clone micro-mall-shop 仓库"
git clone git@gitee.com:cristiane/micro-mall-shop.git
git clone git@gitee.com:cristiane/micro-mall-shop-proto.git
echo "开始clone micro-mall-sku 仓库"
git clone git@gitee.com:cristiane/micro-mall-sku.git
git clone git@gitee.com:cristiane/micro-mall-sku-proto.git
git clone git@gitee.com:cristiane/micro-mall-sku-cron.git
echo "开始clone micro-mall-trolley 仓库"
git clone git@gitee.com:cristiane/micro-mall-trolley.git
git clone git@gitee.com:cristiane/micro-mall-trolley-proto.git
echo "开始clone micro-mall-order 仓库"
git clone git@gitee.com:cristiane/micro-mall-order.git
git clone git@gitee.com:cristiane/micro-mall-order-proto.git
git clone git@gitee.com:cristiane/micro-mall-order-cron.git
git clone git@gitee.com:cristiane/micro-mall-order-consumer.git
echo "开始clone micro-mall-pay 仓库"
git clone git@gitee.com:cristiane/micro-mall-pay.git
git clone git@gitee.com:cristiane/micro-mall-pay-proto.git
git clone git@gitee.com:cristiane/micro-mall-pay-consumer.git
echo "开始clone micro-mall-logistics 仓库"
git clone git@gitee.com:cristiane/micro-mall-logistics.git
git clone git@gitee.com:cristiane/micro-mall-logistics-proto.git
echo "开始clone micro-mall-search 仓库"
git clone git@gitee.com:cristiane/micro-mall-search.git
git clone git@gitee.com:cristiane/micro-mall-search-proto.git
git clone git@gitee.com:cristiane/micro-mall-search-cron.git
echo "开始clone micro-mall-comments 仓库"
git clone git@gitee.com:cristiane/micro-mall-comments.git
git clone git@gitee.com:cristiane/micro-mall-comments-proto.git
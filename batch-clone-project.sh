#! /bin/bash
# shellcheck disable=SC2034
# shellcheck disable=SC2153

git branch

GIT_CLONE_METHOD="$1"
GIT_CLONE_METHOD_URL="git@gitee.com:"

if [ "$GIT_CLONE_METHOD" = "http" ] || [ "$GIT_CLONE_METHOD" = "ssh" ] ;then
    # git clone https
    if [ "$GIT_CLONE_METHOD" = "http" ];then
       GIT_CLONE_METHOD_URL="https://gitee.com/"
    fi
    # git clone ssh
    if [ "$GIT_CLONE_METHOD" = "ssh" ];then
       GIT_CLONE_METHOD_URL="git@gitee.com:"
    fi
else
    echo "Reference:"
    echo "  ./batch-clone-project.sh http (default)"
    GIT_CLONE_METHOD_URL="https://gitee.com/"
fi

#echo ${GIT_CLONE_METHOD_URL}

if [ ! -d "$GOPATH" ]; then
  echo "没有找到 GOPATH 目录,exit"
  exit
fi


if [ ! -d "$GOPATH/src/gitee.com/kelvins-io" ]; then
  mkdir -p "$GOPATH/src/gitee.com/kelvins-io"
fi
cd "$GOPATH/src/gitee.com/kelvins-io" || exit
echo "开始clone kelvins-io/common 仓库"
git clone ${GIT_CLONE_METHOD_URL}kelvins-io/common.git

if [ ! -d "$GOPATH/src/gitee.com/cristiane" ]; then
  mkdir -p "$GOPATH/src/gitee.com/cristiane"
fi
cd "$GOPATH/src/gitee.com/cristiane" || exit
echo "开始clone micro-mall-api 仓库"
git clone ${GIT_CLONE_METHOD_URL}cristiane/micro-mall-api.git
echo "开始clone micro-mall-users 仓库"
git clone ${GIT_CLONE_METHOD_URL}cristiane/micro-mall-users.git
git clone ${GIT_CLONE_METHOD_URL}cristiane/micro-mall-users-proto.git
git clone ${GIT_CLONE_METHOD_URL}cristiane/micro-mall-users-consumer.git
echo "开始clone micro-mall-shop 仓库"
git clone ${GIT_CLONE_METHOD_URL}cristiane/micro-mall-shop.git
git clone ${GIT_CLONE_METHOD_URL}cristiane/micro-mall-shop-proto.git
echo "开始clone micro-mall-sku 仓库"
git clone ${GIT_CLONE_METHOD_URL}cristiane/micro-mall-sku.git
git clone ${GIT_CLONE_METHOD_URL}cristiane/micro-mall-sku-proto.git
git clone ${GIT_CLONE_METHOD_URL}cristiane/micro-mall-sku-cron.git
echo "开始clone micro-mall-trolley 仓库"
git clone ${GIT_CLONE_METHOD_URL}cristiane/micro-mall-trolley.git
git clone ${GIT_CLONE_METHOD_URL}cristiane/micro-mall-trolley-proto.git
echo "开始clone micro-mall-order 仓库"
git clone ${GIT_CLONE_METHOD_URL}cristiane/micro-mall-order.git
git clone ${GIT_CLONE_METHOD_URL}cristiane/micro-mall-order-proto.git
git clone ${GIT_CLONE_METHOD_URL}cristiane/micro-mall-order-cron.git
git clone ${GIT_CLONE_METHOD_URL}cristiane/micro-mall-order-consumer.git
echo "开始clone micro-mall-pay 仓库"
git clone ${GIT_CLONE_METHOD_URL}cristiane/micro-mall-pay.git
git clone ${GIT_CLONE_METHOD_URL}cristiane/micro-mall-pay-proto.git
git clone ${GIT_CLONE_METHOD_URL}cristiane/micro-mall-pay-consumer.git
echo "开始clone micro-mall-logistics 仓库"
git clone ${GIT_CLONE_METHOD_URL}cristiane/micro-mall-logistics.git
git clone ${GIT_CLONE_METHOD_URL}cristiane/micro-mall-logistics-proto.git
echo "开始clone micro-mall-search 仓库"
git clone ${GIT_CLONE_METHOD_URL}cristiane/micro-mall-search.git
git clone ${GIT_CLONE_METHOD_URL}cristiane/micro-mall-search-proto.git
git clone ${GIT_CLONE_METHOD_URL}cristiane/micro-mall-search-cron.git
echo "开始clone micro-mall-comments 仓库"
git clone ${GIT_CLONE_METHOD_URL}cristiane/micro-mall-comments.git
git clone ${GIT_CLONE_METHOD_URL}cristiane/micro-mall-comments-proto.git

exit 0

# shellcheck disable=SC2034
# shellcheck disable=SC2153

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

echo 拉取依赖
go mod vendor

read -r -p "是否只构建当前平台版本? [Y-y/N-n] " input
sysOS=$(uname -s)
case $input in
    [yY][eE][sS]|[yY])
    echo "Yes"
    echo 开始构建版本
    if [ "$sysOS" == "Darwin" ] ; then
    CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -o micro-mall-api-darwin-amd64 main.go
    elif [ "$sysOS" == "Linux" ]; then
    CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o micro-mall-api-linux-amd64 main.go
    else
    CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o micro-mall-api-windows-amd64.exe main.go
    fi
    ;;

    [nN][oO]|[nN])
    echo "No"
    echo 开始构建Darwin版本
    CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -o micro-mall-api-darwin-amd64 main.go
    echo 开始构建Linux版本
    CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o micro-mall-api-linux-amd64 main.go
    echo 开始构建Windows版本
    CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o micro-mall-api-windows-amd64.exe main.go
#    exit 1
    ;;

    *)
    echo "Invalid input..."
    exit 1
    ;;
esac


read -r -p "你需要运行此项目吗? [Y-y/N-n] " input
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

echo 运行项目前请确保你已经配置好了./etc/app.ini项目配置文件
read -r -p "你配置妥当了吗? [Y-y/N-n] " input
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


mkdir -p logs
echo 开始运行micro-mall-api
sysOS=$(uname -s)
if [ "$sysOS" == "Darwin" ] ; then
./micro-mall-api-darwin-amd64
elif [ "$sysOS" == "Linux" ]; then
./micro-mall-api-linux-amd64
else
./micro-mall-api-windows-amd64.exe
fi

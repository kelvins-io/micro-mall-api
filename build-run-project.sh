# shellcheck disable=SC2034
# shellcheck disable=SC2153
# shellcheck disable=SC2046
# shellcheck disable=SC2003

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

echo 开始生成pb,gw文件
python genpb.py ../micro-mall-comments-proto
python genpb.py ../micro-mall-logistics-proto
python genpb.py ../micro-mall-order-proto
python genpb.py ../micro-mall-pay-proto
python genpb.py ../micro-mall-sku-proto
python genpb.py ../micro-mall-shop-proto
python genpb.py ../micro-mall-trolley-proto
python genpb.py ../micro-mall-users-proto

echo 拉取依赖
go mod vendor

echo 开始构建linux-amd64版本
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o micro-mall-api-linux-amd64 main.go
echo 开始构建macOS-amd64版本
CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -o micro-mall-api-darwin-amd64 main.go
echo 开始构建windows-amd64版本
CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o micro-mall-api-windows-amd64.exe main.go
echo 构建完毕

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
if [ "$(uname)" == "Darwin" ] ; then
./micro-mall-api-darwin-amd64
elif [ "$(expr substr $(uname -s) 1 5)" == "Linux" ]; then
./micro-mall-api-linux-amd64
elif [ "$(expr substr $(uname -s) 1 10)" == "MINGW32_NT" ]; then
./micro-mall-api-windows-amd64.exe
fi

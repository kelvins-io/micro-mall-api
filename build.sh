echo 拉取依赖
go mod vendor

cp -n ./etc/app.ini.example ./etc/app.ini

mkdir -p logs
mkdir -p static

echo 开始构建版本
go build -o micro-mall-api main.go

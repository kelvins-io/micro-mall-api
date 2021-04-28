echo 拉取依赖
go mod vendor

cp -n ./etc/app.ini.example ./etc/app.ini

sysOS=$(uname -s)

echo 开始构建版本
if [ "$sysOS" == "Darwin" ] ; then
CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -o micro-mall-api-darwin-amd64 main.go
elif [ "$sysOS" == "Linux" ]; then
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o micro-mall-api-linux-amd64 main.go
else
CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o micro-mall-api-windows-amd64.exe main.go
fi
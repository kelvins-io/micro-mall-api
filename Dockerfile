FROM golang:latest

WORKDIR $GOPATH/src/gitee.com/cristiane/micro-mall-api
COPY . $GOPATH/src/gitee.com/cristiane/micro-mall-api
RUN go build -o micro-mall-api .

EXPOSE 52001
ENTRYPOINT ["./micro-mall-api"]
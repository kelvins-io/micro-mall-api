package router

import (
	"gitee.com/cristiane/micro-mall-api/middleware"
	v1 "gitee.com/cristiane/micro-mall-api/router/api/v1"
	"gitee.com/cristiane/micro-mall-api/vars"
	"io"
	"os"

	"github.com/gin-gonic/gin"
)

func InitRouter(accessInfoLogger, accessErrLogger io.Writer) *gin.Engine {

	gin.DefaultWriter = io.MultiWriter(os.Stdout, accessInfoLogger)
	gin.DefaultErrorWriter = io.MultiWriter(os.Stderr, accessErrLogger)

	gin.SetMode(gin.ReleaseMode) // 默认生产
	if vars.ServerSetting != nil && vars.ServerSetting.Mode != "" {
		gin.SetMode(vars.ServerSetting.Mode)
	}

	r := gin.Default()
	r.Use(middleware.Cors())

	r.GET("/", v1.IndexApi)
	r.GET("/ping", v1.PingApi)
	apiG := r.Group("/api")
	apiV1 := apiG.Group("/v1")
	apiV1.GET("/options", v1.IndexApi)

	return r
}

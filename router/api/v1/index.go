package v1

import (
	"gitee.com/cristiane/micro-mall-api/pkg/app"
	"gitee.com/cristiane/micro-mall-api/pkg/code"
	"gitee.com/cristiane/micro-mall-api/vars"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func IndexApi(c *gin.Context) {
	app.JsonResponse(c, http.StatusOK, code.SUCCESS, "Welcome to "+vars.App.Name)
	return
}

func PingApi(c *gin.Context) {
	app.JsonResponse(c, http.StatusOK, code.SUCCESS, time.Now())
}

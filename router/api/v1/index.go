package v1

import (
	"net/http"

	"gitee.com/cristiane/micro-mall-api/pkg/app"
	"gitee.com/cristiane/micro-mall-api/pkg/code"
	"github.com/gin-gonic/gin"
)

func IndexApi(c *gin.Context) {
	app.JsonResponse(c, http.StatusOK, code.SUCCESS, code.GetMsg(code.SUCCESS), nil)
	return
}

func PingApi(c *gin.Context) {
	app.JsonResponse(c, http.StatusOK, code.SUCCESS, "PONG", nil)
}

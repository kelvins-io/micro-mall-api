package v1

import (
	"net/http"

	"gitee.com/cristiane/micro-mall-api/pkg/app"
	"gitee.com/cristiane/micro-mall-api/pkg/code"
	"github.com/gin-gonic/gin"
)

func checkUserLogin(c *gin.Context) (uid int) {
	value, exist := c.Get("uid")
	if !exist {
		app.JsonResponse(c, http.StatusOK, code.ErrorTokenEmpty, code.GetMsg(code.ErrorTokenEmpty), nil)
		return
	}
	uid, ok := value.(int)
	if !ok {
		app.JsonResponse(c, http.StatusOK, code.ErrorTokenEmpty, code.GetMsg(code.ErrorTokenEmpty), nil)
		return
	}
	return
}

func getUserId(c *gin.Context) (uid int) {
	value, exist := c.Get("uid")
	if !exist {
		return
	}
	uid, ok := value.(int)
	if !ok {
		return
	}
	return
}

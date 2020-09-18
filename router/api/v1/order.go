package v1

import (
	"gitee.com/cristiane/micro-mall-api/model/args"
	"gitee.com/cristiane/micro-mall-api/pkg/app"
	"gitee.com/cristiane/micro-mall-api/pkg/code"
	"gitee.com/cristiane/micro-mall-api/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateTradeOrderApi(c *gin.Context) {
	var uid int
	value, exist := c.Get("uid")
	if !exist {
		app.JsonResponse(c, http.StatusOK, code.ERROR_TOKEN_EMPTY, nil)
		return
	}
	uid, ok := value.(int)
	if !ok {
		app.JsonResponse(c, http.StatusOK, code.ERROR_TOKEN_EMPTY, nil)
		return
	}
	var form args.CreateTradeOrderArgs
	var err error
	err = app.BindAndValid(c, &form)
	if err != nil {
		app.JsonResponse(c, http.StatusOK, code.INVALID_PARAMS, err.Error())
		return
	}
	form.Uid = int64(uid)
	form.ClientIp = c.ClientIP()
	rsp, retCode := service.CreateTradeOrder(c, &form)
	app.JsonResponse(c, http.StatusOK, retCode, rsp)
}

func OrderTradeApi(c *gin.Context) {
	var uid int
	value, exist := c.Get("uid")
	if !exist {
		app.JsonResponse(c, http.StatusOK, code.ERROR_TOKEN_EMPTY, nil)
		return
	}
	uid, ok := value.(int)
	if !ok {
		app.JsonResponse(c, http.StatusOK, code.ERROR_TOKEN_EMPTY, nil)
		return
	}
	var form args.OrderTradeArgs
	var err error
	err = app.BindAndValid(c, &form)
	if err != nil {
		app.JsonResponse(c, http.StatusOK, code.INVALID_PARAMS, err.Error())
		return
	}
	form.OpUid = int64(uid)
	form.OpIp = c.ClientIP()
	rsp, retCode := service.OrderTrade(c, &form)
	app.JsonResponse(c, http.StatusOK, retCode, rsp)
}

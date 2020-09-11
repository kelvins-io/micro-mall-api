package v1

import (
	"gitee.com/cristiane/micro-mall-api/model/args"
	"gitee.com/cristiane/micro-mall-api/pkg/app"
	"gitee.com/cristiane/micro-mall-api/pkg/code"
	"gitee.com/cristiane/micro-mall-api/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

func SkuJoinUserTrolleyApi(c *gin.Context) {
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
	var form args.SkuJoinUserTrolleyArgs
	form.Uid = uid
	var err error
	err = app.BindAndValid(c, &form)
	if err != nil {
		app.JsonResponse(c, http.StatusOK, code.INVALID_PARAMS, err.Error())
		return
	}
	rsp, retCode := service.SkuJoinUserTrolley(c, &form)
	app.JsonResponse(c, http.StatusOK, retCode, rsp)
}

func SkuRemoveUserTrolleyApi(c *gin.Context) {
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
	var form args.SkuRemoveUserTrolleyArgs
	form.Uid = uid
	var err error
	err = app.BindAndValid(c, &form)
	if err != nil {
		app.JsonResponse(c, http.StatusOK, code.INVALID_PARAMS, err.Error())
		return
	}
	rsp, retCode := service.SkuRemoveUserTrolley(c, &form)
	app.JsonResponse(c, http.StatusOK, retCode, rsp)
}

func GetUserTrolleyListApi(c *gin.Context) {
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
	rsp, retCode := service.GetUserTrolleyList(c, int64(uid))
	app.JsonResponse(c, http.StatusOK, retCode, rsp)
}

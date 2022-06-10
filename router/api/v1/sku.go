package v1

import (
	"net/http"

	"gitee.com/cristiane/micro-mall-api/model/args"
	"gitee.com/cristiane/micro-mall-api/pkg/app"
	"gitee.com/cristiane/micro-mall-api/pkg/code"
	"gitee.com/cristiane/micro-mall-api/service"
	"github.com/gin-gonic/gin"
)

func SkuBusinessPutAwayApi(c *gin.Context) {
	uid := checkUserLogin(c)
	if uid <= 0 {
		return
	}
	var form args.SkuBusinessPutAwayArgs
	form.Uid = uid
	form.OpIp = c.ClientIP()
	var err error
	err = app.BindAndValid(c, &form)
	if err != nil {
		app.JsonResponse(c, http.StatusOK, code.InvalidParams, err.Error(), nil)
		return
	}
	rsp, retCode := service.SkuPutAway(c, &form)
	app.JsonResponse(c, http.StatusOK, retCode, code.GetMsg(retCode), rsp)
}

func SkuBusinessSupplementPropertyApi(c *gin.Context) {
	uid := checkUserLogin(c)
	if uid <= 0 {
		return
	}
	var form args.SkuPropertyExArgs
	form.Uid = uid
	form.OpIp = c.ClientIP()
	var err error
	err = app.BindAndValid(c, &form)
	if err != nil {
		app.JsonResponse(c, http.StatusOK, code.InvalidParams, err.Error(), nil)
		return
	}
	rsp, retCode := service.SkuSupplementProperty(c, &form)
	app.JsonResponse(c, http.StatusOK, retCode, code.GetMsg(retCode), rsp)
}

func GetSkuListApi(c *gin.Context) {
	var form args.GetSkuListArgs
	var err error
	err = app.BindAndValid(c, &form)
	if err != nil {
		app.JsonResponse(c, http.StatusOK, code.InvalidParams, err.Error(), nil)
		return
	}
	rsp, retCode := service.GetSkuList(c, &form)
	app.JsonResponse(c, http.StatusOK, retCode, code.GetMsg(retCode), rsp)
}

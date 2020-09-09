package v1

import (
	"gitee.com/cristiane/micro-mall-api/model/args"
	"gitee.com/cristiane/micro-mall-api/pkg/app"
	"gitee.com/cristiane/micro-mall-api/pkg/code"
	"gitee.com/cristiane/micro-mall-api/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

func MerchantsMaterialApi(c *gin.Context) {
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
	var form args.MerchantsMaterialArgs
	form.Uid = uid
	var err error
	err = app.BindAndValid(c, &form)
	if err != nil {
		app.JsonResponse(c, http.StatusOK, code.INVALID_PARAMS, err.Error())
		return
	}
	rsp, retCode := service.MerchantsMaterial(c, &form)
	app.JsonResponse(c, http.StatusOK, retCode, rsp)
}

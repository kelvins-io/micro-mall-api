package v1

import (
	"net/http"

	"gitee.com/cristiane/micro-mall-api/model/args"
	"gitee.com/cristiane/micro-mall-api/pkg/app"
	"gitee.com/cristiane/micro-mall-api/pkg/code"
	"gitee.com/cristiane/micro-mall-api/service"
	"github.com/gin-gonic/gin"
)

func ShopApplyApi(c *gin.Context) {
	uid := checkUserLogin(c)
	if uid <= 0 {
		return
	}
	var form args.ShopBusinessInfoArgs
	form.Uid = uid
	form.OpIp = c.ClientIP()
	var err error
	err = app.BindAndValid(c, &form)
	if err != nil {
		app.JsonResponse(c, http.StatusOK, code.InvalidParams, err.Error(), nil)
		return
	}
	rsp, retCode := service.ShopBusinessApply(c, &form)
	app.JsonResponse(c, http.StatusOK, retCode, code.GetMsg(retCode), rsp)
}

func ShopPledgeApi(c *gin.Context) {

}

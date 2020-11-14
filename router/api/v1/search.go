package v1

import (
	"gitee.com/cristiane/micro-mall-api/model/args"
	"gitee.com/cristiane/micro-mall-api/pkg/app"
	"gitee.com/cristiane/micro-mall-api/pkg/code"
	"gitee.com/cristiane/micro-mall-api/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

func SearchSkuInventoryApi(c *gin.Context) {
	var form args.SearchSkuInventoryArgs
	var err error
	err = app.BindAndValid(c, &form)
	if err != nil {
		app.JsonResponse(c, http.StatusOK, code.InvalidParams, err.Error())
		return
	}
	rsp, retCode := service.SearchSkuInventory(c, &form)
	app.JsonResponse(c, http.StatusOK, retCode, rsp)
}

func SearchShopApi(c *gin.Context) {
	var form args.SearchShopArgs
	var err error
	err = app.BindAndValid(c, &form)
	if err != nil {
		app.JsonResponse(c, http.StatusOK, code.InvalidParams, err.Error())
		return
	}
	rsp, retCode := service.SearchShop(c, &form)
	app.JsonResponse(c, http.StatusOK, retCode, rsp)
}

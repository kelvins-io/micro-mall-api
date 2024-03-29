package v1

import (
	"net/http"

	"gitee.com/cristiane/micro-mall-api/model/args"
	"gitee.com/cristiane/micro-mall-api/pkg/app"
	"gitee.com/cristiane/micro-mall-api/pkg/code"
	"gitee.com/cristiane/micro-mall-api/service"
	"github.com/gin-gonic/gin"
)

func SearchSkuInventoryApi(c *gin.Context) {
	var form args.SearchSkuInventoryArgs
	var err error
	err = app.BindAndValid(c, &form)
	if err != nil {
		app.JsonResponse(c, http.StatusOK, code.InvalidParams, err.Error(), nil)
		return
	}
	rsp, retCode := service.SearchSkuInventory(c, &form)
	app.JsonResponse(c, http.StatusOK, retCode, code.GetMsg(retCode), rsp)
}

func SearchShopApi(c *gin.Context) {
	var form args.SearchShopArgs
	var err error
	err = app.BindAndValid(c, &form)
	if err != nil {
		app.JsonResponse(c, http.StatusOK, code.InvalidParams, err.Error(), nil)
		return
	}
	rsp, retCode := service.SearchShop(c, &form)
	app.JsonResponse(c, http.StatusOK, retCode, code.GetMsg(retCode), rsp)
}

func SearchUserInfoApi(c *gin.Context) {
	var form args.SearchUserInfoArgs
	var err error
	err = app.BindAndValid(c, &form)
	if err != nil {
		app.JsonResponse(c, http.StatusOK, code.InvalidParams, err.Error(), nil)
		return
	}
	list, retCode := service.SearchUserInfo(c, form.Query)
	app.JsonResponse(c, http.StatusOK, retCode, code.GetMsg(retCode), list)
}

func SearchMerchantInfoApi(c *gin.Context) {
	var form args.SearchMerchantInfoArgs
	var err error
	err = app.BindAndValid(c, &form)
	if err != nil {
		app.JsonResponse(c, http.StatusOK, code.InvalidParams, err.Error(), nil)
		return
	}
	list, retCode := service.SearchMerchantInfo(c, form.Query)
	app.JsonResponse(c, http.StatusOK, retCode, code.GetMsg(retCode), list)
}

func SearchTradeOrderApi(c *gin.Context) {
	var form args.SearchTradeOrderArgs
	var err error
	err = app.BindAndValid(c, &form)
	if err != nil {
		app.JsonResponse(c, http.StatusOK, code.InvalidParams, err.Error(), nil)
		return
	}
	list, retCode := service.SearchTradeOrderInfo(c, form.Query)
	app.JsonResponse(c, http.StatusOK, retCode, code.GetMsg(retCode), list)
}

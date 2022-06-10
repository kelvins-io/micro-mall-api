package v1

import (
	"net/http"

	"gitee.com/cristiane/micro-mall-api/model/args"
	"gitee.com/cristiane/micro-mall-api/pkg/app"
	"gitee.com/cristiane/micro-mall-api/pkg/code"
	"gitee.com/cristiane/micro-mall-api/service"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func CreateTradeOrderApi(c *gin.Context) {
	uid := checkUserLogin(c)
	if uid <= 0 {
		return
	}
	var form args.CreateTradeOrderArgs
	var err error
	err = app.BindAndValid(c, &form)
	if err != nil {
		app.JsonResponse(c, http.StatusOK, code.InvalidParams, err.Error(), nil)
		return
	}
	if form.OrderTxCode == "" {
		form.OrderTxCode = uuid.New().String()
	}
	form.Uid = int64(uid)
	form.ClientIp = c.ClientIP()
	rsp, retCode := service.CreateTradeOrder(c, &form)
	app.JsonResponse(c, http.StatusOK, retCode, code.GetMsg(retCode), rsp)
}

func GenTradeOrderCodeApi(c *gin.Context) {
	uid := checkUserLogin(c)
	if uid <= 0 {
		return
	}
	rsp, retCode := service.GenOrderCode(c, int64(uid))
	app.JsonResponse(c, http.StatusOK, retCode, code.GetMsg(retCode), rsp)
}

func OrderTradeApi(c *gin.Context) {
	uid := checkUserLogin(c)
	if uid <= 0 {
		return
	}
	var form args.OrderTradeArgs
	var err error
	err = app.BindAndValid(c, &form)
	if err != nil {
		app.JsonResponse(c, http.StatusOK, code.InvalidParams, err.Error(), nil)
		return
	}
	form.OpUid = int64(uid)
	form.OpIp = c.ClientIP()
	rsp, retCode := service.OrderTrade(c, &form)
	app.JsonResponse(c, http.StatusOK, retCode, code.GetMsg(retCode), rsp)
}

func GetOrderReportApi(c *gin.Context) {
	uid := checkUserLogin(c)
	if uid <= 0 {
		return
	}
	var form args.GetOrderReportArgs
	form.Uid = int64(uid)
	var err error
	err = app.BindAndValid(c, &form)
	if err != nil {
		app.JsonResponse(c, http.StatusOK, code.InvalidParams, err.Error(), nil)
		return
	}
	rsp, retCode := service.GetOrderReport(c, &form)
	if retCode != code.SUCCESS {
		app.JsonResponse(c, http.StatusOK, retCode, code.GetMsg(retCode), rsp)
	} else {
		app.RedirectResponse(c, http.StatusMovedPermanently, rsp.ReportFilePath)
	}
}

func GetOrderShopRankApi(c *gin.Context) {
	var form args.OrderShopRankArgs
	var err error
	err = app.BindAndValid(c, &form)
	if err != nil {
		app.JsonResponse(c, http.StatusOK, code.InvalidParams, err.Error(), nil)
		return
	}
	rsp, retCode := service.OrderShopRank(c, &form)
	app.JsonResponse(c, http.StatusOK, retCode, code.GetMsg(retCode), rsp)
}

func GetOrderSkuRankApi(c *gin.Context) {
	var form args.OrderSkuRankArgs
	var err error
	err = app.BindAndValid(c, &form)
	if err != nil {
		app.JsonResponse(c, http.StatusOK, code.InvalidParams, err.Error(), nil)
		return
	}
	rsp, retCode := service.OrderSkuRank(c, &form)
	app.JsonResponse(c, http.StatusOK, retCode, code.GetMsg(retCode), rsp)
}

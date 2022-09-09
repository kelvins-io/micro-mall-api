package v1

import (
	"net/http"

	"gitee.com/cristiane/micro-mall-api/model/args"
	"gitee.com/cristiane/micro-mall-api/pkg/app"
	"gitee.com/cristiane/micro-mall-api/pkg/code"
	"gitee.com/cristiane/micro-mall-api/service"
	"github.com/gin-gonic/gin"
)

func RegisterUserApi(c *gin.Context) {
	var form args.RegisterUserArgs
	var err error
	err = app.BindAndValid(c, &form)
	if err != nil {
		app.JsonResponse(c, http.StatusOK, code.InvalidParams, err.Error(), nil)
		return
	}
	rsp, retCode := service.CreateUser(c, &form)
	app.JsonResponse(c, http.StatusOK, retCode, code.GetMsg(retCode), rsp)
}

func LoginUserWithVerifyCodeApi(c *gin.Context) {
	var form args.LoginUserWithVerifyCodeArgs
	var err error
	err = app.BindAndValid(c, &form)
	if err != nil {
		app.JsonResponse(c, http.StatusOK, code.InvalidParams, err.Error(), nil)
		return
	}
	loginInfo, retCode := service.LoginUserWithVerifyCode(c, &form)
	app.JsonResponse(c, http.StatusOK, retCode, code.GetMsg(retCode), loginInfo)
}

func LoginUserWithPhoneApi(c *gin.Context) {
	var form args.LoginUserWithPhoneArgs
	var err error
	err = app.BindAndValid(c, &form)
	if err != nil {
		app.JsonResponse(c, http.StatusOK, code.InvalidParams, err.Error(), nil)
		return
	}
	loginInfo, retCode := service.LoginUserWithPhone(c, &form)
	app.JsonResponse(c, http.StatusOK, retCode, code.GetMsg(retCode), loginInfo)
}

func LoginUserWithAccountApi(c *gin.Context) {
	var form args.LoginUserWithAccountArgs
	var err error
	err = app.BindAndValid(c, &form)
	if err != nil {
		app.JsonResponse(c, http.StatusOK, code.InvalidParams, err.Error(), nil)
		return
	}
	loginInfo, retCode := service.LoginUserWithAccount(c, &form)
	app.JsonResponse(c, http.StatusOK, retCode, code.GetMsg(retCode), loginInfo)
}

func GetVerifyCodeApi(c *gin.Context) {
	var form args.GenVerifyCodeArgs
	uid := getUserId(c)
	form.Uid = uid
	var err error
	err = app.BindAndValid(c, &form)
	if err != nil {
		app.JsonResponse(c, http.StatusOK, code.InvalidParams, err.Error(), nil)
		return
	}
	retCode := service.GenVerifyCode(c, &form)
	app.JsonResponse(c, http.StatusOK, retCode, code.GetMsg(retCode), nil)
}

func PasswordResetApi(c *gin.Context) {
	uid := checkUserLogin(c)
	if uid <= 0 {
		return
	}
	var form args.PasswordResetArgs
	form.Uid = uid
	var err error
	err = app.BindAndValid(c, &form)
	if err != nil {
		app.JsonResponse(c, http.StatusOK, code.InvalidParams, err.Error(), nil)
		return
	}

	retCode := service.PasswordReset(c, &form)
	app.JsonResponse(c, http.StatusOK, retCode, code.GetMsg(retCode), nil)
}

func GetUserInfoApi(c *gin.Context) {
	uid := checkUserLogin(c)
	if uid <= 0 {
		return
	}
	userInfo, retCode := service.GetUserInfo(c, uid)
	app.JsonResponse(c, http.StatusOK, retCode, code.GetMsg(retCode), userInfo)
}

func ListUserInfoApi(c *gin.Context) {
	uid := checkUserLogin(c)
	if uid <= 0 {
		return
	}
	var form args.ListUserInfoArgs
	var err error
	err = app.BindAndValid(c, &form)
	if err != nil {
		app.JsonResponse(c, http.StatusOK, code.InvalidParams, err.Error(), nil)
		return
	}

	userInfoList, retCode := service.ListUserInfo(c, &form)
	app.JsonResponse(c, http.StatusOK, retCode, code.GetMsg(retCode), userInfoList)
}

func LoadBalanceTestApi(c *gin.Context) {
	query, _ := c.GetQuery("query")
	result, retCode := service.LoadBalanceTest(c, query)
	app.JsonResponse(c, http.StatusOK, retCode, code.GetMsg(retCode), result)
}

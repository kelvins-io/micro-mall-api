package v1

import (
	"gitee.com/cristiane/micro-mall-api/model/args"
	"gitee.com/cristiane/micro-mall-api/pkg/app"
	"gitee.com/cristiane/micro-mall-api/pkg/code"
	"gitee.com/cristiane/micro-mall-api/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

func RegisterUserApi(c *gin.Context) {
	var form args.RegisterUserArgs
	var err error
	err = app.BindAndValid(c, &form)
	if err != nil {
		app.JsonResponse(c, http.StatusOK, code.INVALID_PARAMS, err.Error())
		return
	}

	retCode := service.CreateUser(c, &form)
	app.JsonResponse(c, http.StatusOK, retCode, code.GetMsg(retCode))
}

func LoginUserWithVerifyCodeApi(c *gin.Context) {
	var form args.LoginUserWithVerifyCodeArgs
	var err error
	err = app.BindAndValid(c, &form)
	if err != nil {
		app.JsonResponse(c, http.StatusOK, code.INVALID_PARAMS, err.Error())
		return
	}
	token, retCode := service.LoginUserWithVerifyCode(c, &form)
	if retCode != code.SUCCESS {
		app.JsonResponse(c, http.StatusOK, retCode, code.GetMsg(retCode))
		return
	}
	c.Writer.Header().Add("token", token)
	app.JsonResponse(c, http.StatusOK, code.SUCCESS, token)
}

func LoginUserWithPwdApi(c *gin.Context) {
	var form args.LoginUserWithPwdArgs
	var err error
	err = app.BindAndValid(c, &form)
	if err != nil {
		app.JsonResponse(c, http.StatusOK, code.INVALID_PARAMS, err.Error())
		return
	}
	token, retCode := service.LoginUserWithPwd(c, &form)
	if retCode != code.SUCCESS {
		app.JsonResponse(c, http.StatusOK, retCode, code.GetMsg(retCode))
		return
	}
	c.Writer.Header().Add("token", token)
	app.JsonResponse(c, http.StatusOK, code.SUCCESS, token)
}

func GetVerifyCodeApi(c *gin.Context) {
	var form args.GenVerifyCodeArgs
	var err error
	err = app.BindAndValid(c, &form)
	if err != nil {
		app.JsonResponse(c, http.StatusOK, code.INVALID_PARAMS, err.Error())
		return
	}
	retCode := service.GenVerifyCode(c, &form)
	app.JsonResponse(c, http.StatusOK, retCode, code.GetMsg(retCode))
}

func PasswordResetApi(c *gin.Context) {
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
	var form args.PasswordResetArgs
	form.Uid = uid
	var err error
	err = app.BindAndValid(c, &form)
	if err != nil {
		app.JsonResponse(c, http.StatusOK, code.INVALID_PARAMS, err.Error())
		return
	}

	retCode := service.PasswordReset(c, &form)
	app.JsonResponse(c, http.StatusOK, retCode, code.GetMsg(retCode))
}

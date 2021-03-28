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
		app.JsonResponse(c, http.StatusOK, code.InvalidParams, err.Error())
		return
	}
	rsp, retCode := service.CreateUser(c, &form)
	app.JsonResponse(c, http.StatusOK, retCode, rsp)
}

func LoginUserWithVerifyCodeApi(c *gin.Context) {
	var form args.LoginUserWithVerifyCodeArgs
	var err error
	err = app.BindAndValid(c, &form)
	if err != nil {
		app.JsonResponse(c, http.StatusOK, code.InvalidParams, err.Error())
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
		app.JsonResponse(c, http.StatusOK, code.InvalidParams, err.Error())
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
		app.JsonResponse(c, http.StatusOK, code.InvalidParams, err.Error())
		return
	}
	retCode, verifyCode := service.GenVerifyCode(c, &form)
	app.JsonResponse(c, http.StatusOK, retCode, verifyCode)
}

func PasswordResetApi(c *gin.Context) {
	var uid int
	value, exist := c.Get("uid")
	if !exist {
		app.JsonResponse(c, http.StatusOK, code.ErrorTokenEmpty, nil)
		return
	}
	uid, ok := value.(int)
	if !ok {
		app.JsonResponse(c, http.StatusOK, code.ErrorTokenEmpty, nil)
		return
	}
	var form args.PasswordResetArgs
	form.Uid = uid
	var err error
	err = app.BindAndValid(c, &form)
	if err != nil {
		app.JsonResponse(c, http.StatusOK, code.InvalidParams, err.Error())
		return
	}

	retCode := service.PasswordReset(c, &form)
	app.JsonResponse(c, http.StatusOK, retCode, code.GetMsg(retCode))
}

func GetUserInfoApi(c *gin.Context) {
	var uid int
	value, exist := c.Get("uid")
	if !exist {
		app.JsonResponse(c, http.StatusOK, code.ErrorTokenEmpty, nil)
		return
	}
	uid, ok := value.(int)
	if !ok {
		app.JsonResponse(c, http.StatusOK, code.ErrorTokenEmpty, nil)
		return
	}

	userInfo, retCode := service.GetUserInfo(c, uid)
	app.JsonResponse(c, http.StatusOK, retCode, userInfo)
}

func ListUserInfoApi(c *gin.Context) {
	_, exist := c.Get("uid")
	if !exist {
		app.JsonResponse(c, http.StatusOK, code.ErrorTokenEmpty, nil)
		return
	}
	var form args.ListUserInfoArgs
	var err error
	err = app.BindAndValid(c, &form)
	if err != nil {
		app.JsonResponse(c, http.StatusOK, code.InvalidParams, err.Error())
		return
	}

	userInfoList, retCode := service.ListUserInfo(c, &form)
	app.JsonResponse(c, http.StatusOK, retCode, userInfoList)
}

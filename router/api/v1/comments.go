package v1

import (
	"gitee.com/cristiane/micro-mall-api/model/args"
	"gitee.com/cristiane/micro-mall-api/pkg/app"
	"gitee.com/cristiane/micro-mall-api/pkg/code"
	"gitee.com/cristiane/micro-mall-api/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateOrderCommentsApi(c *gin.Context) {
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
	var form args.CreateOrderCommentsArgs
	form.Uid = int64(uid)
	var err error
	err = app.BindAndValid(c, &form)
	if err != nil {
		app.JsonResponse(c, http.StatusOK, code.InvalidParams, err.Error())
		return
	}
	retCode := service.CreateOrderComments(c, &form)
	app.JsonResponse(c, http.StatusOK, retCode, "")
}

func GetShopCommentsListApi(c *gin.Context) {
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
	var form args.GetShopCommentsListArgs
	form.Uid = int64(uid)
	var err error
	err = app.BindAndValid(c, &form)
	if err != nil {
		app.JsonResponse(c, http.StatusOK, code.InvalidParams, err.Error())
		return
	}
	rsp, retCode := service.GetOrderCommentsList(c, &form)
	app.JsonResponse(c, http.StatusOK, retCode, rsp)
}

func ModifyCommentsTagsApi(c *gin.Context) {
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
	var form args.ModifyCommentsTagsArgs
	form.Uid = int64(uid)
	var err error
	err = app.BindAndValid(c, &form)
	if err != nil {
		app.JsonResponse(c, http.StatusOK, code.InvalidParams, err.Error())
		return
	}
	retCode := service.ModifyCommentsTags(c, &form)
	app.JsonResponse(c, http.StatusOK, retCode, "")
}

func GetCommentsTagsListApi(c *gin.Context) {
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
	var form args.GetCommentsTagsListArgs
	form.Uid = int64(uid)
	var err error
	err = app.BindAndValid(c, &form)
	if err != nil {
		app.JsonResponse(c, http.StatusOK, code.InvalidParams, err.Error())
		return
	}
	rsp, retCode := service.GetCommentsTagsList(c, &form)
	app.JsonResponse(c, http.StatusOK, retCode, rsp)
}

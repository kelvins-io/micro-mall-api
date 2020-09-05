package app

import (
	"context"
	"gitee.com/cristiane/micro-mall-api/pkg/code"
	"gitee.com/cristiane/micro-mall-api/vars"
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
)

func MarkErrors(ctx context.Context, errors []*validation.Error) {
	for _, err := range errors {
		vars.AccessLogger.Error(ctx, err.Key, err.Message)
	}
	return
}

func JsonResponse(ctx *gin.Context, httpCode, retCode int, data interface{}) {
	ctx.JSON(httpCode, gin.H{
		"code": retCode,
		"msg":  code.GetMsg(retCode),
		"data": data,
	})
}

func ProtoBufResponse(ctx *gin.Context, httpCode int, data interface{}) {
	ctx.ProtoBuf(httpCode, data)
}

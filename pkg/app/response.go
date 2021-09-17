package app

import (
	"gitee.com/cristiane/micro-mall-api/pkg/code"
	"github.com/gin-gonic/gin"
)

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

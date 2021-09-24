package app

import (
	"fmt"
	varsInternal "gitee.com/cristiane/micro-mall-api/internal/vars"
	"gitee.com/cristiane/micro-mall-api/pkg/code"
	"github.com/gin-gonic/gin"
	"time"
)

func JsonResponse(ctx *gin.Context, httpCode, retCode int, data interface{}) {
	echoStatics(ctx)
	ctx.JSON(httpCode, gin.H{
		"code": retCode,
		"msg":  code.GetMsg(retCode),
		"data": data,
	})
}

func ProtoBufResponse(ctx *gin.Context, httpCode int, data interface{}) {
	echoStatics(ctx)
	ctx.ProtoBuf(httpCode, data)
}

func echoStatics(ctx *gin.Context) {
	outcomeTime := time.Now()
	ctx.Header("X-Response-Time", outcomeTime.Format(varsInternal.ResponseTimeLayout))
	incomeTime, exist := ctx.Get(varsInternal.AccessIncomeTimeKey)
	if exist {
		handleTime := outcomeTime.Sub(incomeTime.(time.Time)).Seconds()
		ctx.Header("X-Handle-Time", fmt.Sprintf("%f/s", handleTime))
	}
}

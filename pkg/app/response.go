package app

import (
	"fmt"
	"time"

	varsInternal "gitee.com/cristiane/micro-mall-api/internal/vars"
	"github.com/gin-gonic/gin"
)

func JsonResponse(ctx *gin.Context, httpCode, retCode int, msg string, data interface{}) {
	echoStatics(ctx)
	ctx.JSON(httpCode, gin.H{
		"code": retCode,
		"msg":  msg,
		"data": data,
	})
}

func ProtoBufResponse(ctx *gin.Context, httpCode int, data interface{}) {
	echoStatics(ctx)
	ctx.ProtoBuf(httpCode, data)
}

func RedirectResponse(ctx *gin.Context, httpCode int, location string) {
	echoStatics(ctx)
	ctx.Redirect(httpCode, location)
}

func HtmlResponse(ctx *gin.Context, httpCode int, data string) {
	echoStatics(ctx)
	ctx.String(httpCode, data)
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

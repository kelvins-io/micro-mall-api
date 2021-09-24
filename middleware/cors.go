package middleware

import (
	"fmt"
	varsInternal "gitee.com/cristiane/micro-mall-api/internal/vars"
	"gitee.com/cristiane/micro-mall-api/vars"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
	"strings"
	"time"
)

func Cors() gin.HandlerFunc {
	// 跨域处理
	return func(c *gin.Context) {
		method := c.Request.Method
		origin := c.Request.Header.Get("Origin")
		var headerKeys []string
		for k, _ := range c.Request.Header {
			headerKeys = append(headerKeys, k)
		}
		headerStr := strings.Join(headerKeys, ", ")
		if headerStr != "" {
			headerStr = fmt.Sprintf("access-control-allow-origin, access-control-allow-headers, %s", headerStr)
		} else {
			headerStr = "access-control-allow-origin, access-control-allow-headers"
		}
		if origin != "" {
			c.Header("Access-Control-Allow-Origin", "*")
			c.Header("Access-Control-Allow-Headers", headerStr)
			c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
			c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type")
			c.Header("Access-Control-Allow-Credentials", "true")
			c.Header("Content-Type", "application/json")
		}
		requestId := uuid.New().String()
		c.Header("X-Request-Id", requestId)
		c.Set("X-Request-Id", requestId)
		c.Header("X-Powered-By", "web_gin_template/gin "+vars.Version)
		incomeTime := time.Now()
		c.Set(varsInternal.AccessIncomeTimeKey, incomeTime)
		// 放行所有OPTIONS方法
		if method == "OPTIONS" {
			c.JSON(http.StatusOK, "Options Request!")
		}
		c.Next()
	}
}

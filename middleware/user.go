package middleware

import (
	"gitee.com/cristiane/micro-mall-api/pkg/app"
	"gitee.com/cristiane/micro-mall-api/pkg/code"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func CheckUserToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("token")
		if token == "" {
			app.JsonResponse(c, http.StatusUnauthorized, code.ERROR_TOKEN_EMPTY, code.GetMsg(code.ERROR_TOKEN_EMPTY))
			c.Abort()
			return
		}
		claims, err := ParseToken(token)
		if err != nil {
			app.JsonResponse(c, http.StatusUnauthorized, code.ERROR_TOKEN_INVALID, code.GetMsg(code.ERROR_TOKEN_INVALID))
			c.Abort()
			return
		} else if time.Now().Unix() > claims.ExpiresAt {
			app.JsonResponse(c, http.StatusUnauthorized, code.ERROR_TOKEN_EXPIRE, code.GetMsg(code.ERROR_TOKEN_EXPIRE))
			c.Abort()
			return
		}
		if claims == nil || claims.Uid == 0 {
			app.JsonResponse(c, http.StatusUnauthorized, code.ERROR_USER_NOT_EXIST, code.GetMsg(code.ERROR_USER_NOT_EXIST))
			c.Abort()
			return
		}

		c.Set("uid", claims.Uid)
		c.Next()
	}
}

package middleware

import (
	"gitee.com/cristiane/micro-mall-api/pkg/app"
	"gitee.com/cristiane/micro-mall-api/pkg/code"
	"gitee.com/cristiane/micro-mall-api/pkg/util"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func CheckUserToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("token")
		if token == "" {
			app.JsonResponse(c, http.StatusOK, code.ErrorTokenEmpty, code.GetMsg(code.ErrorTokenEmpty))
			c.Abort()
			return
		}
		claims, err := util.ParseToken(token)
		if err != nil {
			app.JsonResponse(c, http.StatusOK, code.ErrorTokenInvalid, code.GetMsg(code.ErrorTokenInvalid))
			c.Abort()
			return
		} else if claims == nil || claims.Uid == 0 {
			app.JsonResponse(c, http.StatusOK, code.ErrorUserNotExist, code.GetMsg(code.ErrorUserNotExist))
			c.Abort()
			return
		} else if time.Now().Unix() > claims.ExpiresAt {
			app.JsonResponse(c, http.StatusOK, code.ErrorTokenExpire, code.GetMsg(code.ErrorTokenExpire))
			c.Abort()
			return
		}

		c.Set("uid", claims.Uid)
		c.Next()
	}
}

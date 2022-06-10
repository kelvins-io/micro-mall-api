package middleware

import (
	"net/http"
	"time"

	"gitee.com/cristiane/micro-mall-api/pkg/app"
	"gitee.com/cristiane/micro-mall-api/pkg/code"
	"gitee.com/cristiane/micro-mall-api/pkg/util"
	"gitee.com/cristiane/micro-mall-api/service"
	"github.com/gin-gonic/gin"
)

func CheckUserToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("token")
		if token == "" {
			token = c.Query("token")
		}
		if token == "" {
			app.JsonResponse(c, http.StatusUnauthorized, code.ErrorTokenEmpty, code.GetMsg(code.ErrorTokenEmpty), nil)
			c.Abort()
			return
		}
		claims, err := util.ParseToken(token)
		if err != nil {
			app.JsonResponse(c, http.StatusForbidden, code.ErrorTokenInvalid, code.GetMsg(code.ErrorTokenInvalid), nil)
			c.Abort()
			return
		} else if claims == nil || claims.Uid == 0 {
			app.JsonResponse(c, http.StatusForbidden, code.ErrorUserNotExist, code.GetMsg(code.ErrorUserNotExist), nil)
			c.Abort()
			return
		} else if time.Now().Unix() > claims.ExpiresAt {
			app.JsonResponse(c, http.StatusForbidden, code.ErrorTokenExpire, code.GetMsg(code.ErrorTokenExpire), nil)
			c.Abort()
			return
		}

		// 校验用户状态
		retCode := service.VerifyUserState(c, int64(claims.Uid))
		if retCode != code.SUCCESS {
			app.JsonResponse(c, http.StatusForbidden, code.ErrUserStateNotVerify, code.GetMsg(code.ErrUserStateNotVerify), nil)
			c.Abort()
			return
		}

		c.Set("uid", claims.Uid)
		c.Next()
	}
}

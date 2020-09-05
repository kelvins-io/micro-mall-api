package router

import (
	"gitee.com/cristiane/micro-mall-api/middleware"
	v1 "gitee.com/cristiane/micro-mall-api/router/api/v1"
	"gitee.com/cristiane/micro-mall-api/vars"
	"io"
	"os"

	"github.com/gin-gonic/gin"
)

func InitRouter(accessInfoLogger, accessErrLogger io.Writer) *gin.Engine {

	gin.DefaultWriter = io.MultiWriter(os.Stdout, accessInfoLogger)
	gin.DefaultErrorWriter = io.MultiWriter(os.Stderr, accessErrLogger)

	gin.SetMode(gin.ReleaseMode) // 默认生产
	if vars.ServerSetting != nil && vars.ServerSetting.Mode != "" {
		gin.SetMode(vars.ServerSetting.Mode)
	}

	r := gin.Default()
	r.Use(middleware.Cors())

	r.GET("/", v1.IndexApi)
	r.GET("/ping", v1.PingApi)
	apiG := r.Group("/api")
	apiV1 := apiG.Group("/v1")
	apiCommon := apiV1.Group("/common")
	{
		apiCommon.POST("/verify_code/send", v1.GetVerifyCodeApi)
		apiCommon.POST("/register", v1.RegisterUserApi)
		apiCommon.POST("/login/verify_code", v1.LoginUserWithVerifyCodeApi)
		apiCommon.POST("/login/pwd", v1.LoginUserWithPwdApi)
	}
	apiUser := apiV1.Group("/user")
	apiUser.Use(middleware.CheckUserToken())
	{
		apiUser.PUT("/password/reset", v1.PasswordResetApi)
		apiMerchants := apiUser.Group("/merchants")
		{
			apiMerchants.PUT("/material", v1.MerchantsMaterialApi) // 商户提交材料
		}
		apiShop := apiUser.Group("/shop")
		{
			apiShop.POST("/apply", v1.ShopApplyApi)  // 申请店铺
			apiShop.PUT("/pledge", v1.ShopPledgeApi) // 店铺质押，交保证金
		}
	}

	return r
}

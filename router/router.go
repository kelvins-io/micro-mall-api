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
		apiUser.GET("/user_info", v1.GetUserInfoApi)
		apiMerchants := apiUser.Group("/merchants")
		{
			apiMerchants.PUT("/material", v1.MerchantsMaterialApi) // 商户提交材料
		}
		apiTrolley := apiUser.Group("/trolley")
		{
			apiTrolley.PUT("/sku/join", v1.SkuJoinUserTrolleyApi)        // 商品添加到购物车
			apiTrolley.DELETE("/sku/remove", v1.SkuRemoveUserTrolleyApi) // 从购物车移除商品
			apiTrolley.GET("/sku/list", v1.GetUserTrolleyListApi)        // 获取用户购物车中商品列表
		}
	}

	apiShopBusiness := apiV1.Group("/shop_business")
	apiShopBusiness.Use(middleware.CheckUserToken())
	{
		apiShop := apiShopBusiness.Group("/shop")
		{
			apiShop.POST("/apply", v1.ShopApplyApi)  // 申请店铺
			apiShop.PUT("/pledge", v1.ShopPledgeApi) // 店铺质押，交保证金
		}
	}

	apiSkuBusiness := apiV1.Group("/sku_business")
	apiSkuBusiness.Use(middleware.CheckUserToken())
	{
		apiSku := apiSkuBusiness.Group("/sku")
		{
			apiSku.POST("/put_away", v1.SkuBusinessPutAwayApi)             // 上架商品
			apiSku.PUT("/supplement", v1.SkuBusinessSupplementPropertyApi) // 补充商品属性
			apiSku.GET("/list", v1.GetSkuListApi)                          // 获取sku
		}
	}

	apiOrder := apiV1.Group("/order")
	apiOrder.Use(middleware.CheckUserToken())
	{
		apiOrder.POST("/create", v1.CreateTradeOrderApi)
	}

	return r
}

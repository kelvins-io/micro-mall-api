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
	r.GET("/ping", v1.PingApi) // ping
	apiG := r.Group("/api")
	apiV1 := apiG.Group("/v1")
	apiV1.POST("/verify_code/send", v1.GetVerifyCodeApi)            // 验证码发送
	apiV1.POST("/register", v1.RegisterUserApi)                     // 注册
	apiV1.POST("/login/verify_code", v1.LoginUserWithVerifyCodeApi) // 验证码登陆
	apiV1.POST("/login/pwd", v1.LoginUserWithPwdApi)                // 密码登陆
	apiUser := apiV1.Group("/user")
	apiUser.Use(middleware.CheckUserToken())
	{
		apiUser.PUT("/password/reset", v1.PasswordResetApi) // 重置密码
		apiUser.GET("/user_info", v1.GetUserInfoApi)        // 获取用户信息
		userSetting := apiUser.Group("/setting")
		{
			userSetting.POST("/address", v1.UserSettingAddressModifyApi) // 更新用户地址
			userSetting.GET("/address", v1.UserSettingAddressGetApi)     // 获取用户地址
		}
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
		apiShopBusiness := apiUser.Group("/shop_business")
		{
			apiShop := apiShopBusiness.Group("/shop")
			{
				apiShop.POST("/apply", v1.ShopApplyApi)  // 申请店铺
				apiShop.PUT("/pledge", v1.ShopPledgeApi) // 店铺质押，交保证金
			}
		}
		apiSkuBusiness := apiUser.Group("/sku_business")
		{
			apiSku := apiSkuBusiness.Group("/sku")
			{
				apiSku.POST("/put_away", v1.SkuBusinessPutAwayApi)             // 上架商品
				apiSku.PUT("/supplement", v1.SkuBusinessSupplementPropertyApi) // 补充商品属性
				apiSku.GET("/list", v1.GetSkuListApi)                          // 获取sku
			}
		}
		apiOrder := apiUser.Group("/order")
		{
			apiOrder.GET("/code/gen", v1.GenTradeOrderCodeApi) // 生成订单号
			apiOrder.POST("/create", v1.CreateTradeOrderApi)   // 生成订单
			apiOrder.POST("/trade", v1.OrderTradeApi)          // 订单支付
		}
		apiLogistics := apiUser.Group("/logistics")
		{
			apiLogistics.POST("/apply", v1.ApplyLogisticsApi)               // 申请物流
			apiLogistics.GET("/record/query", v1.QueryLogisticsRecordApi)   // 查询物流
			apiLogistics.PUT("/record/update", v1.UpdateLogisticsRecordApi) // 更新物流
		}
	}
	search := apiV1.Group("/search")
	{
		search.GET("/sku_inventory", v1.SearchSkuInventoryApi) // 搜索商品库存
		search.GET("/shop", v1.SearchShopApi)                  // 搜索店铺
	}

	return r
}

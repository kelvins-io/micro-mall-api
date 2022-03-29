package router

import (
	"context"
	"io"
	"os"

	"gitee.com/cristiane/micro-mall-api/internal/config"
	"gitee.com/cristiane/micro-mall-api/middleware"
	v1 "gitee.com/cristiane/micro-mall-api/router/api/v1"
	"gitee.com/cristiane/micro-mall-api/router/process"
	"gitee.com/cristiane/micro-mall-api/vars"
	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	ginEngineInit()

	r := gin.Default()
	r.Use(middleware.Cors())
	if vars.RateLimitSetting != nil && vars.RateLimitSetting.MaxConcurrent > 0 {
		r.Use(middleware.RateLimit(vars.RateLimitSetting.MaxConcurrent))
	}
	r.GET("/", v1.IndexApi)
	pprof.Register(r, "/debug")
	r.GET("/debug/metrics", process.MetricsApi)
	r.GET("/ping", v1.PingApi) // ping
	r.Static("/static", "./static")
	apiG := r.Group("/api")
	apiV1 := apiG.Group("/v1")
	apiV1.POST("/verify_code/send", v1.GetVerifyCodeApi) // 验证码发送
	apiV1.POST("/register", v1.RegisterUserApi)          // 注册
	apiUserLogin := apiV1.Group("/login")
	{
		apiUserLogin.POST("/verify_code", v1.LoginUserWithVerifyCodeApi) // 验证码登陆
		apiUserLogin.POST("/pwd", v1.LoginUserWithPwdApi)                // 密码登陆
	}
	apiUser := apiV1.Group("/user")
	apiUser.Use(middleware.CheckUserToken())
	{
		apiUser.PUT("/password/reset", v1.PasswordResetApi) // 重置密码
		apiUser.GET("/user_info", v1.GetUserInfoApi)        // 获取用户信息
		apiUser.GET("/user_info/list", v1.ListUserInfoApi)  // 列举用户
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
			apiOrder.POST("/report", v1.GetOrderReportApi)     // 获取订单列表
			apiOrder.GET("/rank/shop", v1.GetOrderShopRankApi) // 订单店铺业绩排行
			apiOrder.GET("/rank/sku", v1.GetOrderSkuRankApi)   // 订单商品排行
		}
		apiLogistics := apiUser.Group("/logistics")
		{
			apiLogistics.POST("/apply", v1.ApplyLogisticsApi)               // 申请物流
			apiLogistics.GET("/record/query", v1.QueryLogisticsRecordApi)   // 查询物流
			apiLogistics.PUT("/record/update", v1.UpdateLogisticsRecordApi) // 更新物流
		}
		apiAccount := apiUser.Group("/account")
		{
			apiAccount.PUT("/charge", v1.UserAccountChargeApi) // 账户充值
		}
		apiComments := apiUser.Group("/comments")
		{
			apiComments.POST("/tags/modify", v1.ModifyCommentsTagsApi)   // 修改评论标签
			apiComments.GET("/shop/list", v1.GetShopCommentsListApi)     // 获取评论标签
			apiComments.POST("/order/create", v1.CreateOrderCommentsApi) // 创建订单评论
			apiComments.GET("/tags/list", v1.GetCommentsTagsListApi)     // 获取店铺评论列表
		}
	}
	search := apiV1.Group("/search")
	{
		search.GET("/sku_inventory", v1.SearchSkuInventoryApi) // 搜索商品库存
		search.GET("/shop", v1.SearchShopApi)                  // 搜索店铺
		search.GET("/user_info", v1.SearchUserInfoApi)         // 搜索用户
		search.GET("/merchant_info", v1.SearchMerchantInfoApi) // 商户搜索
		search.GET("/trade_order", v1.SearchTradeOrderApi)     // 订单搜索
	}

	return r
}

func ginEngineInit() {
	var errLogWriter io.Writer = &AccessErrLogger{}
	environ := vars.Environment
	if environ == config.DefaultEnvironmentDev || environ == config.DefaultEnvironmentTest {
		var accessLogWriter io.Writer = &AccessInfoLogger{}
		accessLogWriter = io.MultiWriter(accessLogWriter, os.Stdout)
		errLogWriter = io.MultiWriter(errLogWriter, os.Stdout)
		gin.DefaultWriter = accessLogWriter
	}
	gin.DefaultErrorWriter = errLogWriter

	gin.SetMode(gin.ReleaseMode) // 默认生产
	if environ != "" {
		switch environ {
		case config.DefaultEnvironmentDev:
			gin.SetMode(gin.DebugMode)
		case config.DefaultEnvironmentTest:
			gin.SetMode(gin.TestMode)
		case config.DefaultEnvironmentRelease:
			gin.SetMode(gin.ReleaseMode)
		default:
			gin.SetMode(gin.ReleaseMode)
		}
	}
}

type AccessInfoLogger struct{}

func (a *AccessInfoLogger) Write(p []byte) (n int, err error) {
	vars.AccessLogger.Infof(context.Background(), "[gin-info] %s", p)
	return 0, nil
}

type AccessErrLogger struct{}

func (a *AccessErrLogger) Write(p []byte) (n int, err error) {
	vars.AccessLogger.Errorf(context.Background(), "[gin-err] %s", p)
	return 0, nil
}

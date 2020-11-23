package client

const (
	baseUrlProd    = "https://xxx.xxx.xx/api"
	baseUrlTestAli = "http://xx.xx.xx.xx:xxx/api"
	baseUrlDev     = "http://xx.xx.xx.56:xx/api"
	baseUrlLocal   = "http://localhost:52001/api"
)

const (
	verifyCodeSend          = "/verify_code/send"
	registerUser            = "/register"
	loginUserWithVerifyCode = "/login/verify_code"
	loginUserWithPwd        = "/login/pwd"
	userPwdReset            = "/user/password/reset"
	userInfo                = "/user/user_info"
	merchantsMaterial       = "/user/merchants/material"
	shopBusinessApply       = "/user/shop_business/shop/apply"
	skuBusinessPutAway      = "/user/sku_business/sku/put_away"
	skuBusinessGetSkuList   = "/user/sku_business/sku/list"
	skuBusinessSupplement   = "/user/sku_business/sku/supplement"
	skuJoinUserTrolley      = "/user/trolley/sku/join"
	skuRemoveUserTrolley    = "/user/trolley/sku/remove"
	skuUserTrolleyList      = "/user/trolley/sku/list"
	tradeCreateOrder        = "/user/order/create"
	tradeOrderCodeGen       = "/user/order/code/gen"
	tradeOrderPay           = "/user/order/trade"
	logisticsApply          = "/user/logistics/apply"
	userSettingAddress      = "/user/setting/address"
	searchSkuInventory      = "/search/sku_inventory"
	searchShop              = "/search/shop"
	reportOrder             = "/user/order/report"
	userAccountCharge       = "/user/account/charge"
)

const (
	apiV1 = "/v1"
	apiV2 = "/v2"
)

var apiVersion = apiV1
var qToken = token_10047
var baseUrl = baseUrlLocal + apiVersion

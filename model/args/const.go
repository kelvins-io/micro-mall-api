package args

const (
	Unknown            = 0
	VerifyCodeRegister = 1
	VerifyCodeLogin    = 2
	VerifyCodePassword = 3
	VerifyCodeTemplate = "【%v】验证码 %v，用于%v，%v分钟内有效，验证码提供给其他人可能导致账号被盗，请勿泄漏，谨防被骗。"
)

const (
	UserStateEventTypeRegister  = 10010
	UserStateEventTypeLogin     = 10011
	UserStateEventTypeLogout    = 10012
	UserStateEventTypePwdModify = 10013
)

const (
	RpcServiceMicroMallUsers     = "micro-mall-users"     // 用户服务
	RpcServiceMicroMallShop      = "micro-mall-shop"      // 店铺服务
	RpcServiceMicroMallSku       = "micro-mall-sku"       // 商品服务
	RpcServiceMicroMallTrolley   = "micro-mall-trolley"   // 购物车
	RpcServiceMicroMallOrder     = "micro-mall-order"     // 订单服务
	RpcServiceMicroMallPay       = "micro-mall-pay"       // 支付服务
	RpcServiceMicroMallLogistics = "micro-mall-logistics" // 物流服务
)

const (
	CNY = 0
	USD = 1
)

var (
	VerifyCodeTypes = []int{VerifyCodeRegister, VerifyCodeLogin, VerifyCodePassword}
	CoinTypes       = []int{CNY, USD}
)

var MsgFlags = map[int]string{
	Unknown:                     "未知",
	VerifyCodeRegister:          "注册",
	VerifyCodeLogin:             "登录",
	VerifyCodePassword:          "修改/重置密码",
	UserStateEventTypeRegister:  "注册",
	UserStateEventTypePwdModify: "修改密码",
	UserStateEventTypeLogin:     "登录上线",
	UserStateEventTypeLogout:    "退出登录",
}

type CommonBusinessMsg struct {
	Type int    `json:"type"`
	Tag  string `json:"tag"`
	UUID string `json:"uuid"`
	Msg  string `json:"msg"`
}

type UserRegisterNotice struct {
	CountryCode string `json:"country_code"`
	Phone       string `json:"phone"`
	Time        string `json:"time"`
	State       int    `json:"state"`
}

type UserStateNotice struct {
	Uid  int    `json:"uid"`
	Time string `json:"time"`
}

type UserOnlineState struct {
	Uid   int    `json:"uid"`
	State string `json:"state"`
	Time  string `json:"time"`
}

type SkuInventoryInfo struct {
	SkuCode       string `json:"sku_code"`
	Name          string `json:"name"`
	Price         string `json:"price"`
	Title         string `json:"title"`
	SubTitle      string `json:"sub_title"`
	Desc          string `json:"desc"`
	Production    string `json:"production"`
	Supplier      string `json:"supplier"`
	Category      int32  `json:"category"`
	Color         string `json:"color"`
	ColorCode     int32  `json:"color_code"`
	Specification string `json:"specification"`
	DescLink      string `json:"desc_link"`
	State         int32  `json:"state"`
	Amount        int64  `json:"amount"`
	ShopId        int64  `json:"shop_id"`
	Version       int64  `json:"version"`
}

func GetMsg(code int) string {
	msg, ok := MsgFlags[code]
	if ok {
		return msg
	}
	return MsgFlags[Unknown]
}

package args

const (
	Unknown            = 0
	VerifyCodeRegister = 1
	VerifyCodeLogin    = 2
	VerifyCodePassword = 3
)

const (
	UserStateEventTypeRegister  = 10010
	UserStateEventTypeLogin     = 10011
	UserStateEventTypeLogout    = 10012
	UserStateEventTypePwdModify = 10013
)

const (
	RpcServiceMicroMallUsers     = "micro-mall-users"     // 用户服务,商户服务
	RpcServiceMicroMallShop      = "micro-mall-shop"      // 店铺服务
	RpcServiceMicroMallSku       = "micro-mall-sku"       // 商品服务
	RpcServiceMicroMallTrolley   = "micro-mall-trolley"   // 购物车
	RpcServiceMicroMallOrder     = "micro-mall-order"     // 订单服务
	RpcServiceMicroMallPay       = "micro-mall-pay"       // 支付服务
	RpcServiceMicroMallLogistics = "micro-mall-logistics" // 物流服务
	RpcServiceMicroMallComments  = "micro-mall-comments"  // 评论服务
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

type UserVerifyCode struct {
	VerifyCode string `json:"verify_code"`
	Expire     int64  `json:"expire"`
}

type CommonBusinessMsg struct {
	Type    int    `json:"type"`
	Tag     string `json:"tag"`
	UUID    string `json:"uuid"`
	Content string `json:"content"`
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

type TradeOrderDetail struct {
	UserAccount string                `json:"user_account"`
	CoinType    int                   `json:"coin_type"`
	OrderList   []TradeShopOrderEntry `json:"order_list"`
}

type TradeShopOrderEntry struct {
	OrderCode   string `json:"order_code"`
	Description string `json:"description"`
	ShopAccount string `json:"shop_account"`
	Attach      string `json:"attach"`
	Money       string `json:"money"`
	Reduction   string `json:"reduction"`
}

type TaskGroupErr struct {
	errMsg  string
	retCode int
}

func (t *TaskGroupErr) Error() string {
	return t.errMsg
}

func (t *TaskGroupErr) RetCode() int {
	return t.retCode
}

func (t *TaskGroupErr) ErrMsg() string {
	return t.errMsg
}

func NewTaskGroupErr(errMsg string, retCode int) error {
	return &TaskGroupErr{
		errMsg:  errMsg,
		retCode: retCode,
	}
}

func GetMsg(code int) string {
	msg, ok := MsgFlags[code]
	if ok {
		return msg
	}
	return MsgFlags[Unknown]
}

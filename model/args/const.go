package args

const (
	Unknown            = 0
	VerifyCodeRegister = 1
	VerifyCodeLogin    = 2
	VerifyCodePassword = 3
	VerifyCodeTemplate = "【%v】验证码 %v，用于%v，%v分钟内有效，验证码提供给其他人可能导致账号被盗，请勿泄漏，谨防被骗。"
)

const (
	TaskNameUserRegisterNotice    = "task_user_register_notice"
	TaskNameUserRegisterNoticeErr = "task_user_register_notice_err"

	TaskNameUserStateNotice    = "task_user_state_notice"
	TaskNameUserStateNoticeErr = "task_user_state_notice_err"
)

const (
	RpcServiceMicroMallUsers   = "micro-mall-users"
	RpcServiceMicroMallShop    = "micro-mall-shop"
	RpcServiceMicroMallSku     = "micro-mall-sku"
	RpcServiceMicroMallTrolley = "micro-mall-trolley"
)

const (
	CacheKeyUserSate = "user_state_"
)

const (
	UserStateEventTypeLogin     = 10010
	UserStateEventTypeLogout    = 10011
	UserStateEventTypePwdModify = 10012
)

var (
	VerifyCodeTypes = []int{VerifyCodeRegister, VerifyCodeLogin, VerifyCodePassword}
)

var MsgFlags = map[int]string{
	Unknown:                     "未知",
	VerifyCodeRegister:          "注册",
	VerifyCodeLogin:             "登录",
	VerifyCodePassword:          "修改/重置密码",
	UserStateEventTypePwdModify: "修改密码",
	UserStateEventTypeLogin:     "登录上线",
	UserStateEventTypeLogout:    "退出登录",
}

type UserRegisterNotice struct {
	CountryCode string `json:"country_code"`
	Phone       string `json:"phone"`
	Time        string `json:"time"`
	State       int    `json:"state"`
}

type UserStateNotice struct {
	Uid       int    `json:"uid"`
	EventType int    `json:"event_type"`
	Event     string `json:"event"`
	Time      string `json:"time"`
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
	// 其他属性...

	Amount int64 `json:"amount"`
	ShopId int64 `json:"shop_id"`
}

func GetMsg(code int) string {
	msg, ok := MsgFlags[code]
	if ok {
		return msg
	}
	return MsgFlags[Unknown]
}

package args

type OrderSku struct {
	OrderCode string          `json:"order_code"`
	SkuList   []OrderSkuEntry `json:"sku_list"`
}

type OrderSkuEntry struct {
	SkuCode   string `json:"sku_code"`
	Amount    int    `json:"amount"`
	Name      string `json:"name"`
	Price     string `json:"price"`
	Reduction string `json:"reduction"`
}

type OrderSkuRsp struct {
	SkuList []OrderSku `json:"sku_list"`
}

type CreateOrderRsp struct {
	TxCode string
}

type ShopOrderDetail struct {
	ShopCode    string `json:"shop_code"`
	OrderCode   string `json:"order_code"`
	TimeExpire  string `json:"time_expire"`
	Description string `json:"description"`
	Amount      string `json:"amount"`
	CoinType    int    `json:"coin_type"`
	NotifyUrl   string `json:"notify_url"`
}

type OrderDetailRsp struct {
	UserCode string            `json:"user_code"`
	CoinType int               `json:"coin_type"`
	List     []ShopOrderDetail `json:"list"`
}

const (
	RpcServiceMicroMallUsers = "micro-mall-users"
	RpcServiceMicroMallShop  = "micro-mall-shop"
	RpcServiceMicroMallSku   = "micro-mall-sku"
)

const (
	TaskNameTradeOrderNotice    = "task_trade_order_notice"
	TaskNameTradeOrderNoticeErr = "task_trade_order_notice_err"
)

const (
	ConfigKvShopOrderNotifyUrl = "shop_order_notify_url"
)

type CommonBusinessMsg struct {
	Type int    `json:"type"`
	Tag  string `json:"tag"`
	UUID string `json:"uuid"`
	Msg  string `json:"msg"`
}

type TradeOrderDetail struct {
	ShopId    int64  `json:"shop_id"`
	OrderCode string `json:"order_code"`
}

type TradeOrderNotice struct {
	Uid  int64  `json:"uid"`
	Time string `json:"time"`
	// 9-19 修改，直接通知交易号
	TxCode string `json:"tx_code"`
}

const (
	Unknown                   = 0
	TradeOrderEventTypeCreate = 10014
	TradeOrderEventTypeExpire = 10015
)

var MsgFlags = map[int]string{
	Unknown:                   "未知",
	TradeOrderEventTypeCreate: "交易订单创建",
	TradeOrderEventTypeExpire: "交易订单过期",
}

func GetMsg(code int) string {
	msg, ok := MsgFlags[code]
	if ok {
		return msg
	}
	return MsgFlags[Unknown]
}

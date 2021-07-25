package code

var MsgFlags = map[int]string{
	SUCCESS:                   "ok",
	ERROR:                     "服务器出错",
	InvalidParams:             "请求参数错误",
	IdNotEmpty:                "ID为空",
	ErrorTokenEmpty:           "用户token为空",
	ErrorTokenInvalid:         "用户token无效",
	ErrorTokenExpire:          "用户token过期",
	ErrorUserNotExist:         "用户不存在",
	ErrorUserExist:            "用户已存在",
	UserLoginNotAllow:         "用户暂时不允许登录",
	ErrorEmailSend:            "邮件发送错误",
	ErrorVerifyCodeEmpty:      "验证码为空",
	ErrorVerifyCodeInvalid:    "验证码无效",
	ErrorVerifyCodeExpire:     "验证码过期",
	ErrorVerifyCodeInterval:   "验证码仍在请求时间间隔内",
	ErrorVerifyCodeLimited:    "验证码在请求时间段达到最大限制",
	DbDuplicateEntry:          "Duplicate entry",
	ErrorUserPwd:              "用户密码错误",
	ErrorMerchantNotExist:     "商户未提交过认证资料",
	ErrorMerchantExist:        "商户认证资料已存在",
	ErrorShopBusinessExist:    "店铺认证资料已存在",
	ErrorShopBusinessNotExist: "商户未提交过店铺认证资料",
	ErrorSkuCodeExist:         "商品sku已入库",
	ErrorSkuCodeNotExist:      "商品sku未入库",
	ErrorShopIdNotExist:       "店铺ID不存在",
	ErrorShopIdExist:          "店铺ID已存在",
	ErrorInviteCodeNotExist:   "邀请码不存在",
	ErrorSkuAmountNotEnough:   "商品库存不够",
	UserBalanceNotEnough:      "用户余额不足",
	UserAccountStateLock:      "用户账户被锁定",
	UserAccountNotExist:       "用户账户不存在",
	MerchantAccountNotExist:   "商户账户不存在",
	MerchantAccountStateLock:  "商户账户被锁定",
	DecimalParseErr:           "金额格式解析错误",
	TransactionFailed:         "事务执行失败",
	TradePayExpire:            "支付时间过期",
	TxCodeNotExist:            "交易号不存在",
	TradeOrderTxCodeEmpty:     "订单事务号为空",
	TradeOrderExist:           "订单已存在",
	TradePayRun:               "订单正在支付中",
	TradePaySuccess:           "订单已完成支付",
	LogisticsRecordExist:      "物流记录已存在",
	LogisticsRecordNotExist:   "物流记录不存在",
	UserSettingInfoExist:      "用户设置信息已存在",
	UserSettingInfoNotExist:   "用户设置记录不存在",
	UserDeliveryInfoNotExist:  "用户物流收货地址不存在",
	TradeOrderNotMatchUser:    "交易订单不匹配当前用户",
	SkuPriceVersionNotExist:   "商品价格版本不存在",
	OrderStateInvalid:         "订单状态无效",
	OrderStateLock:            "订单状态被锁定",
	OrderExpire:               "订单过期",
	OrderPayCompleted:         "订单已完成支付",
	UserAccountStateInvalid:   "用户账户无效",
	CommentsTagExist:          "评论标签已存在",
	CommentsTagNotExist:       "评论标签不存在",
	UserOrderNotExist:         "用户订单不存在",
	OutTradeEmpty:             "外部交易号为空",
	UserStateNotVerify:        "用户状态未验证或审核或被锁定",
	ShopStateNotVerify:        "店铺状态未审核或被冻结",
	OrderPayIng:               "交易单号正在支付中",
}

//func init() {
//	for k, v := range MsgFlags {
//		fmt.Println(k, "\t\t", v, "\t\t\t")
//	}
//}

func GetMsg(code int) string {
	msg, ok := MsgFlags[code]
	if ok {
		return msg
	}
	return MsgFlags[ERROR]
}

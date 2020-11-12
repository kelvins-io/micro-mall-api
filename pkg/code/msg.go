package code

import "fmt"

var MsgFlags = map[int]string{
	SUCCESS:                   "ok",
	ERROR:                     "服务器出错",
	InvalidParams:             "请求参数错误",
	IdNotEmpty:                "ID为空",
	ErrorTokenEmpty:           "token为空",
	ErrorTokenInvalid:         "token无效",
	ErrorTokenExpire:          "token过期",
	ErrorUserNotExist:         "用户不存在",
	ErrorUserExist:            "用户已存在",
	UserLoginNotAllow:         "用户暂时不允许登录",
	ErrorEmailSend:            "邮件发送错误",
	ErrorVerifyCodeEmpty:      "验证码为空",
	ErrorVerifyCodeInvalid:    "验证码无效",
	ErrorVerifyCodeExpire:     "验证码过期",
	DbDuplicateEntry:          "Duplicate entry",
	ErrorUserPwd:              "用户密码错误",
	ErrorMerchantNotExist:     "商户未提交过认证资料",
	ErrorMerchantExist:        "商户认证资料已存在",
	ErrorShopBusinessExist:    "店铺认证资料已存在",
	ErrorShopBusinessNotExist: "商户未提交过店铺认证资料",
	ErrorSkuCodeExist:         "商品sku-code已存在系统",
	ErrorSkuCodeNotExist:      "商品sku-code不存在",
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
	TxcodeNotExist:            "交易号不存在",
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
}

func init() {
	for k, v := range MsgFlags {
		fmt.Println(k, "\t\t", v, "\t\t\t")
	}
}

func GetMsg(code int) string {
	msg, ok := MsgFlags[code]
	if ok {
		return msg
	}
	return MsgFlags[ERROR]
}

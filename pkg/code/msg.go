package code

import "fmt"

var MsgFlags = map[int]string{
	SUCCESS:                       "ok",
	ERROR:                         "服务器出错",
	INVALID_PARAMS:                "请求参数错误",
	ID_NOT_EMPTY:                  "ID为空",
	ERROR_TOKEN_EMPTY:             "token为空",
	ERROR_TOKEN_INVALID:           "token无效",
	ERROR_TOKEN_EXPIRE:            "token过期",
	ERROR_USER_NOT_EXIST:          "用户不存在",
	ERROR_USER_EXIST:              "用户已存在",
	ERROR_EMAIL_SEND:              "邮件发送错误",
	ERROR_VERIFY_CODE_EMPTY:       "验证码为空",
	ERROR_VERIFY_CODE_INVALID:     "验证码无效",
	ERROR_VERIFY_CODE_EXPIRE:      "验证码过期",
	DB_DUPLICATE_ENTRY:            "Duplicate entry",
	ERROR_USER_PWD:                "用户密码错误",
	ERROR_MERCHANT_NOT_EXIST:      "商户未提交过认证资料",
	ERROR_MERCHANT_EXIST:          "商户认证资料已存在",
	ERROR_SHOP_BUSINESS_EXIST:     "店铺认证资料已存在",
	ERROR_SHOP_BUSINESS_NOT_EXIST: "商户未提交过店铺认证资料",
	ERROR_SKU_CODE_EXIST:          "商品唯一code已存在系统",
	ERROR_SKU_CODE_NOT_EXIST:      "商品唯一code不存在",
	ERROR_SHOP_ID_NOT_EXIST:       "店铺ID不存在",
	ERROR_SHOP_ID_EXIST:           "店铺ID已存在",
	ERROR_INVITE_CODE_NOT_EXIST:   "邀请码不存在",
	ERROR_SKU_AMOUNT_NOT_ENOUGH:   "商品库存不够",
	USER_BALANCE_NOT_ENOUGH:       "用户余额不足",
	USER_ACCOUNT_STATE_LOCK:       "用户账户被锁定",
	USER_ACCOUNT_NOT_EXIST:        "用户账户不存在",
	MERCHANT_ACCOUNT_NOT_EXIST:    "商户账户不存在",
	MERCHANT_ACCOUNT_STATE_LOCK:   "商户账户被锁定",
	DECIMAL_PARSE_ERR:             "金额格式解析错误",
	TRANSACTION_FAILED:            "事务执行失败",
	TXCODE_NOT_EXIST:              "交易号不存在",
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

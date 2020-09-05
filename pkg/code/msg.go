package code

var MsgFlags = map[int]string{
	SUCCESS:                   "ok",
	ERROR:                     "服务器出错",
	INVALID_PARAMS:            "请求参数错误",
	ID_NOT_EMPTY:              "ID为空",
	ERROR_TOKEN_EMPTY:         "token为空",
	ERROR_TOKEN_INVALID:       "token无效",
	ERROR_TOKEN_EXPIRE:        "token过期",
	ERROR_USER_NOT_EXIST:      "用户不存在",
	ERROR_USER_EXIST:          "用户已存在",
	ERROR_EMAIL_SEND:          "邮件发送错误",
	ERROR_VERIFY_CODE_EMPTY:   "验证码为空",
	ERROR_VERIFY_CODE_INVALID: "验证码无效",
	ERROR_VERIFY_CODE_EXPIRE:  "验证码过期",
	DB_DUPLICATE_ENTRY:        "Duplicate entry",
	ERROR_USER_PWD:            "用户密码错误",
}

func GetMsg(code int) string {
	msg, ok := MsgFlags[code]
	if ok {
		return msg
	}
	return MsgFlags[ERROR]
}

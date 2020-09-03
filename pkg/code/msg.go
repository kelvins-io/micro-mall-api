package code

var MsgFlags = map[int]string{
	SUCCESS:              "ok",
	ERROR:                "服务器出错",
	INVALID_PARAMS:       "请求参数错误",
	ID_NOT_EMPTY:         "ID为空",
	ERROR_TOKEN_EMPTY:    "token为空",
	ERROR_TOKEN_INVALID:  "token无效",
	ERROR_TOKEN_EXPIRE:   "token过期",
	ERROR_USER_NOT_EXIST: "用户不存在",
}

func GetMsg(code int) string {
	msg, ok := MsgFlags[code]
	if ok {
		return msg
	}
	return MsgFlags[ERROR]
}

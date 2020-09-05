package mysql

import "time"

type VerifyCodeRecord struct {
	Id           int       `xorm:"'id' not null pk autoincr comment('自增id') INT"`
	Uid          int       `xorm:"'uid' not null comment('用户UID') INT"`
	BusinessType int       `xorm:"'business_type' comment('验证类型，1-注册登录，2-购买商品') TINYINT"`
	VerifyCode   string    `xorm:"'verify_code' comment('验证码') index CHAR(6)"`
	Expire       int       `xorm:"'expire' comment('过期时间unix') INT"`
	CountryCode  string    `xorm:"'country_code' comment('验证码下发手机国际码') index(country_code_phone_index) CHAR(5)"`
	Phone        string    `xorm:"'phone' comment('验证码下发手机号') index(country_code_phone_index) CHAR(11)"`
	Email        string    `xorm:"'email' comment('验证码下发邮箱') index VARCHAR(255)"`
	CreateTime   time.Time `xorm:"'create_time' comment('创建时间') DATETIME"`
	UpdateTime   time.Time `xorm:"'update_time' comment('修改时间') DATETIME"`
}

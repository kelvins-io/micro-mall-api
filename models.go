package models

import (
	"time"
)

type MerchantInfo struct {
	MerchantId   int       `xorm:"not null pk autoincr comment('商户号ID') INT"`
	Uid          int       `xorm:"not null comment('用户ID') unique INT"`
	RegisterAddr string    `xorm:"not null comment('注册地址') TEXT"`
	HealthCardNo string    `xorm:"not null comment('健康证号') index CHAR(30)"`
	Identity     int       `xorm:"comment('身份属性，1-临时店员，2-正式店员，3-经理，4-店长') TINYINT"`
	State        int       `xorm:"comment('状态，0-未审核，1-审核中，2-审核不通过，3-已审核') TINYINT"`
	TaxCardNo    string    `xorm:"comment('纳税账户号') index CHAR(30)"`
	CreateTime   time.Time `xorm:"default CURRENT_TIMESTAMP comment('创建时间') DATETIME"`
	UpdateTime   time.Time `xorm:"default CURRENT_TIMESTAMP comment('修改时间') DATETIME"`
}

type UserInfo struct {
	Id           int       `xorm:"not null pk autoincr comment('自增ID') INT"`
	AccountId    string    `xorm:"not null comment('账户ID，全局唯一') unique CHAR(36)"`
	UserName     string    `xorm:"not null comment('用户名') index VARCHAR(255)"`
	Password     string    `xorm:"not null comment('用户密码md5值') VARCHAR(255)"`
	PasswordSalt string    `xorm:"comment('密码salt值') VARCHAR(255)"`
	Sex          int       `xorm:"comment('性别，1-男，2-女') TINYINT"`
	Phone        string    `xorm:"comment('手机号') unique(country_code_phone_index) CHAR(11)"`
	CountryCode  string    `xorm:"comment('手机区号') unique(country_code_phone_index) CHAR(5)"`
	Email        string    `xorm:"comment('邮箱') index VARCHAR(255)"`
	State        int       `xorm:"comment('状态，0-未激活，1-审核中，2-审核未通过，3-已审核') TINYINT"`
	IdCardNo     string    `xorm:"comment('身份证号') unique CHAR(18)"`
	Inviter      int       `xorm:"comment('邀请人uid') INT"`
	InviteCode   string    `xorm:"comment('邀请码') CHAR(20)"`
	CreateTime   time.Time `xorm:"not null default CURRENT_TIMESTAMP comment('创建时间') DATETIME"`
	UpdateTime   time.Time `xorm:"not null default CURRENT_TIMESTAMP comment('修改时间') DATETIME"`
	ContactAddr  string    `xorm:"comment('联系地址') TEXT"`
	Age          int       `xorm:"comment('年龄') INT"`
}

type VerifyCodeRecord struct {
	Id           int       `xorm:"not null pk autoincr comment('自增id') INT"`
	Uid          int       `xorm:"not null comment('用户UID') INT"`
	BusinessType int       `xorm:"comment('验证类型，1-注册登录，2-购买商品') TINYINT"`
	VerifyCode   string    `xorm:"comment('验证码') index CHAR(6)"`
	Expire       int       `xorm:"comment('过期时间unix') INT"`
	CountryCode  string    `xorm:"comment('验证码下发手机国际码') index(country_code_phone_index) CHAR(5)"`
	Phone        string    `xorm:"comment('验证码下发手机号') index(country_code_phone_index) CHAR(11)"`
	Email        string    `xorm:"comment('验证码下发邮箱') index VARCHAR(255)"`
	CreateTime   time.Time `xorm:"not null default CURRENT_TIMESTAMP comment('创建时间') DATETIME"`
	UpdateTime   time.Time `xorm:"not null default CURRENT_TIMESTAMP comment('修改时间') DATETIME"`
}

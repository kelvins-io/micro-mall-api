package mysql

import (
	"database/sql"
	"time"
)

const (
	TableConfigKvStore    = "qqq"
	TableUser             = "user_info"
	TableMerchantInfo     = "merchant_info"
	TableVerifyCodeRecord = "verify_code_record"
)

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
type UserInfo struct {
	Id           int            `xorm:"'id' not null pk autoincr comment('用户id') INT"`
	AccountId    string         `xorm:"'account_id' not null comment('账户ID，全局唯一') unique CHAR(36)"`
	UserName     string         `xorm:"'user_name' not null comment('用户名') index VARCHAR(255)"`
	Password     string         `xorm:"'password' not null comment('用户密码md5值') VARCHAR(255)"`
	PasswordSalt string         `xorm:"'password_salt' comment('密码salt值') VARCHAR(255)"`
	Sex          int            `xorm:"'sex' comment('性别，1-男，2-女') TINYINT"`
	Phone        string         `xorm:"'phone' comment('手机号') unique(country_code_phone_index) CHAR(11)"`
	CountryCode  string         `xorm:"'country_code' comment('手机区号') unique(country_code_phone_index) CHAR(5)"`
	Email        string         `xorm:"'email' comment('邮箱') index VARCHAR(255)"`
	State        int            `xorm:"'state' comment('状态，0-未激活，1-审核中，2-审核未通过，3-已审核') TINYINT"`
	IdCardNo     sql.NullString `xorm:"'id_card_no' comment('身份证号') unique CHAR(18)"`
	Inviter      int            `xorm:"'inviter' comment('邀请人uid') INT"`
	InviteCode   string         `xorm:"'invite_code' comment('邀请码') CHAR(20)"`
	ContactAddr  string         `xorm:"'contact_addr' comment('联系地址') TEXT"`
	Age          int            `xorm:"'age' comment('年龄') INT"`
	CreateTime   time.Time      `xorm:"'create_time' not null comment('创建时间') DATETIME"`
	UpdateTime   time.Time      `xorm:"'update_time' not null comment('修改时间') DATETIME"`
}

type MerchantInfo struct {
	MerchantId   int       `xorm:"'merchant_id' not null pk autoincr comment('商户号ID') INT"`
	Uid          int       `xorm:"'uid' not null comment('用户ID') unique INT"`
	RegisterAddr string    `xorm:"'register_addr' not null comment('注册地址') TEXT"`
	HealthCardNo string    `xorm:"'health_card_no' not null comment('健康证号') index CHAR(30)"`
	Identity     int       `xorm:"'identity' comment('身份属性，1-临时店员，2-正式店员，3-经理，4-店长') TINYINT"`
	State        int       `xorm:"'state' comment('状态，0-未审核，1-审核中，2-审核不通过，3-已审核') TINYINT"`
	TaxCardNo    string    `xorm:"'tax_card_no' comment('纳税账户号') index CHAR(30)"`
	CreateTime   time.Time `xorm:"'create_time' default CURRENT_TIMESTAMP comment('创建时间') DATETIME"`
	UpdateTime   time.Time `xorm:"'update_time' default CURRENT_TIMESTAMP comment('修改时间') DATETIME"`
}

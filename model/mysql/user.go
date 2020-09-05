package mysql

import (
	"database/sql"
	"time"
)

type User struct {
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
	CreateTime   time.Time      `xorm:"'create_time' not null comment('创建时间') DATETIME"`
	UpdateTime   time.Time      `xorm:"'update_time' not null comment('修改时间') DATETIME"`
}

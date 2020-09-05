package args

import (
	"gitee.com/cristiane/micro-mall-api/pkg/util"
	"github.com/astaxie/beego/validation"
)

type RegisterUserArgs struct {
	UserName    string `form:"user_name" json:"user_name"`
	Password    string `form:"password" json:"password"`
	Sex         int    `form:"sex" json:"sex"`
	Email       string `form:"email" json:"email"`
	CountryCode string `form:"country_code" json:"country_code"`
	Phone       string `form:"phone" json:"phone"`
	VerifyCode  string `form:"verify_code" json:"verify_code"`
	IdCardNo    string `form:"id_card_no" json:"id_card_no"`
	Inviter     int    `form:"inviter" json:"inviter"`
}

func (t *RegisterUserArgs) Valid(v *validation.Validation) {
	if t.UserName == "" {
		v.SetError("UserName", "姓名不能为空")
	}
	if len(t.Password) < 6 {
		v.SetError("Password", "密码长度不能少于6位")
	}
	if len(t.CountryCode) < 2 {
		v.SetError("CountryCode", "国际码不能少于2位")
	}
	if len(t.Phone) < 11 {
		v.SetError("Phone", "手机号不能少于11位")
	}
	if len(t.VerifyCode) < 6 {
		v.SetError("VerifyCode", "验证码不能少于6位")
	}
	if t.IdCardNo != "" {
		if len(t.IdCardNo) != 18 {
			v.SetError("IdCardNo", "身份证号码必须18位")
		}
	}
	if t.Inviter < 0 {
		v.SetError("Inviter", "邀请人UID必须大于0")
	}
}

type LoginUserWithVerifyCodeArgs struct {
	CountryCode string `form:"country_code" json:"country_code"`
	Phone       string `form:"phone" json:"phone"`
	VerifyCode  string `form:"verify_code" json:"verify_code"`
}

func (t *LoginUserWithVerifyCodeArgs) Valid(v *validation.Validation) {
	if len(t.CountryCode) < 2 {
		v.SetError("CountryCode", "国际码不能少于2位")
	}
	if len(t.Phone) < 11 {
		v.SetError("Phone", "手机号不能少于11位")
	}
	if len(t.VerifyCode) < 6 {
		v.SetError("VerifyCode", "验证码不能少于6位")
	}
}

type LoginUserWithPwdArgs struct {
	CountryCode string `form:"country_code" json:"country_code"`
	Phone       string `form:"phone" json:"phone"`
	Password    string `form:"password" json:"password"`
}

func (t *LoginUserWithPwdArgs) Valid(v *validation.Validation) {
	if len(t.Password) < 6 {
		v.SetError("Password", "密码长度不能少于6位")
	}
	if len(t.CountryCode) < 2 {
		v.SetError("CountryCode", "国际码不能少于2位")
	}
	if len(t.Phone) < 11 {
		v.SetError("Phone", "手机号不能少于11位")
	}
}

type GenVerifyCodeArgs struct {
	CountryCode  string `form:"country_code" json:"country_code"`
	Phone        string `form:"phone" json:"phone"`
	BusinessType int    `form:"business_type" json:"business_type"`
	ReceiveEmail string `form:"receive_email" json:"receive_email"`
}

func (t *GenVerifyCodeArgs) Valid(v *validation.Validation) {
	if len(t.CountryCode) < 2 {
		v.SetError("CountryCode", "国际码不能少于2位")
	}
	if len(t.Phone) < 11 {
		v.SetError("Phone", "手机号不能少于11位")
	}
	if !util.IntSliceContainsItem(VerifyCodeTypes, t.BusinessType) {
		v.SetError("BusinessType", "不支持的获取验证码类型")
	}
}

type PasswordResetArgs struct {
	Uid        int    `json:"uid"`
	VerifyCode string `form:"verify_code" json:"verify_code"`
	Password   string `form:"password" json:"password"`
}

func (t *PasswordResetArgs) Valid(v *validation.Validation) {
	if len(t.Password) < 6 {
		v.SetError("Password", "密码长度不能少于6位")
	}
	if len(t.VerifyCode) < 6 {
		v.SetError("VerifyCode", "验证码不能少于6位")
	}
}

// 商户申请
type MerchantsCertificationApplyArgs struct {
}

type UserInfoRsp struct {
	Id          int    `json:"id"`
	AccountId   string `json:"account_id"`
	UserName    string `json:"user_name"`
	Sex         int    `json:"sex"`
	Phone       string `json:"phone"`
	CountryCode string `json:"country_code"`
	Email       string `json:"email"`
	State       int    `json:"state"`
	IdCardNo    string `json:"id_card_no"`
	Inviter     int    `json:"inviter"`
	InviteCode  string `json:"invite_code"`
	ContactAddr string `json:"contact_addr"`
	Age         int    `json:"age"`
	CreateTime  string `json:"create_time"`
	UpdateTime  string `json:"update_time"`
}

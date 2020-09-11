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
type MerchantsMaterialArgs struct {
	OperationType int    `form:"operation_type" json:"operation_type"`
	Uid           int    `json:"uid"`
	RegisterAddr  string `form:"register_addr" json:"register_addr"`
	HealthCardNo  string `form:"health_card_no" json:"health_card_no"`
	Identity      int    `form:"identity" json:"identity"`
	TaxCardNo     string `form:"tax_card_no" json:"tax_card_no"`
}

func (t *MerchantsMaterialArgs) Valid(v *validation.Validation) {
	if len(t.RegisterAddr) < 1 {
		v.SetError("RegisterAddr", "注册地址不能为空")
	}
	if len(t.HealthCardNo) > 30 {
		v.SetError("HealthCardNo", "健康证号需要小于30位")
	}
	if len(t.HealthCardNo) < 10 {
		v.SetError("HealthCardNo", "健康证号需要大于等于10位")
	}
	if !util.IntSliceContainsItem([]int{1, 2, 3, 4, 5}, t.Identity) {
		v.SetError("Identity", "身份属性必须大于0,小于5")
	}
	if len(t.TaxCardNo) < 15 {
		v.SetError("TaxCardNo", "商户纳税号不能小于15位")
	}
	if !util.IntSliceContainsItem([]int{0, 1, 2, 3}, t.OperationType) {
		v.SetError("OperationType", "不支持的操作类型")
	}
}

type ShopBusinessInfoArgs struct {
	Uid              int
	OpIp             string
	OperationType    int    `form:"operation_type" json:"operation_type"`
	ShopId           int    `form:"shop_id" json:"shop_id"`
	NickName         string `form:"nick_name" json:"nick_name"`
	FullName         string `form:"full_name" json:"full_name"`
	RegisterAddr     string `form:"register_addr" json:"register_addr"`
	MerchantId       int    `form:"merchant_id" json:"merchant_id"`
	BusinessAddr     string `form:"business_addr" json:"business_addr"`
	BusinessLicense  string `form:"business_license" json:"business_license"`
	TaxCardNo        string `form:"tax_card_no" json:"tax_card_no"`
	BusinessDesc     string `form:"business_desc" json:"business_desc"`
	SocialCreditCode string `form:"social_credit_code" json:"social_credit_code"`
	OrganizationCode string `form:"organization_code" json:"organization_code"`
}

func (t *ShopBusinessInfoArgs) Valid(v *validation.Validation) {
	if !util.IntSliceContainsItem([]int{0, 1, 2, 3}, t.OperationType) {
		v.SetError("OperationType", "不支持的操作类型")
	}
	if util.IntSliceContainsItem([]int{1, 2, 3}, t.OperationType) {
		if t.ShopId <= 0 {
			v.SetError("ShopId", "需要大于0")
		}
	} else {
		if t.MerchantId <= 0 {
			v.SetError("MerchantId", "需要大于0")
		}
	}
	if t.NickName == "" {
		v.SetError("NickName", "不能为空")
	}
	if t.FullName == "" {
		v.SetError("FullName", "不能为空")
	}
}

type SkuBusinessPutAwayArgs struct {
	Uid           int
	OpIp          string
	OperationType int32  `form:"operation_type" json:"operation_type"`
	SkuCode       string `form:"sku_code" json:"sku_code"`
	Name          string `form:"name" json:"name"`
	Price         string `form:"price" json:"price"`
	Title         string `form:"title" json:"title"`
	SubTitle      string `form:"sub_title" json:"sub_title"`
	Desc          string `form:"desc" json:"desc"`
	Production    string `form:"production" json:"production"`
	Supplier      string `form:"supplier" json:"supplier"`
	Category      int32  `form:"category" json:"category"`
	Color         string `form:"color" json:"color"`
	ColorCode     int32  `form:"color_code" json:"color_code"`
	Specification string `form:"specification" json:"specification"`
	DescLink      string `form:"desc_link" json:"desc_link"`
	State         int32  `form:"state" json:"state"`

	Amount int64 `form:"amount" json:"amount"`
	ShopId int64 `form:"shop_id" json:"shop_id"`
}

type SkuJoinUserTrolleyArgs struct {
	Uid      int
	SkuCode  string `form:"sku_code" json:"sku_code"`
	ShopId   int    `form:"shop_id" json:"shop_id"`
	Count    int    `form:"count" json:"count"`
	Time     string `form:"time" json:"time"`
	Selected bool   `form:"selected" json:"selected"`
}

func (t *SkuJoinUserTrolleyArgs) Valid(v *validation.Validation) {
	if t.SkuCode == "" {
		v.SetError("SkuCode", "商品唯一code不能为空")
	}
	if t.Count <= 0 {
		v.SetError("Count", "数量需要大于0")
	}
	if t.ShopId <= 0 {
		v.SetError("ShopId", "店铺ID需要大于0")
	}
	if t.Time == "" {
		v.SetError("Time", "时间不能为空")
	}
}

type SkuJoinUserTrolleyRsp struct {
}

type SkuRemoveUserTrolleyArgs struct {
	Uid     int
	SkuCode string `form:"sku_code" json:"sku_code"`
	ShopId  int    `form:"shop_id" json:"shop_id"`
}

func (t *SkuRemoveUserTrolleyArgs) Valid(v *validation.Validation) {
	if t.SkuCode == "" {
		v.SetError("SkuCode", "商品唯一code不能为空")
	}
	if t.ShopId <= 0 {
		v.SetError("ShopId", "店铺ID需要大于0")
	}
}

type SkuRemoveUserTrolleyRsp struct {
}

type UserTrolleyListRsp struct {
	List []UserTrolleyRecord `json:"list"`
}

type UserTrolleyRecord struct {
	SkuCode  string `json:"sku_code"`
	ShopId   int64  `json:"shop_id"`
	Count    int64  `json:"count"`
	Time     string `json:"time"`
	Selected bool   `json:"selected"`
}

type SkuPropertyExArgs struct {
	Uid               int
	OpIp              string
	OperationType     int32  `form:"operation_type" json:"operation_type"`
	ShopId            int64  `form:"shop_id" json:"shop_id"`
	SkuCode           string `form:"sku_code" json:"sku_code"`
	Name              string `form:"name" json:"name"`
	Size              string `form:"size" json:"size"`
	Shape             string `form:"shape" json:"shape"`
	ProductionCountry string `form:"production_country" json:"production_country"`
	ProductionDate    string `form:"production_date" json:"production_date"`
	ShelfLife         string `form:"shelf_life" json:"shelf_life"`
}

func (t *SkuPropertyExArgs) Valid(v *validation.Validation) {
	if !util.IntSliceContainsItem([]int{0, 1, 2, 3}, int(t.OperationType)) {
		v.SetError("OperationType", "不支持的操作类型")
	}
	if t.SkuCode == "" {
		v.SetError("SkuCode", "商品唯一code不能为空")
	}
	if t.OperationType == 0 {
		if t.ShopId <= 0 {
			v.SetError("ShopId", "上架商品店铺ID需要大于0")
		}
	}
}

type SkuPropertyExRsp struct {
}

func (t *SkuBusinessPutAwayArgs) Valid(v *validation.Validation) {
	if !util.IntSliceContainsItem([]int{0, 1, 2, 3}, int(t.OperationType)) {
		v.SetError("OperationType", "不支持的操作类型")
	}
	if t.SkuCode == "" {
		v.SetError("SkuCode", "商品唯一code不能为空")
	}
	if t.Amount <= 0 {
		v.SetError("Amount", "商品数量需要大于0")
	}
	if t.OperationType == 0 {
		if t.ShopId <= 0 {
			v.SetError("ShopId", "上架商品店铺ID需要大于0")
		}
	}
	if t.Price == "" {
		v.SetError("Price", "需要大于0")
	}
	if t.Name == "" {
		v.SetError("Name", "不能为空")
	}
	if t.Title == "" {
		v.SetError("Title", "不能为空")
	}
	if t.SubTitle == "" {
		v.SetError("SubTitle", "不能为空")
	}
}

type SkuBusinessPutAwayRsp struct {
}

type GetSkuListArgs struct {
	ShopId int64 `form:"shop_id" json:"shop_id"`
}

type GetSkuListRsp struct {
	SkuInventoryInfoList []SkuInventoryInfo `json:"sku_inventory_info_list"`
}

type ShopBusinessInfoRsp struct {
	ShopId int `json:"shop_id"`
}

type MerchantsMaterialRsp struct {
	MerchantId int64 `json:"merchant_id"`
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

package args

import (
	"strconv"
	"time"

	"gitee.com/cristiane/micro-mall-api/pkg/util"
	"github.com/astaxie/beego/validation"
)

type UserAccountChargeArgs struct {
	Uid            int    `json:"uid"`
	Ip             string `json:"ip"`
	DeviceCode     string `form:"device_code" json:"device_code"`
	DevicePlatform string `form:"device_platform" json:"device_platform"`
	AccountType    int    `form:"account_type" json:"account_type"`
	CoinType       int    `form:"coin_type" json:"coin_type"`
	OutTradeNo     string `form:"out_trade_no" json:"out_trade_no"`
	Amount         string `form:"amount" json:"amount"`
}

func (t *UserAccountChargeArgs) Valid(v *validation.Validation) {
	if t.Amount == "" {
		v.SetError("Amount", "充值金额不能为空")
	} else {
		_, err := strconv.ParseFloat(t.Amount, 64)
		if err != nil {
			v.SetError("Amount", "充值金额不是有效数字")
		}
	}
	//if t.OutTradeNo == "" {
	//	v.SetError("OutTradeNo", "外部交易号不能为空")
	//}
	if !util.IntSliceContainsItem([]int{0, 1}, t.CoinType) {
		v.SetError("CoinType", "币种不支持")
	}
	if !util.IntSliceContainsItem([]int{0, 1, 2}, t.AccountType) {
		v.SetError("AccountType", "账户类型不支持")
	}
}

type SearchShopArgs struct {
	Keyword string `form:"keyword" json:"keyword"`
}

type SearchSkuInventoryArgs struct {
	Keyword string `form:"keyword" json:"keyword"`
}

type UserDeliveryInfo struct {
	Id           int64    `form:"id" json:"id"`
	DeliveryUser string   `form:"delivery_user" json:"delivery_user"`
	MobilePhone  string   `form:"mobile_phone" json:"mobile_phone"`
	Area         string   `form:"area" json:"area"`
	DetailedArea string   `form:"detailed_area" json:"detailed_area"`
	Label        []string `form:"label" json:"label"`
	IsDefault    bool     `form:"is_default" json:"is_default"`
}

type UserSettingAddressPutArgs struct {
	Uid int `json:"uid"`
	UserDeliveryInfo
	// 0-新增，1-修改，2-删除
	OperationType int `form:"operation_type" json:"operation_type"`
}

func (t *UserSettingAddressPutArgs) Valid(v *validation.Validation) {
	if !util.IntSliceContainsItem([]int{0, 1, 2}, t.OperationType) {
		v.SetError("OperationType", "不支持的操作类型")
	}
	if t.OperationType == 0 {
		if t.DeliveryUser == "" {
			v.SetError("DeliveryUser", "收货人不能为空")
		}
		if t.MobilePhone == "" {
			v.SetError("MobilePhone", "联系人电话不能为空")
		}
		if t.Area == "" {
			v.SetError("Area", "区域不能为空")
		}
		if t.DetailedArea == "" {
			v.SetError("DetailedArea", "详细地址不能为空")
		}
	} else {
		if t.Id <= 0 {
			v.SetError("Id", "记录ID不能为空")
		}
	}
}

type UserSettingAddressGetArgs struct {
	Uid        int `json:"uid"`
	DeliveryId int `form:"delivery_id" json:"delivery_id"`
}

type UpdateLogisticsRecordArgs struct {
	Uid int `json:"uid"`
}

type UpdateLogisticsRecordRsp struct {
}

type QueryLogisticsRecordArgs struct {
	Uid int `json:"uid"`
}

type QueryLogisticsRecordRsp struct {
}

type ApplyLogisticsArgs struct {
	Uid           int              `json:"uid"`
	OutTradeNo    string           `json:"out_trade_no" form:"out_trade_no"`
	Courier       string           `json:"courier" form:"courier"`
	CourierType   int              `json:"courier_type" form:"courier_type"`
	ReceiveType   int              `json:"receive_type" form:"receive_type"`
	SendUserId    int64            `json:"send_user_id" form:"send_user_id"`
	SendUser      string           `json:"send_user" form:"send_user"`
	SendAddr      string           `json:"send_addr" form:"send_addr"`
	SendPhone     string           `json:"send_phone" form:"send_phone"`
	SendTime      string           `json:"send_time" form:"send_time"`
	ReceiveUser   string           `json:"receive_user" form:"receive_user"`
	ReceiveUserId int64            `json:"receive_user_id" form:"receive_user_id"`
	ReceiveAddr   string           `json:"receive_addr" form:"receive_addr"`
	ReceivePhone  string           `json:"receive_phone" form:"receive_phone"`
	Goods         []GoodsLogistics `json:"goods" form:"goods"`
}

type GoodsLogistics struct {
	SkuCode string `json:"sku_code" form:"sku_code"`
	Name    string `json:"name" form:"name"`
	Kind    string `json:"kind" form:"kind"`
	Count   int64  `json:"count" form:"count"`
}

type ApplyLogisticsRsp struct {
	LogisticsCode string `json:"logistics_code"`
}

type RegisterUserArgs struct {
	UserName    string `form:"user_name" json:"user_name"`
	Password    string `form:"password" json:"password"`
	Sex         int    `form:"sex" json:"sex"`
	Email       string `form:"email" json:"email"`
	CountryCode string `form:"country_code" json:"country_code"`
	Phone       string `form:"phone" json:"phone"`
	Age         int    `json:"age" form:"age"`
	ContactAddr string `form:"contact_addr" json:"contact_addr"`
	VerifyCode  string `form:"verify_code" json:"verify_code"`
	IdCardNo    string `form:"id_card_no" json:"id_card_no"`
	InviteCode  string `form:"invite_code" json:"invite_code"`
	AccountId   string `form:"account_id" json:"account_id"`
}

func (t *RegisterUserArgs) Valid(v *validation.Validation) {
	if t.UserName == "" {
		v.SetError("UserName", "姓名不能为空")
	}
	if len(t.Password) < 6 {
		v.SetError("Password", "密码长度不能少于6位")
	}
	if len(t.Password) > 128 {
		v.SetError("Password", "密码长度不能多余128位")
	}
	if t.CountryCode == "" {
		t.CountryCode = "86"
	}
	if t.AccountId != "" {
		if len(t.AccountId) > 36 {
			v.SetError("AccountId", "账号长度不能多余36位")
		}
	}
	if len(t.CountryCode) < 2 {
		v.SetError("CountryCode", "国际码不能少于2位")
	}
	if len(t.Phone) != 11 {
		v.SetError("Phone", "手机号只能为11位")
	}
	if len(t.VerifyCode) < 6 {
		v.SetError("VerifyCode", "验证码不能少于6位")
	}
	if t.IdCardNo != "" {
		if len(t.IdCardNo) != 18 {
			v.SetError("IdCardNo", "身份证号码必须18位")
		}
	}
}

type RegisterUserRsp struct {
	InviteCode string `json:"invite_code"` // 注册成功返回邀请码
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
	if len(t.Phone) != 11 {
		v.SetError("Phone", "手机号只能为111位")
	}
	if len(t.VerifyCode) < 6 {
		v.SetError("VerifyCode", "验证码不能少于6位")
	}
}

type LoginUserWithAccountArgs struct {
	AccountId string `json:"account_id" form:"account_id"`
	Password  string `form:"password" json:"password"`
}

func (t *LoginUserWithAccountArgs) Valid(v *validation.Validation) {
	if len(t.AccountId) == 0 {
		v.SetError("AccountId", "账号不能为空")
	} else {
		if len(t.AccountId) > 36 {
			v.SetError("AccountId", "账号不能多余36位")
		}
	}
	if len(t.Password) < 6 {
		v.SetError("Password", "密码长度不能少于6位")
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
	if len(t.Phone) != 11 {
		v.SetError("Phone", "手机号只能为11位")
	}
}

type GenVerifyCodeArgs struct {
	Uid          int    `json:"uid"`
	CountryCode  string `form:"country_code" json:"country_code"`
	Phone        string `form:"phone" json:"phone"`
	BusinessType int    `form:"business_type" json:"business_type"`
	ReceiveEmail string `form:"receive_email" json:"receive_email"`
}

func (t *GenVerifyCodeArgs) Valid(v *validation.Validation) {
	if t.Uid <= 0 {
		if len(t.CountryCode) < 2 {
			v.SetError("CountryCode", "国际码不能少于2位")
		}
		if len(t.Phone) != 11 {
			v.SetError("Phone", "手机号只能为11位")
		}
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

func (t *SkuBusinessPutAwayArgs) Valid(v *validation.Validation) {
	if !util.IntSliceContainsItem([]int{0, 1, 2, 3, 4}, int(t.OperationType)) {
		v.SetError("OperationType", "不支持的操作类型")
	}
	if t.SkuCode == "" {
		v.SetError("SkuCode", "商品唯一code不能为空")
	}
	if t.Amount <= 0 {
		v.SetError("Amount", "商品数量需要大于0")
	}
	if t.ShopId <= 0 {
		v.SetError("ShopId", "上架商品店铺ID需要大于0")
	}
	if t.OperationType == 0 {
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
	Count   int    `json:"count" form:"count"`
}

func (t *SkuRemoveUserTrolleyArgs) Valid(v *validation.Validation) {
	if t.Count == 0 { // -1 表示全部移除，否则按数量
		t.Count = 1
	}
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

type SkuBusinessPutAwayRsp struct {
}

type GetSkuListArgs struct {
	ShopId      int64    `form:"shop_id" json:"shop_id"`
	SkuCodeList []string `form:"sku_code_list" json:"sku_code_list"`
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

type ListUserInfoArgs struct {
	PageMeta
	Token string `form:"token" json:"token"`
}

type ListUserInfoRsp struct {
	UserInfoList []UserMobilePhone `json:"user_info_list"`
}

type SearchUserInfoArgs struct {
	Query string `json:"query" form:"query"`
}

type SearchMerchantInfoArgs struct {
	Query string `json:"query" form:"query"`
}

type SearchTradeOrderArgs struct {
	Query string `json:"query" form:"query"`
}

type UserMobilePhone struct {
	CountryCode string `json:"country_code"`
	Phone       string `json:"phone"`
}

type PageMeta struct {
	PageSize int32 `form:"page_size" json:"page_size"`
	PageNum  int32 `form:"page_num" json:"page_num"`
}

func (p *PageMeta) Valid(v *validation.Validation) {
	if p.PageNum < 1 {
		v.SetError("PageNum", "PageNum不能小于1")
	}
	if p.PageSize <= 0 {
		v.SetError("PageSize", "PageSize 无效")
	}
	if p.PageSize > 1000 {
		v.SetError("PageSize", "PageSize最大为1000")
	}
}

type OrderShopGoods struct {
	SkuCode string `json:"sku_code"`
	Price   string `json:"price"`
	Amount  int64  `json:"amount"`
	Name    string `json:"name"`
	Version int64  `json:"version"`
}

type OrderShopSceneInfo struct {
	StoreInfo *OrderShopStoreInfo `json:"store_info"`
}

type OrderShopStoreInfo struct {
	Id       int64  `json:"id"`
	Name     string `json:"name"`
	AreaCode string `json:"area_code"`
	Address  string `json:"address"`
}

type OrderShopDetail struct {
	ShopId    int64               `json:"shop_id"`
	CoinType  int32               `json:"coin_type"`
	Goods     []*OrderShopGoods   `json:"goods"`
	SceneInfo *OrderShopSceneInfo `json:"scene_info"`
}

type CreateTradeOrderArgs struct {
	Uid            int64
	ClientIp       string             `form:"client_ip" json:"client_ip"`
	Description    string             `form:"description" json:"description"`
	DeviceId       string             `form:"device_id" json:"device_id"`
	OrderTxCode    string             `form:"order_tx_code" json:"order_tx_code"`
	UserDeliveryId int32              `form:"user_delivery_id" json:"user_delivery_id"`
	Detail         []*OrderShopDetail `form:"detail" json:"detail"`
}

func (t *CreateTradeOrderArgs) Valid(v *validation.Validation) {
	if len(t.Detail) <= 0 {
		v.SetError("Detail", "至少包含一个店铺的订单")
		return
	}
	//if t.OrderTxCode == "" {
	//	v.SetError("OrderTxCode", "订单事务ID不能为空")
	//	return
	//}
	if t.UserDeliveryId <= 0 {
		v.SetError("UserDeliveryId", "订单交付（收货地址）信息不能为空")
		return
	}
	for i := 0; i < len(t.Detail); i++ {
		if t.Detail[i].ShopId <= 0 {
			v.SetError("ShopId", "店铺ID需要大于0")
			return
		}
		if !util.IntSliceContainsItem(CoinTypes, int(t.Detail[i].CoinType)) {
			v.SetError("CoinTypes", "不支持的币种符号")
			return
		}
		if len(t.Detail[i].Goods) <= 0 {
			v.SetError("Goods", "一个店铺至少包含一个商品")
			return
		}
		goods := t.Detail[i].Goods
		for j := 0; j < len(goods); j++ {
			if goods[j].SkuCode == "" {
				v.SetError("SkuCode", "商品sku_code不能为空")
				return
			}
			if goods[j].Name == "" {
				v.SetError("Name", "商品名称不能为空")
				return
			}
			if goods[j].Amount <= 0 {
				v.SetError("Amount", "商品数量需要大于0")
				return
			}
			if goods[j].Price == "" {
				_ = v.SetError("Price", "商品价格不能为空")
				return
			}
			if goods[j].Version <= 0 {
				v.SetError("Version", "商品价格版本需大于0")
				return
			}
			p, err := strconv.ParseFloat(goods[j].Price, 64)
			if err != nil {
				_ = v.SetError("Price", "商品价格格式不正确")
				return
			}
			if p < 0 {
				_ = v.SetError("Price", "商品价格需要大于等于0")
				return
			}
		}
	}
}

type CreateTradeOrderRsp struct {
	TxCode string `json:"tx_code"` // 9-19修改为返回交易号
	//OrderEntryList []OrderEntry `json:"order_entry_list"`
}

type OrderTradeArgs struct {
	TxCode string `form:"tx_code" json:"tx_code"`
	OpUid  int64  `json:"op_uid"`
	OpIp   string `json:"op_ip"`
}

func (t *OrderTradeArgs) Valid(v *validation.Validation) {
	if t.TxCode == "" {
		_ = v.SetError("TxCode", "订单交易号不能为空")
		return
	}
}

type OrderTradeRsp struct {
	IsSuccess bool `json:"is_success"`
}

type GetOrderReportArgs struct {
	Uid       int64
	ShopId    int64  `form:"shop_id" json:"shop_id"`
	StartTime string `form:"start_time" json:"start_time"`
	EndTime   string `form:"end_time" json:"end_time"`
	PageSize  int    `form:"page_size" json:"page_size"`
	PageNum   int    `form:"page_num" json:"page_num"`
}

func (t *GetOrderReportArgs) Valid(v *validation.Validation) {
	if t.StartTime == "" {
		t.StartTime = time.Now().AddDate(0, 0, -3).Format("2006-01-02 15:04:05")
	}
	if t.EndTime == "" {
		t.EndTime = time.Now().Format("2006-01-02 15:04:05")
	}
	if t.PageSize <= 0 {
		t.PageSize = 100
	}
	if t.PageSize > 10000 {
		_ = v.SetError("PageSize", "分页大小超过限制(10000)")
		return
	}
	if t.ShopId <= 0 {
		_ = v.SetError("ShopId", "店铺ID不能为空")
		return
	}
	if t.PageNum < 1 {
		t.PageNum = 1
	}
}

type OrderShopRankArgs struct {
	Uid       int64  `form:"uid" json:"uid"`
	ShopId    int64  `form:"shop_id" json:"shop_id"`
	StartTime string `form:"start_time" json:"start_time"`
	EndTime   string `form:"end_time" json:"end_time"`
	PageSize  int    `form:"page_size" json:"page_size"`
	PageNum   int    `form:"page_num" json:"page_num"`
}

func (t *OrderShopRankArgs) Valid(v *validation.Validation) {
	if t.PageSize > 1000 {
		_ = v.SetError("PageSize", "分页大小超过限制(1000)")
		return
	}
	if t.PageSize < 1 {
		_ = v.SetError("PageSize", "分页单页数量应该大于0")
		return
	}
	if t.ShopId < 0 {
		_ = v.SetError("ShopId", "店铺ID不能为空")
		return
	}
	if t.PageNum < 1 {
		_ = v.SetError("PageNum", "分页起始页码需要大于0")
		return
	}
}

type OrderSkuRankArgs struct {
	ShopId    int64  `form:"shop_id" json:"shop_id"`
	SkuCode   string `form:"sku_code" json:"sku_code"`
	Name      string `form:"name" json:"name"`
	StartTime string `form:"start_time" json:"start_time"`
	EndTime   string `form:"end_time" json:"end_time"`
	PageSize  int    `form:"page_size" json:"page_size"`
	PageNum   int    `form:"page_num" json:"page_num"`
}

func (t *OrderSkuRankArgs) Valid(v *validation.Validation) {
	if t.PageSize > 1000 {
		_ = v.SetError("PageSize", "分页大小超过限制(1000)")
		return
	}
	if t.PageSize < 1 {
		_ = v.SetError("PageSize", "分页单页数量应该大于0")
		return
	}
	if t.ShopId < 0 {
		_ = v.SetError("ShopId", "店铺ID不能为空")
		return
	}
	if t.PageNum < 1 {
		_ = v.SetError("PageNum", "分页起始页码需要大于0")
		return
	}
}

type GetOrderReportRsp struct {
	ReportFilePath string `json:"report_file_path"`
}

type CommentsTags struct {
	TagCode              string `form:"tag_code" json:"tag_code"`
	ClassificationMajor  string `form:"classification_major" json:"classification_major"`
	ClassificationMedium string `form:"classification_medium" json:"classification_medium"`
	ClassificationMinor  string `form:"classification_minor" json:"classification_minor"`
	Content              string `form:"content" json:"content"`
}

type LogisticsCommentsInfo struct {
	LogisticsCode        string   `form:"logistics_code" json:"logistics_code"`
	FedexPack            int8     `form:"fedex_pack_star" json:"fedex_pack"`
	FedexPackLabel       []string `form:"fedex_pack_label" json:"fedex_pack_label"`
	DeliverySpeed        int8     `form:"delivery_speed" json:"delivery_speed"`
	DeliverySpeedLabel   []string `form:"delivery_speed_label" json:"delivery_speed_label"`
	DeliveryService      int8     `form:"delivery_service" json:"delivery_service"`
	DeliveryServiceLabel []string `form:"delivery_service_label" json:"delivery_service_label"`
	Comment              string   `form:"comment" json:"comment"`
}

type OrderCommentsInfo struct {
	ShopId    int64    `form:"shop_id" json:"shop_id"`
	OrderCode string   `form:"order_code" json:"order_code"`
	Star      int8     `form:"star" json:"star"`
	Content   string   `form:"content" json:"content"`
	ImgList   []string `form:"img_list" json:"img_list"`
	CommentId string   `form:"comment_id" json:"comment_id"`
}

type CreateOrderCommentsArgs struct {
	Uid                   int64 `json:"uid"`
	Anonymity             bool  `form:"anonymity" json:"anonymity"`
	OrderCommentsInfo     OrderCommentsInfo
	LogisticsCommentsInfo LogisticsCommentsInfo
}

type GetShopCommentsListArgs struct {
	Uid    int64 `json:"uid"`
	ShopId int64 `form:"shop_id" json:"shop_id"`
}

type ModifyCommentsTagsArgs struct {
	Uid           int64 `json:"uid"`
	OperationType int8  `form:"operation_type" json:"operation_type"`
	CommentsTags
}

type GetCommentsTagsListArgs struct {
	Uid int64 `json:"uid"`
	CommentsTags
}

package client

import (
	"fmt"
	"gitee.com/cristiane/go-common/json"
	"github.com/google/uuid"
	"io/ioutil"
	"math"
	"math/rand"
	"net/http"
	"net/url"
	"strings"
	"testing"
	"time"
)

const (
	baseUrlProd    = "https://xxx.xxx.xx/api"
	baseUrlTestAli = "http://xx.xx.xx.xx:xxx/api"
	baseUrlDev     = "http://xx.xx.xx.56:xx/api"
	baseUrlLocal   = "http://localhost:52001/api"
)

const (
	verifyCodeSend          = "/verify_code/send"
	registerUser            = "/register"
	loginUserWithVerifyCode = "/login/verify_code"
	loginUserWithPwd        = "/login/pwd"
	userPwdReset            = "/user/password/reset"
	userInfo                = "/user/user_info"
	merchantsMaterial       = "/user/merchants/material"
	shopBusinessApply       = "/user/shop_business/shop/apply"
	skuBusinessPutAway      = "/user/sku_business/sku/put_away"
	skuBusinessGetSkuList   = "/user/sku_business/sku/list"
	skuBusinessSupplement   = "/user/sku_business/sku/supplement"
	skuJoinUserTrolley      = "/user/trolley/sku/join"
	skuRemoveUserTrolley    = "/user/trolley/sku/remove"
	skuUserTrolleyList      = "/user/trolley/sku/list"
	tradeCreateOrder        = "/user/order/create"
	tradeOrderCodeGen       = "/user/order/code/gen"
	tradeOrderPay           = "/user/order/trade"
	logisticsApply          = "/user/logistics/apply"
	userSettingAddress      = "/user/setting/address"
	searchSkuInventory      = "/search/sku_inventory"
	searchShop              = "/search/shop"
)

const (
	apiV1 = "/v1"
	apiV2 = "/v2"
)

var apiVersion = apiV1
var qToken = token_10041
var baseUrl = baseUrlLocal + apiVersion

func TestMain(m *testing.M) {
	m.Run()
}

func TestGateway(t *testing.T) {
	t.Run("发送验证码", TestVerifyCodeSend)
	t.Run("注册用户", TestRegisterUser)
	t.Run("登录用户-验证码", TestLoginUserWithVerifyCode)
	t.Run("登录用户-密码", TestLoginUserWithPwd)
	t.Run("重置密码", TestLoginUserPwdReset)
	t.Run("获取用户信息", TestGetUserInfo)
	t.Run("用户申请提交审核资料", TestMerchantsMaterial)
	t.Run("商户提交开店材料", TestShopBusinessApply)
	t.Run("店铺上架商品", TestSkuBusinessPutAway)
	t.Run("获取店铺上架商品列表", TestGetSkuList)
	t.Run("补充商品属性", TestSkuBusinessSupplement)
	t.Run("添加商品到购物车", TestSkuJoinUserTrolley)
	t.Run("从购物车移除商品", TestSkuRemoveUserTrolley)
	t.Run("获取用户购物车列表", TestGetUserTrolleyList)
	t.Run("生成唯一订单号", TestGenOrderCode)
	t.Run("创建交易订单", TestTradeCreateOrder)
	t.Run("交易订单支付", TestOrderTradePay)
	t.Run("申请物流", TestLogisticsApply)
	t.Run("用户设置-地址变更", TestUserSettingAddress)
	t.Run("用户设置-获取收货地址", TestUserSettingAddressGet)
	t.Run("搜索-商品库存", TestSearchSkuInventory)
	t.Run("搜索-店铺", TestSearchShop)
}

func TestGetUserInfo(t *testing.T) {
	r := baseUrl + userInfo
	t.Logf("request url: %s", r)
	req, err := http.NewRequest("GET", r, nil)
	if err != nil {
		t.Error(err)
		return
	}
	req.Header.Set("token", qToken)
	commonTest(r, req, t)
}

func TestSearchShop(t *testing.T) {
	r := baseUrl + searchShop + "?keyword=交个朋友"
	t.Logf("request url: %s", r)
	req, err := http.NewRequest("GET", r, nil)
	if err != nil {
		t.Error(err)
		return
	}
	//req.Header.Set("token", qToken)
	commonTest(r, req, t)
}

func TestSearchSkuInventory(t *testing.T) {
	r := baseUrl + searchSkuInventory + "?keyword=剃须刀"
	t.Logf("request url: %s", r)
	req, err := http.NewRequest("GET", r, nil)
	if err != nil {
		t.Error(err)
		return
	}
	//req.Header.Set("token", qToken)
	commonTest(r, req, t)
}

func TestGetUserTrolleyList(t *testing.T) {
	r := baseUrl + skuUserTrolleyList
	t.Logf("request url: %s", r)
	req, err := http.NewRequest("GET", r, nil)
	if err != nil {
		t.Error(err)
		return
	}
	req.Header.Set("token", qToken)
	commonTest(r, req, t)
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
	UserDeliveryInfo
	// 0-新增，1-修改，2-删除
	OperationType int `form:"operation_type" json:"operation_type"`
}

func TestUserSettingAddress(t *testing.T) {
	r := baseUrl + userSettingAddress
	t.Logf("request url: %s", r)
	args := UserSettingAddressPutArgs{
		UserDeliveryInfo: UserDeliveryInfo{
			Id:           101,
			DeliveryUser: "张6丰",
			MobilePhone:  "15501707785",
			Area:         "广东省广州市",
			DetailedArea: "上海路步行街111号",
			Label:        []string{"公司", "住宅", "生活"},
			IsDefault:    true,
		},
		OperationType: 0,
	}
	data := json.MarshalToStringNoError(args)
	t.Logf("req data: \n%v", data)
	req, err := http.NewRequest("POST", r, strings.NewReader(data))
	if err != nil {
		t.Error(err)
		return
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("token", qToken)
	commonTest(r, req, t)
}

type OrderShopGoods struct {
	SkuCode string `form:"sku_code" json:"sku_code"`
	Price   string `form:"price" json:"price"`
	Amount  int64  `form:"amount" json:"amount"`
	Name    string `form:"name" json:"name"`
	Version int64  `form:"version" json:"version"`
}

type OrderShopSceneInfo struct {
	StoreInfo *OrderShopStoreInfo `form:"store_info" json:"store_info"`
}

type OrderShopStoreInfo struct {
	Id       int64  `form:"id" json:"id"`
	Name     string `form:"name" json:"name"`
	AreaCode string `form:"area_code" json:"area_code"`
	Address  string `form:"address" json:"address"`
}

type OrderShopDetail struct {
	ShopId    int64               `form:"shop_id" json:"shop_id"`
	CoinType  int32               `form:"coin_type" json:"coin_type"`
	Goods     []*OrderShopGoods   `form:"goods" json:"goods"`
	SceneInfo *OrderShopSceneInfo `form:"scene_info" json:"scene_info"`
}

type CreateTradeOrderArgs struct {
	Uid            int64              `json:"uid"`
	ClientIp       string             `json:"client_ip"`
	Description    string             `form:"description" json:"description"`
	DeviceId       string             `form:"device_id" json:"device_id"`
	OrderTxCode    string             `form:"order_tx_code" json:"order_tx_code"`
	UserDeliveryId int32              `form:"user_delivery_id" json:"user_delivery_id"`
	Detail         []*OrderShopDetail `json:"detail"`
}

func TestOrderTradePay(t *testing.T) {
	r := baseUrl + tradeOrderPay
	t.Logf("request url: %s", r)
	data := url.Values{}
	data.Set("tx_code", "c92b7b49-497b-4d1e-98f3-ef16fb140f6a")
	t.Logf("req data: %v", data)
	req, err := http.NewRequest("POST", r, strings.NewReader(data.Encode()))
	if err != nil {
		t.Error(err)
		return
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("token", qToken)
	commonTest(r, req, t)
}

type ApplyLogisticsArgs struct {
	Uid          int              `json:"uid"`
	OutTradeNo   string           `json:"out_trade_no" form:"out_trade_no"`
	Courier      string           `json:"courier" form:"courier"`
	CourierType  int              `json:"courier_type" form:"courier_type"`
	ReceiveType  int              `json:"receive_type" form:"receive_type"`
	SendUser     string           `json:"send_user" form:"send_user"`
	SendAddr     string           `json:"send_addr" form:"send_addr"`
	SendPhone    string           `json:"send_phone" form:"send_phone"`
	SendTime     string           `json:"send_time" form:"send_time"`
	ReceiveUser  string           `json:"receive_user" form:"receive_user"`
	ReceiveAddr  string           `json:"receive_addr" form:"receive_addr"`
	ReceivePhone string           `json:"receive_phone" form:"receive_phone"`
	Goods        []GoodsLogistics `json:"goods" form:"goods"`
}

type GoodsLogistics struct {
	SkuCode string `json:"sku_code" form:"sku_code"`
	Name    string `json:"name" form:"name"`
	Kind    string `json:"kind" form:"kind"`
	Count   int64  `json:"count" form:"count"`
}

func TestLogisticsApply(t *testing.T) {
	r := baseUrl + logisticsApply
	t.Logf("request url: %s", r)
	applyReq := ApplyLogisticsArgs{
		Uid:          0,
		OutTradeNo:   uuid.New().String(),
		Courier:      "微商城快递",
		CourierType:  1,
		ReceiveType:  1,
		SendUser:     "李云龙",
		SendAddr:     "河北省石家庄市丰县迎宾路123号",
		SendPhone:    "18319430520",
		SendTime:     "2020-10-09 12:12:12",
		ReceiveUser:  "马司令",
		ReceiveAddr:  "浙江省杭州市余杭区西湖南路111雅静别院",
		ReceivePhone: "18319430520",
		Goods: []GoodsLogistics{
			{
				SkuCode: "2131d-f111-45e1-b68a-d602c2f0f1b3",
				Name:    "怡宝矿泉水",
				Kind:    "饮用水",
				Count:   98,
			},
		},
	}
	data := json.MarshalToStringNoError(applyReq)
	t.Logf("req data: %v", data)
	req, err := http.NewRequest("POST", r, strings.NewReader(data))
	if err != nil {
		t.Error(err)
		return
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("token", qToken)
	commonTest(r, req, t)
}

func TestTradeCreateOrder(t *testing.T) {
	r := baseUrl + tradeCreateOrder
	t.Logf("request url: %s", r)
	goods1 := OrderShopGoods{
		SkuCode: "3084638f-b8a9-4ce1-a7d3-271ec37e1a0e",
		Price:   "184.32",
		Amount:  1,
		Name:    "盼盼铜锣烧--01",
		Version: 1,
	}
	goods2 := OrderShopGoods{
		SkuCode: "db696ade-3de2-4417-bea0-9b015da37191",
		Price:   "184.32",
		Amount:  1,
		Name:    "盼盼铜锣烧--02",
		Version: 1,
	}
	// b882a5c9-564a-4912-a5d4-ce77de71577c
	detail := OrderShopDetail{
		ShopId:   30059,
		CoinType: 0, // 0-rmb,1-usdt
		Goods:    []*OrderShopGoods{&goods1, &goods2},
		SceneInfo: &OrderShopSceneInfo{
			StoreInfo: &OrderShopStoreInfo{
				Id:       30059,
				Name:     "良品铺子京东旗舰店1",
				AreaCode: "深圳",
				Address:  "深圳市宝安区",
			},
		},
	}
	goods3 := OrderShopGoods{
		SkuCode: "5d6f191d-7d4e-49e5-9384-04abb7ba8cfb",
		Price:   "1099.55",
		Amount:  2,
		Name:    "爱思席梦思",
		Version: 1,
	}
	detail2 := OrderShopDetail{
		ShopId:   30068,
		CoinType: 0,
		Goods:    []*OrderShopGoods{&goods3},
		SceneInfo: &OrderShopSceneInfo{
			StoreInfo: &OrderShopStoreInfo{
				Id:       30068,
				Name:     "爱思席梦思",
				AreaCode: "广东省广州市",
				Address:  "广州",
			},
		},
	}
	data := CreateTradeOrderArgs{
		Description: "双11活动",
		DeviceId:    "xiaomi-10",
		OrderTxCode: uuid.New().String(),
		//OrderTxCode: "84fd4745-f0c0-4a7c-a522-16ef02d058e09",
		UserDeliveryId: 105,
		Detail:         []*OrderShopDetail{&detail, &detail2},
	}
	//log.Println(json.MarshalToStringNoError(data))
	t.Logf("req data: %v", json.MarshalToStringNoError(data))
	req, err := http.NewRequest("POST", r, strings.NewReader(json.MarshalToStringNoError(data)))
	if err != nil {
		t.Error(err)
		return
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("token", qToken)
	commonTest(r, req, t)
}

func TestVerifyCodeSend(t *testing.T) {
	r := baseUrl + verifyCodeSend
	t.Logf("request url: %s", r)
	data := url.Values{}
	data.Set("country_code", "86")
	data.Set("phone", "15501707783")
	data.Set("business_type", "1")
	data.Set("receive_email", "565608463@qq.com")
	t.Logf("req data: %v", data)
	req, err := http.NewRequest("POST", r, strings.NewReader(data.Encode()))
	if err != nil {
		t.Error(err)
		return
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("token", qToken)
	commonTest(r, req, t)
}

func TestRegisterUser(t *testing.T) {
	r := baseUrl + registerUser
	t.Logf("request url: %s", r)
	data := url.Values{}
	data.Set("user_name", "张全蛋")
	data.Set("password", "07030501310")
	data.Set("sex", "1")
	data.Set("age", "33")
	data.Set("country_code", "86")
	data.Set("phone", "15501707783")
	data.Set("email", "31342314@qq.com")
	data.Set("verify_code", "148300")
	data.Set("id_card_no", "513913899384938899")
	data.Set("contact_addr", "廊坊市淮南路清明河畔李家大院")
	data.Set("invite_code", "489f043b3000065")
	t.Logf("req data: %v", data)
	req, err := http.NewRequest("POST", r, strings.NewReader(data.Encode()))
	if err != nil {
		t.Error(err)
		return
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("token", qToken)
	commonTest(r, req, t)
}

func TestLoginUserWithVerifyCode(t *testing.T) {
	r := baseUrl + loginUserWithVerifyCode
	t.Logf("request url: %s", r)
	data := url.Values{}
	data.Set("country_code", "86")
	data.Set("phone", "18319430520")
	data.Set("verify_code", "876306")
	t.Logf("req data: %v", data)
	req, err := http.NewRequest("POST", r, strings.NewReader(data.Encode()))
	if err != nil {
		t.Error(err)
		return
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("token", qToken)
	commonTest(r, req, t)
}

func TestLoginUserWithPwd(t *testing.T) {
	r := baseUrl + loginUserWithPwd
	t.Logf("request url: %s", r)
	data := url.Values{}
	data.Set("country_code", "86")
	data.Set("phone", "15501707783")
	data.Set("password", "07030501310")
	t.Logf("req data: %v", data)
	req, err := http.NewRequest("POST", r, strings.NewReader(data.Encode()))
	if err != nil {
		t.Error(err)
		return
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("token", qToken)
	commonTest(r, req, t)
}

func BenchmarkTestLoginUserWithPwd(b *testing.B) {
	r := baseUrl + loginUserWithPwd
	b.Logf("request url: %s", r)
	data := url.Values{}
	data.Set("country_code", "86")
	data.Set("phone", "15501707783")
	data.Set("password", "07030501310")
	b.Logf("req data: %v", data)
	req, err := http.NewRequest("POST", r, strings.NewReader(data.Encode()))
	if err != nil {
		b.Error(err)
		return
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("token", qToken)
	for i := 0; i < math.MaxInt32; i++ {
		//b.Logf("request token=%v", qToken)
		rsp, err := http.DefaultClient.Do(req)
		if err != nil {
			b.Error(err)
			return
		}
		//b.Logf("req url: %v status : %v", r, rsp.Status)
		if rsp.StatusCode != http.StatusOK {
			b.Error("StatusCode != 200")
			return
		}
		body, err := ioutil.ReadAll(rsp.Body)
		defer rsp.Body.Close()
		if err != nil {
			b.Error(err)
			return
		}
		//b.Logf("req url: %v body : \n%s", r, body)
		var obj HttpCommonRsp
		err = json.Unmarshal(string(body), &obj)
		if err != nil {
			b.Error(err)
			return
		}
		if obj.Code != SuccessBusinessCode {
			b.Errorf("business code != %v", SuccessBusinessCode)
			b.Errorf("obj ==%+v,obj", obj)
			return
		}
	}
}

func TestLoginUserPwdReset(t *testing.T) {
	r := baseUrl + userPwdReset
	t.Logf("request url: %s", r)
	data := url.Values{}
	data.Set("verify_code", "381825")
	data.Set("password", "12345678")
	t.Logf("req data: %v", data)
	req, err := http.NewRequest("PUT", r, strings.NewReader(data.Encode()))
	if err != nil {
		t.Error(err)
		return
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("token", qToken)
	commonTest(r, req, t)
}

func TestMerchantsMaterial(t *testing.T) {
	r := baseUrl + merchantsMaterial
	t.Logf("request url: %s", r)
	data := url.Values{}
	data.Set("operation_type", "0")
	data.Set("register_addr", "深圳市宝安区兴业路宝源二区72栋-深圳星光无限实业有限责任公司")
	data.Set("health_card_no", "R8nJ65TDUGAlqrwerSdb9")
	data.Set("identity", "1")
	data.Set("tax_card_no", "qX2Mr545kznWrlvO4sIp7")
	t.Logf("req data: %v", data)
	req, err := http.NewRequest("PUT", r, strings.NewReader(data.Encode()))
	if err != nil {
		t.Error(err)
		return
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("token", qToken)
	commonTest(r, req, t)
}

func TestShopBusinessApply(t *testing.T) {
	r := baseUrl + shopBusinessApply
	t.Logf("request url: %s", r)
	data := url.Values{}
	data.Set("operation_type", "0")
	data.Set("shop_id", "123")
	data.Set("nick_name", "深圳市有他没我科技有限公司")
	data.Set("full_name", "深圳市有他没我科技有限公司")
	data.Set("register_addr", "深圳市宝安区兴业路宝源二区72栋")
	data.Set("merchant_id", "1069")
	data.Set("business_addr", "深圳市宝安区宝源二区73栋111号")
	data.Set("business_license", "qX2MkznWrlvO4sIp7")
	data.Set("tax_card_no", "qX2MkznWrlvO4sIp7")
	data.Set("business_desc", "qX2MkznWrlvO4sIp7")
	data.Set("social_credit_code", "qX2MkznWrlvO4sIp7")
	data.Set("organization_code", "qX2MkznWrlvO4sIp7")
	t.Logf("req data: %v", data)
	req, err := http.NewRequest("POST", r, strings.NewReader(data.Encode()))
	if err != nil {
		t.Error(err)
		return
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("token", qToken)
	commonTest(r, req, t)
}

func TestSkuBusinessPutAway(t *testing.T) {
	r := baseUrl + skuBusinessPutAway
	t.Logf("request url: %s", r)
	data := url.Values{}
	data.Set("operation_type", "0")
	data.Set("shop_id", "30068")
	data.Set("sku_code", uuid.New().String())
	data.Set("name", "海飞丝洗发水")
	data.Set("price", "59")
	data.Set("title", "海飞丝洗发水，洗发露")
	data.Set("sub_title", "海飞丝洗发水套装怡神冰凉薄荷500ml*2送无硅油80ml（持久去屑去油止痒清爽 男士女士通用）清香型")
	data.Set("desc", "海飞丝是宝洁公司的一款洗发精产品。1963年，全球第一支含有活性去屑成分，可有效去除头屑的洗发露诞生，自此，有效去除头屑成为海飞丝深受全球消费者喜爱的最出色的功效。海飞丝品牌1986年进入台湾，1991年进入中国大陆，其系列产品包括：洗发露、护发素、头皮头发按摩膏、头皮修护精华乳。始创于1837年的宝洁公司，是世界最大的日用消费品公司之一。所经营的300多个品牌的产品畅销160多个国家和地区，其中包括织物及家居护理、美发美容、婴儿及家庭护理、健康护理、食品及饮料等。")
	data.Set("production", "广州宝洁有限公司")
	data.Set("supplier", "宝洁拼多多专营店")
	data.Set("category", "11010")
	data.Set("color", "白色")
	data.Set("color_code", "199")
	data.Set("specification", "瓶装，560ml")
	data.Set("desc_link", "https://item.jd.com/1750531.html")
	data.Set("state", "1")
	data.Set("amount", "100")
	t.Logf("req data: %v", data)
	req, err := http.NewRequest("POST", r, strings.NewReader(data.Encode()))
	if err != nil {
		t.Error(err)
		return
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("token", qToken)
	commonTest(r, req, t)
}

func TestGetSkuList(t *testing.T) {
	r := baseUrl + skuBusinessGetSkuList + "?shop_id=30068"
	t.Logf("request url: %s", r)
	req, err := http.NewRequest("GET", r, nil)
	if err != nil {
		t.Error(err)
		return
	}
	req.Header.Set("token", qToken)
	commonTest(r, req, t)
}

func TestUserSettingAddressGet(t *testing.T) {
	r := baseUrl + userSettingAddress + "?delivery_id="
	t.Logf("request url: %s", r)
	req, err := http.NewRequest("GET", r, nil)
	if err != nil {
		t.Error(err)
		return
	}
	req.Header.Set("token", qToken)
	commonTest(r, req, t)
}

func TestGenOrderCode(t *testing.T) {
	r := baseUrl + tradeOrderCodeGen
	t.Logf("request url: %s", r)
	req, err := http.NewRequest("GET", r, nil)
	if err != nil {
		t.Error(err)
		return
	}
	req.Header.Set("token", qToken)
	commonTest(r, req, t)
}

func TestSkuBusinessSupplement(t *testing.T) {
	r := baseUrl + skuBusinessSupplement
	t.Logf("request url: %s", r)
	data := url.Values{}
	data.Set("operation_type", "0")
	data.Set("shop_id", "30068")
	data.Set("sku_code", "5d6f191d-7d4e-49e5-9384-04abb7ba8cfb")
	data.Set("name", "爱思席梦思")
	data.Set("size", "1.8m")
	data.Set("shape", "长方形")
	data.Set("production_country", "河南爱思实业有限公司")
	data.Set("production_date", "2020/10/19 15:20")
	data.Set("shelf_life", "3.3年")
	t.Logf("req data: %v", data)
	req, err := http.NewRequest("PUT", r, strings.NewReader(data.Encode()))
	if err != nil {
		t.Error(err)
		return
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("token", qToken)
	commonTest(r, req, t)
}

func TestSkuJoinUserTrolley(t *testing.T) {
	r := baseUrl + skuJoinUserTrolley
	t.Logf("request url: %s", r)
	data := url.Values{}
	data.Set("shop_id", "30059")
	data.Set("sku_code", "b363e9f4-3bae-4103-86a6-5e4b83b70303-11")
	data.Set("count", "2020")
	data.Set("time", "2020-09-08 23:32:35")
	data.Set("selected", "false")
	t.Logf("req data: %v", data)
	req, err := http.NewRequest("PUT", r, strings.NewReader(data.Encode()))
	if err != nil {
		t.Error(err)
		return
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("token", qToken)
	commonTest(r, req, t)
}

func TestSkuRemoveUserTrolley(t *testing.T) {
	r := baseUrl + skuRemoveUserTrolley
	t.Logf("request url: %s", r)
	skuCode := "df1a9633-b060-4682-9502-bc934f89392b"
	shopId := "30059"
	r += "?sku_code=" + skuCode + "&shop_id=" + shopId
	req, err := http.NewRequest("DELETE", r, nil)
	if err != nil {
		t.Error(err)
		return
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("token", qToken)
	commonTest(r, req, t)
}

func BenchmarkTestShopBusinessApply(b *testing.B) {
	r := baseUrl + shopBusinessApply
	b.Logf("request url: %s", r)
	for i := 0; i < math.MaxInt16; i++ {
		time.Sleep(time.Duration(rand.Intn(5)) * time.Second)
		data := url.Values{}
		data.Set("operation_type", "0")
		data.Set("shop_id", "123")
		data.Set("nick_name", fmt.Sprintf("良品铺子京东旗舰店-%v", uuid.New().String()))
		data.Set("full_name", "武汉市良品铺子食品股份有限公司深圳分公司宝安店")
		data.Set("register_addr", "深圳市宝安区兴业路宝源二区72栋-良品铺子京东旗舰店")
		data.Set("merchant_id", "1009")
		data.Set("business_addr", "深圳市宝安区宝源二区73栋111号")
		data.Set("business_license", "qX2MkznWrlvO4sIp7")
		data.Set("tax_card_no", "qX2MkznWrlvO4sIp7")
		data.Set("business_desc", "qX2MkznWrlvO4sIp7")
		data.Set("social_credit_code", "qX2MkznWrlvO4sIp7")
		data.Set("organization_code", "qX2MkznWrlvO4sIp7")
		b.Logf("req data: %v", data)
		req, err := http.NewRequest("POST", r, strings.NewReader(data.Encode()))
		if err != nil {
			b.Error(err)
			return
		}
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
		req.Header.Set("token", qToken)
		b.Logf("request token=%v", qToken)
		rsp, err := http.DefaultClient.Do(req)
		if err != nil {
			b.Error(err)
			return
		}
		b.Logf("req url: %v status : %v", r, rsp.Status)
		if rsp.StatusCode != http.StatusOK {
			b.Error("StatusCode != 200")
			return
		}
		body, err := ioutil.ReadAll(rsp.Body)
		defer rsp.Body.Close()
		if err != nil {
			b.Error(err)
			return
		}
		b.Logf("req url: %v body : \n%s", r, body)
		var obj HttpCommonRsp
		err = json.Unmarshal(string(body), &obj)
		if err != nil {
			b.Error(err)
			return
		}
		if obj.Code != SuccessBusinessCode {
			b.Errorf("business code != %v", SuccessBusinessCode)
			b.Errorf("obj ==%+v,obj", obj)
			return
		}
	}
}

const (
	SuccessBusinessCode = 200
)

type HttpCommonRsp struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
	Msg  string      `json:"msg"`
}

func commonTest(r string, req *http.Request, t *testing.T) {
	t.Logf("request token=%v", qToken)
	rsp, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Error(err)
		return
	}
	t.Logf("req url: %v status : %v", r, rsp.Status)
	if rsp.StatusCode != http.StatusOK {
		t.Error("StatusCode != 200")
		return
	}
	body, err := ioutil.ReadAll(rsp.Body)
	defer rsp.Body.Close()
	if err != nil {
		t.Error(err)
		return
	}
	t.Logf("req url: %v body : \n%s", r, body)
	var obj HttpCommonRsp
	err = json.Unmarshal(string(body), &obj)
	if err != nil {
		t.Error(err)
		return
	}
	if obj.Code != SuccessBusinessCode {
		t.Errorf("business code != %v", SuccessBusinessCode)
		t.Errorf("obj ==%+v,obj", obj)
		return
	}
}

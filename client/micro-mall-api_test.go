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
	tradeOrderPay           = "/user/order/trade"
	logisticsApply          = "/user/logistics/apply"
)

const (
	apiV1 = "/v1"
	apiV2 = "/v2"
)

var apiVersion = apiV1
var qToken = token_10036
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
	t.Run("创建交易订单", TestTradeCreateOrder)
	t.Run("交易订单支付", TestOrderTradePay)
	t.Run("申请物流", TestLogisticsApply)
}

const (
	SuccessBusinessCode = 200
)

type HttpCommonRsp struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
	Msg  string      `json:"msg"`
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

type OrderShopGoods struct {
	SkuCode string `form:"sku_code" json:"sku_code"`
	Price   string `form:"price" json:"price"`
	Amount  int64  `form:"amount" json:"amount"`
	Name    string `form:"name" json:"name"`
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
	Uid         int64              `json:"uid"`
	ClientIp    string             `json:"client_ip"`
	Description string             `form:"description" json:"description"`
	DeviceId    string             `form:"device_id" json:"device_id"`
	Detail      []*OrderShopDetail `json:"detail"`
}

func TestOrderTradePay(t *testing.T) {
	r := baseUrl + tradeOrderPay
	t.Logf("request url: %s", r)
	data := url.Values{}
	data.Set("tx_code", "5255cf91-6ab9-4cd2-adb5-3d9c8069d1fd")
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
	}
	goods2 := OrderShopGoods{
		SkuCode: "db696ade-3de2-4417-bea0-9b015da37191",
		Price:   "184.32",
		Amount:  2,
		Name:    "盼盼铜锣烧--02",
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
		SkuCode: "5a1d0ae1-9f1c-497b-a191-673b3572b0f9",
		Price:   "184.32",
		Amount:  3,
		Name:    "盼盼铜锣烧--03",
	}
	detail2 := OrderShopDetail{
		ShopId:   30060,
		CoinType: 0,
		Goods:    []*OrderShopGoods{&goods3},
		SceneInfo: &OrderShopSceneInfo{
			StoreInfo: &OrderShopStoreInfo{
				Id:       29911,
				Name:     "良品铺子京东旗舰店-2",
				AreaCode: "广州",
				Address:  "广州市海珠区",
			},
		},
	}
	data := CreateTradeOrderArgs{
		Description: "网络购物",
		DeviceId:    "HUAWEI",
		Detail:      []*OrderShopDetail{&detail, &detail2},
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
	data.Set("phone", "18319430520")
	data.Set("business_type", "3")
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
	data.Set("user_name", "杨子")
	data.Set("password", "07030501310")
	data.Set("sex", "1")
	data.Set("age", "28")
	data.Set("country_code", "86")
	data.Set("phone", "18319430521")
	data.Set("email", "1225807604@qq.com")
	data.Set("verify_code", "311126")
	data.Set("invite_code", "")
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
	data.Set("phone", "18319430520")
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
	data.Set("health_card_no", "R8nJ65TDUGAlqSdb9")
	data.Set("identity", "1")
	data.Set("tax_card_no", "qX2MkznWrlvO4sIp7")
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
	data.Set("nick_name", "福建赚它一个亿科技有限公司")
	data.Set("full_name", "武汉市良品铺子食品股份有限公司深圳分公司宝安店")
	data.Set("register_addr", "深圳市宝安区兴业路宝源二区72栋-良品铺子京东旗舰店")
	data.Set("merchant_id", "1037")
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
	data.Set("shop_id", "30062")
	data.Set("sku_code", uuid.New().String())
	data.Set("name", "盼盼铜锣烧")
	data.Set("price", "29.32")
	data.Set("title", "盼盼，铜锣烧，办公室零食,盼盼铜锣烧红豆味量贩箱装1000g*1")
	data.Set("sub_title", "盼盼 铜锣烧 面包饼干休闲零食量贩装红豆味1000g")
	data.Set("desc", "满满的一箱，铜锣烧，独立包装，味道不错，软软的，甜甜的，豆沙馅儿，倍儿好吃。不错的休闲食品，出去游玩携带方便，首选休闲食品。京东快递速度快，昨天晚上拍的，今天就到了。非常时期，宅在家里，享受着美味食品，支持京东")
	data.Set("production", "福建盼盼食品股份有限公司")
	data.Set("supplier", "京东盼盼食品旗舰店")
	data.Set("category", "11010")
	data.Set("color", "黄色")
	data.Set("color_code", "100")
	data.Set("specification", "整箱/30包装")
	data.Set("desc_link", "https://item.jd.com/3230143.html")
	data.Set("state", "1")
	data.Set("amount", "89")
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
	r := baseUrl + skuBusinessGetSkuList + "?"
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
	data.Set("shop_id", "30059")
	data.Set("sku_code", uuid.New().String())
	data.Set("name", "农夫山泉-矿泉水")
	data.Set("size", "200cm x 189cm")
	data.Set("shape", "完整包装")
	data.Set("production_country", "农夫山泉")
	data.Set("production_date", "2019/10/19 15:20")
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

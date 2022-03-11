package main

import (
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"math/rand"
	"net"
	"net/http"
	"net/url"
	"strings"
	"time"

	"gitee.com/kelvins-io/common/json"
	"github.com/google/uuid"
	"golang.org/x/net/http2"
)

var benchCount = math.MaxInt32
var randSleep = time.Duration(rand.Intn(200)) * time.Millisecond

func main() {
	monkey()
}

func monkey() {
	createOrderUrl := baseUrl + tradeCreateOrder
	orderTradeUrl := baseUrl + tradeOrderPay

	for i := 0; i < benchCount; i++ {
		//time.Sleep(randSleep)
		goods1 := OrderShopGoods{
			SkuCode: "dd13b4aa-4121-4898-a2b5-bcfebccb713b",
			Price:   "2.9",
			Amount:  1,
			Name:    GetFullName(),
			Version: 1,
		}
		goods2 := OrderShopGoods{
			SkuCode: "a3e5da0a-d3aa-43e2-a7b8-2c5e264e2a09",
			Price:   "23.78",
			Amount:  1,
			Name:    GetFullName(),
			Version: 1,
		}
		detail := OrderShopDetail{
			ShopId:   30071,
			CoinType: 0, // 0-rmb,1-usdt
			Goods:    []*OrderShopGoods{&goods1, &goods2},
			SceneInfo: &OrderShopSceneInfo{
				StoreInfo: &OrderShopStoreInfo{
					Id:       30071,
					Name:     GetFullName(),
					AreaCode: GetFullName(),
					Address:  GetFullName(),
				},
			},
		}
		goods3 := OrderShopGoods{
			SkuCode: "dd13b4aa-4121-a2b5-a2b5-bcfebccb4898",
			Price:   "19.9",
			Amount:  1,
			Name:    GetFullName(),
			Version: 1,
		}
		detail2 := OrderShopDetail{
			ShopId:   30072,
			CoinType: 0,
			Goods:    []*OrderShopGoods{&goods3},
			SceneInfo: &OrderShopSceneInfo{
				StoreInfo: &OrderShopStoreInfo{
					Id:       30072,
					Name:     GetFullName(),
					AreaCode: GetFullName(),
					Address:  GetFullName(),
				},
			},
		}
		data := CreateTradeOrderArgs{
			Description:    GetFullName(),
			DeviceId:       GetFullName(),
			UserDeliveryId: 220,
			Detail:         []*OrderShopDetail{&detail, &detail2},
		}
		data.OrderTxCode = uuid.New().String()
		req, err := http.NewRequest("POST", createOrderUrl, strings.NewReader(json.MarshalToStringNoError(data)))
		if err != nil {
			fmt.Println(err)
			return
		}
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("token", qToken)
		rsp, err := clientH2.Do(req)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Printf("req url: %v status : %v\n", createOrderUrl, rsp.Status)
		if rsp.StatusCode != http.StatusOK {
			fmt.Println("StatusCode != 200")
			return
		}
		body, err := ioutil.ReadAll(rsp.Body)
		rsp.Body.Close()
		if err != nil {
			fmt.Println(err)
			return
		}
		var obj CreateOrderRsp
		err = json.Unmarshal(string(body), &obj)
		if err != nil {
			fmt.Println(err)
			return
		}
		if obj.Code != SuccessBusinessCode {
			log.Printf("business code != %v", SuccessBusinessCode)
			log.Printf("obj ==%+v,obj", obj)
			continue
		}
		if obj.Data.TxCode == "" {
			fmt.Println("创建订单交易号为空")
			continue
		}
		if i%2 == 0 {
			continue
		}
		orderTradeReq := url.Values{}
		orderTradeReq.Set("tx_code", obj.Data.TxCode)
		req, err = http.NewRequest("POST", orderTradeUrl, strings.NewReader(orderTradeReq.Encode()))
		if err != nil {
			fmt.Println(err)
			return
		}
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
		req.Header.Set("token", qToken)
		commonBenchmarkTest(orderTradeUrl, req)
	}
}
func commonBenchmarkTest(r string, req *http.Request) {
	rsp, err := clientH2.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("req url: %v status : %v", r, rsp.Status)
	if rsp.StatusCode != http.StatusOK {
		fmt.Printf("StatusCode != 200 %v\n", rsp.StatusCode)
		return
	}
	body, err := ioutil.ReadAll(rsp.Body)
	defer rsp.Body.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("req url: %v body : \n%s", r, body)
	var obj HttpCommonRsp
	err = json.Unmarshal(string(body), &obj)
	if err != nil {
		fmt.Println(err)
		return
	}
	if obj.Code != SuccessBusinessCode {
		log.Printf("business code != %v", SuccessBusinessCode)
		log.Printf("obj ==%+v,obj", obj)
		return
	}
}

type HttpCommonRsp struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
	Msg  string      `json:"msg"`
}

const (
	SuccessBusinessCode = 200
)

type CreateOrderRsp struct {
	Code int `json:"code"`
	Data struct {
		TxCode string `json:"tx_code"`
	} `json:"data"`
	Msg string `json:"msg"`
}

var (
	clientH2 = http.Client{
		Transport: &http2.Transport{
			AllowHTTP: true,
			DialTLS: func(network, addr string, cfg *tls.Config) (net.Conn, error) {
				return net.Dial(network, addr)
			},
		},
	}
)

const (
	token_79857 = "eyJhbGciOiJIUzM4NCIsInR5cCI6IkpXVCJ9.eyJ1c2VyX25hbWUiOiLlhazlrZnmgZIiLCJ1aWQiOjc5ODU3LCJleHAiOjE2NDU5MTUzOTUsImlzcyI6IndlYl9naW5fdGVtcGxhdGUifQ.c7GQiSFqiFRG0YsoE0os2Y0qr_afXTkw312628ljrGBBX4Pte2nCTFApd2dqQaQ_"
)

var qToken = token_79857
var lastName = []string{
	"赵", "钱", "孙", "李", "周", "吴", "郑", "王", "冯", "陈", "褚", "卫", "蒋",
	"沈", "韩", "杨", "朱", "秦", "尤", "许", "何", "吕", "施", "张", "孔", "曹", "严", "华", "金", "魏",
	"陶", "姜", "戚", "谢", "邹", "喻", "柏", "水", "窦", "章", "云", "苏", "潘", "葛", "奚", "范", "彭",
	"郎", "鲁", "韦", "昌", "马", "苗", "凤", "花", "方", "任", "袁", "柳", "鲍", "史", "唐", "费", "薛",
	"雷", "贺", "倪", "汤", "滕", "殷", "罗", "毕", "郝", "安", "常", "傅", "卞", "齐", "元", "顾", "孟",
	"平", "黄", "穆", "萧", "尹", "姚", "邵", "湛", "汪", "祁", "毛", "狄", "米", "伏", "成", "戴", "谈",
	"宋", "茅", "庞", "熊", "纪", "舒", "屈", "项", "祝", "董", "梁", "杜", "阮", "蓝", "闵", "季", "贾",
	"路", "娄", "江", "童", "颜", "郭", "梅", "盛", "林", "钟", "徐", "邱", "骆", "高", "夏", "蔡", "田",
	"樊", "胡", "凌", "霍", "虞", "万", "支", "柯", "管", "卢", "莫", "柯", "房", "裘", "缪", "解", "应",
	"宗", "丁", "宣", "邓", "单", "杭", "洪", "包", "诸", "左", "石", "崔", "吉", "龚", "程", "嵇", "邢",
	"裴", "陆", "荣", "翁", "荀", "于", "惠", "甄", "曲", "封", "储", "仲", "伊", "宁", "仇", "甘", "武",
	"符", "刘", "景", "詹", "龙", "叶", "幸", "司", "黎", "溥", "印", "怀", "蒲", "邰", "从", "索", "赖",
	"卓", "屠", "池", "乔", "胥", "闻", "莘", "党", "翟", "谭", "贡", "劳", "逄", "姬", "申", "扶", "堵",
	"冉", "宰", "雍", "桑", "寿", "通", "燕", "浦", "尚", "农", "温", "别", "庄", "晏", "柴", "瞿", "阎",
	"连", "习", "容", "向", "古", "易", "廖", "庾", "终", "步", "都", "耿", "满", "弘", "匡", "国", "文",
	"寇", "广", "禄", "阙", "东", "欧", "利", "师", "巩", "聂", "关", "荆", "司马", "上官", "欧阳", "夏侯",
	"诸葛", "闻人", "东方", "赫连", "皇甫", "尉迟", "公羊", "澹台", "公冶", "宗政", "濮阳", "淳于", "单于",
	"太叔", "申屠", "公孙", "仲孙", "轩辕", "令狐", "徐离", "宇文", "长孙", "慕容", "司徒", "司空"}
var firstName = []string{
	"伟", "刚", "勇", "毅", "俊", "峰", "强", "军", "平", "保", "东", "文", "辉", "力", "明", "永", "健", "世", "广", "志", "义",
	"兴", "良", "海", "山", "仁", "波", "宁", "贵", "福", "生", "龙", "元", "全", "国", "胜", "学", "祥", "才", "发", "武", "新",
	"利", "清", "飞", "彬", "富", "顺", "信", "子", "杰", "涛", "昌", "成", "康", "星", "光", "天", "达", "安", "岩", "中", "茂",
	"进", "林", "有", "坚", "和", "彪", "博", "诚", "先", "敬", "震", "振", "壮", "会", "思", "群", "豪", "心", "邦", "承", "乐",
	"绍", "功", "松", "善", "厚", "庆", "磊", "民", "友", "裕", "河", "哲", "江", "超", "浩", "亮", "政", "谦", "亨", "奇", "固",
	"之", "轮", "翰", "朗", "伯", "宏", "言", "若", "鸣", "朋", "斌", "梁", "栋", "维", "启", "克", "伦", "翔", "旭", "鹏", "泽",
	"晨", "辰", "士", "以", "建", "家", "致", "树", "炎", "德", "行", "时", "泰", "盛", "雄", "琛", "钧", "冠", "策", "腾", "楠",
	"榕", "风", "航", "弘", "秀", "娟", "英", "华", "慧", "巧", "美", "娜", "静", "淑", "惠", "珠", "翠", "雅", "芝", "玉", "萍",
	"红", "娥", "玲", "芬", "芳", "燕", "彩", "春", "菊", "兰", "凤", "洁", "梅", "琳", "素", "云", "莲", "真", "环", "雪", "荣",
	"爱", "妹", "霞", "香", "月", "莺", "媛", "艳", "瑞", "凡", "佳", "嘉", "琼", "勤", "珍", "贞", "莉", "桂", "娣", "叶", "璧",
	"璐", "娅", "琦", "晶", "妍", "茜", "秋", "珊", "莎", "锦", "黛", "青", "倩", "婷", "姣", "婉", "娴", "瑾", "颖", "露", "瑶",
	"怡", "婵", "雁", "蓓", "纨", "仪", "荷", "丹", "蓉", "眉", "君", "琴", "蕊", "薇", "菁", "梦", "岚", "苑", "婕", "馨", "瑗",
	"琰", "韵", "融", "园", "艺", "咏", "卿", "聪", "澜", "纯", "毓", "悦", "昭", "冰", "爽", "琬", "茗", "羽", "希", "欣", "飘",
	"育", "滢", "馥", "筠", "柔", "竹", "霭", "凝", "晓", "欢", "霄", "枫", "芸", "菲", "寒", "伊", "亚", "宜", "可", "姬", "舒",
	"影", "荔", "枝", "丽", "阳", "妮", "宝", "贝", "初", "程", "梵", "罡", "恒", "鸿", "桦", "骅", "剑", "娇", "纪", "宽", "苛",
	"灵", "玛", "媚", "琪", "晴", "容", "睿", "烁", "堂", "唯", "威", "韦", "雯", "苇", "萱", "阅", "彦", "宇", "雨", "洋", "忠",
	"宗", "曼", "紫", "逸", "贤", "蝶", "菡", "绿", "蓝", "儿", "翠", "烟", "小", "轩"}
var lastNameLen = len(lastName)
var firstNameLen = len(firstName)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func GetFullName() string {
	var first string
	for i := 0; i <= rand.Intn(1); i++ {
		first = fmt.Sprint(firstName[rand.Intn(firstNameLen-1)])
	}
	return fmt.Sprintf("%s%s", fmt.Sprint(lastName[rand.Intn(lastNameLen-1)]), first)
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

const (
	baseUrlLocal = "http://localhost:52001/api"
)

const (
	verifyCodeSend          = "/verify_code/send"
	registerUser            = "/register"
	loginUserWithVerifyCode = "/login/verify_code"
	loginUserWithPwd        = "/login/pwd"
	userPwdReset            = "/user/password/reset"
	userInfo                = "/user/user_info"
	userInfoList            = "/user/user_info/list"
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
	searchUserInfo          = "/search/user_info"
	searchTradeOrder        = "/search/trade_order"
	searchMerchantInfo      = "/search/merchant_info"
	reportOrder             = "/user/order/report"
	rankOrderShop           = "/user/order/rank/shop"
	rankOrderSku            = "/user/order/rank/sku"
	userAccountCharge       = "/user/account/charge"
	commentsOrderCreate     = "/user/comments/order/create"
	commentsShopList        = "/user/comments/shop/list"
	commentsTagsModify      = "/user/comments/tags/modify"
	commentsTagsList        = "/user/comments/tags/list"
)

const (
	apiV1 = "/v1"
)

var apiVersion = apiV1

var baseUrl = baseUrlLocal + apiVersion

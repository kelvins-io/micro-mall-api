package code

import "gitee.com/kelvins-io/common/errcode"

const (
	Success              = 29000000
	ErrorServer          = 29000001
	UserNotExist         = 29000005
	UserExist            = 29000006
	DBDuplicateEntry     = 29000007
	MerchantExist        = 29000008
	MerchantNotExist     = 29000009
	ShopBusinessExist    = 29000010
	ShopBusinessNotExist = 29000011
	SkuCodeEmpty         = 29000012
	SkuCodeNotExist      = 29000013
	SkuCodeExist         = 29000014
	SkuAmountNotEnough   = 29000015
	TransactionFailed    = 29000016
)

var ErrMap = make(map[int]string)

func init() {
	dict := map[int]string{
		Success:              "OK",
		ErrorServer:          "服务器错误",
		UserNotExist:         "用户不存在",
		DBDuplicateEntry:     "Duplicate entry",
		UserExist:            "已存在用户记录，请勿重复创建",
		MerchantExist:        "商户认证材料已存在",
		MerchantNotExist:     "商户未提交材料",
		ShopBusinessExist:    "店铺申请材料已存在",
		ShopBusinessNotExist: "商户未提交店铺材料",
		SkuCodeEmpty:         "商品唯一code为空",
		SkuCodeNotExist:      "商品唯一code在系统找不到",
		SkuCodeExist:         "商品唯一code已存在系统",
		SkuAmountNotEnough:   "商品库存不够",
		TransactionFailed:    "事务执行失败",
	}
	errcode.RegisterErrMsgDict(dict)
	for key, _ := range dict {
		ErrMap[key] = dict[key]
	}
}

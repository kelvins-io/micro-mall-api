package models

import (
	"time"
)

type Account struct {
	AccountCode string    `xorm:"not null pk comment('账户主键') CHAR(50)"`
	Owner       string    `xorm:"not null comment('账户所有者') unique(account_index) CHAR(36)"`
	Balance     string    `xorm:"comment('账户余额') DECIMAL(32,16)"`
	CoinType    int       `xorm:"not null default 1 comment('币种类型，1-rmb，2-usdt') unique(account_index) TINYINT"`
	CoinDesc    string    `xorm:"comment('币种描述') VARCHAR(64)"`
	State       int       `xorm:"comment('状态，1无效，2锁定，3正常') TINYINT"`
	AccountType int       `xorm:"not null comment('账户类型，1-个人账户，2-公司账户，3-系统账户') unique(account_index) TINYINT"`
	CreateTime  time.Time `xorm:"not null default CURRENT_TIMESTAMP comment('创建时间') index DATETIME"`
	UpdateTime  time.Time `xorm:"not null default CURRENT_TIMESTAMP comment('更新时间') DATETIME"`
}

type ConfigKvStore struct {
	Id          int       `xorm:"not null pk autoincr comment('主键') INT"`
	ConfigKey   string    `xorm:"not null comment('配置键') unique VARCHAR(255)"`
	ConfigValue string    `xorm:"not null comment('配置值') VARCHAR(255)"`
	Prefix      string    `xorm:"not null comment('配置前缀') VARCHAR(255)"`
	Suffix      string    `xorm:"not null comment('配置后缀') VARCHAR(255)"`
	Status      int       `xorm:"not null default 1 comment('是否启用 1是 0否') TINYINT"`
	IsDelete    int       `xorm:"not null default 0 comment('是否删除 1是 0否') TINYINT"`
	CreateTime  time.Time `xorm:"default CURRENT_TIMESTAMP comment('创建时间') DATETIME"`
	UpdateTime  time.Time `xorm:"default CURRENT_TIMESTAMP comment('更新时间') DATETIME"`
}

type Merchant struct {
	MerchantId   int64     `xorm:"not null pk autoincr comment('商户号ID') BIGINT"`
	MerchantCode string    `xorm:"not null comment('商户唯一code') index CHAR(36)"`
	Uid          int64     `xorm:"not null comment('用户ID') unique BIGINT"`
	RegisterAddr string    `xorm:"not null comment('注册地址') TEXT"`
	HealthCardNo string    `xorm:"not null comment('健康证号') CHAR(30)"`
	Identity     int       `xorm:"comment('身份属性，1-临时店员，2-正式店员，3-经理，4-店长') TINYINT"`
	State        int       `xorm:"comment('状态，0-未审核，1-审核中，2-审核不通过，3-已审核') TINYINT"`
	TaxCardNo    string    `xorm:"comment('纳税账户号') CHAR(30)"`
	CreateTime   time.Time `xorm:"not null default CURRENT_TIMESTAMP comment('创建时间') DATETIME"`
	UpdateTime   time.Time `xorm:"not null default CURRENT_TIMESTAMP comment('修改时间') DATETIME"`
}

type Order struct {
	Id           int64     `xorm:"pk autoincr comment('自增ID') BIGINT"`
	TxCode       string    `xorm:"not null comment('交易号') unique(tx_code_order_code_index) CHAR(40)"`
	OrderCode    string    `xorm:"not null comment('订单code') unique unique(tx_code_order_code_index) CHAR(40)"`
	Uid          int64     `xorm:"not null comment('用户UID') index BIGINT"`
	OrderTime    time.Time `xorm:"not null default CURRENT_TIMESTAMP comment('下单时间') index DATETIME"`
	Description  string    `xorm:"comment('订单描述') index VARCHAR(255)"`
	ClientIp     string    `xorm:"comment('客户端IP') CHAR(16)"`
	DeviceCode   string    `xorm:"comment('客户端设备code') VARCHAR(512)"`
	ShopId       int64     `xorm:"not null comment('门店ID') index BIGINT"`
	ShopName     string    `xorm:"not null comment('门店名称') index VARCHAR(255)"`
	ShopAreaCode string    `xorm:"comment('门店区域编号') VARCHAR(255)"`
	ShopAddress  string    `xorm:"comment('门店地址') TEXT"`
	State        int       `xorm:"not null default 0 comment('订单状态，0-有效，1-锁定中，2-无效') TINYINT"`
	PayExpire    time.Time `xorm:"not null comment('支付有效期，默认30分钟内有效') DATETIME"`
	PayState     int       `xorm:"not null default 0 comment('支付状态，0-未支付，1-支付中，2-支付失败，3-已支付') TINYINT"`
	Amount       int       `xorm:"comment('订单关联商品数量') INT"`
	TotalAmount  string    `xorm:"not null default 0.0000000000000000 comment('订单总金额') DECIMAL(32,16)"`
	CoinType     int       `xorm:"default 1 comment(' 订单币种，1-CNY，2-USD') TINYINT"`
	CreateTime   time.Time `xorm:"not null default CURRENT_TIMESTAMP comment('创建时间') DATETIME"`
	UpdateTime   time.Time `xorm:"not null default CURRENT_TIMESTAMP comment('修改时间') DATETIME"`
}

type OrderSku struct {
	Id         int64     `xorm:"pk autoincr comment('自增ID') BIGINT"`
	OrderCode  string    `xorm:"not null comment('对应订单code') unique(order_unique) CHAR(64)"`
	ShopId     int64     `xorm:"not null comment('店铺ID') unique(order_unique) index BIGINT"`
	SkuCode    string    `xorm:"not null comment('商品sku') unique(order_unique) index CHAR(64)"`
	Price      string    `xorm:"not null default 0.0000000000000000 comment('商品单价') DECIMAL(32,16)"`
	Amount     int       `xorm:"not null comment('商品数量') INT"`
	Name       string    `xorm:"comment('商品名称') index VARCHAR(255)"`
	CreateTime time.Time `xorm:"not null default CURRENT_TIMESTAMP comment('创建时间') DATETIME"`
	UpdateTime time.Time `xorm:"not null default CURRENT_TIMESTAMP comment('修改时间') DATETIME"`
}

type PayRecord struct {
	Id          int64     `xorm:"pk autoincr comment('自增ID') BIGINT"`
	PayId       string    `xorm:"not null comment('账单ID') unique CHAR(40)"`
	OutTradeNo  string    `xorm:"not null comment('外部商户订单号') index CHAR(40)"`
	TimeExpire  time.Time `xorm:"not null default CURRENT_TIMESTAMP comment('支付过期时间') DATETIME"`
	NotifyUrl   string    `xorm:"comment('交易结果通知地址') VARCHAR(255)"`
	Description string    `xorm:"comment('交易描述') VARCHAR(255)"`
	Merchant    string    `xorm:"not null comment('交易商户ID') index CHAR(40)"`
	Attach      string    `xorm:"comment('交易留言') TEXT"`
	User        string    `xorm:"not null comment('交易用户ID') index CHAR(40)"`
	Amount      string    `xorm:"not null comment('交易数量') DECIMAL(32,16)"`
	CoinType    int       `xorm:"not null default 0 comment('交易币种，0-cny,1-usd') TINYINT"`
	Reduction   string    `xorm:"comment('满减优惠') DECIMAL(32,16)"`
	PayType     int       `xorm:"not null comment('交易类型，1入账，2退款') TINYINT"`
	PayState    int       `xorm:"comment('支付状态，0-未支付，1-支付中，2-支付失败，3-支付成功') TINYINT"`
	CreateTime  time.Time `xorm:"not null default CURRENT_TIMESTAMP comment('创建时间') DATETIME"`
	UpdateTime  time.Time `xorm:"not null default CURRENT_TIMESTAMP comment('修改时间') DATETIME"`
}

type ShopBusiness struct {
	ShopId           int64     `xorm:"not null pk autoincr comment('店铺ID') BIGINT"`
	NickName         string    `xorm:"not null comment('简称') unique(legal_person_nick_name_index) VARCHAR(512)"`
	ShopCode         string    `xorm:"not null comment('店铺唯一code') unique CHAR(36)"`
	FullName         string    `xorm:"not null comment('店铺全称') TEXT"`
	RegisterAddr     string    `xorm:"not null comment('注册地址') TEXT"`
	BusinessAddr     string    `xorm:"not null comment('实际经营地址') TEXT"`
	LegalPerson      int64     `xorm:"not null comment('店铺法人') index unique(legal_person_nick_name_index) BIGINT"`
	BusinessLicense  string    `xorm:"not null comment('经营许可证') CHAR(36)"`
	TaxCardNo        string    `xorm:"not null comment('纳税号') CHAR(36)"`
	BusinessDesc     string    `xorm:"not null comment('经营描述') TEXT"`
	SocialCreditCode string    `xorm:"not null comment('统一社会信用代码') CHAR(36)"`
	OrganizationCode string    `xorm:"not null comment('组织机构代码') CHAR(36)"`
	State            int       `xorm:"not null default 0 comment('状态，0-未审核，1-审核不通过，2-审核通过') TINYINT"`
	CreateTime       time.Time `xorm:"not null default CURRENT_TIMESTAMP comment('创建时间') DATETIME"`
	UpdateTime       time.Time `xorm:"not null default CURRENT_TIMESTAMP comment('修改时间') DATETIME"`
}

type SkuInventory struct {
	Id         int64     `xorm:"pk autoincr comment('商品库存ID') BIGINT"`
	SkuCode    string    `xorm:"not null comment('商品编码') unique unique(sku_code_shop_id_index) CHAR(64)"`
	Amount     int64     `xorm:"comment('库存数量') BIGINT"`
	Price      string    `xorm:"comment('入库单价') DECIMAL(32,16)"`
	ShopId     int64     `xorm:"not null comment('所属店铺ID') index unique(sku_code_shop_id_index) BIGINT"`
	CreateTime time.Time `xorm:"not null default CURRENT_TIMESTAMP comment('创建时间') DATETIME"`
	UpdateTime time.Time `xorm:"not null default CURRENT_TIMESTAMP comment('修改时间') DATETIME"`
}

type SkuPriceHistory struct {
	Id         int64     `xorm:"pk autoincr comment('自增ID') BIGINT"`
	ShopId     int64     `xorm:"not null comment('调价的店铺id') unique(shop_id_sku_code_index) BIGINT"`
	SkuCode    string    `xorm:"not null comment('商品sku_code') unique(shop_id_sku_code_index) index CHAR(40)"`
	Price      string    `xorm:"not null comment('商品价格') DECIMAL(32,16)"`
	Tsp        int       `xorm:"not null comment('价格变化时的时间戳') index INT"`
	Reason     string    `xorm:"comment('调价说明') TEXT"`
	CreateTime time.Time `xorm:"not null default CURRENT_TIMESTAMP comment('创建时间') DATETIME"`
	UpdateTime time.Time `xorm:"not null default CURRENT_TIMESTAMP comment('更新时间') DATETIME"`
	OpUid      int64     `xorm:"comment('操作员UID') BIGINT"`
	OpIp       string    `xorm:"comment('操作员IP') CHAR(16)"`
}

type SkuProperty struct {
	Id            int64     `xorm:"pk autoincr comment('ID') BIGINT"`
	Code          string    `xorm:"not null comment('商品唯一编号') index CHAR(64)"`
	Price         string    `xorm:"comment('商品当前价格') DECIMAL(32,16)"`
	Name          string    `xorm:"comment('商品名称') index VARCHAR(255)"`
	Desc          string    `xorm:"comment('商品描述') TEXT"`
	Production    string    `xorm:"comment('生产企业') VARCHAR(1024)"`
	Supplier      string    `xorm:"comment('供应商') VARCHAR(1024)"`
	Category      int       `xorm:"comment('商品类别') INT"`
	Title         string    `xorm:"comment('商品标题') VARCHAR(255)"`
	SubTitle      string    `xorm:"comment('商品副标题') VARCHAR(255)"`
	Color         string    `xorm:"comment('商品颜色') VARCHAR(64)"`
	ColorCode     int       `xorm:"comment('商品颜色代码') INT"`
	Specification string    `xorm:"comment('商品规格') VARCHAR(255)"`
	DescLink      string    `xorm:"comment('商品介绍链接') VARCHAR(255)"`
	State         int       `xorm:"default 0 comment('商品状态，0-有效，1-无效，2-锁定') TINYINT"`
	CreateTime    time.Time `xorm:"not null comment('创建时间') DATETIME"`
	UpdateTime    time.Time `xorm:"not null comment('更新时间') DATETIME"`
}

type Transaction struct {
	Id              int64     `xorm:"pk comment('交易ID') BIGINT"`
	FromAccountCode string    `xorm:"not null default '0' comment('转出账户ID') CHAR(36)"`
	FromBalance     string    `xorm:"default 0.0000000000000000 comment('转出后账户余额') DECIMAL(32,16)"`
	ToAccountCode   string    `xorm:"not null default '0' comment('转入账户ID') CHAR(36)"`
	ToBalance       string    `xorm:"comment('转入后账户余额') DECIMAL(32,16)"`
	Amount          string    `xorm:"comment('交易金额') DECIMAL(32,16)"`
	Meta            string    `xorm:"comment('转账说明') VARCHAR(255)"`
	Scene           string    `xorm:"comment('支付场景') VARCHAR(64)"`
	OpUid           int64     `xorm:"not null comment('操作用户UID') BIGINT"`
	OpIp            string    `xorm:"comment('操作的IP') VARCHAR(16)"`
	TxId            string    `xorm:"comment('对应交易号') CHAR(36)"`
	Fingerprint     string    `xorm:"not null comment('防篡改指纹') VARCHAR(32)"`
	PayType         int       `xorm:"default 0 comment('支付方式，0系统操作，1-银行卡，2-信用卡,3-支付宝,4-微信支付,5-京东支付') TINYINT"`
	PayDesc         string    `xorm:"comment('支付方式描述') VARCHAR(36)"`
	CreateTime      time.Time `xorm:"not null default CURRENT_TIMESTAMP comment('创建时间') DATETIME"`
	UpdateTime      time.Time `xorm:"not null default CURRENT_TIMESTAMP comment('修改时间') DATETIME"`
}

type User struct {
	Id           int64     `xorm:"pk autoincr comment('自增ID') BIGINT"`
	AccountId    string    `xorm:"not null comment('账户ID，全局唯一') unique CHAR(36)"`
	UserName     string    `xorm:"not null comment('用户名') index VARCHAR(255)"`
	Password     string    `xorm:"not null comment('用户密码md5值') VARCHAR(255)"`
	PasswordSalt string    `xorm:"comment('密码salt值') VARCHAR(255)"`
	Sex          int       `xorm:"comment('性别，1-男，2-女') TINYINT(1)"`
	Phone        string    `xorm:"comment('手机号') unique(country_code_phone_index) CHAR(11)"`
	CountryCode  string    `xorm:"comment('手机区号') unique(country_code_phone_index) CHAR(5)"`
	Email        string    `xorm:"comment('邮箱') index VARCHAR(255)"`
	State        int       `xorm:"comment('状态，0-未激活，1-审核中，2-审核未通过，3-已审核') TINYINT(1)"`
	IdCardNo     string    `xorm:"comment('身份证号') unique CHAR(18)"`
	Inviter      int64     `xorm:"comment('邀请人uid') BIGINT"`
	InviteCode   string    `xorm:"comment('邀请码') CHAR(20)"`
	CreateTime   time.Time `xorm:"not null default CURRENT_TIMESTAMP comment('创建时间') DATETIME"`
	UpdateTime   time.Time `xorm:"not null default CURRENT_TIMESTAMP comment('修改时间') DATETIME"`
	ContactAddr  string    `xorm:"comment('联系地址') TEXT"`
	Age          int       `xorm:"comment('年龄') INT"`
}

type UserTrolley struct {
	Id         int64     `xorm:"pk autoincr comment('自增ID') BIGINT"`
	Uid        int64     `xorm:"not null comment('用户ID') index(shop_id_sku_uid_index) BIGINT"`
	ShopId     int64     `xorm:"not null comment('店铺ID') index(shop_id_sku_index) index(shop_id_sku_uid_index) BIGINT"`
	SkuCode    string    `xorm:"not null comment('商品sku') index(shop_id_sku_index) index(shop_id_sku_uid_index) index CHAR(40)"`
	Count      int       `xorm:"not null default 1 comment('商品数量') INT"`
	JoinTime   time.Time `xorm:"not null default CURRENT_TIMESTAMP comment('加入时间') DATETIME"`
	Selected   int       `xorm:"default 1 comment('是否选中，1-未选中，2-选中') TINYINT(1)"`
	State      int       `xorm:"default 1 comment('状态，1-有效，2-移除') TINYINT"`
	CreateTime time.Time `xorm:"not null default CURRENT_TIMESTAMP comment('创建时间') DATETIME"`
	UpdateTime time.Time `xorm:"not null default CURRENT_TIMESTAMP comment('更新时间') DATETIME"`
}

type VerifyCodeRecord struct {
	Id           int64     `xorm:"pk autoincr comment('自增id') BIGINT"`
	Uid          int64     `xorm:"not null comment('用户UID') BIGINT"`
	BusinessType int       `xorm:"comment('验证类型，1-注册登录，2-购买商品') TINYINT"`
	VerifyCode   string    `xorm:"comment('验证码') index CHAR(6)"`
	Expire       int       `xorm:"comment('过期时间unix') INT"`
	CountryCode  string    `xorm:"comment('验证码下发手机国际码') index(country_code_phone_index) CHAR(5)"`
	Phone        string    `xorm:"comment('验证码下发手机号') index(country_code_phone_index) CHAR(11)"`
	Email        string    `xorm:"comment('验证码下发邮箱') index VARCHAR(255)"`
	CreateTime   time.Time `xorm:"not null default CURRENT_TIMESTAMP comment('创建时间') DATETIME"`
	UpdateTime   time.Time `xorm:"not null default CURRENT_TIMESTAMP comment('修改时间') DATETIME"`
}

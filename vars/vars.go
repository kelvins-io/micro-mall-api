package vars

import (
	"gitee.com/cristiane/micro-mall-api/config/setting"
	"gitee.com/cristiane/micro-mall-api/pkg/util/goroutine"
	"gitee.com/kelvins-io/common/log"
	"gitee.com/kelvins-io/g2cache"
)

var (
	App *WEBApplication
	//DBEngineXORM     xorm.EngineInterface
	//DBEngineGORM     *gorm.DB
	LoggerSetting    *setting.LoggerSettingS
	AccessLogger     log.LoggerContextIface
	ErrorLogger      log.LoggerContextIface
	BusinessLogger   log.LoggerContextIface
	ServerSetting    *setting.ServerSettingS
	RateLimitSetting *setting.RateLimitSettingS
	JwtSetting       *setting.JwtSettingS
	//MysqlSettingMicroMall *setting.MysqlSettingS
	//RedisSettingMicroMall *setting.RedisSettingS
	G2CacheSetting     *setting.G2CacheSettingS
	EmailConfigSetting *EmailConfigSettingS
	//RedisPoolMicroMall    *redis.Pool
	GPool         *goroutine.Pool
	G2CacheEngine *g2cache.G2Cache
	AppCloseCh    <-chan struct{}
	LoggerLevel   string
	Environment   string
)

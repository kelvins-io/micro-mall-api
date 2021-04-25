package vars

import (
	"gitee.com/cristiane/micro-mall-api/config/setting"
	"gitee.com/cristiane/micro-mall-api/pkg/util/goroutine"
	"github.com/gomodule/redigo/redis"
	"github.com/jinzhu/gorm"
	"xorm.io/xorm"

	"gitee.com/kelvins-io/common/log"
	"gitee.com/kelvins-io/g2cache"
)

var (
	App                   *WEBApplication
	DBEngineXORM          xorm.EngineInterface
	DBEngineGORM          *gorm.DB
	LoggerSetting         *setting.LoggerSettingS
	AccessLogger          log.LoggerContextIface
	ErrorLogger           log.LoggerContextIface
	BusinessLogger        log.LoggerContextIface
	ServerSetting         *setting.ServerSettingS
	JwtSetting            *setting.JwtSettingS
	MysqlSettingMicroMall *setting.MysqlSettingS
	RedisSettingMicroMall *setting.RedisSettingS
	G2CacheSetting		  *setting.G2CacheSettingS
	EmailConfigSetting    *EmailConfigSettingS
	VerifyCodeSetting     *VerifyCodeSettingS
	RedisPoolMicroMall    *redis.Pool
	GPool                 *goroutine.Pool
	G2CacheEngine		  *g2cache.G2Cache
)

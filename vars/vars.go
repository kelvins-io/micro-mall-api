package vars

import (
	"gitee.com/cristiane/micro-mall-api/config/setting"
	"gitee.com/cristiane/micro-mall-api/pkg/util/groutine"
	"github.com/gomodule/redigo/redis"
	"github.com/jinzhu/gorm"
	"xorm.io/xorm"

	"gitee.com/kelvins-io/common/log"
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
	EmailConfigSetting    *EmailConfigSettingS
	VerifyCodeSetting     *VerifyCodeSettingS
	RedisPoolMicroMall    *redis.Pool
	GPool                 *goroute.Pool
)

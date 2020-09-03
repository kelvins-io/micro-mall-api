package vars

import (
	"gitee.com/cristiane/micro-mall-api/config/setting"
	"github.com/gomodule/redigo/redis"
	"github.com/jinzhu/gorm"
	"net/http"
	"time"
	"xorm.io/xorm"

	"gitee.com/cristiane/go-common/log"
)

var (
	App               *WEBApplication
	DBEngineXORM      xorm.EngineInterface
	DBEngineGORM      *gorm.DB
	LoggerSetting     *setting.LoggerSettingS
	AccessLogger      log.LoggerContextIface
	ErrorLogger       log.LoggerContextIface
	BusinessLogger    log.LoggerContextIface
	ServerSetting     *setting.ServerSettingS
	JwtSetting        *setting.JwtSettingS
	MysqlSettingXXXDB *setting.MysqlSettingS
	RedisSetting      *setting.RedisSettingS
	RedisPool         *redis.Pool
	HttpClient        = &http.Client{Timeout: 30 * time.Second}
)

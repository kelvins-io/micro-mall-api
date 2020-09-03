package startup

import (
	"gitee.com/cristiane/micro-mall-api/config"
	"gitee.com/cristiane/micro-mall-api/config/setting"
	"gitee.com/cristiane/micro-mall-api/vars"
	"log"
)

const (
	SectionMysqlXXXDB = "product-mysql"
	SectionRedisXXX   = "product-redis"
)

// LoadConfig 加载自定义配置项
func LoadConfig() error {

	// 外部MySQL数据源
	log.Printf("[info] Load default config %s", SectionMysqlXXXDB)
	vars.MysqlSettingXXXDB = new(setting.MysqlSettingS)
	config.MapConfig(SectionMysqlXXXDB, vars.MysqlSettingXXXDB)

	// 加载外部Redis数据源
	log.Printf("[info] Load default config %s", SectionRedisXXX)
	vars.RedisSetting = new(setting.RedisSettingS)
	config.MapConfig(SectionRedisXXX, vars.RedisSetting)
	return nil
}

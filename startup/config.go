package startup

import (
	"gitee.com/cristiane/micro-mall-api/config"
	"gitee.com/cristiane/micro-mall-api/config/setting"
	"gitee.com/cristiane/micro-mall-api/vars"
	"log"
)

const (
	SectionMysqlMicroMall          = "micro-mall-mysql"
	SectionRedisMicroMall          = "micro-mall-redis"
	SectionEmailConfig             = "email-config"
	SectionQueueUserRegisterNotice = "queue-user-register-notice"
	SectionQueueUserStateNotice    = "queue-user-state-notice"
)

// LoadConfig 加载自定义配置项
func LoadConfig() error {

	// 外部MySQL数据源
	log.Printf("[info] Load default config %s", SectionMysqlMicroMall)
	vars.MysqlSettingMicroMall = new(setting.MysqlSettingS)
	config.MapConfig(SectionMysqlMicroMall, vars.MysqlSettingMicroMall)

	// 加载外部Redis数据源
	log.Printf("[info] Load default config %s", SectionRedisMicroMall)
	vars.RedisSettingMicroMall = new(setting.RedisSettingS)
	config.MapConfig(SectionRedisMicroMall, vars.RedisSettingMicroMall)
	// 加载email数据源
	log.Printf("[info] Load default config %s", SectionEmailConfig)
	vars.EmailConfigSetting = new(vars.EmailConfigSettingS)
	config.MapConfig(SectionEmailConfig, vars.EmailConfigSetting)
	// 用户注册通知
	log.Printf("[info] Load default config %s", SectionQueueUserRegisterNotice)
	vars.QueueAMQPSettingUserRegisterNotice = new(setting.QueueAMQPSettingS)
	config.MapConfig(SectionQueueUserRegisterNotice, vars.QueueAMQPSettingUserRegisterNotice)
	// 用户事件通知
	log.Printf("[info] Load default config %s", SectionQueueUserStateNotice)
	vars.QueueAMQPSettingUserStateNotice = new(setting.QueueAMQPSettingS)
	config.MapConfig(SectionQueueUserStateNotice, vars.QueueAMQPSettingUserStateNotice)

	return nil
}

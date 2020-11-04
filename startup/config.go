package startup

import (
	"gitee.com/cristiane/micro-mall-api/config"
	"gitee.com/cristiane/micro-mall-api/config/setting"
	"gitee.com/cristiane/micro-mall-api/vars"
	"log"
)

const (
	SectionMysqlMicroMall = "micro-mall-mysql"
	SectionRedisMicroMall = "micro-mall-redis"
	SectionEmailConfig    = "email-config"
	SectionVerifyCode     = "micro-mall-verify_code"
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
	// 加载验证码配置
	log.Printf("[info] Load default config %s", SectionVerifyCode)
	vars.VerifyCodeSetting = new(vars.VerifyCodeSettingS)
	config.MapConfig(SectionVerifyCode, vars.VerifyCodeSetting)

	return nil
}

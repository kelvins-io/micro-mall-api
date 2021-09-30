package startup

import (
	"gitee.com/cristiane/micro-mall-api/config"
	"gitee.com/cristiane/micro-mall-api/config/setting"
	"gitee.com/cristiane/micro-mall-api/vars"
)

const (
	SectionMysqlMicroMall = "micro-mall-mysql"
	SectionRedisMicroMall = "micro-mall-redis"
	SectionEmailConfig    = "email-config"
	SectionVerifyCode     = "micro-mall-verify_code"
	SectionG2Cache        = "micro-mall-g2cache"
)

// LoadConfig 加载自定义配置项
func LoadConfig() error {

	// 外部MySQL数据源
	vars.MysqlSettingMicroMall = new(setting.MysqlSettingS)
	config.MapConfig(SectionMysqlMicroMall, vars.MysqlSettingMicroMall)
	// 加载外部Redis数据源
	vars.RedisSettingMicroMall = new(setting.RedisSettingS)
	config.MapConfig(SectionRedisMicroMall, vars.RedisSettingMicroMall)
	//加载G2Cache二级缓存配置
	vars.G2CacheSetting = new(setting.G2CacheSettingS)
	config.MapConfig(SectionG2Cache, vars.G2CacheSetting)
	// 加载email数据源
	vars.EmailConfigSetting = new(vars.EmailConfigSettingS)
	config.MapConfig(SectionEmailConfig, vars.EmailConfigSetting)
	// 加载验证码配置
	vars.VerifyCodeSetting = new(vars.VerifyCodeSettingS)
	config.MapConfig(SectionVerifyCode, vars.VerifyCodeSetting)

	return nil
}

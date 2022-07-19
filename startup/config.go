package startup

import (
	"gitee.com/cristiane/micro-mall-api/config"
	"gitee.com/cristiane/micro-mall-api/config/setting"
	"gitee.com/cristiane/micro-mall-api/vars"
)

const (
	SectionEmailConfig = "email-config"
	SectionG2Cache     = "micro-mall-g2cache"
)

// LoadConfig 加载自定义配置项
func LoadConfig() error {
	//加载G2Cache二级缓存配置
	vars.G2CacheSetting = new(setting.G2CacheSettingS)
	config.MapConfig(SectionG2Cache, vars.G2CacheSetting)
	// 加载email数据源
	vars.EmailConfigSetting = new(vars.EmailConfigSettingS)
	config.MapConfig(SectionEmailConfig, vars.EmailConfigSetting)

	return nil
}

package config

import (
	"gitee.com/cristiane/micro-mall-api/internal/config"
	"log"
	"strings"
)

// 对外提供加载自定义配置功能入口
func MapConfig(section string, v interface{}) {
	// 检测是否加载了系统级别配置
	if strings.HasPrefix(section, "web-") {
		log.Fatalf("[err] section name can't have web- prefix")
	}
	config.MapConfig(section, v)
}

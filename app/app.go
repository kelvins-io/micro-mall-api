package app

import (
	"flag"
	"fmt"
	"gitee.com/kelvins-io/common/log"
	"gitee.com/cristiane/micro-mall-api/internal/setup"
	"gitee.com/cristiane/micro-mall-api/vars"
)

var (
	port       = flag.Int64("p", 0, "Set server port.")
	loggerPath = flag.String("logger_path", "", "Set Logger Root Path.")
)

// 初始化application--日志部分
func initApplication(application *vars.Application) error {
	const DefaultLoggerRootPath = "./logs"
	const DefaultLoggerLevel = "debug"

	rootPath := DefaultLoggerRootPath
	if vars.LoggerSetting != nil && vars.LoggerSetting.RootPath != "" {
		rootPath = vars.LoggerSetting.RootPath
	}
	loggerLevel := DefaultLoggerLevel
	if vars.LoggerSetting != nil && vars.LoggerSetting.Level != "" {
		loggerLevel = vars.LoggerSetting.Level
	}

	err := log.InitGlobalConfig(rootPath, loggerLevel, application.Name)
	if err != nil {
		return fmt.Errorf("log.InitGlobalConfig: %v", err)
	}

	return nil
}

// 初始化全局配置
func setupCommonVars(application *vars.WEBApplication) error {
	if vars.ServerSetting != nil {
		vars.App.EndPort = vars.ServerSetting.EndPort
		vars.App.MonitorEndPort = vars.ServerSetting.MonitorEndPort
		if vars.ServerSetting.MonitorEndPort != 0 {
			application.Mux = setup.NewServerMux()
		}
	}
	return nil
}

package app

import (
	"flag"
	"fmt"
	"os"
	"sync"
	"time"

	"gitee.com/cristiane/micro-mall-api/config/setting"
	"gitee.com/cristiane/micro-mall-api/internal/config"
	"gitee.com/cristiane/micro-mall-api/internal/logging"
	"gitee.com/cristiane/micro-mall-api/internal/util/startup"
	varsInternal "gitee.com/cristiane/micro-mall-api/internal/vars"
	"gitee.com/cristiane/micro-mall-api/vars"
	"gitee.com/kelvins-io/common/log"
)

var (
	flagLoggerLevel = flag.String("logger_level", "", "set logger level eg: debug,warn,error,info")
	flagLoggerPath  = flag.String("logger_path", "", "set logger root path eg: /tmp/kelvins-app")
	flagEnv         = flag.String("env", "", "set exec environment eg: dev,test,prod")
)

// 初始化application--日志部分
func initApplication(application *vars.Application) error {
	flag.Parse()
	rootPath := config.DefaultLoggerRootPath
	if application.LoggerPath != "" {
		rootPath = application.LoggerPath
	}
	if vars.LoggerSetting != nil && vars.LoggerSetting.RootPath != "" {
		rootPath = vars.LoggerSetting.RootPath
	}
	if *flagLoggerPath != "" {
		rootPath = *flagLoggerPath
	}
	application.LoggerPath = rootPath

	loggerLevel := config.DefaultLoggerLevel
	if application.LoggerLevel != "" {
		loggerLevel = application.LoggerLevel
	}
	if vars.LoggerSetting != nil && vars.LoggerSetting.Level != "" {
		loggerLevel = vars.LoggerSetting.Level
	}
	if *flagLoggerLevel != "" {
		loggerLevel = *flagLoggerLevel
	}
	application.LoggerLevel = loggerLevel
	vars.LoggerLevel = loggerLevel

	environment := config.DefaultEnvironmentRelease
	if application.Environment != "" {
		environment = application.Environment
	}
	if vars.ServerSetting != nil && vars.ServerSetting.Environment != "" {
		environment = vars.ServerSetting.Environment
	}
	if *flagEnv != "" {
		environment = *flagEnv
	}
	application.Environment = environment
	vars.Environment = environment

	if vars.ServerSetting == nil {
		vars.ServerSetting = new(setting.ServerSettingS)
	}

	err := log.InitGlobalConfig(rootPath, loggerLevel, application.Name)
	if err != nil {
		return fmt.Errorf("log.InitGlobalConfig: %v", err)
	}

	return nil
}

var appCloseChOnce sync.Once

func appShutdown(application *vars.Application) error {
	if appCloseCh != nil {
		appCloseChOnce.Do(func() {
			close(appCloseCh)
		})
	}
	if application.StopFunc != nil {
		return application.StopFunc()
	}
	return nil
}

func appPrepareForceExit() {
	if !execStopFunc {
		return
	}
	time.AfterFunc(10*time.Second, func() {
		logging.Info("App server Shutdown timeout, force exit")
		os.Exit(1)
	})
}

// 初始化全局配置
func setupCommonVars(application *vars.WEBApplication) error {
	var err error
	vars.ErrorLogger, err = log.GetErrLogger("err")
	if err != nil {
		return err
	}
	varsInternal.ErrorLogger = vars.ErrorLogger

	vars.BusinessLogger, err = log.GetBusinessLogger("business")
	if err != nil {
		return err
	}

	vars.AccessLogger, err = log.GetAccessLogger("access")
	if err != nil {
		return err
	}
	if vars.ServerSetting.PIDFile == "" {
		wd, _ := os.Getwd()
		vars.ServerSetting.PIDFile = fmt.Sprintf("%s/%s.pid", wd, application.Name)
	}
	return nil
}

var execStopFunc bool

var appCloseCh = make(chan struct{})

func startUpControl(pidFile string) (next bool, err error) {
	vars.AppCloseCh = appCloseCh
	next, err = startup.ParseCliCommand(pidFile)
	if next {
		execStopFunc = true
	}
	return
}

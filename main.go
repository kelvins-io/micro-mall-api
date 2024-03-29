package main

import (
	"gitee.com/cristiane/micro-mall-api/app"
	"gitee.com/cristiane/micro-mall-api/startup"
	"gitee.com/cristiane/micro-mall-api/vars"
)

// 服务名
const AppName = "micro-mall-api"

func main() {
	application := &vars.WEBApplication{
		Application: &vars.Application{
			Name:       AppName,
			LoadConfig: startup.LoadConfig,
			SetupVars:  startup.SetupVars,
			StopFunc:   startup.SetStopFunc,
		},
		RegisterHttpRoute: startup.RegisterHttpRoute,
		RegisterTasks:     startup.RegisterTasks,
	}
	app.RunApplication(application)
}

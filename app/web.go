package app

import (
	"flag"
	"fmt"
	"gitee.com/cristiane/micro-mall-api/internal/config"
	"gitee.com/cristiane/micro-mall-api/internal/logging"
	"gitee.com/cristiane/micro-mall-api/vars"
	"net/http"

	"strconv"
)

func RunApplication(application *vars.WEBApplication) {
	defer func() {
		if application.StopFunc != nil {
			application.StopFunc()
		}
	}()
	if application.Name == "" {
		logging.Fatal("Application name can't not be empty")
	}

	flag.Parse()
	application.Type = vars.AppTypeWeb
	vars.App = application
	err := runApp(application)
	if err != nil {
		logging.Fatalf("App.RunListenerApplication err: %v", err)
	}
}

func runApp(webApp *vars.WEBApplication) error {
	// 1. load config
	err := config.LoadDefaultConfig(webApp.Application)
	if err != nil {
		return err
	}
	if webApp.LoadConfig != nil {
		err = webApp.LoadConfig()
		if err != nil {
			return err
		}
	}

	// 2. init application
	err = initApplication(webApp.Application)
	if err != nil {
		return err
	}

	// 3. setup vars
	err = setupWEBVars(webApp)
	if err != nil {
		return err
	}
	if webApp.SetupVars != nil {
		err = webApp.SetupVars()
		if err != nil {
			return fmt.Errorf("App.SetupVars err: %v", err)
		}
	}

	//4.  setup server monitor
	go func() {
		addr := "127.0.0.1:" + strconv.Itoa(webApp.MonitorEndPort)
		logging.Infof("App run monitor server addr: %v", addr)
		err := http.ListenAndServe(addr, webApp.Mux)
		if err != nil {
			logging.Fatalf("App run monitor server err: %v", err)
		}
	}()

	// 5 run task
	//cn := cron.New(cron.WithSeconds())
	//cronTasks := webApp.RegisterTasks()
	//for i := 0;i<len(cronTasks);i++{
	//	if cronTasks[i].TaskFunc != nil {
	//		_,err = cn.AddFunc(cronTasks[i].Cron,cronTasks[i].TaskFunc)
	//		if err != nil {
	//			logging.Fatalf("App run cron task err: %v",err)
	//		}
	//	}
	//}
	//cn.Start()

	// 6. set init service port
	var addr string
	if webApp.EndPort != 0 {
		addr = "127.0.0.1:" + strconv.Itoa(webApp.EndPort)
	} else if vars.ServerSetting.EndPort != 0 {
		addr = "127.0.0.1:" + strconv.Itoa(vars.ServerSetting.EndPort)
	}

	// 7. run http server
	if webApp.RegisterHttpRoute == nil {
		logging.Fatalf("App RegisterHttpRoute nil ??")
	}
	err = webApp.RegisterHttpRoute().Run(addr)

	return err
}

// setupGRPCVars ...
func setupWEBVars(webApp *vars.WEBApplication) error {
	err := setupCommonVars(webApp)
	if err != nil {
		return err
	}

	return nil
}

package app

import (
	"context"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"gitee.com/cristiane/micro-mall-api/internal/config"
	"gitee.com/cristiane/micro-mall-api/internal/logging"
	"gitee.com/cristiane/micro-mall-api/pkg/util/kprocess"
	"gitee.com/cristiane/micro-mall-api/vars"
	"github.com/robfig/cron/v3"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
)

const localAddr = "0.0.0.0:"

func RunApplication(application *vars.WEBApplication) {
	if application == nil || application.Application == nil {
		panic("webApplication is nil or application is nil")
	}
	if application.Name == "" {
		logging.Fatal("Application name can't not be empty")
	}

	application.Type = vars.AppTypeWeb
	vars.App = application
	err := runApp(application)
	if err != nil {
		logging.Infof("App exit over: %v\n", err)
	}
	logging.Info("App exit over")
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

	// startup control
	next, err := startUpControl(vars.ServerSetting.PIDFile)
	if err != nil {
		return err
	}
	if !next {
		return nil
	}

	// 5 run task
	if webApp.RegisterTasks != nil {
		cronTasks := webApp.RegisterTasks()
		if len(cronTasks) != 0 {
			cn := cron.New(cron.WithSeconds())
			for i := 0; i < len(cronTasks); i++ {
				if cronTasks[i].TaskFunc != nil {
					_, err = cn.AddFunc(cronTasks[i].Cron, cronTasks[i].TaskFunc)
					if err != nil {
						logging.Fatalf("App run cron task err: %v", err)
					}
				}
			}
			cn.Start()
			logging.Info("App run cron task")
		}
	}

	// 6. set init service port
	var addr string
	if webApp.EndPort != 0 {
		addr = localAddr + strconv.Itoa(webApp.EndPort)
	} else if vars.ServerSetting.EndPort != 0 {
		addr = localAddr + strconv.Itoa(vars.ServerSetting.EndPort)
	}

	// 7. run http server
	if webApp.RegisterHttpRoute == nil {
		logging.Fatalf("App RegisterHttpRoute nil ??")
	}

	kp := new(kprocess.KProcess)
	network := "tcp"
	if vars.ServerSetting != nil && vars.ServerSetting.Network != "" {
		network = vars.ServerSetting.Network
	}
	ln, err := kp.Listen(network, addr, vars.ServerSetting.PIDFile)
	if err != nil {
		logging.Fatalf("App kprocess listen err: %v", err)
	}
	logging.Infof("server process listen network: %v, addr: %v\n", network, addr)
	ginEngine := webApp.RegisterHttpRoute()
	var handler http.Handler
	if vars.ServerSetting != nil && vars.ServerSetting.SupportH2 {
		logging.Info("server http handler support h2")
		handler = h2c.NewHandler(ginEngine, &http2.Server{IdleTimeout: time.Duration(vars.ServerSetting.IdleTimeout) * time.Second})
	} else {
		handler = ginEngine
	}
	serve := http.Server{
		Handler:      handler,
		ReadTimeout:  time.Duration(vars.ServerSetting.ReadTimeout) * time.Second,
		WriteTimeout: time.Duration(vars.ServerSetting.WriteTimeout) * time.Second,
		IdleTimeout:  time.Duration(vars.ServerSetting.IdleTimeout) * time.Second,
	}
	serverCloe := make(chan struct{})
	go func() {
		defer func() {
			close(serverCloe)
		}()
		err = serve.Serve(ln)
		if err != nil {
			logging.Infof("App run Serve: %v\n", err)
		}
	}()

	select {
	case <-kp.Exit():
	case <-serverCloe:
	}

	appPrepareForceExit()
	err = serve.Shutdown(context.Background())
	if err != nil {
		logging.Infof("App server Shutdown: %v\n", err)
	}
	logging.Info("App server Shutdown ok")

	err = appShutdown(webApp.Application)

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

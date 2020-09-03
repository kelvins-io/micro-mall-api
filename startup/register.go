package startup

import (
	"context"
	"gitee.com/cristiane/micro-mall-api/router"
	"gitee.com/cristiane/micro-mall-api/vars"
	"github.com/gin-gonic/gin"
)

// RegisterHttpRoute 此处注册http接口
func RegisterHttpRoute() *gin.Engine {
	accessInfoLogger := &AccessInfoLogger{}
	accessErrLogger := &AccessErrLogger{}
	ginRouter := router.InitRouter(accessInfoLogger, accessErrLogger)
	return ginRouter
}

type AccessInfoLogger struct{}

func (a *AccessInfoLogger) Write(p []byte) (n int, err error) {
	vars.AccessLogger.Infof(context.Background(), "[gin-info] %s", p)
	return 0, nil
}

type AccessErrLogger struct{}

func (a *AccessErrLogger) Write(p []byte) (n int, err error) {
	vars.AccessLogger.Errorf(context.Background(), "[gin-err] %s", p)
	return 0, nil
}

// 注册定时任务
func RegisterTasks() []vars.CronTask {
	var tasks = make([]vars.CronTask, 0)
	tasks = append(tasks) //TestCronTask(), // 测试定时任务

	return tasks
}

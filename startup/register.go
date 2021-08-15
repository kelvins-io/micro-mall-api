package startup

import (
	"gitee.com/cristiane/micro-mall-api/router"
	"gitee.com/cristiane/micro-mall-api/vars"
	"github.com/gin-gonic/gin"
)

// RegisterHttpRoute 此处注册http接口
func RegisterHttpRoute() *gin.Engine {
	return router.InitRouter()
}

// 注册定时任务
func RegisterTasks() []vars.CronTask {
	var tasks = make([]vars.CronTask, 0)
	tasks = append(tasks) //TestCronTask(), // 测试定时任务

	return tasks
}

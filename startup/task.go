package startup

import (
	"gitee.com/cristiane/micro-mall-api/vars"
	"log"
)

const (
	Cron = "0 */10 * * * *"
)

func TestCronTask() vars.CronTask {
	var task = vars.CronTask{
		Cron: Cron,
		TaskFunc: func() {
			log.Println("需要一些定时任务支持xxx")
		},
	}
	return task
}

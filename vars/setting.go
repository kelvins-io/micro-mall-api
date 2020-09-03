package vars

type CronTask struct {
	Cron string
	TaskFunc func()
}
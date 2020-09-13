package vars

type CronTask struct {
	Cron     string
	TaskFunc func()
}

type EmailConfigSettingS struct {
	User     string `json:"user"`
	Password string `json:"password"`
	Host     string `json:"host"`
	Port     string `json:"port"`
}

type VerifyCodeSettingS struct {
	ExpireMinute int `json:"expire_minute"`
}

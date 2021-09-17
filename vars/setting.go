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
	//number of requests to send verification code in a period of time
	SendPeriodLimitCount int `json:"send_period_limit_count"`
	//limit the timeout period of request to send verification code within a period of time
	SendPeriodLimitExpireSecond int64 `json:"send_period_limit_expire_second"`
	//request to send verification code interval time
	SendIntervalExpireSecond int64 `json:"send_interval_expire_second"`
}

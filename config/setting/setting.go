package setting

type ServerSettingS struct {
	Network      string
	EndPort      int
	ReadTimeout  int
	WriteTimeout int
	IdleTimeout  int
	Mode         string
	PIDFile      string
}

type JwtSettingS struct {
	Secret            string
	TokenExpireSecond int
}

type LoggerSettingS struct {
	RootPath string
	Level    string
}

// MysqlSettingS defines for connecting mysql.
type MysqlSettingS struct {
	Host              string
	UserName          string
	Password          string
	DBName            string
	Charset           string
	PoolNum           int
	Loc               string
	MaxIdleConns      int
	ConnMaxLifeSecond int
	MultiStatements   bool
	ParseTime         bool
}

// RedisSettingS defines for connecting redis.
type RedisSettingS struct {
	Host     string
	Password string
	DB       int
	PoolNum  int
}

// QueueRedisSettingS defines for redis queue.
type QueueRedisSettingS struct {
	Broker          string
	DefaultQueue    string
	ResultBackend   string
	ResultsExpireIn int
}

type QueueAMQPSettingS struct {
	Broker           string
	DefaultQueue     string
	ResultBackend    string
	ResultsExpireIn  int
	Exchange         string
	ExchangeType     string
	BindingKey       string
	PrefetchCount    int
	TaskRetryCount   int
	TaskRetryTimeout int
}

// QueueAliAMQPSettingS defines for aliyun AMQP queue
type QueueAliAMQPSettingS struct {
	AccessKey       string
	SecretKey       string
	AliUid          int
	EndPoint        string
	VHost           string
	DefaultQueue    string
	ResultBackend   string
	ResultsExpireIn int
	Exchange        string
	ExchangeType    string
	BindingKey      string
	PrefetchCount   int
}

//g2cache config
type G2CacheSettingS struct {
	CacheDebug           bool
	CacheMonitor         bool
	OutCachePubSub       bool
	CacheMonitorSecond   int
	EntryLazyFactor      int
	GPoolWorkerNum       int
	GPoolJobQueueChanLen int
	FreeCacheSize        int // 100MB
	PubSubRedisChannel   string
	RedisConfDSN         string
	RedisConfDB          int
	RedisConfPwd         string
	RedisConfMaxConn     int
}

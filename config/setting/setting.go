package setting

type ServerSettingS struct {
	Network        string
	EndPort        int
	MonitorEndPort int
	ReadTimeout    int
	WriteTimeout   int
	Mode           string
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

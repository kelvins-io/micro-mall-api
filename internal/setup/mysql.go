package setup

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"gitee.com/cristiane/micro-mall-api/config/setting"
	"gitee.com/cristiane/micro-mall-api/internal/config"
	"gitee.com/kelvins-io/common/log"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"io"
	"net/url"
	"os"
	"strconv"
	"sync"
	"time"
	"xorm.io/xorm"
	xormLog "xorm.io/xorm/log"
)

// NewMySQL returns *gorm.DB instance.
func NewMySQLWithGORM(mysqlSetting *setting.MysqlSettingS) (*gorm.DB, error) {
	if mysqlSetting == nil {
		return nil, fmt.Errorf("mysqlSetting is nil")
	}
	if mysqlSetting.UserName == "" {
		return nil, fmt.Errorf("lack of mysqlSetting.UserName")
	}
	if mysqlSetting.Password == "" {
		return nil, fmt.Errorf("lack of mysqlSetting.Password")
	}
	if mysqlSetting.Host == "" {
		return nil, fmt.Errorf("lack of mysqlSetting.Host")
	}
	if mysqlSetting.DBName == "" {
		return nil, fmt.Errorf("lack of mysqlSetting.DBName")
	}
	if mysqlSetting.Charset == "" {
		return nil, fmt.Errorf("lack of mysqlSetting.Charset")
	}

	var buf bytes.Buffer
	buf.WriteString(mysqlSetting.UserName)
	buf.WriteString(":")
	buf.WriteString(mysqlSetting.Password)
	buf.WriteString("@tcp(")
	buf.WriteString(mysqlSetting.Host)
	buf.WriteString(")/")
	buf.WriteString(mysqlSetting.DBName)
	buf.WriteString("?charset=")
	buf.WriteString(mysqlSetting.Charset)
	buf.WriteString("&parseTime=" + strconv.FormatBool(mysqlSetting.ParseTime))
	buf.WriteString("&multiStatements=" + strconv.FormatBool(mysqlSetting.MultiStatements))
	if mysqlSetting.Loc == "" {
		buf.WriteString("&loc=Local")
	} else {
		buf.WriteString("&loc=" + url.QueryEscape(mysqlSetting.Loc))
	}

	db, err := gorm.Open("mysql", buf.String())
	if err != nil {
		return nil, err
	}

	logger, _ := log.GetCustomLogger("db-log", "gorm")
	gormLogger := &gormLogger{
		logger: logger,
	}
	if mysqlSetting.Environment == config.DefaultEnvironmentDev {
		db.LogMode(true)
		gormLogger.out = os.Stdout
	}
	db.SetLogger(gormLogger)

	db.DB().SetConnMaxLifetime(3600 * time.Second)
	if mysqlSetting.ConnMaxLifeSecond > 0 {
		db.DB().SetConnMaxLifetime(time.Duration(mysqlSetting.ConnMaxLifeSecond) * time.Second)
	}

	db.DB().SetMaxIdleConns(10)
	if mysqlSetting.MaxIdleConns > 0 {
		db.DB().SetMaxIdleConns(mysqlSetting.MaxIdleConns)
	}

	db.DB().SetMaxOpenConns(10)
	if mysqlSetting.MaxIdleConns > 0 {
		db.DB().SetMaxOpenConns(mysqlSetting.MaxIdleConns)
	}

	return db, nil
}

type gormLogger struct {
	logger log.LoggerContextIface
	out    io.Writer
}

var logBufPool = sync.Pool{
	New: func() interface{} {
		b := make([]byte, 32*1024)
		return &b
	},
}

var gormLoggerCtx = context.Background()

func (l *gormLogger) Print(vv ...interface{}) {
	l.logger.Info(gormLoggerCtx, vv)
	if l.out != nil {
		buf := logBufPool.Get().(*[]byte)
		defer logBufPool.Put(buf)
		w := bytes.NewBuffer(*buf)
		for _, v := range vv {
			vLog, _ := json.Marshal(v)
			w.Write(vLog)
		}
		_, _ = l.out.Write(w.Bytes())
	}
}

var xormLogLevel = map[string]xormLog.LogLevel{
	"debug": xormLog.LOG_DEBUG,
	"info":  xormLog.LOG_INFO,
	"warn":  xormLog.LOG_WARNING,
	"error": xormLog.LOG_ERR,
}

// NewMySQL returns *xorm.DB instance.
func NewMySQLWithXORM(mysqlSetting *setting.MysqlSettingS) (xorm.EngineInterface, error) {
	if mysqlSetting == nil {
		return nil, fmt.Errorf("mysqlSetting is nil")
	}
	if mysqlSetting.UserName == "" {
		return nil, fmt.Errorf("lack of mysqlSetting.UserName")
	}
	if mysqlSetting.Password == "" {
		return nil, fmt.Errorf("lack of mysqlSetting.Password")
	}
	if mysqlSetting.Host == "" {
		return nil, fmt.Errorf("lack of mysqlSetting.Host")
	}
	if mysqlSetting.DBName == "" {
		return nil, fmt.Errorf("lack of mysqlSetting.DBName")
	}
	if mysqlSetting.Charset == "" {
		return nil, fmt.Errorf("lack of mysqlSetting.Charset")
	}

	var buf bytes.Buffer
	buf.WriteString(mysqlSetting.UserName)
	buf.WriteString(":")
	buf.WriteString(mysqlSetting.Password)
	buf.WriteString("@tcp(")
	buf.WriteString(mysqlSetting.Host)
	buf.WriteString(")/")
	buf.WriteString(mysqlSetting.DBName)
	buf.WriteString("?charset=")
	buf.WriteString(mysqlSetting.Charset)
	buf.WriteString("&parseTime=" + strconv.FormatBool(mysqlSetting.ParseTime))
	buf.WriteString("&multiStatements=" + strconv.FormatBool(mysqlSetting.MultiStatements))
	if mysqlSetting.Loc == "" {
		buf.WriteString("&loc=Local")
	} else {
		buf.WriteString("&loc=" + url.QueryEscape(mysqlSetting.Loc))
	}

	engine, err := xorm.NewEngine("mysql", buf.String())
	if err != nil {
		return nil, err
	}

	logger, _ := log.GetCustomLogger("db-log", "xorm")
	xormlogger := &xormLogger{
		logger: logger,
	}
	engine.SetLogLevel(xormLogLevel[mysqlSetting.LoggerLevel])
	var writer io.Writer
	writer = xormlogger
	if mysqlSetting.Environment == config.DefaultEnvironmentDev {
		writer = io.MultiWriter(writer, os.Stdout)
	}
	engine.SetLogger(xormLog.NewSimpleLogger(writer))
	if mysqlSetting.Environment == config.DefaultEnvironmentDev {
		engine.ShowSQL(true)
	}

	engine.SetConnMaxLifetime(3600 * time.Second)
	if mysqlSetting.ConnMaxLifeSecond > 0 {
		engine.SetConnMaxLifetime(time.Duration(mysqlSetting.ConnMaxLifeSecond) * time.Second)
	}

	engine.SetMaxIdleConns(10)
	if mysqlSetting.MaxIdleConns > 0 {
		engine.SetMaxIdleConns(mysqlSetting.MaxIdleConns)
	}

	engine.SetMaxOpenConns(10)
	if mysqlSetting.MaxIdleConns > 0 {
		engine.SetMaxOpenConns(mysqlSetting.MaxIdleConns)
	}

	return engine, nil
}

var xormLoggerCtx = context.Background()

type xormLogger struct {
	logger log.LoggerContextIface
}

func (l *xormLogger) Write(p []byte) (n int, err error) {
	l.logger.Info(xormLoggerCtx, string(p))
	return 0, nil
}

// SetGORMCreateCallback is set create callback
func SetGORMCreateCallback(db *gorm.DB, callback func(scope *gorm.Scope)) {
	db.Callback().Create().Replace("gorm:update_time_stamp", callback)
}

// SetGORMUpdateCallback is set update callback
func SetGORMUpdateCallback(db *gorm.DB, callback func(scope *gorm.Scope)) {
	db.Callback().Update().Replace("gorm:update_time_stamp", callback)
}

// SetGORMDeleteCallback is set delete callback
func SetGORMDeleteCallback(db *gorm.DB, callback func(scope *gorm.Scope)) {
	db.Callback().Delete().Replace("gorm:delete", callback)
}

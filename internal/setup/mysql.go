package setup

import (
	"bytes"
	"fmt"
	"gitee.com/cristiane/micro-mall-api/config/setting"
	"github.com/jinzhu/gorm"
	"net/url"
	"strconv"
	"time"

	"gitee.com/kelvins-io/common/env"
	_ "github.com/go-sql-driver/mysql"
	"xorm.io/xorm"
	xormLog "xorm.io/xorm/log"
)

// NewMySQL returns *xorm.DB instance.
func NewMySQLXORMEngine(mysqlSetting *setting.MysqlSettingS) (*xorm.Engine, error) {
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
	if mysqlSetting.PoolNum == 0 {
		return nil, fmt.Errorf("lack of mysqlSetting.PoolNum")
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
	if env.IsDevMode() {
		engine.SetLogLevel(xormLog.LOG_DEBUG)
		engine.ShowSQL(true)
	}

	engine.SetConnMaxLifetime(time.Duration(30) * time.Second)
	if mysqlSetting.ConnMaxLifeSecond > 0 {
		engine.SetConnMaxLifetime(time.Duration(mysqlSetting.ConnMaxLifeSecond) * time.Second)
	}

	engine.SetMaxIdleConns(10)
	if mysqlSetting.MaxIdleConns > 0 {
		engine.SetMaxIdleConns(mysqlSetting.MaxIdleConns)
	}

	engine.SetMaxOpenConns(10)
	if mysqlSetting.PoolNum > 0 {
		engine.SetMaxOpenConns(mysqlSetting.PoolNum)
	}

	return engine, nil
}

// NewMySQL returns *gorm.DB instance.
func NewMySQLGORMEngine(mysqlSetting *setting.MysqlSettingS) (*gorm.DB, error) {
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
	if mysqlSetting.PoolNum == 0 {
		return nil, fmt.Errorf("lack of mysqlSetting.PoolNum")
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
	if env.IsDevMode() {
		db.LogMode(true)
	}

	db.DB().SetConnMaxLifetime(30 * time.Second)
	if mysqlSetting.ConnMaxLifeSecond > 0 {
		db.DB().SetConnMaxLifetime(time.Duration(mysqlSetting.ConnMaxLifeSecond) * time.Second)
	}

	db.DB().SetMaxIdleConns(10)
	if mysqlSetting.MaxIdleConns > 0 {
		db.DB().SetMaxIdleConns(mysqlSetting.MaxIdleConns)
	}

	db.DB().SetMaxOpenConns(10)
	if mysqlSetting.PoolNum > 0 {
		db.DB().SetMaxOpenConns(mysqlSetting.PoolNum)
	}

	return db, nil
}

// SetCreateCallback is set create callback
//func SetCreateCallback(db *xorm.Engine, callback func(scope *x)) {
//	db.Callback().Create().Replace("gorm:update_time_stamp", callback)
//}
//
//// SetUpdateCallback is set update callback
//func SetUpdateCallback(db *gorm.DB, callback func(scope *gorm.Scope)) {
//	db.Callback().Update().Replace("gorm:update_time_stamp", callback)
//}
//
//// SetDeleteCallback is set delete callback
//func SetDeleteCallback(db *gorm.DB, callback func(scope *gorm.Scope)) {
//	db.Callback().Delete().Replace("gorm:delete", callback)
//}

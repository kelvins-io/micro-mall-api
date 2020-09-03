package startup

import (
	"gitee.com/cristiane/go-common/log"
	"gitee.com/cristiane/micro-mall-api/internal/setup"
	"gitee.com/cristiane/micro-mall-api/vars"
)

// SetupVars 加载变量
func SetupVars() error {
	var err error

	vars.ErrorLogger, err = log.GetErrLogger("err")
	if err != nil {
		return err
	}

	vars.BusinessLogger, err = log.GetBusinessLogger("business")
	if err != nil {
		return err
	}

	vars.AccessLogger, err = log.GetAccessLogger("access")
	if err != nil {
		return err
	}

	if vars.MysqlSettingXXXDB != nil && vars.MysqlSettingXXXDB.Host != "" {
		vars.DBEngineXORM, err = setup.NewMySQLXORMEngine(vars.MysqlSettingXXXDB)
		if err != nil {
			return err
		}
		vars.DBEngineGORM, err = setup.NewMySQLGORMEngine(vars.MysqlSettingXXXDB)
		if err != nil {
			return err
		}
	}
	if vars.RedisSetting != nil && vars.RedisSetting.Host != "" {
		vars.RedisPool, err = setup.NewRedis(vars.RedisSetting)
		if err != nil {
			return err
		}
	}

	return nil
}

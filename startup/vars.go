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

	if vars.MysqlSettingMicroMall != nil && vars.MysqlSettingMicroMall.Host != "" {
		vars.DBEngineXORM, err = setup.NewMySQLXORMEngine(vars.MysqlSettingMicroMall)
		if err != nil {
			return err
		}
		vars.DBEngineGORM, err = setup.NewMySQLGORMEngine(vars.MysqlSettingMicroMall)
		if err != nil {
			return err
		}
	}
	if vars.RedisSettingMicroMall != nil && vars.RedisSettingMicroMall.Host != "" {
		vars.RedisPoolMicroMall, err = setup.NewRedis(vars.RedisSettingMicroMall)
		if err != nil {
			return err
		}
	}
	if vars.QueueAMQPSettingUserRegisterNotice != nil && vars.QueueAMQPSettingUserRegisterNotice.Broker != "" {
		vars.QueueServerUserRegisterNotice = setup.SetUpAMQPQueue(vars.QueueAMQPSettingUserRegisterNotice, nil)
	}

	if vars.QueueAMQPSettingUserStateNotice != nil && vars.QueueAMQPSettingUserStateNotice.Broker != "" {
		vars.QueueServerUserStateNotice = setup.SetUpAMQPQueue(vars.QueueAMQPSettingUserStateNotice, nil)
	}

	return nil
}

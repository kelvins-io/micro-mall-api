package startup

import (
	"gitee.com/cristiane/micro-mall-api/internal/setup"
	"gitee.com/cristiane/micro-mall-api/pkg/util/goroutine"
	"gitee.com/cristiane/micro-mall-api/vars"
	"gitee.com/kelvins-io/common/log"
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
	if vars.G2CacheSetting != nil && vars.G2CacheSetting.RedisConfDSN != "" {
		vars.G2CacheEngine, err = setup.NewG2Cache(vars.G2CacheSetting,nil,nil)
		if err != nil {
			return err
		}
	}
	vars.GPool = goroutine.NewPool(20, 1000)

	return nil
}

func SetStopFunc() (err error) {
	if vars.GPool != nil {
		vars.GPool.Release()
		vars.GPool.WaitAll()
	}
	if vars.G2CacheEngine != nil {
		vars.G2CacheEngine.Close()
	}
	if vars.RedisPoolMicroMall != nil {
		err = vars.RedisPoolMicroMall.Close()
		if err != nil {
			return err
		}
	}

	return nil
}

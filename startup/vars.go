package startup

import (
	"gitee.com/cristiane/micro-mall-api/internal/setup"
	"gitee.com/cristiane/micro-mall-api/pkg/util/goroutine"
	"gitee.com/cristiane/micro-mall-api/vars"
)

// SetupVars 加载变量
func SetupVars() error {
	var err error
	if vars.G2CacheSetting != nil && vars.G2CacheSetting.RedisConfDSN != "" {
		vars.G2CacheEngine, err = setup.NewG2Cache(vars.G2CacheSetting, nil, nil)
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

	return nil
}

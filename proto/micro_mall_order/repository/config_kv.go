package repository

import (
	"gitee.com/cristiane/micro-mall-order/model/mysql"
	"gitee.com/kelvins-io/kelvins"
)

func GetConfigKv(key string) (*mysql.ConfigKvStore, error) {
	var model mysql.ConfigKvStore
	var err error
	_, err = kelvins.XORM_DBEngine.Table(mysql.TableConfigKv).
		Select("config_value").
		Where("status = 1 AND is_delete = 0").
		Where("config_key = ?", key).Get(&model)
	return &model, err
}

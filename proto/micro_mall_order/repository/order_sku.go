package repository

import (
	"gitee.com/cristiane/micro-mall-order/model/mysql"
	"gitee.com/kelvins-io/kelvins"
	"xorm.io/xorm"
)

func CreateOrderSku(tx *xorm.Session, models []mysql.OrderSku) (err error) {
	_, err = tx.Table(mysql.TableOrderSku).Insert(models)
	return
}

func GetOrderSkuListByOrderCode(orderCode []string) ([]mysql.OrderSku, error) {
	var result = make([]mysql.OrderSku, 0)
	var err error
	err = kelvins.XORM_DBEngine.Table(mysql.TableOrderSku).In("order_code", orderCode).Find(&result)
	return result, err
}

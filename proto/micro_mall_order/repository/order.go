package repository

import (
	"gitee.com/cristiane/micro-mall-order/model/mysql"
	"gitee.com/kelvins-io/kelvins"
	"xorm.io/xorm"
)

func CreateOrder(tx *xorm.Session, models []mysql.Order) (err error) {
	_, err = tx.Table(mysql.TableOrder).Insert(models)
	return
}

func GetOrderByOrderCode(orderCode string) (*mysql.Order, error) {
	var model mysql.Order
	var err error
	_, err = kelvins.XORM_DBEngine.Table(mysql.TableOrder).Where("order_code = ?", orderCode).Get(&model)
	return &model, err
}

func GetOrderListByTxCode(txCode string) ([]mysql.Order, error) {
	var result = make([]mysql.Order, 0)
	var err error
	err = kelvins.XORM_DBEngine.Table(mysql.TableOrder).Where("tx_code = ?", txCode).Find(&result)
	return result, err
}

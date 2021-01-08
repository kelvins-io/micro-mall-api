package server

import (
	"context"
	"gitee.com/cristiane/micro-mall-order/pkg/code"
	"gitee.com/cristiane/micro-mall-order/proto/micro_mall_order_proto/order_business"
	"gitee.com/cristiane/micro-mall-order/service"
	"gitee.com/kelvins-io/common/errcode"
)

type OrderServer struct {
}

func NewOrderServer() order_business.OrderBusinessServiceServer {
	return new(OrderServer)
}

func (o *OrderServer) CreateOrder(ctx context.Context, req *order_business.CreateOrderRequest) (*order_business.CreateOrderResponse, error) {
	var result = order_business.CreateOrderResponse{
		Common: &order_business.CommonResponse{
			Code: order_business.RetCode_SUCCESS,
		},
	}
	rsp, retCode := service.CreateOrder(ctx, req)
	if retCode != code.Success {
		if retCode == code.UserNotExist {
			result.Common.Code = order_business.RetCode_USER_NOT_EXIST
		} else if retCode == code.SkuAmountNotEnough {
			result.Common.Code = order_business.RetCode_SKU_AMOUNT_NOT_ENOUGH
		} else if retCode == code.TransactionFailed {
			result.Common.Code = order_business.RetCode_TRANSACTION_FAILED
		} else {
			result.Common.Code = order_business.RetCode_ERROR
		}
		result.Common.Msg = errcode.GetErrMsg(retCode)
		return &result, nil
	}
	result.TxCode = rsp.TxCode
	return &result, nil
}

func (o *OrderServer) GetOrderDetail(ctx context.Context, req *order_business.GetOrderDetailRequest) (*order_business.GetOrderDetailResponse, error) {
	var result order_business.GetOrderDetailResponse
	result.Common = &order_business.CommonResponse{
		Code: order_business.RetCode_SUCCESS,
	}
	rsp, retCode := service.GetOrderDetail(ctx, req)
	if retCode != code.Success {
		result.Common.Code = order_business.RetCode_ERROR
		result.Common.Msg = errcode.GetErrMsg(code.ErrorServer)
		return &result, nil
	}
	result.Account = rsp.UserCode
	result.CoinType = order_business.CoinType(rsp.CoinType)
	result.List = make([]*order_business.ShopOrderDetail, len(rsp.List))
	for i := 0; i < len(rsp.List); i++ {
		shopOrderDe := &order_business.ShopOrderDetail{
			OrderCode:   rsp.List[i].OrderCode,
			Merchant:    rsp.List[i].ShopCode,
			TimeExpire:  rsp.List[i].TimeExpire,
			NotifyUrl:   rsp.List[i].NotifyUrl,
			Description: rsp.List[i].Description,
			Detail: &order_business.TradeGoodsDetail{
				Money: rsp.List[i].Amount,
			},
		}
		result.List[i] = shopOrderDe
	}
	return &result, nil
}

func (o *OrderServer) GetOrderSku(ctx context.Context, req *order_business.GetOrderSkuRequest) (*order_business.GetOrderSkuResponse, error) {
	result := &order_business.GetOrderSkuResponse{
		Common: &order_business.CommonResponse{
			Code: order_business.RetCode_SUCCESS,
		},
	}
	orderSku, retCode := service.GetOrderSku(ctx, req)
	if retCode != code.Success {
		result.Common.Code = order_business.RetCode_ERROR
		result.Common.Msg = errcode.GetErrMsg(retCode)
		return result, nil
	}
	result.OrderList = make([]*order_business.OrderSku, len(orderSku.SkuList))
	for i := 0; i < len(orderSku.SkuList); i++ {
		row := orderSku.SkuList[i]
		orderSku := &order_business.OrderSku{
			OrderCode: row.OrderCode,
			Goods:     nil,
		}
		goods := make([]*order_business.OrderGoods, len(row.SkuList))
		for j := 0; j < len(row.SkuList); j++ {
			orderGoods := &order_business.OrderGoods{
				SkuCode: row.SkuList[j].SkuCode,
				Price:   row.SkuList[j].Price,
				Amount:  int64(row.SkuList[j].Amount),
				Name:    row.SkuList[j].Name,
			}
			goods[j] = orderGoods
		}
		orderSku.Goods = goods
		result.OrderList[i] = orderSku
	}
	return result, nil
}

package service

import (
	"context"
	"gitee.com/cristiane/micro-mall-api/model/args"
	"gitee.com/cristiane/micro-mall-api/pkg/code"
	"gitee.com/cristiane/micro-mall-api/pkg/util"
	"gitee.com/cristiane/micro-mall-api/proto/micro_mall_order_proto/order_business"
	"gitee.com/cristiane/micro-mall-api/proto/micro_mall_pay_proto/pay_business"
	"gitee.com/cristiane/micro-mall-api/vars"
	"time"
)

func GenOrderCode(ctx context.Context, uid int64) (string, int) {
	result := ""
	serverName := args.RpcServiceMicroMallOrder
	conn, err := util.GetGrpcClient(serverName)
	if err != nil {
		vars.ErrorLogger.Errorf(ctx, "GetGrpcClient %v,err: %v", serverName, err)
		return result, code.ERROR
	}
	defer conn.Close()
	client := order_business.NewOrderBusinessServiceClient(conn)
	r := order_business.GenOrderTxCodeRequest{Uid: uid}
	rsp, err := client.GenOrderTxCode(ctx, &r)
	if err != nil || rsp.Common.Code != order_business.RetCode_SUCCESS {
		return "", code.ERROR
	}
	if rsp.OrderTxCode == "" {
		return "", code.ERROR
	}
	result = rsp.OrderTxCode

	return result, code.SUCCESS
}

func CreateTradeOrder(ctx context.Context, req *args.CreateTradeOrderArgs) (*args.CreateTradeOrderRsp, int) {
	var result args.CreateTradeOrderRsp
	serverName := args.RpcServiceMicroMallOrder
	conn, err := util.GetGrpcClient(serverName)
	if err != nil {
		vars.ErrorLogger.Errorf(ctx, "GetGrpcClient %v,err: %v", serverName, err)
		return &result, code.ERROR
	}
	defer conn.Close()
	client := order_business.NewOrderBusinessServiceClient(conn)
	r := order_business.CreateOrderRequest{
		Uid:           req.Uid,
		Time:          util.ParseTimeOfStr(time.Now().Unix()),
		Description:   req.Description,
		PayerClientIp: req.ClientIp,
		DeviceId:      req.DeviceId,
		OrderTxCode:   req.OrderTxCode,
		Detail: &order_business.OrderDetail{
			ShopDetail: nil,
		},
		DeliveryInfo: &order_business.OrderDeliveryInfo{UserDeliveryId: req.UserDeliveryId},
	}
	r.Detail.ShopDetail = make([]*order_business.OrderShopDetail, len(req.Detail))
	for i := 0; i < len(req.Detail); i++ {
		shopDetail := &order_business.OrderShopDetail{
			ShopId:   req.Detail[i].ShopId,
			CoinType: order_business.CoinType(req.Detail[i].CoinType),
			Goods:    nil,
			SceneInfo: &order_business.OrderSceneInfo{
				StoreInfo: &order_business.StoreInfo{
					Id:       req.Detail[i].ShopId,
					Name:     req.Detail[i].SceneInfo.StoreInfo.Name,
					AreaCode: req.Detail[i].SceneInfo.StoreInfo.AreaCode,
					Address:  req.Detail[i].SceneInfo.StoreInfo.Address,
				},
			},
		}
		goods := req.Detail[i].Goods
		orderGoods := make([]*order_business.OrderGoods, len(goods))
		for j := 0; j < len(goods); j++ {
			orderG := &order_business.OrderGoods{
				SkuCode: goods[j].SkuCode,
				Price:   goods[j].Price,
				Amount:  goods[j].Amount,
				Name:    goods[j].Name,
				Version: goods[j].Version,
			}
			orderGoods[j] = orderG
		}
		shopDetail.Goods = orderGoods
		r.Detail.ShopDetail[i] = shopDetail
	}
	rsp, err := client.CreateOrder(ctx, &r)
	if err != nil {
		vars.ErrorLogger.Errorf(ctx, "CreateOrder %v,err: %v, req: %+v", serverName, err, r)
		return &result, code.ERROR
	}
	if rsp.Common.Code == order_business.RetCode_ERROR {
		vars.ErrorLogger.Errorf(ctx, "CreateOrder %v,err: %v, rsp: %+v", serverName, err, rsp)
		return &result, code.ERROR
	}
	result.TxCode = rsp.TxCode
	switch rsp.Common.Code {
	case order_business.RetCode_SKU_PRICE_VERSION_NOT_EXIST:
		return &result, code.SkuPriceVersionNotExist
	case order_business.RetCode_ORDER_DELIVERY_NOT_EXIST:
		return &result, code.UserDeliveryInfoNotExist
	case order_business.RetCode_ORDER_TX_CODE_EMPTY:
		return &result, code.TradeOrderTxCodeEmpty
	case order_business.RetCode_ORDER_EXIST: // 如果订单已存在，显示创建成功，防止客户端反复重试
		return &result, code.TradeOrderExist
	case order_business.RetCode_USER_NOT_EXIST:
		return &result, code.ErrorUserNotExist
	case order_business.RetCode_USER_EXIST:
		return &result, code.ErrorUserExist
	case order_business.RetCode_SHOP_EXIST:
		return &result, code.ErrorShopBusinessExist
	case order_business.RetCode_SHOP_NOT_EXIST:
		return &result, code.ErrorShopBusinessNotExist
	case order_business.RetCode_SKU_AMOUNT_NOT_ENOUGH:
		return &result, code.ErrorSkuAmountNotEnough
	case order_business.RetCode_TRANSACTION_FAILED:
		return &result, code.TransactionFailed
	}

	return &result, code.SUCCESS
}

func OrderTrade(ctx context.Context, req *args.OrderTradeArgs) (result *args.OrderTradeRsp, retCode int) {
	result = &args.OrderTradeRsp{}
	retCode = code.SUCCESS
	// 根据交易号获取订单详情
	serverName := args.RpcServiceMicroMallOrder
	conn, err := util.GetGrpcClient(serverName)
	if err != nil {
		vars.ErrorLogger.Errorf(ctx, "GetGrpcClient %v,err: %v", serverName, err)
		retCode = code.ERROR
		return
	}
	defer conn.Close()
	client := order_business.NewOrderBusinessServiceClient(conn)
	r := order_business.GetOrderDetailRequest{TxCode: req.TxCode}
	rsp, err := client.GetOrderDetail(ctx, &r)
	if err != nil {
		vars.ErrorLogger.Errorf(ctx, "GetOrderDetail %v,err: %v, req: %+v", serverName, err, r)
		retCode = code.ERROR
		return
	}
	if rsp == nil || rsp.Common == nil || rsp.Common.Code == order_business.RetCode_ERROR {
		vars.ErrorLogger.Errorf(ctx, "GetOrderDetail %v,err: %v, rsp: %+v", serverName, err, rsp)
		retCode = code.ERROR
		return
	}
	if rsp.Common.Code != order_business.RetCode_SUCCESS {
		switch rsp.Common.Code {
		case order_business.RetCode_ORDER_STATE_INVALID:
			retCode = code.OrderStateInvalid
			return
		case order_business.RetCode_ORDER_STATE_LOCKED:
			retCode = code.OrderStateLock
			return
		case order_business.RetCode_ORDER_PAY_COMPLETED:
			retCode = code.OrderPayCompleted
			return
		case order_business.RetCode_ORDER_EXPIRE:
			retCode = code.OrderExpire
			return
		default:
			retCode = code.ERROR
			return
		}
	}
	if len(rsp.List) == 0 {
		retCode = code.TxCodeNotExist
		return
	}
	// 发起支付流程
	serverName = args.RpcServiceMicroMallPay
	conn, err = util.GetGrpcClient(serverName)
	if err != nil {
		vars.ErrorLogger.Errorf(ctx, "GetGrpcClient %v,err: %v", serverName, err)
		retCode = code.ERROR
		return
	}
	defer conn.Close()
	payClient := pay_business.NewPayBusinessServiceClient(conn)
	payR := pay_business.TradePayRequest{
		Account:   rsp.Account,
		CoinType:  pay_business.CoinType(rsp.CoinType),
		EntryList: nil,
		OpUid:     req.OpUid,
		OpIp:      req.OpIp,
		OutTxCode: req.TxCode,
	}
	payR.EntryList = make([]*pay_business.TradePayEntry, len(rsp.List))
	for i := 0; i < len(rsp.List); i++ {
		tradeEntry := &pay_business.TradePayEntry{
			OutTradeNo:  rsp.List[i].OrderCode,
			TimeExpire:  rsp.List[i].TimeExpire,
			NotifyUrl:   rsp.List[i].NotifyUrl,
			Description: rsp.List[i].Description,
			Merchant:    rsp.List[i].Merchant,
			Attach:      rsp.List[i].Description,
			Detail: &pay_business.TradeGoodsDetail{
				Amount:    rsp.List[i].Detail.Money,
				Reduction: "0",
			},
		}
		payR.EntryList[i] = tradeEntry
	}
	payRsp, err := payClient.TradePay(ctx, &payR)
	if err != nil {
		vars.ErrorLogger.Errorf(ctx, "TradePay %v,err: %v, req: %+v", serverName, err, payR)
		retCode = code.ERROR
		return
	}
	if payRsp == nil || payRsp.Common == nil || payRsp.Common.Code == pay_business.RetCode_ERROR {
		vars.ErrorLogger.Errorf(ctx, "TradePay %v,err: %v, rsp: %+v", serverName, err, payRsp)
		retCode = code.ERROR
		return
	}
	switch payRsp.Common.Code {
	case pay_business.RetCode_TRADE_ORDER_NOT_MATCH_USER:
		retCode = code.TradeOrderNotMatchUser
		return
	case pay_business.RetCode_USER_NOT_EXIST:
		retCode = code.ErrorUserNotExist
		return
	case pay_business.RetCode_USER_ACCOUNT_NOT_EXIST:
		retCode = code.UserAccountNotExist
		return
	case pay_business.RetCode_USER_BALANCE_NOT_ENOUGH:
		retCode = code.UserBalanceNotEnough
		return
	case pay_business.RetCode_USER_ACCOUNT_STATE_LOCK:
		retCode = code.UserAccountStateLock
		return
	case pay_business.RetCode_MERCHANT_ACCOUNT_NOT_EXIST:
		retCode = code.MerchantAccountNotExist
		return
	case pay_business.RetCode_MERCHANT_ACCOUNT_STATE_LOCK:
		retCode = code.MerchantAccountStateLock
		return
	case pay_business.RetCode_DECIMAL_PARSE_ERR:
		retCode = code.DecimalParseErr
		return
	case pay_business.RetCode_TRANSACTION_FAILED:
		retCode = code.TransactionFailed
		return
	case pay_business.RetCode_TRADE_PAY_RUN:
		retCode = code.TradePayRun
		return
	case pay_business.RetCode_TRADE_PAY_EXPIRE:
		retCode = code.TradePayExpire
		return
	case pay_business.RetCode_TRADE_PAY_SUCCESS:
		retCode = code.TradePaySuccess
		return
	}
	result.IsSuccess = true

	return
}

func GetOrderReport(ctx context.Context, req *args.GetOrderReportArgs) (*args.GetOrderReportRsp, int) {
	return getOrderReport(ctx, req)
}

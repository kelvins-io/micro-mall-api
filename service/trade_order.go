package service

import (
	"context"
	"fmt"
	"gitee.com/cristiane/micro-mall-api/model/args"
	"gitee.com/cristiane/micro-mall-api/pkg/code"
	"gitee.com/cristiane/micro-mall-api/pkg/util"
	"gitee.com/cristiane/micro-mall-api/pkg/util/goroutine"
	"gitee.com/cristiane/micro-mall-api/proto/micro_mall_order_proto/order_business"
	"gitee.com/cristiane/micro-mall-api/proto/micro_mall_pay_proto/pay_business"
	"gitee.com/cristiane/micro-mall-api/proto/micro_mall_shop_proto/shop_business"
	"gitee.com/cristiane/micro-mall-api/proto/micro_mall_users_proto/users"
	"gitee.com/cristiane/micro-mall-api/vars"
	"gitee.com/kelvins-io/common/json"
	"gitee.com/kelvins-io/kelvins"
	"golang.org/x/sync/errgroup"
	"time"
)

func GenOrderCode(ctx context.Context, uid int64) (string, int) {
	result := ""
	serverName := args.RpcServiceMicroMallOrder
	conn, err := util.GetGrpcClient(ctx, serverName)
	if err != nil {
		vars.ErrorLogger.Errorf(ctx, "GetGrpcClient %q err: %v", serverName, err)
		return result, code.ERROR
	}
	//defer conn.Close()
	client := order_business.NewOrderBusinessServiceClient(conn)
	r := order_business.GenOrderTxCodeRequest{Uid: uid}
	rsp, err := client.GenOrderTxCode(ctx, &r)
	if err != nil {
		vars.ErrorLogger.Errorf(ctx, "GenOrderTxCode err: %v, req: %d", err, uid)
		return "", code.ERROR
	}
	if rsp.Common.Code == order_business.RetCode_SUCCESS {
		result = rsp.OrderTxCode
		return result, code.SUCCESS
	}
	vars.ErrorLogger.Errorf(ctx, "GenOrderTxCode req: %d ,resp: %v", uid, json.MarshalToStringNoError(rsp))
	if rsp.OrderTxCode == "" {
		return "", code.ERROR
	}
	return result, code.ERROR
}

func CreateTradeOrder(ctx context.Context, req *args.CreateTradeOrderArgs) (*args.CreateTradeOrderRsp, int) {
	var result args.CreateTradeOrderRsp
	taskGroup, errCtx := errgroup.WithContext(ctx)
	taskGroup.Go(func() error {
		err := goroutine.CheckGoroutineErr(errCtx)
		if err != nil {
			return nil
		}
		// 验证用户物流投递信息
		ret := verifyUserDeliveryInfo(ctx, req.Uid, req.UserDeliveryId)
		if ret != code.SUCCESS {
			return args.NewTaskGroupErr(code.GetMsg(ret), ret)
		}
		return nil
	})
	taskGroup.Go(func() error {
		err := goroutine.CheckGoroutineErr(errCtx)
		if err != nil {
			return nil
		}
		// 验证用户身份
		ret := verifyUserState(ctx, req.Uid)
		if ret != code.SUCCESS {
			return args.NewTaskGroupErr(code.GetMsg(ret), ret)
		}
		return nil
	})
	err := taskGroup.Wait()
	if err != nil {
		taskGroupErr, ok := err.(*args.TaskGroupErr)
		if ok {
			return &result, taskGroupErr.RetCode()
		}
		return &result, code.ERROR
	}

	return createTradeOrder(ctx, req)
}

func createTradeOrder(ctx context.Context, req *args.CreateTradeOrderArgs) (*args.CreateTradeOrderRsp, int) {
	var result args.CreateTradeOrderRsp
	serverName := args.RpcServiceMicroMallOrder
	conn, err := util.GetGrpcClient(ctx, serverName)
	if err != nil {
		vars.ErrorLogger.Errorf(ctx, "GetGrpcClient %q  err: %v", serverName, err)
		return &result, code.ERROR
	}
	//defer conn.Close()
	client := order_business.NewOrderBusinessServiceClient(conn)
	r := order_business.CreateOrderRequest{
		Uid:           req.Uid,
		Time:          util.ParseTimeOfStr(time.Now().Unix()),
		Description:   req.Description,
		PayerClientIp: req.ClientIp,
		DeviceId:      req.DeviceId,
		OrderTxCode:   req.OrderTxCode,
		Detail:        &order_business.OrderDetail{},
		DeliveryInfo:  &order_business.OrderDeliveryInfo{UserDeliveryId: req.UserDeliveryId},
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
		vars.ErrorLogger.Errorf(ctx, "CreateOrder err: %v, req: %v", err, json.MarshalToStringNoError(req))
		return &result, code.ERROR
	}
	if rsp.Common.Code == order_business.RetCode_SUCCESS {
		result.TxCode = rsp.TxCode
		return &result, code.SUCCESS
	}
	vars.ErrorLogger.Errorf(ctx, "CreateOrder req: %v, rsp: %v", json.MarshalToStringNoError(req), json.MarshalToStringNoError(rsp))
	switch rsp.Common.Code {
	case order_business.RetCode_SKU_PRICE_VERSION_NOT_EXIST:
		return &result, code.SkuPriceVersionNotExist
	case order_business.RetCode_ORDER_DELIVERY_NOT_EXIST:
		return &result, code.UserDeliveryInfoNotExist
	case order_business.RetCode_ORDER_TX_CODE_EMPTY:
		return &result, code.TradeOrderTxCodeEmpty
	case order_business.RetCode_ORDER_EXIST: // 如果订单已存在，显示创建成功，防止客户端反复重试
		return &result, code.SUCCESS
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
	default:
		return &result, code.ERROR
	}
}

func verifyTradeOrder(ctx context.Context, uid int64, txCode string) (result args.TradeOrderDetail, retCode int) {
	// 根据交易号获取订单详情
	retCode = code.SUCCESS
	serverName := args.RpcServiceMicroMallOrder
	conn, err := util.GetGrpcClient(ctx, serverName)
	if err != nil {
		vars.ErrorLogger.Errorf(ctx, "GetGrpcClient %q err: %v", serverName, err)
		retCode = code.ERROR
		return
	}
	//defer conn.Close()
	client := order_business.NewOrderBusinessServiceClient(conn)
	r := order_business.GetOrderDetailRequest{TxCode: txCode, Uid: uid}
	rsp, err := client.GetOrderDetail(ctx, &r)
	if err != nil {
		vars.ErrorLogger.Errorf(ctx, "GetOrderDetail err: %v, req: %s %d", err, txCode, uid)
		retCode = code.ERROR
		return
	}
	if rsp.Common.Code != order_business.RetCode_SUCCESS {
		vars.ErrorLogger.Errorf(ctx, "GetOrderDetail req: %s, rsp: %v", txCode, json.MarshalToStringNoError(rsp))
		switch rsp.Common.Code {
		case order_business.RetCode_ORDER_TX_CODE_NOT_EXIST:
			retCode = code.TxCodeNotExist
			return
		case order_business.RetCode_ORDER_PAY_ING:
			retCode = code.OrderPayIng
			return
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
	result.CoinType = int(rsp.CoinType)
	orderList := rsp.List
	// 获取店铺code
	shopIdList := make([]int64, len(orderList))
	for i := 0; i < len(orderList); i++ {
		shopIdList[i] = orderList[i].ShopId
	}
	shopIdToShopCode := make(map[int64]string)
	serverName = args.RpcServiceMicroMallShop
	conn, err = util.GetGrpcClient(ctx, serverName)
	if err != nil {
		kelvins.ErrLogger.Errorf(ctx, "GetGrpcClient %v,err: %v", serverName, err)
		retCode = code.ERROR
		return
	}
	//defer conn.Close()
	serveShop := shop_business.NewShopBusinessServiceClient(conn)
	rShop := shop_business.GetShopMajorInfoRequest{
		ShopIds: shopIdList,
	}
	rspShop, err := serveShop.GetShopMajorInfo(ctx, &rShop)
	if err != nil {
		kelvins.ErrLogger.Errorf(ctx, "GetShopMajorInfo %v,err: %v", serverName, err)
		retCode = code.ERROR
		return
	}
	if rspShop.Common.Code != shop_business.RetCode_SUCCESS {
		kelvins.ErrLogger.Errorf(ctx, "GetShopMajorInfo  req %v,rspShop: %v", json.MarshalToStringNoError(shopIdList), json.MarshalToStringNoError(rspShop))
		switch rspShop.Common.Code {
		case shop_business.RetCode_SHOP_STATE_NOT_VERIFY:
			retCode = code.ShopStateNotVerify
			return
		case shop_business.RetCode_SHOP_NOT_EXIST:
			retCode = code.ErrorShopIdNotExist
			return
		default:
			retCode = code.ERROR
			return
		}
	}
	// 店铺ID和店铺code映射关系
	for i := 0; i < len(rspShop.InfoList); i++ {
		shopIdToShopCode[rspShop.InfoList[i].ShopId] = rspShop.InfoList[i].ShopCode
	}
	if len(shopIdToShopCode) == 0 {
		retCode = code.ErrorShopBusinessNotExist
		return
	}
	result.OrderList = make([]args.TradeShopOrderEntry, len(orderList))
	for i := 0; i < len(orderList); i++ {
		detail := args.TradeShopOrderEntry{
			ShopAccount: shopIdToShopCode[orderList[i].ShopId],
			OrderCode:   orderList[i].OrderCode,
			Description: orderList[i].Description,
			Money:       orderList[i].Money,
		}
		result.OrderList[i] = detail
	}
	return
}

func tradePayVerifyUser(ctx context.Context, uid int64) (account string, retCode int) {
	retCode = code.SUCCESS

	retCode = verifyUserState(ctx, uid)
	if retCode != code.SUCCESS {
		return
	}

	serverName := args.RpcServiceMicroMallUsers
	conn, err := util.GetGrpcClient(ctx, serverName)
	if err != nil {
		kelvins.ErrLogger.Errorf(ctx, "GetGrpcClient %q err: %v", serverName, err)
		retCode = code.ERROR
		return
	}
	//defer conn.Close()
	client := users.NewUsersServiceClient(conn)
	r := users.GetUserAccountIdRequest{
		UidList: []int64{uid},
	}
	rsp, err := client.GetUserAccountId(ctx, &r)
	if err != nil {
		kelvins.ErrLogger.Errorf(ctx, "GetUserAccountId err: %v, req: %v", err, uid)
		retCode = code.ERROR
		return
	}
	switch rsp.Common.Code {
	case users.RetCode_SUCCESS:
		if rsp.InfoList[0].AccountId == "" {
			retCode = code.UserAccountNotExist
			return
		}
		account = rsp.InfoList[0].AccountId
	case users.RetCode_USER_NOT_EXIST:
		retCode = code.ErrorUserNotExist
	default:
		retCode = code.ERROR
	}
	return
}

func decimalZeroCovert(amount string) string {
	if amount == "" {
		return "0"
	}
	return amount
}

func OrderTrade(ctx context.Context, req *args.OrderTradeArgs) (result *args.OrderTradeRsp, retCode int) {
	result = &args.OrderTradeRsp{}
	retCode = code.SUCCESS
	// 验证交易单状态
	orderDetail, ret := verifyTradeOrder(ctx, req.OpUid, req.TxCode)
	if ret != code.SUCCESS {
		retCode = ret
		return
	}
	// 核实用户
	userAccount, ret := tradePayVerifyUser(ctx, req.OpUid)
	if ret != code.SUCCESS {
		retCode = ret
		return
	}
	// 交易支付
	return orderTradePay(ctx, req, userAccount, orderDetail)
}

func orderTradePay(ctx context.Context, req *args.OrderTradeArgs, userAccount string, orderDetail args.TradeOrderDetail) (result *args.OrderTradeRsp, retCode int) {
	result = &args.OrderTradeRsp{}
	// 发起支付流程
	serverName := args.RpcServiceMicroMallPay
	conn, err := util.GetGrpcClient(ctx, serverName)
	if err != nil {
		vars.ErrorLogger.Errorf(ctx, "GetGrpcClient %q err: %v", serverName, err)
		retCode = code.ERROR
		return
	}
	//defer conn.Close()
	payClient := pay_business.NewPayBusinessServiceClient(conn)
	payReq := pay_business.TradePayRequest{
		Account:   userAccount,
		CoinType:  pay_business.CoinType(orderDetail.CoinType),
		OpUid:     req.OpUid,
		OpIp:      req.OpIp,
		OutTxCode: req.TxCode,
	}
	payReq.EntryList = make([]*pay_business.TradePayEntry, len(orderDetail.OrderList))
	for i := 0; i < len(orderDetail.OrderList); i++ {
		tradeEntry := &pay_business.TradePayEntry{
			OutTradeNo:  orderDetail.OrderList[i].OrderCode,
			Description: orderDetail.OrderList[i].Description,
			Merchant:    orderDetail.OrderList[i].ShopAccount,
			Attach:      orderDetail.OrderList[i].Description,
			Detail: &pay_business.TradeGoodsDetail{
				Amount:    decimalZeroCovert(orderDetail.OrderList[i].Money),
				Reduction: decimalZeroCovert(orderDetail.OrderList[i].Reduction),
			},
		}
		payReq.EntryList[i] = tradeEntry
	}
	payRsp, err := payClient.TradePay(ctx, &payReq)
	if err != nil {
		vars.ErrorLogger.Errorf(ctx, "TradePay err: %v, req: %v", err, *req)
		retCode = code.ERROR
		return
	}
	if payRsp.Common.Code == pay_business.RetCode_SUCCESS {
		result.IsSuccess = true
		retCode = code.SUCCESS
		return
	}
	vars.ErrorLogger.Errorf(ctx, "TradePay req: %v, rsp: %v", json.MarshalToStringNoError(req), json.MarshalToStringNoError(payRsp))
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
	case pay_business.RetCode_TRADE_UUID_EMPTY:
		retCode = code.OutTradeEmpty
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
	default:
		retCode = code.ERROR
		return
	}
}

func GetOrderReport(ctx context.Context, req *args.GetOrderReportArgs) (*args.GetOrderReportRsp, int) {
	return getOrderReport(ctx, req)
}

func SearchTradeOrderInfo(ctx context.Context, keyWord string) (result interface{}, retCode int) {
	retCode = code.SUCCESS
	result = make([]*order_business.SearchTradeOrderInfo, 0)
	searchKey := "micro-mall-api:search-trade-order:" + keyWord
	err := vars.G2CacheEngine.Get(searchKey, 15, &result, func() (interface{}, error) {
		ctx, cancel := context.WithTimeout(context.TODO(), 5*time.Second)
		defer cancel()
		list, ret := searchTradeOrderInfo(ctx, keyWord)
		if ret != code.SUCCESS {
			return &list, fmt.Errorf("%v", ret)
		}
		return &list, nil
	})
	if err != nil {
		retCode = code.ERROR
		return
	}
	return
}

func searchTradeOrderInfo(ctx context.Context, query string) (result interface{}, retCode int) {
	retCode = code.SUCCESS
	serverName := args.RpcServiceMicroMallOrder
	conn, err := util.GetGrpcClient(ctx, serverName)
	if err != nil {
		retCode = code.ERROR
		vars.ErrorLogger.Errorf(ctx, "GetGrpcClient %q  err: %v", serverName, err)
		return
	}
	client := order_business.NewOrderBusinessServiceClient(conn)
	rsp, err := client.SearchTradeOrder(ctx, &order_business.SearchTradeOrderRequest{Query: query})
	if err != nil {
		retCode = code.ERROR
		vars.ErrorLogger.Errorf(ctx, "SearchTradeOrder err: %v, query: %v", err, query)
		return
	}
	if rsp.Common.Code != order_business.RetCode_SUCCESS {
		retCode = code.ERROR
		vars.ErrorLogger.Errorf(ctx, "SearchTradeOrder err: %v, query: %v, rsp: %v", err, query, json.MarshalToStringNoError(rsp))
		return
	}
	result = rsp.GetList()
	return
}

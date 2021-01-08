package service

import (
	"context"
	"gitee.com/cristiane/micro-mall-order/model/args"
	"gitee.com/cristiane/micro-mall-order/model/mysql"
	"gitee.com/cristiane/micro-mall-order/pkg/code"
	"gitee.com/cristiane/micro-mall-order/pkg/util"
	"gitee.com/cristiane/micro-mall-order/proto/micro_mall_order_proto/order_business"
	"gitee.com/cristiane/micro-mall-order/proto/micro_mall_shop_proto/shop_business"
	"gitee.com/cristiane/micro-mall-order/proto/micro_mall_sku_proto/sku_business"
	"gitee.com/cristiane/micro-mall-order/proto/micro_mall_users_proto/users"
	"gitee.com/cristiane/micro-mall-order/repository"
	"gitee.com/cristiane/micro-mall-order/vars"
	"gitee.com/kelvins-io/common/errcode"
	"gitee.com/kelvins-io/common/json"
	"gitee.com/kelvins-io/kelvins"
	"github.com/shopspring/decimal"
	"time"
)

func CreateOrder(ctx context.Context, req *order_business.CreateOrderRequest) (result *args.CreateOrderRsp, retCode int) {
	var err error
	result = &args.CreateOrderRsp{
		TxCode: "",
	}
	retCode = code.Success
	// 检查用户
	serverName := args.RpcServiceMicroMallUsers
	conn, err := util.GetGrpcClient(serverName)
	if err != nil {
		kelvins.ErrLogger.Errorf(ctx, "GetGrpcClient %v,err: %v", serverName, err)
		retCode = code.ErrorServer
		return
	}
	defer conn.Close()
	client := users.NewUsersServiceClient(conn)
	r := users.GetUserInfoRequest{
		Uid: req.Uid,
	}
	rsp, err := client.GetUserInfo(ctx, &r)
	if err != nil {
		kelvins.ErrLogger.Errorf(ctx, "GetUserInfo %v,err: %v", serverName, err)
		retCode = code.ErrorServer
		return
	}
	if rsp == nil || rsp.Common.Code != users.RetCode_SUCCESS {
		retCode = code.ErrorServer
		return
	}
	if rsp.Info == nil || rsp.Info.Uid <= 0 {
		retCode = code.UserNotExist
		return
	}
	// 初始订单和订单明细
	shops := req.Detail.ShopDetail
	orderList := make([]mysql.Order, len(shops))
	orderSkuList := make([]mysql.OrderSku, 0)
	txCode := util.GetUUID()
	deductInventoryList := make([]*sku_business.InventoryEntryShop, 0)
	for i := 0; i < len(shops); i++ {
		orderCode := util.GetUUID()
		totalAmount := decimal.NewFromInt(0)
		deductEntryShop := &sku_business.InventoryEntryShop{
			ShopId: shops[i].ShopId,
			Detail: nil,
		}
		deductEntryList := make([]*sku_business.InventoryEntryDetail, 0)
		var skuAmount int64
		for j := 0; j < len(shops[i].Goods); j++ {
			// 统计订单包含商品个数
			skuAmount += shops[i].Goods[j].Amount
			goods := shops[i].Goods[j]
			if shops[i].Goods[j].Price == "" {
				shops[i].Goods[j].Price = "0"
			}
			price, err := decimal.NewFromString(shops[i].Goods[j].Price)
			if err != nil {
				kelvins.ErrLogger.Errorf(ctx, "decimal NewFromString err: %v, Price: %v", err, shops[i].Goods[j].Price)
				retCode = code.ErrorServer
				return
			}
			if shops[i].Goods[j].Reduction == "" {
				shops[i].Goods[j].Reduction = "0"
			}
			reduction, err := decimal.NewFromString(shops[i].Goods[j].Reduction)
			if err != nil {
				kelvins.ErrLogger.Errorf(ctx, "decimal NewFromString err: %v, Reduction: %v", err, shops[i].Goods[j].Reduction)
				retCode = code.ErrorServer
				return
			}
			price = util.DecimalSub(price, reduction)
			temp := util.DecimalMul(price, decimal.NewFromInt(shops[i].Goods[j].Amount))
			totalAmount = util.DecimalAdd(totalAmount, temp)
			orderSku := mysql.OrderSku{
				OrderCode:  orderCode,
				ShopId:     shops[i].ShopId,
				SkuCode:    goods.SkuCode,
				Price:      price.String(), // 满减后的价格
				Amount:     int(goods.Amount),
				Name:       goods.Name,
				CreateTime: time.Now(),
				UpdateTime: time.Now(),
			}
			deductEntry := &sku_business.InventoryEntryDetail{
				SkuCode: goods.SkuCode,
				Amount:  goods.Amount,
			}
			deductEntryList = append(deductEntryList, deductEntry)
			orderSkuList = append(orderSkuList, orderSku)
		}
		deductEntryShop.Detail = deductEntryList
		payExpire := time.Now().Add(30 * time.Minute)
		order := mysql.Order{
			TxCode:       txCode, // 同一个批次下单的订单对应同一个交易code
			OrderCode:    orderCode,
			Uid:          req.Uid,
			OrderTime:    time.Now(),
			Description:  req.Description,
			ClientIp:     req.PayerClientIp,
			DeviceCode:   req.DeviceId,
			ShopId:       shops[i].ShopId,
			ShopName:     shops[i].SceneInfo.StoreInfo.Name,
			ShopAreaCode: shops[i].SceneInfo.StoreInfo.AreaCode,
			ShopAddress:  shops[i].SceneInfo.StoreInfo.Address,
			State:        0,
			PayExpire:    payExpire,
			PayState:     0,
			Amount:       int(skuAmount),
			TotalAmount:  totalAmount.String(),
			CoinType:     int(shops[i].CoinType),
			CreateTime:   time.Now(),
			UpdateTime:   time.Now(),
		}
		orderList[i] = order
		deductInventoryList = append(deductInventoryList, deductEntryShop)
	}

	// 扣减库存
	serverName = args.RpcServiceMicroMallSku
	conn, err = util.GetGrpcClient(serverName)
	if err != nil {
		kelvins.ErrLogger.Errorf(ctx, "GetGrpcClient %v,err: %v", serverName, err)
		retCode = code.ErrorServer
		return
	}
	defer conn.Close()
	skuSer := sku_business.NewSkuBusinessServiceClient(conn)
	skuR := sku_business.DeductInventoryRequest{
		List: deductInventoryList,
		OperationMeta: &sku_business.OperationMeta{
			OpUid: req.Uid,
			OpIp:  req.PayerClientIp,
		},
	}
	skuRsp, err := skuSer.DeductInventory(ctx, &skuR)
	if err != nil {
		kelvins.ErrLogger.Errorf(ctx, "DeductInventory %v,err: %v", serverName, err)
		retCode = code.ErrorServer
		return
	}
	if skuRsp == nil || skuRsp.Common == nil || skuRsp.Common.Code == sku_business.RetCode_ERROR {
		retCode = code.ErrorServer
		return
	}
	if skuRsp.Common.Code == sku_business.RetCode_SKU_AMOUNT_NOT_ENOUGH {
		retCode = code.SkuAmountNotEnough
		return
	}
	if skuRsp.Common.Code == sku_business.RetCode_TRANSACTION_FAILED {
		retCode = code.TransactionFailed
		return
	}

	tx := kelvins.XORM_DBEngine.NewSession()
	err = tx.Begin()
	if err != nil {
		kelvins.ErrLogger.Errorf(ctx, "CreateOrder Begin err: %v", err)
		retCode = code.ErrorServer
		return
	}
	// 创建订单
	err = repository.CreateOrder(tx, orderList)
	if err != nil {
		errRollback := tx.Rollback()
		if errRollback != nil {
			kelvins.ErrLogger.Errorf(ctx, "CreateOrder Rollback err: %v", errRollback)
		}
		kelvins.ErrLogger.Errorf(ctx, "CreateOrder err: %v, orderList: %+v", err, orderList)
		retCode = code.ErrorServer
		return
	}
	// 创建订单明细
	err = repository.CreateOrderSku(tx, orderSkuList)
	if err != nil {
		errRollback := tx.Rollback()
		if errRollback != nil {
			kelvins.ErrLogger.Errorf(ctx, "CreateOrder Rollback err: %v", errRollback)
		}
		kelvins.ErrLogger.Errorf(ctx, "CreateOrderSku err: %v, orderSkuList: %+v", err, orderSkuList)
		retCode = code.ErrorServer
		return
	}
	err = tx.Commit()
	if err != nil {
		kelvins.ErrLogger.Errorf(ctx, "CreateOrder commit err: %v", err)
		retCode = code.ErrorServer
		return
	}
	result.TxCode = txCode

	go func() {
		// 触发订单事件
		pushSer := NewPushNoticeService(vars.TradeOrderQueueServer, PushMsgTag{
			DeliveryTag:    args.TaskNameTradeOrderNotice,
			DeliveryErrTag: args.TaskNameTradeOrderNoticeErr,
			RetryCount:     kelvins.QueueAMQPSetting.TaskRetryCount,
			RetryTimeout:   kelvins.QueueAMQPSetting.TaskRetryTimeout,
		})
		businessMsg := args.CommonBusinessMsg{
			Type: args.TradeOrderEventTypeCreate,
			Tag:  args.GetMsg(args.TradeOrderEventTypeCreate),
			UUID: util.GetUUID(),
			Msg: json.MarshalToStringNoError(args.TradeOrderNotice{
				Uid:    req.Uid,
				Time:   util.ParseTimeOfStr(time.Now().Unix()),
				TxCode: txCode,
			}),
		}
		taskUUID, retCode := pushSer.PushMessage(ctx, businessMsg)
		if retCode != code.Success {
			kelvins.ErrLogger.Errorf(ctx, "trade order businessMsg: %+v  notice send err: ", businessMsg, errcode.GetErrMsg(retCode))
		} else {
			kelvins.BusinessLogger.Infof(ctx, "trade order businessMsg businessMsg: %+v  taskUUID :%v", businessMsg, taskUUID)
		}
	}()

	return
}

func GetOrderDetail(ctx context.Context, req *order_business.GetOrderDetailRequest) (result *args.OrderDetailRsp, retCode int) {
	result = &args.OrderDetailRsp{}
	result.List = make([]args.ShopOrderDetail, 0)
	retCode = code.Success
	// 通过交易号获取订单详细
	orderList, err := repository.GetOrderListByTxCode(req.TxCode)
	if err != nil {
		kelvins.ErrLogger.Errorf(ctx, "GetOrderListByTxCode err: %v, TxCode: %+v", err, req.TxCode)
		retCode = code.ErrorServer
		return
	}
	if len(orderList) <= 0 {
		return
	}
	uid := orderList[0].Uid
	coinType := orderList[0].CoinType
	// 获取订单用户code
	serverName := args.RpcServiceMicroMallUsers
	conn, err := util.GetGrpcClient(serverName)
	if err != nil {
		kelvins.ErrLogger.Errorf(ctx, "GetGrpcClient %v,err: %v", serverName, err)
		retCode = code.ErrorServer
		return
	}
	defer conn.Close()
	serve := users.NewUsersServiceClient(conn)
	r := users.GetUserInfoRequest{
		Uid: uid,
	}
	rsp, err := serve.GetUserInfo(ctx, &r)
	if err != nil {
		kelvins.ErrLogger.Errorf(ctx, "GetUserInfo %v,err: %v", serverName, err)
		retCode = code.ErrorServer
		return
	}
	if rsp == nil || rsp.Common == nil || rsp.Common.Code == users.RetCode_ERROR {
		kelvins.ErrLogger.Errorf(ctx, "GetUserInfo %v, rsp: %v", serverName, rsp.Common.Msg)
		retCode = code.ErrorServer
		return
	}
	if rsp.Common.Code == users.RetCode_USER_EXIST {
		retCode = code.UserNotExist
		return
	}
	if rsp.Info.AccountId == "" {
		retCode = code.UserExist
		return
	}
	result.UserCode = rsp.Info.AccountId
	result.CoinType = coinType
	// 获取店铺code
	shopIdList := make([]int64, len(orderList))
	for i := 0; i < len(orderList); i++ {
		shopIdList[i] = orderList[i].ShopId
	}
	serverName = args.RpcServiceMicroMallShop
	conn, err = util.GetGrpcClient(serverName)
	if err != nil {
		kelvins.ErrLogger.Errorf(ctx, "GetGrpcClient %v,err: %v", serverName, err)
		retCode = code.ErrorServer
		return
	}
	defer conn.Close()
	serveShop := shop_business.NewShopBusinessServiceClient(conn)
	rShop := shop_business.GetShopInfoRequest{
		ShopIds: shopIdList,
	}
	rspShop, err := serveShop.GetShopInfo(ctx, &rShop)
	if err != nil {
		kelvins.ErrLogger.Errorf(ctx, "GetShopInfo %v,err: %v", serverName, err)
		retCode = code.ErrorServer
		return
	}
	if rspShop == nil || rspShop.Common == nil || rspShop.Common.Code == shop_business.RetCode_ERROR {
		kelvins.ErrLogger.Errorf(ctx, "GetShopInfo %v,rspShop: %v", serverName, rspShop.Common.Code)
		retCode = code.ErrorServer
		return
	}
	// 店铺ID和店铺code映射关系
	shopIdToShopCode := make(map[int64]string)
	for i := 0; i < len(rspShop.InfoList); i++ {
		shopIdToShopCode[rspShop.InfoList[i].ShopId] = rspShop.InfoList[i].ShopCode
	}
	//key := args.ConfigKvShopOrderNotifyUrl
	//config, err := repository.GetConfigKv(key)
	//if err != nil {
	//	kelvins.ErrLogger.Errorf(ctx, "GetConfigKv err: %v ,key: %v", err, key)
	//	retCode = code.ErrorServer
	//	return
	//}
	result.List = make([]args.ShopOrderDetail, len(orderList))
	for i := 0; i < len(orderList); i++ {
		detail := args.ShopOrderDetail{
			ShopCode:    shopIdToShopCode[orderList[i].ShopId],
			OrderCode:   orderList[i].OrderCode,
			TimeExpire:  util.ParseTimeOfStr(orderList[i].PayExpire.Unix()),
			Description: orderList[i].Description,
			Amount:      orderList[i].Money,
			CoinType:    orderList[i].CoinType,
			//NotifyUrl:   config.ConfigValue,
		}
		result.List[i] = detail
	}
	return
}

func GetOrderSku(ctx context.Context, req *order_business.GetOrderSkuRequest) (*args.OrderSkuRsp, int) {
	result := &args.OrderSkuRsp{SkuList: make([]args.OrderSku, 0)}
	retCode := code.Success
	orderList, err := repository.GetOrderListByTxCode(req.TxCode)
	if err != nil {
		kelvins.ErrLogger.Errorf(ctx, "GetOrderListByTxCode err: %v ,tx-code: %v", err, req.TxCode)
		retCode = code.ErrorServer
		return result, retCode
	}
	result.SkuList = make([]args.OrderSku, len(orderList))
	for i := 0; i < len(orderList); i++ {
		orderSkuList, err := repository.GetOrderSkuListByOrderCode([]string{orderList[i].OrderCode})
		if err != nil {
			kelvins.ErrLogger.Errorf(ctx, "GetOrderSkuListByOrderCode err: %v ,orderCodeList: %v", err, orderList[i].OrderCode)
			retCode = code.ErrorServer
			return result, retCode
		}
		orderSku := args.OrderSku{
			OrderCode: orderList[i].OrderCode,
			SkuList:   make([]args.OrderSkuEntry, len(orderSkuList)),
		}
		for j := 0; j < len(orderSkuList); j++ {
			entry := args.OrderSkuEntry{
				SkuCode: orderSkuList[j].SkuCode,
				Amount:  orderSkuList[j].Amount,
				Name:    orderSkuList[j].Name,
				Price:   orderSkuList[j].Price,
			}
			orderSku.SkuList[j] = entry
		}
		result.SkuList[i] = orderSku
	}

	return result, retCode
}

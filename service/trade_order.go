package service

import (
	"context"
	"fmt"
	"gitee.com/cristiane/micro-mall-api/model/args"
	"gitee.com/cristiane/micro-mall-api/pkg/code"
	"gitee.com/cristiane/micro-mall-api/pkg/util"
	"gitee.com/cristiane/micro-mall-api/proto/micro_mall_order_proto/order_business"
	"gitee.com/cristiane/micro-mall-api/vars"
	"time"
)

func CreateTradeOrder(ctx context.Context, req *args.CreateTradeOrderArgs) (*args.CreateTradeOrderRsp, int) {
	var result args.CreateTradeOrderRsp
	result.OrderEntryList = make([]args.OrderEntry, 0)
	serverName := args.RpcServiceMicroMallOrder
	conn, err := util.GetGrpcClient(serverName)
	if err != nil {
		vars.ErrorLogger.Errorf(ctx, "GetGrpcClient %v,err: %v", serverName, err)
		return &result, code.ERROR
	}
	defer conn.Close()

	client := order_business.NewOrderBusinessServiceClient(conn)
	fmt.Println("uid ==", req.Uid)
	r := order_business.CreateOrderRequest{
		Uid:           req.Uid,
		Time:          util.ParseTimeOfStr(time.Now().Unix()),
		Description:   req.Description,
		PayerClientIp: req.ClientIp,
		DeviceId:      req.DeviceId,
		Detail: &order_business.OrderDetail{
			ShopDetail: nil,
		},
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

	if rsp == nil || rsp.Common == nil {
		vars.ErrorLogger.Errorf(ctx, "CreateOrder %v,err: %v, rsp: %+v", serverName, err, rsp)
		return &result, code.ERROR
	}
	if rsp.Common.Code != order_business.RetCode_SUCCESS {
		if rsp.Common.Code == order_business.RetCode_USER_NOT_EXIST {
			return &result, code.ERROR_USER_NOT_EXIST
		} else if rsp.Common.Code == order_business.RetCode_USER_EXIST {
			return &result, code.ERROR_USER_EXIST
		} else if rsp.Common.Code == order_business.RetCode_SHOP_EXIST {
			return &result, code.ERROR_SHOP_BUSINESS_EXIST
		} else if rsp.Common.Code == order_business.RetCode_SHOP_NOT_EXIST {
			return &result, code.ERROR_SHOP_BUSINESS_NOT_EXIST
		} else if rsp.Common.Code == order_business.RetCode_SKU_AMOUNT_NOT_ENOUGH {
			return &result, code.ERROR_SKU_AMOUNT_NOT_ENOUGH
		} else {
			return &result, code.ERROR
		}
	}

	result.OrderEntryList = make([]args.OrderEntry, len(rsp.OrderList))
	for i := 0; i < len(rsp.OrderList); i++ {
		orderEntry := args.OrderEntry{
			OrderCode:   rsp.OrderList[i].OrderCode,
			OrderExpire: rsp.OrderList[i].TimeExpire,
		}
		result.OrderEntryList[i] = orderEntry
	}

	return &result, code.SUCCESS
}

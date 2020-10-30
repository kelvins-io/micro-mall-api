package service

import (
	"context"
	"gitee.com/cristiane/micro-mall-api/model/args"
	"gitee.com/cristiane/micro-mall-api/pkg/code"
	"gitee.com/cristiane/micro-mall-api/pkg/util"
	"gitee.com/cristiane/micro-mall-api/proto/micro_mall_logistics_proto/logistics_business"
	"gitee.com/cristiane/micro-mall-api/vars"
)

func ApplyLogistics(ctx context.Context, req *args.ApplyLogisticsArgs) (result *args.ApplyLogisticsRsp, retCode int) {
	result = &args.ApplyLogisticsRsp{}
	retCode = code.SUCCESS
	serverName := args.RpcServiceMicroMallLogistics
	conn, err := util.GetGrpcClient(serverName)
	if err != nil {
		vars.ErrorLogger.Errorf(ctx, "GetGrpcClient %v,err: %v", serverName, err)
		retCode = code.ERROR
		return
	}
	defer conn.Close()
	client := logistics_business.NewLogisticsBusinessServiceClient(conn)
	goods := make([]*logistics_business.GoodsInfo, len(req.Goods))
	for i := 0; i < len(req.Goods); i++ {
		goods[i] = &logistics_business.GoodsInfo{
			SkuCode: req.Goods[i].SkuCode,
			Count:   req.Goods[i].Count,
			Name:    req.Goods[i].Name,
			Kind:    req.Goods[i].Kind,
		}
	}
	r := logistics_business.ApplyLogisticsRequest{
		OutTradeNo:  req.OutTradeNo,
		Courier:     req.Courier,
		CourierType: int32(req.CourierType),
		ReceiveType: int32(req.ReceiveType),
		SendTime:    req.SendTime,
		Customer: &logistics_business.CustomerInfo{
			SendUser:     req.SendUser,
			SendAddr:     req.SendAddr,
			SendPhone:    req.SendPhone,
			SendTime:     req.SendTime,
			ReceiveUser:  req.ReceiveUser,
			ReceiveAddr:  req.ReceiveAddr,
			ReceivePhone: req.ReceivePhone,
		},
		Goods: goods,
	}
	rsp, err := client.ApplyLogistics(ctx, &r)
	if err != nil {
		vars.ErrorLogger.Errorf(ctx, "ApplyLogistics %v,err: %v, req: %+v", serverName, err, r)
		retCode = code.ERROR
		return
	}
	if rsp == nil || rsp.Common == nil || rsp.Common.Code == logistics_business.RetCode_ERROR {
		vars.ErrorLogger.Errorf(ctx, "ApplyLogistics %v,err: %v, rsp: %+v", serverName, err, rsp)
		retCode = code.ERROR
		return
	}
	result.LogisticsCode = rsp.LogisticsCode
	return result, retCode
}

func QueryLogisticsRecord(ctx context.Context, req *args.QueryLogisticsRecordArgs) (*args.QueryLogisticsRecordRsp, int) {
	result := &args.QueryLogisticsRecordRsp{}
	retCode := code.SUCCESS

	return result, retCode
}

func UpdateLogisticsRecord(ctx context.Context, req *args.UpdateLogisticsRecordArgs) (*args.UpdateLogisticsRecordRsp, int) {
	result := &args.UpdateLogisticsRecordRsp{}
	retCode := code.SUCCESS

	return result, retCode
}

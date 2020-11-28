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
	logisticsReq := logistics_business.ApplyLogisticsRequest{
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
	logisticsRsp, err := client.ApplyLogistics(ctx, &logisticsReq)
	if err != nil {
		vars.ErrorLogger.Errorf(ctx, "ApplyLogistics %v,err: %v, req: %+v", serverName, err, logisticsReq)
		retCode = code.ERROR
		return
	}
	if logisticsRsp.Common.Code == logistics_business.RetCode_SUCCESS {
		result.LogisticsCode = logisticsRsp.LogisticsCode
		return
	}
	vars.ErrorLogger.Errorf(ctx, "ApplyLogistics %v,err: %v, rsp: %+v", serverName, err, logisticsRsp)
	switch logisticsRsp.Common.Code {
	case logistics_business.RetCode_LOGISTICS_CODE_EXIST:
		retCode = code.LogisticsRecordExist
	case logistics_business.RetCode_LOGISTICS_CODE_NOT_EXIST:
		retCode = code.LogisticsRecordNotExist
	default:
		retCode = code.ERROR
	}
	return
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

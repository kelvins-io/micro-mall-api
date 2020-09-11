package service

import (
	"context"
	"gitee.com/cristiane/micro-mall-api/model/args"
	"gitee.com/cristiane/micro-mall-api/pkg/code"
	"gitee.com/cristiane/micro-mall-api/pkg/util"
	"gitee.com/cristiane/micro-mall-api/proto/micro_mall_trolley_proto/trolley_business"
	"gitee.com/cristiane/micro-mall-api/vars"
)

func SkuJoinUserTrolley(ctx context.Context, req *args.SkuJoinUserTrolleyArgs) (*args.SkuJoinUserTrolleyRsp, int) {
	var result args.SkuJoinUserTrolleyRsp
	serverName := args.RpcServiceMicroMallTrolley
	conn, err := util.GetGrpcClient(serverName)
	if err != nil {
		vars.ErrorLogger.Errorf(ctx, "GetGrpcClient %v,err: %v", serverName, err)
		return &result, code.ERROR
	}
	defer conn.Close()

	client := trolley_business.NewTrolleyBusinessServiceClient(conn)
	r := trolley_business.JoinSkuRequest{
		Uid:      int64(req.Uid),
		SkuCode:  req.SkuCode,
		ShopId:   int64(req.ShopId),
		Time:     req.Time,
		Count:    int64(req.Count),
		Selected: req.Selected,
	}
	rsp, err := client.JoinSku(ctx, &r)
	if err != nil {
		vars.ErrorLogger.Errorf(ctx, "JoinSku %v,err: %v, req: %+v", serverName, err, r)
		return &result, code.ERROR
	}
	if rsp.Common.Code == trolley_business.RetCode_ERROR {
		return &result, code.ERROR
	}

	if rsp.Common.Code == trolley_business.RetCode_SHOP_NOT_EXIST {
		return &result, code.ERROR_SHOP_ID_NOT_EXIST
	} else if rsp.Common.Code == trolley_business.RetCode_SKU_EXIST {
		return &result, code.ERROR_SHOP_ID_EXIST
	} else if rsp.Common.Code == trolley_business.RetCode_SKU_EXIST {
		return &result, code.ERROR_SKU_CODE_EXIST
	}
	return &result, code.SUCCESS
}

func SkuRemoveUserTrolley(ctx context.Context, req *args.SkuRemoveUserTrolleyArgs) (*args.SkuRemoveUserTrolleyRsp, int) {
	var result args.SkuRemoveUserTrolleyRsp
	serverName := args.RpcServiceMicroMallTrolley
	conn, err := util.GetGrpcClient(serverName)
	if err != nil {
		vars.ErrorLogger.Errorf(ctx, "GetGrpcClient %v,err: %v", serverName, err)
		return &result, code.ERROR
	}
	defer conn.Close()

	client := trolley_business.NewTrolleyBusinessServiceClient(conn)
	r := trolley_business.RemoveSkuRequest{
		Uid:     int64(req.Uid),
		SkuCode: req.SkuCode,
		ShopId:  int64(req.ShopId),
	}
	rsp, err := client.RemoveSku(ctx, &r)
	if err != nil {
		vars.ErrorLogger.Errorf(ctx, "RemoveSku %v,err: %v, req: %+v", serverName, err, r)
		return &result, code.ERROR
	}
	if rsp.Common.Code == trolley_business.RetCode_ERROR {
		return &result, code.ERROR
	}

	if rsp.Common.Code == trolley_business.RetCode_SHOP_NOT_EXIST {
		return &result, code.ERROR_SHOP_ID_NOT_EXIST
	} else if rsp.Common.Code == trolley_business.RetCode_SKU_EXIST {
		return &result, code.ERROR_SHOP_ID_EXIST
	} else if rsp.Common.Code == trolley_business.RetCode_SKU_EXIST {
		return &result, code.ERROR_SKU_CODE_EXIST
	}
	return &result, code.SUCCESS
}

func GetUserTrolleyList(ctx context.Context, uid int64) (*args.UserTrolleyListRsp, int) {
	var result args.UserTrolleyListRsp
	result.List = make([]args.UserTrolleyRecord, 0)
	serverName := args.RpcServiceMicroMallTrolley
	conn, err := util.GetGrpcClient(serverName)
	if err != nil {
		vars.ErrorLogger.Errorf(ctx, "GetGrpcClient %v,err: %v", serverName, err)
		return &result, code.ERROR
	}
	defer conn.Close()

	client := trolley_business.NewTrolleyBusinessServiceClient(conn)
	r := trolley_business.GetUserTrolleyListRequest{
		Uid: uid,
	}
	rsp, err := client.GetUserTrolleyList(ctx, &r)
	if err != nil {
		vars.ErrorLogger.Errorf(ctx, "GetUserTrolleyList %v,err: %v, req: %+v", serverName, err, r)
		return &result, code.ERROR
	}
	if rsp.Common.Code == trolley_business.RetCode_ERROR {
		return &result, code.ERROR
	}
	if rsp.Common.Code == trolley_business.RetCode_USER_NOT_EXIST {
		return &result, code.ERROR_USER_NOT_EXIST
	}
	result.List = make([]args.UserTrolleyRecord, len(rsp.Records))
	for i := 0; i < len(rsp.Records); i++ {
		record := args.UserTrolleyRecord{
			SkuCode:  rsp.Records[i].GetSkuCode(),
			ShopId:   rsp.Records[i].GetShopId(),
			Count:    rsp.Records[i].GetCount(),
			Time:     rsp.Records[i].GetTime(),
			Selected: rsp.Records[i].GetSelected(),
		}
		result.List[i] = record
	}
	return &result, code.SUCCESS
}

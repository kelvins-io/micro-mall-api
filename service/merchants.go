package service

import (
	"context"
	"gitee.com/cristiane/micro-mall-api/model/args"
	"gitee.com/cristiane/micro-mall-api/pkg/code"
	"gitee.com/cristiane/micro-mall-api/pkg/util"
	"gitee.com/cristiane/micro-mall-api/proto/micro_mall_users_proto/users"
	"gitee.com/cristiane/micro-mall-api/vars"
	"gitee.com/kelvins-io/common/json"
)

func MerchantsMaterial(ctx context.Context, req *args.MerchantsMaterialArgs) (*args.MerchantsMaterialRsp, int) {
	var result args.MerchantsMaterialRsp
	serverName := args.RpcServiceMicroMallUsers
	conn, err := util.GetGrpcClient(ctx, serverName)
	if err != nil {
		vars.ErrorLogger.Errorf(ctx, "GetGrpcClient  %q err: %v", serverName, err)
		return &result, code.ERROR
	}
	//defer conn.Close()
	client := users.NewMerchantsServiceClient(conn)
	merchantReq := users.MerchantsMaterialRequest{
		Info: &users.MerchantsMaterialInfo{
			Uid:          int64(req.Uid),
			RegisterAddr: req.RegisterAddr,
			HealthCardNo: req.HealthCardNo,
			Identity:     int32(req.Identity),
			State:        3,
			TaxCardNo:    req.TaxCardNo,
		},
		OperationType: users.OperationType(req.OperationType),
	}
	rsp, err := client.MerchantsMaterial(ctx, &merchantReq)
	if err != nil {
		vars.ErrorLogger.Errorf(ctx, "MerchantsMaterial  err: %v, req: %v", err, json.MarshalToStringNoError(req))
		return &result, code.ERROR
	}

	if rsp.Common.Code == users.RetCode_SUCCESS {
		result.MerchantId = rsp.MaterialId
		return &result, code.SUCCESS
	}

	vars.ErrorLogger.Errorf(ctx, "MerchantsMaterial  req: %v, resp: %v", json.MarshalToStringNoError(req), json.MarshalToStringNoError(rsp))
	switch rsp.Common.Code {
	case users.RetCode_USER_NOT_EXIST:
		return &result, code.ErrorUserNotExist
	case users.RetCode_MERCHANT_EXIST:
		return &result, code.ErrorMerchantExist
	case users.RetCode_MERCHANT_NOT_EXIST:
		return &result, code.ErrorMerchantNotExist
	default:
		return &result, code.ERROR
	}
}

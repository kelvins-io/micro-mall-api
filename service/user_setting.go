package service

import (
	"context"
	"gitee.com/cristiane/micro-mall-api/model/args"
	"gitee.com/cristiane/micro-mall-api/pkg/code"
	"gitee.com/cristiane/micro-mall-api/pkg/util"
	"gitee.com/cristiane/micro-mall-api/proto/micro_mall_users_proto/users"
	"gitee.com/cristiane/micro-mall-api/vars"
)

func ModifyUserSettingAddress(ctx context.Context, req *args.UserSettingAddressPutArgs) int {
	serverName := args.RpcServiceMicroMallUsers
	conn, err := util.GetGrpcClient(serverName)
	if err != nil {
		vars.ErrorLogger.Errorf(ctx, "GetGrpcClient %v,err: %v", serverName, err)
		return code.ERROR
	}
	defer conn.Close()
	client := users.NewUsersServiceClient(conn)
	userReq := users.ModifyUserDeliveryInfoRequest{
		OperationType: users.OperationType(req.OperationType),
		Info: &users.UserDeliveryInfo{
			Id:           req.Id,
			DeliveryUser: req.DeliveryUser,
			MobilePhone:  req.MobilePhone,
			Area:         req.Area,
			DetailedArea: req.DetailedArea,
			Label:        req.Label,
		},
		Uid: int64(req.Uid),
	}
	if req.IsDefault {
		userReq.Info.IsDefault = users.IsDefaultType_DEFAULT_TYPE_TRUE
	} else {
		userReq.Info.IsDefault = users.IsDefaultType_DEFAULT_TYPE_FALSE
	}
	rsp, err := client.ModifyUserDeliveryInfo(ctx, &userReq)
	if err != nil || rsp.Common.Code == users.RetCode_ERROR {
		vars.ErrorLogger.Errorf(ctx, "GetGrpcClient %v,err: %v, req: %+v", serverName, err, userReq)
		return code.ERROR
	}
	if rsp.Common.Code != users.RetCode_SUCCESS {
		switch rsp.Common.Code {
		case users.RetCode_USER_DELIVERY_INFO_EXIST:
			return code.USER_SETTING_INFO_EXIST
		case users.RetCode_USER_DELIVERY_INFO_NOT_EXIST:
			return code.USER_SETTING_INFO_NOT_EXIST
		case users.RetCode_TRANSACTION_FAILED:
			return code.TRANSACTION_FAILED
		}
		vars.ErrorLogger.Errorf(ctx, "GetGrpcClient %v,err: %v, rsp: %+v", serverName, err, rsp)
	}
	return code.SUCCESS
}

func GetUserSettingAddress(ctx context.Context, req *args.UserSettingAddressGetArgs) ([]args.UserDeliveryInfo, int) {
	result := make([]args.UserDeliveryInfo, 0)
	serverName := args.RpcServiceMicroMallUsers
	conn, err := util.GetGrpcClient(serverName)
	if err != nil {
		vars.ErrorLogger.Errorf(ctx, "GetGrpcClient %v,err: %v", serverName, err)
		return result, code.ERROR
	}
	defer conn.Close()
	client := users.NewUsersServiceClient(conn)
	userReq := users.GetUserDeliveryInfoRequest{Uid: int64(req.Uid), UserDeliveryId: int32(req.DeliveryId)}
	rsp, err := client.GetUserDeliveryInfo(ctx, &userReq)
	if err != nil || rsp.Common.Code == users.RetCode_ERROR {
		vars.ErrorLogger.Errorf(ctx, "GetGrpcClient %v,err: %v, req: %+v", serverName, err, userReq)
		return result, code.ERROR
	}
	if rsp.Common.Code != users.RetCode_SUCCESS {
		vars.ErrorLogger.Errorf(ctx, "GetGrpcClient %v,err: %v, rsp: %+v", serverName, err, rsp)
		return result, code.SUCCESS
	}
	result = make([]args.UserDeliveryInfo, len(rsp.Info))
	for i := 0; i < len(rsp.Info); i++ {
		deliveryInfo := args.UserDeliveryInfo{
			Id:           rsp.Info[i].Id,
			DeliveryUser: rsp.Info[i].DeliveryUser,
			MobilePhone:  rsp.Info[i].MobilePhone,
			Area:         rsp.Info[i].Area,
			DetailedArea: rsp.Info[i].DetailedArea,
			Label:        rsp.Info[i].Label,
			IsDefault:    false,
		}
		if rsp.Info[i].IsDefault == users.IsDefaultType_DEFAULT_TYPE_TRUE {
			deliveryInfo.IsDefault = true
		} else {
			deliveryInfo.IsDefault = false
		}
		result[i] = deliveryInfo
	}
	return result, code.SUCCESS
}

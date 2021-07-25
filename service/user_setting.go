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
		vars.ErrorLogger.Errorf(ctx, "GetGrpcClient %q err: %v", serverName, err)
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
	if err != nil {
		vars.ErrorLogger.Errorf(ctx, "ModifyUserDeliveryInfo err: %v, req: %+v, resp: %+v", err, *req, rsp)
		return code.ERROR
	}
	switch rsp.Common.Code {
	case users.RetCode_SUCCESS:
		return code.SUCCESS
	default:
		vars.ErrorLogger.Errorf(ctx, "ModifyUserDeliveryInfo  req: %+v, resp: %+v", *req, rsp)
	}

	switch rsp.Common.Code {
	case users.RetCode_USER_DELIVERY_INFO_EXIST:
		return code.UserSettingInfoExist
	case users.RetCode_USER_DELIVERY_INFO_NOT_EXIST:
		return code.UserSettingInfoNotExist
	case users.RetCode_TRANSACTION_FAILED:
		return code.TransactionFailed
	default:
		return code.ERROR
	}
}

func GetUserSettingAddress(ctx context.Context, req *args.UserSettingAddressGetArgs) ([]args.UserDeliveryInfo, int) {
	result := make([]args.UserDeliveryInfo, 0)
	serverName := args.RpcServiceMicroMallUsers
	conn, err := util.GetGrpcClient(serverName)
	if err != nil {
		vars.ErrorLogger.Errorf(ctx, "GetGrpcClient %q err: %v", serverName, err)
		return result, code.ERROR
	}
	defer conn.Close()
	client := users.NewUsersServiceClient(conn)
	userReq := users.GetUserDeliveryInfoRequest{Uid: int64(req.Uid), UserDeliveryId: int32(req.DeliveryId)}
	rsp, err := client.GetUserDeliveryInfo(ctx, &userReq)
	if err != nil {
		vars.ErrorLogger.Errorf(ctx, "GetUserDeliveryInfo err: %v, req: %+v", err, *req)
		return result, code.ERROR
	}
	if rsp.Common.Code != users.RetCode_SUCCESS {
		vars.ErrorLogger.Errorf(ctx, "GetUserDeliveryInfo  req: %+v, resp: %+v", *req, rsp)
		switch rsp.Common.Code {
		case users.RetCode_USER_NOT_EXIST:
			return result, code.ErrorUserNotExist
		default:
			return result, code.ERROR
		}
	}
	result = make([]args.UserDeliveryInfo, len(rsp.InfoList))
	for i := 0; i < len(rsp.InfoList); i++ {
		deliveryInfo := args.UserDeliveryInfo{
			Id:           rsp.InfoList[i].Id,
			DeliveryUser: rsp.InfoList[i].DeliveryUser,
			MobilePhone:  rsp.InfoList[i].MobilePhone,
			Area:         rsp.InfoList[i].Area,
			DetailedArea: rsp.InfoList[i].DetailedArea,
			Label:        rsp.InfoList[i].Label,
			IsDefault:    false,
		}
		if rsp.InfoList[i].IsDefault == users.IsDefaultType_DEFAULT_TYPE_TRUE {
			deliveryInfo.IsDefault = true
		} else {
			deliveryInfo.IsDefault = false
		}
		result[i] = deliveryInfo
	}
	return result, code.SUCCESS
}

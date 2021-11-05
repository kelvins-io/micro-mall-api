package service

import (
	"context"
	"fmt"
	"gitee.com/cristiane/micro-mall-api/model/args"
	"gitee.com/cristiane/micro-mall-api/pkg/code"
	"gitee.com/cristiane/micro-mall-api/pkg/util"
	"gitee.com/cristiane/micro-mall-api/proto/micro_mall_users_proto/users"
	"gitee.com/cristiane/micro-mall-api/vars"
	"gitee.com/kelvins-io/common/json"
	"time"
)

func ModifyUserSettingDeliveryAddress(ctx context.Context, req *args.UserSettingAddressPutArgs) int {
	serverName := args.RpcServiceMicroMallUsers
	conn, err := util.GetGrpcClient(ctx, serverName)
	if err != nil {
		vars.ErrorLogger.Errorf(ctx, "GetGrpcClient %q err: %v", serverName, err)
		return code.ERROR
	}
	//defer conn.Close()
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
		vars.ErrorLogger.Errorf(ctx, "ModifyUserDeliveryInfo err: %v, req: %v, resp: %v", err, json.MarshalToStringNoError(req), json.MarshalToStringNoError(rsp))
		return code.ERROR
	}
	switch rsp.Common.Code {
	case users.RetCode_SUCCESS:
		return code.SUCCESS
	default:
		vars.ErrorLogger.Errorf(ctx, "ModifyUserDeliveryInfo  req: %v, resp: %v", json.MarshalToStringNoError(req), json.MarshalToStringNoError(rsp))
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

const userDeliveryInfoKey = "micro-mall-api:user-delivery-info:%v-%v"

func GetUserSettingDeliveryInfoAddress(ctx context.Context, uid, deliveryId int) (result []args.UserDeliveryInfo, retCode int) {
	retCode = code.SUCCESS
	result = make([]args.UserDeliveryInfo, 0)
	cacheKey := fmt.Sprintf(userDeliveryInfoKey, uid, deliveryId)
	err := vars.G2CacheEngine.Get(cacheKey, 30, &result, func() (interface{}, error) {
		ctx, cancel := context.WithTimeout(context.TODO(), 5*time.Second)
		defer cancel()
		list, ret := getUserSettingDeliveryInfoAddress(ctx, uid, deliveryId)
		if ret != code.SUCCESS {
			retCode = ret
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

func getUserSettingDeliveryInfoAddress(ctx context.Context, uid, deliveryId int) (result []args.UserDeliveryInfo, retCode int) {
	retCode = code.SUCCESS
	result = make([]args.UserDeliveryInfo, 0)
	serverName := args.RpcServiceMicroMallUsers
	conn, err := util.GetGrpcClient(ctx, serverName)
	if err != nil {
		retCode = code.ERROR
		vars.ErrorLogger.Errorf(ctx, "GetGrpcClient %q err: %v", serverName, err)
		return
	}
	//defer conn.Close()
	client := users.NewUsersServiceClient(conn)
	userReq := users.GetUserDeliveryInfoRequest{Uid: int64(uid), UserDeliveryId: int32(deliveryId)}
	rsp, err := client.GetUserDeliveryInfo(ctx, &userReq)
	if err != nil {
		retCode = code.ERROR
		vars.ErrorLogger.Errorf(ctx, "GetUserDeliveryInfo err: %v, req: %v", err, json.MarshalToStringNoError(fmt.Sprintf("%v-%v", uid, deliveryId)))
		return
	}
	if rsp.Common.Code != users.RetCode_SUCCESS {
		retCode = code.ERROR
		vars.ErrorLogger.Errorf(ctx, "GetUserDeliveryInfo  req: %v, resp: %v", json.MarshalToStringNoError(fmt.Sprintf("%v-%v", uid, deliveryId)), json.MarshalToStringNoError(rsp))
		switch rsp.Common.Code {
		case users.RetCode_USER_NOT_EXIST:
			retCode = code.ErrorUserNotExist
			return
		case users.RetCode_USER_DELIVERY_INFO_NOT_EXIST:
			retCode =  code.UserDeliveryInfoNotExist
			return
		default:
			return
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
	return
}

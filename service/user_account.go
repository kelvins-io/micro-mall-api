package service

import (
	"context"
	"gitee.com/cristiane/micro-mall-api/model/args"
	"gitee.com/cristiane/micro-mall-api/pkg/code"
	"gitee.com/cristiane/micro-mall-api/pkg/util"
	"gitee.com/cristiane/micro-mall-api/proto/micro_mall_users_proto/users"
	"gitee.com/cristiane/micro-mall-api/vars"
)

func UserAccountCharge(ctx context.Context, req *args.UserAccountChargeArgs) (retCode int) {
	retCode = code.SUCCESS
	userInfo, retCode := GetUserInfo(ctx, req.Uid)
	if retCode != code.SUCCESS {
		return
	}
	if userInfo.Id <= 0 {
		retCode = code.ErrorUserNotExist
		return
	}
	serverName := args.RpcServiceMicroMallUsers
	conn, err := util.GetGrpcClient(serverName)
	if err != nil {
		vars.ErrorLogger.Errorf(ctx, "GetGrpcClient %v,err: %v", serverName, err)
		return code.ERROR
	}
	defer conn.Close()
	client := users.NewUsersServiceClient(conn)
	usersReq := users.UserAccountChargeRequest{
		UidList:     []int64{int64(req.Uid)},
		AccountType: users.AccountType(req.AccountType),
		CoinType:    users.CoinType(req.CoinType),
		OutTradeNo: req.OutTradeNo,
		Amount:      req.Amount,
		OpMeta: &users.OperationMeta{
			OpUid:      int64(req.Uid),
			OpIp:       req.Ip,
			OpPlatform: req.DevicePlatform,
			OpDevice:   req.DeviceCode,
		},
	}
	usersRsp, err := client.UserAccountCharge(ctx, &usersReq)
	if err != nil {
		vars.ErrorLogger.Errorf(ctx, "UserAccountCharge %v,err: %v, req: %+v", serverName, err, req)
		return code.ERROR
	}
	if usersRsp.Common.Code == users.RetCode_SUCCESS {
		return
	} else {
		vars.ErrorLogger.Errorf(ctx, "UserAccountCharge %v,err: %v, req: %+v, rsp: %+v", serverName, err, req, usersRsp)
		switch usersRsp.Common.Code {
		case users.RetCode_USER_NOT_EXIST:
			retCode = code.ErrorUserNotExist
		case users.RetCode_ACCOUNT_LOCK:
			retCode = code.UserAccountStateLock
		case users.RetCode_ACCOUNT_INVALID:
			retCode = code.UserAccountStateInvalid
		case users.RetCode_TRANSACTION_FAILED:
			retCode = code.TransactionFailed
		case users.RetCode_ACCOUNT_NOT_EXIST:
			retCode = code.UserAccountNotExist
		case users.RetCode_USER_CHARGE_SUCCESS:
			retCode = code.TradePaySuccess
		case users.RetCode_USER_CHARGE_RUN:
			retCode = code.TradePayRun
		case users.RetCode_USER_CHARGE_TRADE_NO_EMPTY:
			retCode = code.OutTradeEmpty
		default:
			retCode = code.ERROR
		}
		return
	}
}

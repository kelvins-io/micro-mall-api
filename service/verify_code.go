package service

import (
	"context"

	"gitee.com/cristiane/micro-mall-api/model/args"
	"gitee.com/cristiane/micro-mall-api/pkg/code"
	"gitee.com/cristiane/micro-mall-api/pkg/util"
	"gitee.com/cristiane/micro-mall-api/proto/micro_mall_users_proto/users"
	"gitee.com/cristiane/micro-mall-api/vars"
)

func GenVerifyCode(ctx context.Context, req *args.GenVerifyCodeArgs) (retCode int, verifyCode args.UserVerifyCode) {
	retCode = code.SUCCESS
	serverName := args.RpcServiceMicroMallUsers
	conn, err := util.GetGrpcClient(ctx, serverName)
	if err != nil {
		vars.ErrorLogger.Errorf(ctx, "GetGrpcClient %q err: %v", serverName, err)
		return
	}
	client := users.NewUsersServiceClient(conn)
	verifyCodeReq := &users.GenVerifyCodeRequest{
		Uid:          int64(req.Uid),
		CountryCode:  req.CountryCode,
		Phone:        req.Phone,
		BusinessType: int32(req.BusinessType),
		Receiver:     req.ReceiveEmail,
	}
	verifyCodeRsp, err := client.GenVerifyCode(ctx, verifyCodeReq)
	if err != nil {
		vars.ErrorLogger.Errorf(ctx, "GenVerifyCode rpc err: %v, args: %+v", err, *req)
		retCode = code.ERROR
		return
	}

	if verifyCodeRsp.Common.Code == users.RetCode_SUCCESS {
		verifyCode.VerifyCode = verifyCodeRsp.GetVerifyCode()
		verifyCode.Expire = verifyCodeRsp.GetExpire()
		return
	}

	vars.ErrorLogger.Errorf(ctx, "GenVerifyCode rsp: %v, args: %+v", verifyCodeRsp.Common.Code, *req)
	switch verifyCodeRsp.Common.Code {
	case users.RetCode_USER_VERIFY_CODE_LIMITED:
		retCode = code.ErrorVerifyCodeLimited
	case users.RetCode_USER_VERIFY_CODE_FORBIDDEN:
		retCode = code.ErrorVerifyCodeForbidden
	case users.RetCode_USER_VERIFY_CODE_INTERVAL:
		retCode = code.ErrorVerifyCodeInterval
	default:
		retCode = code.ERROR
	}

	return
}

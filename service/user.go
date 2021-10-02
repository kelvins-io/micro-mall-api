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

func CreateUser(ctx context.Context, req *args.RegisterUserArgs) (*args.RegisterUserRsp, int) {
	var result args.RegisterUserRsp
	// 检查验证码
	reqCheckVerifyCode := checkVerifyCodeArgs{
		businessType: args.VerifyCodeRegister,
		countryCode:  req.CountryCode,
		phone:        req.Phone,
		verifyCode:   req.VerifyCode,
	}
	if retCode := checkVerifyCode(ctx, &reqCheckVerifyCode); retCode != code.SUCCESS {
		return &result, retCode
	}

	serverName := args.RpcServiceMicroMallUsers
	conn, err := util.GetGrpcClient(ctx, serverName)
	if err != nil {
		vars.ErrorLogger.Errorf(ctx, "GetGrpcClient %q  err: %v", serverName, err)
		return &result, code.ERROR
	}
	client := users.NewUsersServiceClient(conn)
	//defer conn.Close()
	inviteId := int64(0)
	if req.InviteCode != "" {
		// 检查邀请码
		inviteUserReq := &users.GetUserByInviteCodeRequest{InviteCode: req.InviteCode}
		inviteUser, err := client.GetUserInfoByInviteCode(ctx, inviteUserReq)
		if err != nil {
			vars.ErrorLogger.Errorf(ctx, "GetUserInfoByInviteCode err: %v,req: %q", err, req.InviteCode)
			return &result, code.ERROR
		}
		if inviteUser.Common.Code != users.RetCode_SUCCESS {
			vars.ErrorLogger.Errorf(ctx, "GetUserInfoByInviteCode req: %q, resp: %v", req.InviteCode, json.MarshalToStringNoError(inviteUser))
			return &result, code.ERROR
		}
		if inviteUser.Info.Uid <= 0 {
			return &result, code.ErrorInviteCodeNotExist
		}
		inviteId = int64(int(inviteUser.Info.Uid))
	}
	// 注册用户
	registerReq := &users.RegisterRequest{
		UserName:    req.UserName,
		Sex:         int32(req.Sex),
		CountryCode: req.CountryCode,
		Phone:       req.Phone,
		Email:       req.Email,
		IdCardNo:    req.IdCardNo,
		InviterUser: inviteId,
		ContactAddr: req.ContactAddr,
		Age:         int32(req.Age),
		Password:    req.Password,
	}
	registerRsp, err := client.Register(ctx, registerReq)
	if err != nil {
		vars.ErrorLogger.Errorf(ctx, "GetUserInfoByInviteCode err:%v, req: %v", err, json.MarshalToStringNoError(registerReq))
		return &result, code.ERROR
	}
	if registerRsp.Common.Code == users.RetCode_SUCCESS {
		result.InviteCode = registerRsp.Result.InviteCode
		return &result, code.SUCCESS
	}
	vars.ErrorLogger.Errorf(ctx, "GetUserInfoByInviteCode req: %v, resp: %v", json.MarshalToStringNoError(registerReq), json.MarshalToStringNoError(registerRsp))
	switch registerRsp.Common.Code {
	case users.RetCode_USER_EXIST:
		return &result, code.ErrorUserExist
	default:
		return &result, code.ERROR
	}
}

func LoginUserWithVerifyCode(ctx context.Context, req *args.LoginUserWithVerifyCodeArgs) (string, int) {
	var token string
	reqCheckVerifyCode := checkVerifyCodeArgs{
		businessType: args.VerifyCodeLogin,
		countryCode:  req.CountryCode,
		phone:        req.Phone,
		verifyCode:   req.VerifyCode,
	}
	if retCode := checkVerifyCode(ctx, &reqCheckVerifyCode); retCode != code.SUCCESS {
		return token, retCode
	}
	serverName := args.RpcServiceMicroMallUsers
	conn, err := util.GetGrpcClient(ctx, serverName)
	if err != nil {
		vars.ErrorLogger.Errorf(ctx, "GetGrpcClient %q err: %v", serverName, err)
		return "", code.ERROR
	}
	//defer conn.Close()
	client := users.NewUsersServiceClient(conn)
	loginReq := &users.LoginUserRequest{
		LoginType: users.LoginType_VERIFY_CODE,
		LoginInfo: &users.LoginUserRequest_VerifyCode{
			VerifyCode: &users.LoginVerifyCode{
				Phone: &users.MobilePhone{
					CountryCode: req.CountryCode,
					Phone:       req.Phone,
				},
				VerifyCode: req.VerifyCode,
			},
		},
	}
	loginRsp, err := client.LoginUser(ctx, loginReq)
	if err != nil {
		vars.ErrorLogger.Errorf(ctx, "LoginUser err: %v,req: %v", err, json.MarshalToStringNoError(req))
		return "", code.ERROR
	}
	if loginRsp.Common.Code == users.RetCode_SUCCESS {
		token = loginRsp.IdentityToken
		return token, code.SUCCESS
	}

	vars.ErrorLogger.Errorf(ctx, "LoginUser req: %v,resp: %v", json.MarshalToStringNoError(req), json.MarshalToStringNoError(loginRsp))
	switch loginRsp.Common.Code {
	case users.RetCode_USER_NOT_EXIST:
		return "", code.ErrorUserNotExist
	case users.RetCode_USER_PWD_NOT_MATCH:
		return "", code.ErrorUserPwd
	case users.RetCode_USER_LOGIN_NOT_ALLOW:
		return "", code.UserLoginNotAllow
	default:
		return "", code.ERROR
	}
}

func updateUserStateLogin(ctx context.Context, uid int) int {
	serverName := args.RpcServiceMicroMallUsers
	conn, err := util.GetGrpcClient(ctx, serverName)
	if err != nil {
		vars.ErrorLogger.Errorf(ctx, "GetGrpcClient %q err: %v", serverName, err)
		return code.ERROR
	}
	//defer conn.Close()
	client := users.NewUsersServiceClient(conn)
	req := &users.UpdateUserLoginStateRequest{
		Uid: int64(uid),
		State: &users.UserLoginState{
			Content: "online",
			Time:    time.Now().Unix(),
		},
	}
	rsp, err := client.UpdateUserLoginState(ctx, req)
	if err != nil {
		vars.ErrorLogger.Errorf(ctx, "UpdateUserLoginState err: %v, req: %d", err, uid)
		return code.ERROR
	}

	if rsp.Common.Code == users.RetCode_SUCCESS {
		return code.SUCCESS
	}

	vars.ErrorLogger.Errorf(ctx, "UpdateUserLoginState req: %v, resp: %v", json.MarshalToStringNoError(req), json.MarshalToStringNoError(rsp))
	switch rsp.Common.Code {
	case users.RetCode_USER_NOT_EXIST:
		return code.ErrorUserNotExist
	default:
		return code.ERROR
	}
}

func LoginUserWithPwd(ctx context.Context, req *args.LoginUserWithPwdArgs) (string, int) {
	var token string
	serverName := args.RpcServiceMicroMallUsers
	conn, err := util.GetGrpcClient(ctx, serverName)
	if err != nil {
		vars.ErrorLogger.Errorf(ctx, "GetGrpcClient  %q err: %v", serverName, err)
		return "", code.ERROR
	}
	//defer conn.Close()
	client := users.NewUsersServiceClient(conn)
	loginReq := &users.LoginUserRequest{
		LoginType: users.LoginType_PWD,
		LoginInfo: &users.LoginUserRequest_Pwd{
			Pwd: &users.LoginByPassword{
				LoginKind: users.LoginPwdKind_MOBILE_PHONE,
				Info: &users.LoginByPassword_Phone{
					Phone: &users.MobilePhone{
						CountryCode: req.CountryCode,
						Phone:       req.Phone,
					},
				},
				Pwd: req.Password,
			},
		},
	}
	loginRsp, err := client.LoginUser(ctx, loginReq)
	if err != nil {
		vars.ErrorLogger.Errorf(ctx, "LoginUser err: %v,req: %v", err, json.MarshalToStringNoError(req))
		return "", code.ERROR
	}
	if loginRsp.Common.Code == users.RetCode_SUCCESS {
		token = loginRsp.IdentityToken
		return token, code.SUCCESS
	}

	vars.ErrorLogger.Errorf(ctx, "LoginUser req: %v,resp: %v", json.MarshalToStringNoError(req), json.MarshalToStringNoError(loginRsp))

	switch loginRsp.Common.Code {
	case users.RetCode_USER_NOT_EXIST:
		return "", code.ErrorUserNotExist
	case users.RetCode_USER_PWD_NOT_MATCH:
		return "", code.ErrorUserPwd
	case users.RetCode_USER_LOGIN_NOT_ALLOW:
		return "", code.UserLoginNotAllow
	default:
		return "", code.ERROR
	}
}

func PasswordReset(ctx context.Context, req *args.PasswordResetArgs) int {
	userInfoRsp, ret := GetUserInfo(ctx, req.Uid)
	if ret != code.SUCCESS {
		return ret
	}
	serverName := args.RpcServiceMicroMallUsers
	conn, err := util.GetGrpcClient(ctx, serverName)
	if err != nil {
		vars.ErrorLogger.Errorf(ctx, "GetGrpcClient %q err: %v", serverName, err)
		return code.ERROR
	}
	//defer conn.Close()
	client := users.NewUsersServiceClient(conn)

	reqCheckVerifyCode := checkVerifyCodeArgs{
		businessType: args.VerifyCodePassword,
		countryCode:  userInfoRsp.CountryCode,
		phone:        userInfoRsp.Phone,
		verifyCode:   req.VerifyCode,
	}
	if retCode := checkVerifyCode(ctx, &reqCheckVerifyCode); retCode != code.SUCCESS {
		return retCode
	}
	pwdResetReq := &users.PasswordResetRequest{
		Uid: int64(req.Uid),
		Pwd: req.Password,
	}
	pwdResetRsp, err := client.PasswordReset(ctx, pwdResetReq)
	if err != nil {
		vars.ErrorLogger.Errorf(ctx, "PasswordReset err: %v, req: %v", err, json.MarshalToStringNoError(req))
		return code.ERROR
	}

	if pwdResetRsp.Common.Code == users.RetCode_SUCCESS {
		return code.SUCCESS
	}
	vars.ErrorLogger.Errorf(ctx, "PasswordReset req: %v, resp: %v", json.MarshalToStringNoError(req), json.MarshalToStringNoError(pwdResetRsp))
	switch pwdResetRsp.Common.Code {
	case users.RetCode_USER_NOT_EXIST:
		return code.ErrorUserNotExist
	default:
		return code.ERROR
	}

}

const userInfoCachePhoneKeyPrefix = "micro-mall-api:user_info:phone:%s-%s"

func GetUserInfoByPhone(ctx context.Context, countryCode, phone string) (*users.GetUserInfoByPhoneResponse, int) {
	var result users.GetUserInfoByPhoneResponse
	var err error
	var userInfoCacheKey = fmt.Sprintf(userInfoCachePhoneKeyPrefix, countryCode, phone)
	err = vars.G2CacheEngine.Get(userInfoCacheKey, 60, &result, func() (interface{}, error) {
		serverName := args.RpcServiceMicroMallUsers
		conn, err := util.GetGrpcClient(ctx, serverName)
		if err != nil {
			vars.ErrorLogger.Errorf(ctx, "GetGrpcClient %q err: %v", serverName, err)
			return &result, err
		}
		//defer conn.Close()
		client := users.NewUsersServiceClient(conn)
		userReq := &users.GetUserInfoByPhoneRequest{
			CountryCode: countryCode,
			Phone:       phone,
		}
		userRsp, err := client.GetUserInfoByPhone(ctx, userReq)
		if err != nil {
			vars.ErrorLogger.Errorf(ctx, "GetUserInfoByPhone err:%v, req: %v", err, json.MarshalToStringNoError(userReq))
			return &result, fmt.Errorf("GetUserInfoByPhone err: %v countryCode:%v,phone: %v", err, countryCode, phone)
		}
		if userRsp.Common.Code != users.RetCode_SUCCESS {
			vars.ErrorLogger.Errorf(ctx, "GetUserInfoByPhone userRsp: %v, countryCode: %v,phone: %v", json.MarshalToStringNoError(userRsp), countryCode, phone)
			return &result, fmt.Errorf("GetUserInfoByPhone ret: %d", userRsp.Common.Code)
		}
		if userRsp != nil {
			return userRsp, nil
		}
		return &result, nil
	})
	if err != nil {
		return &result, code.ERROR
	}

	return &result, code.SUCCESS
}

const userInfoCacheUidKeyPrefix = "micro-mall-api:user_info:uid:%d"

func GetUserInfo(ctx context.Context, uid int) (*args.UserInfoRsp, int) {
	var result args.UserInfoRsp
	var userInfoCacheKey = fmt.Sprintf(userInfoCacheUidKeyPrefix, uid)
	var err error
	err = vars.G2CacheEngine.Get(userInfoCacheKey, 60, &result, func() (interface{}, error) {
		serverName := args.RpcServiceMicroMallUsers
		conn, err := util.GetGrpcClient(ctx, serverName)
		if err != nil {
			vars.ErrorLogger.Errorf(ctx, "GetGrpcClient %q err: %v", serverName, err)
			return &result, err
		}
		//defer conn.Close()
		client := users.NewUsersServiceClient(conn)
		req := users.GetUserInfoRequest{
			Uid: int64(uid),
		}
		userInfo, err := client.GetUserInfo(ctx, &req)
		if err != nil {
			vars.ErrorLogger.Errorf(ctx, "GetUserInfo err: %v, req: %d", err, uid)
			return &result, err
		}
		if userInfo.Common.Code != users.RetCode_SUCCESS {
			vars.ErrorLogger.Errorf(ctx, "GetUserInfo  req: %d, rsp: %v", uid, json.MarshalToStringNoError(userInfo))
			return &result, fmt.Errorf("GetUserInfo  uid: %d, resp: %v", uid, json.MarshalToStringNoError(userInfo))
		}
		result = args.UserInfoRsp{
			Id:          uid,
			AccountId:   userInfo.GetInfo().GetAccountId(),
			UserName:    userInfo.GetInfo().GetUserName(),
			Sex:         int(userInfo.GetInfo().GetSex()),
			Phone:       userInfo.GetInfo().GetPhone(),
			CountryCode: userInfo.GetInfo().GetCountryCode(),
			Email:       userInfo.GetInfo().GetEmail(),
			State:       int(userInfo.GetInfo().GetState()),
			IdCardNo:    userInfo.GetInfo().GetIdCardNo(),
			Inviter:     int(userInfo.GetInfo().GetInviter()),
			InviteCode:  userInfo.GetInfo().GetInviterCode(),
			ContactAddr: userInfo.GetInfo().GetContactAddr(),
			Age:         int(userInfo.GetInfo().GetAge()),
			CreateTime:  userInfo.GetInfo().GetCreateTime(),
			UpdateTime:  userInfo.GetInfo().GetUpdateTime(),
		}
		return &result, nil
	})
	if err != nil {
		return &result, code.ERROR
	}

	return &result, code.SUCCESS
}

func ListUserInfo(ctx context.Context, req *args.ListUserInfoArgs) (result args.ListUserInfoRsp, retCode int) {
	result.UserInfoList = make([]args.UserMobilePhone, 0)
	retCode = code.SUCCESS

	serverName := args.RpcServiceMicroMallUsers
	conn, err := util.GetGrpcClient(ctx, serverName)
	if err != nil {
		vars.ErrorLogger.Errorf(ctx, "GetGrpcClient %q err: %v", serverName, err)
		return result, code.ERROR
	}
	//defer conn.Close()
	client := users.NewUsersServiceClient(conn)
	reqList := users.ListUserInfoRequest{
		PageMeta: &users.PageMeta{
			PageNum:  req.PageNum,
			PageSize: req.PageSize,
		},
		Token: req.Token,
	}
	userInfo, err := client.ListUserInfo(ctx, &reqList)
	if err != nil {
		vars.ErrorLogger.Errorf(ctx, "ListUserInfo err: %v, req: %v", err, json.MarshalToStringNoError(req))
		retCode = code.ERROR
		return
	}
	if userInfo.Common.Code != users.RetCode_SUCCESS {
		vars.ErrorLogger.Errorf(ctx, "ListUserInfo  req: %v, rsp: %v", json.MarshalToStringNoError(req), json.MarshalToStringNoError(userInfo))
		retCode = code.ERROR
		return
	}

	infoList := userInfo.UserInfoList
	result.UserInfoList = make([]args.UserMobilePhone, len(infoList))
	for i := 0; i < len(infoList); i++ {
		m := args.UserMobilePhone{
			CountryCode: infoList[i].GetCountryCode(),
			Phone:       infoList[i].GetPhone(),
		}
		result.UserInfoList[i] = m
	}
	return
}

func verifyUserDeliveryInfo(ctx context.Context, uid int64, userDeliveryId int32) int {
	serverName := args.RpcServiceMicroMallUsers
	conn, err := util.GetGrpcClient(ctx, serverName)
	if err != nil {
		vars.ErrorLogger.Errorf(ctx, "GetGrpcClient %q  err: %v", serverName, err)
		return code.ERROR
	}
	//defer conn.Close()
	client := users.NewUsersServiceClient(conn)
	r := users.GetUserDeliveryInfoRequest{
		Uid:            uid,
		UserDeliveryId: userDeliveryId,
	}
	resp, err := client.GetUserDeliveryInfo(ctx, &r)
	if err != nil {
		vars.ErrorLogger.Errorf(ctx, "GetUserDeliveryInfo err: %v, req: %v", err, json.MarshalToStringNoError(r))
		return code.ERROR
	}

	switch resp.Common.Code {
	case users.RetCode_SUCCESS:
		if len(resp.InfoList) == 0 || resp.InfoList[0].Id <= 0 {
			return code.UserDeliveryInfoNotExist
		}
		return code.SUCCESS
	case users.RetCode_USER_NOT_EXIST:
		return code.ErrorUserNotExist
	case users.RetCode_USER_DELIVERY_INFO_NOT_EXIST:
		return code.UserDeliveryInfoNotExist
	default:
		return code.ERROR
	}
}

func verifyUserState(ctx context.Context, uid int64) int {
	serverName := args.RpcServiceMicroMallUsers
	conn, err := util.GetGrpcClient(ctx, serverName)
	if err != nil {
		vars.ErrorLogger.Errorf(ctx, "GetGrpcClient %q  err: %v", serverName, err)
		return code.ERROR
	}
	//defer conn.Close()
	client := users.NewUsersServiceClient(conn)
	req := users.CheckUserStateRequest{UidList: []int64{uid}}
	resp, err := client.CheckUserState(ctx, &req)
	if err != nil {
		vars.ErrorLogger.Errorf(ctx, "CheckUserState err: %v, req: %d", err, uid)
		return code.ERROR
	}
	switch resp.Common.Code {
	case users.RetCode_SUCCESS:
		return code.SUCCESS
	case users.RetCode_USER_NOT_EXIST:
		return code.ErrorUserNotExist
	case users.RetCode_USER_STATE_NOT_VERIFY:
		return code.UserStateNotVerify
	default:
		return code.ERROR
	}
}

func SearchUserInfo(ctx context.Context, query string) (result interface{}, retCode int) {
	retCode = code.SUCCESS
	serverName := args.RpcServiceMicroMallUsers
	conn, err := util.GetGrpcClient(ctx, serverName)
	if err != nil {
		vars.ErrorLogger.Errorf(ctx, "GetGrpcClient %q  err: %v", serverName, err)
		return nil, code.ERROR
	}
	client := users.NewUsersServiceClient(conn)
	rsp, err := client.SearchUserInfo(ctx, &users.SearchUserInfoRequest{Query: query})
	if err != nil {
		vars.ErrorLogger.Errorf(ctx, "UserSearch err: %v, query: %v", err, query)
		return nil, code.ERROR
	}
	if rsp.Common.Code != users.RetCode_SUCCESS {
		vars.ErrorLogger.Errorf(ctx, "UserSearch err: %v, query: %v, rsp: %v", err, query, json.MarshalToStringNoError(rsp))
		return nil, code.ERROR
	}
	return rsp.List, code.SUCCESS
}

func SearchMerchantInfo(ctx context.Context, query string) (result interface{}, retCode int) {
	retCode = code.SUCCESS
	serverName := args.RpcServiceMicroMallUsers
	conn, err := util.GetGrpcClient(ctx, serverName)
	if err != nil {
		vars.ErrorLogger.Errorf(ctx, "GetGrpcClient %q  err: %v", serverName, err)
		return nil, code.ERROR
	}
	client := users.NewMerchantsServiceClient(conn)
	rsp, err := client.SearchMerchantInfo(ctx, &users.SearchMerchantInfoRequest{Query: query})
	if err != nil {
		vars.ErrorLogger.Errorf(ctx, "UserSearch err: %v, query: %v", err, query)
		return nil, code.ERROR
	}
	if rsp.Common.Code != users.RetCode_SUCCESS {
		vars.ErrorLogger.Errorf(ctx, "UserSearch err: %v, query: %v, rsp: %v", err, query, json.MarshalToStringNoError(rsp))
		return nil, code.ERROR
	}
	return rsp.List, code.SUCCESS
}

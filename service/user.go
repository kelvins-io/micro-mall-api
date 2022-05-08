package service

import (
	"context"
	"fmt"
	"time"

	"gitee.com/cristiane/micro-mall-api/model/args"
	"gitee.com/cristiane/micro-mall-api/pkg/code"
	"gitee.com/cristiane/micro-mall-api/pkg/util"
	"gitee.com/cristiane/micro-mall-api/proto/micro_mall_users_proto/users"
	"gitee.com/cristiane/micro-mall-api/vars"
	"gitee.com/kelvins-io/common/json"
)

func CreateUser(ctx context.Context, req *args.RegisterUserArgs) (*args.RegisterUserRsp, int) {
	var result args.RegisterUserRsp
	serverName := args.RpcServiceMicroMallUsers
	conn, err := util.GetGrpcClient(ctx, serverName)
	if err != nil {
		vars.ErrorLogger.Errorf(ctx, "GetGrpcClient %q  err: %v", serverName, err)
		return &result, code.ERROR
	}
	client := users.NewUsersServiceClient(conn)
	// 注册用户
	registerReq := &users.RegisterRequest{
		AccountId:   req.AccountId,
		UserName:    req.UserName,
		Sex:         int32(req.Sex),
		CountryCode: req.CountryCode,
		Phone:       req.Phone,
		Email:       req.Email,
		IdCardNo:    req.IdCardNo,
		InviterUser: req.InviteCode,
		ContactAddr: req.ContactAddr,
		Age:         int32(req.Age),
		Password:    req.Password,
		VerifyCode:  req.VerifyCode,
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
	case users.RetCode_USER_INVITE_CODE_INVALID:
		return &result, code.ErrorInviteCodeNotExist
	case users.RetCode_USER_VERIFY_CODE_INVALID:
		return &result, code.ErrorVerifyCodeInvalid
	case users.RetCode_USER_VERIFY_CODE_EXPIRE:
		return &result, code.ErrorVerifyCodeExpire
	case users.RetCode_USER_VERIFY_CODE_FORBIDDEN:
		return &result, code.ErrorVerifyCodeForbidden
	case users.RetCode_TRANSACTION_FAILED:
		return &result, code.TransactionFailed
	default:
		return &result, code.ERROR
	}
}

func LoginUserWithVerifyCode(ctx context.Context, req *args.LoginUserWithVerifyCodeArgs) (string, int) {
	var token string
	serverName := args.RpcServiceMicroMallUsers
	conn, err := util.GetGrpcClient(ctx, serverName)
	if err != nil {
		vars.ErrorLogger.Errorf(ctx, "GetGrpcClient %q err: %v", serverName, err)
		return "", code.ERROR
	}
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
	case users.RetCode_USER_STATE_NOT_VERIFY:
		return "", code.ErrUserStateNotVerify
	case users.RetCode_USER_NOT_EXIST:
		return "", code.ErrorUserNotExist
	case users.RetCode_USER_VERIFY_CODE_INVALID:
		return "", code.ErrorVerifyCodeInvalid
	case users.RetCode_USER_VERIFY_CODE_EXPIRE:
		return "", code.ErrorVerifyCodeExpire
	case users.RetCode_USER_VERIFY_CODE_FORBIDDEN:
		return "", code.ErrorVerifyCodeForbidden
	case users.RetCode_USER_PWD_NOT_MATCH:
		return "", code.ErrorUserPwd
	case users.RetCode_USER_STATE_FORBIDDEN_LOGIN:
		return "", code.UserStateForbiddenLogin
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

func LoginUserWithAccount(ctx context.Context, req *args.LoginUserWithAccountArgs) (string, int) {
	var token string
	serverName := args.RpcServiceMicroMallUsers
	conn, err := util.GetGrpcClient(ctx, serverName)
	if err != nil {
		vars.ErrorLogger.Errorf(ctx, "GetGrpcClient  %q err: %v", serverName, err)
		return "", code.ERROR
	}
	client := users.NewUsersServiceClient(conn)
	loginReq := &users.LoginUserRequest{
		LoginType: users.LoginType_PWD,
		LoginInfo: &users.LoginUserRequest_Pwd{
			Pwd: &users.LoginByPassword{
				LoginKind: users.LoginPwdKind_ACCOUNT,
				Info: &users.LoginByPassword_Account{
					Account: &users.Account{
						AccountId: req.AccountId,
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
	case users.RetCode_USER_STATE_NOT_VERIFY:
		return "", code.ErrUserStateNotVerify
	case users.RetCode_USER_STATE_FORBIDDEN_LOGIN:
		return "", code.UserStateForbiddenLogin
	default:
		return "", code.ERROR
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
	case users.RetCode_USER_STATE_NOT_VERIFY:
		return "", code.ErrUserStateNotVerify
	case users.RetCode_USER_STATE_FORBIDDEN_LOGIN:
		return "", code.UserStateForbiddenLogin
	default:
		return "", code.ERROR
	}
}

func PasswordReset(ctx context.Context, req *args.PasswordResetArgs) int {
	serverName := args.RpcServiceMicroMallUsers
	conn, err := util.GetGrpcClient(ctx, serverName)
	if err != nil {
		vars.ErrorLogger.Errorf(ctx, "GetGrpcClient %q err: %v", serverName, err)
		return code.ERROR
	}
	//defer conn.Close()
	client := users.NewUsersServiceClient(conn)
	pwdResetReq := &users.PasswordResetRequest{
		Uid:        int64(req.Uid),
		Pwd:        req.Password,
		VerifyCode: req.VerifyCode,
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
	case users.RetCode_USER_STATE_NOT_VERIFY:
		return code.ErrUserStateNotVerify
	case users.RetCode_USER_VERIFY_CODE_INVALID:
		return code.ErrorVerifyCodeInvalid
	case users.RetCode_USER_VERIFY_CODE_EXPIRE:
		return code.ErrorVerifyCodeExpire
	case users.RetCode_USER_VERIFY_CODE_FORBIDDEN:
		return code.ErrorVerifyCodeForbidden
	case users.RetCode_USER_NOT_EXIST:
		return code.ErrorUserNotExist
	case users.RetCode_TRANSACTION_FAILED:
		return code.TransactionFailed
	default:
		return code.ERROR
	}
}

func GetUserInfoByPhone(ctx context.Context, countryCode, phone string) (*users.GetUserInfoByPhoneResponse, int) {
	var result users.GetUserInfoByPhoneResponse
	var err error
	var userInfoCacheKey = fmt.Sprintf(userInfoCachePhoneKeyPrefix, countryCode, phone)
	err = vars.G2CacheEngine.Get(userInfoCacheKey, 15, &result, func() (interface{}, error) {
		ctx, cancel := context.WithTimeout(context.TODO(), 5*time.Second)
		defer cancel()
		return getUserInfoByPhone(ctx, countryCode, phone)
	})
	if err != nil {
		return &result, code.ERROR
	}
	return &result, code.SUCCESS
}

const userInfoCachePhoneKeyPrefix = "micro-mall-api:user_info:phone:%s-%s"

func getUserInfoByPhone(ctx context.Context, countryCode, phone string) (*users.GetUserInfoByPhoneResponse, error) {
	result := users.GetUserInfoByPhoneResponse{}
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
}

func GetUserInfo(ctx context.Context, uid int) (result *args.UserInfoRsp, retCode int) {
	retCode = code.SUCCESS
	result = &args.UserInfoRsp{}
	const userInfoCacheUidKeyPrefix = "micro-mall-api:user_info:uid:%d"
	var userInfoCacheKey = fmt.Sprintf(userInfoCacheUidKeyPrefix, uid)
	err := vars.G2CacheEngine.Get(userInfoCacheKey, 15, result, func() (interface{}, error) {
		ctx, cancel := context.WithTimeout(context.TODO(), 5*time.Second)
		defer cancel()
		s, ret := getUserInfo(ctx, uid)
		if ret != code.SUCCESS {
			return s, fmt.Errorf("%v", ret)
		}
		return s, nil
	})
	if err != nil {
		retCode = code.ERROR
		return
	}
	return
}

func getUserInfo(ctx context.Context, uid int) (result *args.UserInfoRsp, retCode int) {
	retCode = code.SUCCESS
	var err error
	result = &args.UserInfoRsp{}
	serverName := args.RpcServiceMicroMallUsers
	conn, err := util.GetGrpcClient(ctx, serverName)
	if err != nil {
		retCode = code.ERROR
		vars.ErrorLogger.Errorf(ctx, "GetGrpcClient %q err: %v", serverName, err)
		return
	}
	//defer conn.Close()
	client := users.NewUsersServiceClient(conn)
	req := users.GetUserInfoRequest{
		Uid: int64(uid),
	}
	userInfo, err := client.GetUserInfo(ctx, &req)
	if err != nil {
		retCode = code.ERROR
		vars.ErrorLogger.Errorf(ctx, "GetUserInfo err: %v, req: %d", err, uid)
		return
	}
	if userInfo.Common.Code != users.RetCode_SUCCESS {
		retCode = code.ERROR
		vars.ErrorLogger.Errorf(ctx, "GetUserInfo  req: %d, rsp: %v", uid, json.MarshalToStringNoError(userInfo))
		return
	}
	result = &args.UserInfoRsp{
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
	return
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
	list, ret := GetUserSettingDeliveryInfoAddress(ctx, int(uid), int(userDeliveryId))
	if ret != code.SUCCESS {
		return ret
	}
	if len(list) == 0 {
		return code.UserDeliveryInfoNotExist
	}
	return code.SUCCESS
}

func VerifyUserState(ctx context.Context, uid int64) int {
	return verifyUserState(ctx, uid)
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
		return code.ErrUserStateNotVerify
	case users.RetCode_USER_STATE_FORBIDDEN_LOGIN:
		return code.UserStateForbiddenLogin
	default:
		return code.ERROR
	}
}

func SearchUserInfo(ctx context.Context, keyWord string) (result interface{}, retCode int) {
	retCode = code.SUCCESS
	result = make([]*users.SearchUserInfoEntry, 0)
	searchKey := "micro-mall-api:search-user:" + keyWord
	err := vars.G2CacheEngine.Get(searchKey, 15, &result, func() (interface{}, error) {
		ctx, cancel := context.WithTimeout(context.TODO(), 5*time.Second)
		defer cancel()
		list, ret := searchUserInfo(ctx, keyWord)
		if ret != code.SUCCESS {
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

func searchUserInfo(ctx context.Context, query string) (result interface{}, retCode int) {
	retCode = code.SUCCESS
	serverName := args.RpcServiceMicroMallUsers
	conn, err := util.GetGrpcClient(ctx, serverName)
	if err != nil {
		retCode = code.ERROR
		vars.ErrorLogger.Errorf(ctx, "GetGrpcClient %q  err: %v", serverName, err)
		return
	}
	client := users.NewUsersServiceClient(conn)
	rsp, err := client.SearchUserInfo(ctx, &users.SearchUserInfoRequest{Query: query})
	if err != nil {
		retCode = code.ERROR
		vars.ErrorLogger.Errorf(ctx, "UserSearch err: %v, query: %v", err, query)
		return
	}
	if rsp.Common.Code != users.RetCode_SUCCESS {
		retCode = code.ERROR
		vars.ErrorLogger.Errorf(ctx, "UserSearch err: %v, query: %v, rsp: %v", err, query, json.MarshalToStringNoError(rsp))
		return
	}
	result = rsp.GetList()
	return
}

func SearchMerchantInfo(ctx context.Context, keyWord string) (result interface{}, retCode int) {
	retCode = code.SUCCESS
	result = make([]*users.SearchMerchantsInfoEntry, 0)
	searchKey := "micro-mall-api:search-merchant:" + keyWord
	err := vars.G2CacheEngine.Get(searchKey, 15, &result, func() (interface{}, error) {
		ctx, cancel := context.WithTimeout(context.TODO(), 5*time.Second)
		defer cancel()
		list, ret := searchMerchantInfo(ctx, keyWord)
		if ret != code.SUCCESS {
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

func searchMerchantInfo(ctx context.Context, query string) (result interface{}, retCode int) {
	retCode = code.SUCCESS
	serverName := args.RpcServiceMicroMallUsers
	conn, err := util.GetGrpcClient(ctx, serverName)
	if err != nil {
		retCode = code.ERROR
		vars.ErrorLogger.Errorf(ctx, "GetGrpcClient %q  err: %v", serverName, err)
		return
	}
	client := users.NewMerchantsServiceClient(conn)
	rsp, err := client.SearchMerchantInfo(ctx, &users.SearchMerchantInfoRequest{Query: query})
	if err != nil {
		retCode = code.ERROR
		vars.ErrorLogger.Errorf(ctx, "UserSearch err: %v, query: %v", err, query)
		return
	}
	if rsp.Common.Code != users.RetCode_SUCCESS {
		retCode = code.ERROR
		vars.ErrorLogger.Errorf(ctx, "UserSearch err: %v, query: %v, rsp: %v", err, query, json.MarshalToStringNoError(rsp))
		return
	}
	result = rsp.GetList()
	return
}

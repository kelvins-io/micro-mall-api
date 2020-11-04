package service

import (
	"context"
	"fmt"
	"gitee.com/cristiane/go-common/random"
	"gitee.com/cristiane/micro-mall-api/model/args"
	"gitee.com/cristiane/micro-mall-api/model/mysql"
	"gitee.com/cristiane/micro-mall-api/pkg/code"
	"gitee.com/cristiane/micro-mall-api/pkg/util"
	"gitee.com/cristiane/micro-mall-api/pkg/util/email"
	"gitee.com/cristiane/micro-mall-api/proto/micro_mall_users_proto/users"
	"gitee.com/cristiane/micro-mall-api/repository"
	"gitee.com/cristiane/micro-mall-api/vars"
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
	// 通过手机号查询用户是否存在
	serverName := args.RpcServiceMicroMallUsers
	conn, err := util.GetGrpcClient(serverName)
	if err != nil {
		vars.ErrorLogger.Errorf(ctx, "GetGrpcClient %v,err: %v", serverName, err)
		return &result, code.ERROR
	}
	defer conn.Close()

	client := users.NewUsersServiceClient(conn)
	checkUserReq := users.CheckUserByPhoneRequest{
		CountryCode: req.CountryCode,
		Phone:       req.Phone,
	}
	checkResult, err := client.CheckUserByPhone(ctx, &checkUserReq)
	if err != nil || checkResult.Common.Code != users.RetCode_SUCCESS {
		vars.ErrorLogger.Errorf(ctx, "CheckUserByPhone %v,err: %v,r : %+v", serverName, checkUserReq)
		return &result, code.ERROR
	}
	if checkResult.IsExist {
		return &result, code.ERROR_USER_EXIST
	}
	inviteId := int64(0)
	if req.InviteCode != "" {
		// 检查邀请码
		inviteUserReq := &users.GetUserByInviteCodeRequest{InviteCode: req.InviteCode}
		inviteUser, err := client.GetUserInfoByInviteCode(ctx, inviteUserReq)
		if err != nil || inviteUser.Common.Code != users.RetCode_SUCCESS {
			vars.ErrorLogger.Errorf(ctx, "GetUserInfoByInviteCode %v,err: %v,r : %+v", serverName, inviteUserReq)
			return &result, code.ERROR
		}
		if inviteUser.Info.Uid <= 0 {
			return &result, code.ERROR_INVITE_CODE_NOT_EXIST
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
		ContactAddr: req.Email,
		Age:         int32(req.Age),
		Password:    req.Password,
	}
	registerRsp, err := client.Register(ctx, registerReq)
	if err != nil || registerRsp.Common.Code == users.RetCode_ERROR {
		vars.ErrorLogger.Errorf(ctx, "GetUserInfoByInviteCode %v,err: %v,r : %+v", serverName, registerReq)
		return &result, code.ERROR
	}
	switch registerRsp.Common.Code {
	case users.RetCode_USER_EXIST:
		return &result, code.ERROR_USER_EXIST
	}
	result.InviteCode = registerRsp.Result.InviteCode

	return &result, code.SUCCESS
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
	conn, err := util.GetGrpcClient(serverName)
	if err != nil {
		vars.ErrorLogger.Errorf(ctx, "GetGrpcClient %v,err: %v", serverName, err)
		return "", code.ERROR
	}
	defer conn.Close()
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
	if err != nil || loginRsp.Common.Code == users.RetCode_ERROR {
		vars.ErrorLogger.Errorf(ctx, "LoginUser %v,err: %v,r : %+v", serverName, loginReq)
		return "", code.ERROR
	}
	token = loginRsp.IdentityToken
	switch loginRsp.Common.Code {
	case users.RetCode_USER_NOT_EXIST:
		return "", code.ERROR_USER_NOT_EXIST
	case users.RetCode_USER_PWD_NOT_MATCH:
		return "", code.ERROR_USER_PWD
	case users.RetCode_USER_LOGIN_NOT_ALLOW:
		return "", code.USER_LOGIN_NOT_ALLOW
	}

	return token, code.SUCCESS
}

func updateUserStateLogin(ctx context.Context, uid int) int {
	serverName := args.RpcServiceMicroMallUsers
	conn, err := util.GetGrpcClient(serverName)
	if err != nil {
		vars.ErrorLogger.Errorf(ctx, "GetGrpcClient %v,err: %v", serverName, err)
		return code.ERROR
	}
	defer conn.Close()
	client := users.NewUsersServiceClient(conn)
	req := &users.UpdateUserLoginStateRequest{
		Uid: int64(uid),
		State: &users.UserLoginState{
			Content: "online",
			Time:    time.Now().Unix(),
		},
	}
	rsp, err := client.UpdateUserLoginState(ctx, req)
	if err != nil || rsp.Common.Code == users.RetCode_ERROR {
		vars.ErrorLogger.Errorf(ctx, "UpdateUserLoginState %v,err: %v, req: %+v", serverName, err, req)
		return code.ERROR
	}
	return code.SUCCESS
}

func LoginUserWithPwd(ctx context.Context, req *args.LoginUserWithPwdArgs) (string, int) {
	var token string
	serverName := args.RpcServiceMicroMallUsers
	conn, err := util.GetGrpcClient(serverName)
	if err != nil {
		vars.ErrorLogger.Errorf(ctx, "GetGrpcClient %v,err: %v", serverName, err)
		return "", code.ERROR
	}
	defer conn.Close()
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
	if err != nil || loginRsp.Common.Code == users.RetCode_ERROR {
		vars.ErrorLogger.Errorf(ctx, "LoginUser %v,err: %v,r : %+v", serverName, loginReq)
		return "", code.ERROR
	}
	token = loginRsp.IdentityToken
	switch loginRsp.Common.Code {
	case users.RetCode_USER_NOT_EXIST:
		return "", code.ERROR_USER_NOT_EXIST
	case users.RetCode_USER_PWD_NOT_MATCH:
		return "", code.ERROR_USER_PWD
	case users.RetCode_USER_LOGIN_NOT_ALLOW:
		return "", code.USER_LOGIN_NOT_ALLOW
	}

	return token, code.SUCCESS
}

func PasswordReset(ctx context.Context, req *args.PasswordResetArgs) int {
	conn, err := util.GetGrpcClient(args.RpcServiceMicroMallUsers)
	if err != nil {
		vars.ErrorLogger.Errorf(ctx, "GetGrpcClient %v,err: %v", args.RpcServiceMicroMallUsers, err)
		return code.ERROR
	}
	defer conn.Close()
	client := users.NewUsersServiceClient(conn)
	userInfoReq := &users.GetUserInfoRequest{Uid: int64(req.Uid)}
	userInfoRsp, err := client.GetUserInfo(ctx, userInfoReq)
	if err != nil || userInfoRsp.Common.Code == users.RetCode_ERROR {
		vars.ErrorLogger.Errorf(ctx, "GetUserInfo %v,err: %v, req: %+v", args.RpcServiceMicroMallUsers, err, userInfoReq)
		return code.ERROR
	}
	if userInfoRsp.Common.Code == users.RetCode_USER_NOT_EXIST || userInfoRsp.Info.Uid <= 0 {
		return code.ERROR_USER_NOT_EXIST
	}
	reqCheckVerifyCode := checkVerifyCodeArgs{
		businessType: args.VerifyCodePassword,
		countryCode:  userInfoRsp.Info.CountryCode,
		phone:        userInfoRsp.Info.Phone,
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
	if err != nil || pwdResetRsp.Common.Code == users.RetCode_ERROR {
		vars.ErrorLogger.Errorf(ctx, "PasswordReset %v,err: %v, req: %+v", args.RpcServiceMicroMallUsers, err, pwdResetReq)
		return code.ERROR
	}
	if pwdResetRsp.Common.Code == users.RetCode_USER_NOT_EXIST {
		return code.ERROR_USER_NOT_EXIST
	}
	return code.SUCCESS
}

type checkVerifyCodeArgs struct {
	businessType                   int
	countryCode, phone, verifyCode string
}

func checkVerifyCode(ctx context.Context, req *checkVerifyCodeArgs) int {
	record, err := repository.GetVerifyCode(req.businessType, req.countryCode, req.phone, req.verifyCode)
	if err != nil {
		vars.ErrorLogger.Errorf(ctx, "GetVerifyCode err: %v, req: %+v", err, req)
		return code.ERROR
	}
	if record.Id == 0 {
		return code.ERROR_VERIFY_CODE_INVALID
	}
	if int64(record.Expire) < time.Now().Unix() {
		return code.ERROR_VERIFY_CODE_EXPIRE
	}
	return code.SUCCESS
}

func GenVerifyCode(ctx context.Context, req *args.GenVerifyCodeArgs) (retCode int) {
	var err error
	user, err := repository.GetUserByPhone(req.CountryCode, req.Phone)
	if err != nil {
		vars.ErrorLogger.Errorf(ctx, "GetUserByPhone err: %v, req: %+v", err, req)
		return code.ERROR
	}
	//if user.Id == 0 {
	//	return code.ERROR_USER_NOT_EXIST
	//}

	verifyCode := random.KrandNum(6)
	notice := fmt.Sprintf(args.VerifyCodeTemplate, vars.App.Name, verifyCode, args.GetMsg(req.BusinessType), vars.VerifyCodeSetting.ExpireMinute)

	err = email.SendEmailNotice(ctx, req.ReceiveEmail, vars.App.Name, notice)
	if err != nil {
		vars.ErrorLogger.Errorf(ctx, "SendEmailNotice err: %v, req: %+v", err, req)
		return code.ERROR_EMAIL_SEND
	}

	verifyCodeRecord := mysql.VerifyCodeRecord{
		Uid:          user.Id,
		BusinessType: req.BusinessType,
		VerifyCode:   verifyCode,
		Expire:       int(time.Now().Add(time.Duration(vars.VerifyCodeSetting.ExpireMinute) * time.Minute).Unix()),
		CountryCode:  req.CountryCode,
		Phone:        req.Phone,
		Email:        req.ReceiveEmail,
		CreateTime:   time.Now(),
		UpdateTime:   time.Now(),
	}
	err = repository.CreateVerifyCodeRecord(&verifyCodeRecord)
	if err != nil {
		vars.ErrorLogger.Errorf(ctx, "CreateVerifyCodeRecord err: %v, req: %+v", err, req)
		return code.ERROR
	}

	return code.SUCCESS
}

func GetUserInfo(ctx context.Context, uid int) (*args.UserInfoRsp, int) {
	var result args.UserInfoRsp
	serverName := args.RpcServiceMicroMallUsers
	conn, err := util.GetGrpcClient(serverName)
	if err != nil {
		vars.ErrorLogger.Errorf(ctx, "GetGrpcClient %v,err: %v", serverName, err)
		return &result, code.ERROR
	}
	defer conn.Close()

	client := users.NewUsersServiceClient(conn)
	r := users.GetUserInfoRequest{
		Uid: int64(uid),
	}
	userInfo, err := client.GetUserInfo(ctx, &r)
	if err != nil {
		vars.ErrorLogger.Errorf(ctx, "GetUserInfo %v,err: %v, req: %+v", serverName, err, r)
		return &result, code.ERROR
	} else {
		if userInfo != nil && userInfo.Common != nil && userInfo.Common.Code == users.RetCode_SUCCESS {
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
			return &result, code.SUCCESS
		}
		vars.ErrorLogger.Errorf(ctx, "GetUserInfo %v,err: %v, userInfo: %+v", serverName, err, userInfo)
		return &result, code.ERROR
	}
}

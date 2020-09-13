package service

import (
	"context"
	"database/sql"
	"fmt"
	"gitee.com/cristiane/go-common/json"
	"gitee.com/cristiane/go-common/password"
	"gitee.com/cristiane/go-common/random"
	"gitee.com/cristiane/micro-mall-api/model/args"
	"gitee.com/cristiane/micro-mall-api/model/mysql"
	"gitee.com/cristiane/micro-mall-api/pkg/code"
	"gitee.com/cristiane/micro-mall-api/pkg/util"
	"gitee.com/cristiane/micro-mall-api/pkg/util/cache"
	"gitee.com/cristiane/micro-mall-api/pkg/util/email"
	"gitee.com/cristiane/micro-mall-api/proto/micro_mall_users_proto/users"
	"gitee.com/cristiane/micro-mall-api/repository"
	"gitee.com/cristiane/micro-mall-api/vars"
	"strings"
	"time"
)

func CreateUser(ctx context.Context, userInfo *args.RegisterUserArgs) (*args.RegisterUserRsp, int) {
	var result args.RegisterUserRsp
	reqCheckVerifyCode := checkVerifyCodeArgs{
		businessType: args.VerifyCodeRegister,
		countryCode:  userInfo.CountryCode,
		phone:        userInfo.Phone,
		verifyCode:   userInfo.VerifyCode,
	}
	if retCode := checkVerifyCode(ctx, &reqCheckVerifyCode); retCode != code.SUCCESS {
		return &result, retCode
	}

	exist, err := repository.CheckUserExistByPhone(userInfo.CountryCode, userInfo.Phone)
	if err != nil {
		vars.ErrorLogger.Errorf(ctx, "CheckUserExistByPhone err: %v, userInfo: %+v", err, userInfo)
		return &result, code.ERROR
	}
	if exist {
		return &result, code.ERROR_USER_EXIST
	}

	// 检查邀请码
	userRecord, err := repository.GetUserByInviteCode(userInfo.InviteCode)
	if err != nil {
		vars.ErrorLogger.Errorf(ctx, "GetUserByInviteCode err: %v, InviteCode: %+v", err, userInfo.InviteCode)
		return &result, code.ERROR
	}
	if userRecord.Id <= 0 {
		return &result, code.ERROR_INVITE_CODE_NOT_EXIST
	}

	salt := password.GenerateSalt()
	pwd := password.GeneratePassword(userInfo.Password, salt)

	var user = mysql.UserInfo{
		AccountId:    GenAccountId(),
		UserName:     userInfo.UserName,
		Password:     pwd,
		PasswordSalt: salt,
		Sex:          userInfo.Sex,
		Phone:        userInfo.Phone,
		CountryCode:  userInfo.CountryCode,
		Email:        userInfo.Email,
		State:        0,
		IdCardNo: sql.NullString{
			String: userInfo.IdCardNo,
		},
		Inviter:    userRecord.Id,
		InviteCode: GenInviterCode(),
		CreateTime: time.Now(),
		UpdateTime: time.Now(),
	}
	err = repository.CreateUser(&user)
	if err != nil {
		vars.ErrorLogger.Errorf(ctx, "CreateUser err: %v, user: %+v", err, user)
		if strings.Contains(err.Error(), code.GetMsg(code.DB_DUPLICATE_ENTRY)) {
			return &result, code.ERROR_USER_EXIST
		}
		return &result, code.ERROR
	}
	result.InviteCode = user.InviteCode

	pushNoticeService := NewPushNoticeService(vars.QueueServerUserRegisterNotice, PushMsgTag{
		DeliveryTag:    args.TaskNameUserRegisterNotice,
		DeliveryErrTag: args.TaskNameUserRegisterNoticeErr,
		RetryCount:     vars.QueueAMQPSettingUserRegisterNotice.TaskRetryCount,
		RetryTimeout:   vars.QueueAMQPSettingUserRegisterNotice.TaskRetryTimeout,
	})

	businessMsg := args.CommonBusinessMsg{
		Type: args.UserStateEventTypeRegister,
		Tag:  args.GetMsg(args.UserStateEventTypeRegister),
		UUID: genUUID(),
		Msg: json.MarshalToStringNoError(args.UserRegisterNotice{
			CountryCode: userInfo.CountryCode,
			Phone:       userInfo.Phone,
			Time:        util.ParseTimeOfStr(time.Now().Unix()),
			State:       0,
		}),
	}
	taskUUID, retCode := pushNoticeService.PushMessage(ctx, businessMsg)
	if retCode != code.SUCCESS {
		vars.ErrorLogger.Errorf(ctx, "businessMsg: %+v register notice send err: ", businessMsg, code.GetMsg(retCode))
		return &result, code.ERROR
	}
	vars.BusinessLogger.Infof(ctx, "businessMsg: %+v register notice taskUUID :%v", businessMsg, taskUUID)

	return &result, code.SUCCESS
}

func LoginUserWithVerifyCode(ctx context.Context, userInfo *args.LoginUserWithVerifyCodeArgs) (string, int) {
	var token string

	reqCheckVerifyCode := checkVerifyCodeArgs{
		businessType: args.VerifyCodeLogin,
		countryCode:  userInfo.CountryCode,
		phone:        userInfo.Phone,
		verifyCode:   userInfo.VerifyCode,
	}
	if retCode := checkVerifyCode(ctx, &reqCheckVerifyCode); retCode != code.SUCCESS {
		return token, retCode
	}

	user, err := repository.GetUserByPhone(userInfo.CountryCode, userInfo.Phone)
	if err != nil {
		vars.ErrorLogger.Errorf(ctx, "GetUserByPhone err: %v, userInfo: %+v", err, userInfo)
		return token, code.ERROR
	}
	if user.Id == 0 {
		return token, code.ERROR_USER_NOT_EXIST
	}

	token, err = util.GenerateToken(user.UserName, user.Id)
	if err != nil {
		vars.ErrorLogger.Errorf(ctx, "GenerateToken err: %v, user: %+v", err, user)
		return token, code.ERROR
	}

	return token, updateUserStateLogin(ctx, user.Id)
}

func updateUserStateLogin(ctx context.Context, uid int) int {
	state := args.UserOnlineState{
		Uid:   uid,
		State: "online",
		Time:  util.ParseTimeOfStr(time.Now().Unix()),
	}
	userLoginKey := fmt.Sprintf("%v%d", args.CacheKeyUserSate, uid)
	err := cache.Set(vars.RedisPoolMicroMall, userLoginKey, json.MarshalToStringNoError(state), 7200)
	if err != nil {
		vars.ErrorLogger.Errorf(ctx, "setUserState err: %v, userLoginKey: %+v", err, userLoginKey)
		return code.ERROR
	}
	return code.SUCCESS
}

func LoginUserWithPwd(ctx context.Context, userInfo *args.LoginUserWithPwdArgs) (string, int) {
	var token string
	user, err := repository.GetUserByPhone(userInfo.CountryCode, userInfo.Phone)
	if err != nil {
		vars.ErrorLogger.Errorf(ctx, "GetUserByPhone err: %v, userInfo: %+v", err, userInfo)
		return token, code.ERROR
	}
	if user.Id == 0 {
		return token, code.ERROR_USER_NOT_EXIST
	}

	if !password.Check(user.Password, user.PasswordSalt, userInfo.Password) {
		return token, code.ERROR_USER_PWD
	}

	token, err = util.GenerateToken(user.UserName, user.Id)
	if err != nil {
		vars.ErrorLogger.Errorf(ctx, "GenerateToken err: %v, user: %+v", err, user)
		return token, code.ERROR
	}

	return token, updateUserStateLogin(ctx, user.Id)
}

func PasswordReset(ctx context.Context, req *args.PasswordResetArgs) int {
	user, err := repository.GetUserByUid(req.Uid)
	if err != nil {
		vars.ErrorLogger.Errorf(ctx, "GetUserByPhone err: %v, req: %+v", err, req)
		return code.ERROR
	}
	if user.Id == 0 {
		return code.ERROR_USER_NOT_EXIST
	}

	reqCheckVerifyCode := checkVerifyCodeArgs{
		businessType: args.VerifyCodePassword,
		countryCode:  user.CountryCode,
		phone:        user.Phone,
		verifyCode:   req.VerifyCode,
	}
	if retCode := checkVerifyCode(ctx, &reqCheckVerifyCode); retCode != code.SUCCESS {
		return retCode
	}

	query := map[string]interface{}{
		"country_code": user.CountryCode,
		"phone":        user.Phone,
	}
	pwd := password.GeneratePassword(req.Password, user.PasswordSalt)
	maps := map[string]interface{}{
		"password": pwd,
	}
	err = repository.UpdateUserInfo(query, maps)
	if err != nil {
		vars.ErrorLogger.Errorf(ctx, "UpdateUserInfo err: %v, query: %+v，maps： %+v", err, query, maps)
		return code.ERROR
	}

	// 触发密码变更消息
	pushNoticeService := NewPushNoticeService(vars.QueueServerUserStateNotice, PushMsgTag{
		DeliveryTag:    args.TaskNameUserStateNotice,
		DeliveryErrTag: args.TaskNameUserStateNoticeErr,
		RetryCount:     vars.QueueAMQPSettingUserStateNotice.TaskRetryCount,
		RetryTimeout:   vars.QueueAMQPSettingUserStateNotice.TaskRetryTimeout,
	})

	businessMsg := args.CommonBusinessMsg{
		Type: args.UserStateEventTypePwdModify,
		Tag:  args.GetMsg(args.UserStateEventTypePwdModify),
		UUID: genUUID(),
		Msg: json.MarshalToStringNoError(args.UserStateNotice{
			Uid:  user.Id,
			Time: util.ParseTimeOfStr(time.Now().Unix()),
		}),
	}

	taskUUID, retCode := pushNoticeService.PushMessage(ctx, businessMsg)
	if retCode != code.SUCCESS {
		vars.ErrorLogger.Errorf(ctx, "Password Reset businessMsg: %+v  notice send err: ", businessMsg, code.GetMsg(retCode))
		return code.ERROR
	}
	vars.BusinessLogger.Infof(ctx, "Password Reset businessMsg: %+v  taskUUID :%v", businessMsg, taskUUID)

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
		Uid: 10009,
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

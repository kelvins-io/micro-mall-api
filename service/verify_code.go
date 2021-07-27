package service

import (
	"context"
	"fmt"
	"gitee.com/cristiane/micro-mall-api/model/args"
	"gitee.com/cristiane/micro-mall-api/model/mysql"
	"gitee.com/cristiane/micro-mall-api/pkg/code"
	"gitee.com/cristiane/micro-mall-api/pkg/util/email"
	"gitee.com/cristiane/micro-mall-api/repository"
	"gitee.com/cristiane/micro-mall-api/vars"
	"gitee.com/kelvins-io/common/random"
	"strings"
	"time"
)

type checkVerifyCodeArgs struct {
	businessType                   int
	countryCode, phone, verifyCode string
}

func checkVerifyCode(ctx context.Context, req *checkVerifyCodeArgs) int {
	key := fmt.Sprintf("%s:%s-%s-%d", mysql.TableVerifyCodeRecord, req.countryCode, req.phone, req.businessType)
	var obj mysql.VerifyCodeRecord
	err := vars.G2CacheEngine.Get(key, 60*vars.VerifyCodeSetting.ExpireMinute, &obj, func() (interface{}, error) {
		record, err := repository.GetVerifyCode(req.businessType, req.countryCode, req.phone, req.verifyCode)
		return record, err
	})
	if err != nil {
		vars.ErrorLogger.Errorf(ctx, "GetVerifyCode err: %v, req: %+v", err, req)
		return code.ERROR
	}

	if obj.Id == 0 {
		return code.ErrorVerifyCodeInvalid
	}
	if int64(obj.Expire) < time.Now().Unix() {
		return code.ErrorVerifyCodeExpire
	}
	return code.SUCCESS
}

func checkVerifyCodeLimit(limiter repository.CheckVerifyCodeLimiter, key string, limitCount int) (int, int) {
	if limitCount <= 0 {
		limitCount = repository.DefaultVerifyCodeSendPeriodLimitCount
	}
	count, err := limiter.GetVerifyCodePeriodLimitCount(key)
	if err != nil {
		return code.ERROR, count
	}
	if count >= limitCount {
		return code.ErrorVerifyCodeLimited, count
	}

	intervalTime, err := limiter.GetVerifyCodeInterval(key)
	if err != nil {
		return code.ERROR, count
	}
	if intervalTime == 0 {
		return code.SUCCESS, count
	}

	return code.ErrorVerifyCodeInterval, count
}

func GenVerifyCode(ctx context.Context, req *args.GenVerifyCodeArgs) (retCode int, verifyCode string) {
	retCode = code.SUCCESS
	var (
		err error
		//redis : new(repository.CheckVerifyCodeRedisLimiter)
		//local cache map: new(repository.CheckVerifyCodeRedisLimiter)
		limiter = new(repository.CheckVerifyCodeMapCacheLimiter)
	)

	//Limits on the number of verification code requests and time interval
	limitKey := fmt.Sprintf("%s%s", req.CountryCode, req.Phone)
	limitRetCode, limitCount := checkVerifyCodeLimit(limiter, limitKey, vars.VerifyCodeSetting.SendPeriodLimitCount)
	if limitRetCode != code.SUCCESS {
		vars.ErrorLogger.Infof(ctx, "checkVerifyCodeLimit %v %v is limited", req.CountryCode, req.Phone)
		retCode = limitRetCode
		return
	}

	var uid int
	verifyCode = random.KrandNum(6)
	if req.Uid <= 0 {
		userRsp, ret := GetUserInfoByPhone(ctx, req.CountryCode, req.Phone)
		if ret != code.SUCCESS {
			retCode = ret
			return
		}
		if userRsp != nil && userRsp.Info != nil {
			uid = int(userRsp.Info.Uid)
		}
	}
	verifyCodeRecord := mysql.VerifyCodeRecord{
		Uid:          uid,
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
		retCode = code.ERROR
		return
	}

	key := fmt.Sprintf("%s:%s-%s-%d", mysql.TableVerifyCodeRecord, req.CountryCode, req.Phone, req.BusinessType)
	err = vars.G2CacheEngine.Set(key, &verifyCodeRecord, 60*vars.VerifyCodeSetting.ExpireMinute, false)
	if err != nil {
		vars.ErrorLogger.Errorf(ctx, "G2CacheEngine Set err: %v, key: %s,val: %+v", err, key, verifyCodeRecord)
		retCode = code.ERROR
		return
	}

	if req.ReceiveEmail != "" {
		job := func() {
			notice := fmt.Sprintf(args.VerifyCodeTemplate, vars.App.Name, verifyCode, args.GetMsg(req.BusinessType), vars.VerifyCodeSetting.ExpireMinute)
			for _, receiver := range strings.Split(req.ReceiveEmail, ",") {
				err = email.SendEmailNotice(ctx, receiver, vars.App.Name, notice)
				if err != nil {
					vars.ErrorLogger.Errorf(ctx, "SendEmailNotice err %v,receiver:%v, emailNotice: %v", err, receiver, notice)
				}
			}
		}
		vars.GPool.SendJob(job)
	}

	err = limiter.SetVerifyCodeInterval(limitKey, vars.VerifyCodeSetting.SendIntervalExpireSecond)
	if err != nil {
		vars.ErrorLogger.Errorf(ctx, "SetVerifyCodeInterval err: %v, req: %+v", err, req)
		retCode = code.ERROR
		return
	}
	err = limiter.SetVerifyCodePeriodLimitCount(limitKey, limitCount+1, vars.VerifyCodeSetting.SendPeriodLimitExpireSecond)
	if err != nil {
		vars.ErrorLogger.Errorf(ctx, "SetVerifyCodePeriodLimitCount err: %v, req: %+v", err, req)
		retCode = code.ERROR
		return
	}

	return
}
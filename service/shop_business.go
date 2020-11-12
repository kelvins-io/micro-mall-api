package service

import (
	"context"
	"gitee.com/cristiane/micro-mall-api/model/args"
	"gitee.com/cristiane/micro-mall-api/pkg/code"
	"gitee.com/cristiane/micro-mall-api/pkg/util"
	"gitee.com/cristiane/micro-mall-api/proto/micro_mall_shop_proto/shop_business"
	"gitee.com/cristiane/micro-mall-api/vars"
)

func ShopBusinessApply(ctx context.Context, req *args.ShopBusinessInfoArgs) (*args.ShopBusinessInfoRsp, int) {
	var result args.ShopBusinessInfoRsp
	serverName := args.RpcServiceMicroMallShop
	conn, err := util.GetGrpcClient(serverName)
	if err != nil {
		vars.ErrorLogger.Errorf(ctx, "GetGrpcClient %v,err: %v", serverName, err)
		return &result, code.ERROR
	}
	defer conn.Close()

	client := shop_business.NewShopBusinessServiceClient(conn)
	r := shop_business.ShopApplyRequest{
		OperationType:    shop_business.OperationType(req.OperationType),
		OpUid:            int64(req.Uid),
		OpIp:             req.OpIp,
		ShopId:           int64(req.ShopId),
		MerchantId:       int64(req.MerchantId),
		NickName:         req.NickName,
		FullName:         req.FullName,
		RegisterAddr:     req.RegisterAddr,
		BusinessAddr:     req.BusinessAddr,
		BusinessLicense:  req.BusinessLicense,
		TaxCardNo:        req.TaxCardNo,
		BusinessDesc:     req.BusinessDesc,
		SocialCreditCode: req.SocialCreditCode,
		OrganizationCode: req.OrganizationCode,
	}
	rsp, err := client.ShopApply(ctx, &r)
	if err != nil {
		vars.ErrorLogger.Errorf(ctx, "MerchantsMaterial %v,err: %v, req: %+v", serverName, err, r)
		return &result, code.ERROR
	}
	if rsp == nil || rsp.Common == nil || rsp.Common.Code == shop_business.RetCode_ERROR {
		vars.ErrorLogger.Errorf(ctx, "MerchantsMaterial %v,err: %v, rsp: %+v", serverName, err, rsp)
		return &result, code.ERROR
	}
	result.ShopId = int(rsp.ShopId)
	if rsp.Common.Code == shop_business.RetCode_USER_NOT_EXIST {
		return &result, code.ErrorUserNotExist
	} else if rsp.Common.Code == shop_business.RetCode_USER_EXIST {
		return &result, code.ErrorUserExist
	} else if rsp.Common.Code == shop_business.RetCode_MERCHANT_EXIST {
		return &result, code.ErrorMerchantExist
	} else if rsp.Common.Code == shop_business.RetCode_MERCHANT_NOT_EXIST {
		return &result, code.ErrorMerchantNotExist
	} else if rsp.Common.Code == shop_business.RetCode_SHOP_EXIST {
		return &result, code.ErrorShopBusinessExist
	} else if rsp.Common.Code == shop_business.RetCode_SHOP_NOT_EXIST {
		return &result, code.ErrorShopBusinessNotExist
	}
	return &result, code.SUCCESS
}

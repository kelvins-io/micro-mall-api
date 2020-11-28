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
	shopApplyReq := shop_business.ShopApplyRequest{
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
	rsp, err := client.ShopApply(ctx, &shopApplyReq)
	if err != nil {
		vars.ErrorLogger.Errorf(ctx, "ShopApply %v,err: %v, req: %+v", serverName, err, shopApplyReq)
		return &result, code.ERROR
	}
	if rsp.Common.Code == shop_business.RetCode_SUCCESS {
		result.ShopId = int(rsp.ShopId)
		return &result, code.SUCCESS
	}
	vars.ErrorLogger.Errorf(ctx, "ShopApply %v,err: %v, rsp: %+v", serverName, err, rsp)
	switch rsp.Common.Code {
	case shop_business.RetCode_USER_NOT_EXIST:
		return &result, code.ErrorUserNotExist
	case shop_business.RetCode_USER_EXIST:
		return &result, code.ErrorUserExist
	case shop_business.RetCode_MERCHANT_EXIST:
		return &result, code.ErrorMerchantExist
	case shop_business.RetCode_MERCHANT_NOT_EXIST:
		return &result, code.ErrorMerchantNotExist
	case shop_business.RetCode_SHOP_EXIST:
		return &result, code.ErrorShopBusinessExist
	case shop_business.RetCode_SHOP_NOT_EXIST:
		return &result, code.ErrorShopBusinessNotExist
	default:
		return &result, code.ERROR
	}
}

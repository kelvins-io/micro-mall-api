package service

import (
	"context"
	"fmt"
	"gitee.com/cristiane/micro-mall-api/model/args"
	"gitee.com/cristiane/micro-mall-api/pkg/code"
	"gitee.com/cristiane/micro-mall-api/pkg/util"
	"gitee.com/cristiane/micro-mall-api/proto/micro_mall_shop_proto/shop_business"
	"gitee.com/cristiane/micro-mall-api/proto/micro_mall_users_proto/users"
	"gitee.com/cristiane/micro-mall-api/vars"
	"gitee.com/kelvins-io/common/json"
	"time"
)

func ShopBusinessApply(ctx context.Context, req *args.ShopBusinessInfoArgs) (*args.ShopBusinessInfoRsp, int) {
	var result args.ShopBusinessInfoRsp
	if req.MerchantId > 0 {
		serverName := args.RpcServiceMicroMallUsers
		conn, err := util.GetGrpcClient(ctx, serverName)
		if err != nil {
			vars.ErrorLogger.Errorf(ctx, "GetGrpcClient %q err: %v", serverName, err)
			return &result, code.ERROR
		}
		//defer conn.Close()
		client := users.NewMerchantsServiceClient(conn)
		r := users.GetMerchantsMaterialRequest{MaterialId: int64(req.MerchantId)}
		resp, err := client.GetMerchantsMaterial(ctx, &r)
		if err != nil {
			vars.ErrorLogger.Errorf(ctx, "GetMerchantsMaterial err: %v, req: %d", err, r.MaterialId)
			return &result, code.ERROR
		}
		if resp.Common.Code != users.RetCode_SUCCESS {
			return &result, code.ERROR
		}
		vars.ErrorLogger.Errorf(ctx, "GetMerchantsMaterial  req: %d, resp: %v", r.MaterialId, json.MarshalToStringNoError(resp))
		if resp.Info == nil || resp.Info.Uid != int64(req.Uid) {
			return &result, code.MerchantAccountNotExist
		}
	}

	serverName := args.RpcServiceMicroMallShop
	conn, err := util.GetGrpcClient(ctx, serverName)
	if err != nil {
		vars.ErrorLogger.Errorf(ctx, "GetGrpcClient %q err: %v", serverName, err)
		return &result, code.ERROR
	}
	//defer conn.Close()
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
		vars.ErrorLogger.Errorf(ctx, "ShopApply err: %v, req: %v", err, json.MarshalToStringNoError(req))
		return &result, code.ERROR
	}
	if rsp.Common.Code == shop_business.RetCode_SUCCESS {
		result.ShopId = int(rsp.ShopId)
		return &result, code.SUCCESS
	}
	vars.ErrorLogger.Errorf(ctx, "ShopApply req: %v, resp: %v", json.MarshalToStringNoError(req), json.MarshalToStringNoError(rsp))
	switch rsp.Common.Code {
	case shop_business.RetCode_SHOP_EXIST:
		return &result, code.ErrorShopBusinessExist
	case shop_business.RetCode_SHOP_NOT_EXIST:
		return &result, code.ErrorShopBusinessNotExist
	case shop_business.RetCode_TRANSACTION_FAILED:
		return &result, code.TransactionFailed
	default:
		return &result, code.ERROR
	}
}

func SearchShop(ctx context.Context, req *args.SearchShopArgs) (interface{}, int) {
	result := make([]*shop_business.SearchShopInfo, 0)
	keyWord := req.Keyword
	searchKey := "micro-mall-api:search-shop:" + keyWord
	err := vars.G2CacheEngine.Get(searchKey, 15, &result, func() (interface{}, error) {
		ctx, cancel := context.WithTimeout(context.TODO(), 5*time.Second)
		defer cancel()
		list, ret := searchShop(ctx, keyWord)
		if ret != code.SUCCESS {
			return &list, fmt.Errorf("%v", ret)
		}
		return &list, nil
	})
	if err != nil {
		return nil, code.ERROR
	}
	return result, code.SUCCESS
}

func searchShop(ctx context.Context, keyWord string) (result []*shop_business.SearchShopInfo, retCode int) {
	retCode = code.SUCCESS
	result = make([]*shop_business.SearchShopInfo, 0)
	serverName := args.RpcServiceMicroMallShop
	var err error
	conn, err := util.GetGrpcClient(ctx, serverName)
	if err != nil {
		retCode = code.ERROR
		vars.ErrorLogger.Errorf(ctx, "GetGrpcClient %q err: %v", serverName, err)
		return
	}
	//defer conn.Close()
	client := shop_business.NewShopBusinessServiceClient(conn)
	searchReq := &shop_business.SearchShopRequest{Keyword: keyWord}
	searchResp, err := client.SearchShop(ctx, searchReq)
	if err != nil {
		retCode = code.ERROR
		vars.ErrorLogger.Errorf(ctx, "SearchShop  err: %v req: %v", err, json.MarshalToStringNoError(keyWord))
		return
	}
	if searchResp.Common.Code != shop_business.RetCode_SUCCESS {
		retCode = code.ERROR
		vars.ErrorLogger.Errorf(ctx, "SearchShop req: %v, rsp: %v", json.MarshalToStringNoError(keyWord), json.MarshalToStringNoError(searchResp))
		return
	}
	result = searchResp.GetList()
	return
}

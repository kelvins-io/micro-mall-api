package service

import (
	"context"
	"gitee.com/cristiane/micro-mall-api/model/args"
	"gitee.com/cristiane/micro-mall-api/pkg/code"
	"gitee.com/cristiane/micro-mall-api/pkg/util"
	"gitee.com/cristiane/micro-mall-api/proto/micro_mall_shop_proto/shop_business"
	"gitee.com/cristiane/micro-mall-api/proto/micro_mall_sku_proto/sku_business"
	"gitee.com/cristiane/micro-mall-api/vars"
)

func SearchSkuInventory(ctx context.Context, req *args.SearchSkuInventoryArgs) (interface{}, int) {
	serverName := args.RpcServiceMicroMallSku
	conn, err := util.GetGrpcClient(serverName)
	if err != nil {
		vars.ErrorLogger.Errorf(ctx, "GetGrpcClient %v,err: %v", serverName, err)
		return "", code.ERROR
	}
	defer conn.Close()
	client := sku_business.NewSkuBusinessServiceClient(conn)
	searchReq := &sku_business.SearchSkuInventoryRequest{Keyword: req.Keyword}
	searchRsp, err := client.SearchSkuInventory(ctx, searchReq)
	if err != nil {
		vars.ErrorLogger.Errorf(ctx, "SearchSkuInventory %v,err: %v", serverName, err)
		return nil, code.ERROR
	}
	if searchRsp.Common.Code != sku_business.RetCode_SUCCESS {
		vars.ErrorLogger.Errorf(ctx, "SearchSkuInventory %v,err: %v, rsp: %+v", serverName, err, searchRsp)
		return nil, code.ERROR
	}
	return searchRsp.List, code.SUCCESS
}

func SearchShop(ctx context.Context, req *args.SearchShopArgs) (interface{}, int) {
	serverName := args.RpcServiceMicroMallShop
	conn, err := util.GetGrpcClient(serverName)
	if err != nil {
		vars.ErrorLogger.Errorf(ctx, "GetGrpcClient %v,err: %v", serverName, err)
		return "", code.ERROR
	}
	defer conn.Close()
	client := shop_business.NewShopBusinessServiceClient(conn)
	searchReq := &shop_business.SearchShopRequest{Keyword: req.Keyword}
	searchRsp, err := client.SearchShop(ctx, searchReq)
	if err != nil {
		vars.ErrorLogger.Errorf(ctx, "SearchShop %v,err: %v", serverName, err)
		return nil, code.ERROR
	}
	if searchRsp.Common.Code != shop_business.RetCode_SUCCESS {
		vars.ErrorLogger.Errorf(ctx, "SearchShop %v,err: %v, rsp: %+v", serverName, err, searchRsp)
		return nil, code.ERROR
	}
	return searchRsp.List, code.SUCCESS
}

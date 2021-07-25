package service

import (
	"context"
	"gitee.com/cristiane/micro-mall-api/model/args"
	"gitee.com/cristiane/micro-mall-api/pkg/code"
	"gitee.com/cristiane/micro-mall-api/pkg/util"
	"gitee.com/cristiane/micro-mall-api/pkg/util/goroutine"
	"gitee.com/cristiane/micro-mall-api/proto/micro_mall_shop_proto/shop_business"
	"gitee.com/cristiane/micro-mall-api/proto/micro_mall_sku_proto/sku_business"
	"gitee.com/cristiane/micro-mall-api/proto/micro_mall_trolley_proto/trolley_business"
	"gitee.com/cristiane/micro-mall-api/vars"
	"golang.org/x/sync/errgroup"
)

func SkuJoinUserTrolley(ctx context.Context, req *args.SkuJoinUserTrolleyArgs) (*args.SkuJoinUserTrolleyRsp, int) {
	var result args.SkuJoinUserTrolleyRsp

	taskGroup, errCtx := errgroup.WithContext(ctx)
	taskGroup.Go(func() error {
		err := goroutine.CheckGoroutineErr(errCtx)
		if err != nil {
			return nil
		}
		ret := verifyShopBusiness(ctx, []int64{int64(req.ShopId)})
		if ret != code.SUCCESS {
			return args.NewTaskGroupErr(code.GetMsg(ret), ret)
		}
		return nil
	})
	taskGroup.Go(func() error {
		err := goroutine.CheckGoroutineErr(errCtx)
		if err != nil {
			return nil
		}
		ret := verifySkuBusiness(ctx, int64(req.ShopId), req.SkuCode)
		if ret != code.SUCCESS {
			return args.NewTaskGroupErr(code.GetMsg(ret), ret)
		}
		return nil
	})
	err := taskGroup.Wait()
	if err != nil {
		taskGroupErr, ok := err.(*args.TaskGroupErr)
		if ok {
			return &result, taskGroupErr.RetCode()
		}
		return &result, code.ERROR
	}

	serverName := args.RpcServiceMicroMallTrolley
	conn, err := util.GetGrpcClient(serverName)
	if err != nil {
		vars.ErrorLogger.Errorf(ctx, "GetGrpcClient %q err: %v", serverName, err)
		return &result, code.ERROR
	}
	defer conn.Close()

	client := trolley_business.NewTrolleyBusinessServiceClient(conn)
	r := trolley_business.JoinSkuRequest{
		Uid:      int64(req.Uid),
		SkuCode:  req.SkuCode,
		ShopId:   int64(req.ShopId),
		Time:     req.Time,
		Count:    int64(req.Count),
		Selected: req.Selected,
	}
	rsp, err := client.JoinSku(ctx, &r)
	if err != nil {
		vars.ErrorLogger.Errorf(ctx, "JoinSku err: %v, req: %+v", err, *req)
		return &result, code.ERROR
	}

	if rsp.Common.Code == trolley_business.RetCode_SUCCESS {
		return &result, code.SUCCESS
	}

	vars.ErrorLogger.Errorf(ctx, "JoinSku req: %+v, resp: %+v", *req, rsp)
	switch rsp.Common.Code {
	case trolley_business.RetCode_SKU_NOT_EXIST:
		return &result, code.ErrorSkuCodeNotExist
	default:
		return &result, code.ERROR
	}
}

func verifySkuBusiness(ctx context.Context, shopId int64, skuCode string) (retCode int) {
	retCode = code.SUCCESS
	serverName := args.RpcServiceMicroMallSku
	conn, err := util.GetGrpcClient(serverName)
	if err != nil {
		vars.ErrorLogger.Errorf(ctx, "GetGrpcClient %q err: %v", serverName, err)
		retCode = code.ERROR
		return
	}
	defer conn.Close()
	client := sku_business.NewSkuBusinessServiceClient(conn)
	req := sku_business.GetSkuListRequest{
		ShopId:      shopId,
		SkuCodeList: []string{skuCode},
	}
	resp, err := client.GetSkuList(ctx, &req)
	if err != nil {
		vars.ErrorLogger.Errorf(ctx, "GetSkuList err: %v, req: %v %v", err, shopId, skuCode)
		retCode = code.ERROR
		return
	}
	if resp.Common.Code == sku_business.RetCode_SUCCESS {
		if len(resp.List) != 1 {
			retCode = code.ErrorSkuCodeNotExist
			return
		}
		if resp.List[0].SkuCode != skuCode {
			retCode = code.ErrorSkuCodeNotExist
			return
		}
		return
	}
	vars.ErrorLogger.Errorf(ctx, "GetSkuList req: %v %v resp: %v", err, shopId, skuCode, resp)
	switch resp.Common.Code {
	case sku_business.RetCode_INVALID_PARAMETER:
		retCode = code.InvalidParams
	case sku_business.RetCode_SKU_NOT_EXIST:
		retCode = code.ErrorSkuCodeNotExist
	default:
		retCode = code.ERROR
	}
	return
}

func verifyShopBusiness(ctx context.Context, shopIdList []int64) (retCode int) {
	retCode = code.SUCCESS
	serverName := args.RpcServiceMicroMallShop
	conn, err := util.GetGrpcClient(serverName)
	if err != nil {
		vars.ErrorLogger.Errorf(ctx, "GetGrpcClient %q err: %v", serverName, err)
		retCode = code.ERROR
		return
	}
	defer conn.Close()
	shopClient := shop_business.NewShopBusinessServiceClient(conn)
	shopReq := shop_business.GetShopMajorInfoRequest{ShopIds: shopIdList}
	shopResp, err := shopClient.GetShopMajorInfo(ctx, &shopReq)
	if err != nil {
		vars.ErrorLogger.Errorf(ctx, "GetShopMajorInfo err: %v, req: %d", err, shopIdList)
		retCode = code.ERROR
		return
	}
	if shopResp.Common.Code == shop_business.RetCode_SUCCESS {
		return
	}
	vars.ErrorLogger.Errorf(ctx, "GetShopMajorInfo  req: %d, resp: %+v", err, shopIdList, shopResp)
	switch shopResp.Common.Code {
	case shop_business.RetCode_SHOP_NOT_EXIST:
		retCode = code.ErrorShopIdNotExist
		return
	case shop_business.RetCode_SHOP_STATE_NOT_VERIFY:
		retCode = code.ShopStateNotVerify
		return
	default:
		retCode = code.ERROR
		return
	}

}

func SkuRemoveUserTrolley(ctx context.Context, req *args.SkuRemoveUserTrolleyArgs) (*args.SkuRemoveUserTrolleyRsp, int) {
	var result args.SkuRemoveUserTrolleyRsp
	serverName := args.RpcServiceMicroMallTrolley
	conn, err := util.GetGrpcClient(serverName)
	if err != nil {
		vars.ErrorLogger.Errorf(ctx, "GetGrpcClient %q err: %v", serverName, err)
		return &result, code.ERROR
	}
	defer conn.Close()

	client := trolley_business.NewTrolleyBusinessServiceClient(conn)
	r := trolley_business.RemoveSkuRequest{
		Uid:     int64(req.Uid),
		SkuCode: req.SkuCode,
		ShopId:  int64(req.ShopId),
	}
	rsp, err := client.RemoveSku(ctx, &r)
	if err != nil {
		vars.ErrorLogger.Errorf(ctx, "RemoveSku err: %v, req: %+v", err, req)
		return &result, code.ERROR
	}

	if rsp.Common.Code != trolley_business.RetCode_SUCCESS {
		vars.ErrorLogger.Errorf(ctx, "RemoveSku  req: %+v, resp: %+v", req, rsp)
	}
	switch rsp.Common.Code {
	case trolley_business.RetCode_SUCCESS:
		return &result, code.SUCCESS
	case trolley_business.RetCode_SHOP_NOT_EXIST:
		return &result, code.ErrorShopIdNotExist
	case trolley_business.RetCode_SKU_NOT_EXIST:
		return &result, code.ErrorSkuCodeNotExist
	default:
		return &result, code.ERROR
	}
}

func GetUserTrolleyList(ctx context.Context, uid int64) (*args.UserTrolleyListRsp, int) {
	var result args.UserTrolleyListRsp
	result.List = make([]args.UserTrolleyRecord, 0)
	serverName := args.RpcServiceMicroMallTrolley
	conn, err := util.GetGrpcClient(serverName)
	if err != nil {
		vars.ErrorLogger.Errorf(ctx, "GetGrpcClient %q err: %v", serverName, err)
		return &result, code.ERROR
	}
	defer conn.Close()

	client := trolley_business.NewTrolleyBusinessServiceClient(conn)
	r := trolley_business.GetUserTrolleyListRequest{
		Uid: uid,
	}
	rsp, err := client.GetUserTrolleyList(ctx, &r)
	if err != nil {
		vars.ErrorLogger.Errorf(ctx, "GetUserTrolleyList err: %v, req: %d", err, uid)
		return &result, code.ERROR
	}
	switch rsp.Common.Code {
	case trolley_business.RetCode_ERROR:
		return &result, code.ERROR
	default:
	}

	result.List = make([]args.UserTrolleyRecord, len(rsp.Records))
	for i := 0; i < len(rsp.Records); i++ {
		record := args.UserTrolleyRecord{
			SkuCode:  rsp.Records[i].GetSkuCode(),
			ShopId:   rsp.Records[i].GetShopId(),
			Count:    rsp.Records[i].GetCount(),
			Time:     rsp.Records[i].GetTime(),
			Selected: rsp.Records[i].GetSelected(),
		}
		result.List[i] = record
	}
	return &result, code.SUCCESS
}

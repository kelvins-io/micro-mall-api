package service

import (
	"context"
	"gitee.com/cristiane/micro-mall-api/model/args"
	"gitee.com/cristiane/micro-mall-api/pkg/code"
	"gitee.com/cristiane/micro-mall-api/pkg/util"
	"gitee.com/cristiane/micro-mall-api/proto/micro_mall_sku_proto/sku_business"
	"gitee.com/cristiane/micro-mall-api/vars"
	"gitee.com/kelvins-io/common/json"
)

func SkuPutAway(ctx context.Context, req *args.SkuBusinessPutAwayArgs) (*args.SkuBusinessPutAwayRsp, int) {
	var result args.SkuBusinessPutAwayRsp
	if req.ShopId > 0 {
		ret := verifyShopBusiness(ctx, []int64{req.ShopId})
		if ret != code.SUCCESS {
			return &result, ret
		}
	}

	serverName := args.RpcServiceMicroMallSku
	conn, err := util.GetGrpcClient(ctx, serverName)
	if err != nil {
		vars.ErrorLogger.Errorf(ctx, "GetGrpcClient %q err: %v", serverName, err)
		return &result, code.ERROR
	}
	//defer conn.Close()
	client := sku_business.NewSkuBusinessServiceClient(conn)
	r := sku_business.PutAwaySkuRequest{
		OperationType: sku_business.OperationType(req.OperationType),
		OperationMeta: &sku_business.OperationMeta{
			OpUid: int64(req.Uid),
			OpIp:  req.OpIp,
		},
		Sku: &sku_business.SkuInventoryInfo{
			SkuCode:       req.SkuCode,
			Name:          req.Name,
			Price:         req.Price,
			Title:         req.Title,
			SubTitle:      req.SubTitle,
			Desc:          req.Desc,
			Production:    req.Production,
			Supplier:      req.Supplier,
			Category:      req.Category,
			Color:         req.Color,
			ColorCode:     req.ColorCode,
			Specification: req.Specification,
			DescLink:      req.DescLink,
			State:         req.State,
			Amount:        req.Amount,
			ShopId:        req.ShopId,
		},
	}
	rsp, err := client.PutAwaySku(ctx, &r)
	if err != nil {
		vars.ErrorLogger.Errorf(ctx, "PutAwaySku err: %v, req: %v", err, json.MarshalToStringNoError(req))
		return &result, code.ERROR
	}
	if rsp.Common.Code == sku_business.RetCode_SUCCESS {
		return &result, code.SUCCESS
	}
	vars.ErrorLogger.Errorf(ctx, "PutAwaySku req: %v, rsp: %v", json.MarshalToStringNoError(req), json.MarshalToStringNoError(rsp))
	switch rsp.Common.Code {
	case sku_business.RetCode_SKU_EXIST:
		return &result, code.ErrorSkuCodeExist
	case sku_business.RetCode_SKU_NOT_EXIST:
		return &result, code.ErrorSkuCodeNotExist
	case sku_business.RetCode_TRANSACTION_FAILED:
		return &result, code.TransactionFailed
	default:
		return &result, code.ERROR
	}
}

func GetSkuList(ctx context.Context, req *args.GetSkuListArgs) (*args.GetSkuListRsp, int) {
	var result args.GetSkuListRsp
	result.SkuInventoryInfoList = make([]args.SkuInventoryInfo, 0)
	serverName := args.RpcServiceMicroMallSku
	conn, err := util.GetGrpcClient(ctx, serverName)
	if err != nil {
		vars.ErrorLogger.Errorf(ctx, "GetGrpcClient  %q err: %v", serverName, err)
		return &result, code.ERROR
	}
	//defer conn.Close()

	client := sku_business.NewSkuBusinessServiceClient(conn)
	r := sku_business.GetSkuListRequest{
		ShopId:      req.ShopId,
		SkuCodeList: req.SkuCodeList,
	}
	rsp, err := client.GetSkuList(ctx, &r)
	if err != nil {
		vars.ErrorLogger.Errorf(ctx, "GetSkuList err: %v, req: %v", err, json.MarshalToStringNoError(req))
		return &result, code.ERROR
	}
	if rsp.Common.Code != sku_business.RetCode_SUCCESS {
		vars.ErrorLogger.Errorf(ctx, "GetSkuList req: %v, resp: %v", json.MarshalToStringNoError(req), json.MarshalToStringNoError(rsp))
		return &result, code.ERROR
	}

	result.SkuInventoryInfoList = make([]args.SkuInventoryInfo, len(rsp.List))
	for i := 0; i < len(rsp.List); i++ {
		info := args.SkuInventoryInfo{
			SkuCode:       rsp.List[i].GetSkuCode(),
			Name:          rsp.List[i].GetName(),
			Price:         rsp.List[i].GetPrice(),
			Title:         rsp.List[i].GetTitle(),
			SubTitle:      rsp.List[i].GetSubTitle(),
			Desc:          rsp.List[i].GetDesc(),
			Production:    rsp.List[i].GetProduction(),
			Supplier:      rsp.List[i].GetSupplier(),
			Category:      rsp.List[i].GetCategory(),
			Color:         rsp.List[i].GetColor(),
			ColorCode:     rsp.List[i].GetColorCode(),
			Specification: rsp.List[i].GetSpecification(),
			DescLink:      rsp.List[i].GetDescLink(),
			State:         rsp.List[i].GetState(),
			Amount:        rsp.List[i].GetAmount(),
			ShopId:        rsp.List[i].GetShopId(),
			Version:       rsp.List[i].GetVersion(),
		}
		result.SkuInventoryInfoList[i] = info
	}

	return &result, code.SUCCESS
}

func SkuSupplementProperty(ctx context.Context, req *args.SkuPropertyExArgs) (*args.SkuPropertyExRsp, int) {
	var result args.SkuPropertyExRsp
	if req.ShopId > 0 {
		ret := verifyShopBusiness(ctx, []int64{req.ShopId})
		if ret != code.SUCCESS {
			return &result, ret
		}
	}
	serverName := args.RpcServiceMicroMallSku
	conn, err := util.GetGrpcClient(ctx, serverName)
	if err != nil {
		vars.ErrorLogger.Errorf(ctx, "GetGrpcClient %q err: %v", serverName, err)
		return &result, code.ERROR
	}
	//defer conn.Close()
	client := sku_business.NewSkuBusinessServiceClient(conn)
	r := sku_business.SupplementSkuPropertyRequest{
		OperationMeta: &sku_business.OperationMeta{
			OpUid: int64(req.Uid),
			OpIp:  req.OpIp,
		},
		ShopId:            req.ShopId,
		SkuCode:           req.SkuCode,
		Size:              req.Size,
		Shape:             req.Shape,
		ProductionCountry: req.ProductionCountry,
		ProductionDate:    req.ProductionDate,
		ShelfLife:         req.ShelfLife,
		Name:              req.Name,
		OperationType:     sku_business.OperationType(req.OperationType),
	}
	rsp, err := client.SupplementSkuProperty(ctx, &r)
	if err != nil {
		vars.ErrorLogger.Errorf(ctx, "SupplementSkuProperty err: %v, req: %v", err, json.MarshalToStringNoError(req))
		return &result, code.ERROR
	}
	if rsp.Common.Code == sku_business.RetCode_SUCCESS {
		return &result, code.SUCCESS
	}
	vars.ErrorLogger.Errorf(ctx, "SupplementSkuProperty  req: %v, resp: %v", json.MarshalToStringNoError(req), json.MarshalToStringNoError(rsp))
	switch rsp.Common.Code {
	case sku_business.RetCode_SKU_EXIST:
		return &result, code.ErrorSkuCodeExist
	case sku_business.RetCode_SKU_NOT_EXIST:
		return &result, code.ErrorSkuCodeNotExist
	case sku_business.RetCode_TRANSACTION_FAILED:
		return &result, code.TransactionFailed
	default:
		return &result, code.ERROR
	}
}


func SearchSkuInventory(ctx context.Context, req *args.SearchSkuInventoryArgs) (interface{}, int) {
	serverName := args.RpcServiceMicroMallSku
	conn, err := util.GetGrpcClient(ctx, serverName)
	if err != nil {
		vars.ErrorLogger.Errorf(ctx, "GetGrpcClient %q err: %v", serverName, err)
		return "", code.ERROR
	}
	//defer conn.Close()
	client := sku_business.NewSkuBusinessServiceClient(conn)
	searchReq := &sku_business.SearchSkuInventoryRequest{Keyword: req.Keyword}
	searchRsp, err := client.SearchSkuInventory(ctx, searchReq)
	if err != nil {
		vars.ErrorLogger.Errorf(ctx, "SearchSkuInventory err:%v req: %v", err, json.MarshalToStringNoError(req))
		return nil, code.ERROR
	}
	if searchRsp.Common.Code != sku_business.RetCode_SUCCESS {
		vars.ErrorLogger.Errorf(ctx, "SearchSkuInventory req: %v, rsp: %v", json.MarshalToStringNoError(req), json.MarshalToStringNoError(searchRsp))
		return nil, code.ERROR
	}
	return searchRsp.List, code.SUCCESS
}
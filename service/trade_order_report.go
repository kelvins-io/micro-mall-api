package service

import (
	"context"
	"fmt"
	"gitee.com/cristiane/micro-mall-api/model/args"
	"gitee.com/cristiane/micro-mall-api/pkg/code"
	"gitee.com/cristiane/micro-mall-api/pkg/util"
	goroute "gitee.com/cristiane/micro-mall-api/pkg/util/goroutine"
	"gitee.com/cristiane/micro-mall-api/proto/micro_mall_order_proto/order_business"
	"gitee.com/cristiane/micro-mall-api/proto/micro_mall_shop_proto/shop_business"
	"gitee.com/cristiane/micro-mall-api/proto/micro_mall_users_proto/users"
	"gitee.com/cristiane/micro-mall-api/vars"
	"golang.org/x/sync/errgroup"
	"strconv"
	"time"
)

func getOrderReport(ctx context.Context, req *args.GetOrderReportArgs) (result *args.GetOrderReportRsp, retCode int) {
	result = &args.GetOrderReportRsp{
		ReportFilePath: "暂无报告",
	}
	retCode = code.SUCCESS
	// 查找订单信息
	serverName := args.RpcServiceMicroMallOrder
	conn, err := util.GetGrpcClient(serverName)
	if err != nil {
		vars.ErrorLogger.Errorf(ctx, "GetGrpcClient %v,err: %v", serverName, err)
		retCode = code.ERROR
		return
	}
	defer conn.Close()
	orderClient := order_business.NewOrderBusinessServiceClient(conn)
	findOrderReq := order_business.FindOrderListRequest{
		ShopIdList: []int64{req.ShopId},
		UidList:    []int64{req.Uid},
		TimeMeta: &order_business.FiltrateTimeMeta{
			StartTime: req.StartTime,
			EndTime:   req.EndTime,
		},
		PageMeta: &order_business.PageMeta{
			PageNum:  int32(req.PageNum),
			PageSize: int32(req.PageSize),
		},
	}
	findOrderRsp, err := orderClient.FindOrderList(ctx, &findOrderReq)
	if err != nil {
		vars.ErrorLogger.Errorf(ctx, "FindOrderList %v,err: %v, findOrderReq: %+v", serverName, err, findOrderReq)
		retCode = code.ERROR
		return
	}
	if findOrderRsp.Common.Code != order_business.RetCode_SUCCESS {
		vars.ErrorLogger.Errorf(ctx, "FindOrderList %v,err: %v, findOrderRsp: %+v", serverName, err, findOrderRsp)
		retCode = code.ERROR
		return
	}
	if len(findOrderRsp.List) == 0 {
		return
	}
	uidList := make([]int64, 0)
	uidListSet := map[int64]struct{}{}
	shopIdList := make([]int64, 0)
	shopIdListSet := map[int64]struct{}{}
	for i := 0; i < len(findOrderRsp.List); i++ {
		if _, ok := uidListSet[findOrderRsp.List[i].Uid]; !ok {
			uidListSet[findOrderRsp.List[i].Uid] = struct{}{}
			uidList = append(uidList, findOrderRsp.List[i].Uid)
		}
		if _, ok := shopIdListSet[findOrderRsp.List[i].ShopId]; !ok {
			shopIdListSet[findOrderRsp.List[i].ShopId] = struct{}{}
			shopIdList = append(shopIdList, findOrderRsp.List[i].ShopId)
		}
	}
	taskGroup, errCtx := errgroup.WithContext(ctx)
	uidToUserInfo := map[int64]users.UserInfoMain{}
	taskGroup.Go(func() error {
		err := goroute.CheckGoroutineErr(errCtx)
		if err != nil {
			return nil
		}
		retCode = orderReportGetUserInfo(ctx, uidList, uidToUserInfo)
		if retCode == code.SUCCESS {
			return nil
		}
		return fmt.Errorf(code.GetMsg(retCode))
	})
	shopIdToShopInfo := make(map[int64]shop_business.ShopInfo)
	taskGroup.Go(func() error {
		err := goroute.CheckGoroutineErr(errCtx)
		if err != nil {
			return nil
		}
		retCode = orderReportGetShopInfo(ctx, shopIdList, shopIdToShopInfo)
		if retCode == code.SUCCESS {
			return nil
		}
		return fmt.Errorf(code.GetMsg(retCode))
	})
	err = taskGroup.Wait()
	if err != nil {
		vars.ErrorLogger.Errorf(ctx, "taskGroup wait err: %v", err)
		retCode = code.ERROR
		return
	}
	if len(uidToUserInfo) == 0 {
		return
	}
	if len(shopIdToShopInfo) == 0 {
		return
	}
	filePath, retCode := orderReportOutExcel(ctx, req.ShopId, findOrderRsp, uidToUserInfo, shopIdToShopInfo)
	if retCode != code.SUCCESS {
		return
	}
	addr := "localhost:52001"
	if vars.ServerSetting != nil && vars.ServerSetting.EndPort != 0 {
		addr = "localhost:" + strconv.Itoa(vars.ServerSetting.EndPort)
	}
	result.ReportFilePath = "http://" + addr + "/static/" + filePath

	return
}

func orderReportOutExcel(
	ctx context.Context, shopId int64, findOrderRsp *order_business.FindOrderListResponse,
	uidToUserInfo map[int64]users.UserInfoMain,
	shopIdToShopInfo map[int64]shop_business.ShopInfo) (filePath string, retCode int) {
	retCode = code.SUCCESS
	// 汇总信息
	excelSheets := make([]ExcelDataSheet, 0)
	reportSheet := ExcelDataSheet{
		Rows:  nil,
		Sheet: "订单报告",
	}
	reportSheet.Rows = append(reportSheet.Rows, []string{
		"订单号",
		"用户",
		"ip",
		"设备",
		"时间",
		"店铺",
		"总额",
		"描述",
		"状态",
		"支付状态",
	})
	for i := 0; i < len(findOrderRsp.List); i++ {
		record := findOrderRsp.List[i]
		state := ""
		payState := ""
		switch record.State {
		case order_business.OrderStateType_ORDER_EFFECTIVE:
			state = "有效"
		case order_business.OrderStateType_ORDER_LOCKED:
			state = "锁定"
		case order_business.OrderStateType_ORDER_INVALID:
			state = "无效"
		default:
			state = "未知"
		}
		switch record.PayState {
		case order_business.OrderPayStateType_PAY_CANCEL:
			payState = "取消支付"
		case order_business.OrderPayStateType_PAY_FAILED:
			payState = "支付失败"
		case order_business.OrderPayStateType_PAY_READY:
			payState = "未支付"
		case order_business.OrderPayStateType_PAY_RUN:
			payState = "支付中"
		case order_business.OrderPayStateType_PAY_SUCCESS:
			payState = "支付成功"
		default:
			payState = "未知"
		}
		row := []string{
			record.OrderCode,
			uidToUserInfo[record.Uid].Name,
			record.ClientIp,
			record.DeviceCode,
			record.CreateTime,
			shopIdToShopInfo[record.ShopId].FullName,
			record.Money,
			record.Description,
			state,
			payState,
		}
		reportSheet.Rows = append(reportSheet.Rows, row)
	}
	excelSheets = append(excelSheets, reportSheet)
	reportFile := fmt.Sprintf("%v%d-%d.xlsx", "order-report-", shopId, time.Now().Unix())
	reportFilePath := "./static/" + reportFile
	excelReq := ExcelDataArgs{
		OutFilePath: reportFilePath,
		Sheets:      excelSheets,
	}
	err := GenExcelFile(&excelReq)
	if err != nil {
		vars.ErrorLogger.Errorf(ctx, "GenExcelFile err: %v, req: %+v", err, excelReq)
		retCode = code.ERROR
		return
	}
	filePath = reportFile
	return
}

func orderReportGetUserInfo(ctx context.Context, uidList []int64, uidToUserInfo map[int64]users.UserInfoMain) (retCode int) {
	retCode = code.SUCCESS
	// 查找用户信息
	serverName := args.RpcServiceMicroMallUsers
	conn, err := util.GetGrpcClient(serverName)
	if err != nil {
		vars.ErrorLogger.Errorf(ctx, "GetGrpcClient %v,err: %v", serverName, err)
		retCode = code.ERROR
		return
	}
	defer conn.Close()
	userClient := users.NewUsersServiceClient(conn)
	userReq := users.FindUserInfoRequest{UidList: uidList}
	userRsp, err := userClient.FindUserInfo(ctx, &userReq)
	if err != nil {
		vars.ErrorLogger.Errorf(ctx, "FindUserInfo %v,err: %v, req: %+v", serverName, err, userReq)
		retCode = code.ERROR
		return
	}
	if userRsp.Common.Code != users.RetCode_SUCCESS {
		retCode = code.ERROR
		return
	}
	if len(userRsp.InfoList) == 0 {
		return
	}
	for i := 0; i < len(userRsp.InfoList); i++ {
		uidToUserInfo[userRsp.InfoList[i].Uid] = *userRsp.InfoList[i]
	}
	return
}

func orderReportGetShopInfo(ctx context.Context, shopIdList []int64, shopIdToShopInfo map[int64]shop_business.ShopInfo) (retCode int) {
	// 查找店铺信息
	retCode = code.SUCCESS
	serverName := args.RpcServiceMicroMallShop
	conn, err := util.GetGrpcClient(serverName)
	if err != nil {
		vars.ErrorLogger.Errorf(ctx, "GetGrpcClient %v,err: %v", serverName, err)
		retCode = code.ERROR
		return
	}
	defer conn.Close()
	shopClient := shop_business.NewShopBusinessServiceClient(conn)
	shopReq := shop_business.GetShopInfoRequest{ShopIds: shopIdList}
	shopRsp, err := shopClient.GetShopInfo(ctx, &shopReq)
	if err != nil {
		vars.ErrorLogger.Errorf(ctx, "GetShopInfo %v,err: %v, req: %+v", serverName, err, shopReq)
		retCode = code.ERROR
		return
	}
	if shopRsp.Common.Code != shop_business.RetCode_SUCCESS {
		retCode = code.ERROR
		return
	}
	if len(shopRsp.InfoList) == 0 {
		return
	}
	for i := 0; i < len(shopRsp.InfoList); i++ {
		shopIdToShopInfo[shopRsp.InfoList[i].ShopId] = *shopRsp.InfoList[i]
	}
	return
}

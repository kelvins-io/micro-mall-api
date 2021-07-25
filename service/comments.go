package service

import (
	"context"
	"gitee.com/cristiane/micro-mall-api/model/args"
	"gitee.com/cristiane/micro-mall-api/pkg/code"
	"gitee.com/cristiane/micro-mall-api/pkg/util"
	"gitee.com/cristiane/micro-mall-api/proto/micro_mall_comments_proto/comments_business"
	"gitee.com/cristiane/micro-mall-api/vars"
)

func CreateOrderComments(ctx context.Context, req *args.CreateOrderCommentsArgs) (retCode int) {
	retCode = code.SUCCESS
	serverName := args.RpcServiceMicroMallComments
	conn, err := util.GetGrpcClient(serverName)
	if err != nil {
		vars.ErrorLogger.Errorf(ctx, "GetGrpcClient %q,err: %v", serverName, err)
		retCode = code.ERROR
		return
	}
	defer conn.Close()
	client := comments_business.NewCommentsBusinessServiceClient(conn)
	commentsOrderReq := comments_business.CommentsOrderRequest{
		Uid: req.Uid,
		OrderInfo: &comments_business.OrderCommentsInfo{
			ShopId:    req.OrderCommentsInfo.ShopId,
			OrderCode: req.OrderCommentsInfo.OrderCode,
			StarLevel: comments_business.StarLevel(req.OrderCommentsInfo.Star),
			Content:   req.OrderCommentsInfo.Content,
			ImgList:   req.OrderCommentsInfo.ImgList,
			CommentId: req.OrderCommentsInfo.CommentId,
		},
		LogisticsInfo: &comments_business.LogisticsCommentsInfo{
			LogisticsCode:        req.LogisticsCommentsInfo.LogisticsCode,
			FedexPack:            comments_business.StarLevel(req.LogisticsCommentsInfo.FedexPack),
			FedexLabel:           req.LogisticsCommentsInfo.FedexPackLabel,
			DeliverySpeed:        comments_business.StarLevel(req.LogisticsCommentsInfo.DeliverySpeed),
			DeliverySpeedLabel:   req.LogisticsCommentsInfo.DeliverySpeedLabel,
			DeliveryService:      comments_business.StarLevel(req.LogisticsCommentsInfo.DeliveryService),
			DeliveryServiceLabel: req.LogisticsCommentsInfo.DeliveryServiceLabel,
			Comment:              req.LogisticsCommentsInfo.Comment,
		},
		Anonymity: req.Anonymity,
	}
	commentsOrderRsp, err := client.CommentsOrder(ctx, &commentsOrderReq)
	if err != nil {
		vars.ErrorLogger.Errorf(ctx, "CommentsOrder err: %v,req : %+v", err, *req)
		retCode = code.ERROR
		return
	}
	if commentsOrderRsp.Common.Code == comments_business.RetCode_SUCCESS {
		return
	}
	vars.ErrorLogger.Errorf(ctx, "CommentsOrder req: %+v, commentsOrderRsp : %+v", *req, commentsOrderRsp)
	switch commentsOrderRsp.Common.Code {
	case comments_business.RetCode_USER_ORDER_STATE_INVALID:
		retCode = code.OrderStateInvalid
	case comments_business.RetCode_USER_ORDER_NOT_EXIST:
		retCode = code.UserOrderNotExist
	case comments_business.RetCode_COMMENT_EXIST:
		retCode = code.CommentsExist
	case comments_business.RetCode_COMMENT_NOT_EXIST:
		retCode = code.CommentsNotExist
	case comments_business.RetCode_TRANSACTION_FAILED:
		retCode = code.TransactionFailed
	default:
		retCode = code.ERROR
	}
	return
}

func GetOrderCommentsList(ctx context.Context, req *args.GetShopCommentsListArgs) (result []args.OrderCommentsInfo, retCode int) {
	retCode = code.SUCCESS
	serverName := args.RpcServiceMicroMallComments
	conn, err := util.GetGrpcClient(serverName)
	if err != nil {
		vars.ErrorLogger.Errorf(ctx, "GetGrpcClient %q,err: %v", serverName, err)
		retCode = code.ERROR
		return
	}
	defer conn.Close()
	client := comments_business.NewCommentsBusinessServiceClient(conn)
	commentsReq := &comments_business.FindShopCommentsRequest{
		ShopId: req.ShopId,
		Uid:    req.Uid,
	}
	commentsRsp, err := client.FindShopComments(ctx, commentsReq)
	if err != nil {
		vars.ErrorLogger.Errorf(ctx, "FindShopComments err: %v,req : %+v", err, commentsReq)
		retCode = code.ERROR
		return
	}
	if commentsRsp.Common.Code != comments_business.RetCode_SUCCESS {
		vars.ErrorLogger.Errorf(ctx, "FindShopComments req: %+v, resp : %+v", *req, commentsRsp)
		retCode = code.ERROR
		return
	}

	result = make([]args.OrderCommentsInfo, len(commentsRsp.CommentsList))
	for i := 0; i < len(commentsRsp.CommentsList); i++ {
		row := commentsRsp.CommentsList[i]
		info := args.OrderCommentsInfo{
			ShopId:    row.ShopId,
			OrderCode: row.OrderCode,
			Star:      int8(row.StarLevel),
			Content:   row.Content,
			ImgList:   row.ImgList,
			CommentId: row.CommentId,
		}
		result[i] = info
	}
	return
}

func ModifyCommentsTags(ctx context.Context, req *args.ModifyCommentsTagsArgs) (retCode int) {
	retCode = code.SUCCESS
	serverName := args.RpcServiceMicroMallComments
	conn, err := util.GetGrpcClient(serverName)
	if err != nil {
		vars.ErrorLogger.Errorf(ctx, "GetGrpcClient %q err: %v", serverName, err)
		retCode = code.ERROR
		return
	}
	defer conn.Close()
	client := comments_business.NewCommentsBusinessServiceClient(conn)
	commentsReq := &comments_business.ModifyCommentsTagsRequest{
		OpType: comments_business.OperationType(req.OperationType),
		Tag: &comments_business.CommentsTags{
			TagCode:              req.CommentsTags.TagCode,
			ClassificationMajor:  req.CommentsTags.ClassificationMajor,
			ClassificationMedium: req.CommentsTags.ClassificationMedium,
			ClassificationMinor:  req.CommentsTags.ClassificationMinor,
			Content:              req.CommentsTags.Content,
			State:                true,
		},
	}
	commentsRsp, err := client.ModifyCommentsTags(ctx, commentsReq)
	if err != nil {
		vars.ErrorLogger.Errorf(ctx, "ModifyCommentsTags err: %v,req : %+v", err, commentsReq)
		retCode = code.ERROR
		return
	}
	if commentsRsp.Common.Code == comments_business.RetCode_SUCCESS {
		return
	}

	vars.ErrorLogger.Errorf(ctx, "ModifyCommentsTags req: %+v, rsp : %+v", *req, commentsRsp)
	switch commentsRsp.Common.Code {
	case comments_business.RetCode_COMMENT_TAG_NOT_EXIST:
		retCode = code.CommentsTagNotExist
	case comments_business.RetCode_COMMENT_TAG_EXIST:
		retCode = code.CommentsTagExist
	default:
		retCode = code.ERROR
	}
	return
}

func GetCommentsTagsList(ctx context.Context, req *args.GetCommentsTagsListArgs) (result []args.CommentsTags, retCode int) {
	retCode = code.SUCCESS
	serverName := args.RpcServiceMicroMallComments
	conn, err := util.GetGrpcClient(serverName)
	if err != nil {
		vars.ErrorLogger.Errorf(ctx, "GetGrpcClient %q err: %v", serverName, err)
		retCode = code.ERROR
		return
	}
	defer conn.Close()
	client := comments_business.NewCommentsBusinessServiceClient(conn)
	commentsReq := &comments_business.FindCommentsTagRequest{
		TagCode:              req.TagCode,
		ClassificationMajor:  req.ClassificationMajor,
		ClassificationMedium: req.ClassificationMedium,
	}
	commentsRsp, err := client.FindCommentsTags(ctx, commentsReq)
	if err != nil {
		vars.ErrorLogger.Errorf(ctx, "FindCommentsTags err: %v,req : %+v", err, *req)
		retCode = code.ERROR
		return
	}
	if commentsRsp.Common.Code != comments_business.RetCode_SUCCESS {
		vars.ErrorLogger.Errorf(ctx, "FindCommentsTags req: %+v,resp : %+v", *req, commentsRsp)
		switch commentsRsp.Common.Code {
		case comments_business.RetCode_COMMENT_TAG_NOT_EXIST:
			retCode = code.CommentsTagNotExist
		default:
			retCode = code.ERROR
		}
		return
	}

	result = make([]args.CommentsTags, len(commentsRsp.Tags))
	for i := 0; i < len(commentsRsp.Tags); i++ {
		row := commentsRsp.Tags[i]
		tags := args.CommentsTags{
			TagCode:              row.TagCode,
			ClassificationMajor:  row.ClassificationMajor,
			ClassificationMedium: row.ClassificationMedium,
			ClassificationMinor:  row.ClassificationMinor,
			Content:              row.Content,
		}
		result[i] = tags
	}
	return
}

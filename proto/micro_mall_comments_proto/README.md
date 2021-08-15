# micro-mall-comments-proto

#### 介绍
评论服务proto仓库

#### 软件架构
micro-mall-comments服务的接口定义

#### 使用说明
接口说明
```protobuf
    // 订单评价
    rpc CommentsOrder(CommentsOrderRequest) returns (CommentsOrderResponse) {
        option (google.api.http) = {
            post: "/v1/comments/order"
            body:"*"
        };
    }
    // 获取店铺评论列表
    rpc FindShopComments(FindShopCommentsRequest) returns (FindShopCommentsResponse) {
        option (google.api.http) = {
            get: "/v1/comments/list"
        };
    }
    // 获取评论标签
    rpc FindCommentsTags(FindCommentsTagRequest) returns (FindCommentsTagResponse) {
        option (google.api.http) = {
            get: "/v1/comments/tags/list"
        };
    }
    // 修改评论标签
    rpc ModifyCommentsTags(ModifyCommentsTagsRequest) returns (ModifyCommentsTagsResponse) {
        option (google.api.http) = {
            post: "/v1/comments/tags/modify"
            body:"*"
        };
    }
```

#### 参与贡献

1.  Fork 本仓库
2.  新建 Feat_xxx 分支
3.  提交代码
4.  新建 Pull Request


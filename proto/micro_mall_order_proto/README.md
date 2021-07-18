# micro-mall-order-proto

#### 介绍
微商城-订单系统proto

#### 软件架构
订单系统proto


#### 使用说明
接口定义
```protobuf
    // 生成唯一订单事务号
    rpc GenOrderTxCode(GenOrderTxCodeRequest) returns (GenOrderTxCodeResponse) {
        option (google.api.http) = {
            get: "/v1/order/code"
        };
    }
    // 检查外部订单号是否存在
    rpc CheckOrderExist(CheckOrderExistRequest) returns (CheckOrderExistResponse) {
        option (google.api.http) = {
            get: "/v1/order/code/exist"
        };
    }
    // 创建订单
    rpc CreateOrder(CreateOrderRequest) returns (CreateOrderResponse) {
        option (google.api.http) = {
            post: "/v1/order/create"
            body:"*"
        };
    }
    // 获取订单详情
    rpc GetOrderDetail(GetOrderDetailRequest) returns (GetOrderDetailResponse) {
        option (google.api.http) = {
            get: "/v1/order/detail"
        };
    }
    // 获取订单商品
    rpc GetOrderSku(GetOrderSkuRequest) returns (GetOrderSkuResponse) {
        option (google.api.http) = {
            get: "/v1/order/sku"
        };
    }
    // 更新订单状态
    rpc UpdateOrderState(UpdateOrderStateRequest) returns (UpdateOrderStateResponse) {
        option (google.api.http) = {
            post: "/v1/order/state"
            body:"*"
        };
    }
    // 订单支付通知
    rpc OrderTradeNotice(OrderTradeNoticeRequest) returns (OrderTradeNoticeResponse) {
        option (google.api.http) = {
            post: "/v1/order/notice"
            body:"*"
        };
    }
    // 订单状态检查
    rpc CheckOrderState(CheckOrderStateRequest) returns (CheckOrderStateResponse) {
        option (google.api.http) = {
            post: "/v1/order/check"
            body:"*"
        };
    }
    // 获取订单
    rpc FindOrderList(FindOrderListRequest) returns (FindOrderListResponse) {
        option (google.api.http) = {
            get: "/v1/order/list"
        };
    }
    // 店铺订单存在检查
    rpc InspectShopOrder(InspectShopOrderRequest) returns (InspectShopOrderResponse) {
        option (google.api.http) = {
            get: "/v1/order/shop/inspect"
        };
    }
```

#### 参与贡献

1.  Fork 本仓库
2.  新建 Feat_xxx 分支
3.  提交代码
4.  新建 Pull Request


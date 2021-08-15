# micro-mall-trolley-proto

#### 介绍
微商城--购物车proto

#### 软件架构
软件架构说明


#### 使用说明
接口定义
```protobuf
    // 添加商品到购物车
    rpc JoinSku(JoinSkuRequest) returns (JoinSkuResponse) {
        option (google.api.http) = {
            put: "/v1/trolley/sku/join"
            body:"*"
        };
    }
    // 从购物车移除商品
    rpc RemoveSku(RemoveSkuRequest) returns (RemoveSkuResponse) {
        option (google.api.http) = {
            delete: "/v1/trolley/sku/remove"
        };
    }
    // 获取用户购物车中的商品
    rpc GetUserTrolleyList(GetUserTrolleyListRequest) returns (GetUserTrolleyListResponse) {
        option (google.api.http) = {
            get: "/v1/trolley/sku/list"
        };
    }
```

#### 参与贡献

1.  Fork 本仓库
2.  新建 Feat_xxx 分支
3.  提交代码
4.  新建 Pull Request

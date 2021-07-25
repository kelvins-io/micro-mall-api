# micro-mall-shop-proto

#### 介绍
微商城-店铺系统proto

#### 软件架构
软件架构说明


#### 安装教程

1. sh build.sh

#### 使用说明
接口定义
```protobuf
    // 提交店铺申请材料
    rpc ShopApply(ShopApplyRequest) returns (ShopApplyResponse) {
        option (google.api.http) = {
            post: "/v1/shop_business/shop/apply"
            body:"*"
        };
    }
    // 店铺质押保证金
    rpc ShopPledge(ShopPledgeRequest) returns (ShopPledgeResponse) {
        option (google.api.http) = {
            put: "/v1/shop_business/shop/pledge"
            body:"*"
        };
    }
    // 获取店铺材料
    rpc GetShopMaterial(GetShopMaterialRequest) returns (GetShopMaterialResponse) {
        option (google.api.http) = {
            get: "/v1/shop_business/shop/material"
        };
    }
    // 获取店铺数据
    rpc GetShopInfo(GetShopInfoRequest) returns (GetShopInfoResponse) {
        option (google.api.http) = {
            get: "/v1/shop_business/shop/info"
        };
    }
    // 搜索同步店铺数据
    rpc SearchSyncShop(SearchSyncShopRequest) returns (SearchSyncShopResponse) {
        option (google.api.http) = {
            get: "/v1/shop_business/shop/search/sync"
        };
    }
    // 搜索店铺
    rpc SearchShop(SearchShopRequest) returns (SearchShopResponse) {
        option (google.api.http) = {
            get: "/v1/shop_business/shop/search"
        };
    }
    // 获取店铺主要信息
    rpc GetShopMajorInfo(GetShopMajorInfoRequest) returns (GetShopMajorInfoResponse) {
        option (google.api.http) = {
            get: "/v1/shop_business/shop/info/major"
        };
    }
```

#### 参与贡献

1.  Fork 本仓库
2.  新建 Feat_xxx 分支
3.  提交代码
4.  新建 Pull Request


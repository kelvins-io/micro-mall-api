# micro-mall-sku-proto

#### 介绍
微商城-商品系统proto

#### 软件架构
软件架构说明


#### 使用说明
接口定义
```protobuf
    // 上架商品
    rpc PutAwaySku(PutAwaySkuRequest) returns (PutAwaySkuResponse) {
        option (google.api.http) = {
            post: "/v1/sku/inventory/put_away"
            body:"*"
        };
    }
    // 获取店铺sku列表
    rpc GetSkuList(GetSkuListRequest) returns (GetSkuListResponse) {
        option (google.api.http) = {
            get: "/v1/sku/inventory/list"
        };
    }
    // 补充sku商品属性
    rpc SupplementSkuProperty(SupplementSkuPropertyRequest) returns (SupplementSkuPropertyResponse) {
        option (google.api.http) = {
            put: "/v1/sku/property/supplement"
            body:"*"
        };
    }
    // 扣减库存
    rpc DeductInventory(DeductInventoryRequest) returns (DeductInventoryResponse) {
        option (google.api.http) = {
            put: "/v1/sku/inventory/deduct"
            body:"*"
        };
    }
    // 恢复库存
    rpc RestoreInventory(RestoreInventoryRequest) returns (RestoreInventoryResponse) {
        option (google.api.http) = {
            put: "/v1/sku/inventory/restore"
            body:"*"
        };
    }
    // 按策略筛选商品价格版本
    rpc FiltrateSkuPriceVersion(FiltrateSkuPriceVersionRequest) returns (FiltrateSkuPriceVersionResponse) {
        option (google.api.http) = {
            post: "/v1/sku/price/filtrate"
            body:"*"
        };
    }
    // 商品库存搜索同步数据(请在业务不繁忙时调用)
    rpc SearchSyncSkuInventory(SearchSyncSkuInventoryRequest) returns (SearchSyncSkuInventoryResponse) {
        option (google.api.http) = {
            post: "/v1/sku/inventory/search/sync"
            body:"*"
        };
    }
    // 商品搜索
    rpc SearchSkuInventory(SearchSkuInventoryRequest) returns (SearchSkuInventoryResponse) {
        option (google.api.http) = {
            post: "/v1/sku/inventory/search"
            body:"*"
        };
    }
    // 确认库存
    rpc ConfirmSkuInventory(ConfirmSkuInventoryRequest) returns (ConfirmSkuInventoryResponse) {
        option (google.api.http) = {
            post: "/v1/sku/inventory/confirm"
            body:"*"
        };
    }
```

#### 参与贡献

1.  Fork 本仓库
2.  新建 Feat_xxx 分支
3.  提交代码
4.  新建 Pull Request

# micro-mall-pay-proto

#### 介绍
微商城-支付系统proto

#### 软件架构
软件架构说明

#### 使用说明
接口定义
```protobuf
    // 统一收单支付
    rpc TradePay(TradePayRequest) returns (TradePayResponse) {
        option (google.api.http) = {
            post: "/v1/trade/pay"
            body:"*"
        };
    }
    // 创建账户
    rpc CreateAccount(CreateAccountRequest) returns (CreateAccountResponse) {
        option (google.api.http) = {
            post: "/v1/account/init"
            body:"*"
        };
    }
    // 获取账户
    rpc FindAccount(FindAccountRequest) returns (FindAccountResponse) {
        option (google.api.http) = {
            get: "/v1/account"
        };
    }
    // 账户充值
    rpc AccountCharge(AccountChargeRequest) returns (AccountChargeResponse) {
        option (google.api.http) = {
            post: "/v1/account/charge"
            body:"*"
        };
    }
    // 获取交易唯一ID
    rpc GetTradeUUID(GetTradeUUIDRequest) returns(GetTradeUUIDResponse) {
        option (google.api.http) = {
            get: "/v1/trade/uuid"
        };
    }
```

#### 参与贡献

1.  Fork 本仓库
2.  新建 Feat_xxx 分支
3.  提交代码
4.  新建 Pull Request


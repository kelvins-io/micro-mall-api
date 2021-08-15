# micro-mall-logistics-proto

#### 介绍
微商城-物流系统proto

#### 软件架构
物流服务的接口定义

#### 使用说明
接口列表
```protobuf
    // 申请物流
    rpc ApplyLogistics(ApplyLogisticsRequest) returns (ApplyLogisticsResponse) {
        option (google.api.http) = {
            post: "/v1/logistics/apply"
            body:"*"
        };
    }
    // 查询物流记录
    rpc QueryRecord(QueryRecordRequest) returns (QueryRecordResponse) {
        option (google.api.http) = {
            get: "/v1/logistics/query"
        };
    }
    // 更新物流状态
    rpc UpdateState(UpdateStateRequest) returns (UpdateStateResponse) {
        option (google.api.http) = {
            put: "/v1/logistics/update"
            body:"*"
        };
    }
```

#### 参与贡献

1.  Fork 本仓库
2.  新建 Feat_xxx 分支
3.  提交代码
4.  新建 Pull Request
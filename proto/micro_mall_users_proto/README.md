# micro-mall-users-proto

#### 介绍
微商城-用户系统proto

#### 软件架构
软件架构说明

#### 使用说明
接口定义
```protobuf
    // 获取用户信息
    rpc GetUserInfo(GetUserInfoRequest) returns (GetUserInfoResponse) {
        option (google.api.http) = {
            get: "/v1/user/info"
        };
    }
    // 获取用户信息-手机号
    rpc GetUserInfoByPhone(GetUserInfoByPhoneRequest) returns (GetUserInfoByPhoneResponse) {
        option (google.api.http) = {
            get: "/v1/user/info/phone"
        };
    }
    // 检查用户是否注册-手机号
    rpc CheckUserByPhone(CheckUserByPhoneRequest) returns(CheckUserByPhoneResponse) {
        option (google.api.http) = {
            get: "/v1/user/exist/phone"
        };
    }
    // 根据邀请码获取邀请人
    rpc GetUserInfoByInviteCode(GetUserByInviteCodeRequest) returns (GetUserByInviteCodeResponse) {
        option (google.api.http) = {
            get: "/v1/user/info/invite_code"
        };
    }
    // 注册用户
    rpc Register(RegisterRequest) returns (RegisterResponse) {
        option (google.api.http) = {
            post: "/v1/user/register"
            body: "*"
        };
    }
    // 登录
    rpc LoginUser(LoginUserRequest) returns (LoginUserResponse) {
        option (google.api.http) = {
            put: "/v1/user/login"
            body: "*"
        };
    }
    // 检查用户身份
    rpc CheckUserIdentity(CheckUserIdentityRequest) returns (CheckUserIdentityResponse) {
        option (google.api.http) = {
            get: "/v1/user/identity"
        };
    }
    // 重置登录密码
    rpc PasswordReset(PasswordResetRequest) returns (PasswordResetResponse) {
        option (google.api.http) = {
            put: "/v1/user/password/reset"
            body: "*"
        };
    }
    // 更新用户登录态
    rpc UpdateUserLoginState(UpdateUserLoginStateRequest) returns (UpdateUserLoginStateResponse) {
        option (google.api.http) = {
            put: "/v1/user/state/login"
            body: "*"
        };
    }
    // 修改用户收货信息
    rpc ModifyUserDeliveryInfo(ModifyUserDeliveryInfoRequest) returns (ModifyUserDeliveryInfoResponse) {
        option (google.api.http) = {
            put: "/v1/user/info/delivery"
            body: "*"
        };
    }
    // 获取用户收货信息
    rpc GetUserDeliveryInfo(GetUserDeliveryInfoRequest) returns (GetUserDeliveryInfoResponse) {
        option (google.api.http) = {
            get: "/v1/user/info/delivery"
        };
    }
    // 查找用户信息
    rpc FindUserInfo(FindUserInfoRequest) returns (FindUserInfoResponse) {
        option (google.api.http) = {
            get: "/v1/user/info/find"
        };
    }
    // 用户账户充值
    rpc UserAccountCharge(UserAccountChargeRequest) returns (UserAccountChargeResponse) {
        option (google.api.http) = {
            put: "/v1/user/account/charge"
            body: "*"
        };
    }
    // 检查用户收货地址
    rpc CheckUserDeliveryInfo(CheckUserDeliveryInfoRequest) returns (CheckUserDeliveryInfoResponse) {
        option (google.api.http) = {
            get: "/v1/user/info/check"
        };
    }
    // 检查用户
    rpc CheckUserState(CheckUserStateRequest) returns (CheckUserStateResponse) {
        option (google.api.http) = {
            get: "/v1/user/state/check"
        };
    }
    // 获取用户accountId
    rpc GetUserAccountId(GetUserAccountIdRequest) returns (GetUserAccountIdResponse) {
        option (google.api.http) = {
            get: "/v1/user/info/account_id"
        };
    }
    // 列表所有用户
    rpc ListUserInfo(ListUserInfoRequest) returns (ListUserInfoResponse) {
        option (google.api.http) = {
            get: "/v1/user/info/list"
        };
    }
```

#### 参与贡献

1.  Fork 本仓库
2.  新建 Feat_xxx 分支
3.  提交代码
4.  新建 Pull Request


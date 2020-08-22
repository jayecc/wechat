> [小程序服务端官方文档](https://developers.weixin.qq.com/miniprogram/dev/api-backend/) 

## 安装

```sh

go get -u -v github.com/jayecc/wechat

```

## 目录

- [登录](#登录)
  - [code2Session](#code2Session)
- [用户信息](#用户信息)
  - [getPaidUnionId](#getPaidUnionId) 
- [接口调用凭证](#接口调用凭证)
  - [getAccessToken](#getAccessToken)
  
---

## 登陆

#### [code2Session](https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/login/auth.code2Session.html)

```go
import "github.com/jayecc/wechat"

req := new(Code2SessionRequest)
resp := new(Code2SessionResponse)

req = &Code2SessionRequest{
    AppID:  "appid",
    Secret: "secret",
    JsCode: "js_code",
}

if err := Code2Session(req, resp); err != nil {
    t.Fatalf("%v", err)
}

```

---

## 用户信息

#### [getPaidUnionId](https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/user-info/auth.getPaidUnionId.html)
> 调用时两种方式任选其一

```go
import "github.com/jayecc/wechat"

req := new(GetPaidUnionIdRequest)
resp := new(GetPaidUnionIdResponse)

req = &GetPaidUnionIdRequest{
    AccessToken:   "access_token",
    Openid:        "open_id",
    TransactionID: "transaction_id",
    MchID:         "mch_id",
    OutTradeNO:    "out_trade_no",
}

if err := GetPaidUnionId(req, resp); err != nil {
    t.Fatalf("%v", err)
}

```
 

 ---
 
 ## 调用接口凭证
 
 #### [getAccessToken](https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/access-token/auth.getAccessToken.html)
 > 调用时请注意频率，做好妥善缓存
 
 ```go
 import "github.com/jayecc/wechat"
 
req := new(GetAccessTokenRequest)
resp := new(GetAccessTokenResponse)

req = &GetAccessTokenRequest{
    AppID:  "wxaec93043ddef499d",
    Secret: "a810dda42806e814fbcf1237ee824bba",
}

if err := GetAccessToken(req, resp); err != nil {
    t.Fatalf("%v", err)
}
 ```
  
  
  
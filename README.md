> [小程序服务端官方文档](https://developers.weixin.qq.com/miniprogram/dev/api-backend/) 

## 安装

```sh

go get -u -v github.com/jayecc/wechat

```

## 目录

- [登陆](#登陆)
  - [auth.code2Session](#authcode2Session)
- [用户信息](#用户信息)
  - [auth.getPaidUnionId](#auth.getPaidUnionId) 
- [接口调用凭证](#接口调用凭证)
  - [auth.getAccessToken](#auth.getAccessToken)
- [数据分析](#数据分析)
  - [访问留存](#访问留存)
    - [analysis.getDailyRetain](#analysis.getDailyRetain)
    - [analysis.getMonthlyRetain](#analysis.getMonthlyRetain)
    - [analysis.getWeeklyRetain](#analysis.getWeeklyRetain)
  - [analysis.getDailySummary](#analysis.getDailySummary)
  - [访问趋势](#访问趋势)
    - [analysis.getDailyVisitTrend](#analysis.getDailyVisitTrend)
    - [analysis.getMonthlyVisitTrend](#analysis.getMonthlyVisitTrend)
    - [analysis.getWeeklyVisitTrend](#analysis.getWeeklyVisitTrend)
  - [analysis.getUserPortrait](#analysis.getUserPortrait)
  - [analysis.getVisitDistribution](#analysis.getVisitDistribution)
  - [analysis.getVisitPage](#analysis.getVisitPage)
- [客服消息](#客服消息)
  - [customerServiceMessage.getTempMedia](#customerServiceMessage.getTempMedia)
  - [customerServiceMessage.send](#customerServiceMessage.send)
  - [customerServiceMessage.setTyping](#customerServiceMessage.setTyping)
  - [customerServiceMessage.uploadTempMedia](#customerServiceMessage.uploadTempMedia)
- [统一消息](#统一消息)
  - [uniformMessage.send](#uniformMessage.send)
- [动态消息](#动态消息)
  - [updatableMessage.createActivityId](#updatableMessage.createActivityId)
  - [updatableMessage.setUpdatableMsg](#updatableMessage.setUpdatableMsg)
- [插件管理](#插件管理)
  - [pluginManager.applyPlugin](#pluginManager.applyPlugin)
  - [pluginManager.getPluginDevApplyList](#pluginManager.getPluginDevApplyList)
  - [pluginManager.getPluginList](#pluginManager.getPluginList)
  - [pluginManager.setDevPluginApplyStatus](#pluginManager.setDevPluginApplyStatus)
  - [pluginManager.unbindPlugin](#pluginManager.unbindPlugin)
---

## 登陆

#### [authcode2Session](https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/login/auth.code2Session.html)

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

#### [auth.getPaidUnionId](https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/user-info/auth.getPaidUnionId.html)
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
 
 ## 接口调用凭证
 
 #### [auth.getAccessToken](https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/access-token/auth.getAccessToken.html)
 > 调用时请注意频率，做好妥善缓存
 
 ```go
 import "github.com/jayecc/wechat"
 
req := new(GetAccessTokenRequest)
resp := new(GetAccessTokenResponse)

req = &GetAccessTokenRequest{
    AppID:  "xxx",
    Secret: "xxx",
}

if err := GetAccessToken(req, resp); err != nil {
    t.Fatalf("%v", err)
}
 ```
  

---
 
## 数据分析

### 访问留存
 
#### [analysis.getDailyRetain](https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/data-analysis/visit-retain/analysis.getDailyRetain.html)
 
```go
import "github.com/jayecc/wechat"
 
req := new(GetDailyRetainRequest)
resp := new(GetDailyRetainResponse)

token := "xxxx"

req = &GetDailyRetainRequest{
    BeginDate: "20170313",
    EndDate:   "20170313",
}

if err := GetDailyRetain(token, req, resp); err != nil {
    t.Fatalf("%v", err)
}
```
 
#### [analysis.getMonthlyRetain](https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/data-analysis/visit-retain/analysis.getMonthlyRetain.html)
 
```go
import "github.com/jayecc/wechat"
 
req := new(GetMonthlyRetainRequest)
resp := new(GetMonthlyRetainResponse)

token := "xxxx"

req = &GetMonthlyRetainRequest{
    BeginDate: "20170201",
    EndDate:   "20170228",
}

if err := GetMonthlyRetain(token, req, resp); err != nil {
    t.Fatalf("%v", err)
}
```

#### [analysis.getWeeklyRetain](https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/data-analysis/visit-retain/analysis.getWeeklyRetain.html)
 
```go
import "github.com/jayecc/wechat"
 
req := new(GetWeeklyRetainRequest)
resp := new(GetWeeklyRetainResponse)

token := "xxxx"

req = &GetWeeklyRetainRequest{
    BeginDate: "20170306",
    EndDate:   "20170312",
}

if err := GetWeeklyRetain(token, req, resp); err != nil {
    t.Fatalf("%v", err)
}
```

  
  
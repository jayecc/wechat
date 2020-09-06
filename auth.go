package wechat

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/pkg/errors"
)

// GrantType 类型
type GrantType string

const (
	// GrantTypeAuthorizationCode 授权码
	GrantTypeAuthorizationCode GrantType = "authorization_code"
	// GrantTypeClientCredential 客户端凭证
	GrantTypeClientCredential GrantType = "client_credential"
)

// Code2SessionRequest 登录凭证校验-请求
type Code2SessionRequest struct {
	AppID     string    `json:"appid"`      //小程序 appId
	Secret    string    `json:"secret"`     //小程序 appSecret
	JsCode    string    `json:"js_code"`    //登录时获取的 code
	GrantType GrantType `json:"grant_type"` //授权类型，此处只需填写 authorization_code
}

// Code2SessionResponse 登录凭证校验-响应
type Code2SessionResponse struct {
	OpenID     string `json:"open_id"`     //用户唯一标识
	SessionKey string `json:"session_key"` //会话密钥
	UnionID    string `json:"unionid"`     //用户在开放平台的唯一标识符，在满足 UnionID 下发条件的情况下会返回，详见 UnionID 机制说明。
}

// Code2Session 登录凭证校验。通过 wx.login 接口获得临时登录凭证 code 后传到开发者服务器调用此接口完成登录流程。更多使用方法详见 小程序登录。
// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/login/auth.code2Session.html
func Code2Session(req *Code2SessionRequest, resp *Code2SessionResponse) error {

	req.GrantType = GrantTypeAuthorizationCode

	if err := validation.ValidateStruct(req,
		validation.Field(&req.AppID, validation.Required),
		validation.Field(&req.Secret, validation.Required),
		validation.Field(&req.JsCode, validation.Required),
	); err != nil {
		return errors.Wrap(err, "request param error")
	}

	URL := "https://api.weixin.qq.com/sns/jscode2session"

	return httpGetJSON(DefaultHTTPClient, URL, req, resp)
}

// GetPaidUnionIDRequest 获取UnionID-请求
type GetPaidUnionIDRequest struct {
	AccessToken   string `json:"access_token"`   //接口调用凭证
	Openid        string `json:"openid"`         //支付用户唯一标识
	TransactionID string `json:"transaction_id"` //微信支付订单号
	MchID         string `json:"mch_id"`         //微信支付分配的商户号，和商户订单号配合使用
	OutTradeNO    string `json:"out_trade_no"`   //微信支付商户订单号，和商户号配合使用
}

// GetPaidUnionIDResponse 获取UnionID-响应
type GetPaidUnionIDResponse struct {
	UnionID string `json:"unionid"` //用户唯一标识，调用成功后返回
}

// GetPaidUnionID 用户支付完成后，获取该用户的 UnionId，无需用户授权
// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/user-info/auth.getPaidUnionId.html
func GetPaidUnionID(req *GetPaidUnionIDRequest, resp *GetPaidUnionIDResponse) error {

	if err := validation.ValidateStruct(req,
		validation.Field(&req.AccessToken, validation.Required),
		validation.Field(&req.Openid, validation.Required),
	); err != nil {
		return errors.Wrap(err, "request param error")
	}

	URL := "https://api.weixin.qq.com/wxa/getpaidunionid"

	return httpGetJSON(DefaultHTTPClient, URL, req, resp)
}

// GetAccessTokenRequest 获取凭证-请求
type GetAccessTokenRequest struct {
	GrantType GrantType `json:"grant_type"` //填写 client_credential
	AppID     string    `json:"appid"`      //小程序唯一凭证，即 AppID，可在「微信公众平台 - 设置 - 开发设置」页中获得。（需要已经成为开发者，且帐号没有异常状态）
	Secret    string    `json:"secret"`     //小程序唯一凭证密钥，即 AppSecret，获取方式同 appid
}

// GetAccessTokenResponse 获取凭证-响应
type GetAccessTokenResponse struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   int    `json:"expires_in"`
}

// GetAccessToken 获取小程序全局唯一后台接口调用凭据（access_token）。调用绝大多数后台接口时都需使用 access_token，开发者需要进行妥善保存。
// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/access-token/auth.getAccessToken.html
func GetAccessToken(req *GetAccessTokenRequest, resp *GetAccessTokenResponse) error {

	req.GrantType = GrantTypeClientCredential

	if err := validation.ValidateStruct(req,
		validation.Field(&req.AppID, validation.Required),
		validation.Field(&req.Secret, validation.Required),
	); err != nil {
		return errors.Wrap(err, "request param error")
	}

	URL := "https://api.weixin.qq.com/cgi-bin/token"

	return httpGetJSON(DefaultHTTPClient, URL, req, resp)
}

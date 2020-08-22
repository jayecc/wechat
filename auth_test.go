package wechat

import "testing"

func TestCode2Session(t *testing.T) {

	req := new(Code2SessionRequest)
	resp := new(Code2SessionResponse)

	req = &Code2SessionRequest{
		AppID:  "111",
		Secret: "q",
		JsCode: "w",
	}

	if err := Code2Session(req, resp); err != nil {
		t.Fatalf("%v", err)
	}

	t.Log(resp)
	t.Log("ok")
}

func TestGetPaidUnionId(t *testing.T) {

	req := new(GetPaidUnionIdRequest)
	resp := new(GetPaidUnionIdResponse)

	req = &GetPaidUnionIdRequest{
		AccessToken:   "xxxx",
		Openid:        "22",
		TransactionID: "",
		MchID:         "",
		OutTradeNO:    "",
	}

	if err := GetPaidUnionId(req, resp); err != nil {
		t.Fatalf("%v", err)
	}

	t.Log(resp)
	t.Log("ok")
}

func TestGetAccessToken(t *testing.T) {

	req := new(GetAccessTokenRequest)
	resp := new(GetAccessTokenResponse)

	req = &GetAccessTokenRequest{
		AppID:  "xxxx",
		Secret: "xxxx",
	}

	if err := GetAccessToken(req, resp); err != nil {
		t.Fatalf("%v", err)
	}

	t.Log(resp)
	t.Log("ok")
}

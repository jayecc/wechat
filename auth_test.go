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
		AccessToken:   "36_MqakaCLpfg1pTbdaoBeuaQV_yrM9A5j5kZAUQGhcFef4e3o6DPD0b7k5fMtnYuE-dky2b7fQwad91Ksk5rvOv6x6eSWdqnj7zxgdW-JEJaZDhf_kWm8gG5TpTbRxV59sTTKVH-MWxHWij6kbYOHaAIAOEX",
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
		AppID:  "wxaec93043ddef499d",
		Secret: "a810dda42806e814fbcf1237ee824bba",
	}

	if err := GetAccessToken(req, resp); err != nil {
		t.Fatalf("%v", err)
	}

	t.Log(resp)
	t.Log("ok")
}

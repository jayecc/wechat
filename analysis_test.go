package wechat

import "testing"

func TestGetDailyRetain(t *testing.T) {

	req := new(GetDailyRetainRequest)
	resp := new(GetDailyRetainResponse)

	token := "36_MqakaCLpfg1pTbdaoBeuaQV_yrM9A5j5kZAUQGhcFef4e3o6DPD0b7k5fMtnYuE-dky2b7fQwad91Ksk5rvOv6x6eSWdqnj7zxgdW-JEJaZDhf_kWm8gG5TpTbRxV59sTTKVH-MWxHWij6kbYOHaAIAOEX"

	req = &GetDailyRetainRequest{
		BeginDate: "20170313",
		EndDate:   "20170313",
	}

	if err := GetDailyRetain(token, req, resp); err != nil {
		t.Fatalf("%v", err)
	}

	t.Log(resp)
	t.Log("ok")

}

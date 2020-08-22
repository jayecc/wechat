package wechat

import "testing"

func TestGetDailyRetain(t *testing.T) {

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

	t.Log(resp)
	t.Log("ok")

}

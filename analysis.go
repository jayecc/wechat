package wechat

import (
	"github.com/pkg/errors"
	"net/url"
)

// 获取用户访问小程序日留存-请求
type GetDailyRetainRequest struct {
	BeginDate string `json:"begin_date"` //开始日期。格式为 yyyymmdd
	EndDate   string `json:"end_date"`   //结束日期，限定查询1天数据，允许设置的最大值为昨日。格式为 yyyymmdd
}

// 获取用户访问小程序日留存-响应
type GetDailyRetainResponse struct {
	Error
	RefDate    string                  `json:"ref_date"`     //日期
	VisitUvNew []DailyRetainVisitUvNew `json:"visit_uv_new"` //新增用户留存
	VisitUv    []DailyRetainVisitUv    `json:"visit_uv"`     //活跃用户留存
}

type DailyRetainVisitUvNew struct {
	Key   int `json:"key"`   //标识，0开始，表示当天，1表示1天后。依此类推，key取值分别是：0,1,2,3,4,5,6,7,14,30
	Value int `json:"value"` //key对应日期的新增用户数/活跃用户数（key=0时）或留存用户数（k>0时）
}

type DailyRetainVisitUv struct {
	Key   int `json:"key"`   //标识，0开始，表示当天，1表示1天后。依此类推，key取值分别是：0,1,2,3,4,5,6,7,14,30
	Value int `json:"value"` //key对应日期的新增用户数/活跃用户数（key=0时）或留存用户数（k>0时）
}

// 获取用户访问小程序日留存
// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/data-analysis/visit-retain/analysis.getDailyRetain.html
func GetDailyRetain(accessToken string, req *GetDailyRetainRequest, resp *GetDailyRetainResponse) error {

	if accessToken == "" {
		return errors.Wrap(errors.New("access token is empty"), "request param error")
	}

	if req.BeginDate == "" {
		return errors.Wrap(errors.New("begin date is empty"), "request param error")
	}

	if req.EndDate == "" {
		return errors.Wrap(errors.New("end date is empty"), "request param error")
	}

	u := "https://api.weixin.qq.com/datacube/getweanalysisappiddailyretaininfo?access_token=" + url.QueryEscape(accessToken)

	if err := httpPostJSON(DefaultHttpClient, u, req, resp); err != nil {
		return errors.Wrap(err, "http request error")
	}

	if resp.ErrCode != ErrCodeOK {
		return errors.Wrap(errors.New(resp.Error.Error()), "http response error")
	}

	return nil
}

// 获取用户访问小程序月留存-请求
type GetMonthlyRetainRequest struct {
	BeginDate string `json:"begin_date"` //开始日期，为自然月第一天。格式为 yyyymmdd
	EndDate   string `json:"end_date"`   //结束日期，为自然月最后一天，限定查询一个月数据。格式为 yyyymmdd
}

// 获取用户访问小程序月留存-响应
type GetMonthlyRetainResponse struct {
	Error
	RefDate    string                    `json:"ref_date"`     //时间，如："201702"
	VisitUvNew []MonthlyRetainVisitUvNew `json:"visit_uv_new"` //新增用户留存
	VisitUv    []MonthlyRetainVisitUv    `json:"visit_uv"`     //活跃用户留存
}

type MonthlyRetainVisitUvNew struct {
	Key   int `json:"key"`   //标识，0开始，表示当月，1表示1月后。key取值分别是：0,1
	Value int `json:"value"` //key对应日期的新增用户数/活跃用户数（key=0时）或留存用户数（k>0时）
}

type MonthlyRetainVisitUv struct {
	Key   int `json:"key"`   //标识，0开始，表示当月，1表示1月后。key取值分别是：0,1
	Value int `json:"value"` //key对应日期的新增用户数/活跃用户数（key=0时）或留存用户数（k>0时）
}

// 获取用户访问小程序月留存
// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/data-analysis/visit-retain/analysis.getMonthlyRetain.html
func GetMonthlyRetain(accessToken string, req *GetMonthlyRetainRequest, resp *GetMonthlyRetainResponse) error {

	if accessToken == "" {
		return errors.Wrap(errors.New("access token is empty"), "request param error")
	}

	if req.BeginDate == "" {
		return errors.Wrap(errors.New("begin date is empty"), "request param error")
	}

	if req.EndDate == "" {
		return errors.Wrap(errors.New("end date is empty"), "request param error")
	}

	u := "https://api.weixin.qq.com/datacube/getweanalysisappidmonthlyretaininfo?access_token=" + url.QueryEscape(accessToken)

	if err := httpPostJSON(DefaultHttpClient, u, req, resp); err != nil {
		return errors.Wrap(err, "http request error")
	}

	if resp.ErrCode != ErrCodeOK {
		return errors.Wrap(errors.New(resp.Error.Error()), "http response error")
	}

	return nil
}

// 获取用户访问小程序周留存-请求
type GetWeeklyRetainRequest struct {
	BeginDate string `json:"begin_date"` //开始日期，为周一日期。格式为 yyyymmdd
	EndDate   string `json:"end_date"`   //结束日期，为周日日期，限定查询一周数据。格式为 yyyymmdd
}

// 获取用户访问小程序周留存-响应
type GetWeeklyRetainResponse struct {
	Error
	RefDate    string                   `json:"ref_date"`     //时间，如："20170306-20170312"
	VisitUvNew []WeeklyRetainVisitUvNew `json:"visit_uv_new"` //新增用户留存
	VisitUv    []WeeklyRetainVisitUv    `json:"visit_uv"`     //活跃用户留存
}

type WeeklyRetainVisitUvNew struct {
	Key   int `json:"key"`   //标识，0开始，表示当周，1表示1周后。依此类推，取值分别是：0,1,2,3,4
	Value int `json:"value"` //key对应日期的新增用户数/活跃用户数（key=0时）或留存用户数（k>0时）
}

type WeeklyRetainVisitUv struct {
	Key   int `json:"key"`   //标识，0开始，表示当周，1表示1周后。依此类推，取值分别是：0,1,2,3,4
	Value int `json:"value"` //key对应日期的新增用户数/活跃用户数（key=0时）或留存用户数（k>0时）
}

// 获取用户访问小程序周留存
// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/data-analysis/visit-retain/analysis.getWeeklyRetain.html
func GetWeeklyRetain(accessToken string, req *GetWeeklyRetainRequest, resp *GetWeeklyRetainResponse) error {

	if accessToken == "" {
		return errors.Wrap(errors.New("access token is empty"), "request param error")
	}

	if req.BeginDate == "" {
		return errors.Wrap(errors.New("begin date is empty"), "request param error")
	}

	if req.EndDate == "" {
		return errors.Wrap(errors.New("end date is empty"), "request param error")
	}

	u := "https://api.weixin.qq.com/datacube/getweanalysisappidweeklyretaininfo?access_token=" + url.QueryEscape(accessToken)

	if err := httpPostJSON(DefaultHttpClient, u, req, resp); err != nil {
		return errors.Wrap(err, "http request error")
	}

	if resp.ErrCode != ErrCodeOK {
		return errors.Wrap(errors.New(resp.Error.Error()), "http response error")
	}

	return nil
}

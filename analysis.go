package wechat

import (
	"net/url"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/pkg/errors"
)

// GetDailyRetainRequest 获取用户访问小程序日留存-请求
type GetDailyRetainRequest struct {
	BeginDate string `json:"begin_date"` //开始日期。格式为 yyyymmdd
	EndDate   string `json:"end_date"`   //结束日期，限定查询1天数据，允许设置的最大值为昨日。格式为 yyyymmdd
}

// GetDailyRetainResponse 获取用户访问小程序日留存-响应
type GetDailyRetainResponse struct {
	Error
	RefDate    string                  `json:"ref_date"`     //日期
	VisitUvNew []DailyRetainVisitUvNew `json:"visit_uv_new"` //新增用户留存
	VisitUv    []DailyRetainVisitUv    `json:"visit_uv"`     //活跃用户留存
}

// DailyRetainVisitUvNew 新增用户留存
type DailyRetainVisitUvNew struct {
	Key   int `json:"key"`   //标识，0开始，表示当天，1表示1天后。依此类推，key取值分别是：0,1,2,3,4,5,6,7,14,30
	Value int `json:"value"` //key对应日期的新增用户数/活跃用户数（key=0时）或留存用户数（k>0时）
}

// DailyRetainVisitUv 活跃用户留存
type DailyRetainVisitUv struct {
	Key   int `json:"key"`   //标识，0开始，表示当天，1表示1天后。依此类推，key取值分别是：0,1,2,3,4,5,6,7,14,30
	Value int `json:"value"` //key对应日期的新增用户数/活跃用户数（key=0时）或留存用户数（k>0时）
}

// GetDailyRetain 获取用户访问小程序日留存
// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/data-analysis/visit-retain/analysis.getDailyRetain.html
func GetDailyRetain(accessToken string, req *GetDailyRetainRequest, resp *GetDailyRetainResponse) error {

	if err := validation.Validate(accessToken, validation.Required); err != nil {
		return errors.Wrap(err, "request param error")
	}

	if err := validation.ValidateStruct(req,
		validation.Field(&req.BeginDate, validation.Required),
		validation.Field(&req.EndDate, validation.Required),
	); err != nil {
		return errors.Wrap(err, "request param error")
	}

	URL := "https://api.weixin.qq.com/datacube/getweanalysisappiddailyretaininfo?access_token=" + url.QueryEscape(accessToken)

	if err := httpPostJSON(DefaultHTTPClient, URL, req, resp); err != nil {
		return errors.Wrap(err, "http request error")
	}

	if resp.ErrCode != ErrCodeOK {
		return errors.Wrap(errors.New(resp.Error.Error()), "http response error")
	}

	return nil
}

// GetMonthlyRetainRequest 获取用户访问小程序月留存-请求
type GetMonthlyRetainRequest struct {
	BeginDate string `json:"begin_date"` //开始日期，为自然月第一天。格式为 yyyymmdd
	EndDate   string `json:"end_date"`   //结束日期，为自然月最后一天，限定查询一个月数据。格式为 yyyymmdd
}

// GetMonthlyRetainResponse 获取用户访问小程序月留存-响应
type GetMonthlyRetainResponse struct {
	Error
	RefDate    string                    `json:"ref_date"`     //时间，如："201702"
	VisitUvNew []MonthlyRetainVisitUvNew `json:"visit_uv_new"` //新增用户留存
	VisitUv    []MonthlyRetainVisitUv    `json:"visit_uv"`     //活跃用户留存
}

// MonthlyRetainVisitUvNew 新增用户留存
type MonthlyRetainVisitUvNew struct {
	Key   int `json:"key"`   //标识，0开始，表示当月，1表示1月后。key取值分别是：0,1
	Value int `json:"value"` //key对应日期的新增用户数/活跃用户数（key=0时）或留存用户数（k>0时）
}

// MonthlyRetainVisitUv 活跃用户留存
type MonthlyRetainVisitUv struct {
	Key   int `json:"key"`   //标识，0开始，表示当月，1表示1月后。key取值分别是：0,1
	Value int `json:"value"` //key对应日期的新增用户数/活跃用户数（key=0时）或留存用户数（k>0时）
}

// GetMonthlyRetain 获取用户访问小程序月留存
// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/data-analysis/visit-retain/analysis.getMonthlyRetain.html
func GetMonthlyRetain(accessToken string, req *GetMonthlyRetainRequest, resp *GetMonthlyRetainResponse) error {

	if err := validation.Validate(accessToken, validation.Required); err != nil {
		return errors.Wrap(err, "request param error")
	}

	if err := validation.ValidateStruct(req,
		validation.Field(&req.BeginDate, validation.Required),
		validation.Field(&req.EndDate, validation.Required),
	); err != nil {
		return errors.Wrap(err, "request param error")
	}

	URL := "https://api.weixin.qq.com/datacube/getweanalysisappidmonthlyretaininfo?access_token=" + url.QueryEscape(accessToken)

	if err := httpPostJSON(DefaultHTTPClient, URL, req, resp); err != nil {
		return errors.Wrap(err, "http request error")
	}

	if resp.ErrCode != ErrCodeOK {
		return errors.Wrap(errors.New(resp.Error.Error()), "http response error")
	}

	return nil
}

// GetWeeklyRetainRequest 获取用户访问小程序周留存-请求
type GetWeeklyRetainRequest struct {
	BeginDate string `json:"begin_date"` //开始日期，为周一日期。格式为 yyyymmdd
	EndDate   string `json:"end_date"`   //结束日期，为周日日期，限定查询一周数据。格式为 yyyymmdd
}

// GetWeeklyRetainResponse 获取用户访问小程序周留存-响应
type GetWeeklyRetainResponse struct {
	Error
	RefDate    string                   `json:"ref_date"`     //时间，如："20170306-20170312"
	VisitUvNew []WeeklyRetainVisitUvNew `json:"visit_uv_new"` //新增用户留存
	VisitUv    []WeeklyRetainVisitUv    `json:"visit_uv"`     //活跃用户留存
}

// WeeklyRetainVisitUvNew 新增用户留存
type WeeklyRetainVisitUvNew struct {
	Key   int `json:"key"`   //标识，0开始，表示当周，1表示1周后。依此类推，取值分别是：0,1,2,3,4
	Value int `json:"value"` //key对应日期的新增用户数/活跃用户数（key=0时）或留存用户数（k>0时）
}

// WeeklyRetainVisitUv 活跃用户留存
type WeeklyRetainVisitUv struct {
	Key   int `json:"key"`   //标识，0开始，表示当周，1表示1周后。依此类推，取值分别是：0,1,2,3,4
	Value int `json:"value"` //key对应日期的新增用户数/活跃用户数（key=0时）或留存用户数（k>0时）
}

// GetWeeklyRetain 获取用户访问小程序周留存
// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/data-analysis/visit-retain/analysis.getWeeklyRetain.html
func GetWeeklyRetain(accessToken string, req *GetWeeklyRetainRequest, resp *GetWeeklyRetainResponse) error {

	if err := validation.Validate(accessToken, validation.Required); err != nil {
		return errors.Wrap(err, "request param error")
	}

	if err := validation.ValidateStruct(req,
		validation.Field(&req.BeginDate, validation.Required),
		validation.Field(&req.EndDate, validation.Required),
	); err != nil {
		return errors.Wrap(err, "request param error")
	}

	URL := "https://api.weixin.qq.com/datacube/getweanalysisappidweeklyretaininfo?access_token=" + url.QueryEscape(accessToken)

	if err := httpPostJSON(DefaultHTTPClient, URL, req, resp); err != nil {
		return errors.Wrap(err, "http request error")
	}

	if resp.ErrCode != ErrCodeOK {
		return errors.Wrap(errors.New(resp.Error.Error()), "http response error")
	}

	return nil
}

// GetDailySummaryRequest 获取用户访问小程序数据概况-请求
type GetDailySummaryRequest struct {
	BeginDate string `json:"begin_date"` //开始日期。格式为 yyyymmdd
	EndDate   string `json:"end_date"`   //结束日期，限定查询1天数据，允许设置的最大值为昨日。格式为 yyyymmdd
}

// GetDailySummaryResponse 获取用户访问小程序数据概况-响应
type GetDailySummaryResponse struct {
	Error
	List []DailySummary `json:"list"`
}

// DailySummary 访问数据
type DailySummary struct {
	RefDate    string `json:"ref_date"`    //日期，格式为 yyyymmdd
	VisitTotal int    `json:"visit_total"` //累计用户数
	SharePV    int    `json:"share_pv"`    //转发次数
	ShareUV    int    `json:"share_uv"`    //转发人数
}

// GetDailySummary 获取用户访问小程序数据概况
// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/data-analysis/analysis.getDailySummary.html
func GetDailySummary(accessToken string, req *GetDailySummaryRequest, resp *GetDailySummaryResponse) error {

	if err := validation.Validate(accessToken, validation.Required); err != nil {
		return errors.Wrap(err, "request param error")
	}

	if err := validation.ValidateStruct(req,
		validation.Field(&req.BeginDate, validation.Required),
		validation.Field(&req.EndDate, validation.Required),
	); err != nil {
		return errors.Wrap(err, "request param error")
	}

	URL := "https://api.weixin.qq.com/datacube/getweanalysisappiddailysummarytrend?access_token=" + url.QueryEscape(accessToken)

	if err := httpPostJSON(DefaultHTTPClient, URL, req, resp); err != nil {
		return errors.Wrap(err, "http request error")
	}

	if resp.ErrCode != ErrCodeOK {
		return errors.Wrap(errors.New(resp.Error.Error()), "http response error")
	}

	return nil
}

// GetDailyVisitTrendRequest 获取用户访问小程序数据日趋势-请求
type GetDailyVisitTrendRequest struct {
	BeginDate string `json:"begin_date"` //开始日期。格式为 yyyymmdd
	EndDate   string `json:"end_date"`   //结束日期，限定查询1天数据，允许设置的最大值为昨日。格式为 yyyymmdd
}

// GetDailyVisitTrendResponse 获取用户访问小程序数据日趋势-响应
type GetDailyVisitTrendResponse struct {
	Error
	List []DailyVisitTrend `json:"list"` //数据列表
}

// DailyVisitTrend 用户访问小程序数据日趋势
type DailyVisitTrend struct {
	RefDate         string `json:"ref_date"`          //日期，格式为 yyyymmdd
	SessionCnt      int    `json:"session_cnt"`       //打开次数
	VisitPV         int    `json:"visit_pv"`          //访问次数
	VisitUV         int    `json:"visit_uv"`          //访问人数
	VisitUvNew      int    `json:"visit_uv_new"`      //新用户数
	StayTimeUV      int    `json:"stay_time_uv"`      //人均停留时长 (浮点型，单位：秒)
	StayTimeSession int    `json:"stay_time_session"` //次均停留时长 (浮点型，单位：秒)
	VisitDepth      int    `json:"visit_depth"`       //平均访问深度 (浮点型)
}

// GetDailyVisitTrend 获取用户访问小程序数据日趋势
func GetDailyVisitTrend(accessToken string, req *GetDailyVisitTrendRequest, resp *GetDailyVisitTrendResponse) error {

	if err := validation.Validate(accessToken, validation.Required); err != nil {
		return errors.Wrap(err, "request param error")
	}

	if err := validation.ValidateStruct(req,
		validation.Field(&req.BeginDate, validation.Required),
		validation.Field(&req.EndDate, validation.Required),
	); err != nil {
		return errors.Wrap(err, "request param error")
	}

	URL := "https://api.weixin.qq.com/datacube/getweanalysisappiddailyvisittrend?access_token=" + url.QueryEscape(accessToken)

	if err := httpPostJSON(DefaultHTTPClient, URL, req, resp); err != nil {
		return errors.Wrap(err, "http request error")
	}

	if resp.ErrCode != ErrCodeOK {
		return errors.Wrap(errors.New(resp.Error.Error()), "http response error")
	}

	return nil
}

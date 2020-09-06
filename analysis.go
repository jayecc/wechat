package wechat

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/pkg/errors"
)

type aims int

const (
	// GetDailyRetain 获取用户访问小程序日留存
	// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/data-analysis/visit-retain/analysis.getDailyRetain.html
	GetDailyRetain aims = iota
	// GetMonthlyRetain 获取用户访问小程序月留存
	// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/data-analysis/visit-retain/analysis.getMonthlyRetain.html
	GetMonthlyRetain
	// GetWeeklyRetain 获取用户访问小程序周留存
	// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/data-analysis/visit-retain/analysis.getWeeklyRetain.html
	GetWeeklyRetain
	// GetDailySummary 获取用户访问小程序数据概况
	// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/data-analysis/analysis.getDailySummary.html
	GetDailySummary
	// GetDailyVisitTrend 获取用户访问小程序数据日趋势
	// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/data-analysis/visit-trend/analysis.getDailyVisitTrend.html
	GetDailyVisitTrend
	// GetMonthlyVisitTrend 获取用户访问小程序数据月趋势
	// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/data-analysis/visit-trend/analysis.getMonthlyVisitTrend.html
	GetMonthlyVisitTrend
	// GetWeeklyVisitTrend 获取用户访问小程序数据周趋势
	// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/data-analysis/visit-trend/analysis.getWeeklyVisitTrend.html
	GetWeeklyVisitTrend
	// GetUserPortrait 获取小程序新增或活跃用户的画像分布数据
	// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/data-analysis/analysis.getUserPortrait.html
	GetUserPortrait
	// GetVisitDistribution 获取用户小程序访问分布数据
	// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/data-analysis/analysis.getVisitDistribution.html
	GetVisitDistribution
	// GetVisitPage 访问页面
	// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/data-analysis/analysis.getVisitPage.html
	GetVisitPage
)

var datacubeURL = map[aims]string{
	GetDailyRetain:       "https://api.weixin.qq.com/datacube/getweanalysisappiddailyretaininfo",
	GetMonthlyRetain:     "https://api.weixin.qq.com/datacube/getweanalysisappidmonthlyretaininfo",
	GetWeeklyRetain:      "https://api.weixin.qq.com/datacube/getweanalysisappidweeklyretaininfo",
	GetDailySummary:      "https://api.weixin.qq.com/datacube/getweanalysisappiddailysummarytrend",
	GetDailyVisitTrend:   "https://api.weixin.qq.com/datacube/getweanalysisappiddailyvisittrend",
	GetMonthlyVisitTrend: "https://api.weixin.qq.com/datacube/getweanalysisappidmonthlyvisittrend",
	GetWeeklyVisitTrend:  "https://api.weixin.qq.com/datacube/getweanalysisappidweeklyvisittrend",
	GetUserPortrait:      "https://api.weixin.qq.com/datacube/getweanalysisappiduserportrait",
	GetVisitDistribution: "https://api.weixin.qq.com/datacube/getweanalysisappidvisitdistribution",
	GetVisitPage:         "https://api.weixin.qq.com/datacube/getweanalysisappidvisitpage",
}

// DatacubeURL 数据分析URL
func DatacubeURL(a aims) string {
	return datacubeURL[a]
}

// GetDatacube 数据分析
func GetDatacube(aims aims, accessToken string, req *GetDatacubeRequest, resp interface{}) error {

	if err := validation.Validate(accessToken, validation.Empty); err != nil {
		return errors.Wrap(err, "request param error")
	}

	if err := validation.ValidateStruct(req,
		validation.Field(&req.BeginDate, validation.Required),
		validation.Field(&req.EndDate, validation.Required, validation.Min(req.BeginDate)),
	); err != nil {
		return errors.Wrap(err, "request param error")
	}

	URL, err := encodeURL(DatacubeURL(aims), queryParams{"access_token": accessToken})
	if err != nil {
		return errors.Wrap(err, "encode url error")
	}

	if err = httpPostJSON(DefaultHTTPClient, URL, req, resp); err != nil {
		return errors.Wrap(err, "http request error")
	}

	return nil
}

// GetDatacubeRequest 请求
type GetDatacubeRequest struct {
	BeginDate string `json:"begin_date"` //开始日期。格式为 yyyymmdd
	EndDate   string `json:"end_date"`   //结束日期，格式为 yyyymmdd
}

// GetDailyRetainResponse 获取用户访问小程序日留存-响应
type GetDailyRetainResponse struct {
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

// GetMonthlyRetainResponse 获取用户访问小程序月留存-响应
type GetMonthlyRetainResponse struct {
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

// GetWeeklyRetainResponse 获取用户访问小程序周留存-响应
type GetWeeklyRetainResponse struct {
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

// GetDailySummaryResponse 获取用户访问小程序数据概况-响应
type GetDailySummaryResponse struct {
	List []DailySummary `json:"list"`
}

// DailySummary 访问数据
type DailySummary struct {
	RefDate    string `json:"ref_date"`    //日期，格式为 yyyymmdd
	VisitTotal int    `json:"visit_total"` //累计用户数
	SharePV    int    `json:"share_pv"`    //转发次数
	ShareUV    int    `json:"share_uv"`    //转发人数
}

// GetDailyVisitTrendResponse 获取用户访问小程序数据日趋势-响应
type GetDailyVisitTrendResponse struct {
	List []DailyVisitTrend `json:"list"` //数据列表
}

// DailyVisitTrend 用户访问小程序数据日趋势
type DailyVisitTrend struct {
	RefDate         string  `json:"ref_date"`          //日期，格式为 yyyymmdd
	SessionCnt      int     `json:"session_cnt"`       //打开次数
	VisitPV         int     `json:"visit_pv"`          //访问次数
	VisitUV         int     `json:"visit_uv"`          //访问人数
	VisitUvNew      int     `json:"visit_uv_new"`      //新用户数
	StayTimeUV      float64 `json:"stay_time_uv"`      //人均停留时长 (浮点型，单位：秒)
	StayTimeSession float64 `json:"stay_time_session"` //次均停留时长 (浮点型，单位：秒)
	VisitDepth      float64 `json:"visit_depth"`       //平均访问深度 (浮点型)
}

// GetMonthlyVisitTrendRespone 获取用户访问小程序数据月趋势-响应
type GetMonthlyVisitTrendRespone struct {
	List []MonthlyVisitTrend `json:"list"`
}

// MonthlyVisitTrend 用户访问小程序数据月趋势
type MonthlyVisitTrend struct {
	RefDate         string  `json:"ref_date"`          //时间，格式为 yyyymm，如："201702"
	SessionCnt      int     `json:"session_cnt"`       //打开次数（自然月内汇总）
	VisitPV         int     `json:"visit_pv"`          //访问次数（自然月内汇总）
	VisitUV         int     `json:"visit_uv"`          //访问人数（自然月内去重）
	VisitUvNew      int     `json:"visit_uv_new"`      //新用户数（自然月内去重）
	StayTimeUV      float64 `json:"stay_time_uv"`      //人均停留时长 (浮点型，单位：秒)
	StayTimeSession float64 `json:"stay_time_session"` //次均停留时长 (浮点型，单位：秒)
	VisitDepth      float64 `json:"visit_depth"`       //平均访问深度 (浮点型)
}

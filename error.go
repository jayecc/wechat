package wechat

import (
	"fmt"
)

const (
	// ErrCodeOK 请求成功
	ErrCodeOK = 0
)

// Error 通用错误
type Error struct {
	ErrCode int  `json:"errcode"`
	ErrMsg  string `json:"errmsg"`
}

func (err *Error) Error() string {
	return fmt.Sprintf("errcode: %d, errmsg: %s", err.ErrCode, err.ErrMsg)
}

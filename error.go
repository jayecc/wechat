package wechat

import (
	"fmt"
)

const (
	ErrCodeOK = 0 // 请求成功
)

type Error struct {
	ErrCode int64  `json:"errcode"`
	ErrMsg  string `json:"errmsg"`
}

func (err *Error) Error() string {
	return fmt.Sprintf("errcode: %d, errmsg: %s", err.ErrCode, err.ErrMsg)
}

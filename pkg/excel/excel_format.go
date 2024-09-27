package excel

import (
	"time"

	"github.com/miajio/aide/pkg/utils"
)

// FormatFunc 格式化函数
type FormatFunc func(any) string

// UnixDateFormatYMDHMS 时间戳转时间字符串 2006-01-02 15:04:05
func UnixDateFormatYMDHMS(val any) string {
	if utils.IsZero(val) {
		return ""
	}
	// 时间戳
	ival := val.(int64)
	// 时间戳转时间
	t := time.Unix(ival, 0)
	// 时间转字符串
	return t.Format("2006-01-02 15:04:05")
}

// UnixDateFormatYMD 时间戳转时间字符串 2006-01-02
func UnixDateFormatYMD(val any) string {
	if utils.IsZero(val) {
		return ""
	}
	// 时间戳
	ival := val.(int64)
	// 时间戳转时间
	t := time.Unix(ival, 0)
	// 时间转字符串
	return t.Format("2006-01-02")
}

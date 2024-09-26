package system

import (
	"fmt"
	"math"
	"strings"
	"time"
)

// logn computes the base-b logarithm of n
func computesLogarithm(n, b float64) float64 {
	return math.Log(n) / math.Log(b)
}

// humanateBytes converts bytes to a human readable string
func humanateBytes(s uint64, base float64, sizes []string) string {
	if s < 10 {
		return fmt.Sprintf("%d B", s)
	}
	e := math.Floor(computesLogarithm(float64(s), base))
	suffix := sizes[int(e)]
	val := float64(s) / math.Pow(base, math.Floor(e))
	f := "%.0f"
	if val < 10 {
		f = "%.1f"
	}

	return fmt.Sprintf(f+" %s", val, suffix)
}

// FileSize calculates the file size and generate user-friendly string.
func FileSize(s int64) string {
	sizes := []string{"B", "KB", "MB", "GB", "TB", "PB", "EB"}
	return humanateBytes(uint64(s), 1024, sizes)
}

// Seconds-based time units
const (
	Minute = 60
	Hour   = 60 * Minute
	Day    = 24 * Hour
	Week   = 7 * Day
	Month  = 30 * Day
	Year   = 12 * Month
)

// computeTimeDiff computes the time difference and generate user-friendly string.
func computeTimeDiff(diff int64) (int64, string) {
	diffStr := ""
	switch {
	case diff <= 0:
		diff = 0
		diffStr = "now"
	case diff < 2:
		diff = 0
		diffStr = "1 秒"
	case diff < 1*Minute:
		diffStr = fmt.Sprintf("%d 秒", diff)
		diff = 0

	case diff < 2*Minute:
		diff -= 1 * Minute
		diffStr = "1 分"
	case diff < 1*Hour:
		diffStr = fmt.Sprintf("%d 分", diff/Minute)
		diff -= diff / Minute * Minute

	case diff < 2*Hour:
		diff -= 1 * Hour
		diffStr = "1 小时"
	case diff < 1*Day:
		diffStr = fmt.Sprintf("%d 小时", diff/Hour)
		diff -= diff / Hour * Hour

	case diff < 2*Day:
		diff -= 1 * Day
		diffStr = "1 天"
	case diff < 1*Week:
		diffStr = fmt.Sprintf("%d 天", diff/Day)
		diff -= diff / Day * Day

	case diff < 2*Week:
		diff -= 1 * Week
		diffStr = "1 周"
	case diff < 1*Month:
		diffStr = fmt.Sprintf("%d 周", diff/Week)
		diff -= diff / Week * Week

	case diff < 2*Month:
		diff -= 1 * Month
		diffStr = "1 月"
	case diff < 1*Year:
		diffStr = fmt.Sprintf("%d 月", diff/Month)
		diff -= diff / Month * Month

	case diff < 2*Year:
		diff -= 1 * Year
		diffStr = "1 年"
	default:
		diffStr = fmt.Sprintf("%d 年", diff/Year)
		diff = 0
	}
	return diff, diffStr
}

// TimeSincePro calculates the time interval and generate full user-friendly string.
func TimeSincePro(then time.Time) string {
	now := time.Now()
	diff := now.Unix() - then.Unix()

	if then.After(now) {
		return "future"
	}

	var timeStr, diffStr string
	for {
		if diff == 0 {
			break
		}

		diff, diffStr = computeTimeDiff(diff)
		timeStr += ", " + diffStr
	}
	return strings.TrimPrefix(timeStr, ", ")
}

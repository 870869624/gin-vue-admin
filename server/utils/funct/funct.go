package funct

import (
	"strconv"
	"time"

	userid "github.com/google/uuid"
)

// 将string转换为int64
func StrToInt64(str string) int64 {

	i, _ := strconv.ParseInt(str, 10, 64)

	return i

}

func StrToInt(x string) int {

	i, _ := strconv.Atoi(x)

	return i

}
func Int64ToStr(x int64) string {
	i := strconv.FormatInt(x, 10)
	return i

}
func StrToFloat64(str string) float64 {
	i, _ := strconv.ParseFloat(str, 64)
	return i

}

func GenerateUserUuid(prefix string) string {
	uuidObj := userid.New()
	uid := uuidObj.ID()
	uuidstr := strconv.Itoa(int(uid))
	var id string
	if prefix != "" {
		id = prefix + uuidstr
	} else {
		id = uuidstr
	}
	return id
}

func StrUnixToTime(str string) time.Time {
	return time.Unix(StrToInt64(str), 0)
}

// 寻找本周的开始时间
func GetWeekStartTime() time.Time {
	now := time.Now()
	weekday := int(now.Weekday())
	daysAgo := (weekday + 6) % 7
	weekStart := now.AddDate(0, 0, -daysAgo)
	return time.Date(weekStart.Year(), weekStart.Month(), weekStart.Day(), 0, 0, 0, 0, now.Location())

}

// 本周结束时间
func GetWeekEndTime() time.Time {
	now := time.Now()
	weekday := int(now.Weekday())
	daysUntilSunday := 7 - weekday
	weekEnd := now.AddDate(0, 0, daysUntilSunday)
	return time.Date(weekEnd.Year(), weekEnd.Month(), weekEnd.Day(), 23, 59, 59, 999999999, now.Location())
}

func GetMonthEndTime() time.Time {
	now := time.Now()
	return time.Date(now.Year(), now.Month()+1, 0, 23, 59, 59, 0, now.Location())
}

func SetDayStartTime(date time.Time) time.Time {
	return time.Date(date.Year(), date.Month(), date.Day(), 0, 0, 0, 0, date.Location())
}

func SetDayEndTime(date time.Time) time.Time {
	return time.Date(date.Year(), date.Month(), date.Day(), 23, 59, 59, 999999999, date.Location())
}

func SetMonthStartTime(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month(), 1, 0, 0, 0, 0, t.Location())
}

func SetMonthEndTime(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month()+1, 0, 23, 59, 59, 0, t.Location())
}

// 判断是否是同一天
func IsDay(a, b time.Time) bool {
	return a.Year() == b.Year() && a.Month() == b.Month() && a.Day() == b.Day()
}

// 判断是否是周一
func IsWeekend(weekday time.Time) bool {
	if weekday.Weekday() == time.Monday {
		return true
	}
	return false
}

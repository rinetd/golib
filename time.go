package golib

import (
	"math"
	"time"
)

var DefaultTimeZone = time.UTC

// 把时间转成 unix的浮点,(unix浮点,是以utc时间储存的)
func GetUnixFloat(t1 time.Time) float64 {
	return (float64(t1.Nanosecond()) / 1e9) + float64(t1.Unix())
}

// unix的浮点 转成时间 (unix浮点,是以utc时间储存的,单位秒)
func FromUnixFloat(f float64) time.Time {
	s, ns := math.Modf(f)
	return time.Unix(int64(s), int64(ns*1e9)).In(DefaultTimeZone)
}

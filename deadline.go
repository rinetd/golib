package golib

import (
	"os"
	"time"
)

// DeadLine("2019-01-29 00:00:00")
func DeadLine(date string) {
	deadline, _ := time.ParseInLocation("2006-01-02 15:04:05", date, time.Local)
	if time.Now().After(deadline) {
		print("程序版本不兼容")
		time.Sleep(time.Second)
		os.Exit(0)
	}
	println(date)
}

func Test_Deadline() {
	DeadLine("2019-01-29 00:00:00")
}

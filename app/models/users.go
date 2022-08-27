package models

import (
	"fmt"
	"time"
)

func calculateTime() {
	now := time.Now()
	day1 := time.Date(now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second(), now.Nanosecond(), time.Local)
	day2 := time.Date(2081, 7, 6, 10, 30, 0, 0, time.Local)
	duration := day2.Sub(day1)
	fmt.Println(duration) // => "60h30m0s"

	hours0 := int(duration.Hours())
	days := hours0 / 24
	hours := hours0 % 24
	mins := int(duration.Minutes()) % 60
	secs := int(duration.Seconds()) % 60
	fmt.Printf("残り：%d日%d時間%d分%d秒\n", days, hours, mins, secs)
	// => "2 days + 12 hours + 30 minutes + 0 seconds"
}

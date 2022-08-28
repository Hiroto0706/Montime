package controllers

import (
	"log"
	"net/http"
	"strconv"
	"time"
)

type RemainingData struct {
	Days     int
	Hours    int
	Mins     int
	Sec      int
	Remainig int
	Sex      string
	Duration time.Duration
}

var remainingData RemainingData

func top(w http.ResponseWriter, r *http.Request) {
	generateHTML(w, nil, "layout", "top")
}

func calculate(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Println(err)
	}

	birthdayYear, err := strconv.Atoi(r.PostFormValue("year"))
	if err != nil {
		log.Println(err)
	}
	birthdayMonth, err := strconv.Atoi(r.PostFormValue("month"))
	if err != nil {
		log.Println(err)
	}
	birthdayDay, err := strconv.Atoi(r.PostFormValue("day"))
	if err != nil {
		log.Println(err)
	}

	birthday := time.Date(birthdayYear, time.Month(birthdayMonth), birthdayDay, 0, 0, 0, 0, time.Local)
	sex := r.PostFormValue("sex")
	// fmt.Println(sex)

	now := time.Now()
	day1 := time.Date(now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second(), now.Nanosecond(), time.Local)

	var day2 time.Time
	switch sex {
	case "male":
		day2 = birthday.AddDate(82, 0, 0)
	case "female":
		day2 = birthday.AddDate(88, 0, 0)
	}

	duration := day2.Sub(day1)
	// fmt.Println(duration)

	hours0 := int(duration.Hours())
	days := hours0 / 24
	hours := hours0 % 24
	mins := int(duration.Minutes()) % 60
	secs := int(duration.Seconds()) % 60
	// fmt.Printf("残り：%d日%d時間%d分%d秒\n", days, hours, mins, secs)

	remaining := days / 365
	// fmt.Printf("余命：%d年", remaining)

	remainingData = RemainingData{
		Days:     days,
		Hours:    hours,
		Mins:     mins,
		Sec:      secs,
		Remainig: remaining,
		Sex:      sex,
		Duration: duration,
	}

	// fmt.Println(remainingData)

	generateHTML(w, remainingData, "layout", "top")
}

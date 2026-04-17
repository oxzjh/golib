package golib

import "time"

func GetTime() int64 {
	return time.Now().Unix()
}

func GetToday() int64 {
	now := time.Now()
	hour, min, sec := now.Clock()
	return now.Unix() - int64(hour*3600+min*60+sec)
}

func GetTodayTimeRemaining() int {
	hour, min, sec := time.Now().Clock()
	return 86400 - hour*3600 - min*60 - sec
}

func GetSunday() int64 {
	now := time.Now()
	var weekday int = int(now.Weekday())
	hour, min, sec := now.Clock()
	return now.Unix() - int64(weekday*86400+hour*3600+min*60+sec)
}

func GetMonday() int64 {
	now := time.Now()
	var weekday int = int(now.Weekday())
	if weekday == 0 {
		weekday = 7
	}
	hour, min, sec := now.Clock()
	return now.Unix() - int64((weekday-1)*86400+hour*3600+min*60+sec)
}

func GetMonthFirstDay() int64 {
	now := time.Now()
	day := now.Day()
	hour, min, sec := now.Clock()
	return now.Unix() - int64((day-1)*86400+hour*3600+min*60+sec)
}

func GetTodayUTC() int64 {
	now := time.Now().UTC()
	hour, min, sec := now.Clock()
	return now.Unix() - int64(hour*3600+min*60+sec)
}

func GetSundayUTC() int64 {
	now := time.Now().UTC()
	var weekday int = int(now.Weekday())
	hour, min, sec := now.Clock()
	return now.Unix() - int64(weekday*86400+hour*3600+min*60+sec)
}

func GetMondayUTC() int64 {
	now := time.Now().UTC()
	var weekday int = int(now.Weekday())
	if weekday == 0 {
		weekday = 7
	}
	hour, min, sec := now.Clock()
	return now.Unix() - int64((weekday-1)*86400+hour*3600+min*60+sec)
}

func GetMonthFirstDayUTC() int64 {
	now := time.Now().UTC()
	day := now.Day()
	hour, min, sec := now.Clock()
	return now.Unix() - int64((day-1)*86400+hour*3600+min*60+sec)
}

func GetDayStamp(tz int64) int64 {
	return (GetTime() + tz*3600) / 86400
}

func GetTimeStamp(day, tz int64) int64 {
	return day*86400 - tz*3600
}

package golib

import "time"

func Schedule(d time.Duration, callback func(int64)) {
	go func() {
		for t := range time.Tick(d) {
			callback(t.Unix())
		}
	}()
}

func ScheduleDay(hour int, callback func()) {
	go func() {
		for {
			now := time.Now()
			next := time.Date(now.Year(), now.Month(), now.Day(), hour, 0, 0, 0, now.Location())
			if now.Hour() >= hour {
				next = next.Add(time.Hour * 24)
			}
			t := time.NewTimer(next.Sub(now))
			<-t.C
			callback()
		}
	}()
}

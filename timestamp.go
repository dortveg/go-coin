package main

import (
	"fmt"
	"time"
)

func getPrevTS() string {
	prevTime := time.Now().Add(-1 * time.Hour)
	prevTS := int64(time.Nanosecond) * prevTime.UnixNano() / int64(time.Millisecond)
	prev := fmt.Sprint(prevTS)
	return prev
}

func getLastTS() string {
	prevTime := time.Now().Add(-2 * time.Hour)
	prevTS := int64(time.Nanosecond) * prevTime.UnixNano() / int64(time.Millisecond)
	prev := fmt.Sprint(prevTS)
	return prev
}

func getWeekTS1() string {
	oldTime := time.Now().Add(-169 * time.Hour)
	oldTS := int64(time.Nanosecond) * oldTime.UnixNano() / int64(time.Millisecond)
	week1 := fmt.Sprint(oldTS)
	return week1
}

func getWeekTS2() string {
	oldTime := time.Now().Add(-168 * time.Hour)
	oldTS := int64(time.Nanosecond) * oldTime.UnixNano() / int64(time.Millisecond)
	week2 := fmt.Sprint(oldTS)
	return week2
}

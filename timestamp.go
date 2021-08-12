package main

import (
	"fmt"
	"time"
)

func getCurTS() string {
	curTime := time.Now()
	curTS := int64(time.Nanosecond) * curTime.UnixNano() / int64(time.Millisecond)
	cur := fmt.Sprint(curTS)
	return cur
}

func getPrevTS() string {
	prevTime := time.Now().Add(-1 * time.Hour)
	prevTS := int64(time.Nanosecond) * prevTime.UnixNano() / int64(time.Millisecond)
	prev := fmt.Sprint(prevTS)
	return prev
}

func getOldTS() string {
	oldTime := time.Now().Add(-4 * time.Hour)
	oldTS := int64(time.Nanosecond) * oldTime.UnixNano() / int64(time.Millisecond)
	old := fmt.Sprint(oldTS)
	return old
}

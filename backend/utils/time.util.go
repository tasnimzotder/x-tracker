package utils

import "time"

func GetCurrentTimeMilli() int64 {
	return time.Now().UnixNano() / int64(time.Millisecond)
}

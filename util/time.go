package util

import "time"

//GetMsSinceEpoch returns the number of milliseconds since Jan 01 1970
func GetMsSinceEpoch() int64 {
	return time.Now().UnixNano() / int64(time.Millisecond)
}

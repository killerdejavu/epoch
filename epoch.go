package epoch

import (
	"fmt"
	"time"
)

func GetCurrentTime() int64 {
	now := time.Now()
	secs := now.Unix()
	fmt.Println(secs)
	//nanos := now.UnixNano()
	////
	////// Note that there is no `UnixMillis`, so to get the
	////// milliseconds since epoch you'll need to manually
	////// divide from nanoseconds.
	//millis := nanos / 1000000
	////
	//fmt.Println(millis)
	//fmt.Println(nanos)
	//
	//// You can also convert integer seconds or nanoseconds
	//// since the epoch into the corresponding `time`.
	//fmt.Println(time.Unix(secs, 0))
	//fmt.Println(time.Unix(0, nanos))
	return secs
}
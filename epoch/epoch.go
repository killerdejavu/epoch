package epoch

import (
	"fmt"
	"os"
	"text/tabwriter"
	"time"
)

func GetCurrentTime() int64 {
	now := time.Now()
	secs := now.Unix()
	//fmt.Println(secs)
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

func DisplayDataForEpochTime(epochTime int64) {
	ct := time.Now()
	t := time.Unix(epochTime, 0)
	millis := t.UnixNano() / 1000000
	diff := ct.Sub(t)
	var diffWord string
	if diff > 0 {
		diffWord = "in past"
	} else {
		diffWord = "in future"
		diff = -diff
	}
	w := new(tabwriter.Writer)
	w.Init(os.Stdout, 8, 15, 0, '\t', 0)
	defer w.Flush()

	fmt.Fprintf(w, "epoch time: \t%d\n", t.Unix())
	fmt.Fprintf(w, "epoch time millis: \t%d\n", millis)
	fmt.Fprintf(w, "epoch time nano: \t%d\n", t.UnixNano())
	fmt.Fprintf(w, "epoch time ISO: \t%s\n", t)
	fmt.Fprintf(w, "epoch time ISO UTC: \t%s\n", t.UTC())
	fmt.Fprintf(w, "relative to now: \t%s %s\n", diff, diffWord)
}

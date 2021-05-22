package main

import (
	"fmt"
	"github.com/killerdejavu/epoch/epoch"
	"os"
	"strconv"
)

func main() {
	var epochTime int64
	var err error

	if len(os.Args) > 1 {
		epochTime, err = strconv.ParseInt(os.Args[1], 10, 64)
		if err != nil {
			fmt.Println(os.Args[1], "is not a valid epoch time")
			os.Exit(2)
		}
	} else {
		epochTime = epoch.GetCurrentTime()
	}
	epoch.DisplayDataForEpochTime(epochTime)
}

package helpers

import (
	"time"
)

func StringToTime(dateStr *string, resultTime *time.Time) (ok bool) {
	var err error
	*resultTime, err = time.Parse(time.RFC3339Nano, *dateStr)
	if !CheckErr(err) {
		return
	}

	return true
}

func StringListToDateList(str *string, dateList *[]time.Time, ch chan bool) {
	if *str == "" {
		ch <- false
		return
	}
	var date time.Time
	if StringToTime(str, &date) {
		*dateList = append(*dateList, date)
	}
	ch <- true
}

func TimeToString(dateStr *string, date *time.Time) {
	*dateStr = date.Format(time.RFC3339Nano)
}

package helpers

import (
	"get/set"
	"time"
)

func AskTimeOut(ch chan bool) {
	time.Sleep(set.TimeOutSecForAskMongo)
	ch <- false
}

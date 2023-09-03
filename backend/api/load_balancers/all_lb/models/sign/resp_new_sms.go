package sign

import (
	"all_lb/helpers"
	"encoding/json"
)

type RespNewSms struct {
	Empty   bool
	TimeOut string
}

func (r *RespNewSms) FromJson(data *[]byte) (ok bool) {
	err := json.Unmarshal(*data, r)
	if !helpers.CheckErr(err) {
		return false
	}

	return true
}

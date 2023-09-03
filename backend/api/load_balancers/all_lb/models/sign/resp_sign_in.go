package sign

import (
	"all_lb/helpers"
	"encoding/json"
)

type RespSignIn struct {
	AccessToken  string
	RefreshToken string
	AccessUuid   string
	RefreshUuid  string
	AtExpires    int64
	RtExpires    int64
}

func (r *RespSignIn) FromJson(data *[]byte) (ok bool) {
	err := json.Unmarshal(*data, r)
	if !helpers.CheckErr(err) {
		return false
	}

	return true
}

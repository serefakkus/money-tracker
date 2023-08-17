package sign

import (
	"encoding/json"
	"sign_lb/helpers"
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

package sign

import (
	"encoding/json"
	"sign_lb/helpers"
)

type ReqNewUser struct {
	UserId string
}

func (r *ReqNewUser) ToJson() (data []byte) {
	var err error
	data, err = json.Marshal(r)
	helpers.CheckErr(err)

	return
}

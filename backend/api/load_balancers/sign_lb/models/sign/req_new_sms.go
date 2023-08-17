package sign

import (
	"encoding/json"
	"sign_lb/helpers"
	"sign_lb/set"
)

type ReqNewSms struct {
	Phone string
}

func (r *ReqNewSms) ToJson() (data []byte) {
	var err error
	data, err = json.Marshal(r)
	helpers.CheckErr(err)
	return
}

func (r *ReqNewSms) SendToServer(resp *RespNewSms, isRef bool) (ok bool, code int) {
	if isRef {
		ok, code = helpers.SendToServerPost(r, resp, set.SignUri, set.NewSmsRefPath)
	} else {
		ok, code = helpers.SendToServerPost(r, resp, set.SignUri, set.NewSmsPath)
	}

	if !ok {
		return false, code
	}

	return true, code
}

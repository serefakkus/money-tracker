package sign

import (
	"all_lb/helpers"
	"all_lb/set"
	"encoding/json"
)

type ReqPhoneCode struct {
	Phone string
	Code  string
}

func (r *ReqPhoneCode) ToJson() (data []byte) {
	var err error
	data, err = json.Marshal(r)
	helpers.CheckErr(err)
	return
}

func (r *ReqPhoneCode) SendToServer(isRef bool) (ok bool, code int) {
	if isRef {
		ok, code = helpers.SendToServerPostWithOutResp(r, set.SignUri, set.AskSmsRefPath)
	} else {
		ok, code = helpers.SendToServerPostWithOutResp(r, set.SignUri, set.AskSmsPath)
	}

	if !ok {
		return false, code
	}

	return true, code
}

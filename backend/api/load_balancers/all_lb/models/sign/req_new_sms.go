package sign

import (
	"all_lb/helpers"
	"all_lb/models"
	"all_lb/set"
	"encoding/json"
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

func (r *ReqNewSms) SendToServer(resp *models.RespSign, isRef bool) (ok bool, code int) {
	if isRef {
		ok, code = helpers.SendToServerPost(r, &resp.Message, set.SignUri, set.NewSmsRefPath)
	} else {
		ok, code = helpers.SendToServerPost(r, &resp.Message, set.SignUri, set.NewSmsPath)
	}

	if !ok {
		return false, code
	}

	return true, code
}

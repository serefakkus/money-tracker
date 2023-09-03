package sign

import (
	"all_lb/helpers"
	"all_lb/models"
	"all_lb/set"
	"encoding/json"
)

type ReqSignUp struct {
	Phone string
	Code  string
	Pass  string
	Ip    Ip
}

func (r *ReqSignUp) ToJson() (data []byte) {
	var err error
	data, err = json.Marshal(r)
	helpers.CheckErr(err)

	return
}

func (r *ReqSignUp) SendToServer(resp *models.RespSign, isRef bool) (ok bool, code int) {
	if isRef {
		ok, code = helpers.SendToServerPost(r, &resp.Message, set.SignUri, set.RefPassPath)
	} else {
		ok, code = helpers.SendToServerPost(r, &resp.Message, set.SignUri, set.SignUpPath)
	}

	if !ok {
		return false, code
	}

	return true, code
}

func (r *ReqSignUp) SendToServerNewUser(resp *models.RespSign) (ok bool, code int) {

	ok, code = helpers.SendToServerPost(r, &resp.Message, set.SignUri, set.RefPassPath)

	if !ok {
		return false, code
	}

	return true, code
}

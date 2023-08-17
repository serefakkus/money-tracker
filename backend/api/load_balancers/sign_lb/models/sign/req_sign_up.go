package sign

import (
	"encoding/json"
	"sign_lb/helpers"
	"sign_lb/set"
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

func (r *ReqSignUp) SendToServer(resp *RespSignIn, isRef bool) (ok bool, code int) {
	if isRef {
		ok, code = helpers.SendToServerPost(r, resp, set.SignUri, set.RefPassPath)
	} else {
		ok, code = helpers.SendToServerPost(r, resp, set.SignUri, set.SignUpPath)
	}

	if !ok {
		return false, code
	}

	return true, code
}

func (r *ReqSignUp) SendToServerNewUser(resp *RespSignIn) (ok bool, code int) {

	ok, code = helpers.SendToServerPost(r, resp, set.SignUri, set.RefPassPath)

	if !ok {
		return false, code
	}

	return true, code
}

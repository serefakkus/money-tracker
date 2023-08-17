package sign

import (
	"encoding/json"
	"sign_lb/helpers"
	"sign_lb/set"
)

type ReqSignIn struct {
	Phone string
	Pass  string
	Ip    Ip
}

func (r *ReqSignIn) SendToServer(resp *RespSignIn) (ok bool, code int) {

	ok, code = helpers.SendToServerPost(r, resp, set.SignUri, set.SignInPath)

	if !ok {
		return false, code
	}

	return true, code
}

func (r *ReqSignIn) ToJson() (data []byte) {
	var err error
	data, err = json.Marshal(r)
	helpers.CheckErr(err)
	return
}

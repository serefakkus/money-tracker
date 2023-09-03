package sign

import (
	"all_lb/helpers"
	"all_lb/models"
	"all_lb/set"
	"encoding/json"
)

type ReqSignIn struct {
	Phone string
	Pass  string
	Ip    Ip
}

func (r *ReqSignIn) SendToServer(resp *models.RespSign) (ok bool, code int) {

	ok, code = helpers.SendToServerPost(r, &resp.Message, set.SignUri, set.SignInPath)

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

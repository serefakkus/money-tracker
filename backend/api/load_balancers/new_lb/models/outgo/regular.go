package outgo

import (
	"encoding/json"
	"new_lb/helpers"
	"new_lb/set"
)

type ReqOutRegular struct {
	UserId      string
	RegularId   string
	IntervalDay int
	Category    string
	Emoji       string
	Not         string
	Amount      float64
	Time        string
	Date        string
	To          FromUser
}

func (r *ReqOutRegular) ToJson() (data []byte) {
	var err error
	data, err = json.Marshal(r)
	helpers.CheckErr(err)
	return
}

func (r *ReqOutRegular) FromJson(data *[]byte) bool {
	err := json.Unmarshal(*data, r)
	return helpers.CheckErr(err)
}

func (r *ReqOutRegular) SendToServerNew() (ok bool, code int) {
	return helpers.SendToServerPostWithOutResp(r, set.OutGoBaseUri, set.OutGoRegNewPath)
}

func (r *ReqOutRegular) SendToServerRef() (ok bool, code int) {
	return helpers.SendToServerPostWithOutResp(r, set.OutGoBaseUri, set.OutGoRegRefPath)
}

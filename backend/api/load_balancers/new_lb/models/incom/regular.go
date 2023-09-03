package incom

import (
	"encoding/json"
	"new_lb/helpers"
	"new_lb/set"
)

type ReqInRegularNew struct {
	UserId      string
	RegularId   string
	IntervalDay int
	Category    string
	Emoji       string
	Not         string
	Amount      float64
	Time        string
	Date        string
	From        FromUser
}

func (r *ReqInRegularNew) ToJson() (data []byte) {
	var err error
	data, err = json.Marshal(r)
	helpers.CheckErr(err)
	return
}

func (r *ReqInRegularNew) FromJson(data *[]byte) bool {
	err := json.Unmarshal(*data, r)
	return helpers.CheckErr(err)
}

func (r *ReqInRegularNew) SendToServerNew() (ok bool, code int) {
	return helpers.SendToServerPostWithOutResp(r, set.InComBaseUri, set.InComRegNewPath)
}

func (r *ReqInRegularNew) SendToServerRef() (ok bool, code int) {
	return helpers.SendToServerPostWithOutResp(r, set.InComBaseUri, set.InComRegRefPath)
}

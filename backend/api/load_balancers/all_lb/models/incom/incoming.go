package incom

import (
	"all_lb/helpers"
	"all_lb/set"
	"encoding/json"
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

type ReqInComingNew struct {
	UserId   string
	Category string
	Emoji    string
	Not      string
	Amount   float64
	Time     string
	Date     string
	From     FromUser
}

type FromUser struct {
	Name   string
	UserId string
}

func (r *ReqInComingNew) ToJson() (data []byte) {
	var err error
	data, err = json.Marshal(r)
	helpers.CheckErr(err)
	return
}

func (r *ReqInComingNew) FromJson(data *[]byte) bool {
	err := json.Unmarshal(*data, r)
	return helpers.CheckErr(err)
}

func (r *ReqInComingNew) SendToServer() (ok bool, code int) {
	return helpers.SendToServerPostWithOutResp(r, set.InComBaseUri, set.InComNewPath)
}

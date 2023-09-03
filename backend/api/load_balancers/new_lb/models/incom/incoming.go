package incom

import (
	"encoding/json"
	"new_lb/helpers"
	"new_lb/set"
)

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

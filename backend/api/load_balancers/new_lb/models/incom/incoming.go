package incom

import (
	"encoding/json"
	"new_lb/helpers"
	"new_lb/set"
)

type ReqInComing struct {
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

func (r *ReqInComing) ToJson() (data []byte) {
	var err error
	data, err = json.Marshal(r)
	helpers.CheckErr(err)
	return
}

func (r *ReqInComing) FromJson(data *[]byte) bool {
	err := json.Unmarshal(*data, r)
	return helpers.CheckErr(err)
}

func (r *ReqInComing) SendToServer() (ok bool, code int) {
	return helpers.SendToServerPostWithOutResp(r, set.InComBaseUri, set.InComNewPath)
}

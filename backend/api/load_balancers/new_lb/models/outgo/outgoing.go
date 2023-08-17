package outgo

import (
	"encoding/json"
	"new_lb/helpers"
	"new_lb/set"
)

type ReqOutgoing struct {
	UserId   string
	Category string
	Emoji    string
	Not      string
	Amount   float64
	Time     string
	Date     string
	To       FromUser
}

type FromUser struct {
	Name   string
	UserId string
}

func (r *ReqOutgoing) ToJson() (data []byte) {
	var err error
	data, err = json.Marshal(r)
	helpers.CheckErr(err)
	return
}

func (r *ReqOutgoing) FromJson(data *[]byte) bool {
	err := json.Unmarshal(*data, r)
	return helpers.CheckErr(err)
}

func (r *ReqOutgoing) SendToServer() (ok bool, code int) {
	return helpers.SendToServerPostWithOutResp(r, set.OutGoBaseUri, set.OutGoNewPath)
}

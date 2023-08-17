package models

import (
	"encoding/json"
	"sign_lb/helpers"
)

type Resp struct {
	Status   bool
	ReqId    string
	Message  []byte
	HttpCode int
}

func (r *Resp) ToJson(status bool, reqId string, code int) []byte {
	r.HttpCode = code
	r.ReqId = reqId
	r.Status = status
	data, err := json.Marshal(r)
	helpers.CheckErr(err)
	return data
}

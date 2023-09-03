package models

import (
	"all_lb/helpers"
	"encoding/json"
)

type RespSign struct {
	Status   bool
	ReqId    string
	Message  string
	HttpCode int
}

func (r *RespSign) ToJson(status bool, reqId string, code int) []byte {
	r.HttpCode = code
	r.ReqId = reqId
	r.Status = status
	data, err := json.Marshal(r)
	helpers.CheckErr(err)
	return data
}

type Resp struct {
	Status   bool
	ReqId    string
	Message  string
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

func (r *Resp) FromJson(data *[]byte) (ok bool) {
	err := json.Unmarshal(*data, r)
	if !helpers.CheckErr(err) {
		return false
	}

	return true
}

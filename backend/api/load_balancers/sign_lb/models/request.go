package models

import (
	"encoding/json"
	"sign_lb/helpers"
)

type Req struct {
	ReqId   string
	ReqType string
	Auth    string
	Message []byte
}

func (r *Req) FromJson(data []byte) bool {
	err := json.Unmarshal(data, r)
	return helpers.CheckErr(err)
}

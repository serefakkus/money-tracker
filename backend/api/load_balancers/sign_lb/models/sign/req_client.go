package sign

import (
	"encoding/json"
	"sign_lb/helpers"
)

type ReqSign struct {
	ReqId   string
	ReqType string
	Message []byte
}

func (r *ReqSign) FromJson(data []byte) bool {
	err := json.Unmarshal(data, r)
	return helpers.CheckErr(err)
}

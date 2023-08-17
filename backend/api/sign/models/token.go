package models

import (
	"encoding/json"
	"net/http"
	"sign/helpers"
	"sign/set"
)

type ReqNewToken struct {
	UserId string
}

func (r *ReqNewToken) toJson(data *[]byte) {
	var err error
	*data, err = json.Marshal(r)
	helpers.CheckErr(err)
}

type TokenDetails struct {
	AccessToken  string
	RefreshToken string
	AccessUuid   string
	RefreshUuid  string
	AtExpires    int64
	RtExpires    int64
}

func (r *TokenDetails) fromJson(data *[]byte) bool {
	err := json.Unmarshal(*data, r)
	return helpers.CheckErr(err)
}

func (r *ReqNewToken) getNewToken(userId *string, details *TokenDetails) (ok bool, code int) {
	r.UserId = *userId
	var message []byte
	r.toJson(&message)

	code, ok = helpers.SendToDbPost(&message, set.TokenUri)

	if !ok {
		return
	}

	if !details.fromJson(&message) {
		code = http.StatusBadGateway
		ok = false
		return
	}

	code = http.StatusOK
	return
}

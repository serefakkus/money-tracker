package models

import (
	"encoding/json"
	"get_lb/helpers"
	"net/http"
)

//---------------------------token req for query for token service

type TokenReq struct {
	Auth string
}

func (t *TokenReq) toJson() (data []byte) {
	var err error
	data, err = json.Marshal(t)
	helpers.CheckErr(err)
	return
}

func (t *TokenReq) Get(userId *string) (ok bool, code int) {

	var uri = helpers.GetUriToken("is_ok")

	var resp tokenResp

	ok, code = helpers.SendToServerGet(&resp, uri, t.Auth)

	if !ok {
		return false, code
	}

	*userId = resp.UserId

	if code != http.StatusOK {
		return false, code
	}

	return true, code
}

func (t *TokenReq) Ref(refTokDetails *TokenDetails) (ok bool, code int) {

	var uri = helpers.GetUriToken("ref")

	ok, code = helpers.SendToServerGet(refTokDetails, uri, t.Auth)

	if !ok {
		return false, code
	}

	if code != http.StatusOK {
		return false, code
	}

	return true, code
}

func (t *TokenReq) Del() (ok bool, code int) {
	var refTok tokenResp

	var uri = helpers.GetUriToken("del")

	ok, code = helpers.SendToServerGet(&refTok, uri, t.Auth)

	if !ok {
		return false, code
	}

	if code != http.StatusOK {
		return false, code
	}

	return true, code
}

//---------------------------token resp for result from token service

type tokenResp struct {
	UserId string
}

func (t *tokenResp) FromJson(data *[]byte) (ok bool) {
	err := json.Unmarshal(*data, t)
	return helpers.CheckErr(err)
}

//---------------------------refreshed token result from token service

//---------------------------token details for refresh token result

type TokenDetails struct {
	AtExpires    int64  `json:"at_expires"`
	RtExpires    int64  `json:"rt_expires"`
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	AccessUuid   string `json:"access_uuid"`
	RefreshUuid  string `json:"refresh_uuid"`
}

func (t *TokenDetails) ToJson() []byte {
	data, err := json.Marshal(t)

	helpers.CheckErr(err)

	return data
}

func (t *TokenDetails) FromJson(data *[]byte) (ok bool) {
	err := json.Unmarshal(*data, t)
	return helpers.CheckErr(err)
}

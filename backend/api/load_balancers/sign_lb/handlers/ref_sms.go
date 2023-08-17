package handlers

import (
	"encoding/json"
	"net/http"
	"sign_lb/helpers"
	"sign_lb/models"
	"sign_lb/models/sign"
)

func refNewSms(req *sign.ReqSign) (data []byte) {
	var sendReq sign.ReqNewSms
	var resp models.Resp
	err := json.Unmarshal(req.Message, &sendReq)
	if !helpers.CheckErr(err) {
		return resp.ToJson(false, req.ReqId, http.StatusInternalServerError)
	}

	if !helpers.PhoneControl(&sendReq.Phone) {
		return resp.ToJson(false, req.ReqId, http.StatusRequestedRangeNotSatisfiable)
	}

	var respSign sign.RespNewSms

	ok, code := sendReq.SendToServer(&respSign, true)
	if !ok {
		return resp.ToJson(false, req.ReqId, code)
	}

	resp.Message, _ = json.Marshal(respSign)

	return resp.ToJson(true, req.ReqId, http.StatusOK)
}

func refAskSms(req *sign.ReqSign) (data []byte) {
	var sendReq sign.ReqPhoneCode
	var resp models.Resp
	err := json.Unmarshal(req.Message, &sendReq)
	if !helpers.CheckErr(err) {
		return resp.ToJson(false, req.ReqId, http.StatusInternalServerError)
	}

	if !helpers.PhoneControl(&sendReq.Phone) {
		return resp.ToJson(false, req.ReqId, http.StatusRequestedRangeNotSatisfiable)
	}

	ok, code := sendReq.SendToServer(true)
	if !ok {
		return resp.ToJson(false, req.ReqId, code)
	}

	return resp.ToJson(true, req.ReqId, http.StatusOK)
}

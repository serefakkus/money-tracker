package sign

import (
	"all_lb/helpers"
	"all_lb/models"
	"all_lb/models/sign"
	"encoding/json"
	"net/http"
)

func newSms(req *sign.ReqSign) (data []byte) {
	var sendReq sign.ReqNewSms
	var resp models.RespSign
	err := json.Unmarshal(req.Message, &sendReq)
	if !helpers.CheckErr(err) {
		return resp.ToJson(false, req.ReqId, http.StatusInternalServerError)
	}

	if !helpers.PhoneControl(&sendReq.Phone) {
		return resp.ToJson(false, req.ReqId, http.StatusRequestedRangeNotSatisfiable)
	}

	ok, code := sendReq.SendToServer(&resp, false)
	if !ok {
		return resp.ToJson(false, req.ReqId, code)
	}

	return resp.ToJson(true, req.ReqId, http.StatusOK)
}

func askSms(req *sign.ReqSign) (data []byte) {
	var sendReq sign.ReqPhoneCode
	var resp models.RespSign
	err := json.Unmarshal(req.Message, &sendReq)
	if !helpers.CheckErr(err) {
		return resp.ToJson(false, req.ReqId, http.StatusInternalServerError)
	}

	if !helpers.PhoneControl(&sendReq.Phone) {
		return resp.ToJson(false, req.ReqId, http.StatusRequestedRangeNotSatisfiable)
	}

	ok, code := sendReq.SendToServer(false)
	if !ok {
		return resp.ToJson(false, req.ReqId, code)
	}

	return resp.ToJson(true, req.ReqId, http.StatusOK)
}

package sign

import (
	"all_lb/helpers"
	"all_lb/models"
	"all_lb/models/sign"
	"encoding/json"
	"net/http"
)

func signIn(req *sign.ReqSign, ip *sign.Ip) (data []byte) {

	var sendReq sign.ReqSignIn
	var resp models.RespSign
	err := json.Unmarshal(req.Message, &sendReq)
	if !helpers.CheckErr(err) {
		return resp.ToJson(false, req.ReqId, http.StatusInternalServerError)
	}

	sendReq.Ip = *ip

	if !helpers.SignInControl(&sendReq.Phone, &sendReq.Pass) {
		return resp.ToJson(false, req.ReqId, http.StatusRequestedRangeNotSatisfiable)
	}

	ok, code := sendReq.SendToServer(&resp)
	if !ok {
		return resp.ToJson(false, req.ReqId, code)
	}

	return resp.ToJson(true, req.ReqId, http.StatusOK)
}

func signUp(req *sign.ReqSign, ip *sign.Ip) (data []byte) {

	var sendReq sign.ReqSignUp
	var resp models.RespSign

	err := json.Unmarshal(req.Message, &sendReq)

	if !helpers.CheckErr(err) {
		return resp.ToJson(false, req.ReqId, http.StatusInternalServerError)
	}

	sendReq.Ip = *ip

	if !helpers.SignUpControl(&sendReq.Phone, &sendReq.Pass, &sendReq.Code) {
		return resp.ToJson(false, req.ReqId, http.StatusRequestedRangeNotSatisfiable)
	}

	ok, code := sendReq.SendToServer(&resp, false)

	if !ok {
		return resp.ToJson(false, req.ReqId, code)
	}

	return resp.ToJson(true, req.ReqId, http.StatusOK)
}

func refPass(req *sign.ReqSign, ip *sign.Ip) (data []byte) {

	var sendReq sign.ReqSignUp
	var resp models.RespSign

	err := json.Unmarshal(req.Message, &sendReq)

	if !helpers.CheckErr(err) {
		return resp.ToJson(false, req.ReqId, http.StatusInternalServerError)
	}

	sendReq.Ip = *ip

	if !helpers.SignUpControl(&sendReq.Phone, &sendReq.Pass, &sendReq.Code) {
		return resp.ToJson(false, req.ReqId, http.StatusRequestedRangeNotSatisfiable)
	}

	ok, code := sendReq.SendToServer(&resp, true)

	if !ok {
		return resp.ToJson(false, req.ReqId, code)
	}

	return resp.ToJson(true, req.ReqId, http.StatusOK)
}

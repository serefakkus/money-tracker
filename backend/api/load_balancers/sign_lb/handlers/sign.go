package handlers

import (
	"encoding/json"
	"net/http"
	"sign_lb/helpers"
	"sign_lb/models"
	"sign_lb/models/sign"
)

func signIn(req *sign.ReqSign, ip *sign.Ip) (data []byte) {

	var sendReq sign.ReqSignIn
	var resp models.Resp
	err := json.Unmarshal(req.Message, &sendReq)
	if !helpers.CheckErr(err) {
		return resp.ToJson(false, req.ReqId, http.StatusInternalServerError)
	}

	sendReq.Ip = *ip

	if !helpers.SignInControl(&sendReq.Phone, &sendReq.Pass) {
		return resp.ToJson(false, req.ReqId, http.StatusRequestedRangeNotSatisfiable)
	}

	var respSign sign.RespSignIn

	ok, code := sendReq.SendToServer(&respSign)
	if !ok {
		return resp.ToJson(false, req.ReqId, code)
	}

	resp.Message, _ = json.Marshal(respSign)

	return resp.ToJson(true, req.ReqId, http.StatusOK)
}

func signUp(req *sign.ReqSign, ip *sign.Ip) (data []byte) {

	var sendReq sign.ReqSignUp
	var resp models.Resp

	err := json.Unmarshal(req.Message, &sendReq)

	if !helpers.CheckErr(err) {
		return resp.ToJson(false, req.ReqId, http.StatusInternalServerError)
	}

	sendReq.Ip = *ip

	if !helpers.SignUpControl(&sendReq.Phone, &sendReq.Pass, &sendReq.Code) {
		return resp.ToJson(false, req.ReqId, http.StatusRequestedRangeNotSatisfiable)
	}

	var respSign sign.RespSignIn
	ok, code := sendReq.SendToServer(&respSign, false)

	if !ok {
		return resp.ToJson(false, req.ReqId, code)
	}

	resp.Message, _ = json.Marshal(respSign)

	return resp.ToJson(true, req.ReqId, http.StatusOK)
}

func refPass(req *sign.ReqSign, ip *sign.Ip) (data []byte) {

	var sendReq sign.ReqSignUp
	var resp models.Resp

	err := json.Unmarshal(req.Message, &sendReq)

	if !helpers.CheckErr(err) {
		return resp.ToJson(false, req.ReqId, http.StatusInternalServerError)
	}

	sendReq.Ip = *ip

	if !helpers.SignUpControl(&sendReq.Phone, &sendReq.Pass, &sendReq.Code) {
		return resp.ToJson(false, req.ReqId, http.StatusRequestedRangeNotSatisfiable)
	}

	var respSign sign.RespSignIn
	ok, code := sendReq.SendToServer(&respSign, true)

	if !ok {
		return resp.ToJson(false, req.ReqId, code)
	}

	resp.Message, _ = json.Marshal(respSign)

	return resp.ToJson(true, req.ReqId, http.StatusOK)
}

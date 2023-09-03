package incomingNew

import (
	"all_lb/models"
	"all_lb/models/incom"
	"net/http"
)

func newInComing(req *models.Req, userId *string) (data []byte) {
	var resp models.Resp
	var reqIn incom.ReqInComingNew

	if !reqIn.FromJson(&req.Message) {
		return resp.ToJson(false, req.ReqId, http.StatusBadRequest)
	}

	reqIn.UserId = *userId

	ok, code := reqIn.SendToServer()

	return resp.ToJson(ok, req.ReqId, code)
}

func regNewInComing(req *models.Req, userId *string) (data []byte) {

	var resp models.Resp
	var reqRegIn incom.ReqInRegularNew

	if !reqRegIn.FromJson(&req.Message) {
		return resp.ToJson(false, req.ReqId, http.StatusBadRequest)
	}

	reqRegIn.UserId = *userId

	ok, code := reqRegIn.SendToServerNew()

	return resp.ToJson(ok, req.ReqId, code)
}

func regRefInComing(req *models.Req, userId *string) (data []byte) {
	var resp models.Resp
	var reqRegIn incom.ReqInRegularNew

	if !reqRegIn.FromJson(&req.Message) {
		return resp.ToJson(false, req.ReqId, http.StatusBadRequest)
	}

	reqRegIn.UserId = *userId

	ok, code := reqRegIn.SendToServerRef()

	return resp.ToJson(ok, req.ReqId, code)
}

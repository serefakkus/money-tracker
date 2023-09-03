package outgoingNew

import (
	"net/http"
	"new_lb/models"
	"new_lb/models/outgo"
)

func newOutGoing(req *models.Req, userId *string) (data []byte) {
	var resp models.Resp
	var reqIn outgo.ReqOutgoingNew

	if !reqIn.FromJson(&req.Message) {
		return resp.ToJson(false, req.ReqId, http.StatusBadRequest)
	}

	reqIn.UserId = *userId

	ok, code := reqIn.SendToServer()

	return resp.ToJson(ok, req.ReqId, code)
}

func regNewOutGoing(req *models.Req, userId *string) (data []byte) {
	var resp models.Resp
	var reqRegIn outgo.ReqOutRegularNew

	if !reqRegIn.FromJson(&req.Message) {
		return resp.ToJson(false, req.ReqId, http.StatusBadRequest)
	}

	reqRegIn.UserId = *userId

	ok, code := reqRegIn.SendToServerNew()

	return resp.ToJson(ok, req.ReqId, code)
}

func regRefOutGoing(req *models.Req, userId *string) (data []byte) {
	var resp models.Resp
	var reqRegIn outgo.ReqOutRegularNew

	if !reqRegIn.FromJson(&req.Message) {
		return resp.ToJson(false, req.ReqId, http.StatusBadRequest)
	}

	reqRegIn.UserId = *userId

	ok, code := reqRegIn.SendToServerRef()

	return resp.ToJson(ok, req.ReqId, code)
}

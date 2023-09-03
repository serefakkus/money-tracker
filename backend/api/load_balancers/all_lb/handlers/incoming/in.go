package incoming

import (
	"all_lb/models"
	"all_lb/models/in"
	"net/http"
)

func askIncoming(req *models.Req, userId *string) (data []byte) {
	var resp models.Resp
	var incoming in.ReqIncoming

	if !incoming.FromJson(&req.Message) {
		return resp.ToJson(false, req.ReqId, http.StatusBadRequest)
	}

	incoming.UserId = *userId

	ok, code := incoming.SendToServer(&resp)

	return resp.ToJson(ok, req.ReqId, code)
}

func askRegIncoming(req *models.Req, userId *string) (data []byte) {
	var resp models.Resp
	var incoming in.ReqRegularIn

	if !incoming.FromJson(&req.Message) {
		return resp.ToJson(false, req.ReqId, http.StatusBadRequest)
	}

	incoming.UserId = *userId

	ok, code := incoming.SendToServer(&resp)

	return resp.ToJson(ok, req.ReqId, code)
}

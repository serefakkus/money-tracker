package outgoing

import (
	"all_lb/models"
	"all_lb/models/out"
	"net/http"
)

func askOutgoing(req *models.Req, userId *string) (data []byte) {
	var resp models.Resp
	var outgoing out.ReqOutgo

	if !outgoing.FromJson(&req.Message) {
		return resp.ToJson(false, req.ReqId, http.StatusBadRequest)
	}

	outgoing.UserId = *userId

	ok, code := outgoing.SendToServer(&resp)

	return resp.ToJson(ok, req.ReqId, code)
}

func askRegOutgoing(req *models.Req, userId *string) (data []byte) {
	var resp models.Resp
	var outgoing out.ReqRegularOut

	if !outgoing.FromJson(&req.Message) {
		return resp.ToJson(false, req.ReqId, http.StatusBadRequest)
	}

	outgoing.UserId = *userId

	ok, code := outgoing.SendToServer(&resp)

	return resp.ToJson(ok, req.ReqId, code)
}

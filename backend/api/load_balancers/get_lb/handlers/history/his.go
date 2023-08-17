package history

import (
	"get_lb/models"
	"get_lb/models/his"
	"net/http"
)

func askHistory(req *models.Req, userId *string) (data []byte) {
	var resp models.Resp
	var history his.ReqHis

	if !history.FromJson(&req.Message) {
		return resp.ToJson(false, req.ReqId, http.StatusBadRequest)
	}

	history.UserId = *userId

	ok, code := history.SendToServer(&resp)

	return resp.ToJson(ok, req.ReqId, code)
}

func askInRegHistory(req *models.Req, userId *string) (data []byte) {
	var resp models.Resp
	var history his.ReqRegHis

	if !history.FromJson(&req.Message) {
		return resp.ToJson(false, req.ReqId, http.StatusBadRequest)
	}

	history.UserId = *userId

	ok, code := history.SendToServerIn(&resp)

	return resp.ToJson(ok, req.ReqId, code)
}

func askOutRegHistory(req *models.Req, userId *string) (data []byte) {
	var resp models.Resp
	var history his.ReqRegHis

	if !history.FromJson(&req.Message) {
		return resp.ToJson(false, req.ReqId, http.StatusBadRequest)
	}

	history.UserId = *userId

	ok, code := history.SendToServerOut(&resp)

	return resp.ToJson(ok, req.ReqId, code)
}

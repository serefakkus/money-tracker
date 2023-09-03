package history

import (
	"all_lb/models"
	"all_lb/models/his"
	"log"
	"net/http"
)

func askHistory(req *models.Req, userId *string) (data []byte) {
	var resp models.Resp
	var history his.ReqHis

	history.UserId = *userId

	ok, code := history.SendToServer(&resp)

	log.Println("resp = " + resp.Message)

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

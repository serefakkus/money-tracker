package incomingNew

import (
	"net/http"
	"new_lb/handlers/token"
	"new_lb/models"
)

func RouterInComingNew(req *models.Req) (data []byte) {

	var userId string
	var resp models.Resp

	ok, code := token.IsTokenOk(&req.Auth, &userId)

	if !ok {
		return resp.ToJson(false, req.ReqId, code)
	}

	switch req.ReqType {

	case "is_ok":

		return resp.ToJson(true, req.ReqId, http.StatusOK)

	case "ref_tok":

		return token.TokenRef(req, &resp)

	case "del_tok":

		return token.TokenDel(req)

	case "in":
		return newInComing(req, &userId)

	case "reg_new":

		return regNewInComing(req, &userId)

	case "reg_ref":

		return regRefInComing(req, &userId)

	default:
		return resp.ToJson(false, req.ReqId, http.StatusBadRequest)
	}

}

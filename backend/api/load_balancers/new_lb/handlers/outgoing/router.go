package outgoing

import (
	"net/http"
	"new_lb/handlers/token"
	"new_lb/models"
)

func RouterOutGoing(req *models.Req) (data []byte) {

	var userId string
	var resp models.Resp

	ok, code := token.IsTokenOk(&req.Auth, &userId)

	if !ok {
		return resp.ToJson(false, req.ReqId, code)
	}

	switch req.ReqType {

	case "out":

		return newOutGoing(req, &userId)

	case "reg_new":

		return regNewOutGoing(req, &userId)

	case "reg_ref":

		return regRefOutGoing(req, &userId)

	default:

		return resp.ToJson(false, req.ReqId, http.StatusBadRequest)
	}

}

package outgoing

import (
	"all_lb/handlers/token"
	"all_lb/models"
	"net/http"
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

		return askOutgoing(req, &userId)

	case "reg_out":

		return askRegOutgoing(req, &userId)

	default:

		return resp.ToJson(false, req.ReqId, http.StatusBadRequest)
	}

}

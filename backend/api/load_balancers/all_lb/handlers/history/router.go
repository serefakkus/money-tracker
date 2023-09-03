package history

import (
	"all_lb/handlers/token"
	"all_lb/models"
	"net/http"
)

func RouterHistory(req *models.Req) (data []byte) {

	var userId string
	var resp models.Resp

	ok, code := token.IsTokenOk(&req.Auth, &userId)

	if !ok {
		return resp.ToJson(false, req.ReqId, code)
	}

	switch req.ReqType {

	case "his":

		return askHistory(req, &userId)

	case "in_reg_his":

		return askInRegHistory(req, &userId)

	case "our_reg_his":

		return askOutRegHistory(req, &userId)

	default:
		return resp.ToJson(false, req.ReqId, http.StatusBadRequest)
	}

}

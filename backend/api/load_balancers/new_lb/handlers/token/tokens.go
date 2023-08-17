package token

import "new_lb/models"

func IsTokenOk(auth *string, userId *string) (ok bool, code int) {
	var token models.TokenReq

	token.Auth = *auth

	return token.Get(userId)
}

func TokenRef(req *models.Req, resp *models.Resp) (data []byte) {
	var token models.TokenReq
	var code int

	token.Auth = req.Auth
	var refTokData models.TokenDetails

	resp.Status, code = token.Ref(&refTokData)

	resp.Message = refTokData.ToJson()

	return resp.ToJson(resp.Status, req.ReqId, code)
}

func TokenDel(req *models.Req) (data []byte) {
	var code int
	var token models.TokenReq

	token.Auth = req.Auth

	var resp models.Resp
	resp.Status, code = token.Del()

	return resp.ToJson(resp.Status, req.ReqId, code)
}

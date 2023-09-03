package sign

import (
	"all_lb/models"
	"all_lb/models/sign"
	"net/http"
)

func RouterSign(req *sign.ReqSign, ip *sign.Ip) (data []byte) {

	//var resp models.Resp

	switch req.ReqType {

	case "login":
		return signIn(req, ip)

	case "new_sms":

		return newSms(req)

	case "ask_sms":

		return askSms(req)

	case "signup":

		return signUp(req, ip)

	case "new_sms_ref":

		return refNewSms(req)

	case "ask_sms_ref":

		return refAskSms(req)

	case "ref_pass":

		return refPass(req, ip)

	default:

		var resp models.Resp
		resp.HttpCode = http.StatusNotAcceptable
		return
	}
}

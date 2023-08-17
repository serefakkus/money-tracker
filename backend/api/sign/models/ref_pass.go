package models

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"sign/dataloaders/mongo"
	"sign/helpers"
	"sign/set"
)

func (r *ReqSignUp) refPass(details *TokenDetails) (code int) {
	if !helpers.CheckPass(r.Pass) {
		return http.StatusRequestedRangeNotSatisfiable
	}

	var sms ReqPhoneCode
	sms.Phone = r.Phone
	sms.Code = r.Code

	var ok bool

	code, ok = sms.askCode()
	if !ok {
		return
	}

	var user MongoSignInfo
	var empty bool

	ok, empty = user.isExistUser(r.Phone)

	if empty {
		return http.StatusNotFound
	}

	var refPass MongoSignRefPass
	var filter filterPhoneMongo
	filter.Phone = r.Phone
	refPass.refPass(r.Pass, r.Ip)

	if !mongo.RefMongo(&filter, &refPass, &set.SignTableName, &set.SignDbName) {
		return http.StatusBadGateway
	}

	var token ReqNewToken

	ok, code = token.getNewToken(&user.UserId, details)

	if !ok {
		return
	}

	return http.StatusOK

}

func (r *ReqSignUp) HandlerRefPass(c *gin.Context) {
	err := c.ShouldBindJSON(r)
	if !helpers.CheckErr(err) {
		c.Writer.WriteHeader(http.StatusBadRequest)
		return
	}

	var resp TokenDetails

	c.IndentedJSON(r.refPass(&resp), resp)
}

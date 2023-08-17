package models

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"sign/dataloaders/mongo"
	"sign/helpers"
	"sign/set"
)

type ReqNewSms struct {
	Phone string
}

type RespNewSms struct {
	Empty   bool
	TimeOut string
}

func (r *ReqNewSms) isNotExistUser(resp *RespNewSms, filter *filterSignPhone) (code int, ok bool, isExist bool) {
	var user MongoSignInfo

	filter.Phone = r.Phone

	ok, empty := mongo.SorMongo(filter, &user, &set.SignTableName, &set.SignDbName)

	if !ok {
		err := errors.New("sormongo error")
		helpers.CheckErr(err)
		code = http.StatusInternalServerError
		return
	}

	if empty {
		resp.Empty = false
		isExist = false
		ok = true
		return
	}

	resp.Empty = true
	isExist = true
	ok = true

	return
}

func (r *ReqNewSms) sendToSmsForNew(resp *RespNewSms) (code int) { //send sms for new user
	if !helpers.CheckPhoneNo(r.Phone) {
		return http.StatusRequestedRangeNotSatisfiable
	}

	var filter filterSignPhone
	var result SmsLogSql
	var ok, isExist bool

	result.Phone = r.Phone

	code, ok, isExist = r.isNotExistUser(resp, &filter)

	if !ok {
		return
	}

	if isExist {
		code = http.StatusAlreadyReported
		return
	}

	code, ok = result.getFromDb(&filter, resp)

	if !ok {
		return
	}

	var sms SmsSql

	ok, code = sms.insertNewSms(&filter, &r.Phone)

	if !ok {
		return
	}

	if !helpers.SendSms(&sms.Phone, &sms.Code) {
		code = http.StatusInternalServerError
		return
	}

	resp.Empty = true

	return http.StatusOK
}

func (r *ReqNewSms) sendToSmsRefPass(resp *RespNewSms) (code int) { //send sms for new user
	if !helpers.CheckPhoneNo(r.Phone) {
		return http.StatusRequestedRangeNotSatisfiable
	}

	var filter filterSignPhone
	var result SmsLogSql
	var ok, isExist bool

	result.Phone = r.Phone

	code, ok, isExist = r.isNotExistUser(resp, &filter)

	if !ok {
		return
	}

	if !isExist {
		code = http.StatusNotFound
		return
	}

	code, ok = result.getFromDb(&filter, resp)

	if !ok {
		return
	}

	var sms SmsSql

	ok, code = sms.insertNewSms(&filter, &r.Phone)

	if !ok {
		return
	}

	if !helpers.SendSms(&sms.Phone, &sms.Code) {
		code = http.StatusInternalServerError
		return
	}

	return http.StatusOK
}

func (r *ReqNewSms) HandlerNewSmsSignUp(c *gin.Context) {
	err := c.ShouldBindJSON(r)
	if !helpers.CheckErr(err) {
		c.Writer.WriteHeader(http.StatusBadRequest)
		return
	}

	var resp RespNewSms

	c.IndentedJSON(r.sendToSmsForNew(&resp), resp)
}

func (r *ReqNewSms) HandlerNewSmsRefPass(c *gin.Context) {
	err := c.ShouldBindJSON(r)
	if !helpers.CheckErr(err) {
		c.Writer.WriteHeader(http.StatusBadRequest)
		return
	}

	var resp RespNewSms

	c.IndentedJSON(r.sendToSmsRefPass(&resp), resp)
}

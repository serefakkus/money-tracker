package models

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"sign/dataloaders/sql"
	"sign/helpers"
	"sign/set"
	"time"
)

type ReqPhoneCode struct {
	Phone string
	Code  string
}

type ReqPhone struct {
	Phone string
}

func (r *ReqPhoneCode) askCode() (code int, ok bool) {
	if !helpers.CheckPhoneNo(r.Phone) || !helpers.CheckCode(r.Code) {
		code = http.StatusRequestedRangeNotSatisfiable
		return
	}

	var sms SmsSql
	var filter filterSignPhone
	var empty bool

	filter.Phone = r.Phone
	ok, empty = sql.SelectSql(&sms, &filter, &set.SmsSqlTableName)
	if !ok || empty || sms.Phone != r.Phone || sms.Code != r.Code || r.Code == "" {
		code = http.StatusUnauthorized
		ok = false
		return
	}

	if empty {
		code = http.StatusNotFound
		ok = false
		return
	}

	var date time.Time
	if !sms.getTime(&date) {
		code = http.StatusInternalServerError
		ok = false
		return
	}

	if !time.Now().Before(date) {
		code = http.StatusTooEarly
		ok = false
		return
	}

	ok = true
	code = http.StatusOK
	return
}

func (r *ReqPhoneCode) askCodeSignUp() (code int) {
	var ok, empty bool

	code, ok = r.askCode()

	if !ok {
		return code
	}

	var user MongoSignInfo

	ok, empty = user.isExistUser(r.Phone)
	if !ok {
		return http.StatusInternalServerError
	}

	if !empty {
		return http.StatusAlreadyReported
	}

	code = http.StatusOK
	return
}

func (r *ReqPhoneCode) askCodeRefPass() (code int) {
	var ok, empty bool

	code, ok = r.askCode()

	if !ok {
		return code
	}

	var user MongoSignInfo

	ok, empty = user.isExistUser(r.Phone)

	if !ok {
		return http.StatusInternalServerError
	}

	if empty {
		return http.StatusNotFound
	}

	code = http.StatusOK
	return
}

func (r *ReqPhoneCode) HandlerAskCodeSignUp(c *gin.Context) {
	err := c.ShouldBindJSON(r)
	if !helpers.CheckErr(err) {
		c.Writer.WriteHeader(http.StatusBadRequest)
		return
	}

	c.Writer.WriteHeader(r.askCodeSignUp())
}

func (r *ReqPhoneCode) HandlerAskCodeRefPass(c *gin.Context) {
	err := c.ShouldBindJSON(r)
	if !helpers.CheckErr(err) {
		c.Writer.WriteHeader(http.StatusBadRequest)
		return
	}

	c.Writer.WriteHeader(r.askCodeRefPass())
}

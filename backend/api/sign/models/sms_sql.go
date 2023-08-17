package models

import (
	"net/http"
	"sign/dataloaders/sql"
	"sign/helpers"
	"sign/set"
	"time"
)

type SmsSql struct {
	Phone string
	Code  string
	Time  string
}

func (s *SmsSql) GetPhone() string {
	return s.Phone
}

func (s *SmsSql) GetCode() string {
	return s.Code
}

func (s *SmsSql) generateCode() {
	s.Code = helpers.RandCodeRunes(5)
}

func (s *SmsSql) setTime() { //set time is after 15 minutes from now
	var timeAfter15Minutes = time.Now().Add(time.Minute * 15)
	helpers.TimeToString(&s.Time, &timeAfter15Minutes)
}

func (s *SmsSql) getTime(date *time.Time) (ok bool) {
	return helpers.StringToTime(&s.Time, date)
}

func (s *SmsSql) insertNewSms(filter *filterSignPhone, phone *string) (ok bool, code int) {

	s.Phone = *phone
	s.generateCode()
	s.setTime()

	var oldSmsSql SmsSql
	ok, empty := sql.SelectSql(&oldSmsSql, filter, &set.SmsSqlTableName)
	if !ok {
		code = http.StatusInternalServerError
		return
	}

	if empty {
		if !sql.InsertSql(s, &set.SmsSqlTableName) {
			code = http.StatusInternalServerError
			return
		}
		ok = true
		return
	}

	ok, empty = sql.UpdateSql(s, filter, &set.SmsSqlTableName)
	if !ok {
		code = http.StatusInternalServerError
		return
	}

	if empty {
		if !sql.InsertSql(s, &set.SmsSqlTableName) {
			code = http.StatusInternalServerError
			return
		}

		ok = true
		return
	}

	return
}

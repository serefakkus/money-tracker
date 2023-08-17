package models

import (
	"errors"
	"github.com/lib/pq"
	"net/http"
	"sign/dataloaders/sql"
	"sign/helpers"
	"sign/set"
	"time"
)

type SmsLogSql struct {
	Phone string
	Times []string
}

type smsLogSqlP struct {
	Phone string
	Times pq.StringArray
}

func (s *SmsLogSql) setTime() (added bool) { //set time is after 15 minutes from now

	var now = time.Now()
	var timeAfter15Minutes = time.Now().Add(time.Minute * 15)
	var date time.Time

	if s.Times == nil {
		var list []string
		for i := 0; i < 3; i++ {
			list = append(list, "")
		}
		s.Times = list
	}

	if len(s.Times) != 3 {
		var list []string
		for i := 0; i < 3; i++ {
			list = append(list, "")
		}
		s.Times = list
	}

	var errCount = 0

	for i := 0; i < 3; i++ {

		if s.Times[i] == "" {
			helpers.TimeToString(&s.Times[i], &timeAfter15Minutes)
			added = true
			continue
		}

		if !helpers.StringToTime(&s.Times[i], &date) {
			errCount++
			s.Times[i] = ""
			continue
		}

		if now.After(date) {
			if added {
				s.Times[i] = ""
				continue
			}
			helpers.TimeToString(&s.Times[i], &timeAfter15Minutes)
			added = true
			continue
		}
	}

	if errCount == 3 {
		helpers.TimeToString(&s.Times[0], &timeAfter15Minutes)
		added = true
	}

	return added
}

func (s *SmsLogSql) getFromDb(filter *filterSignPhone, resp *RespNewSms) (code int, ok bool) { //ask sms send limit if limit is not full insert new log
	var empty bool
	var postgres smsLogSqlP
	ok, empty = sql.SelectSql(&postgres, filter, &set.SmsSqlLogTableName)
	if !ok {
		err := errors.New("selectsql error")
		helpers.CheckErr(err)
		code = http.StatusInternalServerError
		return
	}

	s.Times = postgres.Times

	if empty {
		s.setTime()
		if !sql.InsertSql(s, &set.SmsSqlLogTableName) {
			err := errors.New("insert smslog sql error")
			helpers.CheckErr(err)
			code = http.StatusInternalServerError
			return
		}
	}

	if !empty {
		if !s.setTime() {

			if !s.getTime(&resp.TimeOut) {
				err := errors.New("sms sql log get time error")
				helpers.CheckErr(err)
				code = http.StatusInternalServerError
				return
			}

			code = http.StatusTooEarly
			ok = false
			return
		}

		ok, empty = sql.UpdateSql(s, &filter, &set.SmsSqlLogTableName)

		if !ok {
			err := errors.New("sms sql log update error")
			helpers.CheckErr(err)
		}

		if empty {
			err := errors.New("sms sql log update is empty error")
			helpers.CheckErr(err)
		}
	}

	code = http.StatusOK
	ok = true
	return
}

func (s *SmsLogSql) getTime(date *string) (ok bool) {
	var timeList []time.Time
	var ch = make(chan bool)

	for i := range s.Times {
		go helpers.StringListToDateList(&s.Times[i], &timeList, ch)
	}

	for range s.Times {
		switch <-ch {

		case true:

		case false:

		}
	}

	if len(timeList) < 3 {
		return
	}

	var lastDate = timeList[0]

	for i := range timeList {
		if lastDate.Before(timeList[i]) {
			lastDate = timeList[i]
		}
	}

	if lastDate.Before(time.Now()) {
		ok = false
		return
	}

	helpers.TimeToString(date, &lastDate)

	return true
}

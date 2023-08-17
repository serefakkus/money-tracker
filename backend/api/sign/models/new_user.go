package models

import (
	"encoding/json"
	"errors"
	"sign/helpers"
	"sign/set"
	"strconv"
)

type ReqNewUser struct {
	UserId string
}

func (r *ReqNewUser) ToJson() (data []byte) {
	var err error
	data, err = json.Marshal(r)
	helpers.CheckErr(err)
	return
}

func (r *ReqNewUser) SendToDb(userId *string) bool {
	r.UserId = *userId

	var message = r.ToJson()

	code, ok := helpers.SendToDbPost(&message, set.NewUserUri)

	if !ok {
		err := errors.New("User was created but new user cant created pay db !!!!!!!!!!!!!!!!" + " http code = " + strconv.Itoa(code))
		helpers.CheckErr(err)
		return false
	}

	return true
}

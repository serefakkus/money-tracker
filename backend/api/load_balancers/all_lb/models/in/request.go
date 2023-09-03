package in

import (
	"all_lb/helpers"
	"all_lb/models"
	"all_lb/set"
	"encoding/json"
)

type ReqIncoming struct {
	UserId     string
	IncomingId string
}

func (r *ReqIncoming) ToJson() (data []byte) {
	var err error
	data, err = json.Marshal(r)
	helpers.CheckErr(err)
	return
}

func (r *ReqIncoming) FromJson(data *[]byte) bool {
	err := json.Unmarshal(*data, r)
	return helpers.CheckErr(err)
}

func (r *ReqIncoming) SendToServer(resp *models.Resp) (ok bool, code int) {
	return helpers.SendToServerPost(r, &resp.Message, set.GetBaseUri, set.IncomingPath)
}

//--------------------------------------------

type ReqRegularIn struct {
	UserId    string
	RegularId string
	Offset    int
}

func (r *ReqRegularIn) ToJson() (data []byte) {
	var err error
	data, err = json.Marshal(r)
	helpers.CheckErr(err)
	return
}

func (r *ReqRegularIn) FromJson(data *[]byte) bool {
	err := json.Unmarshal(*data, r)
	return helpers.CheckErr(err)
}

func (r *ReqRegularIn) SendToServer(resp *models.Resp) (ok bool, code int) {
	return helpers.SendToServerPost(r, &resp.Message, set.GetBaseUri, set.RegularIncomingPath)
}

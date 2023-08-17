package his

import (
	"encoding/json"
	"get_lb/helpers"
	"get_lb/models"
	"get_lb/set"
)

type ReqHis struct {
	UserId string
}

func (r *ReqHis) ToJson() (data []byte) {
	var err error
	data, err = json.Marshal(r)
	helpers.CheckErr(err)
	return
}

func (r *ReqHis) FromJson(data *[]byte) bool {
	err := json.Unmarshal(*data, r)
	return helpers.CheckErr(err)
}

func (r *ReqHis) SendToServer(resp *models.Resp) (ok bool, code int) {
	return helpers.SendToServerPost(r, &resp.Message, set.GetBaseUri, set.HistoryPath)
}

//--------------------------------

type ReqRegHis struct {
	UserId    string
	RegularId string
}

func (r *ReqRegHis) ToJson() (data []byte) {
	var err error
	data, err = json.Marshal(r)
	helpers.CheckErr(err)
	return
}

func (r *ReqRegHis) FromJson(data *[]byte) bool {
	err := json.Unmarshal(*data, r)
	return helpers.CheckErr(err)
}

func (r *ReqRegHis) SendToServerIn(resp *models.Resp) (ok bool, code int) {
	return helpers.SendToServerPost(r, &resp.Message, set.GetBaseUri, set.ReqInHisPath)
}

func (r *ReqRegHis) SendToServerOut(resp *models.Resp) (ok bool, code int) {
	return helpers.SendToServerPost(r, &resp.Message, set.GetBaseUri, set.ReqOutHisPath)
}

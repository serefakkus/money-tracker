package out

import (
	"encoding/json"
	"get_lb/helpers"
	"get_lb/models"
	"get_lb/set"
)

type ReqOutgo struct {
	UserId     string
	OutgoingId string
}

func (r *ReqOutgo) ToJson() (data []byte) {
	var err error
	data, err = json.Marshal(r)
	helpers.CheckErr(err)
	return
}

func (r *ReqOutgo) FromJson(data *[]byte) bool {
	err := json.Unmarshal(*data, r)
	return helpers.CheckErr(err)
}

func (r *ReqOutgo) SendToServer(resp *models.Resp) (ok bool, code int) {
	return helpers.SendToServerPost(r, &resp.Message, set.GetBaseUri, set.OutgoingPath)
}

//--------------------------------------------

type ReqRegularOut struct {
	UserId    string
	RegularId string
	Offset    int
}

func (r *ReqRegularOut) ToJson() (data []byte) {
	var err error
	data, err = json.Marshal(r)
	helpers.CheckErr(err)
	return
}

func (r *ReqRegularOut) FromJson(data *[]byte) bool {
	err := json.Unmarshal(*data, r)
	return helpers.CheckErr(err)
}

func (r *ReqRegularOut) SendToServer(resp *models.Resp) (ok bool, code int) {
	return helpers.SendToServerPost(r, &resp.Message, set.GetBaseUri, set.RegularOutgoingPath)
}

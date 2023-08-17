package mongo_models

import (
	"get/dataloaders/mongo"
	"get/models/response"
	"get/set"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type RegularInHistoryMongo struct {
	Id          primitive.ObjectID `bson:"_id"`
	UserId      string
	Category    string
	RegularId   string
	IntervalDay int
	Time        string
	Ids         []RegHisIdMongo
}

func (r *RegularInHistoryMongo) AskMongo(userId *string, regularId *string) (ok bool, empty bool) {
	var filter filterRegID
	filter.UserId = *userId
	filter.RegularId = *regularId
	ok, empty = mongo.AskMongo(&filter, r, &set.RegularHistoryIncomingTableName, &set.UserHistoryDBName)
	return
}

func (r *RegularInHistoryMongo) ToResp(resp *response.RegHisInResp) {
	resp.RegularId = r.RegularId
	resp.IntervalDay = r.IntervalDay
	resp.Time = r.Time
	resp.Category = r.Category
	for i := range r.Ids {
		var id response.RegHisIdSResp
		id.Id = r.Ids[i].Id.Hex()
		id.Date = r.Ids[i].Date
		resp.IdS = append(resp.IdS, id)
	}
}

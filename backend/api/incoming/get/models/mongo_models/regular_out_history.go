package mongo_models

import (
	"get/dataloaders/mongo"
	"get/models/response"
	"get/set"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type RegularOutHistoryMongo struct {
	Id          primitive.ObjectID `bson:"_id"`
	UserId      string
	Category    string
	RegularId   string
	IntervalDay int
	Time        string
	IdS         []RegHisIdMongo
}

type RegHisIdMongo struct {
	Id   primitive.ObjectID
	Date string
}

func (r *RegularOutHistoryMongo) AskMongo(userId *string, regularId *string) (ok bool, empty bool) {
	var filter filterRegID
	filter.UserId = *userId
	filter.RegularId = *regularId
	ok, empty = mongo.AskMongo(&filter, r, &set.RegularHistoryOutgoingTableName, &set.UserHistoryDBName)
	return
}

func (r *RegularOutHistoryMongo) ToResp(resp *response.RegHisOutResp) {
	resp.RegularId = r.RegularId
	resp.IntervalDay = r.IntervalDay
	resp.Time = r.Time
	resp.Category = r.Category
	for i := range r.IdS {
		var id response.RegHisIdSResp
		id.Id = r.IdS[i].Id.Hex()
		id.Date = r.IdS[i].Date
		resp.IdS = append(resp.IdS, id)
	}
}

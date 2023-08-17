package mongo_models

import (
	"get/dataloaders/mongo"
	"get/models/response"
	"get/set"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type RegularInMongo struct {
	Id     primitive.ObjectID `bson:"_id"`
	Emoji  string
	Not    string
	Amount float64
	Time   string
	From   response.FromUser
}

func (r *RegularInMongo) GetMongo(MongoId *primitive.ObjectID) (ok bool, empty bool) {
	var filter FilterId
	filter.Id = *MongoId

	return mongo.AskMongo(&filter, r, &set.RegularIncomingTableName, &set.IncomingDBName)
}

func (r *RegularInMongo) ToResp(resp *response.RegInResp) {
	resp.Not = r.Not
	resp.Time = r.Time
	resp.Emoji = r.Emoji
	resp.Amount = r.Amount
	resp.From = r.From
}

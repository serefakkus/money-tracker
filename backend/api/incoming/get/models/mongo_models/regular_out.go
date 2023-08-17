package mongo_models

import (
	"get/dataloaders/mongo"
	"get/models/response"
	"get/set"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type RegularOutMongo struct {
	Id     primitive.ObjectID `bson:"_id"`
	Emoji  string
	Not    string
	Amount float64
	Time   string
	To     response.FromUser
}

func (r *RegularOutMongo) GetMongo(MongoId *primitive.ObjectID) (ok bool, empty bool) {
	var filter FilterId
	filter.Id = *MongoId

	return mongo.AskMongo(&filter, r, &set.RegularOutgoingTableName, &set.OutgoingTableName)
}

func (r *RegularOutMongo) ToResp(resp *response.RegOutResp) {
	resp.Emoji = r.Emoji
	resp.Not = r.Not
	resp.Time = r.Time
	resp.To = r.To
	resp.Amount = r.Amount
}

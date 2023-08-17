package mongo_models

import (
	"get/dataloaders/mongo"
	"get/set"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type DayMongo struct {
	Id  primitive.ObjectID `bson:"_id"`
	Ids []primitive.ObjectID
}

func (d *DayMongo) GetDays(mongoId *primitive.ObjectID, isInCom bool) (ok bool, empty bool) {
	var filter FilterId
	filter.Id = *mongoId
	if isInCom {
		return mongo.AskMongo(&filter, d, &set.IncomingDailyTableName, &set.UserHistoryDBName)
	}
	return mongo.AskMongo(&filter, d, &set.OutgoingDailyTableName, &set.UserHistoryDBName)
}

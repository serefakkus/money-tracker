package mongo_models

import (
	"get/models/response"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type IncomingMongo struct {
	Id        primitive.ObjectID `bson:"_id"`
	Type      bool               //true is regular & false is once
	RegularId string
	Category  string
	Emoji     string
	Not       string
	Amount    float64
	Time      string
	Date      string
	From      response.FromUser
}

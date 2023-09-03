package mongo_models

import "go.mongodb.org/mongo-driver/bson/primitive"

type FilterId struct {
	Id primitive.ObjectID `bson:"_id"`
}

type filterUser struct {
	UserId string
}

type filterRegID struct {
	UserId    string
	RegularId string
}

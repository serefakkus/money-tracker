package mongo

import (
	"context"
	"errors"
	"get/helpers"
	"get/interfaces"
	"get/set"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

func AskMongo(filterData interfaces.IDataLoader, resultData interfaces.IDataLoader, tableName *string, dbName *string) (ok bool, empty bool) {

	if !(helpers.GetModelLengt(filterData) > 0) {
		err := errors.New("filter data is empty")
		helpers.CheckErr(err)
		return false, false
	}

	clientOptions := options.Client().ApplyURI(set.MongoUri)

	// Connect to MongoDB
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, clientOptions)
	if !helpers.CheckErr(err) {
		return
	}

	collection := client.Database(*dbName).Collection(*tableName)

	data := collection.FindOne(context.TODO(), filterData)

	err = client.Disconnect(context.TODO())
	helpers.CheckErr(err)

	err = data.Decode(resultData)

	if err != nil {
		if err.Error() == "mongo: no documents in result" {
			return true, true
		}
		helpers.CheckErr(err)
		return
	}
	return true, false
}

package mongo

import (
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"sign/helpers"
	"sign/interfaces"
	"sign/set"
	"time"
)

func NewMongo(InsertData interfaces.IDataLoader, tableName *string, dbName *string) (ok bool) {

	if !(helpers.GetModelLengt(InsertData) > 0) {
		err := errors.New("filter data is empty")
		helpers.CheckErr(err)
		return false
	}

	clientOptions := options.Client().ApplyURI(set.MongoUri)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, clientOptions)
	if !helpers.CheckErr(err) {
		return
	}

	collection := client.Database(*dbName).Collection(*tableName)

	insertResult, err := collection.InsertOne(context.TODO(), InsertData)
	if !helpers.CheckErr(err) {
		return
	}

	err = client.Disconnect(context.TODO())

	helpers.CheckErr(err)

	if oid, ok2 := insertResult.InsertedID.(primitive.ObjectID); ok2 {
		if oid.IsZero() {
			return
		}
		return true
	}

	return
}

func SorMongo(filterData interfaces.IDataLoader, resultData interfaces.IDataLoader, tableName *string, dbName *string) (ok bool, empty bool) {

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

func RefMongo(filterData interfaces.IDataLoader, newData interfaces.IDataLoader, tableName *string, dbName *string) (ok bool) {

	if !(helpers.GetModelLengt(filterData) > 0) {
		err := errors.New("filter data is empty")
		helpers.CheckErr(err)
		return false
	}

	if !(helpers.GetModelLengt(newData) > 0) {
		err := errors.New("new data is empty")
		helpers.CheckErr(err)
		return false
	}

	updateData := bson.M{
		"$set": newData,
	}

	clientOptions := options.Client().ApplyURI(set.MongoUri)

	// Connect to MongoDB
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, clientOptions)
	if !helpers.CheckErr(err) {
		return false
	}

	collection := client.Database(*dbName).Collection(*tableName)

	update, err := collection.UpdateOne(context.TODO(), filterData, updateData)
	if !helpers.CheckErr(err) {
		return false
	}

	err = client.Disconnect(context.TODO())

	helpers.CheckErr(err)

	d := update.MatchedCount

	if d == 1 {

		return true
	} else {
		return false
	}

}

func DelMongo(filterData interfaces.IDataLoader, tableName *string, dbName *string) (ok bool) {

	if !(helpers.GetModelLengt(filterData) > 0) {
		err := errors.New("filter data is empty")
		helpers.CheckErr(err)
		return false
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

	data, err := collection.DeleteOne(context.TODO(), filterData)

	helpers.CheckErr(err)

	err = client.Disconnect(context.TODO())

	helpers.CheckErr(err)

	count := data.DeletedCount

	if !(count > 0) {
		return true
	}

	return
}

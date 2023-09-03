package helpers

import (
	"all_lb/set"
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

type LogError struct {
	ErrorString string
	Time        string
}

func (l *LogError) New(e error) {
	l.Time = time.Now().String()
	l.ErrorString = e.Error()
}

func CheckErr(e error) bool {
	if e == nil {
		return true
	}
	var logError LogError
	logError.New(e)

	clientOptions := options.Client().ApplyURI(set.MongoUri)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		errLog(e, "")
		errLog(err, "")
		return false
	}

	collection := client.Database("errors").Collection("lb-sign")

	insertResult, err := collection.InsertOne(context.TODO(), &logError)
	if err != nil {
		errLog(e, "")
		errLog(err, "kayıt hatası")
		return false
	}

	err = client.Disconnect(context.TODO())
	if err != nil {
		errLog(err, "")
	}

	if oid, ok2 := insertResult.InsertedID.(primitive.ObjectID); ok2 {
		if oid.IsZero() {
			errLog(e, "")
			errLog(err, "")
			return false
		}
	}
	return false
}

func errLog(err error, str string) {
	if str == "" {
		log.Println(err.Error())
		return
	}
	log.Println(err.Error() + " /*/ " + str)
}

package set

import (
	"os"
)

var MongoUri string

var AccessSecret string
var RefreshSecret string

var TokenDBName string
var TokenDBPort string

func InitAllStrings() {

	MongoUri = os.Getenv("MONGO_URI")

	AccessSecret = os.Getenv("ACCESS_SECRET")
	RefreshSecret = os.Getenv("REFRESH_SECRET")

	TokenDBName = os.Getenv("TOKEN_DB_NAME")
	TokenDBPort = os.Getenv("TOKEN_DB_PORT")

	return
}

package set

import (
	"fmt"
	"os"
	"strconv"
)

var SignTableName string
var SignDbName string

var SmsSqlTableName string
var SmsSqlLogTableName string

var MongoUri string

var psqlHost string
var psqlPort int
var psqlUser string
var psqlPass string
var psqlDbName string

var PsqlInfo string

var TokenUri string
var NewUserUri string

func InitAllStrings() (err error) {
	SignTableName = os.Getenv("SIGN_TABLE_NAME")
	SignDbName = os.Getenv("SIGN_DB_NAME")

	SmsSqlTableName = os.Getenv("SMS_SQL_TABLE_NAME")
	SmsSqlLogTableName = os.Getenv("SMS_SQL_LOG_TABLE_NAME")

	MongoUri = os.Getenv("MONGO_URI")

	psqlHost = os.Getenv("PSQL_HOST")
	portString := os.Getenv("PSQL_PORT")

	psqlPort, err = strconv.Atoi(portString)

	psqlUser = os.Getenv("PSQL_USER")
	psqlPass = os.Getenv("PSQL_PASSWORD")
	psqlDbName = os.Getenv("PSQL_DBNAME")

	TokenUri = os.Getenv("NEW_TOKEN_URI")

	NewUserUri = os.Getenv("NEW_USER_URI")

	PsqlInfo = fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", psqlHost, psqlPort, psqlUser, psqlPass, psqlDbName)

	return
}

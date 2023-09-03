package set

import (
	"os"
)

var MongoUri string

var IncomingDBName string
var OutgoingDBName string
var UserHistoryDBName string

var RegularIncomingTableName string
var RegularHistoryIncomingTableName string
var IncomingTableName string
var UserHistoryTableName string
var IncomingDailyTableName string

var OutgoingDailyTableName string
var RegularOutgoingTableName string
var RegularHistoryOutgoingTableName string
var OutgoingTableName string

func InitAllStrings() (err error) {
	MongoUri = os.Getenv("MONGO_URI")
	IncomingDBName = os.Getenv("INCOMING_DB_NAME")
	RegularIncomingTableName = os.Getenv("REGULAR_INCOMING_TABLE_NAME")
	RegularHistoryIncomingTableName = os.Getenv("REGULAR_HISTORY_INCOMING_TABLE_NAME")
	IncomingTableName = os.Getenv("INCOMING_TABLE_NAME")
	UserHistoryTableName = os.Getenv("USER_HISTORY_TABLE_NAME")
	IncomingDailyTableName = os.Getenv("INCOMING_DAILY_TABLE_NAME")
	UserHistoryDBName = os.Getenv("USER_HISTORY_DB_NAME")

	OutgoingDBName = os.Getenv("OUTGOING_DB_NAME")
	OutgoingDailyTableName = os.Getenv("OUTGOING_DAILY_TABLE_NAME")
	RegularOutgoingTableName = os.Getenv("REGULAR_OUTGOING_TABLE_NAME")
	RegularHistoryOutgoingTableName = os.Getenv("REGULAR_HISTORY_OUTGOING_TABLE_NAME")
	OutgoingTableName = os.Getenv("OUTGOING_TABLE_NAME")
	return
}

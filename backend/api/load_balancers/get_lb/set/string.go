package set

import (
	"os"
)

var MongoUri string

var TokenUri string

var GetBaseUri string
var HistoryPath string
var ReqInHisPath string
var ReqOutHisPath string
var IncomingPath string
var RegularIncomingPath string
var OutgoingPath string
var RegularOutgoingPath string

func InitAllStrings() {
	MongoUri = os.Getenv("MONGO_URI")

	TokenUri = os.Getenv("TOKEN_URI")

	GetBaseUri = os.Getenv("GET_BASE_URI")
	HistoryPath = os.Getenv("HISTORY_PATH")
	ReqInHisPath = os.Getenv("INCOMING_REG_HISTORY_PATH")
	ReqOutHisPath = os.Getenv("OUTGOING_REG_HISTORY_PATH")
	IncomingPath = os.Getenv("INCOMING_PATH")
	RegularIncomingPath = os.Getenv("REGULAR_INCOMING_PATH")
	OutgoingPath = os.Getenv("OUTGOING_PATH")
	RegularOutgoingPath = os.Getenv("REGULAR_OUTGOING_PATH")
}

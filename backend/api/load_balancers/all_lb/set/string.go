package set

import (
	"os"
)

var MongoUri, SignUri, SignInPath, SignUpPath, NewSmsPath, AskSmsPath, NewSmsRefPath, AskSmsRefPath, RefPassPath string

var TokenUri string

var GetBaseUri string
var HistoryPath string
var ReqInHisPath string
var ReqOutHisPath string
var IncomingPath string
var RegularIncomingPath string
var OutgoingPath string
var RegularOutgoingPath string

var InComBaseUri string
var InComNewPath string
var InComRegNewPath string
var InComRegRefPath string

var OutGoBaseUri string
var OutGoNewPath string
var OutGoRegNewPath string
var OutGoRegRefPath string

func InitAllStrings() {
	MongoUri = os.Getenv("MONGO_URI")
	SignUri = os.Getenv("SIGN_URI")
	SignInPath = os.Getenv("SIGN_IN_PATH")
	SignUpPath = os.Getenv("SIGN_UP_PATH")
	NewSmsPath = os.Getenv("NEW_SMS_PATH")
	AskSmsPath = os.Getenv("ASK_SMS_PATH")
	NewSmsRefPath = os.Getenv("NEW_SMS_REF_PATH")
	AskSmsRefPath = os.Getenv("ASK_SMS_REF_PATH")
	RefPassPath = os.Getenv("REF_PASS_PATH")

	TokenUri = os.Getenv("TOKEN_URI")

	GetBaseUri = os.Getenv("GET_BASE_URI")
	HistoryPath = os.Getenv("HISTORY_PATH")
	ReqInHisPath = os.Getenv("INCOMING_REG_HISTORY_PATH")
	ReqOutHisPath = os.Getenv("OUTGOING_REG_HISTORY_PATH")
	IncomingPath = os.Getenv("INCOMING_PATH")
	RegularIncomingPath = os.Getenv("REGULAR_INCOMING_PATH")
	OutgoingPath = os.Getenv("OUTGOING_PATH")
	RegularOutgoingPath = os.Getenv("REGULAR_OUTGOING_PATH")

	InComBaseUri = os.Getenv("INCOMING_BASE_URI")
	InComNewPath = os.Getenv("INCOMING_NEW_PATH")
	InComRegNewPath = os.Getenv("INCOMING_NEW_REG_PATH")
	InComRegRefPath = os.Getenv("INCOMING_REF_REG_PATH")

	OutGoBaseUri = os.Getenv("OUTGOING_BASE_URI")
	OutGoNewPath = os.Getenv("OUTGOING_NEW_PATH")
	OutGoRegNewPath = os.Getenv("OUTGOING_NEW_REG_PATH")
	OutGoRegRefPath = os.Getenv("OUTGOING_REF_REG_PATH")
}

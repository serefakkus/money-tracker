package set

import (
	"os"
)

var MongoUri string

var TokenUri string

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

	TokenUri = os.Getenv("TOKEN_URI")

	InComBaseUri = os.Getenv("INCOMING_BASE_URI")
	InComNewPath = os.Getenv("INCOMING_NEW_PATH")
	InComRegNewPath = os.Getenv("INCOMING_NEW_REG_PATH")
	InComRegRefPath = os.Getenv("INCOMING_REF_REG_PATH")

	OutGoBaseUri = os.Getenv("OUTGOING_BASE_URI")
	OutGoNewPath = os.Getenv("OUTGOING_NEW_PATH")
	OutGoRegNewPath = os.Getenv("OUTGOING_NEW_REG_PATH")
	OutGoRegRefPath = os.Getenv("OUTGOING_REF_REG_PATH")
}

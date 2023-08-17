package set

import (
	"os"
)

var MongoUri, SignUri, SignInPath, SignUpPath, NewSmsPath, AskSmsPath, NewSmsRefPath, AskSmsRefPath, RefPassPath string

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
}

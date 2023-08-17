package models

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
	"sign/dataloaders/mongo"
	"sign/helpers"
	"sign/set"
	"time"
)

type ReqSignUp struct {
	Phone string
	Code  string
	Pass  string
	Ip    ip
}

func (r *ReqSignUp) insertMongo(details *TokenDetails) (code int) {
	if !helpers.CheckPass(r.Pass) {
		return http.StatusRequestedRangeNotSatisfiable
	}

	var sms ReqPhoneCode
	sms.Phone = r.Phone
	sms.Code = r.Code

	var ok bool

	code, ok = sms.askCode()
	if !ok {
		return
	}

	var user MongoSignInfo
	var empty bool

	ok, empty = user.isExistUser(r.Phone)

	if !empty {
		return http.StatusAlreadyReported
	}

	user.Ip = r.Ip
	var now = time.Now()
	helpers.TimeToString(&user.Time, &now)

	user.Pass = r.Pass
	user.UserId = primitive.NewObjectID().Hex()
	user.Phone = r.Phone
	user.newPass()

	if !mongo.NewMongo(&user, &set.SignTableName, &set.SignDbName) {
		return http.StatusInternalServerError
	}

	var newUser ReqNewUser
	if !newUser.SendToDb(&user.UserId) {
		var filter filterUserId
		filter.UserId = user.UserId
		mongo.DelMongo(&filter, &set.SignTableName, &set.SignDbName)
		return http.StatusInternalServerError
	}

	var reqToken ReqNewToken

	reqToken.getNewToken(&user.UserId, details)

	return http.StatusOK
}

func (r *ReqSignUp) HandlerSignUp(c *gin.Context) {
	err := c.ShouldBindJSON(r)
	if !helpers.CheckErr(err) {
		c.Writer.WriteHeader(http.StatusBadRequest)
		return
	}

	var resp TokenDetails

	c.IndentedJSON(r.insertMongo(&resp), resp)
}

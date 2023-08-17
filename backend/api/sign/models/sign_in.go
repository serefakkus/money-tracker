package models

import (
	"crypto/sha1"
	"encoding/hex"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
	"sign/dataloaders/mongo"
	"sign/helpers"
	"sign/set"
)

type ReqSignIn struct {
	Phone string
	Pass  string
	Ip    ip
}

type filterSignPhone struct {
	Phone string
}

func (s *ReqSignIn) HandlerSignIn(c *gin.Context) {
	err := c.ShouldBindJSON(s)
	if !helpers.CheckErr(err) {
		c.Writer.WriteHeader(http.StatusBadRequest)
		return
	}

	var resp TokenDetails

	c.IndentedJSON(s.signIn(&resp), resp)
}

//-----------------------helpers

func (s *ReqSignIn) passConverter(passHelper string) {
	h := sha1.New()
	h.Write([]byte(s.Pass))
	s.Pass = hex.EncodeToString(h.Sum(nil))
	s.Pass += passHelper
	h = sha1.New()
	h.Write([]byte(s.Pass))
	s.Pass = hex.EncodeToString(h.Sum(nil))
}

func (s *ReqSignIn) fromJson(data []byte) bool {
	err := json.Unmarshal(data, s)
	return helpers.CheckErr(err)
}

func (s *ReqSignIn) signIn(tokenDetails *TokenDetails) (code int) {
	if !helpers.CheckPhoneNo(s.Phone) {
		return http.StatusRequestedRangeNotSatisfiable
	}

	var phone filterSignPhone
	var result MongoSignInfo
	phone.Phone = s.Phone

	ok, empty := mongo.SorMongo(&phone, &result, &set.SignTableName, &set.SignDbName)
	if !ok {
		return http.StatusInternalServerError
	}

	if empty {
		return http.StatusNotFound
	}

	s.passConverter(result.PassHelper)

	if result.Pass != s.Pass || result.Phone != s.Phone || result.UserId == "" {
		return http.StatusUnauthorized
	}

	var mongoId filterMongoId

	mongoId.Id = result.Id

	var newIps mongoSignInIps
	newIps.newIp(&result.SignInIps, &s.Ip)

	mongo.RefMongo(&mongoId, &newIps, &set.SignTableName, &set.SignDbName)

	var tokReq ReqNewToken
	ok, code = tokReq.getNewToken(&result.UserId, tokenDetails)

	if !ok {
		return
	}

	return http.StatusOK
}

package models

import (
	"crypto/sha1"
	"encoding/hex"
	"errors"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"sign/dataloaders/mongo"
	"sign/helpers"
	"sign/set"
	"time"
)

type MongoSignInfo struct {
	Id         primitive.ObjectID `bson:"_id"`
	UserId     string
	Phone      string
	Pass       string
	PassHelper string
	Time       string
	Ip         ip
	SignInIps  []ip
}

func (s *MongoSignInfo) newPass() {
	s.Id = primitive.NewObjectID()
	s.PassHelper = helpers.RandStrRunes(30)
	h := sha1.New()
	h.Write([]byte(s.Pass))
	s.Pass = hex.EncodeToString(h.Sum(nil))
	s.Pass += s.PassHelper
	h = sha1.New()
	h.Write([]byte(s.Pass))
	s.Pass = hex.EncodeToString(h.Sum(nil))
}

func (s *MongoSignInfo) isExistUser(phone string) (ok bool, empty bool) {
	s.Phone = phone
	var filter filterSignPhone

	filter.Phone = phone

	ok, empty = mongo.SorMongo(&filter, s, &set.SignTableName, &set.SignDbName)

	if !ok {
		err := errors.New("sormongo error")
		helpers.CheckErr(err)
		return
	}

	return
}

type ip struct {
	Ip   string
	Time string
}

type mongoSignInIps struct {
	SignInIps []ip
}

type filterMongoId struct {
	Id primitive.ObjectID `bson:"_id"`
}

type filterUserId struct {
	UserId string
}

func (m *mongoSignInIps) newIp(ips *[]ip, ip *ip) {
	var now = time.Now()
	helpers.TimeToString(&ip.Time, &now)
	m.SignInIps = append(*ips, *ip)
}

type filterPhoneMongo struct {
	Phone string
}

package models

import (
	"crypto/sha1"
	"encoding/hex"
	"sign/helpers"
)

type MongoSignRefPass struct {
	Pass       string
	PassHelper string
	SignInIps  []ip
}

func (s *MongoSignRefPass) refPass(pass string, ip ip) {
	s.Pass = pass
	s.SignInIps = append(s.SignInIps, ip)

	s.PassHelper = helpers.RandStrRunes(30)
	h := sha1.New()
	h.Write([]byte(s.Pass))
	s.Pass = hex.EncodeToString(h.Sum(nil))
	s.Pass += s.PassHelper
	h = sha1.New()
	h.Write([]byte(s.Pass))
	s.Pass = hex.EncodeToString(h.Sum(nil))

}

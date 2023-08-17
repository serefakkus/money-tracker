package models

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/go-redis/redis/v7"
	"github.com/twinj/uuid"
	"log"
	"time"
	"token/set"
)

var client *redis.Client

type Token struct {
	UserId       string
	TokenDetails TokenDetails
	Auth         string
}

//-----------------------details

type TokenDetails struct {
	AccessToken  string
	RefreshToken string
	AccessUuid   string
	RefreshUuid  string
	AtExpires    int64
	RtExpires    int64
}

func (token *Token) CreateToken() bool {
	token.TokenDetails.AtExpires = time.Now().Add(time.Hour * 24 * 7).Unix()
	token.TokenDetails.AccessUuid = uuid.NewV4().String()

	token.TokenDetails.RtExpires = time.Now().Add(time.Hour * 24 * 365).Unix()
	token.TokenDetails.RefreshUuid = uuid.NewV4().String()

	var err error

	//Creating Access Token

	atClaims := jwt.MapClaims{}
	atClaims["authorized"] = true
	atClaims["access_uuid"] = token.TokenDetails.AccessUuid
	atClaims["user_id"] = token.UserId
	atClaims["exp"] = token.TokenDetails.AtExpires
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	token.TokenDetails.AccessToken, err = at.SignedString([]byte(set.AccessSecret))
	if !CheckErr(err) {
		return false
	}

	//Creating Refresh Token

	rtClaims := jwt.MapClaims{}
	rtClaims["refresh_uuid"] = token.TokenDetails.RefreshUuid
	rtClaims["user_id"] = token.UserId
	rtClaims["exp"] = token.TokenDetails.RtExpires
	rt := jwt.NewWithClaims(jwt.SigningMethodHS256, rtClaims)
	token.TokenDetails.RefreshToken, err = rt.SignedString([]byte(set.RefreshSecret))
	if !CheckErr(err) {
		return false
	}
	return true
}

type AccessDetails struct {
	AccessUuid string
	UserId     string
}

type RefreshDetails struct {
	RefreshUuid string
	UserId      string
}

//--------------------------------redis db

func InitRedis() (ok bool) {
	//Initializing redis

	dsn := set.TokenDBName + ":" + set.TokenDBPort

	client = redis.NewClient(&redis.Options{
		Addr: dsn, //redis port
		DB:   0,
	})
	_, err := client.Ping().Result()

	if !CheckErr(err) {
		panic("redis server conn err")
		return false
	}

	return true
}

func CreateAuth(userid string, td *TokenDetails) error {
	at := time.Unix(td.AtExpires, 0) //converting Unix to UTC(to Time object)
	rt := time.Unix(td.RtExpires, 0)
	now := time.Now()

	errAccess := client.Set(td.AccessUuid, userid, at.Sub(now)).Err()
	if errAccess != nil {
		return errAccess
	}
	errRefresh := client.Set(td.RefreshUuid, userid, rt.Sub(now)).Err()
	if errRefresh != nil {
		return errRefresh
	}
	return nil
}

func FetchAuth(authD *AccessDetails) (string, error) {
	userid, err := client.Get(authD.AccessUuid).Result()
	if err != nil {
		log.Println(err.Error())
		return "", err
	}

	return userid, nil
}

func RefFetchAuth(authD *RefreshDetails) (string, error) {
	userid, err := client.Get(authD.RefreshUuid).Result()
	if err != nil {
		return "", err
	}

	return userid, nil
}

func DeleteAuth(givenUuid string) (int64, error) {
	deleted, err := client.Del(givenUuid).Result()
	if err != nil {
		return 0, err
	}
	return deleted, nil
}

//--------------------------------------------------------heslpers

func VerifyToken(t *Token) (*jwt.Token, error) {
	token, err := jwt.Parse(t.Auth, func(token *jwt.Token) (interface{}, error) {
		//Make sure that the token method conform to "SigningMethodHMAC"
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(set.AccessSecret), nil
	})
	if err != nil {
		return nil, err
	}
	return token, nil
}

func TokenValid(t *Token) error {
	token, err := VerifyToken(t)
	if err != nil {
		return err
	}
	if _, ok := token.Claims.(jwt.Claims); !ok && !token.Valid {
		return err
	}
	return nil
}

func ExtractTokenMetadata(t *Token) (*AccessDetails, error) {
	token, err := VerifyToken(t)
	if err != nil {
		return nil, err
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		accessUuid, ok := claims["access_uuid"].(string)
		if !ok {
			return nil, err
		}
		userId, ok := claims["user_id"].(string)

		if !ok {
			return nil, err
		}

		if err != nil {
			return nil, err
		}
		return &AccessDetails{
			AccessUuid: accessUuid,
			UserId:     userId,
		}, nil
	}
	return nil, err
}

func VerifyRefToken(t *Token) (*jwt.Token, error) {
	token, err := jwt.Parse(t.Auth, func(token *jwt.Token) (interface{}, error) {
		//Make sure that the token method conform to "SigningMethodHMAC"
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(set.RefreshSecret), nil
	})
	if err != nil {
		return nil, err
	}
	return token, nil
}

func RefTokenValid(t *Token) error {
	token, err := VerifyRefToken(t)
	if err != nil {
		return err
	}
	if _, ok := token.Claims.(jwt.Claims); !ok && !token.Valid {
		return err
	}
	return nil
}

func ExtractRefTokenMetadata(t *Token) (*RefreshDetails, error) {
	token, err := VerifyRefToken(t)
	if err != nil {
		return nil, err
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		refreshUuid, ok := claims["refresh_uuid"].(string)
		if !ok {
			return nil, err
		}
		userId, ok := claims["user_id"].(string)
		if !ok {
			return nil, err
		}
		return &RefreshDetails{
			RefreshUuid: refreshUuid,
			UserId:      userId,
		}, nil
	}
	return nil, err
}

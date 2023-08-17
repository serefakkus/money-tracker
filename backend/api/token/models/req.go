package models

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
	"token/helpers"
)

type NewTokenReq struct {
	UserId string
}

func (t *NewTokenReq) fromJson(data []byte) {
	var err error
	err = json.Unmarshal(data, t)
	CheckErr(err)
	return
}

//-------------------------------------------------------------------

type TokenReq struct {
	Auth string
}

func (t *TokenReq) fromJson(data []byte) {
	var err error
	err = json.Unmarshal(data, t)
	CheckErr(err)
	return
}

//--------------------is ok

func (t *TokenReq) isOkToken(resp *TokenResp) (code int) {
	var tok Token
	tok.Auth = t.Auth

	if !helpers.CheckAuth(&tok.Auth) {
		return http.StatusRequestedRangeNotSatisfiable
	}

	err := TokenValid(&tok)

	if !CheckErr(err) {
		return http.StatusUnauthorized
	}

	tokenAuth, err := ExtractTokenMetadata(&tok)

	if !CheckErr(err) {
		return http.StatusUnauthorized
	}

	resp.UserId, err = FetchAuth(tokenAuth)

	if !CheckErr(err) {
		return http.StatusRequestTimeout
	}

	return http.StatusOK
}

func (t *TokenReq) HandlerIsOk(c *gin.Context) {
	t.Auth = c.GetHeader("Authorization")

	var resp TokenResp

	code := t.isOkToken(&resp)

	c.IndentedJSON(code, resp)
}

//-----------------------------ref

func (t *TokenReq) refToken(resp *TokenDetails) (code int) {
	var tok Token
	tok.Auth = t.Auth

	if !helpers.CheckAuth(&tok.Auth) {
		return http.StatusRequestedRangeNotSatisfiable
	}

	err := RefTokenValid(&tok)
	if !CheckErr(err) {
		return http.StatusUnauthorized
	}

	tokenAuth, err := ExtractRefTokenMetadata(&tok)
	if !CheckErr(err) {
		return http.StatusUnauthorized
	}

	tok.UserId, err = RefFetchAuth(tokenAuth)
	if !CheckErr(err) {
		return http.StatusRequestTimeout
	}

	if tok.UserId == "" {
		return http.StatusUnauthorized
	}

	deleted, delErr := DeleteAuth(tokenAuth.RefreshUuid)
	if delErr != nil || deleted == 0 { //if any goes wrong
		CheckErr(delErr)
		return http.StatusRequestTimeout
	}

	if !tok.CreateToken() {
		return http.StatusInternalServerError
	}

	err = CreateAuth(tok.UserId, &tok.TokenDetails)

	if !CheckErr(err) {
		return http.StatusInternalServerError
	}

	*resp = tok.TokenDetails

	return http.StatusOK
}

func (t *TokenReq) HandlerRef(c *gin.Context) {
	t.Auth = c.GetHeader("Authorization")

	var resp TokenDetails

	code := t.refToken(&resp)

	c.IndentedJSON(code, resp)
}

//-----------------------------del

func (t *TokenReq) delToken() (code int) {
	var tok Token
	tok.Auth = t.Auth

	if !helpers.CheckAuth(&tok.Auth) {
		return http.StatusRequestedRangeNotSatisfiable
	}

	au, err := ExtractTokenMetadata(&tok)
	if !CheckErr(err) {
		return http.StatusUnauthorized
	}

	deleted, delErr := DeleteAuth(au.AccessUuid)
	if delErr != nil || deleted == 0 { //if any goes wrong
		CheckErr(err)
		return http.StatusRequestTimeout
	}

	return http.StatusOK
}

func (t *TokenReq) HandlerDel(c *gin.Context) {
	t.Auth = c.GetHeader("Authorization")

	code := t.delToken()

	c.Writer.WriteHeader(code)
}

//-----------------------------newToken

func (t *NewTokenReq) newToken(resp *TokenDetails, c *gin.Context) (code int) {
	err := c.ShouldBindJSON(t)
	if !CheckErr(err) {
		return http.StatusRequestedRangeNotSatisfiable
	}

	var tok Token
	tok.UserId = t.UserId

	if t.UserId == "" {
		return http.StatusUnauthorized
	}

	if !tok.CreateToken() {
		return http.StatusInternalServerError
	}

	err = CreateAuth(tok.UserId, &tok.TokenDetails)
	if !CheckErr(err) {
		return http.StatusInternalServerError
	}

	*resp = tok.TokenDetails

	return http.StatusOK
}

func (t *NewTokenReq) HandlerNew(c *gin.Context) {
	var resp TokenDetails

	c.IndentedJSON(t.newToken(&resp, c), resp)
}

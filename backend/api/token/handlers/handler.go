package handlers

import (
	"github.com/gin-gonic/gin"
	"token/models"
)

func Ping(c *gin.Context) {
	c.Writer.WriteHeader(200)
}

func IsOkToken(c *gin.Context) {
	var req models.TokenReq
	req.HandlerIsOk(c)
}

func RefToken(c *gin.Context) {
	var req models.TokenReq
	req.HandlerRef(c)
}

func DelToken(c *gin.Context) {
	var req models.TokenReq
	req.HandlerDel(c)
}

func NewToken(c *gin.Context) {
	var req models.NewTokenReq
	req.HandlerNew(c)
}

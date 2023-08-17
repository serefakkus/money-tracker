package handlers

import (
	"github.com/gin-gonic/gin"
	"sign/models"
)

func Ping(c *gin.Context) {
	c.Writer.WriteHeader(200)
}

func NewSms(c *gin.Context) {
	var req models.ReqNewSms
	req.HandlerNewSmsSignUp(c)
}

func NewSmsRefPass(c *gin.Context) {
	var req models.ReqNewSms
	req.HandlerNewSmsRefPass(c)
}

func SignIn(c *gin.Context) {
	var req models.ReqSignIn
	req.HandlerSignIn(c)
}

func SignUp(c *gin.Context) {
	var req models.ReqSignUp
	req.HandlerSignUp(c)
}

func RefPass(c *gin.Context) {
	var req models.ReqSignUp
	req.HandlerRefPass(c)
}

func AskSms(c *gin.Context) {
	var req models.ReqPhoneCode
	req.HandlerAskCodeSignUp(c)
}

func AskSmsRef(c *gin.Context) {
	var req models.ReqPhoneCode
	req.HandlerAskCodeRefPass(c)
}

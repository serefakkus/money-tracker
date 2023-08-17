package main

import (
	"github.com/gin-gonic/gin"
	"sign/dataloaders/sql"
	"sign/handlers"
	"sign/helpers"
)

func main() {

	sql.InitEveryThings()
	gin.SetMode(gin.ReleaseMode)
	r := gin.New()

	err := r.SetTrustedProxies(nil)

	if !helpers.CheckErr(err) {
		return
	}

	r.GET("/ping", handlers.Ping)

	r.POST("/sign-in", handlers.SignIn)

	r.POST("/new-sms", handlers.NewSms)

	r.POST("/ask-sms", handlers.AskSms)

	r.POST("/sign-up", handlers.SignUp)

	r.POST("/new-sms-ref", handlers.NewSmsRefPass)

	r.POST("/ask-sms-ref", handlers.AskSmsRef)

	r.POST("/ref-pass", handlers.RefPass)

	err = r.Run("0.0.0.0:80")

	helpers.CheckErr(err)
}

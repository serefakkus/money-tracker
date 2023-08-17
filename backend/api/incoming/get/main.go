package main

import (
	"get/handlers"
	"get/helpers"
	"get/set"
	"github.com/gin-gonic/gin"
)

func main() {

	err := set.InitAllStrings()

	if !helpers.CheckErr(err) {
		panic(err)
	}

	gin.SetMode(gin.ReleaseMode)
	r := gin.New()
	err = r.SetTrustedProxies(nil)

	if !helpers.CheckErr(err) {
		return
	}

	r.GET("/ping", handlers.Ping)

	r.POST("/his", handlers.AskHis)

	r.POST("/in", handlers.AskIn)

	r.POST("/out", handlers.AskOut)

	r.POST("/in-reg-his", handlers.AskInRegHis)

	r.POST("/out-reg-his", handlers.AskOutRegHis)

	r.POST("/in-reg", handlers.AskInReg)

	r.POST("/out-reg", handlers.AskOutReg)

	err = r.Run("0.0.0.0:80")

	helpers.CheckErr(err)
}

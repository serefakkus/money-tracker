package main

import (
	"github.com/gin-gonic/gin"
	"token/handlers"
	"token/models"
	"token/set"
)

func main() {

	set.InitAllStrings()
	if !models.InitRedis() {
		panic("")
	}

	gin.SetMode(gin.ReleaseMode)

	r := gin.New()

	err := r.SetTrustedProxies(nil)
	if !models.CheckErr(err) {
		return
	}

	r.GET("/ping", handlers.Ping)
	r.GET("/is_ok", handlers.IsOkToken)
	r.GET("/ref", handlers.RefToken)
	r.GET("/del", handlers.DelToken)
	r.POST("/new", handlers.NewToken)

	err = r.Run("0.0.0.0:80")

	models.CheckErr(err)
}

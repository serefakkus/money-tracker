package handlers

import (
	"get/models/request"
	"github.com/gin-gonic/gin"
)

func Ping(c *gin.Context) {
	c.Writer.WriteHeader(200)
}

func AskHis(c *gin.Context) {
	var req request.ReqAsk
	req.HandlerAsk(c)
}

func AskOut(c *gin.Context) {
	var req request.OutgoReq
	req.HandlerAsk(c)
}

func AskIn(c *gin.Context) {
	var req request.IncomingReq
	req.HandlerAsk(c)
}

func AskOutRegHis(c *gin.Context) {
	var req request.RegularHisReq
	req.HandlerOutHis(c)
}

func AskOutReg(c *gin.Context) {
	var req request.RegularOutReq
	req.HandlerAsk(c)
}

func AskInRegHis(c *gin.Context) {
	var req request.RegularHisReq
	req.HandlerInHis(c)
}

func AskInReg(c *gin.Context) {
	var req request.RegularInReq
	req.HandlerAsk(c)
}

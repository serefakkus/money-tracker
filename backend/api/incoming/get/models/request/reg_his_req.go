package request

import (
	"get/helpers"
	"get/models/mongo_models"
	"get/models/response"
	"github.com/gin-gonic/gin"
	"net/http"
)

type RegularHisReq struct {
	UserId    string
	RegularId string
}

func (r *RegularHisReq) HandlerOutHis(c *gin.Context) {
	err := c.ShouldBindJSON(r)
	if !helpers.CheckErr(err) {
		c.Writer.WriteHeader(http.StatusBadRequest)
		return
	}

	var resp response.RegHisOutResp

	code := r.getOutHis(&resp)

	c.IndentedJSON(code, resp)
}

func (r *RegularHisReq) getOutHis(regHis *response.RegHisOutResp) (code int) {
	var regMon mongo_models.RegularOutHistoryMongo
	var ok, empty = regMon.AskMongo(&r.UserId, &r.RegularId)
	if !ok {
		return http.StatusInternalServerError
	}

	if empty {
		return http.StatusNotFound
	}

	regMon.ToResp(regHis)

	return http.StatusOK
}

func (r *RegularHisReq) HandlerInHis(c *gin.Context) {
	err := c.ShouldBindJSON(r)
	if !helpers.CheckErr(err) {
		c.Writer.WriteHeader(http.StatusBadRequest)
		return
	}

	var resp response.RegHisInResp

	code := r.getInHis(&resp)

	c.IndentedJSON(code, resp)
}

func (r *RegularHisReq) getInHis(regHis *response.RegHisInResp) (code int) {
	var regMon mongo_models.RegularInHistoryMongo
	var ok, empty = regMon.AskMongo(&r.UserId, &r.RegularId)
	if !ok {
		return http.StatusInternalServerError
	}

	if empty {
		return http.StatusNotFound
	}

	regMon.ToResp(regHis)

	return http.StatusOK
}

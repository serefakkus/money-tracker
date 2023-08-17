package request

import (
	"get/helpers"
	"get/models/mongo_models"
	"get/models/response"
	"github.com/gin-gonic/gin"
	"net/http"
)

type ReqAsk struct {
	UserId string
}

func (r *ReqAsk) HandlerAsk(c *gin.Context) {
	err := c.ShouldBindJSON(r)
	if !helpers.CheckErr(err) {
		c.Writer.WriteHeader(http.StatusBadRequest)
		return
	}

	var resp response.RespUserHistory

	code := r.getUserHistory(&resp)

	c.IndentedJSON(code, resp)
}

func (r *ReqAsk) getUserHistory(his *response.RespUserHistory) (code int) {
	if r.UserId == "" {
		return http.StatusRequestedRangeNotSatisfiable
	}

	var historyMongo mongo_models.UserHistoryMongo

	ok, empty := historyMongo.AskMongo(r.UserId)

	if !ok {
		return http.StatusInternalServerError
	}

	if empty {
		return http.StatusNotFound
	}

	if !historyMongo.ToResp(his) {
		return http.StatusRequestTimeout
	}

	return http.StatusOK
}

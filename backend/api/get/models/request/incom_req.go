package request

import (
	"get/dataloaders/mongo"
	"get/helpers"
	"get/models/mongo_models"
	"get/models/response"
	"get/set"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
)

type IncomingReq struct {
	UserId     string
	IncomingId string
}

func (i *IncomingReq) HandlerAsk(c *gin.Context) {
	err := c.ShouldBindJSON(i)
	if !helpers.CheckErr(err) {
		c.Writer.WriteHeader(http.StatusBadRequest)
		return
	}

	var resp response.InComInfo

	code := i.getIn(&resp)

	c.IndentedJSON(code, resp)
}

func (i *IncomingReq) getIn(inCom *response.InComInfo) (code int) {
	var filter mongo_models.FilterId
	var err error
	filter.Id, err = primitive.ObjectIDFromHex(i.IncomingId)
	if !helpers.CheckErr(err) {
		return http.StatusBadRequest
	}

	ok, empty := mongo.AskMongo(&filter, inCom, &set.IncomingTableName, &set.IncomingDBName)

	if !ok {
		return http.StatusInternalServerError
	}

	if empty {
		return http.StatusNotFound
	}

	return http.StatusOK
}

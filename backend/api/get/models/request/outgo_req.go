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

type OutgoReq struct {
	UserId     string
	OutgoingId string
}

func (o *OutgoReq) HandlerAsk(c *gin.Context) {
	err := c.ShouldBindJSON(o)
	if !helpers.CheckErr(err) {
		c.Writer.WriteHeader(http.StatusBadRequest)
		return
	}

	var resp response.OutGoInfo

	code := o.getOut(&resp)

	c.IndentedJSON(code, resp)
}

func (o *OutgoReq) getOut(outGo *response.OutGoInfo) (code int) {
	var filter mongo_models.FilterId
	var err error
	filter.Id, err = primitive.ObjectIDFromHex(o.OutgoingId)
	if !helpers.CheckErr(err) {
		return http.StatusBadRequest
	}

	ok, empty := mongo.AskMongo(&filter, outGo, &set.OutgoingTableName, &set.OutgoingDBName)

	if !ok {
		return http.StatusInternalServerError
	}

	if empty {
		return http.StatusNotFound
	}

	return http.StatusOK
}

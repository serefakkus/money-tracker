package request

import (
	"get/helpers"
	"get/models/mongo_models"
	"get/models/response"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
	"net/http"
)

type RegularOutReq struct {
	UserId    string
	RegularId string
	Offset    int
}

func (r *RegularOutReq) HandlerAsk(c *gin.Context) {
	err := c.ShouldBindJSON(r)
	if !helpers.CheckErr(err) {
		c.Writer.WriteHeader(http.StatusBadRequest)
		return
	}

	var respList []response.RegOutResp
	code := r.ask(&respList)

	c.IndentedJSON(code, respList)
}

func (r *RegularOutReq) ask(respList *[]response.RegOutResp) (code int) {
	var userHis mongo_models.UserHistoryMongo
	var regMongo []mongo_models.RegularOutMongo

	ch := make(chan int)

	go r.askMongo(ch, &regMongo)
	go r.isUserHisOk(ch, &userHis)

	for i := 0; i < 2; i++ {
		select {
		case code = <-ch:
			if code != http.StatusOK {
				return
			}
		}
	}

	for i := 0; i < len(regMongo); i++ {
		var resp response.RegOutResp
		regMongo[i].ToResp(&resp)
		*respList = append(*respList, resp)
	}

	return http.StatusOK
}

func (r *RegularOutReq) askMongo(ch chan int, regMongoList *[]mongo_models.RegularOutMongo) (code int) {

	var regHis mongo_models.RegularOutHistoryMongo
	regHis.UserId = r.UserId

	ok, empty := regHis.AskMongo(&r.UserId, &r.RegularId)

	if !ok {
		ch <- http.StatusInternalServerError
		return
	}

	if empty {
		ch <- http.StatusNotFound
		return
	}

	if r.Offset >= len(regHis.IdS) {
		ch <- http.StatusOK
		return
	}

	var newList []mongo_models.RegHisIdMongo

	if r.Offset+10 < len(regHis.IdS) {
		newList = regHis.IdS[r.Offset : r.Offset+10]
	} else {
		newList = regHis.IdS[r.Offset:]
	}

	chBool := make(chan bool)
	var count = 0
	for ; count < len(newList); count++ {
		go r.getReg(chBool, regMongoList, &newList[count].Id)
	}

	for i := 0; i < count; i++ {
		select {
		case ok := <-chBool:
			if !ok {
				log.Println("regular not edded")
			}
		}
	}

	if !ok {
		ch <- http.StatusInternalServerError
		return
	}

	if empty {
		ch <- http.StatusNotFound
		return
	}

	ch <- http.StatusOK
	return
}

func (r *RegularOutReq) getReg(ch chan bool, regMongoList *[]mongo_models.RegularOutMongo, MongoId *primitive.ObjectID) {
	var regMongo mongo_models.RegularOutMongo
	ok, empty := regMongo.GetMongo(MongoId)

	*regMongoList = append(*regMongoList, regMongo)

	if !ok {
		ch <- false
		return
	}

	if empty {
		ch <- false
		return
	}

	ch <- true
	return
}

func (r *RegularOutReq) isUserHisOk(ch chan int, userHis *mongo_models.UserHistoryMongo) {
	ok, empty := userHis.AskMongo(r.UserId)
	if !ok {
		ch <- http.StatusInternalServerError
		return
	}

	if empty {
		ch <- http.StatusNotFound
		return
	}

	for i := range userHis.OutRegular {
		if userHis.OutRegular[i].MongoId.Hex() == r.RegularId {
			ch <- http.StatusOK
			return
		}
	}
	ch <- http.StatusNotFound
	return
}

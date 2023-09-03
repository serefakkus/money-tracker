package mongo_models

import (
	"get/dataloaders/mongo"
	"get/models/response"
	"get/set"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserHistoryMongo struct {
	Id            primitive.ObjectID `bson:"_id"`
	UserId        string
	IncomingYears []HistoryYearsMongo
	OutgoingYears []HistoryYearsMongo
	InRegular     []RegularHistoryMongo
	OutRegular    []RegularHistoryMongo
}

type HistoryYearsMongo struct {
	Year   string
	Months []HistoryMonthsMongo
}

type HistoryMonthsMongo struct {
	Month string
	Days  []HistoryDaysMongo
}

type HistoryDaysMongo struct {
	Day     string
	Mongoid primitive.ObjectID
}

type RegularHistoryMongo struct {
	MongoId  primitive.ObjectID
	Category string
}

func (u *UserHistoryMongo) AskMongo(userId string) (ok bool, empty bool) {
	var filter filterUser
	filter.UserId = userId

	return mongo.AskMongo(&filter, u, &set.UserHistoryTableName, &set.UserHistoryDBName)
}

func (u *UserHistoryMongo) ToResp(r *response.RespUserHistory) bool {
	var inList []primitive.ObjectID
	var outList []primitive.ObjectID

	var year response.HistoryYears
	for i := range u.IncomingYears {
		year.Year = u.IncomingYears[i].Year
		year.Months = nil
		u.IncomingYears[i].ToResp(&year, true, &inList, &outList)
		r.IncomingYears = append(r.IncomingYears, year)
	}

	for i := range u.OutgoingYears {
		year.Year = u.OutgoingYears[i].Year
		year.Months = nil
		u.OutgoingYears[i].ToResp(&year, false, &inList, &outList)
		r.OutgoingYears = append(r.OutgoingYears, year)
	}

	var reg response.RegularHistory
	for i := range u.InRegular {
		u.InRegular[i].ToResp(&reg)
		r.InRegular = append(r.InRegular, reg)
	}

	for i := range u.OutRegular {
		u.OutRegular[i].ToResp(&reg)
		r.OutRegular = append(r.OutRegular, reg)
	}

	return true
}

func (mon *HistoryYearsMongo) ToResp(year *response.HistoryYears, isInCom bool, inList *[]primitive.ObjectID, outList *[]primitive.ObjectID) {
	mon.Year = year.Year
	var month response.HistoryMonths
	for i := range mon.Months {
		mon.Months[i].ToResp(&month, isInCom, inList, outList)
		year.Months = append(year.Months, month)
	}
}

func (mon *HistoryMonthsMongo) ToResp(month *response.HistoryMonths, isInCom bool, inList *[]primitive.ObjectID, outList *[]primitive.ObjectID) {
	month.Month = mon.Month
	var day response.HistoryDays
	for i := range mon.Days {
		mon.Days[i].ToResp(&day, isInCom, inList, outList)
		month.Days = append(month.Days, day)
	}
}

func (mon *HistoryDaysMongo) ToResp(day *response.HistoryDays, isInCom bool, inList *[]primitive.ObjectID, outList *[]primitive.ObjectID) {
	day.Day = mon.Day
	var dayMongo DayMongo

	dayMongo.GetDays(&mon.Mongoid, isInCom)
	var adding = false
	if isInCom {
		if len(*inList) < 10 {
			adding = true
		}
	} else {
		if len(*outList) < 10 {
			adding = true
		}
	}
	for i := range dayMongo.Ids {
		day.Ids = append(day.Ids, dayMongo.Ids[i].Hex())
		if adding {
			if isInCom {
				*inList = append(*inList, dayMongo.Ids[i])
			} else {
				*outList = append(*outList, dayMongo.Ids[i])
			}
		}
	}
}

func (mon *RegularHistoryMongo) ToResp(regular *response.RegularHistory) {
	regular.MongoId = mon.MongoId.Hex()
	regular.Category = mon.Category
}

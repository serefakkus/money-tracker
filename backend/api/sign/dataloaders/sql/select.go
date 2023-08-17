package sql

import (
	"database/sql"
	"encoding/json"
	"errors"
	"github.com/blockloop/scan/v2"
	"sign/helpers"
	"sign/interfaces"
	"sign/set"
	"strconv"
)

func SelectSql(resultData interfaces.IDataLoader, filterData interfaces.IDataLoader, tableName *string) (ok bool, empy bool) {

	if !(helpers.GetModelLengt(filterData) > 0) {
		err := errors.New("filter data is empty")
		helpers.CheckErr(err)
		return false, false
	}

	var queryStr string
	var queryVal []any
	var resKey []string
	ok, queryStr = sqlSelectStr(resultData, &queryVal, filterData, &resKey, tableName)

	if !ok {
		return
	}

	db, err := sql.Open("postgres", set.PsqlInfo)

	if !helpers.CheckErr(err) {
		ok = false
		return
	}

	defer func(db *sql.DB) {
		err := db.Close()
		helpers.CheckErr(err)

	}(db)

	row, err2 := db.Query(queryStr, queryVal...)

	if !helpers.CheckErr(err2) {

		ok = false
		return
	}

	err = scan.Row(resultData, row)

	switch err {

	case sql.ErrNoRows:
		ok = true
		empy = true

	case nil:
		ok = true
		empy = false

	default:
		helpers.CheckErr(err)
		ok = false
	}

	return
}

func sqlSelectStr(resultData interfaces.IDataLoader, queryValList *[]any, filterData interfaces.IDataLoader, resKeyList *[]string, tableName *string) (ok bool, selectStr string) {

	js, err := json.Marshal(&resultData)
	if !helpers.CheckErr(err) {
		return
	}
	var mapJs map[string]interface{}

	err = json.Unmarshal(js, &mapJs)
	if !helpers.CheckErr(err) {
		return
	}

	selectStr = "SELECT "

	for k := range mapJs {
		selectStr += k + ","

		*resKeyList = append(*resKeyList, k)

	}

	selectStr = selectStr[:len(selectStr)-1]

	selectStr += " FROM " + *tableName + " WHERE "

	js, err = json.Marshal(&filterData)
	if !helpers.CheckErr(err) {
		return
	}

	var mapJs2 map[string]interface{}

	err = json.Unmarshal(js, &mapJs2)
	if !helpers.CheckErr(err) {
		return
	}

	var count = 0
	var length = len(mapJs2)

	for k, v := range mapJs2 {
		count++

		selectStr += k + "=$" + strconv.Itoa(count)

		if count != 1 && count != length {
			selectStr += " AND "
		}

		if length == 2 && count == 1 {
			selectStr += " AND "
		}

		*queryValList = append(*queryValList, v)
	}

	selectStr += ";"

	ok = true

	return
}

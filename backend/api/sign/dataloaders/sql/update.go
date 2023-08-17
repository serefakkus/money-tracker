package sql

import (
	"database/sql"
	"encoding/json"
	"errors"
	"github.com/lib/pq"
	"reflect"
	"sign/helpers"
	"sign/interfaces"
	"sign/set"
	"strconv"
)

func UpdateSql(newData interfaces.IDataLoader, filterData interfaces.IDataLoader, tableName *string) (ok bool, empy bool) {

	if !(helpers.GetModelLengt(filterData) > 0) {
		err := errors.New("filter data is empty")
		helpers.CheckErr(err)

		return false, false
	}

	if !(helpers.GetModelLengt(newData) > 0) {
		err := errors.New("new data is empty")
		helpers.CheckErr(err)
		return false, false
	}

	var queryStr string
	var ValList []any

	ok, queryStr = sqlUpdateStr(newData, &ValList, filterData, tableName)

	if !ok {
		return
	}

	db, err := sql.Open("postgres", set.PsqlInfo)

	if !helpers.CheckErr(err) {
		return
	}

	defer func(db *sql.DB) {
		err := db.Close()
		helpers.CheckErr(err)

	}(db)

	stmt, err := db.Prepare(queryStr)

	switch err := stmt.QueryRow(ValList...).Err(); err {

	case sql.ErrNoRows:

		ok = true
		empy = true
		return

	case nil:

		ok = true
		empy = false

		return

	default:
		helpers.CheckErr(err)
		ok = false
		return

	}
}

func sqlUpdateStr(newData interfaces.IDataLoader, valList *[]any, filterData interfaces.IDataLoader, tableName *string) (ok bool, selectStr string) {

	js, err := json.Marshal(&newData)
	if !helpers.CheckErr(err) {
		return
	}
	var mapJs map[string]interface{}

	err = json.Unmarshal(js, &mapJs)
	if !helpers.CheckErr(err) {
		return
	}

	selectStr = "UPDATE " + *tableName + " SET "

	var count = 0

	var length = len(mapJs)

	for k, v := range mapJs {
		count++

		selectStr += k + "=$" + strconv.Itoa(count)

		if reflect.TypeOf(v).Kind().String() == "slice" {
			v = pq.Array(v)
		}

		if count == length {
			*valList = append(*valList, v)
			continue
		}

		selectStr += ","

		*valList = append(*valList, v)
	}

	selectStr += " WHERE "

	js, err = json.Marshal(&filterData)
	if !helpers.CheckErr(err) {
		return
	}

	var mapJs2 map[string]interface{}

	err = json.Unmarshal(js, &mapJs2)
	if !helpers.CheckErr(err) {
		return
	}

	var filterLen = len(mapJs2)

	for k, v := range mapJs2 {
		count++

		selectStr += k + "=$" + strconv.Itoa(count)

		if filterLen == 1 {
			*valList = append(*valList, v)
			continue
		}

		if count != 1 && count != length {
			selectStr += " AND "
		}

		if length == 2 && count == 1 {
			selectStr += " AND "
		}

		*valList = append(*valList, v)
	}

	selectStr += ";"

	ok = true

	return
}

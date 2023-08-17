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

func InsertSql(dataLoader interfaces.IDataLoader, tableName *string) bool {

	if !(helpers.GetModelLengt(dataLoader) > 0) {
		err := errors.New("filter data is empty")
		helpers.CheckErr(err)
		return false
	}

	db, err := sql.Open("postgres", set.PsqlInfo)

	if !helpers.CheckErr(err) {
		return false
	}

	defer func(db *sql.DB) {
		err := db.Close()
		if !helpers.CheckErr(err) {

		}
	}(db)

	var sqlQueryStr string
	var valList []any
	ok := sqlInsertStr(tableName, dataLoader, &sqlQueryStr, &valList)

	if !ok {
		return false
	}

	_, err2 := db.Exec(sqlQueryStr, valList...)

	if !helpers.CheckErr(err2) {
		return false
	}

	return true
}

func sqlInsertStr(tableName *string, dataLoader interfaces.IDataLoader, sqlQueryStr *string, valList *[]any) (ok bool) {

	js, err := json.Marshal(&dataLoader)
	if !helpers.CheckErr(err) {
		return
	}
	var mapJs map[string]interface{}

	err = json.Unmarshal(js, &mapJs)
	if !helpers.CheckErr(err) {
		return
	}

	var kCount = 0
	*sqlQueryStr = "INSERT INTO " + *tableName + "("

	for k, v := range mapJs {
		*sqlQueryStr += k + ","
		kCount++

		if v != nil {
			if reflect.TypeOf(v).Kind().String() == "slice" {
				v = pq.Array(v)
			}
		}

		*valList = append(*valList, v)
	}

	var str = *sqlQueryStr
	*sqlQueryStr = str[:len(str)-1]

	*sqlQueryStr += " ) VALUES ("

	for i := 0; i < kCount; i++ {
		*sqlQueryStr += "$" + strconv.Itoa(i+1) + ","
	}

	str = *sqlQueryStr
	*sqlQueryStr = str[:len(str)-1]
	*sqlQueryStr += ");"

	ok = true

	return
}

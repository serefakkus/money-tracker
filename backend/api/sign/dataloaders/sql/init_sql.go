package sql

import (
	"database/sql"
	_ "github.com/lib/pq"
	"sign/helpers"
	"sign/set"
)

func InitPhoneSql() bool {
	db, err := sql.Open("postgres", set.PsqlInfo)

	if !helpers.CheckErr(err) {
		panic(err)
		return false
	}

	defer func(db *sql.DB) {
		err := db.Close()
		helpers.CheckErr(err)
	}(db)

	_, err = db.Exec("CREATE TABLE IF NOT EXISTS " + set.SmsSqlTableName + " (Phone varchar unique,Code varchar,Time varchar,Ok boolean)")
	if !helpers.CheckErr(err) {
		panic(err)
		return false
	}

	return true
}

func InitPhoneLogSql() bool {
	db, err := sql.Open("postgres", set.PsqlInfo)

	if !helpers.CheckErr(err) {
		panic(err)
		return false
	}

	defer func(db *sql.DB) {
		err := db.Close()
		helpers.CheckErr(err)
	}(db)

	_, err = db.Exec("CREATE TABLE IF NOT EXISTS " + set.SmsSqlLogTableName + " (Phone varchar unique,Times varchar array)")
	if !helpers.CheckErr(err) {
		panic(err)
		return false
	}

	return true
}

func InitEveryThings() {
	err := set.InitAllStrings()
	if !helpers.CheckErr(err) {
		panic(err)
	}

	InitPhoneSql()

	InitPhoneLogSql()
}

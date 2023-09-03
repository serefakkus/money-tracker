package helpers

import (
	"encoding/json"
	"get/interfaces"
)

func GetModelLengt(data interfaces.IDataLoader) (count int) {

	js, err := json.Marshal(&data)
	if !CheckErr(err) {
		return
	}
	var mapJs map[string]interface{}

	err = json.Unmarshal(js, &mapJs)
	if !CheckErr(err) {
		return
	}

	count = 0

	for range mapJs {
		count++
	}
	return
}

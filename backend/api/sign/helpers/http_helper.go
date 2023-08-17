package helpers

import (
	"bytes"
	"errors"
	"io"
	"net/http"
	"strconv"
)

func SendToDbPost(message *[]byte, uri string) (code int, ok bool) {

	r, err := http.NewRequest("POST", uri, bytes.NewBuffer(*message))
	if !CheckErr(err) {
		return http.StatusBadGateway, false
	}

	r.Header.Add("Content-Type", "application/json")

	var client http.Client

	res, err := client.Do(r)

	if !CheckErr(err) {
		return http.StatusBadGateway, false
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		CheckErr(err)
	}(res.Body)

	if res.StatusCode != http.StatusOK {
		err = errors.New("post url = " + uri + " - http status = " + strconv.Itoa(res.StatusCode) + " - " + "data = " + string(*message))
		return http.StatusBadGateway, false
	}

	*message, err = io.ReadAll(res.Body)
	if !CheckErr(err) {
		return http.StatusInternalServerError, false
	}

	return http.StatusOK, true

}

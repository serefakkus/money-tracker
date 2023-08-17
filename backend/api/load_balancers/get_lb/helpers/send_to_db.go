package helpers

import (
	"bytes"
	"encoding/json"
	"get_lb/interfaces"
	"io"
	"net/http"
)

func SendToServerGet(respData interfaces.IResponseServer, uri string, auth string) (ok bool, code int) {

	client := &http.Client{}

	req, err := http.NewRequest(http.MethodGet, uri, nil)

	if !CheckErr(err) {
		return false, http.StatusBadGateway
	}

	req.Header.Add("Authorization", auth)

	resp, err := client.Do(req)

	if !CheckErr(err) {
		return false, http.StatusBadGateway
	}

	body, err := io.ReadAll(resp.Body)

	if !CheckErr(err) {
		return false, resp.StatusCode
	}

	err2 := json.Unmarshal(body, respData)

	return CheckErr(err2), resp.StatusCode
}

func SendToServerPost(req interfaces.IRequestServer, respData *[]byte, uri string, path string) (ok bool, code int) {

	reqData := req.ToJson()

	responseBody := bytes.NewBuffer(reqData)

	post, err := http.Post(uri+path, "application/json", responseBody)

	if !CheckErr(err) {
		return false, http.StatusInternalServerError
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			CheckErr(err)
		}
	}(post.Body)

	*respData, err = io.ReadAll(post.Body)

	if post.StatusCode != http.StatusOK {
		return false, post.StatusCode
	}

	if !CheckErr(err) {
		return false, http.StatusInternalServerError
	}

	return true, http.StatusOK
}

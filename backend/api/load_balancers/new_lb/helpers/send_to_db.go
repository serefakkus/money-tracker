package helpers

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"new_lb/interfaces"
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

func SendToServerPost(req interfaces.IRequestServer, resp interfaces.IResponseServer, uri string, path string) (ok bool, code int) {

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

	respData, err := io.ReadAll(post.Body)

	if !resp.FromJson(&respData) {
		return false, http.StatusBadGateway
	}

	if post.StatusCode != http.StatusOK {
		return false, post.StatusCode
	}

	if !CheckErr(err) {
		return false, http.StatusInternalServerError
	}

	return true, http.StatusOK
}

func SendToServerPostWithOutResp(req interfaces.IRequestServer, uri string, path string) (ok bool, code int) {
	reqData := req.ToJson()

	responseBody := bytes.NewBuffer(reqData)

	post, err := http.Post(uri+path, "application/json", responseBody)

	if !CheckErr(err) {
		return false, http.StatusInternalServerError
	}

	if post.StatusCode != http.StatusOK {
		return false, post.StatusCode
	}

	return true, http.StatusOK
}

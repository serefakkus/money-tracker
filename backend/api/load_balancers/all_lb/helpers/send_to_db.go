package helpers

import (
	"all_lb/interfaces"
	"bytes"
	"io"
	"log"
	"net/http"
)

/*
func SendToServerGet(resp *[]byte, uri string, auth string) (ok bool, code int) {

	post, err := http.Get(uri)

	post.Header.Add("Authorization", auth)

	if !CheckErr(err) {
		return false, post.StatusCode
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			CheckErr(err)
		}
	}(post.Body)

	body, err := io.ReadAll(post.Body)

	if !CheckErr(err) {
		return false, post.StatusCode
	}

	err2 := json.Unmarshal(body, resp)

	return CheckErr(err2), post.StatusCode
}
*/

func SendToServerPost(req interfaces.IRequestServer, respJson *string, uri string, path string) (ok bool, code int) {

	reqData := req.ToJson()

	responseBody := bytes.NewBuffer(reqData)

	post, err := http.Post(uri+"/"+path, "application/json", responseBody)

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

	*respJson = string(respData)

	log.Println(*respJson)

	if !CheckErr(err) {
		return false, http.StatusInternalServerError
	}

	if post.StatusCode != http.StatusOK {
		return false, post.StatusCode
	}

	return true, http.StatusOK
}

func SendToServerPostWithOutResp(req interfaces.IRequestServer, uri string, path string) (ok bool, code int) {

	reqData := req.ToJson()

	responseBody := bytes.NewBuffer(reqData)

	post, err := http.Post(uri+"/"+path, "application/json", responseBody)

	if !CheckErr(err) {
		return false, http.StatusInternalServerError
	}

	if post.StatusCode != http.StatusOK {
		return false, post.StatusCode
	}

	return true, http.StatusOK
}

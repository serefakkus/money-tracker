package helpers

func SendSms(phone *string, code *string) (ok bool) {

	/*

		body := make(map[string]interface{})

		body["api_id"] = "290b4961e5935759b46181dc"
		body["api_key"] = "3c793bf45465bf28f19d7d37"
		body["sender"] = "S.AKKUS"
		body["message_type"] = "normal"
		body["message"] = "ssshht reklam aktivasyon kodunuz : " + *code + "."
		body["phones"] = []string{*phone}

		requestBody, err := json.Marshal(body)
		if !error_log.CheckErr(err) {
			return false
		}

		rawResponse, err := http.Post("https://api.vatansms.net/api/v1/1toN", "application/json", bytes.NewBuffer(requestBody))
		if !error_log.CheckErr(err) {
			return false
		}

		response := make(map[string]interface{})

		err = json.NewDecoder(rawResponse.Body).Decode(&response)

		if !error_log.CheckErr(err) {
			return false
		}

		status := response["status"]

		if status != "success" {
			var errMessage = fmt.Sprintf("send sms error resp = %v", status)
			err = errors.New(errMessage)
			return false
		}

	*/

	return true
}

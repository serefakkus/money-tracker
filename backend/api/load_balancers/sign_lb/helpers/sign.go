package helpers

func SignInControl(phone *string, pass *string) (ok bool) {
	if len(*phone) != 10 || len(*pass) < 6 {
		return false
	}

	return true
}

func SignUpControl(phone *string, pass *string, code *string) (ok bool) {
	if len(*phone) != 10 || len(*pass) < 6 || len(*code) != 5 {
		return false
	}

	return true
}

func PhoneControl(phone *string) (ok bool) {
	if len(*phone) != 10 {
		return false
	}

	return true
}

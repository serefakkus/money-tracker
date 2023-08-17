package helpers

func CheckAuth(auth *string) (ok bool) {
	if *auth != "" {
		ok = true
	} else {
		ok = false
	}
	return
}

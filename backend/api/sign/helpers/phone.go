package helpers

import (
	"strconv"
)

func CheckPhoneNo(phoneNo string) bool { //phone number control use for only turkey
	return len(phoneNo) == 10 && string(phoneNo[0]) == "5"
}

func CheckPass(pass string) bool { //phone number control use for only turkey
	return len(pass) > 5
}

func CheckCode(code string) bool {
	_, err := strconv.Atoi(code)
	if !CheckErr(err) {
		return false
	}

	return len(code) == 5
}

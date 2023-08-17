package helpers

import "new_lb/set"

func GetUriToken(path string) string {
	return set.TokenUri + "/" + path
}

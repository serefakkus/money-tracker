package helpers

import "get_lb/set"

func GetUriToken(path string) string {
	return set.TokenUri + "/" + path
}

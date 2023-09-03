package helpers

import "all_lb/set"

func GetUriToken(path string) string {
	return set.TokenUri + "/" + path
}

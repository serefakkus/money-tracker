package helpers

import "math/rand"

var strRunes = []rune("0123456789qwertyuioplkjhgfdsazxcvbnmQWERTYUIOPLKJHGFDSAZXCVBNM")

func RandStrRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = strRunes[rand.Intn(len(strRunes))]
	}
	return string(b)
}

var codeRunes = []rune("0123456789")

func RandCodeRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = codeRunes[rand.Intn(len(codeRunes))]
	}
	return string(b)
}

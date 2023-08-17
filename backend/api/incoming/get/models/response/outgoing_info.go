package response

type OutGoInfo struct {
	RegularId string
	Category  string
	Emoji     string
	Not       string
	Amount    float64
	Time      string
	Date      string
	To        FromUser
}

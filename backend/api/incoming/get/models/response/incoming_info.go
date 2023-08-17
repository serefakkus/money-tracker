package response

type InComInfo struct {
	RegularId string
	Category  string
	Emoji     string
	Not       string
	Amount    float64
	Time      string
	Date      string
	From      FromUser
}

type FromUser struct {
	Name   string
	UserId string
}

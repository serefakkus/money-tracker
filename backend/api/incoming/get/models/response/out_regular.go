package response

type RegOutResp struct {
	Emoji  string
	Not    string
	Amount float64
	Time   string
	To     FromUser
}

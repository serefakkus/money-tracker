package response

type RegHisOutResp struct {
	Category    string
	RegularId   string
	IntervalDay int
	Time        string
	IdS         []RegHisIdSResp
}

type RegHisIdSResp struct {
	Id   string
	Date string
}

package response

type RegHisInResp struct {
	Category    string
	RegularId   string
	IntervalDay int
	Time        string
	IdS         []RegHisIdSResp
}

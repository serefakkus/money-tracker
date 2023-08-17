package response

type RespUserHistory struct {
	IncomingYears []HistoryYears
	OutgoingYears []HistoryYears
	InRegular     []RegularHistory
	OutRegular    []RegularHistory
}

type HistoryYears struct {
	Year   string
	Months []HistoryMonths
}

type HistoryMonths struct {
	Month string
	Days  []HistoryDays
}

type HistoryDays struct {
	Day string
	Ids []string
}

type RegularHistory struct {
	MongoId  string
	Category string
}

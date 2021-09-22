package model

type Opinion struct {
	QID      string `json:"qid"`
	OptionID string `json:"oid"`
	DateTime string `json:"date_time"`
}

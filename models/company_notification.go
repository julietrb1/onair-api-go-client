package models

type CompanyNotification struct {
	Id             string  `json:"Id"`
	CompanyID      string  `json:"CompanyId"`
	FlightID       string  `json:"FlightId,omitempty"`
	IsRead         bool    `json:"IsRead"`
	IsNotification bool    `json:"IsNotification"`
	ZuluEventTime  string  `json:"ZuluEventTime"`
	Category       int     `json:"Category"`
	Action         int     `json:"Action"`
	Description    string  `json:"Description"`
	Amount         float64 `json:"Amount"`
	AccountID      string  `json:"AccountId,omitempty"`
	PeopleID       string  `json:"PeopleId,omitempty"`
}

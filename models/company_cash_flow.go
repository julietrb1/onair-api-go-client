package models

type CompanyCashFlow struct {
	Entries            []CompanyCashFlowEntry `json:"Entries"`
	CompanyCurrentCash float64                `json:"CompanyCurrentCash"`
	LastReportAmount   float64                `json:"LastReportAmount"`
	LastReportDate     string                 `json:"LastReportDate"`
}

type CompanyCashFlowEntry struct {
	ID           string  `json:"Id"`
	CompanyID    string  `json:"CompanyId"`
	AccountID    string  `json:"AccountId"`
	Amount       float64 `json:"Amount"`
	CreationDate string  `json:"CreationDate"`
	Description  string  `json:"Description"`
	CarryForward bool    `json:"CarryForward"`
	AircraftID   string  `json:"AircraftId,omitempty"`
	Key          string  `json:"Key,omitempty"`
	FboID        string  `json:"FBOId,omitempty"`
	PeopleID     string  `json:"PeopleId,omitempty"`
}

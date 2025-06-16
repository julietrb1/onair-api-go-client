package models

type CompanyIncomeStatement struct {
	REVAccounts     []CompanyIncomeStatementAccount `json:"REVAccounts"`
	EXPAccounts     []CompanyIncomeStatementAccount `json:"EXPAccounts"`
	REVAmount       float64                         `json:"REVAmount"`
	EXPAmount       float64                         `json:"EXPAmount"`
	NetIncomeAmount float64                         `json:"NetIncomeAmount"`
}

type CompanyIncomeStatementAccount struct {
	Name      string                        `json:"Name"`
	ShortName string                        `json:"ShortName"`
	Entries   []CompanyIncomeStatementEntry `json:"Entries"`
	Order     int                           `json:"Order"`
	Amount    float64                       `json:"Amount"`
}

type CompanyIncomeStatementEntry struct {
	Id           string  `json:"Id"`
	CompanyID    string  `json:"CompanyId"`
	AccountID    string  `json:"AccountId"`
	Amount       float64 `json:"Amount"`
	Account      Account `json:"Account"`
	CreationDate string  `json:"CreationDate"`
	Description  string  `json:"Description"`
	CarryForward bool    `json:"CarryForward"`
	AircraftID   string  `json:"AircraftId,omitempty"`
	FBOId        string  `json:"FBOId,omitempty"`
}

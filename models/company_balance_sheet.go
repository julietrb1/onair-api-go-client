package models

type CompanyBalanceSheet struct {
	AssetAccounts     []AccountingGroup `json:"ASSAccounts"`
	LiabilityAccounts []AccountingGroup `json:"LIAAccounts"`
	AssetAmount       float64           `json:"ASSAmount"`
	LiabilityAmount   float64           `json:"LIAAmount"`
	DeltaBalance      float64           `json:"DeltaBalance"`
}

type AccountingGroup struct {
	Name      string            `json:"Name"`
	ShortName string            `json:"ShortName"`
	Entries   []AccountingEntry `json:"Entries"`
	Order     int               `json:"Order"`
	Amount    float64           `json:"Amount"`
}

type AccountingEntry struct {
	Id           string  `json:"Id"`
	CompanyId    string  `json:"CompanyId"`
	AccountId    string  `json:"AccountId"`
	Amount       float64 `json:"Amount"`
	Account      Account `json:"Account"`
	CreationDate string  `json:"CreationDate"`
	Description  string  `json:"Description"`
	CarryForward bool    `json:"CarryForward"`
}

type Account struct {
	Id        string `json:"Id"`
	ShortName string `json:"ShortName"`
	Name      string `json:"Name"`
	Category  int    `json:"Category"`
	Order     int    `json:"Order"`
}

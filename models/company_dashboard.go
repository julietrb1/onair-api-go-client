package models

type CompanyDashboard struct {
	Cash                      float64 `json:"Cash"`
	ShareCapital              float64 `json:"ShareCapital"`
	Value                     float64 `json:"Value"`
	Income1Week               float64 `json:"Incomes1week"`
	Income2Weeks              float64 `json:"Incomes2weeks"`
	IncomeGlobal              float64 `json:"IncomesGlobal"`
	NumberOfCompletedMissions int     `json:"NumberOfCompletedMissions"`
	NumberOfActiveMissions    int     `json:"NumberOfActivesMissions"`
	NumberOfActiveAircraft    int     `json:"NumberOfActiveAircrafts"`
	NumberOfAircraft          int     `json:"NumberOfAircrafts"`
	NumberOfEmployees         int     `json:"NumberOfEmployees"`
	NumberOfFBOs              int     `json:"NumberOfFBOs"`
	ReturnOnAssets            float64 `json:"ReturnOnAssets"`
	DebtRatio                 float64 `json:"DebtRatio"`
	Assets                    float64 `json:"Assets"`
	Loans                     float64 `json:"Loans"`
	Level                     int     `json:"Level"`
	XP                        int     `json:"XP"`
	XPNeeded                  int     `json:"XPNeeded"`
}

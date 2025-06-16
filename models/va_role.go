package models

type VARole struct {
	ID                                     string  `json:"Id"`
	VaID                                   string  `json:"VAId"`
	Name                                   string  `json:"Name"`
	Permission                             int     `json:"Permission"`
	IsDefaultNewRole                       bool    `json:"IsDefaultNewRole"`
	Color                                  string  `json:"Color"`
	PayPercent                             float64 `json:"PayPercent"`
	IsHidden                               bool    `json:"IsHidden"`
	RestrictLoadingVAJobsIntoNonVAAircraft bool    `json:"RestrictLoadingVAJobsIntoNonVAAircraft"`
	RestrictLoadingNonVAJobsIntoVAAircraft bool    `json:"RestrictLoadingNonVAJobsIntoVAAircraft"`
	PayWeekly                              float64 `json:"PayWeekly"`
	PayPerFlightHour                       float64 `json:"PayPerFlightHour"`
}

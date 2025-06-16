package models

type Job struct {
	ID                              string      `json:"Id"`
	WorldID                         string      `json:"WorldId"`
	MissionTypeID                   string      `json:"MissionTypeId"`
	MissionType                     MissionType `json:"MissionType"`
	MainAirportID                   string      `json:"MainAirportId"`
	BaseAirportID                   string      `json:"BaseAirportId"`
	CompletionFlightID              string      `json:"CompletionFlightId"`
	ValuePerLbsPerDistance          float64     `json:"ValuePerLbsPerDistance"`
	IsGoodValue                     bool        `json:"IsGoodValue"`
	MaxDistance                     float64     `json:"MaxDistance"`
	TotalDistance                   float64     `json:"TotalDistance"`
	MainAirportHeading              float64     `json:"MainAirportHeading"`
	Description                     string      `json:"Description"`
	ExpirationDate                  string      `json:"ExpirationDate,omitempty"`
	Pay                             float64     `json:"Pay"`
	PayAirportBoostBonus            float64     `json:"PayAirportBoostBonus,omitempty"`
	Penalty                         float64     `json:"Penality"`
	ReputationImpact                float64     `json:"ReputationImpact"`
	CompanyID                       string      `json:"CompanyId"`
	CreationDate                    string      `json:"CreationDate"`
	TakenDate                       string      `json:"TakenDate"`
	CompletionDate                  string      `json:"CompletionDate"`
	PaidDate                        string      `json:"PaidDate"`
	TotalCargoTransported           float64     `json:"TotalCargoTransported"`
	TotalPaxTransported             int         `json:"TotalPaxTransported"`
	Category                        int         `json:"Category"`
	State                           int         `json:"State"`
	XP                              int         `json:"XP"`
	SkillPoint                      int         `json:"SkillPoint"`
	MinCompanyReputation            float64     `json:"MinCompanyReput"`
	Hash                            string      `json:"Hash,omitempty"`
	RealPay                         float64     `json:"RealPay"`
	RealPenalty                     float64     `json:"RealPenality"`
	CanAccess                       bool        `json:"CanAccess"`
	CanAccessDisplay                string      `json:"CanAccessDisplay,omitempty"`
	IsLastMinute                    bool        `json:"IsLastMinute"`
	IsFavorite                      bool        `json:"IsFavorited"`
	PayAirportBoostBonusDescription string      `json:"PayAirportBoostBonusDescription,omitempty"`
	PayLastMinuteBonus              float64     `json:"PayLastMinuteBonus,omitempty"`
	FBOLogisticQueryID              string      `json:"FBOLogisticQueryId,omitempty"`
	TourID                          string      `json:"TourId,omitempty"`
	CustomID                        string      `json:"CustomId,omitempty"`
}

type MissionType struct {
	ID                   string  `json:"Id"`
	Name                 string  `json:"Name"`
	ShortName            string  `json:"ShortName"`
	Description          string  `json:"Description"`
	BaseReputationImpact float64 `json:"BaseReputationImpact"`
	BasePayFactor        float64 `json:"BasePayFactor"`
	BasePenalityFactor   float64 `json:"BasePenalityFactor"`
}

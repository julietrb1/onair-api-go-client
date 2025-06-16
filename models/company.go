package models

type Company struct {
	ID                                 string  `json:"Id"`
	WorldID                            string  `json:"WorldId"`
	Name                               string  `json:"Name"`
	AirlineCode                        string  `json:"AirlineCode"`
	LastConnection                     string  `json:"LastConnection"`
	LastReportDate                     string  `json:"LastReportDate"`
	Reputation                         float64 `json:"Reputation"`
	CreationDate                       string  `json:"CreationDate"`
	DifficultyLevel                    int     `json:"DifficultyLevel"`
	UTCOffsetInHours                   float64 `json:"UTCOffsetinHours"`
	Paused                             bool    `json:"Paused"`
	Level                              int     `json:"Level"`
	LevelXP                            int     `json:"LevelXP"`
	TransportEmployeeInstant           bool    `json:"TransportEmployeeInstant"`
	TransportPlayerInstant             bool    `json:"TransportPlayerInstant"`
	ForceTimeInSimulator               bool    `json:"ForceTimeInSimulator"`
	UseSmallAirports                   bool    `json:"UseSmallAirports"`
	UseOnlyVanillaAirports             bool    `json:"UseOnlyVanillaAirports"`
	EnableSkillTree                    bool    `json:"EnableSkillTree"`
	CheckrideLevel                     int     `json:"CheckrideLevel"`
	EnableLandingPenalities            bool    `json:"EnableLandingPenalities"`
	EnableEmployeesFlightDutyAndSleep  bool    `json:"EnableEmployeesFlightDutyAndSleep"`
	AircraftRentLevel                  int     `json:"AircraftRentLevel"`
	EnableCargosAndChartersLoadingTime bool    `json:"EnableCargosAndChartersLoadingTime"`
	InSurvival                         bool    `json:"InSurvival"`
	PayBonusFactor                     float64 `json:"PayBonusFactor"`
	EnableSimFailures                  bool    `json:"EnableSimFailures"`
	DisableSeatsConfigCheck            bool    `json:"DisableSeatsConfigCheck"`
	RealisticSimProcedures             bool    `json:"RealisticSimProcedures"`
	TravelTokens                       int     `json:"TravelTokens"`
	LastWeeklyManagementsPaymentDate   string  `json:"LastWeeklyManagementsPaymentDate"`
	SkillTreeResetCount                int     `json:"SkillTreeResetCount"`
	IndustryPoints                     int     `json:"IndustryPoints"`
	TotalIndustryPoints                int     `json:"TotalIndustryPoints"`
	TotalContractsCompleted            int     `json:"TotalContractsCompleted"`
	TotalContractsEarnedCredits        float64 `json:"TotalContractsEarnedCredits"`
	CompanyType                        int     `json:"CompanyType"`
	Callsign                           string  `json:"Callsign"`
	AutoUnloadMerchandises             bool    `json:"AutoUnloadMerchandises"`
	ExcludeAirportsNotInSimbrief       bool    `json:"ExcludeAirportsNotInSimbrief"`
	DisableFlightScoring               bool    `json:"DisableFlightScoring"`
	EnableFlightScoringLights          bool    `json:"EnableFlightScoringLights"`
	ManagementMode                     int     `json:"ManagementMode"`
	IsAdvancedManagementModeUnlocked   bool    `json:"IsAdvancedManagementModeUnlocked"`
}

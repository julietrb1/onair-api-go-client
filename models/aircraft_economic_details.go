package models

type AircraftEconomicDetails struct {
	AircraftType                       AircraftTypeAtAirport `json:"AircraftType"`
	LowestCargoPayCrPerLbsPerNM        float64               `json:"LowestCargoPayCrPerLbsPerNM"`
	OperationalPayloadLbs              float64               `json:"OperationalPayloadLbs"`
	AverageCruiseSpeedKts              float64               `json:"AverageCruiseSpeedKts"`
	AverageWeeklyFlightHours           float64               `json:"AverageWeeklyFlightHours"`
	AverageAnnualFlightHours           float64               `json:"AverageAnnualFlightHours"`
	AverageWeeklyFlightHoursCompressed float64               `json:"AverageWeeklyFlightHoursCompressed"`
	AverageAnnualFlightHoursCompressed float64               `json:"AverageAnnualFlightHoursCompressed"`
	CostHourlyFuel                     float64               `json:"CostHourlyFuel"`
	CostHourlyCrew                     float64               `json:"CostHourlyCrew"`
	HourlyProfit                       float64               `json:"HourlyProfit"`
	AverageDeltaConditionPerFlightHour float64               `json:"AverageDeltaConditionPerFlightHour"`
	CostAirframeRepair1P               float64               `json:"CostAirframeRepair1P"`
	CostEngineRepair1P                 float64               `json:"CostEngineRepair1P"`
	CostHourlyAirframeRepair           float64               `json:"CostHourlyAirframeRepair"`
	CostHourlyEnginesRepair            float64               `json:"CostHourlyEnginesRepair"`
	Cost100HInspection                 float64               `json:"Cost100HInspection"`
	CostAnnualCheckup                  float64               `json:"CostAnnualCheckup"`
	CostHourly100HInspection           float64               `json:"CostHourly100HInspection"`
	CostHourlyAnnualCheckup            float64               `json:"CostHourlyAnnualCheckup"`
	CostHourlyTotalMaintenance         float64               `json:"CostHourlyTotalMaintenance"`
	CostHourlyRentalFee                float64               `json:"CostHourlyRentalFee"`
	CostHourlyRentTotal                float64               `json:"CostHourlyRentTotal"`
	CostWeeklyLease                    float64               `json:"CostWeeklyLease"`
	CostHourlyLease                    float64               `json:"CostHourlyLease"`
	ProfitHourlyLease                  float64               `json:"ProfitHourlyLease"`
	CostWeeklyOwning                   float64               `json:"CostWeeklyOwning"`
	CostHourlyOwning                   float64               `json:"CostHourlyOwning"`
	ProfitHourlyOwning                 float64               `json:"ProfitHourlyOwning"`
}

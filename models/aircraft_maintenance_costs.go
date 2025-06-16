package models

type AircraftMaintenanceCosts struct {
	AnnualCheckupCost          float64 `json:"AnnualCheckupCost"`
	Inspection100HCost         float64 `json:"Inspection100hCost"`
	RepairAirframe1PercentCost float64 `json:"RepairAirframe1PercentCost"`
	RepairEngine1PercentCost   float64 `json:"RepairEngine1PercentCost"`
	EngineReplacementCost      float64 `json:"EngineReplacementCost"`
}

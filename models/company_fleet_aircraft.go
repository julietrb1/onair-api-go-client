package models

type CompanyFleetAircraft struct {
	ID                             string                `json:"Id"`
	AircraftTypeID                 string                `json:"AircraftTypeId"`
	AircraftType                   AircraftTypeAtAirport `json:"AircraftType"`
	Nickname                       string                `json:"Nickname"`
	WorldID                        string                `json:"WorldId"`
	CompanyID                      string                `json:"CompanyId,omitempty"`
	CurrentAirportID               string                `json:"CurrentAirportId,omitempty"`
	CurrentAirport                 Airport               `json:"CurrentAirport,omitempty"`
	AircraftStatus                 int                   `json:"AircraftStatus"`
	LastStatusChange               string                `json:"LastStatusChange"`
	CurrentStatusDurationInMinutes int                   `json:"CurrentStatusDurationInMinutes"`
	AllowSell                      bool                  `json:"AllowSell"`
	AllowRent                      bool                  `json:"AllowRent"`
	AllowLease                     bool                  `json:"AllowLease"`
	SellPrice                      float64               `json:"SellPrice"`
	RentHourPrice                  float64               `json:"RentHourPrice"`
	Identifier                     string                `json:"Identifier"`
	Heading                        float64               `json:"Heading"`
	Longitude                      float64               `json:"Longitude"`
	Latitude                       float64               `json:"Latitude"`
	FuelTotalGallons               float64               `json:"fuelTotalGallons"`
	FuelWeight                     int                   `json:"fuelWeight"`
	Altitude                       float64               `json:"Altitude"`
	IndicatedSpeed                 float64               `json:"IndicatedSpeed"`
	GroundSpeed                    float64               `json:"GroundSpeed"`
	LoadedWeight                   float64               `json:"loadedWeight"`
	ZeroFuelWeight                 int                   `json:"zeroFuelWeight"`
	AirframeHours                  float64               `json:"airframeHours"`
	AirframeCondition              float64               `json:"airframeCondition"`
	AirframeMaxCondition           float64               `json:"AirframeMaxCondition"`
	LastAnnualCheckup              string                `json:"LastAnnualCheckup"`
	Last100HInspection             string                `json:"Last100hInspection"`
	LastWeeklyOwnershipPayment     string                `json:"LastWeeklyOwnershipPayment"`
	LastParkingFeePayment          string                `json:"LastParkingFeePayment"`
	IsControlledByAI               bool                  `json:"IsControlledByAI"`
	HoursBefore100HInspection      float64               `json:"HoursBefore100HInspection"`
	Engines                        []Engines             `json:"Engines"`
	ConfigFirstSeats               int                   `json:"ConfigFirstSeats"`
	ConfigBusSeats                 int                   `json:"ConfigBusSeats"`
	ConfigEcoSeats                 int                   `json:"ConfigEcoSeats"`
	SeatsReservedForEmployees      int                   `json:"SeatsReservedForEmployees"`
	LastMagicTransportationDate    string                `json:"LastMagicTransportationDate"`
	CurrentCompanyID               string                `json:"CurrentCompanyId"`
	CurrentCompanyIDIfAny          string                `json:"CurrentCompanyIdIfAny"`
	ExtraWeightCapacity            float64               `json:"ExtraWeightCapacity"`
	TotalWeightCapacity            float64               `json:"TotalWeightCapacity"`
	CurrentSeats                   int                   `json:"CurrentSeats"`
	MustDoMaintenance              bool                  `json:"MustDoMaintenance"`
	MustDoMaintenanceSoon          bool                  `json:"MustDoMaintenanceSoon"`
	RentAirportID                  string                `json:"RentAirportId,omitempty"`
	RentAirport                    Airport               `json:"RentAirport,omitempty"`
	RentFuelTotalGallons           float64               `json:"RentFuelTotalGallons,omitempty"`
	RentCautionAmount              float64               `json:"RentCautionAmount,omitempty"`
	RentCompanyID                  string                `json:"RentCompanyId,omitempty"`
	RentCompany                    VA                    `json:"RentCompany,omitempty"`
	RentStartDate                  string                `json:"RentStartDate,omitempty"`
	RentLastDailyChargeDate        string                `json:"RentLastDailyChargeDate,omitempty"`
	InFlightStatus                 int                   `json:"InFlightStatus,omitempty"`
}

type Engines struct {
	ID           string  `json:"Id"`
	AircraftID   string  `json:"AircraftId"`
	Number       int     `json:"Number"`
	Condition    float64 `json:"Condition"`
	MaxCondition float64 `json:"MaxCondition"`
	EngineHours  float64 `json:"EngineHours"`
	LastCheckup  string  `json:"LastCheckup"`
}

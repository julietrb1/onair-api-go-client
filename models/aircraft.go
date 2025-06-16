package models

import (
	"github.com/google/uuid"
	"time"
)

type AircraftClass struct {
	ID        uuid.UUID `json:"Id"`
	ShortName string    `json:"ShortName"`
	Name      string    `json:"Name"`
	Order     int       `json:"Order"`
}

type AircraftAddon struct {
	ID                            uuid.UUID `json:"Id"`
	Hash                          string    `json:"Hash"`
	AircraftTypeID                uuid.UUID `json:"AircraftTypeId"`
	CreationDate                  time.Time `json:"CreationDate"`
	LastModerationDate            time.Time `json:"LastModerationDate,omitempty"`
	FuelTotalCapacityInGallons    float64   `json:"FuelTotalCapacityInGallons"`
	DisplayName                   string    `json:"DisplayName"`
	TypeName                      string    `json:"TypeName"`
	AirFileName                   string    `json:"AirFileName"`
	SimulatorVersion              int       `json:"simulatorVersion"`
	EmptyWeight                   int       `json:"emptyWeight"`
	MaximumGrossWeight            int       `json:"maximumGrossWeight"`
	EstimatedCruiseFF             int       `json:"estimatedCruiseFF"`
	EngineType                    int       `json:"engineType"`
	NumberOfEngines               int       `json:"numberOfEngines"`
	FuelType                      int       `json:"fuelType"`
	DesignSpeedVS0                float64   `json:"designSpeedVS0"`
	DesignSpeedVS1                float64   `json:"designSpeedVS1"`
	DesignSpeedVC                 float64   `json:"designSpeedVC"`
	IsDisabled                    bool      `json:"IsDisabled"`
	AircraftDataSheetUrl          string    `json:"AircraftDataSheetUrl,omitempty"`
	AddonUrl                      string    `json:"AddonUrl,omitempty"`
	IsVanilla                     bool      `json:"IsVanilla"`
	CreatedByUserID               uuid.UUID `json:"CreatedByUserId"`
	TestedByUser                  bool      `json:"TestedByUser"`
	LastTestFlightDate            time.Time `json:"LastTestFlightDate,omitempty"`
	ConsolidatedDesignSpeedVC     float64   `json:"ConsolidatedDesignSpeedVC"`
	ConsolidatedEstimatedCruiseFF float64   `json:"ConsolidatedEstimatedCruiseFF"`
	EnableAutoConsolidation       bool      `json:"EnableAutoConsolidation"`
	ComputedMaxPayload            int       `json:"ComputedMaxPayload"`
	ComputedSeats                 int       `json:"ComputedSeats"`
}

type AircraftType struct {
	ID                         uuid.UUID     `json:"Id"`
	Hash                       string        `json:"Hash"`
	AircraftClassID            uuid.UUID     `json:"AircraftClassId"`
	AircraftClass              AircraftClass `json:"AircraftClass"`
	CreationDate               time.Time     `json:"CreationDate"`
	LastModerationDate         time.Time     `json:"LastModerationDate"`
	DisplayName                string        `json:"DisplayName"`
	TypeName                   string        `json:"TypeName"`
	FlightsCount               int           `json:"FlightsCount"`
	TimeBetweenOverhaul        int           `json:"TimeBetweenOverhaul"`
	HightimeAirframe           int           `json:"HightimeAirframe"`
	AirportMinSize             int           `json:"AirportMinSize"`
	EmptyWeight                int           `json:"emptyWeight"`
	MaximumGrossWeight         int           `json:"maximumGrossWeight"`
	EstimatedCruiseFF          int           `json:"estimatedCruiseFF"`
	BasePrice                  float64       `json:"Baseprice"`
	FuelTotalCapacityInGallons float64       `json:"FuelTotalCapacityInGallons"`
	EngineType                 int           `json:"engineType"`
	NumberOfEngines            int           `json:"numberOfEngines"`
	Seats                      int           `json:"seats"`
	NeedsCopilot               bool          `json:"needsCopilot"`
	FuelType                   int           `json:"fuelType"`
	MaximumCargoWeight         int           `json:"maximumCargoWeight"`
	MaximumRangeInHour         float64       `json:"maximumRangeInHour"`
	MaximumRangeInNM           float64       `json:"maximumRangeInNM"`
}

type EmbeddedAircraft struct {
	ID                             string                `json:"Id"`
	AircraftTypeID                 string                `json:"AircraftTypeId"`
	AircraftType                   AircraftTypeAtAirport `json:"AircraftType"`
	Nickname                       string                `json:"Nickname"`
	WorldID                        string                `json:"WorldId"`
	AircraftStatus                 int                   `json:"AircraftStatus"`
	LastStatusChange               string                `json:"LastStatusChange"`
	CurrentStatusDurationInMinutes int                   `json:"CurrentStatusDurationInMinutes"`
	AllowSell                      bool                  `json:"AllowSell"`
	AllowRent                      bool                  `json:"AllowRent"`
	AllowLease                     bool                  `json:"AllowLease"`
	SellPrice                      float64               `json:"SellPrice"`
	RentHourPrice                  float64               `json:"RentHourPrice"`
	RentAirportID                  string                `json:"RentAirportId"`
	RentFuelTotalGallons           float64               `json:"RentFuelTotalGallons"`
	RentCautionAmount              float64               `json:"RentCautionAmount"`
	RentCompanyID                  string                `json:"RentCompanyId"`
	RentStartDate                  string                `json:"RentStartDate"`
	RentLastDailyChargeDate        string                `json:"RentLastDailyChargeDate"`
	Identifier                     string                `json:"Identifier"`
	Heading                        float64               `json:"Heading"`
	Longitude                      float64               `json:"Longitude"`
	Latitude                       float64               `json:"Latitude"`
	FuelTotalGallons               float64               `json:"fuelTotalGallons"`
	FuelWeight                     int                   `json:"fuelWeight"`
	Altitude                       float64               `json:"Altitude"`
	IndicatedSpeed                 float64               `json:"IndicatedSpeed"`
	GroundSpeed                    float64               `json:"GroundSpeed"`
	InFlightStatus                 int                   `json:"InFlightStatus"`
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
}

type EmbeddedAircraftAddon struct {
	ID                            string                `json:"Id"`
	Hash                          string                `json:"Hash"`
	AircraftTypeID                string                `json:"AircraftTypeId"`
	AircraftType                  AircraftTypeAtAirport `json:"AircraftType"`
	CreationDate                  string                `json:"CreationDate"`
	LastModerationDate            string                `json:"LastModerationDate"`
	FuelTotalCapacityInGallons    float64               `json:"FuelTotalCapacityInGallons"`
	DisplayName                   string                `json:"DisplayName"`
	TypeName                      string                `json:"TypeName"`
	AirFileName                   string                `json:"AirFileName"`
	SimulatorVersion              int                   `json:"simulatorVersion"`
	EmptyWeight                   int                   `json:"emptyWeight"`
	MaximumGrossWeight            int                   `json:"maximumGrossWeight"`
	EstimatedCruiseFF             int                   `json:"estimatedCruiseFF"`
	EngineType                    int                   `json:"engineType"`
	NumberOfEngines               int                   `json:"numberOfEngines"`
	FuelType                      int                   `json:"fuelType"`
	DesignSpeedVS0                float64               `json:"designSpeedVS0"`
	DesignSpeedVS1                float64               `json:"designSpeedVS1"`
	DesignSpeedVC                 float64               `json:"designSpeedVC"`
	IsDisabled                    bool                  `json:"IsDisabled"`
	AddonUrl                      string                `json:"AddonUrl"`
	IsVanilla                     bool                  `json:"IsVanilla"`
	CreatedByUserID               string                `json:"CreatedByUserId"`
	TestedByUser                  bool                  `json:"TestedByUser"`
	LastTestFlightDate            string                `json:"LastTestFlightDate"`
	ConsolidatedDesignSpeedVC     float64               `json:"ConsolidatedDesignSpeedVC"`
	ConsolidatedEstimatedCruiseFF float64               `json:"ConsolidatedEstimatedCruiseFF"`
	EnableAutoConsolidation       bool                  `json:"EnableAutoConsolidation"`
	ComputedMaxPayload            int                   `json:"ComputedMaxPayload"`
	ComputedSeats                 int                   `json:"ComputedSeats"`
	ProposedSeats                 int                   `json:"ProposedSeats"`
	ProposedMaxPayload            int                   `json:"ProposedMaxPayload"`
	FlightsCount                  int                   `json:"FlightsCount"`
	DisableHardLanding            bool                  `json:"DisableHardLanding"`
	DisableAutoSet                bool                  `json:"DisableAutoSet"`
	SimVersionShortDisplay        string                `json:"SimVersionShortDisplay"`
	SimVersionDisplay             string                `json:"SimVersionDisplay"`
}

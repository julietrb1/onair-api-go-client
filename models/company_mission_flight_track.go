package models

type CompanyMissionFlightTrack struct {
	Id        string `json:"Id"`
	MissionId string `json:"MissionId"`
	FlightId  string `json:"FlightId"`
}

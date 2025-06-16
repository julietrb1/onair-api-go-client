package api

import (
	"encoding/json"
	"fmt"
	"github.com/julietrb1/onair-api-go-client/models"
	"net/http"
	"os"
)

const onAirBaseURL = "https://server1.models.company/api"
const onAirAuthHeaderName = "oa-apikey"

// OnAirAPI encapsulates the OnAir API client
type OnAirAPI struct {
	apiKey string
	client *http.Client
}

// NewOnAirAPI creates a new OnAir API client with the required API key
func NewOnAirAPI() (*OnAirAPI, error) {
	apiKey := os.Getenv("ONAIR_API_KEY")
	if apiKey == "" {
		return nil, fmt.Errorf("ONAIR_API_KEY environment variable not set")
	}

	client := &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}

	return &OnAirAPI{
		apiKey: apiKey,
		client: client,
	}, nil
}

// OAResponse is a reusable struct for handling OnAir responses.
type OAResponse[T any] struct {
	Content T `json:"Content"`
}

// GetAirport fetches an airport by its ICAO.
func (api *OnAirAPI) GetAirport(icao string) (*models.Airport, error) {
	url := fmt.Sprintf("%s/v1/airports/%s", onAirBaseURL, icao)
	resp, err := getResponse(url, api)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var apiResp OAResponse[models.Airport]
	if err := json.NewDecoder(resp.Body).Decode(&apiResp); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	if apiResp.Content.ICAO == "" {
		return nil, fmt.Errorf("airport with ICAO %s not found in the API", icao)
	}

	if apiResp.Content.Name == "" {
		return nil, fmt.Errorf("airport with ICAO %s has no name in the API", icao)
	}

	return &apiResp.Content, nil
}

// GetAircraftAtAirport fetches all aircraft at a specific airport by its ICAO.
func (api *OnAirAPI) GetAircraftAtAirport(icao string) (*[]models.AircraftTypeAtAirport, error) {
	url := fmt.Sprintf("%s/v1/airports/%s/aircraft", onAirBaseURL, icao)
	resp, err := getResponse(url, api)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var aircraftResponse OAResponse[[]models.AircraftTypeAtAirport]
	if err := json.NewDecoder(resp.Body).Decode(&aircraftResponse); err != nil {
		return nil, fmt.Errorf("error decoding response: %w", err)
	}

	return &aircraftResponse.Content, nil
}

// GetAircraftType fetches an aircraft type by its ID.
func (api *OnAirAPI) GetAircraftType(aircraftTypeID string) (*models.AircraftType, error) {
	url := fmt.Sprintf("%s/v1/aircrafttypes/%s", onAirBaseURL, aircraftTypeID)
	resp, err := getResponse(url, api)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var apiResp OAResponse[models.AircraftType]
	if err := json.NewDecoder(resp.Body).Decode(&apiResp); err != nil {
		return nil, fmt.Errorf("error decoding response: %w", err)
	}

	return &apiResp.Content, nil
}

// GetVA gets details for a given Virtual Airline.
func (api *OnAirAPI) GetVA(vaID string) (*models.VA, error) {
	url := fmt.Sprintf("%s/v1/va/%s", onAirBaseURL, vaID)

	resp, err := getResponse(url, api)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var apiResp OAResponse[models.VA]
	if err := json.NewDecoder(resp.Body).Decode(&apiResp); err != nil {
		return nil, fmt.Errorf("error decoding response: %w", err)
	}

	return &apiResp.Content, nil
}

// GetVAMembers gets details for members of a given Virtual Airline.
func (api *OnAirAPI) GetVAMembers(vaID string) (*[]models.VAMember, error) {
	url := fmt.Sprintf("%s/v1/va/%s/members", onAirBaseURL, vaID)

	resp, err := getResponse(url, api)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var apiResp OAResponse[[]models.VAMember]
	if err := json.NewDecoder(resp.Body).Decode(&apiResp); err != nil {
		return nil, fmt.Errorf("error decoding response: %w", err)
	}

	return &apiResp.Content, nil
}

// GetVAShareholders gets details for shareholders of a given Virtual Airline.
func (api *OnAirAPI) GetVAShareholders(vaID string) (*[]models.VAShareholder, error) {
	url := fmt.Sprintf("%s/v1/va/%s/shareholders", onAirBaseURL, vaID)

	resp, err := getResponse(url, api)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var apiResp OAResponse[[]models.VAShareholder]
	if err := json.NewDecoder(resp.Body).Decode(&apiResp); err != nil {
		return nil, fmt.Errorf("error decoding response: %w", err)
	}

	return &apiResp.Content, nil
}

// GetVARoles gets details for roles of a given Virtual Airline.
func (api *OnAirAPI) GetVARoles(vaID string) (*[]models.VARole, error) {
	url := fmt.Sprintf("%s/v1/va/%s/roles", onAirBaseURL, vaID)

	resp, err := getResponse(url, api)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var apiResp OAResponse[[]models.VARole]
	if err := json.NewDecoder(resp.Body).Decode(&apiResp); err != nil {
		return nil, fmt.Errorf("error decoding response: %w", err)
	}

	return &apiResp.Content, nil
}

// GetAircraftEconomicDetails gets financial and operational information for an aircraft.
func (api *OnAirAPI) GetAircraftEconomicDetails(aircraftID string) (*models.AircraftEconomicDetails, error) {
	url := fmt.Sprintf("%s/v1/aircraft/%s/economic_details", onAirBaseURL, aircraftID)

	resp, err := getResponse(url, api)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var apiResp OAResponse[models.AircraftEconomicDetails]
	if err := json.NewDecoder(resp.Body).Decode(&apiResp); err != nil {
		return nil, fmt.Errorf("error decoding response: %w", err)
	}

	return &apiResp.Content, nil
}

// GetAircraftFlights gets all flights that an aircraft has performed.
func (api *OnAirAPI) GetAircraftFlights(aircraftID string, startIndex int, limit int) (*models.AircraftEconomicDetails, error) {
	url := fmt.Sprintf("%s/v1/aircraft/%s/flights?startIndex=%d&limit=%d", onAirBaseURL, aircraftID, startIndex, limit)

	resp, err := getResponse(url, api)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var apiResp OAResponse[models.AircraftEconomicDetails]
	if err := json.NewDecoder(resp.Body).Decode(&apiResp); err != nil {
		return nil, fmt.Errorf("error decoding response: %w", err)
	}

	return &apiResp.Content, nil
}

// GetAircraftMaintenanceCosts gets repair, engine replacement, inspection, and checkup costs for an aircraft.
func (api *OnAirAPI) GetAircraftMaintenanceCosts(aircraftID string) (*models.AircraftMaintenanceCosts, error) {
	url := fmt.Sprintf("%s/v1/aircraft/%s/maintenance_costs", onAirBaseURL, aircraftID)

	resp, err := getResponse(url, api)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var apiResp OAResponse[models.AircraftMaintenanceCosts]
	if err := json.NewDecoder(resp.Body).Decode(&apiResp); err != nil {
		return nil, fmt.Errorf("error decoding response: %w", err)
	}

	return &apiResp.Content, nil
}

// GetCompany gets a company.
func (api *OnAirAPI) GetCompany(companyID string) (*models.Company, error) {
	url := fmt.Sprintf("%s/v1/company/%s", onAirBaseURL, companyID)

	resp, err := getResponse(url, api)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var apiResp OAResponse[models.Company]
	if err := json.NewDecoder(resp.Body).Decode(&apiResp); err != nil {
		return nil, fmt.Errorf("error decoding response: %w", err)
	}

	return &apiResp.Content, nil
}

// GetCompanyBalanceSheet gets a company's accounting information relating to assets and liabilities.
func (api *OnAirAPI) GetCompanyBalanceSheet(companyID string) (*models.CompanyBalanceSheet, error) {
	url := fmt.Sprintf("%s/v1/company/%s/balancesheet", onAirBaseURL, companyID)

	resp, err := getResponse(url, api)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var apiResp OAResponse[models.CompanyBalanceSheet]
	if err := json.NewDecoder(resp.Body).Decode(&apiResp); err != nil {
		return nil, fmt.Errorf("error decoding response: %w", err)
	}

	return &apiResp.Content, nil
}

// GetCompanyCashFlow gets a company's transaction information including amounts, people involved, the relevant account, and more.
func (api *OnAirAPI) GetCompanyCashFlow(companyID string) (*models.CompanyCashFlow, error) {
	url := fmt.Sprintf("%s/v1/company/%s/cashflow", onAirBaseURL, companyID)

	resp, err := getResponse(url, api)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var apiResp OAResponse[models.CompanyCashFlow]
	if err := json.NewDecoder(resp.Body).Decode(&apiResp); err != nil {
		return nil, fmt.Errorf("error decoding response: %w", err)
	}

	return &apiResp.Content, nil
}

// GetCompanyDashboard gets a company's statistical summary including debt, aircraft, income, and XP.
func (api *OnAirAPI) GetCompanyDashboard(companyID string) (*models.CompanyDashboard, error) {
	url := fmt.Sprintf("%s/v1/company/%s/cashflow", onAirBaseURL, companyID)

	resp, err := getResponse(url, api)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var apiResp OAResponse[models.CompanyDashboard]
	if err := json.NewDecoder(resp.Body).Decode(&apiResp); err != nil {
		return nil, fmt.Errorf("error decoding response: %w", err)
	}

	return &apiResp.Content, nil
}

// GetCompanyEmployees gets a company's employees including salaries, home airport, and fatigue.
func (api *OnAirAPI) GetCompanyEmployees(companyID string) (*[]models.Employee, error) {
	url := fmt.Sprintf("%s/v1/company/%s/cashflow", onAirBaseURL, companyID)

	resp, err := getResponse(url, api)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var apiResp OAResponse[[]models.Employee]
	if err := json.NewDecoder(resp.Body).Decode(&apiResp); err != nil {
		return nil, fmt.Errorf("error decoding response: %w", err)
	}

	return &apiResp.Content, nil
}

// GetCompanyFBOs gets a company's FBOs including fuel quantity, fuel orders, and workshops.
func (api *OnAirAPI) GetCompanyFBOs(companyID string) (*[]models.FBO, error) {
	url := fmt.Sprintf("%s/v1/company/%s/fbos", onAirBaseURL, companyID)

	resp, err := getResponse(url, api)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var apiResp OAResponse[[]models.FBO]
	if err := json.NewDecoder(resp.Body).Decode(&apiResp); err != nil {
		return nil, fmt.Errorf("error decoding response: %w", err)
	}

	return &apiResp.Content, nil
}

// GetCompanyFleet gets aircraft rented, leased, or owned by a company.
func (api *OnAirAPI) GetCompanyFleet(companyID string) (*[]models.CompanyFleetAircraft, error) {
	url := fmt.Sprintf("%s/v1/company/%s/fbos", onAirBaseURL, companyID)

	resp, err := getResponse(url, api)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var apiResp OAResponse[[]models.CompanyFleetAircraft]
	if err := json.NewDecoder(resp.Body).Decode(&apiResp); err != nil {
		return nil, fmt.Errorf("error decoding response: %w", err)
	}

	return &apiResp.Content, nil
}

// GetCompanyFlights gets flights registered by a company, including airborne time, max. bank, and XP earned.
func (api *OnAirAPI) GetCompanyFlights(companyID string) (*[]models.Flight, error) {
	url := fmt.Sprintf("%s/v1/company/%s/flights", onAirBaseURL, companyID)

	resp, err := getResponse(url, api)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var apiResp OAResponse[[]models.Flight]
	if err := json.NewDecoder(resp.Body).Decode(&apiResp); err != nil {
		return nil, fmt.Errorf("error decoding response: %w", err)
	}

	return &apiResp.Content, nil
}

// GetCompanyIncomeStatement gets revenue and expenditure information for a company.
func (api *OnAirAPI) GetCompanyIncomeStatement(companyID string) (*models.CompanyIncomeStatement, error) {
	url := fmt.Sprintf("%s/v1/company/%s/incomestatement", onAirBaseURL, companyID)

	resp, err := getResponse(url, api)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var apiResp OAResponse[models.CompanyIncomeStatement]
	if err := json.NewDecoder(resp.Body).Decode(&apiResp); err != nil {
		return nil, fmt.Errorf("error decoding response: %w", err)
	}

	return &apiResp.Content, nil
}

// GetCompanyCompletedJobs gets completed job information for a company.
func (api *OnAirAPI) GetCompanyCompletedJobs(companyID string) (*[]models.CompanyJob, error) {
	url := fmt.Sprintf("%s/v1/company/%s/jobs/completed", onAirBaseURL, companyID)

	resp, err := getResponse(url, api)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var apiResp OAResponse[[]models.CompanyJob]
	if err := json.NewDecoder(resp.Body).Decode(&apiResp); err != nil {
		return nil, fmt.Errorf("error decoding response: %w", err)
	}

	return &apiResp.Content, nil
}

func getResponse(url string, api *OnAirAPI) (*http.Response, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set(onAirAuthHeaderName, api.apiKey)

	resp, err := api.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to send request: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}
	// Note: The caller is responsible for closing the response body
	return resp, nil
}

# OnAir API Go Client

An unofficial Go client library for interacting with the [OnAir](https://www.onair.company) API. OnAir is a proprietary commercial flight simulation and management platform that provides a realistic airline management experience.

## Installation

```bash
go get github.com/julietrb1/onair-api-go-client
```

## Usage

```go
package main

import (
	"fmt"
	"github.com/julietrb1/onair-api-go-client/api"
	"log"
)

func main() {
	client, err := api.NewOnAirAPI()
	if err != nil {
		log.Fatalf("Failed to create OnAir API client: %v", err)
	}

	airport, err := client.GetAirport("KLAX")
	if err != nil {
		log.Fatalf("Failed to get airport: %v", err)
	}

	fmt.Printf("Airport: %s (%s)\n", airport.Name, airport.ICAO)
}
```

## Available Methods

The client provides the following methods to access various OnAir resources:

### Airport and Aircraft Methods
- `GetAirport(icao string)` - Fetches an airport by its ICAO
- `GetAircraftAtAirport(icao string)` - Fetches all aircraft at a specific airport
- `GetAircraftType(aircraftTypeID string)` - Fetches an aircraft type by its ID
- `GetAircraftEconomicDetails(aircraftID string)` - Gets financial and operational information for an aircraft
- `GetAircraftFlights(aircraftID string, startIndex int, limit int)` - Gets all flights that an aircraft has performed
- `GetAircraftMaintenanceCosts(aircraftID string)` - Gets repair, engine replacement, inspection, and checkup costs for an aircraft

### Virtual Airline Methods
- `GetVA(vaID string)` - Gets details for a given Virtual Airline
- `GetVAMembers(vaID string)` - Gets details for members of a given Virtual Airline
- `GetVAShareholders(vaID string)` - Gets details for shareholders of a given Virtual Airline
- `GetVARoles(vaID string)` - Gets details for roles of a given Virtual Airline

### Company Methods
- `GetCompany(companyID string)` - Gets a company
- `GetCompanyBalanceSheet(companyID string)` - Gets a company's accounting information relating to assets and liabilities
- `GetCompanyCashFlow(companyID string)` - Gets a company's transaction information
- `GetCompanyDashboard(companyID string)` - Gets a company's statistical summary
- `GetCompanyEmployees(companyID string)` - Gets a company's employees
- `GetCompanyFBOs(companyID string)` - Gets a company's FBOs
- `GetCompanyFleet(companyID string)` - Gets aircraft rented, leased, or owned by a company
- `GetCompanyFlights(companyID string)` - Gets flights registered by a company
- `GetCompanyIncomeStatement(companyID string)` - Gets revenue and expenditure information for a company
- `GetCompanyCompletedJobs(companyID string)` - Gets completed job information for a company
- `GetCompanyPendingJobs(companyID string)` - Gets pending job information for a company
- `GetCompanyMissionFlightTracks(companyID string)` - Gets pairs of missions and flights for a company
- `GetCompanyNotifications(companyID string)` - Gets notifications posted for a company

### Employee Methods
- `GetEmployee(employeeID string)` - Gets an employee's information

### FBO Methods
- `GetFBOJobs(fboID string)` - Gets all jobs available at an FBO

### Flight Methods
- `GetFlight(flightID string)` - Gets a flight including airborne time, max. bank, and XP earned

## Methods Yet To Be Implemented

The following API endpoints are not yet implemented with no expected dates:

- `/company/<id>/workorders`
- `/company/<id>/workorders/<aircraftid>`

## Authentication

You need an OnAir API key to use this client. Provide it when calling `NewOnAirAPI`.

## License

See [LICENSE](LICENSE). I do not represent OnAir - this contribution is a work of passion only.

Please note that while the code and documentation of this project are licensed on their own, the concepts and the data pulled from [OnAir](https://www.onair.company) are proprietary and subject to [OnAir's terms and conditions](https://www.onair.company/terms-conditions/).
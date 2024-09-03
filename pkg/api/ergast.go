package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

const baseURL = "https://api.jolpi.ca/ergast/f1"

var currentYear = time.Now().Year()

func GetLatestRaceResult() (*RaceTable, error) {
	result := RaceResultResponse{}
	err := apiCall(fmt.Sprintf("/%d/last/results", currentYear), &result)
	if err != nil {
		return nil, err
	}
	return &result.MRData.RaceTable, nil
}

func GetCurrentDriverStandings() (*DriverStandingsTable, error) {
	result := DriverStandingsResponse{}
	err := apiCall(fmt.Sprintf("/%d/driverstandings", currentYear), &result)
	if err != nil {
		return nil, err
	}
	return &result.MRData.StandingsTable, nil
}

func GetCurrentConstructorStandings() (*ConstructorStandingsTable, error) {
	result := ConstructorStandingsResponse{}
	err := apiCall(fmt.Sprintf("/%d/constructorstandings", currentYear), &result)
	if err != nil {
		return nil, err
	}
	return &result.MRData.StandingsTable, nil
}

func GetCurrentSeasonSchedule() (*ScheduleTable, error) {
	result := ScheduleResponse{}
	err := apiCall("/current", &result)
	if err != nil {
		return nil, err
	}
	return &result.MRData.RaceTable, nil
}

func GetRaceResult(year, round string) (*RaceTable, error) {
	result := RaceResultResponse{}
	err := apiCall(fmt.Sprintf("/%s/%s/results", year, round), &result)
	if err != nil {
		return nil, err
	}
	return &result.MRData.RaceTable, nil
}

func GetQualifyingResult(year, round string) (*QualifyingTable, error) {
	result := QualifyingResponse{}
	err := apiCall(fmt.Sprintf("/%s/%s/qualifying", year, round), &result)
	if err != nil {
		return nil, err
	}
	return &result.MRData.RaceTable, nil
}

func GetDriverRaceResults(year, driverID string) (*RaceTable, error) {
	result := RaceResultResponse{}
	err := apiCall(fmt.Sprintf("/%s/drivers/%s/results", year, driverID), &result)
	if err != nil {
		return nil, err
	}
	return &result.MRData.RaceTable, nil
}

func apiCall(url string, v interface{}) error {
	res, err := http.Get(baseURL + url)

	if err != nil {
		return fmt.Errorf("Error while contacting APIs:\n%v", err)
	}
	defer res.Body.Close()

	err = json.NewDecoder(res.Body).Decode(v)
	if err != nil {
		return fmt.Errorf("Error while reading API response:\n%v", err)
	}
	return nil
}

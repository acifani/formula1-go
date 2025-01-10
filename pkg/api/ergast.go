package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

const baseURL = "https://api.jolpi.ca/ergast/f1"

// getCurrentSeason returns the active F1 season year
func getCurrentSeason() int {
	now := time.Now()
	year := now.Year()

	// F1 season typically starts in March
	if now.Month() < time.March {
		return year - 1
	}
	return year
}

func GetLatestRaceResult() (*RaceTable, error) {
	result := RaceResultResponse{}
	err := apiCall(fmt.Sprintf("/%d/last/results", getCurrentSeason()), &result)
	if err != nil {
		return nil, err
	}
	return &result.MRData.RaceTable, nil
}

func GetCurrentDriverStandings() (*DriverStandingsTable, error) {
	result := DriverStandingsResponse{}
	err := apiCall(fmt.Sprintf("/%d/driverstandings", getCurrentSeason()), &result)
	if err != nil {
		return nil, err
	}
	return &result.MRData.StandingsTable, nil
}

func GetCurrentConstructorStandings() (*ConstructorStandingsTable, error) {
	result := ConstructorStandingsResponse{}
	err := apiCall(fmt.Sprintf("/%d/constructorstandings", getCurrentSeason()), &result)
	if err != nil {
		return nil, err
	}
	return &result.MRData.StandingsTable, nil
}

func GetCurrentSeasonSchedule() (*ScheduleTable, error) {
	result := ScheduleResponse{}
	err := apiCall(fmt.Sprintf("/%d", getCurrentSeason()), &result)
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

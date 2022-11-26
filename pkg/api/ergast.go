package api

import (
	"encoding/json"
	"net/http"
)

const baseURL = "http://ergast.com/api/f1"

func GetLatestRaceResult() (*RaceTable, error) {
	result := RaceResultResponse{}
	err := apiCall("/current/last/results.json", &result)
	if err != nil {
		return nil, err
	}
	return &result.MRData.RaceTable, nil
}

func GetCurrentDriverStandings() (*DriverStandingsTable, error) {
	result := DriverStandingsResponse{}
	err := apiCall("/current/driverStandings.json", &result)
	if err != nil {
		return nil, err
	}
	return &result.MRData.StandingsTable, nil
}

func GetCurrentConstructorStandings() (*ConstructorStandingsTable, error) {
	result := ConstructorStandingsResponse{}
	err := apiCall("/current/constructorStandings.json", &result)
	if err != nil {
		return nil, err
	}
	return &result.MRData.StandingsTable, nil
}

func apiCall(url string, v interface{}) error {
	res, err := http.Get(baseURL + url)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	err = json.NewDecoder(res.Body).Decode(v)
	if err != nil {
		return err
	}
	return nil
}

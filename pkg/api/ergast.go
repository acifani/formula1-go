package api

import (
	"encoding/json"
	"net/http"
)

func GetLatestRaceResult() (*RaceTable, error) {
	res, err := http.Get("https://ergast.com/api/f1/current/last/results.json")
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	result := RaceResultResponse{}
	err = json.NewDecoder(res.Body).Decode(&result)
	if err != nil {
		return nil, err
	}
	return &result.MRData.RaceTable, nil
}

func GetDriverInfo(id string) (*Driver, error) {
	res, err := http.Get("https://ergast.com/api/f1/drivers/" + id + ".json")
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	result := DriverInfo{}
	err = json.NewDecoder(res.Body).Decode(&result)
	if err != nil {
		return nil, err
	}
	return &result.MRData.DriverTable.Drivers[0], nil
}

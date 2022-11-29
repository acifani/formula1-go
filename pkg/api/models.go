package api

type RaceResultResponse struct {
	MRData struct {
		Xmlns     string    `json:"xmlns"`
		Series    string    `json:"series"`
		URL       string    `json:"url"`
		Limit     string    `json:"limit"`
		Offset    string    `json:"offset"`
		Total     string    `json:"total"`
		RaceTable RaceTable `json:"RaceTable"`
	} `json:"MRData"`
}

type RaceTable struct {
	Season string `json:"season"`
	Round  string `json:"round"`
	Races  []struct {
		Season   string `json:"season"`
		Round    string `json:"round"`
		URL      string `json:"url"`
		RaceName string `json:"raceName"`
		Circuit  struct {
			CircuitID   string `json:"circuitId"`
			URL         string `json:"url"`
			CircuitName string `json:"circuitName"`
			Location    struct {
				Lat      string `json:"lat"`
				Long     string `json:"long"`
				Locality string `json:"locality"`
				Country  string `json:"country"`
			} `json:"Location"`
		} `json:"Circuit"`
		Date    string `json:"date"`
		Time    string `json:"time"`
		Results []struct {
			Number       string `json:"number"`
			Position     string `json:"position"`
			PositionText string `json:"positionText"`
			Points       string `json:"points"`
			Driver       struct {
				DriverID        string `json:"driverId"`
				PermanentNumber string `json:"permanentNumber"`
				Code            string `json:"code"`
				URL             string `json:"url"`
				GivenName       string `json:"givenName"`
				FamilyName      string `json:"familyName"`
				DateOfBirth     string `json:"dateOfBirth"`
				Nationality     string `json:"nationality"`
			} `json:"Driver"`
			Constructor struct {
				ConstructorID string `json:"constructorId"`
				URL           string `json:"url"`
				Name          string `json:"name"`
				Nationality   string `json:"nationality"`
			} `json:"Constructor"`
			Grid   string `json:"grid"`
			Laps   string `json:"laps"`
			Status string `json:"status"`
			Time   struct {
				Millis string `json:"millis"`
				Time   string `json:"time"`
			} `json:"Time,omitempty"`
			FastestLap struct {
				Rank string `json:"rank"`
				Lap  string `json:"lap"`
				Time struct {
					Time string `json:"time"`
				} `json:"Time"`
				AverageSpeed struct {
					Units string `json:"units"`
					Speed string `json:"speed"`
				} `json:"AverageSpeed"`
			} `json:"FastestLap"`
		} `json:"Results"`
	} `json:"Races"`
}

type DriverStandingsResponse struct {
	MRData struct {
		Xmlns          string               `json:"xmlns"`
		Series         string               `json:"series"`
		URL            string               `json:"url"`
		Limit          string               `json:"limit"`
		Offset         string               `json:"offset"`
		Total          string               `json:"total"`
		StandingsTable DriverStandingsTable `json:"StandingsTable"`
	} `json:"MRData"`
}

type DriverStandingsTable struct {
	Season         string `json:"season"`
	StandingsLists []struct {
		Season          string `json:"season"`
		Round           string `json:"round"`
		DriverStandings []struct {
			Position     string `json:"position"`
			PositionText string `json:"positionText"`
			Points       string `json:"points"`
			Wins         string `json:"wins"`
			Driver       struct {
				DriverID        string `json:"driverId"`
				PermanentNumber string `json:"permanentNumber"`
				Code            string `json:"code"`
				URL             string `json:"url"`
				GivenName       string `json:"givenName"`
				FamilyName      string `json:"familyName"`
				DateOfBirth     string `json:"dateOfBirth"`
				Nationality     string `json:"nationality"`
			} `json:"Driver"`
			Constructors []struct {
				ConstructorID string `json:"constructorId"`
				URL           string `json:"url"`
				Name          string `json:"name"`
				Nationality   string `json:"nationality"`
			} `json:"Constructors"`
		} `json:"DriverStandings"`
	} `json:"StandingsLists"`
}

type ConstructorStandingsResponse struct {
	MRData struct {
		Xmlns          string                    `json:"xmlns"`
		Series         string                    `json:"series"`
		URL            string                    `json:"url"`
		Limit          string                    `json:"limit"`
		Offset         string                    `json:"offset"`
		Total          string                    `json:"total"`
		StandingsTable ConstructorStandingsTable `json:"StandingsTable"`
	} `json:"MRData"`
}

type ConstructorStandingsTable struct {
	Season         string `json:"season"`
	StandingsLists []struct {
		Season               string `json:"season"`
		Round                string `json:"round"`
		ConstructorStandings []struct {
			Position     string `json:"position"`
			PositionText string `json:"positionText"`
			Points       string `json:"points"`
			Wins         string `json:"wins"`
			Constructor  struct {
				ConstructorID string `json:"constructorId"`
				URL           string `json:"url"`
				Name          string `json:"name"`
				Nationality   string `json:"nationality"`
			} `json:"Constructor"`
		} `json:"ConstructorStandings"`
	} `json:"StandingsLists"`
}

type ScheduleResponse struct {
	MRData struct {
		Xmlns     string        `json:"xmlns"`
		Series    string        `json:"series"`
		URL       string        `json:"url"`
		Limit     string        `json:"limit"`
		Offset    string        `json:"offset"`
		Total     string        `json:"total"`
		RaceTable ScheduleTable `json:"RaceTable"`
	} `json:"MRData"`
}

type ScheduleTable struct {
	Season string `json:"season"`
	Races  []struct {
		Season   string `json:"season"`
		Round    string `json:"round"`
		URL      string `json:"url"`
		RaceName string `json:"raceName"`
		Circuit  struct {
			CircuitID   string `json:"circuitId"`
			URL         string `json:"url"`
			CircuitName string `json:"circuitName"`
			Location    struct {
				Lat      string `json:"lat"`
				Long     string `json:"long"`
				Locality string `json:"locality"`
				Country  string `json:"country"`
			} `json:"Location"`
		} `json:"Circuit"`
		Date          string `json:"date"`
		Time          string `json:"time"`
		FirstPractice struct {
			Date string `json:"date"`
			Time string `json:"time"`
		} `json:"FirstPractice"`
		SecondPractice struct {
			Date string `json:"date"`
			Time string `json:"time"`
		} `json:"SecondPractice"`
		ThirdPractice struct {
			Date string `json:"date"`
			Time string `json:"time"`
		} `json:"ThirdPractice,omitempty"`
		Qualifying struct {
			Date string `json:"date"`
			Time string `json:"time"`
		} `json:"Qualifying"`
		Sprint struct {
			Date string `json:"date"`
			Time string `json:"time"`
		} `json:"Sprint,omitempty"`
	} `json:"Races"`
}

type QualifyingResponse struct {
	MRData struct {
		Xmlns     string          `json:"xmlns"`
		Series    string          `json:"series"`
		URL       string          `json:"url"`
		Limit     string          `json:"limit"`
		Offset    string          `json:"offset"`
		Total     string          `json:"total"`
		RaceTable QualifyingTable `json:"RaceTable"`
	} `json:"MRData"`
}

type QualifyingTable struct {
	Season string `json:"season"`
	Round  string `json:"round"`
	Races  []struct {
		Season   string `json:"season"`
		Round    string `json:"round"`
		URL      string `json:"url"`
		RaceName string `json:"raceName"`
		Circuit  struct {
			CircuitID   string `json:"circuitId"`
			URL         string `json:"url"`
			CircuitName string `json:"circuitName"`
			Location    struct {
				Lat      string `json:"lat"`
				Long     string `json:"long"`
				Locality string `json:"locality"`
				Country  string `json:"country"`
			} `json:"Location"`
		} `json:"Circuit"`
		Date              string `json:"date"`
		Time              string `json:"time"`
		QualifyingResults []struct {
			Number   string `json:"number"`
			Position string `json:"position"`
			Driver   struct {
				DriverID        string `json:"driverId"`
				PermanentNumber string `json:"permanentNumber"`
				Code            string `json:"code"`
				URL             string `json:"url"`
				GivenName       string `json:"givenName"`
				FamilyName      string `json:"familyName"`
				DateOfBirth     string `json:"dateOfBirth"`
				Nationality     string `json:"nationality"`
			} `json:"Driver"`
			Constructor struct {
				ConstructorID string `json:"constructorId"`
				URL           string `json:"url"`
				Name          string `json:"name"`
				Nationality   string `json:"nationality"`
			} `json:"Constructor"`
			Q1 string `json:"Q1"`
			Q2 string `json:"Q2,omitempty"`
			Q3 string `json:"Q3,omitempty"`
		} `json:"QualifyingResults"`
	} `json:"Races"`
}

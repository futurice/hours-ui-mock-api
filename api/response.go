package api

import (
	"time"
)

type Meta struct {
	Limit      int    `json:"limit"`
	Next       string `json:"next"`
	Offset     int    `json:"offset"`
	Previous   string `json:"previous"`
	TotalCount int    `json:"total_count"`
}

type HoursResponse struct {
	Meta    Meta   `json:"meta"`
	Objects []Hour `json:"objects"`
}

type Hour struct {
	Absence     bool    `json:"absence"`
	Billable    bool    `json:"billable"`
	Day         string  `json:"day"`
	Description string  `json:"description"`
	Hours       float64 `json:"hours"`
}

func MockHoursResponse(startDate, endDate string) (HoursResponse, error) {
	timeFormat := "2006-01-02"
	start, err := time.Parse(timeFormat, startDate)
	if err != nil {
		return HoursResponse{}, err
	}
	end, err := time.Parse(timeFormat, endDate)
	if err != nil {
		return HoursResponse{}, err
	}

	duration := int(end.Sub(start).Hours()/24) + 1

	objects := make([]Hour, duration)

	for i := 0; i < duration; i++ {
		date := end.AddDate(0, 0, -i).Format(timeFormat)
		objects[i] = Hour{
			false,
			false,
			date,
			"Description",
			7.5,
		}
	}

	response := HoursResponse{
		Meta{300, "", 0, "", duration},
		objects,
	}

	return response, nil
}

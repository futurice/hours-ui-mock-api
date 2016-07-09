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

type UserResponse struct {
	FirstName       string    `json:"firstName"`
	LastName        string    `json:"lastName"`
	Balance         float64   `json:"balance"`
	HolidaysLeft    int       `json:"holidaysLeft"`
	UtilizationRate float64   `json:"utilizationRate"`
	Projects        []Project `json:"projects"`
}

type Project struct {
	ID int `json:"id"`
	Name string `json:"name"`
	Active bool `json:"active"`
	Tasks []Task `json:"tasks"`
}

type Task struct {
	ID int `json:"id"`
	Name string `json:"name"`
	LatestMarking string `json:"latestMarking"`
}

func MockUserResponse() UserResponse {
	balance := RoundToHalf(RandomFloat64(-10, 40));
	holidaysLeft := int(RandomFloat64(0, 24));
	utilizationRate := RandomFloat64(0, 100);

	projects := []Project{
		Project{
			1,
			"Interal work",
			true,
			[]Task{
				Task{
					11,
					"Things",
					"Doing things",
				},
				Task{
					12,
					"Stuff",
					"Doing stuff",
				},
			},
		},
		Project{
			2,
			"Actual customer work",
			true,
			[]Task{
				Task{
					13,
					"Development",
					"Developing",
				},
				Task{
					14,
					"On-Call",
					"Long weekend :(",
				},
			},
		},
		Project{
			3,
			"Not active project",
			false,
			[]Task{
				Task{
					15,
					"Work",
					"Doing work",
				},
			},
		},
	}

	return UserResponse{
		"Test",
		"User",
		balance,
		holidaysLeft,
		utilizationRate,
		projects,

	}
}

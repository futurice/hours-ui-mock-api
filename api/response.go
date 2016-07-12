package api

import (
	"time"
)

type ErrorResponse struct {
	Status     int    `json:"status"`
	StatusText string `json:"statusText"`
}

type HoursResponse struct {
	Projects []Project        `json:"projects"`
	Months   map[string]Month `json:"months"`
}

type Month struct {
	Hours           float64        `json:"hours"`
	UtilizationRate float64        `json:"utilizationRate"`
	Days            map[string]Day `json:"days"`
}

type Day struct {
	HolidayName     string  `json:"holidayName,omitempty"`
	Hours           float64 `json:"hours"`
	UtilizationRate float64 `json:"utilizationRate"`
	Entries         []Entry `json:"entries"`
}

type Entry struct {
	ID          int     `json:"id"`
	ProjectID   int     `json:"projectID"`
	TaskID      int     `json:"taskID"`
	Description string  `json:"description"`
	Hours       float64 `json:"hours"`
	Editable    bool    `json:"editable"`
}

type Project struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Active bool   `json:"active"`
	Tasks  []Task `json:"tasks"`
}

type Task struct {
	ID            int    `json:"id"`
	Name          string `json:"name"`
	LatestMarking string `json:"latestMarking"`
}

func MockHoursResponse(startDate, endDate string) (HoursResponse, error) {
	dateFormat := "2006-01-02"
	monthFormat := "2006-01"
	start, err := time.Parse(dateFormat, startDate)
	if err != nil {
		return HoursResponse{}, err
	}
	end, err := time.Parse(dateFormat, endDate)
	if err != nil {
		return HoursResponse{}, err
	}

	duration := int(end.Sub(start).Hours()/24) + 1

	months := make(map[string]Month)

	for i := 0; i < duration; i++ {
		day := start.AddDate(0, 0, i)

		month, ok := months[day.Format(monthFormat)]

		if ok == false {
			months[day.Format(monthFormat)] = Month{
				Hours:           RoundToHalf(RandomFloat64(0, 150)),
				UtilizationRate: RandomFloat64(0, 100),
				Days:            make(map[string]Day),
			}
			month = months[day.Format(monthFormat)]
		}
		month.Days[day.Format(dateFormat)] = days[i%len(days)]
	}

	response := HoursResponse{
		Projects: projects,
		Months:   months,
	}

	return response, nil
}

type UserResponse struct {
	FirstName       string  `json:"firstName"`
	LastName        string  `json:"lastName"`
	Balance         float64 `json:"balance"`
	HolidaysLeft    int     `json:"holidaysLeft"`
	UtilizationRate float64 `json:"utilizationRate"`
}

func MockUserResponse() UserResponse {
	return UserResponse{
		FirstName:       "Test",
		LastName:        "User",
		Balance:         RoundToHalf(RandomFloat64(-10, 40)),
		HolidaysLeft:    int(RandomFloat64(0, 24)),
		UtilizationRate: RandomFloat64(0, 100),
	}
}

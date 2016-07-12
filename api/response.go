package api

import (
	"strconv"
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

const DATE_FORMAT = "2006-01-02"
const MONTH_FORMAT = "2006-01"

func MockHoursResponse(startDate, endDate string) (HoursResponse, error) {
	start, err := time.Parse(DATE_FORMAT, startDate)
	if err != nil {
		return HoursResponse{}, err
	}
	end, err := time.Parse(DATE_FORMAT, endDate)
	if err != nil {
		return HoursResponse{}, err
	}

	duration := int(end.Sub(start).Hours()/24) + 1

	months := make(map[string]Month)

	for i := 0; i < duration; i++ {
		day := start.AddDate(0, 0, i)

		month, ok := months[day.Format(MONTH_FORMAT)]

		if ok == false {
			months[day.Format(MONTH_FORMAT)] = Month{
				Hours:           RoundToHalf(RandomFloat64(0, 150)),
				UtilizationRate: RandomFloat64(0, 100),
				Days:            make(map[string]Day),
			}
			month = months[day.Format(MONTH_FORMAT)]
		}
		month.Days[day.Format(DATE_FORMAT)] = days[i%len(days)]
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

type HoursUpdateRequest struct {
	ProjectID   int     `json:"projectID"`
	TaskID      int     `json:"taskID"`
	Description string  `json:"description"`
	Day         string  `json:"day"`
	Hours       float64 `json:"hours"`
}

type HoursUpdateResponse struct {
	User  UserResponse  `json:"user"`
	Hours HoursResponse `json:"hours"`
}

func MockHoursPOSTResponse(request HoursUpdateRequest) (HoursUpdateResponse, error) {
	ShuffleProjects(projects)
	day, err := time.Parse(DATE_FORMAT, request.Day)
	if err != nil {
		return HoursUpdateResponse{}, err
	}

	months := make(map[string]Month)

	months[day.Format(MONTH_FORMAT)] = Month{
		Hours:           RoundToHalf(RandomFloat64(0, 150)),
		UtilizationRate: RandomFloat64(0, 100),
		Days:            make(map[string]Day),
	}

	months[day.Format(MONTH_FORMAT)].Days[day.Format(DATE_FORMAT)] = Day{
		Hours:           request.Hours,
		UtilizationRate: 100.0,
		Entries: []Entry{
			Entry{
				ID:          int(RandomFloat64(0, 100)),
				ProjectID:   request.ProjectID,
				TaskID:      request.TaskID,
				Description: request.Description,
				Hours:       request.Hours,
				Editable:    true,
			},
		},
	}
	response := HoursUpdateResponse{
		User: UserResponse{
			FirstName:       "Test",
			LastName:        "User",
			Balance:         RoundToHalf(RandomFloat64(-10, 40)),
			HolidaysLeft:    int(RandomFloat64(0, 24)),
			UtilizationRate: RandomFloat64(0, 100),
		},
		Hours: HoursResponse{
			Projects: projects,
			Months:   months,
		},
	}

	return response, nil
}

func MockHoursPUTResponse(id string, request HoursUpdateRequest) (HoursUpdateResponse, error) {
	parse, err := strconv.ParseInt(id, 10, 32)
	if err != nil {
		return HoursUpdateResponse{}, nil
	}
	ID := int(parse)
	response, err := MockHoursPOSTResponse(request)
	if err != nil {
		return HoursUpdateResponse{}, nil
	}

	for _, month := range response.Hours.Months {
		for _, day := range month.Days {
			day.Entries[0].ID = ID
		}
	}
	return response, nil
}

func MockHoursDeleteResponse() (HoursUpdateResponse, error) {
	response, err := MockHoursPOSTResponse(HoursUpdateRequest{
		Day:         time.Now().Format(DATE_FORMAT),
		ProjectID:   1,
		TaskID:      1,
		Description: "test",
		Hours:       7.5,
	})
	if err != nil {
		return HoursUpdateResponse{}, nil
	}
	for _, month := range response.Hours.Months {
		for key, day := range month.Days {
			day.Entries = make([]Entry, 0, 0)
			month.Days[key] = day
		}
	}
	return response, nil
}

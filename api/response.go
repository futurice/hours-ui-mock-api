package api

import (
	"strconv"
	"time"
)

type ErrorResponse struct {
	Status int    `json:"status"`
	Error  string `json:"error"`
}

type HoursResponse struct {
	DefaultWorkHours   float64             `json:"defaultWorkHours"`
	MostRecentProjects []MostRecentProject `json:"mostRecentProjects"`
	Projects           []Project           `json:"projects"`
	Months             map[string]Month    `json:"months"`
}

type HoursUpdateResponse struct {
	Months             map[string]MonthUpdate `json:"months"`
	DefaultWorkHours   float64                `json:"defaultWorkHours"`
	MostRecentProjects []MostRecentProject    `json:"mostRecentProjects"`
}

type Month struct {
	Hours           float64        `json:"hours"`
	UtilizationRate float64        `json:"utilizationRate"`
	Days            map[string]Day `json:"days"`
}

type MonthUpdate struct {
	Hours           float64              `json:"hours"`
	UtilizationRate float64              `json:"utilizationRate"`
	Days            map[string]DayUpdate `json:"days"`
}

type Day struct {
	HolidayName string  `json:"holidayName,omitempty"`
	Hours       float64 `json:"hours"`
	Entries     []Entry `json:"entries"`
	Closed      bool    `json:"closed,omitempty"`
}

type DayUpdate struct {
	HolidayName string  `json:"holidayName,omitempty"`
	Hours       float64 `json:"hours"`
	Entry       *Entry  `json:"entry,omitempty"`
}

type Entry struct {
	ID          int     `json:"id"`
	ProjectID   int     `json:"projectID"`
	TaskID      int     `json:"taskID"`
	Description string  `json:"description"`
	Hours       float64 `json:"hours"`
	Closed      bool    `json:"closed,omitempty"`
}

// Every project that is assigned to the user, if closed, don't show in ui
type Project struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Tasks  []Task `json:"tasks"`
	Closed bool   `json:"closed,omitempty"`
}

type MostRecentProject struct {
	ID              int    `json:"id"`
	Name            string `json:"name"`
	MostRecentTasks []Task `json:"mostRecentTasks"`
}

// Every task that is assigned to the user, if closed, don't show in ui
type Task struct {
	ID            int    `json:"id"`
	Name          string `json:"name"`
	LatestMarking string `json:"latestMarking"`
	Absence       bool   `json:"absence,omitempty"`
	Closed        bool   `json:"closed,omitempty"`
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
		Projects:           projects,
		Months:             months,
		DefaultWorkHours:   7.5,
		MostRecentProjects: mostRecentProjects,
	}

	return response, nil
}

type UserResponse struct {
	FirstName       string  `json:"firstName"`
	LastName        string  `json:"lastName"`
	Balance         float64 `json:"balance"`
	HolidaysLeft    int     `json:"holidaysLeft"`
	UtilizationRate float64 `json:"utilizationRate"`
	ProfilePicture  string  `json:"profilePicture"`
}

func MockUserResponse() UserResponse {
	return UserResponse{
		FirstName:       "Test",
		LastName:        "User",
		Balance:         RoundToHalf(RandomFloat64(-10, 40)),
		HolidaysLeft:    int(RandomFloat64(0, 24)),
		UtilizationRate: RandomFloat64(0, 100),
		ProfilePicture:  "https://raw.githubusercontent.com/futurice/spiceprogram/gh-pages/assets/img/logo/chilicorn_no_text-128.png",
	}
}

type EntryUpdateRequest struct {
	ProjectID   int     `json:"projectID"`
	TaskID      int     `json:"taskID"`
	Description string  `json:"description"`
	Date        string  `json:"date"`
	Hours       float64 `json:"hours"`
	// When frontend sends closed entry to be updated, API doesn't do anything, just respond ok
	Closed bool `json:"closed,omitempty"`
}

type EntryUpdateResponse struct {
	User  UserResponse        `json:"user"`
	Hours HoursUpdateResponse `json:"hours"`
}

func MockEntryPOSTResponse(request EntryUpdateRequest) (EntryUpdateResponse, error) {
	date, err := time.Parse(DATE_FORMAT, request.Date)
	if err != nil {
		return EntryUpdateResponse{}, err
	}

	months := make(map[string]MonthUpdate)

	months[date.Format(MONTH_FORMAT)] = MonthUpdate{
		Hours:           RoundToHalf(RandomFloat64(0, 150)),
		UtilizationRate: RandomFloat64(0, 100),
		Days:            make(map[string]DayUpdate),
	}

	months[date.Format(MONTH_FORMAT)].Days[date.Format(DATE_FORMAT)] = DayUpdate{
		Hours: request.Hours,
		Entry: &Entry{
			ID:          int(RandomFloat64(0, 100)),
			ProjectID:   request.ProjectID,
			TaskID:      request.TaskID,
			Description: request.Description,
			Hours:       request.Hours,
			Closed:      request.Closed,
		},
	}

	mostRecent := make([]MostRecentProject, 1)

	for _, project := range projects {
		if project.ID == request.ProjectID {
			for _, task := range project.Tasks {
				if task.ID == request.TaskID {
					if !request.Closed {
						mostRecent[0] = MostRecentProject{
							ID:   project.ID,
							Name: project.Name,
							MostRecentTasks: []Task{
								task,
							},
						}
					} else {
						mostRecent[0] = mostRecentProjects[0]
					}
				}
			}
		}
	}

	response := EntryUpdateResponse{
		User: UserResponse{
			FirstName:       "Test",
			LastName:        "User",
			Balance:         RoundToHalf(RandomFloat64(-10, 40)),
			HolidaysLeft:    int(RandomFloat64(0, 24)),
			UtilizationRate: RandomFloat64(0, 100),
			ProfilePicture:  "https://raw.githubusercontent.com/futurice/spiceprogram/gh-pages/assets/img/logo/chilicorn_no_text-128.png",
		},
		Hours: HoursUpdateResponse{
			MostRecentProjects: mostRecent,
			Months:             months,
			DefaultWorkHours:   7.5,
		},
	}

	return response, nil
}

func MockEntryPUTResponse(id string, request EntryUpdateRequest) (EntryUpdateResponse, error) {
	parse, err := strconv.ParseInt(id, 10, 32)
	if err != nil {
		return EntryUpdateResponse{}, nil
	}
	ID := int(parse)
	response, err := MockEntryPOSTResponse(request)
	if err != nil {
		return EntryUpdateResponse{}, nil
	}

	for _, month := range response.Hours.Months {
		for _, day := range month.Days {
			day.Entry.ID = ID
		}
	}
	return response, nil
}

func MockEntryDELETEResponse() (EntryUpdateResponse, error) {
	response, err := MockEntryPOSTResponse(EntryUpdateRequest{
		Date:        time.Now().Format(DATE_FORMAT),
		ProjectID:   1,
		TaskID:      1,
		Description: "test",
		Hours:       7.5,
	})
	if err != nil {
		return EntryUpdateResponse{}, nil
	}
	for _, month := range response.Hours.Months {
		for key, day := range month.Days {
			day.Entry = nil
			day.Hours = 0
			month.Days[key] = day
		}
	}
	return response, nil
}

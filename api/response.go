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
	DefaultWorkHours float64          `json:"defaultWorkHours"`
	Projects         []Project        `json:"projects"`
	Months           map[string]Month `json:"months"`
}

type HoursUpdateResponse struct {
	DefaultWorkHours float64                `json:"defaultWorkHours"`
	Projects         []Project              `json:"projects"`
	Months           map[string]MonthUpdate `json:"months"`
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

// If the day is closed, one cannot mark entries to it. E.g. it should not be possible
// to mark hours to days that are already invoiced
type Day struct {
	HolidayName string  `json:"holidayName,omitempty"`
	Hours       float64 `json:"hours"`
	Entries     []Entry `json:"entries"`
	Closed      bool    `json:"closed,omitempty"`
}

// Only the created or updated entry is sent as an response. Deleted entry shows as null
type DayUpdate struct {
	HolidayName string  `json:"holidayName,omitempty"`
	Hours       float64 `json:"hours"`
	Entry       *Entry  `json:"entry,omitempty"`
}

// Entries are closed when they are invoiced and one cannot edit them afterwards
type Entry struct {
	ID          int     `json:"id"`
	ProjectID   int     `json:"projectID"`
	TaskID      int     `json:"taskID"`
	Description string  `json:"description"`
	Smiley      int     `json:"smiley,omitempty"`
	Hours       float64 `json:"hours"`
	Closed      bool    `json:"closed,omitempty"`
}

// Every project that is assigned to the user, sorted based on most recent usage.
// UI hides projects that are closed and were the latest entry was made over two months ago
type Project struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Tasks  []Task `json:"tasks"`
	Closed bool   `json:"closed,omitempty"`
}

// Every task that is assigned to the user, sorted based on most recent usage.
// UI hides tasks that are closed and were the latest entry was made over two months ago
type Task struct {
	ID             int         `json:"id"`
	Name           string      `json:"name"`
	Absence        bool        `json:"absence,omitempty"`
	Closed         bool        `json:"closed,omitempty"`
	LatestEntry    LatestEntry `json:"latestEntry"`
	HoursRemaining float64     `json:"hoursRemaining,omitempty"`
}

type LatestEntry struct {
	Date        string  `json:"date"`
	ID          int     `json:"id"`
	ProjectID   int     `json:"projectID"`
	TaskID      int     `json:"taskID"`
	Description string  `json:"description"`
	Smiley      int     `json:"smiley"`
	Hours       float64 `json:"hours"`
	Closed      bool    `json:"closed,omitempty"`
}

const DATE_FORMAT = "2006-01-02"
const MONTH_FORMAT = "2006-01"

func fillProjects() []Project {
	filledProjects := make([]Project, len(projects))
	for i, project := range projects {
		_project := project
		for j, task := range project.Tasks {
			_task := task
			_task.LatestEntry.Date = time.Now().Format(DATE_FORMAT)
			_task.LatestEntry.Hours = RoundToHalf(RandomFloat64(0.5, 7.5))
			if _project.ID != internalProject.ID && _project.ID != absenceProject.ID {
				_task.HoursRemaining = RoundToHalf(RandomFloat64(-10, 20))
			}
			_project.Tasks[j] = _task
		}
		filledProjects[i] = _project
	}
	return filledProjects
}

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
	filledProjects := fillProjects()

	response := HoursResponse{
		Projects:         filledProjects,
		Months:           months,
		DefaultWorkHours: 7.5,
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
	Smiley      int     `json:"smiley"`
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
	filledProjects := fillProjects()
	ShuffleProjects(filledProjects)
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
			Smiley:      request.Smiley,
			Hours:       request.Hours,
			Closed:      request.Closed,
		},
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
			DefaultWorkHours: 7.5,
			Projects:         filledProjects,
			Months:           months,
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
		Smiley:      1,
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

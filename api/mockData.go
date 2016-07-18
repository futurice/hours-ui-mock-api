package api

var internalProject = Project{
	ID:     1,
	Name:   "Internal work",
	Active: true,
	Tasks: []Task{
		Task{
			ID:            1,
			Name:          "Things",
			LatestMarking: "Doing things",
		},
		Task{
			ID:            2,
			Name:          "Stuff",
			LatestMarking: "Doing stuff",
		},
	},
}

var absenceProject = Project{
	ID:     4,
	Name:   "Absences",
	Active: true,
	Tasks: []Task{
		Task{
			ID:            6,
			Name:          "Balance leave",
			LatestMarking: "Balance leave",
		},
		Task{
			ID:            7,
			Name:          "Unpaid holiday",
			LatestMarking: "Unpaid holiday",
		},
		Task{
			ID:            8,
			Name:          "Sick leave",
			LatestMarking: "Sick leave",
		},
	},
}

var customerProject = Project{
	ID:     2,
	Name:   "Actual customer work",
	Active: true,
	Tasks: []Task{
		Task{
			ID:            3,
			Name:          "Development",
			LatestMarking: "Developing",
		},
		Task{
			ID:            4,
			Name:          "On-Call",
			LatestMarking: "Long weekend :(",
		},
	},
}

var inactiveProject = Project{
	ID:     3,
	Name:   "Not active project",
	Active: false,
	Tasks: []Task{
		Task{
			ID:            5,
			Name:          "Work",
			LatestMarking: "Doing work",
		},
	},
}

var projects = []Project{
	internalProject,
	absenceProject,
	customerProject,
	inactiveProject,
}

var days = []Day{
	Day{
		UtilizationRate: 0.0,
		Hours:           7.5,
		Entries: []Entry{
			Entry{
				ID:          1,
				ProjectID:   internalProject.ID,
				ProjectName: internalProject.Name,
				TaskID:      internalProject.Tasks[0].ID,
				TaskName:    internalProject.Tasks[0].Name,
				Description: "Internal work",
				Hours:       7.5,
				Editable:    true,
			},
		},
	},
	Day{
		UtilizationRate: 0.0,
		Hours:           0.0,
		HolidayName:     "Public holiday",
		Entries:         []Entry{},
	},
	Day{
		UtilizationRate: 0.0,
		Hours:           0.0,
		Entries:         []Entry{},
	},
	Day{
		UtilizationRate: 0.0,
		Hours:           7.5,
		Entries: []Entry{
			Entry{
				ID:          2,
				ProjectID:   absenceProject.ID,
				ProjectName: absenceProject.Name,
				TaskID:      absenceProject.Tasks[0].ID,
				TaskName:    absenceProject.Tasks[0].Name,
				Description: absenceProject.Tasks[0].LatestMarking,
				Hours:       7.5,
				Editable:    true,
			},
		},
	},
	Day{
		UtilizationRate: 0.0,
		Hours:           7.5,
		Entries: []Entry{
			Entry{
				ID:          3,
				ProjectID:   absenceProject.ID,
				ProjectName: absenceProject.Name,
				TaskID:      absenceProject.Tasks[0].ID,
				TaskName:    absenceProject.Tasks[0].Name,
				Description: absenceProject.Tasks[0].LatestMarking,
				Hours:       7.5,
				Editable:    true,
			},
		},
	},
	Day{
		UtilizationRate: 66.66666,
		Hours:           7.5,
		Entries: []Entry{
			Entry{
				ID:          4,
				ProjectID:   absenceProject.ID,
				ProjectName: absenceProject.Name,
				TaskID:      absenceProject.Tasks[2].ID,
				TaskName:    absenceProject.Tasks[2].Name,
				Description: absenceProject.Tasks[2].LatestMarking,
				Hours:       2.5,
				Editable:    true,
			},
			Entry{
				ID:          5,
				ProjectID:   customerProject.ID,
				ProjectName: customerProject.Name,
				TaskID:      customerProject.Tasks[0].ID,
				TaskName:    customerProject.Tasks[0].Name,
				Description: "Customer work",
				Hours:       5.0,
				Editable:    true,
			},
		},
	},
	Day{
		UtilizationRate: 100.0,
		Hours:           9.0,
		Entries: []Entry{
			Entry{
				ID:          6,
				ProjectID:   inactiveProject.ID,
				ProjectName: inactiveProject.Name,
				TaskID:      inactiveProject.Tasks[0].ID,
				TaskName:    inactiveProject.Tasks[0].Name,
				Description: inactiveProject.Tasks[0].LatestMarking,
				Hours:       9.0,
				Editable:    false,
			},
		},
	},
}

package api

var internalProject = Project{
	ID:   1,
	Name: "Internal work",
	Tasks: []Task{
		Task{
			ID:                1,
			Name:              "Things",
			LatestDescription: "Doing things",
		},
		Task{
			ID:                2,
			Name:              "Stuff",
			LatestDescription: "Doing stuff",
		},
	},
}

var absenceProject = Project{
	ID:   4,
	Name: "Absences",
	Tasks: []Task{
		Task{
			ID:                6,
			Name:              "Balance leave",
			LatestDescription: "Balance leave",
			Absence:           true,
		},
		Task{
			ID:                7,
			Name:              "Unpaid holiday",
			LatestDescription: "Unpaid holiday",
			Absence:           true,
		},
		Task{
			ID:                8,
			Name:              "Sick leave",
			LatestDescription: "Sick leave",
			Absence:           true,
		},
	},
}

var customerProject = Project{
	ID:   2,
	Name: "Actual customer work",
	Tasks: []Task{
		Task{
			ID:                3,
			Name:              "Development",
			LatestDescription: "Developing",
		},
		Task{
			ID:                4,
			Name:              "On-Call",
			LatestDescription: "Long weekend :(",
			Closed:            true,
		},
	},
}

var inactiveProject = Project{
	ID:     3,
	Name:   "Not active project",
	Closed: true,
	Tasks: []Task{
		Task{
			ID:                5,
			Name:              "Work",
			LatestDescription: "Doing work",
			Closed:            true,
		},
		Task{
			ID:                6,
			Name:              "Design",
			LatestDescription: "Designing",
			Closed:            true,
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
		Hours: 5,
		Entries: []Entry{
			Entry{
				ID:          1,
				ProjectID:   internalProject.ID,
				TaskID:      internalProject.Tasks[0].ID,
				Description: "Internal work",
				Hours:       5,
			},
		},
	},
	Day{
		Hours:       0.0,
		HolidayName: "Public holiday",
	},
	Day{
		Hours: 0.0,
	},
	Day{
		Hours: 0.0,
	},
	Day{
		Closed: true,
		Hours:  0.0,
	},
	Day{
		Hours: 7.5,
		Entries: []Entry{
			Entry{
				ID:          2,
				ProjectID:   absenceProject.ID,
				TaskID:      absenceProject.Tasks[0].ID,
				Description: absenceProject.Tasks[0].LatestDescription,
				Hours:       7.5,
			},
		},
	},
	Day{
		Closed: true,
		Hours:  7.5,
		Entries: []Entry{
			Entry{
				ID:          3,
				ProjectID:   absenceProject.ID,
				TaskID:      absenceProject.Tasks[0].ID,
				Description: absenceProject.Tasks[0].LatestDescription,
				Hours:       7.5,
			},
		},
	},
	Day{
		Hours: 10,
		Entries: []Entry{
			Entry{
				ID:          13,
				ProjectID:   customerProject.ID,
				TaskID:      customerProject.Tasks[1].ID,
				Description: customerProject.Tasks[1].LatestDescription,
				Hours:       10,
				Closed:      true,
			},
		},
	},
	Day{
		Hours: 7.5,
		Entries: []Entry{
			Entry{
				ID:          4,
				ProjectID:   absenceProject.ID,
				TaskID:      absenceProject.Tasks[2].ID,
				Description: absenceProject.Tasks[2].LatestDescription,
				Hours:       2.5,
			},
			Entry{
				ID:          5,
				ProjectID:   customerProject.ID,
				TaskID:      customerProject.Tasks[0].ID,
				Description: "Customer work",
				Hours:       5.0,
			},
		},
	},
	Day{
		Hours: 9.0,
		Entries: []Entry{
			Entry{
				ID:          6,
				ProjectID:   inactiveProject.ID,
				TaskID:      inactiveProject.Tasks[0].ID,
				Description: inactiveProject.Tasks[0].LatestDescription,
				Hours:       9.0,
				Closed:      true,
			},
		},
	},
	Day{
		Hours: 7.5,
		Entries: []Entry{
			Entry{
				ID:          7,
				ProjectID:   inactiveProject.ID,
				TaskID:      inactiveProject.Tasks[0].ID,
				Description: inactiveProject.Tasks[0].LatestDescription,
				Hours:       7.5,
			},
		},
	},
}

package api

var internalProject = Project{
	Id:   1,
	Name: "Internal work",
	Tasks: []Task{
		Task{
			Id:          1,
			Name:        "Things",
			LatestEntry: LatestEntry{Description: "Doing things"},
		},
		Task{
			Id:          2,
			Name:        "Stuff",
			LatestEntry: LatestEntry{Description: "Doing stuff"},
		},
	},
}

var absenceProject = Project{
	Id:   4,
	Name: "Absences",
	Tasks: []Task{
		Task{
			Id:          6,
			Name:        "Balance leave",
			LatestEntry: LatestEntry{Description: "Balance leave"},
			Absence:     true,
		},
		Task{
			Id:          7,
			Name:        "Unpaid holiday",
			LatestEntry: LatestEntry{Description: "Unpaid holiday"},
			Absence:     true,
		},
		Task{
			Id:          8,
			Name:        "Sick leave",
			LatestEntry: LatestEntry{Description: "Sick leave"},
			Absence:     true,
		},
	},
}

var customerProject = Project{
	Id:   2,
	Name: "Actual customer work",
	Tasks: []Task{
		Task{
			Id:          3,
			Name:        "Development",
			LatestEntry: LatestEntry{Description: "Developing"},
		},
		Task{
			Id:          4,
			Name:        "On-Call",
			LatestEntry: LatestEntry{Description: "Long weekend :("},
			Closed:      true,
		},
	},
}

var inactiveProject = Project{
	Id:     3,
	Name:   "Not active project",
	Closed: true,
	Tasks: []Task{
		Task{
			Id:          5,
			Name:        "Work",
			LatestEntry: LatestEntry{Description: "Doing work"},
			Closed:      true,
		},
		Task{
			Id:          6,
			Name:        "Design",
			LatestEntry: LatestEntry{Description: "Designing"},
			Closed:      true,
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
				Id:          1,
				ProjectId:   internalProject.Id,
				TaskId:      internalProject.Tasks[0].Id,
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
				Id:          2,
				ProjectId:   absenceProject.Id,
				TaskId:      absenceProject.Tasks[0].Id,
				Description: absenceProject.Tasks[0].LatestEntry.Description,
				Hours:       7.5,
			},
		},
	},
	Day{
		Closed: true,
		Hours:  7.5,
		Entries: []Entry{
			Entry{
				Id:          3,
				ProjectId:   absenceProject.Id,
				TaskId:      absenceProject.Tasks[0].Id,
				Description: absenceProject.Tasks[0].LatestEntry.Description,
				Hours:       7.5,
			},
		},
	},
	Day{
		Hours: 10,
		Entries: []Entry{
			Entry{
				Id:          13,
				ProjectId:   customerProject.Id,
				TaskId:      customerProject.Tasks[1].Id,
				Description: customerProject.Tasks[1].LatestEntry.Description,
				Hours:       10,
				Closed:      true,
			},
		},
	},
	Day{
		Hours: 7.5,
		Entries: []Entry{
			Entry{
				Id:          4,
				ProjectId:   absenceProject.Id,
				TaskId:      absenceProject.Tasks[2].Id,
				Description: absenceProject.Tasks[2].LatestEntry.Description,
				Hours:       2.5,
			},
			Entry{
				Id:          5,
				ProjectId:   customerProject.Id,
				TaskId:      customerProject.Tasks[0].Id,
				Description: "Customer work",
				Hours:       5.0,
			},
		},
	},
	Day{
		Hours: 9.0,
		Entries: []Entry{
			Entry{
				Id:          6,
				ProjectId:   inactiveProject.Id,
				TaskId:      inactiveProject.Tasks[0].Id,
				Description: inactiveProject.Tasks[0].LatestEntry.Description,
				Hours:       9.0,
				Closed:      true,
			},
		},
	},
	Day{
		Hours: 7.5,
		Entries: []Entry{
			Entry{
				Id:          7,
				ProjectId:   inactiveProject.Id,
				TaskId:      inactiveProject.Tasks[0].Id,
				Description: inactiveProject.Tasks[0].LatestEntry.Description,
				Hours:       7.5,
			},
		},
	},
}

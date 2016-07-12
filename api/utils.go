package api

import (
	"math"
	"math/rand"
	"time"
)

func RandomFloat64(start, end float64) float64 {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return start + r.Float64()*(end-start)
}

func RoundToHalf(val float64) float64 {
	whole, frac := math.Modf(val)
	switch {
	case math.Abs(frac) < 0.25:
		{
			return float64(whole)
		}
	case math.Abs(frac) < 0.75:
		{
			if whole >= 0 {
				return float64(whole) + 0.5
			} else {
				return float64(whole) - 0.5
			}
		}
	default:
		{
			if whole >= 0 {
				return float64(whole + 1)
			} else {
				return float64(whole - 1)
			}
		}
	}
}

func ShuffleProjects(projects []Project) {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i, project := range projects {
		ShuffleTasks(project.Tasks)
		j := r.Intn(i + 1)
		projects[i], projects[j] = projects[j], projects[i]
	}
}

func ShuffleTasks(tasks []Task) {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := range tasks {
		j := r.Intn(i + 1)
		tasks[i], tasks[j] = tasks[j], tasks[i]
	}
}

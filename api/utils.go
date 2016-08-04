package api

import (
	"math"
	"math/rand"
	"os"
	"strconv"
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

func RandomFail() bool {
	r := RandomFloat64(0, 1)
	env := os.Getenv("FAIL_RATIO")
	if env == "" {
		env = "0"
	}

	failRatio, err := strconv.ParseFloat(env, 64)
	if err != nil {
		return true
	}

	return r <= failRatio
}

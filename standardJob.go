package helpers

import (
	"math"
	"time"
)

type standardJob struct {
	JobID       string    `json:"jobid"`
	Title       string    `json:"title"`
	URL         string    `json:"url"`
	Company     string    `json:"company"`
	City        string    `json:"city"`
	CPA         float32   `json:"cpa,omitempty"`
	CPC         float32   `json:"cpc,omitempty"`
	Description string    `json:"description"`
	Date        time.Time `json:"date"`
}

func (sj *standardJob) EstimatedValue() float64 {
	CPAEstimatedValue := sj.CPA * 0.7
	CPCEstimatedValue := sj.CPC * 35
	return math.Max(float64(CPAEstimatedValue), float64(CPCEstimatedValue))
}

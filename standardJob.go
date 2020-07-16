package helpers

import (
	"math"
	"time"
)

// StandardJob generic job to be used accross the system
type StandardJob struct {
	JobID       string    `json:"jobid"`
	Title       string    `json:"title"`
	URL         string    `json:"url"`
	Company     string    `json:"company"`
	Slug     	string    `json:"slug"`
	City        string    `json:"city"`
	CPA         float32   `json:"cpa,omitempty"`
	CPC         float32   `json:"cpc,omitempty"`
	Description string    `json:"description"`
	Date        time.Time `json:"date"`
	Country     string    `json:"country"`
	ZIP     	string    `json:"zip"`
	State     	string    `json:"state"`
}

// EstimatedValue standard calc for the jobs value
func (sj *StandardJob) EstimatedValue() float64 {
	CPAEstimatedValue := sj.CPA * 0.7
	CPCEstimatedValue := sj.CPC * 35
	return math.Max(float64(CPAEstimatedValue), float64(CPCEstimatedValue))
}

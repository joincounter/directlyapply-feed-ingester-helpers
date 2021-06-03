package helpers

import (
	"math"
	"time"
)

// StandardJob generic job to be used accross the system
type StandardJob struct {
	JobID          string    `json:"jobid" db:"external_id"`
	Title          string    `json:"title" db:"title"`
	URL            string    `json:"url" db:"url"`
	Company        string    `json:"company" db:"company"`
	Slug           string    `json:"slug" db:"company_slug"`
	City           string    `json:"city" db:"city"`
	CPA            float32   `json:"cpa,omitempty" db:"cpa"`
	CPC            float32   `json:"cpc,omitempty" db:"cpc"`
	Description    string    `json:"description" db:"description"`
	Date           time.Time `json:"date" db:"date"`
	Country        string    `json:"country" db:"country"`
	ZIP            string    `json:"zip" db:"zip"`
	State          string    `json:"state" db:"state"`
	Location       string    `json:"location" db:"location"`
	RemoteAvalible bool      `json:"remoteAvalible" db:"remote_avalible"`
	Recruiter      bool      `json:"recruiter" db:"recruiter"`
	SalaryMin      float64   `bson:"salaryMinimum" json:"salaryMinimum"`
	SalaryMax      float64   `bson:"salaryMaximum" json:"salaryMaximum"`
	SalaryType     string    `bson:"salaryType" json:"salaryType"`
	Category       string    `json:"category"`
	Expires        time.Time `json:"expires" bson:"expires"`
}

// EstimatedValue standard calc for the jobs value
func (sj *StandardJob) EstimatedValue() float64 {
	CPAEstimatedValue := sj.CPA * 0.7
	CPCEstimatedValue := sj.CPC * 35
	return math.Max(float64(CPAEstimatedValue), float64(CPCEstimatedValue))
}

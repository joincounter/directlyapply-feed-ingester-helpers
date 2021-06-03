package helpers

import (
	"math"
	"time"
)

// StandardJob generic job to be used accross the system
type StandardJob struct {
	JobID          string    `json:"jobid" bson:"jobid" db:"external_id"`
	Title          string    `json:"title" bson:"title" db:"title"`
	URL            string    `json:"url" bson:"url" db:"url"`
	Company        string    `json:"company" bson:"company" db:"company"`
	Slug           string    `json:"slug" bson:"slug" db:"company_slug"`
	City           string    `json:"city" bson:"city" db:"city"`
	CPA            float32   `json:"cpa,omitempty" bson:"cpa" db:"cpa"`
	CPC            float32   `json:"cpc,omitempty" bson:"cpc" db:"cpc"`
	Description    string    `json:"description" bson:"description" db:"description"`
	Date           time.Time `json:"date" bson:"posted" db:"date"`
	Country        string    `json:"country" bson:"country" db:"country"`
	ZIP            string    `json:"zip" bson:"zip" db:"zip"`
	State          string    `json:"state" bson:"state" db:"state"`
	Location       string    `json:"location" bson:"location" db:"location"`
	RemoteAvalible bool      `json:"remoteAvalible" bson:"remoteAvalible" db:"remote_avalible"`
	Recruiter      bool      `json:"recruiter" bson:"recruiter" db:"recruiter"`
	SalaryMin      float64   `json:"salaryMinimum" bson:"salaryMinimum"`
	SalaryMax      float64   `json:"salaryMaximum" bson:"salaryMaximum"`
	SalaryType     string    `json:"salaryType" bson:"salaryType"`
	Category       string    `json:"category" bson:"category"`
	Expires        time.Time `json:"expires" bson:"expires"`
	Summaries      []string  `json:"summaries" bson:"summaries"`
}

// EstimatedValue standard calc for the jobs value
func (sj *StandardJob) EstimatedValue() float64 {
	CPAEstimatedValue := sj.CPA * 0.7
	CPCEstimatedValue := sj.CPC * 35
	return math.Max(float64(CPAEstimatedValue), float64(CPCEstimatedValue))
}

package helpers

import (
	"fmt"
	"math"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
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

type LegacyUploadJob struct {
	Title        string             `bson:"title" json:"title"`
	Company      string             `bson:"company" json:"company"`
	Location     string             `bson:"location" json:"location"`
	City         string             `bson:"city" json:"city"`
	ZIP          string             `bson:"zip" json:"zip"`
	State        string             `bson:"state" json:"state"`
	SourceID     string             `bson:"sourceid" json:"locatsourceidion"`
	ID           primitive.ObjectID `bson:"_id,omitempty" json:"_id"`
	Description  string             `bson:"description" json:"description"`
	Summary      string             `bson:"summary" json:"summary"`
	Posted       time.Time          `bson:"posted" json:"posted"`
	Expires      time.Time          `bson:"expires" json:"expires"`
	EasyApply    bool               `bson:"easyapply" json:"easyapply"`
	SourceURL    string             `bson:"sourceurl" json:"sourceurl"`
	SecondaryURL string             `bson:"secondaryurl" json:"secondaryurl"`
	Slug         string             `bson:"slug" json:"slug"`
	GA           string             `bson:"ga" json:"ga"`
	BadLink      bool               `bson:"badlink" json:"badlink"`
	Logo         string             `bson:"logo" json:"logo"`
	CPC          float64            `bson:"cpc" json:"cpc"`
	CPA          float64            `bson:"cpa" json:"cpa"`
	SalaryMin    float64            `bson:"salaryMinimum" json:"salaryMinimum"`
	SalaryMax    float64            `bson:"salaryMaximum" json:"salaryMaximum"`
	SalaryType   string             `bson:"salaryType" json:"salaryType"`
}

func ConvertToLegacyUploadJob(jobRaw StandardJob) LegacyUploadJob {
	job := LegacyUploadJob{}

	job.Title = jobRaw.Title
	job.Description = jobRaw.Description

	job.Company = jobRaw.Company
	job.Slug = GenerateSlug(job.Company)
	job.Location = jobRaw.Location
	job.City = jobRaw.City
	job.ZIP = jobRaw.ZIP
	job.CPC = float64(jobRaw.CPC)
	job.CPA = float64(jobRaw.CPA)

	job.SalaryMin = jobRaw.SalaryMin
	job.SalaryMax = jobRaw.SalaryMax
	job.SalaryType = jobRaw.SalaryType

	job.Posted = jobRaw.Date
	job.Expires = jobRaw.Date.AddDate(0, 0, 7*4)

	job.Summary = formatListOfStringsAsSummary(jobRaw.Summaries)

	job.SourceURL = jobRaw.URL
	job.SecondaryURL = jobRaw.URL

	return job
}

func formatListOfStringsAsSummary(strs []string) string {
	innerhtml := ""
	for i := 0; i < len(strs); i++ {
		innerhtml = innerhtml + fmt.Sprintf("<li>%s</li>", strs[i])
	}
	return fmt.Sprintf("<ul>%s</ul>", innerhtml)
}

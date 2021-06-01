package helpers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"time"
)

type workableJobs struct {
	Jobs []workableJob `json:"jobs"`
}

type workableJob struct {
	Jobid        string          `json:"id"`
	Title        string          `json:"title"`
	Description  string          `json:"description"`
	Benefits     string          `json:"benefitsSection"`
	Requirements string          `json:"requirementsSection"`
	URL          string          `json:"applyUrl"`
	Location     []string        `json:"locations"`
	Company      workableCompany `json:"company"`
	Date         string          `json:"created"`
	// Category    string  `json:"category"`
}

type workableCompany struct {
	Name    string `json:"title"`
	Website string `json:"website"`
	Logo    string `json:"logo"`
}

// NeuvooConverter convert Neuvoo jobs to standard jobs
func WorkableConverter(file *os.File) (*[]StandardJob, error) {

	var jobsTemp workableJobs
	jobsFinal := make([]StandardJob, 0)
	byteValue, _ := ioutil.ReadAll(file)

	err := json.Unmarshal(byteValue, &jobsTemp)
	if err != nil {
		fmt.Println(err)
	}

	for i := 0; i < len(jobsTemp.Jobs); i++ {
		job := jobsTemp.Jobs[i]
		date, err := time.Parse(time.RFC3339, job.Date)
		if err != nil {
			fmt.Printf("error parsing date: title: %s err: %s", job.Title, err.Error())
		}
		jobsFinal = append(jobsFinal, StandardJob{
			Title:       job.Title,
			Description: job.Description + job.Requirements + job.Benefits,
			Company:     job.Company.Name,
			JobID:       job.Jobid,
			City:        strings.Split(job.Location[0], ", ")[0],
			URL:         job.URL,
			Slug:        GenerateSlug(job.Company.Name),
			CPA:         0,
			CPC:         0,
			Country:     strings.Split(job.Location[0], ", ")[len(strings.Split(job.Location[0], ", "))-1],
			Date:        date,
		})
	}

	return &jobsFinal, nil
}

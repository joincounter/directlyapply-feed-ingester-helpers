package converters

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"time"

	extern_helpers "github.com/joincounter/directlyapply-feed-ingester-helpers"
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
func WorkableConverter(file *os.File) (*[]extern_helpers.StandardJob, error) {

	var jobsTemp workableJobs
	jobsFinal := make([]extern_helpers.StandardJob, 0)
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
		jobsFinal = append(jobsFinal, extern_helpers.StandardJob{
			Title:       job.Title,
			Description: job.Description + job.Requirements + job.Benefits,
			Company:     job.Company.Name,
			JobID:       job.Jobid,
			City:        strings.Split(job.Location[0], ", ")[0],
			URL:         job.URL,
			Slug:        extern_helpers.GenerateSlug(job.Company.Name),
			CPA:         0,
			CPC:         0,
			Country:     strings.Split(job.Location[0], ", ")[len(strings.Split(job.Location[0], ", "))-1],
			Date:        date,
		})
	}

	return &jobsFinal, nil
}

package helpers

import (
	"encoding/xml"
	"fmt"
	"io"
	"os"
	"strings"
	"time"
)

type resumeLibJobs struct {
	XMLName xml.Name   `xml:"source"`
	Text    string     `xml:",chardata"`
	Job     []resumeLibJob `xml:"job"`
}

type resumeLibJob struct {
	// Included
	Jobid       string `xml:"job_id"`
	Title       string `xml:"job_title"`
	Jobtype     string `xml:"job_type"`
	Description string `xml:"description"`
	Country     string `xml:"country"`
	Location    string  `xml:"location_text"`
	Company     string `xml:"company"`
	Date        string `xml:"post_date"`
	URL         string `xml:"apply_url"`
	// Addedd Below
	City        string `xml:"city"`
	State       string `xml:"state"`
	// Excluded
	ZIP         string `xml:"postalcode"`
	Category    string `xml:"category"`
	Sponsored   string   `xml:"sponsored"`
	CPC         float32 `xml:"cpc"`
	CPA         float32 `xml:"cpa"`
}

// RecruiticsConverter convert Recruitics jobs to standard jobs
func ResumeLibraryConverter(file *os.File) (*[]StandardJob, error) {

	jobs := make([]StandardJob, 0)

	decoder := xml.NewDecoder(file)

	for {
		// Read tokens from the XML document in a stream.
		token, err := decoder.Token()
		if token == nil || err == io.EOF {
			// EOF means we're done.
			break
		} else if err != nil {
			return nil, err
		}

		// Inspect the type of the token just read.
		switch se := token.(type) {
		case xml.StartElement:
			if se.Name.Local == "job" {
				var job recruiticsJob
				err = decoder.DecodeElement(&job, &se)

				if err != nil {
					fmt.Printf("continuing: error occured while decoding xml: %s", err)
					continue
				}

				date, timeError := time.Parse("Mon, 2 Jan 2006 15:04:05 MST", job.Date)

				if timeError != nil {
					fmt.Printf("error parsing date: title: %s err: %s", job.Title, timeError.Error())
					return nil, err
				}

				job.City = strings.Split(job.Location, ",")[0]
				job.CPA = 1.0

				jobs = append(jobs, StandardJob{
					Title:       job.Title,
					JobID:       job.Jobid,
					URL:         job.URL,
					Company:     job.Company,
					Slug:        GenerateSlug(job.Company),
					City:        job.City,
					Location:    job.Location,
					CPA:         job.CPA,
					CPC:         job.CPC,
					Description: job.Description,
					Date:        date,
					Country:     job.Country,
					ZIP:         job.ZIP,
					State:       job.State,
				})
			}
		default:
		}
	}

	return &jobs, nil
}

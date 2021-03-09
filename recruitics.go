package helpers

import (
	"encoding/xml"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
	"time"
)

type recruiticsJobs struct {
	XMLName xml.Name   `xml:"source"`
	Text    string     `xml:",chardata"`
	Job     []recruiticsJob `xml:"job"`
}

type recruiticsJob struct {
	Title       string `xml:"title"`
	Date        string `xml:"date"`
	Jobid       string `xml:"referencenumber"`
	URL         string `xml:"url"`
	Company     string `xml:"company"`
	City        string `xml:"city"`
	Location    string  `xml:"location"`
	State       string `xml:"state"`
	Country     string `xml:"country"`
	ZIP         string `xml:"postalcode"`
	Description string `xml:"description"`
	Category    string `xml:"category"`
	Jobtype     string `xml:"jobtype"`
	Sponsored   string   `xml:"sponsored"`
	CPC         float32 `xml:"cpc"`
	CPA         float32 `xml:"cpa"`
	
}

// RecruiticsConverter convert Recruitics jobs to standard jobs
func RecruiticsConverter(file *os.File) (*[]StandardJob, error) {

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

				if job.Sponsored != "" {
					cpcString := strings.Split(job.Sponsored, " ")[1]
					float, _ := strconv.ParseFloat(cpcString, 32)
					job.CPC = float32(float)
					job.CPA = float32(float)
				}


				job.Location = job.City + ", " + job.State + ", " + job.Country

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

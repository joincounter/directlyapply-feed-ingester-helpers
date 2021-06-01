package helpers

import (
	"encoding/xml"
	"fmt"
	"io"
	"os"
	"strings"
	"time"
)

type rawCVLibrary struct {
	Text        string `xml:",chardata"`
	Jobref      string `xml:"jobref"`
	Date        string `xml:"date"`
	Title       string `xml:"title"`
	Company     string `xml:"company"`
	Email       string `xml:"email"`
	URL         string `xml:"url"`
	Salarymin   string `xml:"salarymin"`
	Salarymax   string `xml:"salarymax"`
	Benefits    string `xml:"benefits"`
	Salary      string `xml:"salary"`
	Jobtype     string `xml:"jobtype"`
	FullPart    string `xml:"full_part"`
	SalaryPer   string `xml:"salary_per"`
	Location    string `xml:"location"`
	City        string `xml:"city"`
	County      string `xml:"county"`
	Country     string `xml:"country"`
	Description string `xml:"description"`
	Category    string `xml:"category"`
	Image       string `xml:"image"`
}

// RecruiticsConverter convert Recruitics jobs to standard jobs
func CVLibraryConverter(file *os.File) (*[]StandardJob, error) {

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
				var job rawCVLibrary
				err = decoder.DecodeElement(&job, &se)

				if err != nil {
					fmt.Printf("continuing: error occured while decoding xml: %s", err)
					continue
				}

				date, timeError := time.Parse("2006-01-02 15:04:05", job.Date)

				if timeError != nil {
					fmt.Printf("error parsing date: title: %s err: %s", job.Title, timeError.Error())
					return nil, err
				}

				job.City = strings.Split(job.Location, ",")[0]

				jobs = append(jobs, StandardJob{
					Title:       job.Title,
					JobID:       job.Jobref,
					URL:         job.URL,
					Company:     job.Company,
					Slug:        GenerateSlug(job.Company),
					City:        job.City,
					Location:    job.Location,
					CPA:         2.0,
					CPC:         2.0,
					Description: job.Description,
					Date:        date,
					Country:     job.Country,
					State:       job.County,
				})
			}
		default:
		}
	}

	return &jobs, nil
}

package helpers

import (
	"encoding/xml"
	"fmt"
	"io"
	"os"
	"time"

	"github.com/dylankbuckley/zippia/helpers"
)

type neuvooJobs struct {
	XMLName xml.Name    `xml:"jobs"`
	Text    string      `xml:",chardata"`
	Job     []neuvooJob `xml:"job"`
}

type neuvooJob struct {
	Text        string  `xml:",chardata"`
	Jobid       string  `xml:"jobid"`
	Title       string  `xml:"title"`
	Category    string  `xml:"category"`
	City        string  `xml:"city"`
	State       string  `xml:"state"`
	Country     string  `xml:"country"`
	Jobtype     string  `xml:"jobtype"`
	Description string  `xml:"description"`
	URL         string  `xml:"url"`
	Date        string  `xml:"date"`
	CPC         float32 `xml:"cpc"`
	Currency    string  `xml:"currency"`
	Logo        string  `xml:"logo"`
	Company     string  `xml:"company"`
}

// NeuvooConverter convert Neuvoo jobs to standard jobs
func NeuvooConverter(file *os.File) (*[]StandardJob, error) {

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
				var job neuvooJob
				err = decoder.DecodeElement(&job, &se)

				if err != nil {
					fmt.Printf("continuing: error occured while decoding xml: %s", err)
					continue
				}

				date, err := time.Parse(time.RFC3339, job.Date)

				if err != nil {
					fmt.Printf("error parsing date: title: %s err: %s", job.Title, err.Error())
				} else {
					jobs = append(jobs, StandardJob{
						Title:       job.Title,
						JobID:       job.Jobid,
						URL:         job.URL,
						Company:     job.Company,
						Slug:		 helpers.GenerateSlug(job.Company),
						City:        job.City,
						CPA:         0,
						CPC:         job.CPC,
						Description: job.Description,
						Date:        date,
						Country:     job.Country,
					})
				}
			}
		default:
		}
	}

	return &jobs, nil
}

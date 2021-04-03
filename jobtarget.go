package helpers

import (
	"encoding/xml"
	"fmt"
	"io"
	"os"
	"time"
)

type jobtargetJobs struct {
	XMLName xml.Name    `xml:"jobs"`
	Text    string      `xml:",chardata"`
	Job     []jobtargetJob `xml:"job"`
}

type jobtargetJob struct {
	Text        string  `xml:",chardata"`
	Jobid       string  `xml:"name"`
	Title       string  `xml:"position"`
	Category    string  `xml:"function"`
	City        string  `xml:"city"`
	State       string  `xml:"state"`
	Country     string  `xml:"country"`
	ZIP     	string  `xml:"zip"`
	Jobtype     string  `xml:"jobtype"`
	Description string  `xml:"description"`
	URL         string  `xml:"apply_url"`
	Date        string  `xml:"start"`
	CPC         float32 `xml:"cpc"`
	Currency    string  `xml:"currency"`
	Logo        string  `xml:"logo"`
	Company     string  `xml:"company"`
}

// JobTargetConverter convert JobTarget jobs to standard jobs
func JobTargetConverter(file *os.File) (*[]StandardJob, error) {

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
				var job jobtargetJob
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
						Slug:        GenerateSlug(job.Company),
						City:        job.City,
						CPA:         0,
						CPC:         job.CPC,
						Description: job.Description,
						Date:        date,
						Country:     job.Country,
						Category:    job.Category,
						ZIP: 		 job.ZIP,
						State: 		 job.State,
					})
				}
			}
		default:
		}
	}

	return &jobs, nil
}

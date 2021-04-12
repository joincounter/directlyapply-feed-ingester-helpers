package helpers

import (
	"encoding/xml"
	"fmt"
	"io"
	"os"
	"time"
)

type indeedJobs struct {
	XMLName xml.Name    `xml:"rss"`
	Text    string      `xml:",chardata"`
	Job     []indeedJob `xml:"item"`
}

type indeedJob struct {

	Jobid       string  `xml:"id"`
	Date        string  `xml:"date"`
	Title       string  `xml:"title"`
	Company     string  `xml:"company"`
	URL         string  `xml:"url"`
	Jobtype     string  `xml:"jobtype"`
	Country     string  `xml:"country"`
	City        string  `xml:"location"`
	Description string  `xml:"description"`
	CPC         float32 `xml:"bid-value"`




	// Text        string  `xml:",chardata"`
	// Category    string  `xml:"category"`
	// State       string  `xml:"state"`
	// ZIP     	string  `xml:"zipcode"`
	// Currency    string  `xml:"currency"`
	// Logo        string  `xml:"logo"`
	
}

// apploiConverter convert apploi jobs to standard jobs
func IndeedConverter(file *os.File) (*[]StandardJob, error) {

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
				var job apploiJob
				err = decoder.DecodeElement(&job, &se)

				if err != nil {
					fmt.Printf("continuing: error occured while decoding xml: %s", err)
					continue
				}

				// Wed, 27 Jan 2021 05:00:03 GMT

				date, err := time.Parse(time.RFC1123, job.Date)

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
						CPA:         job.CPC,
						CPC:         0,
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

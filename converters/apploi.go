package converters

import (
	"encoding/xml"
	"fmt"
	"io"
	"os"
	"time"

	extern_helpers "github.com/joincounter/directlyapply-feed-ingester-helpers"
)

type apploiJobs struct {
	XMLName xml.Name    `xml:"rss"`
	Text    string      `xml:",chardata"`
	Job     []apploiJob `xml:"item"`
}

type apploiJob struct {
	Text        string `xml:",chardata"`
	Jobid       string `xml:"referencenumber"`
	Title       string `xml:"title"`
	Date        string `xml:"pubDate"`
	Category    string `xml:"category"`
	Company     string `xml:"company"`
	City        string `xml:"city"`
	State       string `xml:"state"`
	ZIP         string `xml:"zipcode"`
	Description string `xml:"description"`
	URL         string `xml:"link"`
	// Not Included
	Country  string  `xml:"country"`
	Jobtype  string  `xml:"jobtype"`
	CPC      float32 `xml:"cpc"`
	Currency string  `xml:"currency"`
	Logo     string  `xml:"logo"`
}

// apploiConverter convert apploi jobs to standard jobs
func ApploiConverter(file *os.File) (*[]extern_helpers.StandardJob, error) {

	fmt.Println("Sup Apploi")

	jobs := make([]extern_helpers.StandardJob, 0)

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
			if se.Name.Local == "item" {
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
					jobs = append(jobs, extern_helpers.StandardJob{
						Title:       job.Title,
						JobID:       job.Jobid,
						URL:         job.URL,
						Company:     job.Company,
						Slug:        extern_helpers.GenerateSlug(job.Company),
						City:        job.City,
						CPA:         0,
						CPC:         0,
						Description: job.Description,
						Date:        date,
						Country:     job.Country,
						Category:    job.Category,
						ZIP:         job.ZIP,
						State:       job.State,
					})
				}
			}
		default:
		}
	}

	return &jobs, nil
}

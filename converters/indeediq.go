package converters

import (
	"encoding/xml"
	"fmt"
	"io"
	"os"
	"time"

	extern_helpers "github.com/joincounter/directlyapply-feed-ingester-helpers"
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
func IndeedConverter(file *os.File) (*[]extern_helpers.StandardJob, error) {

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
			if se.Name.Local == "job" {
				var job indeedJob
				err = decoder.DecodeElement(&job, &se)

				if err != nil {
					fmt.Printf("continuing: error occured while decoding xml: %s", err)
					continue
				}

				//08/03/2021

				date, err := time.Parse("02/01/2006", job.Date)

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

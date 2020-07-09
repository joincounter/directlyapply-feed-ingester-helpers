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

type joveoJobs struct {
	XMLName xml.Name    `xml:"source"`
	Text    string      `xml:",chardata"`
	Job     []joveoJob `xml:"job"`
}

type joveoJob struct {
	Date        string  `xml:"date"`
	Country     string  `xml:"country"`
	City        string  `xml:"city"`
	Jobid       string  `xml:"referencenumber"`
	Description string  `xml:"description"`
	Title       string  `xml:"title"`
	Jobtype     string  `xml:"type"`
	URL         string  `xml:"url"`
	ZIP         string  `xml:"postalcode"`
	CPC         string `xml:"cpc"`
	CPA         string `xml:"cpa"`
	Company     string  `xml:"company"`
	Category    string  `xml:"category"`
	State       string  `xml:"state"`
}

// JoveoConverter convert Joveo jobs to standard jobs
func JoveoConverter(file *os.File) (*[]StandardJob, error) {

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
				var job joveoJob
				err = decoder.DecodeElement(&job, &se)

				if err != nil {
					fmt.Printf("continuing: error occured while decoding xml: %s", err)
					continue
				}


				date, timeError := time.Parse(time.RFC3339, job.Date)

				cpa, _ := strconv.ParseFloat(strings.Split(job.CPA, " ")[1], 32)
				cpc, _ := strconv.ParseFloat(strings.Split(job.CPC, " ")[1], 32)

				if timeError != nil {
					fmt.Printf("error parsing date: title: %s err: %s", job.Title, timeError.Error())
				} else {
					jobs = append(jobs, StandardJob{
						Title:       job.Title,
						JobID:       job.Jobid,
						URL:         job.URL,
						Company:     job.Company,
						City:        job.City,
						CPA:         float32(cpa),
						CPC:         float32(cpc),
						Description: job.Description,
						Date:        date,
						Country:     job.Country,
						ZIP:		 job.ZIP,
						State:		 job.State,
					})
				}
			}
		default:
		}
	}

	return &jobs, nil
}

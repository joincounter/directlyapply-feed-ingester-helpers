package converters

import (
	"encoding/xml"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
	"time"

	extern_helpers "github.com/joincounter/directlyapply-feed-ingester-helpers"
)

type joveoJobs struct {
	XMLName xml.Name   `xml:"source"`
	Text    string     `xml:",chardata"`
	Job     []joveoJob `xml:"job"`
}

type joveoJob struct {
	Date        string `xml:"date"`
	Country     string `xml:"country"`
	City        string `xml:"city"`
	Jobid       string `xml:"referencenumber"`
	Description string `xml:"description"`
	Title       string `xml:"title"`
	Jobtype     string `xml:"type"`
	URL         string `xml:"url"`
	ZIP         string `xml:"postalcode"`
	CPC         string `xml:"cpc"`
	CPA         string `xml:"cpa"`
	Company     string `xml:"company"`
	Category    string `xml:"category"`
	State       string `xml:"state"`
}

// JoveoConverter convert Joveo jobs to standard jobs
func JoveoConverter(file *os.File) (*[]extern_helpers.StandardJob, error) {

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
				var job joveoJob
				err = decoder.DecodeElement(&job, &se)

				if err != nil {
					fmt.Printf("continuing: error occured while decoding xml: %s", err)
					continue
				}

				date, timeError := time.Parse("2006-01-02 15:04:05.000 MST", job.Date)

				if timeError != nil {
					fmt.Printf("error parsing date: title: %s err: %s", job.Title, timeError.Error())
					return nil, err
				}

				var tryParseCPA = func(s string) float32 {
					split := strings.Split(job.CPA, " ")
					lastBit := split[len(split)-1]
					cpa, err := strconv.ParseFloat(lastBit, 32)
					if err != nil {
						return 0
					}
					return float32(cpa)
				}

				var tryParseCPC = func(s string) float32 {
					split := strings.Split(job.CPC, " ")
					lastBit := split[len(split)-1]
					cpc, err := strconv.ParseFloat(lastBit, 32)
					if err != nil {
						return 0
					}
					return float32(cpc)
				}

				jobs = append(jobs, extern_helpers.StandardJob{
					Title:       job.Title,
					JobID:       job.Jobid,
					URL:         job.URL,
					Company:     job.Company,
					Slug:        extern_helpers.GenerateSlug(job.Company),
					City:        job.City,
					CPA:         tryParseCPA(job.CPA),
					CPC:         tryParseCPC(job.CPC),
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

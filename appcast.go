package helpers

import (
	"encoding/xml"
	"fmt"
	"io"
	"os"
	"strconv"
	"time"
)

type appcastRoot struct {
	XMLName xml.Name    `xml:"source"`
	Jobs    appcastJobs `xml:"jobs"`
}

type appcastJobs struct {
	XMLName xml.Name     `xml:"jobs"`
	Jobs    []rawAppCast `xml:"job"`
}

type rawAppCast struct {
	XMLName     xml.Name `xml:"job"`
	Title       string   `xml:"title"`
	Company     string   `xml:"company"`
	Description string   `xml:"body"`
	URL         string   `xml:"url"`
	City        string   `xml:"city"`
	Posted      string   `xml:"posted_at"`
	Type        string   `xml:"job_type"`
	SourceID    string   `xml:"job_reference"`
	Country     string   `xml:"country"`
	Zip         string   `xml:"zip"`
	Location    string   `xml:"location"`
	State       string   `xml:"state"`
	Category    string   `xml:"category"`
	CPC         string   `xml:"cpc"`
	CPA         string   `xml:"cpa"`
}

// AppcastConverter convert Appcast jobs to standard
func AppcastConverter(file *os.File) (*[]StandardJob, error) {

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
				var job rawAppCast
				err = decoder.DecodeElement(&job, &se)

				if err != nil {
					fmt.Printf("continuing: error occured while decoding xml: %s", err)
					continue
				}

				date, err := time.Parse("2006-01-02", job.Posted)

				if err != nil {
					fmt.Printf("error parsing date: title: %s err: %s", job.Title, err.Error())
				} else {
					newCpa, _ := strconv.ParseFloat(job.CPA, 32)
					newCpc, _ := strconv.ParseFloat(job.CPC, 32)

					jobs = append(jobs, StandardJob{
						Title:       job.Title,
						JobID:       job.SourceID,
						URL:         job.URL,
						Company:     job.Company,
						Slug:        GenerateSlug(job.Company),
						City:        job.City,
						State:       job.State,
						ZIP:         job.Zip,
						Location:    job.Location,
						CPA:         float32(newCpa),
						CPC:         float32(newCpc),
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

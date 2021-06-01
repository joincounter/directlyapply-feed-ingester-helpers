package helpers

import (
	"encoding/xml"
	"fmt"
	"io"
	"os"
	"time"
)

type zipRecruiterJob struct {
	Text            string  `xml:",chardata"`
	City            string  `xml:"city"`
	Referencenumber string  `xml:"referencenumber"`
	Category        string  `xml:"category"`
	State           string  `xml:"state"`
	Postalcode      string  `xml:"postalcode"`
	Company         string  `xml:"company"`
	Title           string  `xml:"title"`
	URL             string  `xml:"url"`
	Campaign        string  `xml:"campaign"`
	CPC             float32 `xml:"cpc"`
	Country         string  `xml:"country"`
	Description     string  `xml:"description"`
	Date            string  `xml:"date"`
}

func ZipRecruiterConverter(file *os.File) (*[]StandardJob, error) {

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
				var job zipRecruiterJob
				err = decoder.DecodeElement(&job, &se)

				if err != nil {
					fmt.Printf("continuing: error occured while decoding xml: %s", err)
					continue
				}

				date, err := time.Parse("Mon, 2 Jan 2006 15:04:05 MST", job.Date)

				if err != nil {
					fmt.Printf("error parsing date: title: %s err: %s", job.Title, err.Error())
				} else {
					jobs = append(jobs, StandardJob{
						Title:       job.Title,
						JobID:       job.Referencenumber,
						URL:         job.URL,
						Company:     job.Company,
						Slug:        GenerateSlug(job.Company),
						City:        job.City,
						CPA:         0,
						CPC:         job.CPC,
						Description: job.Description,
						Date:        date,
						Country:     job.Country,
						ZIP:         job.Postalcode,
						State:       job.State,
					})
				}
			}
		default:
		}
	}

	return &jobs, nil
}

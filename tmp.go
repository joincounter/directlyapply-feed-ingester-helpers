package helpers

import (
	"encoding/xml"
	"fmt"
	"io"
	"os"
	"time"
)

type tmpJob struct {
	Text        string  `xml:",chardata"`
	Title       string  `xml:"title"`
	ID          string  `xml:"id"`
	Date        string  `xml:"date"`
	URL         string  `xml:"url"`
	Company     string  `xml:"company"`
	City        string  `xml:"city"`
	State       string  `xml:"state"`
	Country     string  `xml:"country"`
	PostalCode  string  `xml:"postalcode"`
	Description string  `xml:"description"`
	Salary      string  `xml:"salary"`
	CPC         float32 `xml:"cpc"`
}

func TmpConverter(file *os.File) (*[]StandardJob, error) {

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
				var job tmpJob
				err = decoder.DecodeElement(&job, &se)

				if err != nil {
					fmt.Printf("continuing: error occured while decoding xml: %s", err)
					continue
				}

				date, err := time.Parse("Mon, 02 Jan 2006 15:04:05 MST", job.Date)

				if err != nil {
					fmt.Printf("error parsing date: title: %s err: %s", job.Title, err.Error())
				} else {
					jobs = append(jobs, StandardJob{
						Title:       job.Title,
						JobID:       job.ID,
						URL:         job.URL,
						Company:     job.Company,
						Slug:        GenerateSlug(job.Company),
						City:        job.City,
						CPA:         0,
						CPC:         job.CPC,
						Description: job.Description,
						Date:        date,
						Country:     job.Country,
						ZIP:         job.PostalCode,
						State:       job.State,
					})
				}
			}
		default:
		}
	}

	return &jobs, nil
}

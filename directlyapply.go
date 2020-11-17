package helpers

import (
	"encoding/xml"
	"fmt"
	"io"
	"os"
	"time"
)

// AppcastConverter convert Appcast jobs to standard
func DirectlyApplyConverter(file *os.File) (*[]StandardJob, error) {

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

				date, err := time.Parse("Mon, 2 Jan 2006 15:04:05 MST", job.Posted)

				if err != nil {
					fmt.Printf("error parsing date: title: %s err: %s", job.Title, err.Error())
				} else {
					jobs = append(jobs, StandardJob{
						Title:       job.Title,
						JobID:       job.SourceID,
						URL:         job.URL,
						Company:     job.Company,
						Slug:        GenerateSlug(job.Company),
						City:        job.City,
						State:        job.State,
						ZIP:		job.Zip,
						Location: job.Location,
						CPA:         job.CPA,
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

package converters

import (
	"encoding/xml"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"
	"time"

	extern_helpers "github.com/joincounter/directlyapply-feed-ingester-helpers"
	"golang.org/x/text/encoding/charmap"
)

func makeCharsetReader(charset string, input io.Reader) (io.Reader, error) {
	if charset == "Windows-1252" {
		// Windows-1252 is a superset of ISO-8859-1, so should do here
		return charmap.Windows1252.NewDecoder().Reader(input), nil
	}
	return nil, fmt.Errorf("unknown charset: %s", charset)
}

func decodeWindows1252String(originalString string) string {
	sr := strings.NewReader(originalString)
	tr := charmap.Windows1252.NewDecoder().Reader(sr)
	b, e := ioutil.ReadAll(tr)
	if e != nil {
		fmt.Println("error:", e)
	}
	return string(b)
}

type rawPando struct {
	XMLName     xml.Name `xml:"job"`
	Title       string   `xml:"title"`
	Date        string   `xml:"date"`
	JobID       string   `xml:"referencenumber"`
	URL         string   `xml:"url"`
	State       string   `xml:"state"`
	Country     string   `xml:"country"`
	ZIP         string   `xml:"postalcode"`
	City        string   `xml:"city"`
	Description string   `xml:"description"`
	CPC         float32  `xml:"cpc"`
	CPA         float32  `xml:"cpa"`
	Category    string   `xml:"category"`
	Company     string   `xml:"company"`
}

type jobsPando struct {
	XMLName      xml.Name   `xml:"source"`
	Text         string     `xml:",chardata"`
	Fo           string     `xml:"fo,attr"`
	Publisher    string     `xml:"publisher"`
	Publisherurl string     `xml:"publisherurl"`
	Job          []rawPando `xml:"job"`
}

func PandoLogicConverter(file *os.File) (*[]extern_helpers.StandardJob, error) {
	jobs := make([]extern_helpers.StandardJob, 0)

	var j1 jobsPando

	decoder := xml.NewDecoder(file)
	decoder.CharsetReader = makeCharsetReader
	err := decoder.Decode(&j1)
	if err != nil {
		return nil, err
	}

	for _, pandoJob := range j1.Job {
		date, err := time.Parse("2006-01-02T15:04:05", strings.Split(pandoJob.Date, ".")[0])
		if err != nil {
			fmt.Println(err.Error())
		}

		sj := extern_helpers.StandardJob{
			JobID:       pandoJob.JobID,
			Title:       decodeWindows1252String(pandoJob.Title),
			Description: decodeWindows1252String(pandoJob.Description),
			Date:        date,
			URL:         pandoJob.URL,
			Company:     pandoJob.Company,
			Slug:        extern_helpers.GenerateSlug(pandoJob.Company),
			City:        pandoJob.City,
			CPA:         pandoJob.CPA,
			CPC:         pandoJob.CPC,
			Country:     pandoJob.Country,
			ZIP:         pandoJob.ZIP,
			State:       pandoJob.State,
		}
		jobs = append(jobs, sj)
	}

	return &jobs, nil
}

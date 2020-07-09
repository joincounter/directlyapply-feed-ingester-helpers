package helpers

import (
	"bytes"
	"encoding/json"
	"net/http"
)

// SendToScraper is the stand way to send a job to the scrapper
func SendToScraper(scrapperEndpoint string, job StandardJob) (*http.Response, error) {
	jsonValue, err := json.Marshal(job)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest("POST", scrapperEndpoint, bytes.NewBuffer(jsonValue))
	if err != nil {
		return nil, err
	}
	req.Header.Add("content-type", "application/json")
	req.Header.Add("cache-control", "no-cache")
	return http.DefaultClient.Do(req)
}

package helpers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

var countryMap = map[string]string{
	"US": "us-directlyapply-indexer.herokuapp.com",
	"UK": "directlyapply-indexer.herokuapp.com",
	"CA": "ca-directlyapply-indexer.herokuapp.com",
}

func SendToIndexerUpdate(countryCode string, job LegacyUploadJob) (*http.Response, error) {
	indexexURL := fmt.Sprintf("https://%s/v2", countryMap[countryCode])
	buf := new(bytes.Buffer)
	json.NewEncoder(buf).Encode(job)
	req, _ := http.NewRequest("POST", indexexURL, buf)
	req.Header.Add("content-type", "application/json")
	return http.DefaultClient.Do(req)
}

func SendToIndexerNew(countryCode string, job LegacyUploadJob) (*http.Response, error) {
	indexexURL := fmt.Sprintf("https://%s/v1", countryMap[countryCode])
	buf := new(bytes.Buffer)
	json.NewEncoder(buf).Encode(job)
	req, _ := http.NewRequest("POST", indexexURL, buf)
	req.Header.Add("content-type", "application/json")
	return http.DefaultClient.Do(req)
}

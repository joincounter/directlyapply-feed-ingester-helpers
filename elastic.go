package helpers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

// RemoveFromElastic remove list of IDs from elastic
func RemoveFromElastic(jobIDs []string, countryCode string, token string) (*http.Response, error) {
	type deletionStruct struct {
		Query struct {
			Terms struct {
				ID []string `json:"_id"`
			} `json:"terms"`
		} `json:"query"`
	}
	var deletion deletionStruct
	deletion.Query.Terms.ID = jobIDs
	var jsonData []byte
	jsonData, err := json.Marshal(deletion)
	if err != nil {
		return nil, err
	}
	payload := strings.NewReader(string(jsonData))
	req, err := http.NewRequest("POST", fmt.Sprintf("https://76b76771cccc46ca8ef47110d5142527.eu-west-1.aws.found.io:9243/%s/_delete_by_query", elasticDBs[countryCode]), payload)
	req.Header.Add("Content-Type", "application/json")
	if err != nil {
		return nil, err
	}
	req.Header.Add("authorization", fmt.Sprintf("Basic %s", token))
	return http.DefaultClient.Do(req)
}

func PushToElastic(job LegacyUploadJob, countryCode string, token string) (*http.Response, error) {
	if countryCode == "UK" || countryCode == "CA" {
		job.State = ""
	}
	jsonData, err := json.Marshal(job)
	if err != nil {
		return nil, err
	}
	payload := strings.NewReader(string(jsonData))
	req, err := http.NewRequest("POST", fmt.Sprintf("https://76b76771cccc46ca8ef47110d5142527.eu-west-1.aws.found.io:9243/%s/_doc/%s", elasticDBs[countryCode], job.ID.Hex()), payload)
	req.Header.Add("Content-Type", "application/json")
	if err != nil {
		return nil, err
	}
	req.Header.Add("authorization", fmt.Sprintf("Basic %s", token))
	return http.DefaultClient.Do(req)
}

var elasticDBs = map[string]string{
	"US": "jobs_us",
	"UK": "jobs",
	"CA": "jobs_ca",
}

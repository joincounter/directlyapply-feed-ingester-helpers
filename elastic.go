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
	req, err := http.NewRequest("POST", elasticURLs[countryCode], payload)
	req.Header.Add("Content-Type", "application/json")
	if err != nil {
		return nil, err
	}
	req.Header.Add("authorization", fmt.Sprintf("Basic %s", token))
	return http.DefaultClient.Do(req)
}

var elasticURLs = map[string]string{
	"US": "https://76b76771cccc46ca8ef47110d5142527.eu-west-1.aws.found.io:9243/jobs_us/_delete_by_query",
	"UK": "https://76b76771cccc46ca8ef47110d5142527.eu-west-1.aws.found.io:9243/jobs/_delete_by_query",
	"CA": "https://76b76771cccc46ca8ef47110d5142527.eu-west-1.aws.found.io:9243/jobs_ca/_delete_by_query",
}

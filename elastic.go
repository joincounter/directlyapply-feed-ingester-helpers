package helpers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

// RemoveFromElastic remove list of IDs from elastic
func RemoveFromElastic(jobIDs []string, hostname string, token string) (*http.Response, error) {
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
	url := fmt.Sprintf("https://%s/jobs_ca/_delete_by_query", hostname)
	req, err := http.NewRequest("POST", url, payload)
	req.Header.Add("Content-Type", "application/json")
	if err != nil {
		return nil, err
	}
	req.Header.Add("authorization", fmt.Sprintf("Basic %s", token))
	return http.DefaultClient.Do(req)
}

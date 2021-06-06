package helpers

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"mime/multipart"
	"net/http"
)

func FormatJobDescription(jobDescription string, endpoint string) (string, error) {
	payload := &bytes.Buffer{}
	writer := multipart.NewWriter(payload)
	err := writer.WriteField("jd", jobDescription)
	if err != nil {
		return jobDescription, err
	}
	err = writer.Close()
	if err != nil {
		return jobDescription, err
	}

	client := &http.Client{}
	req, err := http.NewRequest("POST", endpoint, payload)

	if err != nil {
		return jobDescription, err
	}
	req.Header.Set("Content-Type", writer.FormDataContentType())
	res, err := client.Do(req)
	if err != nil {
		return jobDescription, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return jobDescription, err
	}
	return string(body), nil
}

func GetSummaries(jobDescription string, endpoint string) (*[]string, error) {
	payload := &bytes.Buffer{}
	writer := multipart.NewWriter(payload)
	err := writer.WriteField("jd", jobDescription)
	if err != nil {
		return nil, err
	}
	err = writer.WriteField("format", "")
	if err != nil {
		return nil, err
	}
	err = writer.Close()
	if err != nil {
		return nil, err
	}

	client := &http.Client{}
	req, err := http.NewRequest("POST", endpoint, payload)

	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", writer.FormDataContentType())
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var summaries []string

	json.Unmarshal(body, &summaries)

	return &summaries, nil
}

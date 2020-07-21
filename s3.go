package helpers

import (
	"net/http"
)

// FetchFeedLastModified will return the upload time of a public s3 file
func FetchFeedLastModified(url string) (*string, error) {
	res, err := http.Head(url)
	if err != nil {
		return nil, err
	}
	lastModified := res.Header.Get("last-modified")
	return &lastModified, nil
}

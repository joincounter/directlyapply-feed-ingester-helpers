package helpers

import (
	"compress/gzip"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/google/uuid"
)

// DownloadToDisk will download straight to disk
func DownloadToDisk(URL string) (*string, error) {
	os.MkdirAll("./downloadedFiles", os.ModePerm)
	fileName := fmt.Sprintf("downloadedFiles/notzippedfile_%s.xml", uuid.New().String())

	req, requestErr := http.NewRequest("GET", URL, nil)
	if requestErr != nil {
		return nil, requestErr
	}

	res, excutionError := http.DefaultClient.Do(req)
	if excutionError != nil {
		return nil, excutionError
	}
	defer res.Body.Close()

	file, fileOpenError := os.OpenFile(fileName, os.O_RDWR|os.O_CREATE, 0755)

	defer file.Close()

	if fileOpenError != nil {
		return nil, fileOpenError
	}

	_, writeErr := io.Copy(file, res.Body)

	if writeErr != nil {
		return nil, writeErr
	}

	return &fileName, nil
}

// DownloadAndUnzipToDisk will download and uncompress the feed and save to file
func DownloadAndUnzipToDisk(URL string) (*string, error) {

	URLSplit := strings.Split(URL, "/")

	if len(URLSplit) == 1 {
		return nil, errors.New("URL not in correct format")
	}

	fileName := URLSplit[len(URLSplit)-1]

	if strings.HasPrefix(fileName, ".gz") {
		return nil, errors.New("file not a gzip")
	}

	os.MkdirAll("./downloadedFiles", os.ModePerm)
	fileName = "downloadedFiles/" + uuid.New().String() + "_" + strings.TrimRight(fileName, ".gz")

	req, requestErr := http.NewRequest("GET", URL, nil)
	if requestErr != nil {
		return nil, requestErr
	}

	req.Header.Add("Accept-Encoding", "gzip")

	res, excutionError := http.DefaultClient.Do(req)
	if excutionError != nil {
		return nil, excutionError
	}
	defer res.Body.Close()

	// Decompress the GZIP
	gz, err := gzip.NewReader(res.Body)

	if err != nil {
		log.Fatal(err)
	}

	defer gz.Close()

	file, fileOpenError := os.OpenFile(fileName, os.O_RDWR|os.O_CREATE, 0755)

	defer file.Close()

	if fileOpenError != nil {
		return nil, fileOpenError
	}

	_, writeErr := io.Copy(file, gz)

	if writeErr != nil {
		return nil, writeErr
	}

	return &fileName, nil
}

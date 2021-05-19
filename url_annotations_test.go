package helpers

import (
	"fmt"
	"testing"
)

func TestAppendDirectlyApplySourceTagToStandardJobs(t *testing.T) {
	feedname := "directlyapply-uk"
	jobs := []StandardJob{
		{
			URL: "www.neuvoo.co.uk/jobs/123456789",
		},
		{
			URL: "www.neuvoo.co.uk",
		},
	}
	t.Run("Will correctly anotate URL", func(t *testing.T) {
		expectedURL := "www.neuvoo.co.uk/jobs/123456789?da_src=directlyapply-uk"
		output := AppendDirectlyApplySourceTagToStandardJobs(jobs, feedname)
		if output[0].URL != expectedURL {
			t.Errorf("Expected %s but got %s", expectedURL, output[0].URL)
		}
	})
}

func TestAppendDirectlyApplySourceTag(t *testing.T) {
	feedname := "directlyapply-uk"
	t.Run("url with path returns correct url annotation", func(t *testing.T) {
		urlToTest := "www.neuvoo.co.uk/jobs/123456789"
		expectedURL := fmt.Sprintf("www.neuvoo.co.uk/jobs/123456789?da_src=%s", feedname)
		output := AppendDirectlyApplySourceTag(urlToTest, feedname)
		if output != expectedURL {
			t.Errorf("Expected %s but got %s", expectedURL, output)
		}
	})
	t.Run("url no path returns correct url annotation", func(t *testing.T) {
		urlToTest := "www.neuvoo.co.uk"
		expectedURL := fmt.Sprintf("www.neuvoo.co.uk?da_src=%s", feedname)
		output := AppendDirectlyApplySourceTag(urlToTest, feedname)
		if output != expectedURL {
			t.Errorf("Expected %s but got %s", expectedURL, output)
		}
	})
}

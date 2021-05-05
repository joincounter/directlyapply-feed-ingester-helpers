package helpers

import (
	"math"
)

// StringInSlice figures out if string is in slice
func StringInSlice(str string, list []string) bool {
	for _, v := range list {
		if v == str {
			return true
		}
	}
	return false
}

func BatchStandardJobs(jobs []StandardJob, batchSize int) [][]StandardJob {
	batchedJobs := make([][]StandardJob, 0)
	var batchMarker int
	for {
		jobsInThisBatch := int(math.Min(float64(len(jobs)-batchMarker), float64(batchSize)))
		batchedJobs = append(batchedJobs, jobs[batchMarker:batchMarker+jobsInThisBatch])
		batchMarker += batchSize
		if jobsInThisBatch < batchSize {
			break
		}
	}
	return batchedJobs
}

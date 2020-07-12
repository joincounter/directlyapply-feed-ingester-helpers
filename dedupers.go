package helpers

import "strings"

// MainDeduper standared way to remove duplicates
func MainDeduper(jobs []StandardJob) ([]StandardJob, int) {
	jobsMap := make(map[string]StandardJob)
	jobsRemoved := 0
	jobsSlice := make([]StandardJob, 0)
	for _, job := range jobs {
		slug := strings.Join([]string{job.Title, job.Company, job.City}, "|")
		// TODO: I need to check CPAs and CPCs here
		if _, ok := jobsMap[slug]; ok {
			jobsRemoved++
		}
		jobsMap[slug] = job
	}
	for _, job := range jobsMap {
		jobsSlice = append(jobsSlice, job)
	}
	return jobsSlice, jobsRemoved
}

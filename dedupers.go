package helpers

import "strings"

// MainDeduper standared way to remove duplicates
func MainDeduper(jobs []StandardJob) []StandardJob {
	jobsMap := make(map[string]StandardJob)
	jobsSlice := make([]StandardJob, 0)
	for _, job := range jobs {
		slug := strings.Join([]string{job.Title, job.Company, job.City}, "|")
		// TODO: I need to check CPAs and CPCs here
		jobsMap[slug] = job
	}
	for _, job := range jobsMap {
		jobsSlice = append(jobsSlice, job)
	}
	return jobsSlice
}

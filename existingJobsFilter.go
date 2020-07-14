package helpers

// ExistingJobFilter a helper to remove existing jobs
func ExistingJobFilter(jobs []StandardJob, existingJobs []PersistentIndexJob) ([]StandardJob, int) {
	jobsRemovedCount := 0
	jobsNotInDB := make([]StandardJob, 0)
UpperLoop:
	for i := 0; i < len(jobs); i++ {
		for j := 0; j < len(existingJobs); j++ {
			url := existingJobs[j].URL
			if jobs[i].URL == url {
				jobsRemovedCount++
				break UpperLoop
			}
		}
		jobsNotInDB = append(jobsNotInDB, jobs[i])
	}
	return jobsNotInDB, jobsRemovedCount
}

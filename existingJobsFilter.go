package helpers

// ExistingJobFilter a helper to remove existing jobs
func ExistingJobFilter(jobs []StandardJob, existingJobs []PersistentIndexJob) ([]StandardJob, []string, int) {
	jobsRemovedCount := 0
	jobsNotInDB := make([]StandardJob, 0)
	existingJobsMap := make(map[PersistentIndexJob]bool)
	for i := 0; i < len(existingJobs); i++ {
		existingJobsMap[existingJobs[i]] = false
	}
UpperLoop:
	for i := 0; i < len(jobs); i++ {
		for job := range existingJobsMap {
			url := job.URL
			if jobs[i].URL == url {
				jobsRemovedCount++
				existingJobsMap[job] = true
				break UpperLoop
			}
		}
		jobsNotInDB = append(jobsNotInDB, jobs[i])
	}
	jobIDsNotInFeed := make([]string, 0)
	for job, inFeed := range existingJobsMap {
		if !inFeed {
			jobIDsNotInFeed = append(jobIDsNotInFeed, job.ID)
		}
	}
	return jobsNotInDB, jobIDsNotInFeed, jobsRemovedCount
}

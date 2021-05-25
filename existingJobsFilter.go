package helpers

// ExistingJobFilter a helper to remove existing jobs
func ExistingJobFilter(jobs []StandardJob, existingJobs []PersistentIndexJob) ([]StandardJob, []string, int) {
	existingJobsURLMap := make(map[string]string)
	for i := 0; i < len(existingJobs); i++ {
		existingJobsURLMap[existingJobs[i].URL] = existingJobs[i].ID
	}
	feedJobsURLMap := make(map[string]StandardJob)
	for i := 0; i < len(jobs); i++ {
		feedJobsURLMap[jobs[i].URL] = jobs[i]
	}

	jobsNotInDB := make([]StandardJob, 0)
	for i := 0; i < len(jobs); i++ {
		if _, inDatabase := existingJobsURLMap[jobs[i].URL]; !inDatabase {
			jobsNotInDB = append(jobsNotInDB, jobs[i])
		}
	}

	jobIDsNotInFeed := make([]string, 0)
	for i := 0; i < len(existingJobs); i++ {
		if _, inFeed := feedJobsURLMap[existingJobs[i].URL]; !inFeed {
			jobIDsNotInFeed = append(jobIDsNotInFeed, existingJobs[i].ID)
		}
	}

	jobsRemovedCount := len(jobs) - len(jobsNotInDB)
	return jobsNotInDB, jobIDsNotInFeed, jobsRemovedCount
}

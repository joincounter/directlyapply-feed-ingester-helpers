package helpers

// NoCityFilter remove jobs that don't have a city
func NoCityFilter(job *StandardJob) *StandardJob {
	if job.City == "" {
		return nil
	}
	return job
}

// NoJobTitleFilter remove jobs that don't have a title
func NoJobTitleFilter(job *StandardJob) *StandardJob {
	if job.Title == "" {
		return nil
	}
	return job
}

// EvaluateFilter evaluate filter against a slice of jobs
func EvaluateFilter(jobs []StandardJob, filter func(job *StandardJob) *StandardJob) []StandardJob {
	evaluatedJobs := make([]StandardJob, 0)
	for i := 0; i < len(jobs); i++ {
		job := filter(&jobs[i])
		if job != nil {
			evaluatedJobs = append(evaluatedJobs, *job)
		}
	}
	return evaluatedJobs
}

package helpers

import "net/url"

func AppendDirectlyApplySourceTagToStandardJobs(jobs []StandardJob, feedname string) []StandardJob {
	for i := 0; i < len(jobs); i++ {
		jobs[i].URL = AppendDirectlyApplySourceTag(jobs[i].URL, feedname)
	}
	return jobs
}

func AppendDirectlyApplySourceTag(urlStr, feedname string) string {
	urlF, err := url.Parse(urlStr)
	if err != nil {
		println(err.Error())
		return urlStr
	}
	x := urlF.Query()
	x.Add("da_src", feedname)
	urlF.RawQuery = x.Encode()
	return urlF.String()
}

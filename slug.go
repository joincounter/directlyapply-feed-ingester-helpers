package helpers

import (
	"fmt"
	"strings"
)

// JobTitleSlug mcleans up and normalizes Job Title
func JobTitleSlug(jobTitle string) string {
	jobTitle = strings.Split(jobTitle, "--")[0]
	jobTitle = strings.Replace(jobTitle, " ", "-", -1)
	jobTitle = strings.Trim(jobTitle, "-")
	if len(jobTitle) > 49 {
		return jobTitle[:50]
	}
	return jobTitle
}

// LocationSlug makes a unique slug from a region and a city
func LocationSlug(region string, city string) string {
	city = strings.ToLower(city)
	city = strings.Replace(city, " ", "-", -1)
	if region == "" {
		return city
	}
	region = strings.ToLower(region)
	region = strings.Replace(region, " ", "-", -1)
	region = fmt.Sprintf("%s//%s", region, city)
	if len(region) > 49 {
		return region[:50]
	}
	return region
}

// GenerateSlug is a standard slug generating function
func GenerateSlug(companyName string) string {
	slug := strings.ToLower(companyName)
	slug = strings.Replace(slug, "ltd", "", -1)
	slug = strings.Replace(slug, "limited", "", -1)
	slug = strings.Replace(slug, "llp", "", -1)
	slug = strings.Replace(slug, "plc", "", -1)
	slug = strings.Replace(slug, ".", "", -1)
	slug = strings.Replace(slug, "inc", "", -1)
	slug = strings.Replace(slug, "|", "", -1)
	slug = strings.Replace(slug, "(", "", -1)
	slug = strings.Replace(slug, ")", "", -1)
	slug = strings.Split(slug, ",")[0]
	slug = strings.Replace(slug, " ", "-", -1)
	slug = strings.Replace(slug, "--", "-", -1)
	slug = strings.Replace(slug, "'", "", -1)
	slug = strings.Replace(slug, "/", "", -1)
	slug = strings.Replace(slug, `\`, "", -1)
	slug = strings.Replace(slug, "!", "", -1)
	slug = strings.Replace(slug, "&#39;", "", -1)
	slug = strings.Replace(slug, "®", "", -1)
	slug = strings.TrimRight(slug, "-")
	slug = strings.TrimLeft(slug, "-")
	if len(slug) > 49 {
		return slug[:50]
	}
	return slug
}

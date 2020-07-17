package helpers

import (
	"strings"
)

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
	return slug
}

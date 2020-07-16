package helpers

import (
	"strings"
)

func GenerateSlug(companyName string) string {
	lowercase := strings.ToLower(companyName)
	removeLtd := strings.Replace(lowercase, "ltd", "", -1)
	removeLimited := strings.Replace(removeLtd, "limited", "", -1)
	removeLlp := strings.Replace(removeLimited, "llp", "", -1)
	removePlc := strings.Replace(removeLlp, "plc", "", -1)
	removeDots := strings.Replace(removePlc, ".", "", -1)
	removeInc := strings.Replace(removeDots, "inc", "", -1)
	removeLine := strings.Replace(removeInc, "|", "", -1)
	removeLeftBracket := strings.Replace(removeLine, "(", "", -1)
	removeRightBracket := strings.Replace(removeLeftBracket, ")", "", -1)
	splitOnComma := strings.Split(removeRightBracket, ",")[0]
	slug := strings.Replace(splitOnComma, " ", "-", -1)
	removeDoubleDash := strings.Replace(slug, "--", "-", -1)
	removeApostrophe := strings.Replace(removeDoubleDash, "'", "", -1)
	removeSlash1 := strings.Replace(removeApostrophe, "/", "", -1)
	removeSlash2 := strings.Replace(removeSlash1, `\`, "", -1)
	removeExclamation := strings.Replace(removeSlash2, "!", "", -1)
	removeHtml := strings.Replace(removeExclamation, "&#39;", "", -1)
	removeRights := strings.Replace(removeHtml, "®", "", -1)
	slugCleanedRight := strings.TrimRight(removeRights, "-")
	slugCleanedLeft := strings.TrimLeft(slugCleanedRight, "-")
	return slugCleanedLeft
}
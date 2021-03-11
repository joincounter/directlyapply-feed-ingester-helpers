package helpers

import (
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"
)

func ExtractUSDSalaries(description string) SalaryData {
	r, err := regexp.Compile("([£$€][0-9][0-9],[0-9][0-9][0-9]|[£$€][0-9][0-9][0-9][0-9][0-9]|[£$€][0-9][0-9][Kk])")
	if err != nil {
		log.Fatalln("huge error, unrecoverable")
	}

	annualMatches := r.FindAllString(description, -1)
	if len(annualMatches) > 0 {
		curr := annualMatches[0][0]
		annualMatchesNormalised := make([]float64, 0)
		for i := 0; i < len(annualMatches); i++ {
			norm, err := normaliseSalaries(annualMatches[i])
			if err != nil {
				continue
			}
			annualMatchesNormalised = append(annualMatchesNormalised, norm)
		}
		// calculate the average 🙄
		var total float64
		for _, v := range annualMatchesNormalised {
			total += v
		}
		averageAnnualSalary := total / float64(len(annualMatchesNormalised))
		return annualSalaryData{annualRate: averageAnnualSalary, currency: string(curr)}
	}

	return emptySalaryData{}
}

func normaliseSalaries(extracted string) (float64, error) {
	extracted = strings.Replace(extracted, "£", "", 1)
	extracted = strings.Replace(extracted, "$", "", 1)
	extracted = strings.Replace(extracted, "€", "", 1)
	extracted = strings.Replace(extracted, ",", "", 1)
	extracted = strings.Replace(extracted, "K", "000", 1)
	extracted = strings.Replace(extracted, "k", "000", 1)
	return strconv.ParseFloat(extracted, 64)
}

type emptySalaryData struct{}

func (esd emptySalaryData) Annual() float64 {
	return 0
}

func (esd emptySalaryData) Hourly() float64 {
	return 0
}

func (esd emptySalaryData) asString() string {
	return "no salary data"
}

func (esd emptySalaryData) getCurrency() string {
	return "None"
}

type annualSalaryData struct {
	annualRate float64
	currency   string
}

func (asd annualSalaryData) Annual() float64 {
	return asd.annualRate
}

func (asd annualSalaryData) Hourly() float64 {
	return asd.annualRate / 1950
}

func (asd annualSalaryData) asString() string {
	return fmt.Sprintf("%s%f per year", asd.currency, asd.Annual())
}

func (asd annualSalaryData) getCurrency() string {
	return asd.currency
}

type hourlySalaryData struct {
	HourlyRate float64
	currency   string
}

func (hsd hourlySalaryData) Annual() float64 {
	return hsd.HourlyRate * 1950
}

func (hsd hourlySalaryData) Hourly() float64 {
	return hsd.HourlyRate
}

func (hsd hourlySalaryData) asString() string {
	return fmt.Sprintf("%s%f per year", hsd.currency, hsd.Annual())
}

func (hsd hourlySalaryData) getCurrency() string {
	return hsd.currency
}

// SalaryData is a representation of a job descriptions salary data
type SalaryData interface {
	Annual() float64
	Hourly() float64
	asString() string
	getCurrency() string
}

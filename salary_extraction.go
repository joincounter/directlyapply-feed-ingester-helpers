package helpers

import (
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"
)

// ExtractUSDSalaries get salary data from job description
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

	wageRedux, err := regexp.Compile(`([£$€][0-9]?[0-9]\.[0-9][0-9]|[£$€][0-9]?[0-9]\s)`)
	if err != nil {
		log.Fatalln("huge error, unrecoverable")
	}

	wageMatches := wageRedux.FindAllString(description, -1)

	if len(wageMatches) > 0 {
		curr := wageMatches[0][0]
		if len(wageMatches) == 1 {
			norm, err := normaliseSalaries(wageMatches[0])
			if err != nil {
				return emptySalaryData{}
			}
			return hourlySalaryData{HourlyRate: norm, currency: string(curr)}
		}
		if len(wageMatches) == 2 {
			wageMin, err := normaliseSalaries(wageMatches[0])
			wageMax, err := normaliseSalaries(wageMatches[1])
			if err != nil {
				return emptySalaryData{}
			}
			return hourlyRangeSalaryData{lowerHourlyRate: wageMin, higherHourlyRate: wageMax, currency: string(curr)}
		}
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

func (esd emptySalaryData) SalaryType() string {
	return "NONE"
}

func (esd emptySalaryData) SalaryMin() float64 {
	return 0
}

func (esd emptySalaryData) SalaryMax() float64 {
	return 0
}

func (esd emptySalaryData) Annual() float64 {
	return 0
}

func (esd emptySalaryData) Hourly() float64 {
	return 0
}

func (esd emptySalaryData) String() string {
	return "no salary data"
}

func (esd emptySalaryData) GetCurrency() string {
	return "None"
}

type annualSalaryData struct {
	annualRate float64
	currency   string
}

func (asd annualSalaryData) SalaryType() string {
	return "ANNUAL"
}

func (asd annualSalaryData) SalaryMin() float64 {
	return asd.annualRate
}

func (asd annualSalaryData) SalaryMax() float64 {
	return asd.annualRate
}

func (asd annualSalaryData) Annual() float64 {
	return asd.annualRate
}

func (asd annualSalaryData) Hourly() float64 {
	return asd.annualRate / 1950
}

func (asd annualSalaryData) String() string {
	return fmt.Sprintf("%s%f per year", asd.currency, asd.Annual())
}

func (asd annualSalaryData) GetCurrency() string {
	return asd.currency
}

type hourlySalaryData struct {
	HourlyRate float64
	currency   string
}

func (hsd hourlySalaryData) SalaryType() string {
	return "HOURLY"
}

func (hsd hourlySalaryData) SalaryMin() float64 {
	return hsd.HourlyRate
}

func (hsd hourlySalaryData) SalaryMax() float64 {
	return hsd.HourlyRate
}

func (hsd hourlySalaryData) Annual() float64 {
	return hsd.HourlyRate * 1950
}

func (hsd hourlySalaryData) Hourly() float64 {
	return hsd.HourlyRate
}

func (hsd hourlySalaryData) String() string {
	return fmt.Sprintf("%s%f per hour", hsd.currency, hsd.HourlyRate)
}

func (hsd hourlySalaryData) GetCurrency() string {
	return hsd.currency
}

type hourlyRangeSalaryData struct {
	lowerHourlyRate  float64
	higherHourlyRate float64
	currency         string
}

func (hsd hourlyRangeSalaryData) SalaryType() string {
	return "HOURLY"
}

func (hsd hourlyRangeSalaryData) SalaryMin() float64 {
	return hsd.lowerHourlyRate
}

func (hsd hourlyRangeSalaryData) SalaryMax() float64 {
	return hsd.higherHourlyRate
}

func (hsd hourlyRangeSalaryData) meanWage() float64 {
	return (hsd.lowerHourlyRate + hsd.higherHourlyRate) / 2
}

func (hsd hourlyRangeSalaryData) Annual() float64 {
	return hsd.meanWage() * 1950
}

func (hsd hourlyRangeSalaryData) Hourly() float64 {
	return hsd.meanWage()
}

func (hsd hourlyRangeSalaryData) String() string {
	return fmt.Sprintf("Between %s%f and %s%f per hour", hsd.currency, hsd.lowerHourlyRate, hsd.currency, hsd.higherHourlyRate)
}

func (hsd hourlyRangeSalaryData) GetCurrency() string {
	return hsd.currency
}

// SalaryData is a representation of a job descriptions salary data
type SalaryData interface {
	Annual() float64
	Hourly() float64
	String() string
	GetCurrency() string
	SalaryType() string
	SalaryMin() float64
	SalaryMax() float64
}

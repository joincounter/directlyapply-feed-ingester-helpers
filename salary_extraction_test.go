package helpers

import "testing"

func TestExtractUSDSalaries(t *testing.T) {
	t.Run("basic test", func(t *testing.T) {
		jobDescription := "Hello, great job, £40,000"
		expected := 40000.0
		output := ExtractUSDSalaries(jobDescription).Annual()
		if output != expected {
			t.Errorf("Expected %f but got %f", expected, output)
		}
	})
	t.Run("wage test", func(t *testing.T) {
		jobDescription := "Hello, great job, £5.56 per hour"
		expected := 5.56
		output := ExtractUSDSalaries(jobDescription).Hourly()
		if output != expected {
			t.Errorf("Expected %f but got %f", expected, output)
		}
	})
	t.Run("wage test, range", func(t *testing.T) {
		jobDescription := "Hello, great job, €12.50 - €13.50 per hour"
		expected := 13.0
		output := ExtractUSDSalaries(jobDescription).Hourly()
		if output != expected {
			t.Errorf("Expected %f but got %f", expected, output)
		}
	})
}

func TestNormaliseSalaries(t *testing.T) {
	examples := []struct {
		name   string
		input  string
		output float64
	}{
		{
			name:   "Test full thing",
			input:  "£30,405.02",
			output: 30405.02,
		},
		{
			name:   "Test shortend thing",
			input:  "£30,475",
			output: 30475,
		},
		{
			name:   "Test euros",
			input:  "€30,405",
			output: 30405,
		},
		{
			name:   "Test dollars",
			input:  "$30,405",
			output: 30405,
		},
		{
			name:   "Test small",
			input:  "$30.34",
			output: 30.34,
		},
	}
	for i := 0; i < len(examples); i++ {
		t.Run(examples[i].name, func(t *testing.T) {
			calculated, err := normaliseSalaries(examples[i].input)
			if err != nil || calculated != examples[i].output {
				t.Errorf("Expected %f but got %f", examples[i].output, calculated)
			}
		})
	}
}

func TextMaxMinFloat64(t *testing.T) {
	examples := []struct{
		name string
		values []float64
		min float64
		max float64
	}{
		{
			name: "nothing",
			values: []float64{},
			min: 0,
			max: 0,
		},
		{
			name: "one value",
			values: []float64{5.67},
			min: 5.67,
			max: 5.67,
		},
		{
			name: "two values",
			values: []float64{7.78, 5.67},
			min: 5.67,
			max: 7.78,
		},
		{
			name: "multiple values",
			values: []float64{7.78, 5.67, 6.76, 11.34},
			min: 5.67,
			max: 11.34,
		},
	}
	for i := 0; i < len(examples); i++ {
		t.Run(examples[i].name, func(t *testing.T) {
			min, max := maxMinFloat64(examples[i].values...)
			if examples[i].min != min {
				t.Errorf("Expected min value of %f but got %f", examples[i].min, min)
			}
			if examples[i].max != max {
				t.Errorf("Expected max value of %f but got %f", examples[i].max, max)
			}
		})
	}
}

package utils

import "testing"

type scenario struct {
	zipCode  string
	expected bool
}

func TestIsValidZipCode(t *testing.T) {
	testScenarios := []scenario{
		{zipCode: "01451918", expected: true},
		{zipCode: "0145191", expected: false},
		{zipCode: "014519180", expected: false},
	}

	for i := range testScenarios {
		got := IsValidZipCode(testScenarios[i].zipCode)
		if got != testScenarios[i].expected {
			t.Fatalf("Expected %t, got %t", testScenarios[i].expected, got)
		}
	}
}

func TestConvertCelciusToFahrenheit(t *testing.T) {
	temperatureInCelcius := 32.0
	expected := 89.60
	got := ConvertCelciusToFahrenheit(temperatureInCelcius)

	if got != expected {
		t.Fatalf("Expected %f, got %f", expected, got)
	}
}

func TestConvertCelciusToKelvin(t *testing.T) {
	temperatureInCelcius := 32.0
	expected := 305.0
	got := ConvertCelciusToKelvin(temperatureInCelcius)

	if got != expected {
		t.Fatalf("Expected %f, got %f", expected, got)
	}
}

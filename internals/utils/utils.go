package utils

import (
	"encoding/json"
	"fmt"
	"regexp"
	"strings"
)

func ConvertCelciusToKelvin(celciusValue float64) float64 {
	return celciusValue + 273.0
}

func ConvertCelciusToFahrenheit(celciusValue float64) float64 {
	return (celciusValue * 1.8) + 32.0
}

func IsValidZipCode(zipCode string) bool {
	zipCode = strings.ReplaceAll(zipCode, "-", "")
	re := regexp.MustCompile(`^\d{8}$`)

	return re.MatchString(zipCode)
}

func BuildJsonMessage(message string) ([]byte, error) {
	preMessage := fmt.Sprintf("Message: %s", message)
	return json.Marshal(preMessage)
}

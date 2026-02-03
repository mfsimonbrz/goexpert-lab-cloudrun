package main

import (
	"net/http"
	"os"

	"github.com/mfsimonbrz/goexpert-lab-cloudrun/internals/handlers"
)

var weatherInfo *handlers.WeatherInfo

func init() {
	apiToken := os.Getenv("API_TOKEN")
	if apiToken == "" {
		panic("Ennvironment variable API_TOKEN not found!")
	}

	weatherInfo = handlers.NewWeatherInfo(apiToken)
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /{zipCode}", weatherInfo.TemperatureHandler)

	http.ListenAndServe(":8080", mux)
}

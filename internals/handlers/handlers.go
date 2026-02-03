package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	"github.com/mfsimonbrz/goexpert-lab-cloudrun/internals/business"
	"github.com/mfsimonbrz/goexpert-lab-cloudrun/internals/models"
	"github.com/mfsimonbrz/goexpert-lab-cloudrun/internals/utils"
)

type WeatherInfo struct {
	businessInfo business.BusinessInfo
}

func NewWeatherInfo(apiToken string) *WeatherInfo {
	return &WeatherInfo{businessInfo: business.BusinessInfo{APIToken: apiToken}}
}

func (wi *WeatherInfo) TemperatureHandler(w http.ResponseWriter, r *http.Request) {
	zipCode := r.PathValue("zipCode")
	if !utils.IsValidZipCode(zipCode) {
		message, _ := utils.BuildJsonMessage(fmt.Sprintf("invalid zip code: %s", zipCode))
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnprocessableEntity)
		w.Write(message)
		return
	}

	endereco, err := wi.businessInfo.GetAddressInformation(zipCode)
	if err != nil {
		message, _ := utils.BuildJsonMessage(fmt.Sprintf("Can not find zip code: %s", err.Error()))
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		w.Write(message)
		return
	}

	cidade := url.QueryEscape(endereco.Localidade)
	weatherInfo, err := wi.businessInfo.GetWeatherInformation(cidade)
	if err != nil {
		message, _ := utils.BuildJsonMessage(err.Error())
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(message)
		return
	}

	tempF := utils.ConvertCelciusToFahrenheit(weatherInfo.Current.TempC)
	tempK := utils.ConvertCelciusToKelvin(weatherInfo.Current.TempC)
	result := models.Result{TemperatureCelcius: weatherInfo.Current.TempC, TemperatureFahrenheit: tempF, TemperatureKelvin: tempK}

	resultInBytes, err := json.Marshal(result)
	if err != nil {
		message, _ := utils.BuildJsonMessage(err.Error())
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(message)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(resultInBytes)
}

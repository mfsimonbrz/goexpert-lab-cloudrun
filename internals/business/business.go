package business

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/mfsimonbrz/goexpert-lab-cloudrun/internals/models"
)

const zipCodeAPI = "http://viacep.com.br/ws/%s/json"
const weatherAPI = "http://api.weatherapi.com/v1/current.json?key=%s&q=%s&aqi=no"

type BusinessInfo struct {
	APIToken string
}

func (b *BusinessInfo) GetWeatherInformation(cityName string) (*models.Weather, error) {
	url := fmt.Sprintf(weatherAPI, b.APIToken, cityName)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	bodyBytes, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	var weather models.Weather
	err = json.Unmarshal(bodyBytes, &weather)
	if err != nil {
		return nil, err
	}

	return &weather, nil
}

func (b *BusinessInfo) GetAddressInformation(zipCode string) (*models.Endereco, error) {
	url := fmt.Sprintf(zipCodeAPI, zipCode)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	bodyBytes, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	var endereco models.Endereco
	json.Unmarshal(bodyBytes, &endereco)

	if endereco.CEP == "" {
		return nil, fmt.Errorf("zip code %s cannot be found", zipCode)
	}

	return &endereco, nil
}

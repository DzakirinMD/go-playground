package api

import (
	"encoding/json"
	"go_playground/mini-project/weather-tracker/weatherstructs"
	"io/ioutil"
	"log"
	"net/http"
)

func Query(city string) (weatherstructs.WeatherData, error) {
	apiConfig, err := loadApiConfig("mini-project/weather-tracker/.apiConfig")
	if err != nil {
		return weatherstructs.WeatherData{}, err
	}

	resp, err := http.Get("http://api.openweathermap.org/data/2.5/weather?APPID=" + apiConfig.OpenWeatherMapApiKey + "&q=" + city)
	if err != nil {
		return weatherstructs.WeatherData{}, err
	}

	log.Println("log Respond from weather api : ", resp)

	defer resp.Body.Close()

	var d weatherstructs.WeatherData
	if err := json.NewDecoder(resp.Body).Decode(&d); err != nil {
		return weatherstructs.WeatherData{}, err
	}
	return d, nil
}

func loadApiConfig(fileName string) (weatherstructs.ApiConfigData, error) {
	bytes, err := ioutil.ReadFile(fileName)

	if err != nil {
		return weatherstructs.ApiConfigData{}, err
	}

	var apikeys weatherstructs.ApiConfigData

	// transforming an object into a specific data format that is suitable for transmission.
	err = json.Unmarshal(bytes, &apikeys)
	if err != nil {
		return weatherstructs.ApiConfigData{}, err
	}
	return apikeys, nil
}

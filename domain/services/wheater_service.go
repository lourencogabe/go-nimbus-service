package services

import (
	"encoding/json"
	"fmt"
	"net/http"

	domain "github.com/lourencogabe/gonimbus/domain/entities"
	"github.com/lourencogabe/gonimbus/dto"
)

// método get realizara a consulta
func GetWeatherData(city string) dto.GetWeather {
	url := fmt.Sprintf("https://wttr.in/%s?format=j1", city)

	resp, err := http.Get(url)

	if err != nil {
		panic(err)
	}

	var wheather dto.GetWeather
	err = json.NewDecoder(resp.Body).Decode(&wheather)

	if err != nil {
		panic(err)
	}

	return wheather
}

func WeatherDataFormating(city string) domain.Weather {
	resp := GetWeatherData(city)

	formattedWeather := domain.Weather{}

	if len(resp.CurrentCondition) > 0 {
		formattedWeather.TempC = resp.CurrentCondition[0].TempC
		formattedWeather.Humidity = resp.CurrentCondition[0].Humidity
	} else {
		fmt.Errorf("Dados de clima não encontrados para %s", city)
	}

	return formattedWeather
}

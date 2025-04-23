package services

import (
	"fmt"

	domain "github.com/lourencogabe/gonimbus/domain/entities"
)

// método get realizara a consulta
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

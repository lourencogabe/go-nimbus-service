package services

import (
	"fmt"

	domain "github.com/lourencogabe/gonimbus/domain/entities"
	"github.com/lourencogabe/gonimbus/infra/handlers/external"
)

// método get realizara a consulta
func WeatherDataFormating(city string) (*domain.Weather, error) {
	resp, err := external.FetchWeatherData(city)
	if err != nil {
		return nil, fmt.Errorf("falha ao solicitar dados: %w", err)
	}

	wthr := domain.Weather{}
	if len(resp.CurrentCondition) > 0 {
		wthr.TempC = resp.CurrentCondition[0].TempC
		wthr.Humidity = resp.CurrentCondition[0].Humidity
	} else {
		return nil, fmt.Errorf("falha ao formatar dados de clima para: %s", city)
	}

	return &wthr, nil
}

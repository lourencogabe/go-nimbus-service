package usecases

import (
	"fmt"

	domain "github.com/lourencogabe/gonimbus/domain/entities"
	"github.com/lourencogabe/gonimbus/domain/services"
)

func GetWeatherData(city string) (*domain.WthrJson, error) {
	wthr, err := services.WeatherDataFormating(city)
	if err != nil {
		return nil, fmt.Errorf("falha ao realizar requisição para dados de clima: %w", err)
	}
	loc, err := services.MapDataFormating(city)
	if err != nil {
		return nil, fmt.Errorf("falha ao realizar requisição para dados do local: %w", err)
	}

	//retorna ponteiro para uma instância preenchida do tipo domain.WthrJson.
	return &domain.WthrJson{
		TempC:      wthr.TempC,
		Humidity:   wthr.Humidity,
		City:       loc.City,
		Population: loc.Population,
		Country:    loc.Country,
	}, nil
}

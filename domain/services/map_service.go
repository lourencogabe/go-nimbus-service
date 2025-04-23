package services

import (
	"fmt"

	domain "github.com/lourencogabe/gonimbus/domain/entities"
	"github.com/lourencogabe/gonimbus/infra/handlers/external"
)

func MapDataFormating(city string) (*domain.Map, error) {
	resp, err := external.FetchWeatherData(city)

	if err != nil {
		return nil, fmt.Errorf("falha ao chamar ao solicitar dados: %w", err)
	}

	loc := domain.Map{}
	if len(resp.NearestArea) > 0 {
		loc.City = resp.NearestArea[0].AreaName[0].Value
		loc.Country = resp.NearestArea[0].Country[0].Value
		loc.Population = resp.NearestArea[0].Population
	} else {
		return nil, fmt.Errorf("falha ao formatar dados para: %s", city)
	}

	return &loc, nil
}

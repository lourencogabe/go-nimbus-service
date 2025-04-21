package services

import (
	"fmt"

	domain "github.com/lourencogabe/gonimbus/domain/entities"
)

func MapDataFormating(city string) domain.Map {
	resp := GetWeatherData(city)

	loc := domain.Map{}

	if len(resp.NearestArea) > 0 {
		loc.City = resp.NearestArea[0].AreaName[0].Value
		loc.Country = resp.NearestArea[0].Country[0].Value
		loc.Population = resp.NearestArea[0].Population
	} else {
		fmt.Errorf("Dados de localização não encontrados para: %s", loc.City)
	}

	return loc
}

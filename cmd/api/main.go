package main

import (
	"fmt"

	"github.com/lourencogabe/gonimbus/domain/services"
)

func main() {
	wheatherData := services.WeatherDataFormating("curitiba")
	locData, err := services.MapDataFormating("curitiba")

	//test := external.FetchWeatherData("curitiba")

	fmt.Printf("Dados: %+v\n", wheatherData)
	fmt.Printf("Mapa: %+v, %+v\n", locData, err)
}

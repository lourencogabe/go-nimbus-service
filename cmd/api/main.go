package main

import (
	"fmt"

	"github.com/lourencogabe/gonimbus/domain/services"
)

func main() {
	wheatherData := services.WeatherDataFormating("curitiba")
	locData := services.MapDataFormating("curitiba")

	fmt.Printf("Dados: %+v\n", wheatherData)
	fmt.Printf("Mapa: %+v\n", locData)
}

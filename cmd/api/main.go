package main

import (
	"fmt"

	"github.com/lourencogabe/gonimbus/domain/services"
)

func main() {
	wheatherData := services.GetWeatherData("curitiba")

	fmt.Printf("Dados: %+v\n", wheatherData)
}

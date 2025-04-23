package main

import (
	"fmt"

	"github.com/lourencogabe/gonimbus/application/usecases"
)

func main() {
	data, err := usecases.GetWeatherData("curitiba")
	if err != nil {
		fmt.Errorf("Erro ao executar: %w", err)
	}

	//test := external.FetchWeatherData("curitiba")

	fmt.Printf("Dados: %+v\n", data)
}

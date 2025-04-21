package main

import (
	"fmt"

	"github.com/lourencogabe/gonimbus/domain/services"
)

func main() {
	wheatherData := services.WheatherDataFormating("0")

	fmt.Printf("Dados: %+v\n", wheatherData)
}

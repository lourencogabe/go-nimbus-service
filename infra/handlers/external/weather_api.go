package external

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/lourencogabe/gonimbus/dto"
)

func FetchWeatherData(city string) (*dto.GetWeather, error) {
	url := fmt.Sprintf("https://wttr.in/%s?format=j1", city)
	resp, err := http.Get(url)

	if err != nil {
		return nil, fmt.Errorf("falha ao realizar requisição: %s", err)
	}
	//realiza o fechamento do corpo da resposta para evitar vazamentos de dados
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("api indisponivel: %v", resp.StatusCode)
	}

	data := dto.GetWeather{}
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return nil, fmt.Errorf("falha ao decodificar a respostam em JSON: %w", err)
	}

	return &data, nil
}

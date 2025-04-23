package domain

type WthrJson struct {
	TempC      string `json:"temp_c"`
	Humidity   string `json:"humidity"`
	City       string `json:"city"`
	Country    string `json:"country"`
	Population string `json:"population"`
}

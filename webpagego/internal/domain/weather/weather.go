package weather

type WeatherRequest struct {
	CityID     string `json:"id"`
	CityName   string `json:"q"`
	Token      string `json:"appid"`
	PersonName string `json:"personname"`
}

type WeatherResponse struct {
	Temperature float64   `json:"temp"`
	Code        int       `json:"cod"`
	Ratio       CityRatio `json:"ratio"`
	Description string
	//TODO FINISH
}

type CityRatio struct {
	CityNumber int
	Number     int
}

package openweather

type WeatherResponse struct {
	Main main    `json:"main"`
	Code int     `json:"cod"`
	W    Weather `json:"weather"`
	//TODO finish!
}

type main struct {
	Temp float64 `json:"temp"`
}
type Weather []Detail

type Detail struct {
	Description string `json:"description"`
	Id          int    `json:"id"`
	//etc
}

package nameday

type NamedayResponse struct {
	Data Data `json:"data"`
}
type Data struct {
	Dates    Dates    `json:"dates"`
	Namedays Langlist `json:"namedays"`
}
type Dates struct {
	Day   int `json:"day"`
	Month int `json:"month"`
}

type Langlist struct {
	//TODO other languages!
	Hungarian string `json:"hu"`
}

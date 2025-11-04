package datasource

type GameDS struct {
	Matrix [3][3]int `json:"matrix"`
	Turn   int       `json:"turn"`
	Winner int       `json:"winner"`
}

package game_post

type GameWebResponse struct {
	Board  [3][3]int `json:"board"`
	Winner int       `json:"winner"`
}
type GameWebRequest struct {
	Board [3][3]int `json:"board"`
}

type ErrorResponse struct {
	Error string `json:"error"`
}

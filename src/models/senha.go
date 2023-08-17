package models

type Senha struct {
	Nova  string `json: "nova"`
	Atual string `json: "atual"`
}

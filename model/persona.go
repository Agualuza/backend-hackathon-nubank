package model

type Persona struct {
	Id          int     `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Goal        string  `json:"goal"`
	Photo       string  `json:"photo"`
	Factor      float64 `json:"factor"`
}

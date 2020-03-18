package model

type Persona struct {
	Id          int     `json:"id"`
	Name        string  `json:"name"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Goal        string  `json:"goal"`
	Photo       string  `json:"photo"`
	Factor      float64 `json:"factor"`
	Payment     float64 `json:"payment"`
	Bill        float64 `json:"bill"`
}

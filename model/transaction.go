package model

type Transaction struct {
	Id       int `json:"id"`
	User     User `json:"user"`
	Type    string `json:"type"`
	Amount float64 `json:"amount"`
	CreatedAt    string `json:"created_at"`
}

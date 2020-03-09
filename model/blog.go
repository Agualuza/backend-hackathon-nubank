package model

type Blog struct {
	Id     int     `json:"id"`
	Title  string  `json:"title"`
	Post   string  `json:"post"`
	Author   string  `json:"author"`
	CreatedAt  string  `json:"created_at"`
}

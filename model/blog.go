package model

type Blog struct {
	Id        int      `json:"id"`
	Title     string   `json:"title"`
	Post      string   `json:"post"`
	Author    string   `json:"author"`
	ReadTime  ReadTime `json:"read_time"`
	CreatedAt string   `json:"created_at"`
}

type ReadTime struct {
	Minutes int    `json:"minutes"`
	Time    string `json:"time"`
}

package endpoint

import (
	"bank/database"
	"bank/model"
	"encoding/json"
	"github.com/labstack/echo"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func LoadPosts(c echo.Context) error {
	var response jsonReponse

	response.Message = MessageSuccess
	response.Status = StatusOk
	response.Response = append(response.Response, getAllPosts())

	c.Response().Header().Set("Access-Control-Allow-Origin","*")
	c.Response().Header().Set(echo.HeaderContentType,echo.MIMEApplicationJSONCharsetUTF8)
	c.Response().WriteHeader(http.StatusBadRequest)
	return json.NewEncoder(c.Response()).Encode(response)
}


func getAllPosts() []model.Blog{
	db := database.ConnectDB()
	var posts []model.Blog

	rows, _ := db.Query("SELECT id,title,post,author,read_time,created_at FROM blog "+
		"ORDER BY created_at DESC")
	defer rows.Close()

	for rows.Next() {
		var b model.Blog
		var t time.Time
		var rt int
		rows.Scan(&b.Id, &b.Title, &b.Post, &b.Author,&rt, &t)
		sDate := strings.Split(t.String(), "-")
		sDay := strings.Split(sDate[2], " ")
		day := sDay[0]
		month := sDate[1]
		year := sDate[0]

		date := day + "/" + month + "/" + year

		b.CreatedAt = date
		rtString := strconv.Itoa(rt) + " minutos"

		var readTime model.ReadTime
		readTime.Minutes = rt
		readTime.Time = rtString

		b.ReadTime = readTime
		posts = append(posts,b)
	}
	defer db.Close()
	return posts
}
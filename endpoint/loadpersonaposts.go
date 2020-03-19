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

func LoadPersonaPosts(c echo.Context) error {
	if len(c.FormValue("persona_id")) == 0 {
		var response jsonReponse
		response.Status = StatusNok
		response.Message = "persona_id is required"
		c.Response().Header().Set("Access-Control-Allow-Origin","*")
		c.Response().Header().Set(echo.HeaderContentType,echo.MIMEApplicationJSONCharsetUTF8)
		c.Response().WriteHeader(http.StatusBadRequest)
		return json.NewEncoder(c.Response()).Encode(response)
	}

	pid, _ := strconv.Atoi(c.FormValue("persona_id"))
	var response jsonReponse

	response.Message = MessageSuccess
	response.Status = StatusOk
	response.Response = append(response.Response, getPostsByPersonaId(pid))

	c.Response().Header().Set("Access-Control-Allow-Origin","*")
	c.Response().Header().Set(echo.HeaderContentType,echo.MIMEApplicationJSONCharsetUTF8)
	c.Response().WriteHeader(http.StatusBadRequest)
	return json.NewEncoder(c.Response()).Encode(response)
}


func getPostsByPersonaId(pid int) []model.Blog{
	db := database.ConnectDB()
	var posts []model.Blog

	rows, _ := db.Query("SELECT b.id,b.title,b.post,b.author,b.read_time,b.created_at FROM blog b INNER JOIN response_blog rb on b.id = rb.blog_id WHERE rb.persona_id = ? "+
		"GROUP BY b.id ORDER BY b.created_at DESC",pid)

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

		b.ReadTime.Minutes = rt
		b.ReadTime.Time = rtString
		posts = append(posts,b)
	}
	defer db.Close()
	return posts
}
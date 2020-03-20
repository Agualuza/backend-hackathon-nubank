package endpoint

import (
	"bank/database"
	"encoding/json"
	"github.com/labstack/echo"
	"net/http"
)


func LoadQuestions(c echo.Context) error {
	var response jsonReponse

	response.Message = MessageSuccess
	response.Status = StatusOk
	response.Response = append(response.Response, getAllQuestions())

	c.Response().Header().Set("Access-Control-Allow-Origin","*")
	c.Response().Header().Set(echo.HeaderContentType,echo.MIMEApplicationJSONCharsetUTF8)
	c.Response().WriteHeader(http.StatusBadRequest)
	return json.NewEncoder(c.Response()).Encode(response)
}

func getAllQuestions() []string{
	db := database.ConnectDB()
	rows, _ := db.Query("SELECT question from question")
	defer rows.Close()
	var q string
	var list []string

	for rows.Next() {
		rows.Scan(&q);
		list = append(list,q)
	}

	defer db.Close()
	return list
}




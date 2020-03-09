package endpoint

import (
	"bank/database"
	"github.com/labstack/echo"
	"net/http"
)


func LoadQuestions(c echo.Context) error {
	var response jsonReponse

	response.Message = MessageSuccess
	response.Status = StatusOk
	response.Response = append(response.Response, getAllQuestions())

	return c.JSON(http.StatusOK, response)
}

func getAllQuestions() []string{
	db := database.ConnectDB()
	rows, _ := db.Query("SELECT question from question")
	var q string
	var list []string

	for rows.Next() {
		rows.Scan(&q);
		list = append(list,q)
	}

	defer db.Close()
	return list
}




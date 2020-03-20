package endpoint

import (
	"bank/database"
	"bank/model"
	"encoding/json"
	"github.com/labstack/echo"
	"net/http"
)


func LoadCategories(c echo.Context) error {
	var response jsonReponse

	response.Message = MessageSuccess
	response.Status = StatusOk
	response.Response = append(response.Response, getAllCategories())

	c.Response().Header().Set("Access-Control-Allow-Origin","*")
	c.Response().Header().Set(echo.HeaderContentType,echo.MIMEApplicationJSONCharsetUTF8)
	c.Response().WriteHeader(http.StatusBadRequest)
	return json.NewEncoder(c.Response()).Encode(response)
}

func getAllCategories() []model.Category{
	db := database.ConnectDB()
	rows, _ := db.Query("SELECT id, name, type from category")
	defer rows.Close()
	var id int
	var name,t string
	var list []model.Category

	for rows.Next() {
		rows.Scan(&id, &name, &t);
		var c model.Category
		c.Id = id
		c.Name = name
		c.Type = t
		list = append(list,c)
	}

	defer db.Close()
	return list
}



package endpoint

import (
	"bank/database"
	"bank/model"
	"github.com/labstack/echo"
	"net/http"
)


func LoadCategories(c echo.Context) error {
	var response jsonReponse

	response.Message = MessageSuccess
	response.Status = StatusOk
	response.Response = append(response.Response, getAllCategories())

	return c.JSON(http.StatusOK, response)
}

func getAllCategories() []model.Category{
	db := database.ConnectDB()
	rows, _ := db.Query("SELECT id, name, type from category")
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



package endpoint

import (
	"bank/database"
	"bank/model"
	"github.com/labstack/echo"
	"net/http"
)


func LoadPersonas(c echo.Context) error {
	var response jsonReponse

	response.Message = MessageSuccess
	response.Status = StatusOk
	response.Response = append(response.Response, getAllPersonas())

	return c.JSON(http.StatusOK, response)
}

func getAllPersonas() []model.Persona{
	db := database.ConnectDB()
	rows, _ := db.Query("SELECT id, name, description,goal,factor,photo from persona")
	var id int
	var name,description,goal,photo string
	var factor float64
	var list []model.Persona

	for rows.Next() {
		rows.Scan(&id, &name, &description, &goal,&factor, &photo);
			var p model.Persona
			p.Id = id
			p.Name = name
			p.Description = description
			p.Goal = goal
			p.Photo = photo
			p.Factor = factor
			list = append(list,p)
	}

	defer db.Close()
	return list
}



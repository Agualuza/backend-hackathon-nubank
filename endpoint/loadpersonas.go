package endpoint

import (
	"bank/database"
	"bank/model"
	"encoding/json"
	"github.com/labstack/echo"
	"net/http"
)


func LoadPersonas(c echo.Context) error {
	var response jsonReponse

	response.Message = MessageSuccess
	response.Status = StatusOk
	response.Response = append(response.Response, getAllPersonas())

	c.Response().Header().Set("Access-Control-Allow-Origin","*")
	c.Response().Header().Set(echo.HeaderContentType,echo.MIMEApplicationJSONCharsetUTF8)
	c.Response().WriteHeader(http.StatusBadRequest)
	return json.NewEncoder(c.Response()).Encode(response)
}

func getAllPersonas() []model.Persona{
	db := database.ConnectDB()
	rows, _ := db.Query("SELECT id, name,title, description,goal,factor,payment,bill,photo from persona")
	defer rows.Close()
	var id int
	var name,title,description,goal,photo string
	var factor,payment,bill float64
	var list []model.Persona

	for rows.Next() {
		rows.Scan(&id, &name, &title, &description, &goal,&factor,&payment,&bill,&photo);
			var p model.Persona
			p.Id = id
			p.Name = name
			p.Description = description
			p.Goal = goal
			p.Photo = photo
			p.Factor = factor
			p.Payment = payment
			p.Bill = bill
			p.Title = title
			list = append(list,p)
	}

	defer db.Close()
	return list
}



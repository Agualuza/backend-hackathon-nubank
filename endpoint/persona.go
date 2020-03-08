package endpoint

import (
	"bank/database"
	"bank/model"
	"github.com/labstack/echo"
	"net/http"
	"strconv"
)

func Persona(c echo.Context) error {
	db := database.ConnectDB()
	var id,count int
	var name,description,goal,photo string

	currentUser := getUserByToken(c.FormValue("token"))

	_ = db.QueryRow("SELECT count(1),id,name,description,goal,photo FROM persona WHERE q1 = ? and q2 = ? and q3 = ? and q4 = ? and q5 = ?",c.FormValue("q1"),c.FormValue("q2"),c.FormValue("q3"),c.FormValue("q4"),c.FormValue("q5")).Scan(&count,&id,&name,&description,&goal,&photo)

	var response jsonReponse
	var persona model.Persona

	persona.Id = id
	persona.Name = name
	persona.Description = description
	persona.Goal = goal
	persona.Photo = photo

	if count == 0 {
		p := returnEstimatedPersona(c.FormValue("q1"),c.FormValue("q2"),c.FormValue("q4"),c.FormValue("q5"))
		persona = p
	}

	savePersonaHistoric(currentUser.Id,persona.Id,c.FormValue("q1"),c.FormValue("q2"),c.FormValue("q3"),c.FormValue("q4"),c.FormValue("q5"))


	response.Message = MessageSuccess
	response.Status = StatusOk
	response.Response = append(response.Response, persona)

	return c.JSON(http.StatusOK, response)
}

func savePersonaHistoric(uid,pid int ,q1,q2,q3,q4,q5 string) {
	db := database.ConnectDB()
	_ , _ = db.Query("INSERT INTO answer (user_id,persona_id,q1,q2,q3,q4,q5) VALUES (?,?,?,?,?,?,?)",uid,pid,q1,q2,q3,q4,q5)
}

func returnEstimatedPersona(q1,q2,q4,q5 string) model.Persona {
	var p model.Persona
	var sum,pid int
	db := database.ConnectDB()

	q1i, _ := strconv.Atoi(q1)
	q2i, _ := strconv.Atoi(q2)
	q4i, _ := strconv.Atoi(q4)
	q5i, _ := strconv.Atoi(q5)

	sum = sumTwoAnswers(q1i,q2i) - sumTwoAnswers(q4i,q5i)

	if sum > 0 {
		pid = 2
	} else if sum < 0{
		pid = 4
	}  else {
		pid = 3
	}

	var name,description,goal,photo string

	_ = db.QueryRow("SELECT name,description,goal,photo from  persona where id = ?",pid).Scan(&name,&description,&goal,&photo)

	p.Id = pid
	p.Photo = photo
	p.Goal = goal
	p.Name = name
	p.Description = description

	return p


}

func sumTwoAnswers(q1,q2 int) int {
	return q1+q2
}

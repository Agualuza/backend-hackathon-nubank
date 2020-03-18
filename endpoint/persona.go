package endpoint

import (
	"bank/database"
	"bank/model"
	"encoding/json"
	"github.com/labstack/echo"
	"net/http"
	"strconv"
)

func Persona(c echo.Context) error {
	db := database.ConnectDB()
	var id,count int
	var name,title,description,goal,photo string
	var payment,bill float64

	currentUser := getUserByToken(c.FormValue("token"))

	_ = db.QueryRow("SELECT count(1),id,name,title,description,goal,payment,bill,photo FROM persona WHERE q1 = ? and q2 = ? and q3 = ? and q4 = ? and q5 = ?",c.FormValue("q1"),c.FormValue("q2"),c.FormValue("q3"),c.FormValue("q4"),c.FormValue("q5")).Scan(&count,&id,&name,&title,&description,&goal,&payment,&bill,&photo)

	rows , _ := db.Query("SELECT id,name,title,description,goal,payment,bill,photo FROM persona WHERE id = ? or id = ?",id-1,id+1)
	defer rows.Close()

	var response jsonReponse
	var persona []model.Persona
	var mainPid int

	if count == 0 {
		p,pid := returnEstimatedPersona(c.FormValue("q1"),c.FormValue("q2"),c.FormValue("q4"),c.FormValue("q5"))
		persona = p
		mainPid = pid
	} else {
		var p model.Persona
		p.Id = id
		p.Name = name
		p.Description = description
		p.Goal = goal
		p.Payment = payment
		p.Bill = bill
		p.Photo = photo
		p.Title = title
		mainPid = id

		persona = append(persona,p)

		for rows.Next() {
			var pAux model.Persona
			rows.Scan(&pAux.Id,&pAux.Name,&pAux.Title,&pAux.Description,&pAux.Goal,&pAux.Payment,&pAux.Bill,&pAux.Photo)
			persona = append(persona,pAux)
		}
	}

	savePersonaHistoric(currentUser.Id,mainPid,c.FormValue("q1"),c.FormValue("q2"),c.FormValue("q3"),c.FormValue("q4"),c.FormValue("q5"))



	response.Message = MessageSuccess
	response.Status = StatusOk
	response.Response = append(response.Response, persona)

	defer db.Close()

	c.Response().Header().Set("Access-Control-Allow-Origin","*")
	c.Response().Header().Set(echo.HeaderContentType,echo.MIMEApplicationJSONCharsetUTF8)
	c.Response().WriteHeader(http.StatusBadRequest)
	return json.NewEncoder(c.Response()).Encode(response)
}

func savePersonaHistoric(uid,pid int ,q1,q2,q3,q4,q5 string) {
	db := database.ConnectDB()
	_ , _ = db.Query("INSERT INTO answer (user_id,persona_id,q1,q2,q3,q4,q5) VALUES (?,?,?,?,?,?,?)",uid,pid,q1,q2,q3,q4,q5)
	defer db.Close()
}

func returnEstimatedPersona(q1,q2,q4,q5 string) ([]model.Persona,int) {
	var list []model.Persona
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

	var name,title,description,goal,photo string
	var factor,payment,bill float64
	var id,count int

	_ = db.QueryRow("SELECT count(1) FROM persona").Scan(&count)

	id0,id1,id2 := getIdsList(pid,count)

	rows , _ := db.Query("SELECT id,name,title,description,goal,factor,payment,bill,photo FROM  persona WHERE id IN (?,?,?) ORDER BY id != ?",id0,id1,id2,id0)

	defer rows.Close()

	for rows.Next() {
		var p model.Persona
		rows.Scan(&id,&name,&title,&description,&goal,&factor,&payment,&bill,&photo)
		p.Id = id
		p.Photo = photo
		p.Goal = goal
		p.Name = name
		p.Description = description
		p.Factor = factor
		p.Payment = payment
		p.Bill = bill
		p.Title = title
		list = append(list,p)
	}

	defer db.Close()
	return list,pid


}

func getIdsList(pid,count int) (int,int,int) {
	if pid > 1 && pid < count{
		return pid,pid-1,pid+1
	} else if pid == 0 {
		return pid,pid+1,0
	} else {
		return pid,pid-1,0
	}
}

func sumTwoAnswers(q1,q2 int) int {
	return q1+q2
}

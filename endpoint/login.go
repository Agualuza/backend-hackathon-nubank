package endpoint

import (
	"bank/database"
	"bank/model"
	"encoding/json"
	"github.com/labstack/echo"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

func Login(c echo.Context) error {
	if len(c.FormValue("token")) > 0 {
		currentUser := getUserByToken(c.FormValue("token"))
		var response jsonReponse
		response.Status = StatusOk
		response.Response = append(response.Response, currentUser)
		response.Message = MessageSuccess

		if currentUser.Id == 0 {
			response.Status = StatusNok
			response.Message = "Token invalid"
		}

		c.Response().Header().Set("Access-Control-Allow-Origin","*")
		c.Response().Header().Set(echo.HeaderContentType,echo.MIMEApplicationJSONCharsetUTF8)
		c.Response().WriteHeader(http.StatusBadRequest)
		return json.NewEncoder(c.Response()).Encode(response)
	}

	currentUser, exists := getCurrentUser(c.FormValue("email"), c.FormValue("password"))

	message := "User/Password invalid"
	status := StatusNok

	if exists {
		message = MessageSuccess
		status = StatusOk
	}

	var response jsonReponse
	response.Status = status
	response.Response = append(response.Response, currentUser)
	response.Message = message

	c.Response().Header().Set("Access-Control-Allow-Origin","*")
	c.Response().Header().Set(echo.HeaderContentType,echo.MIMEApplicationJSONCharsetUTF8)
	c.Response().WriteHeader(http.StatusBadRequest)
	return json.NewEncoder(c.Response()).Encode(response)

}

func getCurrentUser(e, p string) (model.User, bool) {
	db := database.ConnectDB()
	currentUser := model.User{}

	rows, _ := db.Query("SELECT count(1),id, name, email,password,balance FROM user WHERE email = ?", e)

	defer rows.Close()

	var count, id int
	var name, email, password string
	var balance float64

	for rows.Next() {
		if err := rows.Scan(&count, &id, &name, &email, &password, &balance); err != nil {
			return currentUser, false
		}
	}

	if count > 0 && bcrypt.CompareHashAndPassword([]byte(password), []byte(p)) == nil {
		currentUser.Id = id
		currentUser.Name = name
		currentUser.Password = "******"
		currentUser.Email = email
		currentUser.Balance = balance
		var token string
		_, _ = db.Query("UPDATE user SET token = MD5((RAND()+RAND()+RAND()+RAND())*NOW())")
		_ = db.QueryRow("SELECT token FROM user WHERE id = ?", currentUser.Id).Scan(&token)
		currentUser.Token = token
		return currentUser, true
	}

	defer db.Close()
	return currentUser, false
}

func getUserByToken(t string) model.User {
	db := database.ConnectDB()
	currentUser := model.User{}
	var count, id int
	var name, email, password, token string
	var balance float64

	 _ = db.QueryRow("SELECT count(1),id, name, email,password,token,balance FROM user WHERE token = ?", t).Scan(&count, &id, &name, &email, &password, &token,&balance)

	if count > 0 {
		currentUser.Email = email
		currentUser.Id = id
		currentUser.Password = "******"
		currentUser.Name = name
		currentUser.Balance = balance
		currentUser.Token = token
	}

	return currentUser

}

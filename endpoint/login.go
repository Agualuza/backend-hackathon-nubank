package endpoint

import (
	"bank/database"
	"bank/model"
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

		return c.JSON(http.StatusOK, response)
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

	return c.JSON(http.StatusOK, response)

}

func getCurrentUser(e, p string) (model.User, bool) {
	db := database.ConnectDB()
	currentUser := model.User{}

	rows, _ := db.Query("SELECT count(1),id, name, email,password FROM user WHERE email = ?", e)

	defer rows.Close()

	var count, id int
	var name, email, password string

	for rows.Next() {
		if err := rows.Scan(&count, &id, &name, &email, &password); err != nil {
			return currentUser, false
		}
	}

	if count > 0 && bcrypt.CompareHashAndPassword([]byte(password), []byte(p)) == nil {
		currentUser.Id = id
		currentUser.Name = name
		currentUser.Password = "******"
		currentUser.Email = email
		var token string
		_, _ = db.Query("UPDATE user SET token = MD5((RAND()+RAND()+RAND()+RAND())*NOW())")
		_ = db.QueryRow("SELECT token FROM user WHERE id = ?", currentUser.Id).Scan(&token)
		currentUser.Token = token
		return currentUser, true
	}

	return currentUser, false
}

func getUserByToken(t string) model.User {
	db := database.ConnectDB()
	currentUser := model.User{}
	var count, id int
	var name, email, password, token string

	 _ = db.QueryRow("SELECT count(1),id, name, email,password,token FROM user WHERE token = ?", t).Scan(&count, &id, &name, &email, &password, &token)

	if count > 0 {
		currentUser.Email = email
		currentUser.Id = id
		currentUser.Password = "******"
		currentUser.Name = name
		currentUser.Token = token
	}

	return currentUser

}
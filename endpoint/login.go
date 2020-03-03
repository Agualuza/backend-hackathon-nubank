package endpoint

import (
	"bank/database"
	"bank/model"
	"github.com/labstack/echo"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

func Login(c echo.Context) error {
	currentUser, exists := getCurrentUser(c.FormValue("email"),c.FormValue("password"))

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

	var count,id int
	var name,email,password string

	for rows.Next() {
		if err := rows.Scan(&count,&id,&name,&email,&password); err != nil {
			return currentUser,false
		}
	}

	if count > 0 && bcrypt.CompareHashAndPassword([]byte(password),[]byte(p)) == nil {
		currentUser.Id = id
		currentUser.Name = name
		currentUser.Password = "******"
		currentUser.Email = email
		var token string
		_, _ = db.Query("UPDATE user SET token = MD5((RAND()+RAND()+RAND()+RAND())*NOW())")
		_ = db.QueryRow("SELECT token FROM user WHERE id = ?",currentUser.Id).Scan(&token)
		currentUser.Token = token
		return currentUser,true
	}

	return currentUser,false
}

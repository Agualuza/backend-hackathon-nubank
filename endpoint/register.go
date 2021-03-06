package endpoint

import (
	"bank/database"
	"encoding/json"
	"github.com/cheekybits/genny/generic"
	"github.com/labstack/echo"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

func Register(c echo.Context) error {
	var response jsonReponse

	can := verify(c.FormValue("email"))

	if can {
		name := c.FormValue("name")
		password := c.FormValue("password")
		email := c.FormValue("email")
		bcryptPass, _ := bcrypt.GenerateFromPassword([]byte(password), HashCost)

		db := database.ConnectDB()

		rows , err := db.Query("INSERT INTO user (name,email,password) VALUES (?,?,?)", name, email, bcryptPass)

		defer rows.Close()

		response.Status = StatusNok
		response.Message = "Registration has not been possible"

		if err == nil {
			response.Status = StatusOk
			response.Message = MessageSuccess
		}

		defer db.Close()
		return c.JSON(http.StatusOK, response)
	}

	var emptySlice []generic.Type

	response.Status = StatusNok
	response.Response = emptySlice
	response.Message = "You are already registered"

	c.Response().Header().Set("Access-Control-Allow-Origin","*")
	c.Response().Header().Set(echo.HeaderContentType,echo.MIMEApplicationJSONCharsetUTF8)
	c.Response().WriteHeader(http.StatusBadRequest)
	return json.NewEncoder(c.Response()).Encode(response)
}

func verify(email string) bool {
	db := database.ConnectDB()

	rows, _ := db.Query("SELECT COUNT(*) FROM user WHERE email = ?", email)


	defer rows.Close()

	var count int

	for rows.Next() {
		if err := rows.Scan(&count); err != nil {
			return false
		}
	}

	if count > 0 {
		return false
	}

	defer db.Close()
	return true
}

package endpoint

import (
	"bank/database"
	"github.com/labstack/echo"
	"net/http"
)

func MakeTransaction(c echo.Context) error {
	db := database.ConnectDB()
	tp := c.FormValue("type") // D - Debit; C - Credit
	token := c.FormValue("token")
	amount := c.FormValue("amount")

	currentUser := getUserByToken(token)

	_ , err := db.Query("INSERT INTO transaction (user_id,type,amount) VALUES (?,?,?)",currentUser.Id,tp,amount)

	var response jsonReponse
	response.Status = StatusOk
	response.Message = MessageSuccess

	if err != nil {
		response.Status = StatusNok
		response.Message = "An error has occurred. Transaction not completed. Try again later"
	}

	defer db.Close()

	c.Response().Header().Set("Access-Control-Allow-Origin","*")
	c.Response().WriteHeader(http.StatusOK)
	return c.JSON(http.StatusOK, response)
}
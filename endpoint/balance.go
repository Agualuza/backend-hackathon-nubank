package endpoint

import (
	"github.com/labstack/echo"
	"net/http"
)


func Balance(c echo.Context) error {

	currentUser := getUserByToken(c.FormValue("token"))

	var response jsonReponse
	response.Message = MessageSuccess
	response.Status = StatusOk
	response.Response = append(response.Response, currentUser.Balance)

	return c.JSON(http.StatusOK, response)
}

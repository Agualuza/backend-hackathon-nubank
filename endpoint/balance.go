package endpoint

import (
	"encoding/json"
	"github.com/labstack/echo"
	"net/http"
)


func Balance(c echo.Context) error {

	currentUser := getUserByToken(c.FormValue("token"))

	var response jsonReponse
	response.Message = MessageSuccess
	response.Status = StatusOk
	response.Response = append(response.Response, currentUser.Balance)

	c.Response().Header().Set("Access-Control-Allow-Origin","*")
	c.Response().Header().Set(echo.HeaderContentType,echo.MIMEApplicationJSONCharsetUTF8)
	c.Response().WriteHeader(http.StatusBadRequest)
	return json.NewEncoder(c.Response()).Encode(response)
}

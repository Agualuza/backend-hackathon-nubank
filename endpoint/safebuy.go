package endpoint

import (
	"bank/database"
	"github.com/labstack/echo"
	"math"
	"net/http"
	"strconv"
)

func SafeBuy(c echo.Context) error {

	Profiles := []Profile{
		{1, 2000.00},
		{2, 5000.00},
		{3, 10000.00},
	}

	pid, _ := strconv.ParseInt(c.FormValue("profile_id"), 10, 4)
	p := Profiles[pid-1]
	price, _ := strconv.ParseFloat(c.FormValue("product_price"),64)

	var varPayment float64
	var rate float64

	varPayment = p.Payment * SafeBuyVariablePercentage

	rate = varPayment / price

	var response RiskEvaluate
	response.Profile = p
	response.Rate = 1 - math.Min(rate,1)
	saveAnalytics(c.FormValue("token"),p,price,rate)
	return c.JSON(http.StatusOK, response)

}

func saveAnalytics(t string , profile Profile, pp,rt float64){
	currentUser := getUserByToken(t)
	db := database.ConnectDB()
	_ , err := db.Query("INSERT INTO risk_analytic (user_id,payment,product_price,rate) VALUES(?,?,?,?)",currentUser.Id,profile.Payment,pp,rt)

	if(err != nil) {
		return
	}
}
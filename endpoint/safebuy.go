package endpoint

import (
	"bank/database"
	"bank/model"
	"github.com/labstack/echo"
	"net/http"
	"strconv"
	"time"
)

func SafeBuy(c echo.Context) error {
	if len(c.FormValue("persona_id")) == 0 {
		var response jsonReponse
		response.Status = StatusNok
		response.Message = "persona_id is required"
		return c.JSON(http.StatusBadRequest, response)
	}

	if len(c.FormValue("product_price")) == 0 {
		var response jsonReponse
		response.Status = StatusNok
		response.Message = "product_price is required"
		return c.JSON(http.StatusBadRequest, response)
	}

	if len(c.FormValue("category_id")) == 0 {
		var response jsonReponse
		response.Status = StatusNok
		response.Message = "category_id is required"
		return c.JSON(http.StatusBadRequest, response)
	}

	if len(c.FormValue("token")) == 0 {
		var response jsonReponse
		response.Status = StatusNok
		response.Message = "token is required"
		return c.JSON(http.StatusBadRequest, response)
	}

	currentUser := getUserByToken(c.FormValue("token"))
	persona := getPersona(c.FormValue("persona_id"))
	category := getCategory(c.FormValue("category_id"))
	productPrice, _ := strconv.ParseFloat(c.FormValue("product_price"),64)

	evaluation := analyze(persona,category,currentUser, productPrice)

	var response jsonReponse
	var riskEvaluate RiskEvaluate
	riskEvaluate.Evaluation = evaluation
	response.Status = StatusOk
	response.Message = MessageSuccess
	response.Response = append(response.Response,evaluation)
	return c.JSON(http.StatusOK, response)

}

func analyze(p model.Persona, c model.Category, u model.User, price float64) Evaluation {
	evaluations := loadEvaluations()
	payment := getPayment(u)
	var response float64

	response = ( payment * SafeBuyVariablePercentage / price ) * p.Factor
	eval := calculateEvaluation(response)
	bestEvalAllowed := getBestEvalAllowed(eval,c,p)

	saveAnalytics(u,price,payment,eval)

	return evaluations[bestEvalAllowed]
}

func calculateEvaluation(r float64) string {
	if r >= 1.0 {
		return "GE"
	} else if r >= 0.9 {
		return "GM"
	} else if r >= 0.8 {
		return "GB"
	} else if r >= 0.7 {
		return "YE"
	} else if r >= 0.6 {
		return "YM"
	} else if r >= 0.5 {
		return "YB"
	} else if r >= 0.4 {
		return "RE"
	} else if r >= 0.3 {
		return "RM"
	} else{
		return "RB"
	}

}

func getBestEvalAllowed(response string, c model.Category, p model.Persona) string {
	switch p.Id {
	case 1:
		return response
	case 2:
		if c.Type == "L" && ( response == "GE" || response == "GM" ) {
			return "GB"
		}
	case 3:
		if (c.Type == "L" || c.Type == "F" ) && ( response == "GE" || response == "GM" || response == "GB" || response == "YE"  )  {
			return "YM"
		}
	case 4:
		if (c.Type != "P" ) && ( response == "GE" || response == "GM" || response == "GB" || response == "YE" || response == "YM" )  {
			return "YB"
		}
	case 5:
		if (c.Type != "P" ) && ( response == "GE" || response == "GM" || response == "GB" || response == "YE" || response == "YM" || response == "YB"  ) {
			return "RE"
		}
	}

	return response
}

func getPayment(user model.User) float64 {
	db := database.ConnectDB()
	var p float64
	day := time.Now().Day()
	month := time.Now().Month()
	year := time.Now().Year()

	if day > 20 {
		month--
	}

	_ = db.QueryRow("SELECT ((SELECT SUM(amount) FROM `transaction` WHERE MONTH(created_at) = ? "+
		"and YEAR(created_at) = ? and `type` = 'C' and user_id = ?) - (SELECT SUM(amount) FROM `transaction` "+
		"WHERE MONTH(created_at) = ? and YEAR(created_at) = ? and `type` = 'D' and user_id = ?))", month, year, user.Id, month, year, user.Id).Scan(&p)

	return p
}

func loadEvaluations() map[string]Evaluation {
	var m map[string]Evaluation
	m = make(map[string]Evaluation)

	m["RB"] = Evaluation{"RB", "Red Begin"}
	m["RM"] = Evaluation{"RM", "Red Mid"}
	m["RE"] = Evaluation{"RE", "Red End"}

	m["YB"] = Evaluation{"YB", "Yellow Begin"}
	m["YM"] = Evaluation{"YM", "Yellow Mid"}
	m["YE"] = Evaluation{"YE", "Yellow End"}

	m["GB"] = Evaluation{"GB", "Green Begin"}
	m["GM"] = Evaluation{"GM", "Green Mid"}
	m["GE"] = Evaluation{"GE", "Green End"}

	return m
}

func saveAnalytics(u model.User,price,payment float64, response string) {
	db := database.ConnectDB()
	_, err := db.Query("INSERT INTO risk_analytic (user_id,payment,product_price,response) VALUES(?,?,?,?)", u.Id, payment, price, response)

	if err != nil {
		return
	}
}

func getPersona(pid string) model.Persona {
	var p model.Persona
	db := database.ConnectDB()

	_ = db.QueryRow("SELECT id,name,description,goal,factor,photo FROM persona WHERE id = ?", pid).Scan(&p.Id, &p.Name, &p.Description, &p.Goal, &p.Factor,&p.Photo)

	return p
}

func getCategory(cid string) model.Category {
	var c model.Category
	db := database.ConnectDB()

	_ = db.QueryRow("SELECT id,name,type FROM category WHERE id = ?", cid).Scan(&c.Id, &c.Name, &c.Type)

	return c
}

package endpoint

import (
	"bank/database"
	"bank/model"
	"encoding/json"
	"github.com/labstack/echo"
	"math"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func SafeBuy(c echo.Context) error {
	if len(c.FormValue("persona_id")) == 0 {
		var response jsonReponse
		response.Status = StatusNok
		response.Message = "persona_id is required"
		c.Response().Header().Set("Access-Control-Allow-Origin","*")
		c.Response().Header().Set(echo.HeaderContentType,echo.MIMEApplicationJSONCharsetUTF8)
		c.Response().WriteHeader(http.StatusBadRequest)
		return json.NewEncoder(c.Response()).Encode(response)
	}

	if len(c.FormValue("product_price")) == 0 {
		var response jsonReponse
		response.Status = StatusNok
		response.Message = "product_price is required"
		c.Response().Header().Set("Access-Control-Allow-Origin","*")
		c.Response().Header().Set(echo.HeaderContentType,echo.MIMEApplicationJSONCharsetUTF8)
		c.Response().WriteHeader(http.StatusBadRequest)
		return json.NewEncoder(c.Response()).Encode(response)
	}

	if len(c.FormValue("category_id")) == 0 {
		var response jsonReponse
		response.Status = StatusNok
		response.Message = "category_id is required"
		c.Response().Header().Set("Access-Control-Allow-Origin","*")
		c.Response().Header().Set(echo.HeaderContentType,echo.MIMEApplicationJSONCharsetUTF8)
		c.Response().WriteHeader(http.StatusBadRequest)
		return json.NewEncoder(c.Response()).Encode(response)
	}

	if len(c.FormValue("token")) == 0 {
		var response jsonReponse
		response.Status = StatusNok
		response.Message = "token is required"
		c.Response().Header().Set("Access-Control-Allow-Origin","*")
		c.Response().Header().Set(echo.HeaderContentType,echo.MIMEApplicationJSONCharsetUTF8)
		c.Response().WriteHeader(http.StatusBadRequest)
		return json.NewEncoder(c.Response()).Encode(response)
	}

	currentUser := getUserByToken(c.FormValue("token"))
	persona := getPersona(c.FormValue("persona_id"))
	category := getCategory(c.FormValue("category_id"))
	productPrice, _ := strconv.ParseFloat(c.FormValue("product_price"), 64)

	evaluation := analyze(persona, category, currentUser, productPrice)

	blog := getBlogPost(persona.Id, evaluation.Evaluation)

	var response jsonReponse
	var riskEvaluate RiskEvaluate
	riskEvaluate.Evaluation = evaluation
	riskEvaluate.Blog = blog
	response.Status = StatusOk
	response.Message = MessageSuccess
	response.Response = append(response.Response, riskEvaluate)
	c.Response().Header().Set("Access-Control-Allow-Origin","*")
	c.Response().Header().Set(echo.HeaderContentType,echo.MIMEApplicationJSONCharsetUTF8)
	c.Response().WriteHeader(http.StatusBadRequest)
	return json.NewEncoder(c.Response()).Encode(response)

}

func analyze(p model.Persona, c model.Category, u model.User, price float64) Evaluation {
	evaluations := loadEvaluations()
	var response float64

	balance := (p.Payment * SafeBuyVariablePercentage) - p.Bill
	
	balance = math.Max(balance,SafeBuyMinBalance)
	
	response = (balance / price) * p.Factor
	eval := calculateEvaluation(response)
	bestEvalAllowed := getBestEvalAllowed(eval, c, p)

	saveAnalytics(u, price, p.Payment, eval)

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
	} else {
		return "RB"
	}

}

func getBestEvalAllowed(response string, c model.Category, p model.Persona) string {
	switch p.Id {
	case 1:
		return response
	case 2:
		if c.Type == "L" && (response == "GE" || response == "GM") {
			return "GB"
		}
	case 3:
		if (c.Type == "L" || c.Type == "F") && (response == "GE" || response == "GM" || response == "GB" || response == "YE") {
			return "YM"
		}
	case 4:
		if (c.Type != "P") && (response == "GE" || response == "GM" || response == "GB" || response == "YE" || response == "YM") {
			return "YB"
		}
	case 5:
		if (c.Type != "P") && (response == "GE" || response == "GM" || response == "GB" || response == "YE" || response == "YM" || response == "YB") {
			return "RE"
		}
	}

	return response
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

func saveAnalytics(u model.User, price, payment float64, response string) {
	db := database.ConnectDB()
	_, err := db.Query("INSERT INTO risk_analytic (user_id,payment,product_price,response) VALUES(?,?,?,?)", u.Id, payment, price, response)

	defer db.Close()

	if err != nil {
		return
	}
}

func getPersona(pid string) model.Persona {
	var p model.Persona
	db := database.ConnectDB()

	_ = db.QueryRow("SELECT id,name,description,goal,factor,payment,bill,photo FROM persona WHERE id = ?", pid).Scan(&p.Id, &p.Name, &p.Description, &p.Goal, &p.Factor, &p.Payment, &p.Bill, &p.Photo)

	defer db.Close()
	return p
}

func getCategory(cid string) model.Category {
	var c model.Category
	db := database.ConnectDB()

	_ = db.QueryRow("SELECT id,name,type FROM category WHERE id = ?", cid).Scan(&c.Id, &c.Name, &c.Type)

	defer db.Close()
	return c
}

func getBlogPost(pid int, r string) model.Blog {
	db := database.ConnectDB()
	var b model.Blog
	var t time.Time
	runes := []rune(r)

	_ = db.QueryRow("SELECT b.id,b.title,b.post,b.author,b.created_at FROM blog b "+
		"INNER JOIN response_blog rb ON rb.blog_id = b.id "+
		"WHERE rb.persona_id = ? AND rb.response = ?", pid, string(runes[0])).Scan(&b.Id, &b.Title, &b.Post, &b.Author, &t)

	sDate := strings.Split(t.String(), "-")
	sDay := strings.Split(sDate[2], " ")
	day := sDay[0]
	month := sDate[1]
	year := sDate[0]

	date := day + "/" + month + "/" + year

	b.CreatedAt = date

	defer db.Close()
	return b
}
